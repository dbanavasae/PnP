package main

import (
	"fmt"
	"log"
	"runtime"
	"io"
	"time"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"github.com/micro/go-grpc"
	"golang.org/x/net/context"
	"github.com/PnP/common"
	"github.com/PnP/executor"
	"github.com/micro/go-micro"
	"github.com/micro/cli"
	"github.com/micro/go-micro/transport"
	proto "github.com/PnP/pnp-proto"
)

func populateClientDetails() (clientInfo proto.ClientInfo) {
	archType := runtime.GOARCH
	osType := runtime.GOOS
	getOSFlavorCmd := "lsb_release -a | grep Description | awk -F':' '{print $2}'"

	osFlavor, err := executor.ExecuteCommand(getOSFlavorCmd)
	if err != nil {
		log.Fatalf("Error while getting OS type: %v", err)
	}
	// ToDo: Client ID generation...
	clientId := "client1"

	clientInfo = proto.ClientInfo{OsType: osType, ArchType: archType, OsFlavor: osFlavor, ClientId: clientId}
	return
}

func executeServerInstructions(cmdString []string) (exeErr error) {
	var errStr string
	for _, cmd := range cmdString {
		errStr, exeErr = executor.ExecuteCommand(cmd)
		if exeErr != nil {
			fmt.Printf("\nCommand <%v> failed to execute\nErrorString: %v\nError: %v\n", cmd, errStr, exeErr)
			break
		}
	}
	return exeErr
}

func setPkgMsgType(serverPkgOperType proto.PkgOperType, exeErr error) (clientPkgMsgType proto.ClientPkgMsgType) {

	switch serverPkgOperType {
	case proto.PkgOperType_IS_PKG_INSTALLED:
		{
			if exeErr == nil {
				clientPkgMsgType = proto.ClientPkgMsgType_PKG_INSTALLED
			} else {
				clientPkgMsgType = proto.ClientPkgMsgType_PKG_NOT_INSTALLED
			}
		}

	case proto.PkgOperType_INSTALL_PKG:
		fallthrough

	case proto.PkgOperType_INSTALL_PKG_FROM_FILE:
		{
			if exeErr == nil {
				clientPkgMsgType = proto.ClientPkgMsgType_PKG_INSTALL_SUCCESS
			} else {
				clientPkgMsgType = proto.ClientPkgMsgType_PKG_INSTALL_FAILED
				fmt.Printf("\nFailed to install package\n")
			}
		}

	case proto.PkgOperType_GET_NEXT_PKG:
		{
			clientPkgMsgType = proto.ClientPkgMsgType_PKG_INIT
		}
	}
	return
}

func installPackages(pnpClient proto.PnPService) {
	cxt, cancel := context.WithTimeout(context.Background(), time.Minute*20)
	defer cancel()
	stream, err := pnpClient.GetPackages(cxt)
	clientInfo := populateClientDetails()
	clientPkgMsgType := proto.ClientPkgMsgType_PKG_INIT

	clientMsg := &proto.ClientPkgMsg{CommonClientInfo: &proto.CommonClientInfo{RequestHeader:
		common.NewReqHdrGenerateTraceAndMessageID(), ClientInfo: &clientInfo},
			ClientPkgMsgType: clientPkgMsgType}
	serverPkgResp := &proto.ServerPkgResponse{}

	for {
		if err = stream.Send(clientMsg); err != nil {
			log.Fatalf("Failed to send client message, Error: %v", err)
		}

		serverPkgResp, err = stream.Recv()
		if err == io.EOF || serverPkgResp.CommonServerResponse.GetCmdType() == proto.CmdType_CLOSE_CONN {
			fmt.Println("\nClosing connection")
			stream.Close()
			break
		}

		if err != nil {
			fmt.Printf("Error while receiving data from server %v\n",  err)
		}

		var exeErr error
		if serverPkgResp.CommonServerResponse.GetCmdType() == proto.CmdType_RUN {
			fmt.Printf("\nCommand string: %v\n", serverPkgResp.InstructionPayload.Cmd)
			cmdStr := serverPkgResp.InstructionPayload.Cmd
			exeErr = executeServerInstructions(cmdStr)
		}

		clientPkgMsgType = setPkgMsgType(serverPkgResp.GetPkgOperType(), exeErr)

		traceId := serverPkgResp.CommonServerResponse.ResponseHeader.Identifiers.TraceID

		clientMsg = &proto.ClientPkgMsg{CommonClientInfo: &proto.CommonClientInfo{RequestHeader:
			common.NewReqHdrGenerateMessageID(traceId), ClientInfo: &clientInfo},
				ClientPkgMsgType: clientPkgMsgType }
	}
}

