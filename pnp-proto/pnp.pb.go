// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pnp.proto

/*
Package pnp is a generated protocol buffer package.

It is generated from these files:
	pnp.proto

It has these top-level messages:
	ClientInfo
	CommonClientInfo
	ClientPkgMsg
	ClientPlatformMsg
	InstructionPayload
	CommonServerResponse
	Chunk
	ServerPkgResponse
	ServerPlatformResponse
	Query
	ClientRequestMsg
*/
package pnp

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/PnP/common/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ClientPkgMsgType int32

const (
	ClientPkgMsgType_PKG_INIT              ClientPkgMsgType = 0
	ClientPkgMsgType_PKG_NOT_INSTALLED     ClientPkgMsgType = 1
	ClientPkgMsgType_PKG_INSTALLED         ClientPkgMsgType = 2
	ClientPkgMsgType_PKG_VERSION_CHANGED   ClientPkgMsgType = 3
	ClientPkgMsgType_PKG_VERSION_NO_CHANGE ClientPkgMsgType = 4
	ClientPkgMsgType_PKG_UNINSTALL_SUCCESS ClientPkgMsgType = 5
	ClientPkgMsgType_PKG_UNINSTALL_FAILED  ClientPkgMsgType = 6
	ClientPkgMsgType_PKG_INSTALL_SUCCESS   ClientPkgMsgType = 7
	ClientPkgMsgType_PKG_INSTALL_FAILED    ClientPkgMsgType = 8
)

var ClientPkgMsgType_name = map[int32]string{
	0: "PKG_INIT",
	1: "PKG_NOT_INSTALLED",
	2: "PKG_INSTALLED",
	3: "PKG_VERSION_CHANGED",
	4: "PKG_VERSION_NO_CHANGE",
	5: "PKG_UNINSTALL_SUCCESS",
	6: "PKG_UNINSTALL_FAILED",
	7: "PKG_INSTALL_SUCCESS",
	8: "PKG_INSTALL_FAILED",
}
var ClientPkgMsgType_value = map[string]int32{
	"PKG_INIT":              0,
	"PKG_NOT_INSTALLED":     1,
	"PKG_INSTALLED":         2,
	"PKG_VERSION_CHANGED":   3,
	"PKG_VERSION_NO_CHANGE": 4,
	"PKG_UNINSTALL_SUCCESS": 5,
	"PKG_UNINSTALL_FAILED":  6,
	"PKG_INSTALL_SUCCESS":   7,
	"PKG_INSTALL_FAILED":    8,
}

