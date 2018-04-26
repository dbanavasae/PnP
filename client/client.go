package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"io"
	"time"
	//"flag"
	"github.com/micro/go-grpc"
	"golang.org/x/net/context"
	"github.com/PnP/common"
	"github.com/PnP/executor"
	proto "github.com/PnP/pnp-proto"
)

/*var (
	consulAddress string
	pnpServerName string
)

func init() {
	flag.StringVar(&consulAddress, "consulIP", "127.0.0.1",
		"Consul server's IP address at which gRPC client looks for service offered by gRPC server")
	flag.StringVar(&pnpServerName, "PnPServerName", "pnpService",
		"PnP server's registered service name in Consul")
	flag.Parse()

}*/

func init() {
	os.Setenv("MICRO_REGISTRY_ADDRESS", "192.168.50.129")
	os.Setenv("SDP_USER_PASSWD", "sdp")
	os.Setenv("SDP_NETWORK_INTERFACE", "ens33")
}

func populateClientDetails() (proto.ClientInfo) {
	archType := runtime.GOARCH
	osType := runtime.GOOS
	getOSFlavorCmd := "lsb_release -a | grep Description | awk -F':' '{print $2}'"

	osFlavor, err := executor.ExecuteCommand(getOSFlavorCmd)
	if err != nil {
		log.Fatalf("Error while getting OS type: %v", err)
	}
	clientId := "client1"

	clientInfo := proto.ClientInfo{OsType: osType, ArchType: archType, OsFlavor: osFlavor, ClientId: clientId}
	return clientInfo
}

func runServerInstructions(serverPkgResp *proto.ServerPkgResponse) (exeErr error) {
	var cmdString []string

	if serverPkgResp.CommonServerResponse.GetCmdType() == proto.CmdType_RUN {
		fmt.Printf("\nCommand string: %v\n", serverPkgResp.InstructionPayload.Cmd)

		cmdString = serverPkgResp.InstructionPayload.Cmd
		for _, cmd := range cmdString {
			var errStr string
			errStr, exeErr = executor.ExecuteCommand(cmd)
			if exeErr != nil {
				fmt.Printf("\nCommand <%v> failed to execute\nError: %v\n", cmd, errStr)
				break
			}
		}
	}
	return exeErr
}

func setPkgMsgType(serverPkgResp *proto.ServerPkgResponse, exeErr error) (clientPkgMsgType proto.ClientPkgMsgType) {
	switch serverPkgResp.GetPkgOperType() {
	  case proto.PkgOperType_IS_PKG_INSTALLED:
		if exeErr == nil {
			clientPkgMsgType = proto.ClientPkgMsgType_PKG_INSTALLED
		}else {
			clientPkgMsgType = proto.ClientPkgMsgType_PKG_NOT_INSTALLED
		}

	  case proto.PkgOperType_INSTALL_PKG:
		fallthrough

	  case proto.PkgOperType_INSTALL_PKG_FROM_FILE:
		if exeErr == nil {
			clientPkgMsgType = proto.ClientPkgMsgType_PKG_INSTALL_SUCCESS
		} else {
			clientPkgMsgType = proto.ClientPkgMsgType_PKG_INSTALL_FAILED
			fmt.Printf("\nFailed to install package\n")
		}

	  case proto.PkgOperType_GET_NEXT_PKG:
		clientPkgMsgType = proto.ClientPkgMsgType_PKG_INIT
	}
	return clientPkgMsgType
}

func installPackages(pclient proto.PnPService) {
	cxt, cancel := context.WithTimeout(context.Background(), time.Minute*20)
	defer cancel()
	stream, err := pclient.GetPackages(cxt)
	clientInfo := populateClientDetails()
	var traceId string

	clientPkgMsgType := proto.ClientPkgMsgType_PKG_INIT
	clientMsg := &proto.ClientPkgMsg{CommonClientInfo: &proto.CommonClientInfo{RequestHeader: common.NewReqHdrGenerateTraceAndMessageID(),
	ClientInfo: &clientInfo}, ClientPkgMsgType: clientPkgMsgType}
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

		/*if serverPkgResp.GetPkgOperType() == proto.PkgOperType_INSTALL_PKG_FROM_FILE {
			// set clientMsgType to get_instruction_file
			clientPkgMsgType = proto.ClientPkgMsgType_PKG_GET_FILE
		}*/

		exeErr := runServerInstructions(serverPkgResp)
		clientPkgMsgType = setPkgMsgType(serverPkgResp, exeErr)
		traceId = serverPkgResp.CommonServerResponse.ResponseHeader.Identifiers.TraceID

		clientMsg = &proto.ClientPkgMsg{CommonClientInfo: &proto.CommonClientInfo{RequestHeader: common.NewReqHdrGenerateMessageID(traceId),
		ClientInfo: &clientInfo}, ClientPkgMsgType: clientPkgMsgType }
	}
}

func main() {
	service := grpc.NewService()
	service.Init()
	pnpClient := proto.PnPServiceClient("PnPServer", service.Client())
	//ToDo: Provide flags to perform different operations
	installPackages(pnpClient)
}