func setPlatformMsgType(serverPlatformOperType proto.SDPOperType, exeErr error) (clientPlatformMsgType proto.ClientPlatformMsgType) {

	switch serverPlatformOperType {
	case proto.SDPOperType_IS_SDP_PLATFORM_INSTALLED:
		{
			if exeErr != nil {
				clientPlatformMsgType = proto.ClientPlatformMsgType_SDP_PLATFORM_NOT_INSTALLED
			} else {
				clientPlatformMsgType = proto.ClientPlatformMsgType_SDP_PLATFORM_ALREADY_INSTALLED
			}
		}
	case proto.SDPOperType_DOWNLOAD_SDP_PLATFORM_ARTIFACT:
		{
			if exeErr != nil {
				clientPlatformMsgType = proto.ClientPlatformMsgType_SDP_PLATFORM_ARTIFACT_DOWNLOAD_FAILED
			} else {
				clientPlatformMsgType = proto.ClientPlatformMsgType_SDP_PLATFORM_ARTIFACT_DOWNLOAD_SUCCESS
			}
		}
	case proto.SDPOperType_DEPLOY_SDP_PLATFORM:
		{
			if exeErr != nil {
				clientPlatformMsgType = proto.ClientPlatformMsgType_SDP_PLATFORM_DEPLOYMENT_FAILED
			} else {
				clientPlatformMsgType = proto.ClientPlatformMsgType_SDP_PLATFORM_DEPLOYMENT_SUCCESS
			}
		}
	}
	return
}

func deployPlatform(pnpClient proto.PnPService, platformType string) {
	cxt, cancel := context.WithTimeout(context.Background(), time.Minute*120)
	defer cancel()
	stream, err := pnpClient.DeployPlatform(cxt)
	var clientPlatformMsgType proto.ClientPlatformMsgType

	clientInfo := populateClientDetails()

	if platformType == "master" {
		clientPlatformMsgType = proto.ClientPlatformMsgType_SDP_PLATFORM_MASTER_INIT
	} else if platformType == "satellite" {
		clientPlatformMsgType = proto.ClientPlatformMsgType_SDP_PLATFORM_SATELLITE_INIT
	}

	clientMsg := &proto.ClientPlatformMsg{CommonClientInfo: &proto.CommonClientInfo{RequestHeader:
		common.NewReqHdrGenerateTraceAndMessageID(), ClientInfo: &clientInfo}, ClientPlatformMsgType:
			clientPlatformMsgType}

	serverPlatformResponse := &proto.ServerPlatformResponse{}

	for {
		fmt.Printf("\nSending SDP Deploy request to PnP server, Request type: %v\n", clientPlatformMsgType)
		if err = stream.Send(clientMsg); err != nil {
			log.Fatalf("\nFailed to send client message, Error: %v", err)
		}

		serverPlatformResponse, err = stream.Recv()
		if err == io.EOF || serverPlatformResponse.CommonServerResponse.GetCmdType() == proto.CmdType_CLOSE_CONN {
			fmt.Println("\nClosing connection")
			stream.Close()
			break
		}

		if err != nil {
			fmt.Printf("\nError while receiving data from PnP server %v\n",  err)
		}
		fmt.Printf("\nReceived SDP Deploy instructions from PnP server, Instruction type: %v\n", serverPlatformResponse.GetSdpOperType())

		var exeErr error
		if serverPlatformResponse.CommonServerResponse.GetCmdType() == proto.CmdType_RUN {
			fmt.Printf("\nCommand string to execute: %v\n", serverPlatformResponse.InstructionPayload.Cmd)
			cmdStr := serverPlatformResponse.InstructionPayload.Cmd
			exeErr = executeServerInstructions(cmdStr)
		}

		clientPlatformMsgType = setPlatformMsgType(serverPlatformResponse.GetSdpOperType(), exeErr)

		traceId := serverPlatformResponse.CommonServerResponse.ResponseHeader.Identifiers.TraceID

		clientMsg = &proto.ClientPlatformMsg{CommonClientInfo: &proto.CommonClientInfo{RequestHeader:
			common.NewReqHdrGenerateMessageID(traceId), ClientInfo: &clientInfo},
				ClientPlatformMsgType: clientPlatformMsgType }
	}
}

func main() {
	var pnpServer string
	var pnpOpType string
	var serverCert string

	service := grpc.NewService(
		micro.Flags(
			cli.StringFlag{
				Name : "pnp_server",
				Value: "PnPService",
				Usage: "PnP server name registered to registry",
			},
			cli.StringFlag{
				Name : "pnp_op_type",
				Usage: "Specifies pnp operation type, supported values are" +
					"installPackages, deployPlatform",
			},
			cli.StringFlag{
				Name : "server_cert_file",
				Value: "/certs/server.crt",
				Usage: "Path of server certificate file",
			},
		),
	)
	service.Init(
		micro.Action(func(c *cli.Context) {
			pnpServer = c.String("pnp_server")
			pnpOpType = c.String("pnp_op_type")
			serverCert = c.String("server_cert_file")
			if pnpOpType == "" {
				log.Fatalf("PnP operation type not specified, supported values are" +
					"installPackages, deploySDPMaster, deploySDPSatellite")
			}
		}),
	)

	caCert, err := ioutil.ReadFile(serverCert)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	tlsConfig := &tls.Config{
		RootCAs:      caCertPool,
	}
	tlsConfig.BuildNameToCertificate()
	service.Init(
		micro.Transport(transport.NewTransport(transport.Secure(true))),
		grpc.WithTLS(tlsConfig),
	)

	pnpClient := proto.PnPServiceClient(pnpServer, service.Client())

	switch pnpOpType {
	case "installPackages":
		{
			fmt.Println("Initializing package installation..")
			installPackages(pnpClient)
		}
	case "deploySDPMaster":
		{
			fmt.Println("Requesting deployment of SDP Master from PnP server..")
			deployPlatform(pnpClient, "master")
		}
	case "deploySDPSatellite":
		{
			fmt.Println("Requesting deployment of SDP Satellite from PnP server..")
			deployPlatform(pnpClient, "satellite")
		}
	}
}