func (x ClientPkgMsgType) String() string {
	return proto.EnumName(ClientPkgMsgType_name, int32(x))
}
func (ClientPkgMsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClientPlatformMsgType int32

const (
	ClientPlatformMsgType_PLATFORM_INIT                      ClientPlatformMsgType = 0
	ClientPlatformMsgType_PLATFORM_ALREADY_INSTALLED         ClientPlatformMsgType = 1
	ClientPlatformMsgType_PLATFORM_NOT_INSTALLED             ClientPlatformMsgType = 2
	ClientPlatformMsgType_PLATFORM_ARTIFACT_DOWNLOAD_SUCCESS ClientPlatformMsgType = 3
	ClientPlatformMsgType_PLATFORM_ARTIFACT_DOWNLOAD_FAILED  ClientPlatformMsgType = 4
	ClientPlatformMsgType_PLATFORM_DEPLOYMENT_SUCCESS        ClientPlatformMsgType = 5
	ClientPlatformMsgType_PLATFORM_DEPLOYMENT_FAILED         ClientPlatformMsgType = 6
	ClientPlatformMsgType_PLATFORM_SERVICE_UP                ClientPlatformMsgType = 7
	ClientPlatformMsgType_PLATFORM_SERVICE_DOWN              ClientPlatformMsgType = 8
)

var ClientPlatformMsgType_name = map[int32]string{
	0: "PLATFORM_INIT",
	1: "PLATFORM_ALREADY_INSTALLED",
	2: "PLATFORM_NOT_INSTALLED",
	3: "PLATFORM_ARTIFACT_DOWNLOAD_SUCCESS",
	4: "PLATFORM_ARTIFACT_DOWNLOAD_FAILED",
	5: "PLATFORM_DEPLOYMENT_SUCCESS",
	6: "PLATFORM_DEPLOYMENT_FAILED",
	7: "PLATFORM_SERVICE_UP",
	8: "PLATFORM_SERVICE_DOWN",
}
var ClientPlatformMsgType_value = map[string]int32{
	"PLATFORM_INIT":                      0,
	"PLATFORM_ALREADY_INSTALLED":         1,
	"PLATFORM_NOT_INSTALLED":             2,
	"PLATFORM_ARTIFACT_DOWNLOAD_SUCCESS": 3,
	"PLATFORM_ARTIFACT_DOWNLOAD_FAILED":  4,
	"PLATFORM_DEPLOYMENT_SUCCESS":        5,
	"PLATFORM_DEPLOYMENT_FAILED":         6,
	"PLATFORM_SERVICE_UP":                7,
	"PLATFORM_SERVICE_DOWN":              8,
}

func (x ClientPlatformMsgType) String() string {
	return proto.EnumName(ClientPlatformMsgType_name, int32(x))
}
func (ClientPlatformMsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CmdType int32

const (
	CmdType_RUN        CmdType = 0
	CmdType_CONFIG     CmdType = 1
	CmdType_INFO       CmdType = 2
	CmdType_CLOSE_CONN CmdType = 3
)

var CmdType_name = map[int32]string{
	0: "RUN",
	1: "CONFIG",
	2: "INFO",
	3: "CLOSE_CONN",
}
var CmdType_value = map[string]int32{
	"RUN":        0,
	"CONFIG":     1,
	"INFO":       2,
	"CLOSE_CONN": 3,
}

func (x CmdType) String() string {
	return proto.EnumName(CmdType_name, int32(x))
}
func (CmdType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type PkgOperType int32

const (
	PkgOperType_IS_PKG_INSTALLED      PkgOperType = 0
	PkgOperType_INSTALL_PKG           PkgOperType = 1
	PkgOperType_INSTALL_PKG_FROM_FILE PkgOperType = 2
	PkgOperType_IS_PKG_OUTDATED       PkgOperType = 3
	PkgOperType_UNINSTALL_PKG         PkgOperType = 4
	PkgOperType_GET_NEXT_PKG          PkgOperType = 5
)

var PkgOperType_name = map[int32]string{
	0: "IS_PKG_INSTALLED",
	1: "INSTALL_PKG",
	2: "INSTALL_PKG_FROM_FILE",
	3: "IS_PKG_OUTDATED",
	4: "UNINSTALL_PKG",
	5: "GET_NEXT_PKG",
}
var PkgOperType_value = map[string]int32{
	"IS_PKG_INSTALLED":      0,
	"INSTALL_PKG":           1,
	"INSTALL_PKG_FROM_FILE": 2,
	"IS_PKG_OUTDATED":       3,
	"UNINSTALL_PKG":         4,
	"GET_NEXT_PKG":          5,
}

func (x PkgOperType) String() string {
	return proto.EnumName(PkgOperType_name, int32(x))
}
func (PkgOperType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SDPOperType int32

const (
	SDPOperType_IS_PLATFORM_INSTALLED      SDPOperType = 0
	SDPOperType_DOWNLOAD_PLATFORM_ARTIFACT SDPOperType = 1
	SDPOperType_DEPLOY_PLATFORM            SDPOperType = 2
	SDPOperType_CHECK_PLATFORM_STATUS      SDPOperType = 3
)

var SDPOperType_name = map[int32]string{
	0: "IS_PLATFORM_INSTALLED",
	1: "DOWNLOAD_PLATFORM_ARTIFACT",
	2: "DEPLOY_PLATFORM",
	3: "CHECK_PLATFORM_STATUS",
}
var SDPOperType_value = map[string]int32{
	"IS_PLATFORM_INSTALLED":      0,
	"DOWNLOAD_PLATFORM_ARTIFACT": 1,
	"DEPLOY_PLATFORM":            2,
	"CHECK_PLATFORM_STATUS":      3,
}

func (x SDPOperType) String() string {
	return proto.EnumName(SDPOperType_name, int32(x))
}
func (SDPOperType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ClientInfo struct {
	OsType   string `protobuf:"bytes,1,opt,name=osType" json:"osType,omitempty"`
	OsFlavor string `protobuf:"bytes,2,opt,name=osFlavor" json:"osFlavor,omitempty"`
	ArchType string `protobuf:"bytes,3,opt,name=archType" json:"archType,omitempty"`
	ClientId string `protobuf:"bytes,4,opt,name=clientId" json:"clientId,omitempty"`
}

func (m *ClientInfo) Reset()                    { *m = ClientInfo{} }
func (m *ClientInfo) String() string            { return proto.CompactTextString(m) }
func (*ClientInfo) ProtoMessage()               {}
func (*ClientInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientInfo) GetOsType() string {
	if m != nil {
		return m.OsType
	}
	return ""
}

func (m *ClientInfo) GetOsFlavor() string {
	if m != nil {
		return m.OsFlavor
	}
	return ""
}

func (m *ClientInfo) GetArchType() string {
	if m != nil {
		return m.ArchType
	}
	return ""
}

func (m *ClientInfo) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type CommonClientInfo struct {
	RequestHeader *common.RequestHeader `protobuf:"bytes,1,opt,name=requestHeader" json:"requestHeader,omitempty"`
	ClientInfo    *ClientInfo           `protobuf:"bytes,2,opt,name=clientInfo" json:"clientInfo,omitempty"`
}

func (m *CommonClientInfo) Reset()                    { *m = CommonClientInfo{} }
func (m *CommonClientInfo) String() string            { return proto.CompactTextString(m) }
func (*CommonClientInfo) ProtoMessage()               {}
func (*CommonClientInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CommonClientInfo) GetRequestHeader() *common.RequestHeader {
	if m != nil {
		return m.RequestHeader
	}
	return nil
}

func (m *CommonClientInfo) GetClientInfo() *ClientInfo {
	if m != nil {
		return m.ClientInfo
	}
	return nil
}

type ClientPkgMsg struct {
	CommonClientInfo *CommonClientInfo `protobuf:"bytes,1,opt,name=commonClientInfo" json:"commonClientInfo,omitempty"`
	ClientPkgMsgType ClientPkgMsgType  `protobuf:"varint,2,opt,name=clientPkgMsgType,enum=pnp.ClientPkgMsgType" json:"clientPkgMsgType,omitempty"`
}

func (m *ClientPkgMsg) Reset()                    { *m = ClientPkgMsg{} }
func (m *ClientPkgMsg) String() string            { return proto.CompactTextString(m) }
func (*ClientPkgMsg) ProtoMessage()               {}
func (*ClientPkgMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientPkgMsg) GetCommonClientInfo() *CommonClientInfo {
	if m != nil {
		return m.CommonClientInfo
	}
	return nil
}

func (m *ClientPkgMsg) GetClientPkgMsgType() ClientPkgMsgType {
	if m != nil {
		return m.ClientPkgMsgType
	}
	return ClientPkgMsgType_PKG_INIT
}

type ClientPlatformMsg struct {
	CommonClientInfo      *CommonClientInfo     `protobuf:"bytes,1,opt,name=commonClientInfo" json:"commonClientInfo,omitempty"`
	ClientPlatformMsgType ClientPlatformMsgType `protobuf:"varint,2,opt,name=clientPlatformMsgType,enum=pnp.ClientPlatformMsgType" json:"clientPlatformMsgType,omitempty"`
	Error                 string                `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *ClientPlatformMsg) Reset()                    { *m = ClientPlatformMsg{} }
func (m *ClientPlatformMsg) String() string            { return proto.CompactTextString(m) }
func (*ClientPlatformMsg) ProtoMessage()               {}
func (*ClientPlatformMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientPlatformMsg) GetCommonClientInfo() *CommonClientInfo {
	if m != nil {
		return m.CommonClientInfo
	}
	return nil
}

func (m *ClientPlatformMsg) GetClientPlatformMsgType() ClientPlatformMsgType {
	if m != nil {
		return m.ClientPlatformMsgType
	}
	return ClientPlatformMsgType_PLATFORM_INIT
}

func (m *ClientPlatformMsg) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type InstructionPayload struct {
	Cmd []string `protobuf:"bytes,1,rep,name=cmd" json:"cmd,omitempty"`
}

func (m *InstructionPayload) Reset()                    { *m = InstructionPayload{} }
func (m *InstructionPayload) String() string            { return proto.CompactTextString(m) }
func (*InstructionPayload) ProtoMessage()               {}
func (*InstructionPayload) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *InstructionPayload) GetCmd() []string {
	if m != nil {
		return m.Cmd
	}
	return nil
}

type CommonServerResponse struct {
	ResponseHeader *common.ResponseHeader `protobuf:"bytes,1,opt,name=responseHeader" json:"responseHeader,omitempty"`
	CmdType        CmdType                `protobuf:"varint,2,opt,name=cmdType,enum=pnp.CmdType" json:"cmdType,omitempty"`
}

func (m *CommonServerResponse) Reset()                    { *m = CommonServerResponse{} }
func (m *CommonServerResponse) String() string            { return proto.CompactTextString(m) }
func (*CommonServerResponse) ProtoMessage()               {}
func (*CommonServerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CommonServerResponse) GetResponseHeader() *common.ResponseHeader {
	if m != nil {
		return m.ResponseHeader
	}
	return nil
}

func (m *CommonServerResponse) GetCmdType() CmdType {
	if m != nil {
		return m.CmdType
	}
	return CmdType_RUN
}

type Chunk struct {
	FileChunk []byte `protobuf:"bytes,1,opt,name=fileChunk,proto3" json:"fileChunk,omitempty"`
}

func (m *Chunk) Reset()                    { *m = Chunk{} }
func (m *Chunk) String() string            { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()               {}
func (*Chunk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Chunk) GetFileChunk() []byte {
	if m != nil {
		return m.FileChunk
	}
	return nil
}

type ServerPkgResponse struct {
	CommonServerResponse *CommonServerResponse `protobuf:"bytes,1,opt,name=commonServerResponse" json:"commonServerResponse,omitempty"`
	PkgOperType          PkgOperType           `protobuf:"varint,2,opt,name=pkgOperType,enum=pnp.PkgOperType" json:"pkgOperType,omitempty"`
	InstructionPayload   *InstructionPayload   `protobuf:"bytes,3,opt,name=instructionPayload" json:"instructionPayload,omitempty"`
	Chunk                *Chunk                `protobuf:"bytes,4,opt,name=chunk" json:"chunk,omitempty"`
}

func (m *ServerPkgResponse) Reset()                    { *m = ServerPkgResponse{} }
func (m *ServerPkgResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerPkgResponse) ProtoMessage()               {}
func (*ServerPkgResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ServerPkgResponse) GetCommonServerResponse() *CommonServerResponse {
	if m != nil {
		return m.CommonServerResponse
	}
	return nil
}

func (m *ServerPkgResponse) GetPkgOperType() PkgOperType {
	if m != nil {
		return m.PkgOperType
	}
	return PkgOperType_IS_PKG_INSTALLED
}

func (m *ServerPkgResponse) GetInstructionPayload() *InstructionPayload {
	if m != nil {
		return m.InstructionPayload
	}
	return nil
}

func (m *ServerPkgResponse) GetChunk() *Chunk {
	if m != nil {
		return m.Chunk
	}
	return nil
}

type ServerPlatformResponse struct {
	CommonServerResponse *CommonServerResponse `protobuf:"bytes,1,opt,name=commonServerResponse" json:"commonServerResponse,omitempty"`
	SdpOperType          SDPOperType           `protobuf:"varint,2,opt,name=sdpOperType,enum=pnp.SDPOperType" json:"sdpOperType,omitempty"`
	InstructionPayload   *InstructionPayload   `protobuf:"bytes,3,opt,name=instructionPayload" json:"instructionPayload,omitempty"`
}

func (m *ServerPlatformResponse) Reset()                    { *m = ServerPlatformResponse{} }
func (m *ServerPlatformResponse) String() string            { return proto.CompactTextString(m) }
func (*ServerPlatformResponse) ProtoMessage()               {}
func (*ServerPlatformResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *ServerPlatformResponse) GetCommonServerResponse() *CommonServerResponse {
	if m != nil {
		return m.CommonServerResponse
	}
	return nil
}

func (m *ServerPlatformResponse) GetSdpOperType() SDPOperType {
	if m != nil {
		return m.SdpOperType
	}
	return SDPOperType_IS_PLATFORM_INSTALLED
}

func (m *ServerPlatformResponse) GetInstructionPayload() *InstructionPayload {
	if m != nil {
		return m.InstructionPayload
	}
	return nil
}

type Query struct {
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type ClientRequestMsg struct {
	RequestHeader *common.RequestHeader `protobuf:"bytes,1,opt,name=RequestHeader" json:"RequestHeader,omitempty"`
	QueryData     *Query                `protobuf:"bytes,2,opt,name=QueryData" json:"QueryData,omitempty"`
}

func (m *ClientRequestMsg) Reset()                    { *m = ClientRequestMsg{} }
func (m *ClientRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*ClientRequestMsg) ProtoMessage()               {}
func (*ClientRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ClientRequestMsg) GetRequestHeader() *common.RequestHeader {
	if m != nil {
		return m.RequestHeader
	}
	return nil
}

func (m *ClientRequestMsg) GetQueryData() *Query {
	if m != nil {
		return m.QueryData
	}
	return nil
}

func init() {
	proto.RegisterType((*ClientInfo)(nil), "pnp.ClientInfo")
	proto.RegisterType((*CommonClientInfo)(nil), "pnp.CommonClientInfo")
	proto.RegisterType((*ClientPkgMsg)(nil), "pnp.ClientPkgMsg")
	proto.RegisterType((*ClientPlatformMsg)(nil), "pnp.ClientPlatformMsg")
	proto.RegisterType((*InstructionPayload)(nil), "pnp.InstructionPayload")
	proto.RegisterType((*CommonServerResponse)(nil), "pnp.CommonServerResponse")
	proto.RegisterType((*Chunk)(nil), "pnp.Chunk")
	proto.RegisterType((*ServerPkgResponse)(nil), "pnp.ServerPkgResponse")
	proto.RegisterType((*ServerPlatformResponse)(nil), "pnp.ServerPlatformResponse")
	proto.RegisterType((*Query)(nil), "pnp.Query")
	proto.RegisterType((*ClientRequestMsg)(nil), "pnp.ClientRequestMsg")
	proto.RegisterEnum("pnp.ClientPkgMsgType", ClientPkgMsgType_name, ClientPkgMsgType_value)
	proto.RegisterEnum("pnp.ClientPlatformMsgType", ClientPlatformMsgType_name, ClientPlatformMsgType_value)
	proto.RegisterEnum("pnp.CmdType", CmdType_name, CmdType_value)
	proto.RegisterEnum("pnp.PkgOperType", PkgOperType_name, PkgOperType_value)
	proto.RegisterEnum("pnp.SDPOperType", SDPOperType_name, SDPOperType_value)
}

func init() { proto.RegisterFile("pnp.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 996 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xe2, 0x46,
	0x14, 0x8e, 0x31, 0x84, 0xe4, 0x90, 0x64, 0x87, 0xd9, 0xc0, 0x26, 0x6c, 0xb5, 0x4d, 0x2d, 0x6d,
	0x14, 0xe5, 0x22, 0xa9, 0xd2, 0x9b, 0x4a, 0x2b, 0xad, 0x64, 0xd9, 0x86, 0x58, 0x01, 0xdb, 0x1d,
	0x9b, 0x6d, 0xf7, 0xca, 0xf2, 0x1a, 0x87, 0x20, 0xc0, 0x76, 0x0d, 0xac, 0x96, 0x9b, 0xaa, 0xbd,
	0xec, 0x6d, 0xd5, 0xa7, 0xe8, 0x43, 0xf4, 0x69, 0xf6, 0x41, 0x2a, 0x8f, 0xff, 0xc6, 0x80, 0x2a,
	0x55, 0xda, 0xde, 0x71, 0xce, 0xf7, 0x9d, 0x6f, 0xbe, 0x39, 0x33, 0x73, 0x0c, 0x1c, 0x86, 0x7e,
	0x78, 0x13, 0x46, 0xc1, 0x32, 0xc0, 0x7c, 0xe8, 0x87, 0x9d, 0xab, 0xf1, 0x64, 0xf9, 0xb4, 0xfa,
	0x70, 0xe3, 0x06, 0xf3, 0x5b, 0xc3, 0x37, 0x6e, 0xdd, 0x60, 0x3e, 0x0f, 0xfc, 0x5b, 0xca, 0x48,
	0x83, 0x84, 0x2e, 0x7c, 0x02, 0x90, 0x66, 0x13, 0xcf, 0x5f, 0xaa, 0xfe, 0x63, 0x80, 0xdb, 0xb0,
	0x1f, 0x2c, 0xac, 0x75, 0xe8, 0x9d, 0x71, 0x17, 0xdc, 0xd5, 0x21, 0x49, 0x23, 0xdc, 0x81, 0x83,
	0x60, 0xd1, 0x9d, 0x39, 0x1f, 0x83, 0xe8, 0xac, 0x42, 0x91, 0x3c, 0x8e, 0x31, 0x27, 0x72, 0x9f,
	0x68, 0x15, 0x9f, 0x60, 0x59, 0x1c, 0x63, 0x6e, 0xa2, 0x3e, 0x3a, 0xab, 0x26, 0x58, 0x16, 0x0b,
	0xbf, 0x72, 0x80, 0x24, 0x6a, 0x85, 0x31, 0xf0, 0x06, 0x8e, 0x23, 0xef, 0xe7, 0x95, 0xb7, 0x58,
	0xde, 0x7b, 0xce, 0xc8, 0x8b, 0xa8, 0x8f, 0xc6, 0x5d, 0xeb, 0x26, 0x35, 0x4d, 0x58, 0x90, 0x94,
	0xb9, 0xf8, 0x16, 0xc0, 0xcd, 0xa5, 0xa8, 0xcf, 0xc6, 0xdd, 0xb3, 0x9b, 0xb8, 0x35, 0xc5, 0x0a,
	0x84, 0xa1, 0x08, 0x7f, 0x72, 0x70, 0x94, 0x40, 0xc6, 0x74, 0x3c, 0x58, 0x8c, 0xb1, 0x08, 0xc8,
	0xdd, 0xb0, 0x94, 0x3b, 0xa0, 0x3a, 0x1b, 0x20, 0xd9, 0xa2, 0x53, 0x09, 0x46, 0x92, 0xb6, 0x25,
	0xb6, 0x72, 0x92, 0x49, 0x6c, 0x80, 0x64, 0x8b, 0x2e, 0xfc, 0xcd, 0x41, 0x33, 0xa5, 0xcd, 0x9c,
	0xe5, 0x63, 0x10, 0xcd, 0xbf, 0x90, 0x37, 0x03, 0x5a, 0xee, 0xa6, 0x2e, 0x63, 0xb0, 0xc3, 0x1a,
	0x2c, 0x33, 0xc8, 0xee, 0x42, 0x7c, 0x0a, 0x35, 0x2f, 0x8a, 0x82, 0x28, 0x3d, 0xf9, 0x24, 0x10,
	0x2e, 0x01, 0xab, 0xfe, 0x62, 0x19, 0xad, 0xdc, 0xe5, 0x24, 0xf0, 0x0d, 0x67, 0x3d, 0x0b, 0x9c,
	0x11, 0x46, 0xc0, 0xbb, 0xf3, 0xd1, 0x19, 0x77, 0xc1, 0x5f, 0x1d, 0x92, 0xf8, 0xa7, 0xf0, 0x0b,
	0x9c, 0x26, 0xae, 0x4d, 0x2f, 0xfa, 0xe8, 0x45, 0xc4, 0x5b, 0x84, 0x81, 0xbf, 0xf0, 0xf0, 0x5b,
	0x38, 0x89, 0xd2, 0xdf, 0xa5, 0x6b, 0xd0, 0x2e, 0xae, 0x01, 0x8b, 0x92, 0x0d, 0x36, 0xbe, 0x84,
	0xba, 0x3b, 0x1f, 0x31, 0x3b, 0x3b, 0x4a, 0x76, 0x96, 0xe4, 0x48, 0x06, 0x0a, 0xaf, 0xa1, 0x26,
	0x3d, 0xad, 0xfc, 0x29, 0xfe, 0x0a, 0x0e, 0x1f, 0x27, 0x33, 0x8f, 0x06, 0x74, 0xad, 0x23, 0x52,
	0x24, 0x84, 0xdf, 0x2a, 0xd0, 0x4c, 0x1c, 0x1a, 0xd3, 0x71, 0x6e, 0x72, 0x00, 0xa7, 0xee, 0x0e,
	0xf3, 0xa9, 0xd5, 0x73, 0xe6, 0x4c, 0xca, 0x04, 0xb2, 0xb3, 0x0c, 0xdf, 0x41, 0x23, 0x9c, 0x8e,
	0xf5, 0xd0, 0x8b, 0x18, 0xdf, 0x88, 0xaa, 0x18, 0x45, 0x9e, 0xb0, 0x24, 0xdc, 0x03, 0x3c, 0xd9,
	0xea, 0x33, 0x3d, 0x8a, 0xc6, 0xdd, 0x0b, 0x5a, 0xba, 0x7d, 0x0c, 0x64, 0x47, 0x09, 0xbe, 0x80,
	0x9a, 0x4b, 0xf7, 0x5e, 0xa5, 0xb5, 0x90, 0x98, 0x8f, 0x33, 0x24, 0x01, 0x84, 0xcf, 0x1c, 0xb4,
	0xd3, 0x1e, 0xa4, 0x57, 0xe0, 0x7f, 0x6c, 0xc4, 0x62, 0x14, 0xee, 0x6c, 0x84, 0x29, 0x1b, 0x45,
	0x23, 0x18, 0xd2, 0x17, 0x6b, 0x84, 0x50, 0x87, 0xda, 0x0f, 0x2b, 0x2f, 0x5a, 0x0b, 0x6b, 0x40,
	0xc9, 0x43, 0x48, 0x27, 0x4e, 0xfc, 0x02, 0xdf, 0xc0, 0x31, 0xf9, 0x0f, 0xc3, 0xa9, 0x14, 0xe2,
	0x2b, 0x38, 0xa4, 0xca, 0xb2, 0xb3, 0x74, 0xd2, 0xd9, 0x94, 0xb4, 0x99, 0x66, 0x49, 0x01, 0x5e,
	0x7f, 0xe6, 0xb2, 0xb5, 0x8b, 0x99, 0x80, 0x8f, 0xe0, 0xc0, 0x78, 0xe8, 0xd9, 0xaa, 0xa6, 0x5a,
	0x68, 0x0f, 0xb7, 0xa0, 0x19, 0x47, 0x9a, 0x6e, 0xd9, 0xaa, 0x66, 0x5a, 0x62, 0xbf, 0xaf, 0xc8,
	0x88, 0xc3, 0x4d, 0x38, 0x4e, 0x48, 0x59, 0xaa, 0x82, 0x5f, 0xc0, 0xf3, 0x38, 0xf5, 0x4e, 0x21,
	0xa6, 0xaa, 0x6b, 0xb6, 0x74, 0x2f, 0x6a, 0x3d, 0x45, 0x46, 0x3c, 0x3e, 0x87, 0x16, 0x0b, 0x68,
	0x7a, 0x8a, 0xa1, 0x6a, 0x06, 0x0d, 0xb5, 0x54, 0xc8, 0x36, 0x87, 0x92, 0xa4, 0x98, 0x26, 0xaa,
	0xe1, 0x33, 0x38, 0x2d, 0x43, 0x5d, 0x51, 0x8d, 0x17, 0xda, 0xcf, 0x16, 0xda, 0x2c, 0xa9, 0xe3,
	0x36, 0x60, 0x16, 0x48, 0x0b, 0x0e, 0xae, 0xff, 0xaa, 0x40, 0x6b, 0xe7, 0xac, 0xa1, 0xdb, 0xe8,
	0x8b, 0x56, 0x57, 0x27, 0x83, 0x6c, 0xc3, 0xaf, 0xa0, 0x93, 0xa7, 0xc4, 0x3e, 0x51, 0x44, 0xf9,
	0x7d, 0x69, 0xe7, 0x1d, 0x68, 0xe7, 0x78, 0xb9, 0x2b, 0x15, 0x7c, 0x09, 0x42, 0x51, 0x4b, 0x2c,
	0xb5, 0x2b, 0x4a, 0x96, 0x2d, 0xeb, 0x3f, 0x6a, 0x7d, 0x5d, 0x94, 0x73, 0xa3, 0x3c, 0x7e, 0x0d,
	0xdf, 0xfc, 0x0b, 0x2f, 0xf5, 0x5d, 0xc5, 0x5f, 0xc3, 0xcb, 0x9c, 0x26, 0x2b, 0x46, 0x5f, 0x7f,
	0x3f, 0x50, 0x34, 0x8b, 0xe9, 0x11, 0xeb, 0x95, 0x21, 0x94, 0x3a, 0x95, 0xe1, 0xa6, 0x42, 0xde,
	0xa9, 0x92, 0x62, 0x0f, 0x0d, 0x54, 0xa7, 0x7d, 0xdf, 0x04, 0xe2, 0xf5, 0xd1, 0xc1, 0xf5, 0xf7,
	0x50, 0x4f, 0xa7, 0x17, 0xae, 0x03, 0x4f, 0x86, 0x1a, 0xda, 0xc3, 0x00, 0xfb, 0x92, 0xae, 0x75,
	0xd5, 0x1e, 0xe2, 0xf0, 0x01, 0x54, 0x55, 0xad, 0xab, 0xa3, 0x0a, 0x3e, 0x01, 0x90, 0xfa, 0xba,
	0xa9, 0xd8, 0x92, 0xae, 0x69, 0x88, 0xbf, 0xfe, 0x9d, 0x83, 0x06, 0x33, 0x40, 0xf0, 0x29, 0x20,
	0xd5, 0xb4, 0xcb, 0xd7, 0x64, 0x0f, 0x3f, 0x83, 0x46, 0x76, 0x40, 0xc6, 0x43, 0x2c, 0x78, 0x0e,
	0x2d, 0x26, 0x61, 0x77, 0x89, 0x3e, 0xb0, 0xbb, 0x6a, 0x5f, 0x41, 0x15, 0xfc, 0x1c, 0x9e, 0xa5,
	0x0a, 0xfa, 0xd0, 0x92, 0x45, 0x8b, 0x5e, 0xa7, 0x26, 0x1c, 0x17, 0x97, 0x22, 0x96, 0xa8, 0x62,
	0x04, 0x47, 0x3d, 0xc5, 0xb2, 0x35, 0xe5, 0x27, 0x8b, 0x66, 0x6a, 0xd7, 0x9f, 0xa0, 0xc1, 0x3c,
	0x61, 0xba, 0x86, 0x69, 0x33, 0x47, 0x5d, 0xf8, 0x79, 0x05, 0x9d, 0xbc, 0xf3, 0x5b, 0x87, 0x82,
	0xb8, 0xd8, 0x43, 0xd2, 0xda, 0x1c, 0x45, 0x95, 0x58, 0x4f, 0xba, 0x57, 0xa4, 0x87, 0xa2, 0xc2,
	0xb4, 0x44, 0x6b, 0x68, 0x22, 0xfe, 0xee, 0x0f, 0x0e, 0x78, 0xc3, 0x37, 0xf0, 0x5b, 0x68, 0xf4,
	0xbc, 0xa5, 0xe1, 0xb8, 0x53, 0x67, 0xec, 0x2d, 0x70, 0x73, 0xeb, 0x93, 0xdc, 0x69, 0x27, 0x93,
	0x66, 0x73, 0xdc, 0x0b, 0x7b, 0x57, 0xdc, 0xb7, 0x1c, 0x7e, 0x80, 0x13, 0xd9, 0x0b, 0x67, 0xc1,
	0x3a, 0xbb, 0xb3, 0xb8, 0xbd, 0xfb, 0xa3, 0xd9, 0x79, 0xc9, 0xea, 0x6c, 0x8c, 0xcc, 0x44, 0xec,
	0xc3, 0x3e, 0xfd, 0x0b, 0xf6, 0xdd, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11, 0xd4, 0x95, 0xdc,
	0xbe, 0x09, 0x00, 0x00,
}
