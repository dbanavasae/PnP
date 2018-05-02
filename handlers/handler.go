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

func setPkgServerResponse (pkg common.Package,
	clientPkgMsgType proto.ClientPkgMsgType, numPkgsToInstall int) (cmdType proto.CmdType,
		pkgOperType proto.PkgOperType, exeCmd []string){

	switch clientPkgMsgType {
	case proto.ClientPkgMsgType_PKG_INIT:
		{
			cmdType = proto.CmdType_RUN
			pkgOperType = proto.PkgOperType_IS_PKG_INSTALLED
			exeCmd = pkg.CheckInstalledCmd
		}

	case proto.ClientPkgMsgType_PKG_NOT_INSTALLED:
		{
			cmdType = proto.CmdType_RUN
			if pkg.InstallFromFile != "" {
				pkgOperType = proto.PkgOperType_INSTALL_PKG_FROM_FILE

			} else {
				if pkg.UpdateRepo != nil {
					exeCmd = pkg.UpdateRepo
				}
				pkgOperType = proto.PkgOperType_INSTALL_PKG
				for _, cmd := range pkg.InstallInstructions {
					exeCmd = append(exeCmd, cmd)
				}
			}
		}

	case proto.ClientPkgMsgType_PKG_INSTALLED:
		{
			fmt.Printf("Package %v already installed\n", pkg.Name)
			if numPkgsToInstall == 0 {
				cmdType = proto.CmdType_CLOSE_CONN
			} else {
				cmdType = proto.CmdType_INFO
				pkgOperType = proto.PkgOperType_GET_NEXT_PKG
			}
		}

	case proto.ClientPkgMsgType_PKG_INSTALL_SUCCESS:
		{
			fmt.Printf("Package %v installed\n", pkg.Name)
			if numPkgsToInstall == 0 {
				fmt.Println("\nDone with all pkgs\n")
				cmdType = proto.CmdType_CLOSE_CONN
			} else {
				cmdType = proto.CmdType_INFO
				pkgOperType = proto.PkgOperType_GET_NEXT_PKG
			}
		}

	case proto.ClientPkgMsgType_PKG_INSTALL_FAILED:
		{
			fmt.Printf("Installation of package %v failed\n", pkg.Name)
			cmdType = proto.CmdType_CLOSE_CONN
		}
	}
	return
}
//ToDo: Set dead timer value
func (s *PnPService) GetPackages (ctx context.Context, stream proto.PnP_GetPackagesStream) (err error) {
	serverPkgResponse := &proto.ServerPkgResponse{}
	packageInfo := &common.PackageInfo{}
	pwd, _ := os.Getwd()

	if err = common.GetConfigFromJson(pwd + config.PackageFilePath, packageInfo); err != nil {
		log.Fatalf("Unable to get config data from JSON file, Error: %v", err)
	}

	numPkgsToInstall := len(packageInfo.Packages)

	for _, pkg := range packageInfo.Packages {
		numPkgsToInstall = numPkgsToInstall - 1

		for {
			clientPkgMsg, err := stream.Recv()
			if err == io.EOF {
				return nil
			}

			if err != nil {
				fmt.Printf("Error reading data from client, Error : %v", err)
				break
			}

			cmdType, pkgOperType, exeCmd := setPkgServerResponse(pkg, clientPkgMsg.GetClientPkgMsgType(), numPkgsToInstall)

			serverPkgResponse = &proto.ServerPkgResponse{CommonServerResponse: &proto.CommonServerResponse{ResponseHeader:
				&pb.ResponseHeader{Identifiers: &pb.Identifiers{TraceID: clientPkgMsg.CommonClientInfo.RequestHeader.Identifiers.TraceID,
					MessageID: clientPkgMsg.CommonClientInfo.RequestHeader.Identifiers.MessageID}, ResponseTimestamp:
						ptypes.TimestampNow()}, CmdType: cmdType}, InstructionPayload:
							&proto.InstructionPayload{exeCmd},
								PkgOperType: pkgOperType}

			if err = stream.Send(serverPkgResponse); err != nil {
				fmt.Printf("Error while sending response to client, Error: %v", err)
				break
			}

			if pkgOperType == proto.PkgOperType_GET_NEXT_PKG {
				break
			}
		}
		if err != nil {
			break
		}
	}
	stream.Close()
	return nil
}

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
			fmt.Println("Platform already installed.. closing stream")
			cmdType = proto.CmdType_CLOSE_CONN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_NOT_INSTALLED:
		{
			platformOperType = proto.SDPOperType_DOWNLOAD_SDP_PLATFORM_ARTIFACT
			cmdType = proto.CmdType_RUN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_ARTIFACT_DOWNLOAD_SUCCESS:
		{
			platformOperType = proto.SDPOperType_DEPLOY_SDP_PLATFORM
			cmdType = proto.CmdType_RUN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_ARTIFACT_DOWNLOAD_FAILED:
		{
			fmt.Println("Artifact download failed.. closing stream")
			cmdType = proto.CmdType_CLOSE_CONN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_DEPLOYMENT_SUCCESS:
		{
			platformOperType = proto.SDPOperType_CHECK_SDP_PLATFORM_STATUS
			cmdType = proto.CmdType_RUN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_SERVICE_UP:
		{
			fmt.Println("Platform services are up and running.. closing stream")
			cmdType = proto.CmdType_CLOSE_CONN
		}
	case proto.ClientPlatformMsgType_SDP_PLATFORM_SERVICE_DOWN:
		{
			fmt.Println("Platform services are down.. closing stream")
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
			return nil
		}

		if err != nil {
			fmt.Printf("Error reading data from client, Error : %v", err)
			break
		}

		if clientPlatformMsg.GetClientPlatformMsgType() == proto.ClientPlatformMsgType_SDP_PLATFORM_MASTER_INIT {
			sdpDeploymentType = proto.SDPDeploymentType_MASTER
		} else {
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

		if err = stream.Send(serverPlatformResponse); err != nil {
			fmt.Printf("Error while sending response to client, Error: %v", err)
			break
		}
	}
	return nil
}

/*func (s *PnPService) Query (ctx context.Context, in *proto.ClientRequestMsg, out *proto.ServerResponse) (err error) {
	return nil
}*/

