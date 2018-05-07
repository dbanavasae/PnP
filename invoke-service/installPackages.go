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
	case proto.PkgOperType_INSTALL_PKG, proto.PkgOperType_INSTALL_PKG_FROM_FILE:
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

func InstallPackages(pnpClient proto.PnPService) {
	cxt, cancel := context.WithTimeout(context.Background(), time.Minute*20)
	defer cancel()
	stream, err := pnpClient.GetPackages(cxt)
	clientInfo := common.PopulateClientDetails()
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
			cmdStr := serverPkgResp.InstructionPayload.Cmd
			exeErr = common.ExecuteServerInstructions(cmdStr)
		}

		clientPkgMsgType = setPkgMsgType(serverPkgResp.GetPkgOperType(), exeErr)

		traceId := serverPkgResp.CommonServerResponse.ResponseHeader.Identifiers.TraceID

		clientMsg = &proto.ClientPkgMsg{CommonClientInfo: &proto.CommonClientInfo{RequestHeader:
		common.NewReqHdrGenerateMessageID(traceId), ClientInfo: &clientInfo},
			ClientPkgMsgType: clientPkgMsgType }
	}
}
