package invoke

import (
	"time"
	"io"
	"fmt"
	"log"
	"github.com/PnP/common"
	"golang.org/x/net/context"
	proto "github.com/PnP/pnp-proto"
)

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

func DeployPlatform(pnpClient proto.PnPService, platformType string) {
	cxt, cancel := context.WithTimeout(context.Background(), time.Minute*120)
	defer cancel()
	stream, err := pnpClient.DeployPlatform(cxt)
	var clientPlatformMsgType proto.ClientPlatformMsgType

	clientInfo := common.PopulateClientDetails()

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
			exeErr = common.ExecuteServerInstructions(cmdStr)
		}

		clientPlatformMsgType = setPlatformMsgType(serverPlatformResponse.GetSdpOperType(), exeErr)

		traceId := serverPlatformResponse.CommonServerResponse.ResponseHeader.Identifiers.TraceID

		clientMsg = &proto.ClientPlatformMsg{CommonClientInfo: &proto.CommonClientInfo{RequestHeader:
		common.NewReqHdrGenerateMessageID(traceId), ClientInfo: &clientInfo},
			ClientPlatformMsgType: clientPlatformMsgType }
	}
}
