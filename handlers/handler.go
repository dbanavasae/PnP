package handlers

import (
	"context"
	"fmt"
	"io"
	"github.com/golang/protobuf/ptypes"
	"github.com/PnP/common"
	pb "github.com/PnP/common/proto"
	proto "github.com/PnP/pnp-proto"
	"os"
)

type PnPService struct {}

var (
	packageInfoFile = "/PnP/config/packageInfo.json"
	serverPkgResponse = &proto.ServerPkgResponse{}
)

func setServerResponse (pkg common.Package,
	clientPkgMsgType proto.ClientPkgMsgType, len int) (cmdType proto.CmdType,
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

		//ToDo: Update version if required
	  case proto.ClientPkgMsgType_PKG_INSTALLED:
		{
			fmt.Printf("Package %v already installed\n", pkg.Name)
			if len == 0 {
				cmdType = proto.CmdType_CLOSE_CONN
			} else {
				cmdType = proto.CmdType_INFO
				pkgOperType = proto.PkgOperType_GET_NEXT_PKG
			}
		}

	  case proto.ClientPkgMsgType_PKG_INSTALL_SUCCESS:
		{
			fmt.Printf("Package %v installed\n", pkg.Name)
			if len == 0 {
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
		// Add case for get_instruction_file
		//
	}
	return
}
//ToDo: Set dead timer value
func (s *PnPService) GetPackages (ctx context.Context, stream proto.PnP_GetPackagesStream) (err error) {
	pwd, _ := os.Getwd()
	packageInfo, err := common.FromFile(pwd+packageInfoFile)
	if err != nil {
		fmt.Printf("Json marshalling of packageInfo.json failed, %v", err)
	}

	len := len(packageInfo.Packages)
	//pkg := &common.Package{}

	for _, pkg := range packageInfo.Packages {
		var cmdType proto.CmdType
		var pkgOperType proto.PkgOperType
		var exeCmd []string
		len = len - 1

		for {
			clientPkgMsg, err := stream.Recv()
			if err == io.EOF {
				return nil
			}

			if err != nil {
				fmt.Printf("Error reading data from client, Error : %v", err)
				break
			}

			cmdType, pkgOperType, exeCmd = setServerResponse(pkg, clientPkgMsg.GetClientPkgMsgType(), len)

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

func (s *PnPService) InitPlatformDeploy (ctx context.Context, stream proto.PnP_InitPlatformDeployStream) (err error) {
	return nil

}

/*func (s *PnPService) Query (ctx context.Context, in *proto.ClientRequestMsg, out *proto.ServerResponse) (err error) {
	return nil
}*/

