package handlers

import (
	"context"
	"fmt"
	"io"
	"os"
	"log"
	"github.com/golang/protobuf/ptypes"
	"github.com/PnP/common"
	"github.com/PnP/config"
	pb "github.com/PnP/common/proto"
	proto "github.com/PnP/pnp-proto"
)

type PnPService struct {}

func setSDPOperType(msgType proto.ClientPlatformMsgType) (cmdType proto.CmdType,
	platformOperType proto.SDPOperType) {

	switch msgType {
	case proto.ClientPlatformMsgType_SDP_PLATFORM_MASTER_INIT:
		fallthrough
	case proto.ClientPlatformMsgType_SDP_PLATFORM_SATELLITE_INIT:
		{
			platformOperType = proto.SDPOperType_IS_SDP_PLATFORM_INSTALLED
			cmdType = proto.CmdType_RUN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_ALREADY_INSTALLED:
		{
			fmt.Println("SDP Platform is already installed.. Closing connection")
			cmdType = proto.CmdType_CLOSE_CONN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_NOT_INSTALLED:
		{
			fmt.Println("Sending SDP Platform artifact download instructions..")
			platformOperType = proto.SDPOperType_DOWNLOAD_SDP_PLATFORM_ARTIFACT
			cmdType = proto.CmdType_RUN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_ARTIFACT_DOWNLOAD_SUCCESS:
		{
			fmt.Println("SDP Platform artifact downloaded successfully.. Sending Deploy instructions")
			platformOperType = proto.SDPOperType_DEPLOY_SDP_PLATFORM
			cmdType = proto.CmdType_RUN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_ARTIFACT_DOWNLOAD_FAILED:
		{
			fmt.Println("SDP Platform Artifact download failed.. closing stream")
			cmdType = proto.CmdType_CLOSE_CONN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_DEPLOYMENT_SUCCESS:
		{
			fmt.Println("SDP Platform deployment successful.. Checking deployment status")
			platformOperType = proto.SDPOperType_CHECK_SDP_PLATFORM_STATUS
			cmdType = proto.CmdType_RUN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_DEPLOYMENT_FAILED:
		{
			fmt.Println("SDP Platform deployment failed.. Closing stream")
			cmdType = proto.CmdType_CLOSE_CONN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_SERVICE_UP:
		{
			fmt.Println("SDP Platform services are UP and RUNNING.. Closing connection")
			cmdType = proto.CmdType_CLOSE_CONN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_SERVICE_DOWN:
		{
			fmt.Println("SDP Platform services are DOWN.. Closing connection")
			cmdType = proto.CmdType_CLOSE_CONN
		}
	}
	return
}

func getSDPDeployCmd(operType proto.SDPOperType, sdpDeploymentType proto.SDPDeploymentType) (exeCmd []string, err error) {
	platformDeploy := &common.PlatformDeploy{}
	pwd, _ := os.Getwd()
	if err = common.GetConfigFromJson(pwd + config.PlatformDeployFile, platformDeploy); err != nil {
		log.Fatalf("Unable to get config data from JSON file, Error: %v", err)
	}

	switch operType {
	case proto.SDPOperType_IS_SDP_PLATFORM_INSTALLED:
		{
			if sdpDeploymentType == proto.SDPDeploymentType_MASTER {
				exeCmd = platformDeploy.DeployInfo.CheckSDPMasterInstallation
			} else {
				exeCmd = platformDeploy.DeployInfo.CheckSDPSatelliteInstallation
			}
		}
	case proto.SDPOperType_DOWNLOAD_SDP_PLATFORM_ARTIFACT:
		{
			exeCmd = platformDeploy.DeployInfo.DownloadSDPArtifact
		}
	case proto.SDPOperType_DEPLOY_SDP_PLATFORM:
		{
			if sdpDeploymentType == proto.SDPDeploymentType_MASTER {
				exeCmd = platformDeploy.DeployInfo.InstallSDPMasterPlatform
			} else {
				exeCmd = platformDeploy.DeployInfo.InstallSDPSatellitePlatform
			}
		}
	case proto.SDPOperType_CHECK_SDP_PLATFORM_STATUS:
		{
			if sdpDeploymentType == proto.SDPDeploymentType_MASTER {
				exeCmd = platformDeploy.DeployInfo.CheckSDPMasterStatus
			} else {
				exeCmd = platformDeploy.DeployInfo.CheckSDPSatelliteStatus
			}
		}
	}
	return
}

func (s *PnPService) DeployPlatform (ctx context.Context, stream proto.PnP_DeployPlatformStream) (err error) {
	serverPlatformResponse := &proto.ServerPlatformResponse{}
	var exeCmd []string
	var sdpDeploymentType proto.SDPDeploymentType

	for {
		clientPlatformMsg, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("\nDone with platform install.. closing stream\n")
			stream.Close()
			return nil
		}

		if err != nil {
			fmt.Printf("\nError reading data from client, Error : %v", err)
			break
		}

		fmt.Printf("\nReceived SDP Deploy request from client, Request type: %v\n", clientPlatformMsg.GetClientPlatformMsgType())

		if clientPlatformMsg.GetClientPlatformMsgType() == proto.ClientPlatformMsgType_SDP_PLATFORM_MASTER_INIT {
			fmt.Printf("Client %v requested for SDP Master installation, Starting deployment...", clientPlatformMsg.GetCommonClientInfo().ClientInfo.ClientId)
			sdpDeploymentType = proto.SDPDeploymentType_MASTER
		} else if clientPlatformMsg.GetClientPlatformMsgType() == proto.ClientPlatformMsgType_SDP_PLATFORM_SATELLITE_INIT {
			fmt.Printf("Client %v requested for SDP Satellite installation, Starting deployment...", clientPlatformMsg.GetCommonClientInfo().ClientInfo.ClientId)
			sdpDeploymentType = proto.SDPDeploymentType_SATELLITE
		}

		cmdType, platformOperType := setSDPOperType(clientPlatformMsg.GetClientPlatformMsgType())

		if cmdType == proto.CmdType_RUN {
			exeCmd, err = getSDPDeployCmd(platformOperType, sdpDeploymentType)
			if err != nil {
				fmt.Printf("Getting SDP deploy instructions failed, Error: %v", err)
				break
			}
		}

		serverPlatformResponse = &proto.ServerPlatformResponse{CommonServerResponse: &proto.CommonServerResponse{
			ResponseHeader: &pb.ResponseHeader{Identifiers: &pb.Identifiers{TraceID: clientPlatformMsg.
				CommonClientInfo.RequestHeader.Identifiers.TraceID, MessageID: clientPlatformMsg.
				CommonClientInfo.RequestHeader.Identifiers.MessageID}, ResponseTimestamp:
			ptypes.TimestampNow()}, CmdType: cmdType}, SdpOperType:
		platformOperType, InstructionPayload:
		&proto.InstructionPayload{Cmd: exeCmd}}

		fmt.Printf("Sending SDP Deploy instructions to PnP client, Instruction type: %v", platformOperType)
		if err = stream.Send(serverPlatformResponse); err != nil {
			fmt.Printf("Error while sending response to client, Error: %v", err)
			break
		}
	}
	stream.Close()
	return nil
}
