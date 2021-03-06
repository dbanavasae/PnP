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
