// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Command.proto

package LiquidRpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CmdCommand struct {
	UserID               string   `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	UserToken            string   `protobuf:"bytes,2,opt,name=UserToken,proto3" json:"UserToken,omitempty"`
	Platform             string   `protobuf:"bytes,3,opt,name=Platform,proto3" json:"Platform,omitempty"`
	CmdId                string   `protobuf:"bytes,4,opt,name=CmdId,proto3" json:"CmdId,omitempty"`
	CmdSn                uint64   `protobuf:"varint,5,opt,name=CmdSn,proto3" json:"CmdSn,omitempty"`
	CmdName              string   `protobuf:"bytes,6,opt,name=CmdName,proto3" json:"CmdName,omitempty"`
	CmdData              []byte   `protobuf:"bytes,7,opt,name=CmdData,proto3" json:"CmdData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdCommand) Reset()         { *m = CmdCommand{} }
func (m *CmdCommand) String() string { return proto.CompactTextString(m) }
func (*CmdCommand) ProtoMessage()    {}
func (*CmdCommand) Descriptor() ([]byte, []int) {
	return fileDescriptor_8272e154246f9d3c, []int{0}
}

func (m *CmdCommand) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdCommand.Unmarshal(m, b)
}
func (m *CmdCommand) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdCommand.Marshal(b, m, deterministic)
}
func (m *CmdCommand) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdCommand.Merge(m, src)
}
func (m *CmdCommand) XXX_Size() int {
	return xxx_messageInfo_CmdCommand.Size(m)
}
func (m *CmdCommand) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdCommand.DiscardUnknown(m)
}

var xxx_messageInfo_CmdCommand proto.InternalMessageInfo

func (m *CmdCommand) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *CmdCommand) GetUserToken() string {
	if m != nil {
		return m.UserToken
	}
	return ""
}

func (m *CmdCommand) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *CmdCommand) GetCmdId() string {
	if m != nil {
		return m.CmdId
	}
	return ""
}

func (m *CmdCommand) GetCmdSn() uint64 {
	if m != nil {
		return m.CmdSn
	}
	return 0
}

func (m *CmdCommand) GetCmdName() string {
	if m != nil {
		return m.CmdName
	}
	return ""
}

func (m *CmdCommand) GetCmdData() []byte {
	if m != nil {
		return m.CmdData
	}
	return nil
}

type CmdCommandReply struct {
	CmdSn                uint64   `protobuf:"varint,1,opt,name=CmdSn,proto3" json:"CmdSn,omitempty"`
	CmdData              []byte   `protobuf:"bytes,2,opt,name=CmdData,proto3" json:"CmdData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CmdCommandReply) Reset()         { *m = CmdCommandReply{} }
func (m *CmdCommandReply) String() string { return proto.CompactTextString(m) }
func (*CmdCommandReply) ProtoMessage()    {}
func (*CmdCommandReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_8272e154246f9d3c, []int{1}
}

func (m *CmdCommandReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CmdCommandReply.Unmarshal(m, b)
}
func (m *CmdCommandReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CmdCommandReply.Marshal(b, m, deterministic)
}
func (m *CmdCommandReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CmdCommandReply.Merge(m, src)
}
func (m *CmdCommandReply) XXX_Size() int {
	return xxx_messageInfo_CmdCommandReply.Size(m)
}
func (m *CmdCommandReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CmdCommandReply.DiscardUnknown(m)
}

var xxx_messageInfo_CmdCommandReply proto.InternalMessageInfo

func (m *CmdCommandReply) GetCmdSn() uint64 {
	if m != nil {
		return m.CmdSn
	}
	return 0
}

func (m *CmdCommandReply) GetCmdData() []byte {
	if m != nil {
		return m.CmdData
	}
	return nil
}

func init() {
	proto.RegisterType((*CmdCommand)(nil), "RPC.CmdCommand")
	proto.RegisterType((*CmdCommandReply)(nil), "RPC.CmdCommandReply")
}

func init() { proto.RegisterFile("Command.proto", fileDescriptor_8272e154246f9d3c) }

var fileDescriptor_8272e154246f9d3c = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0x80, 0xcd, 0xfe, 0xb4, 0xee, 0xa8, 0x2c, 0x84, 0x45, 0x86, 0xc5, 0x43, 0xe9, 0xa9, 0xa7,
	0x1e, 0xd6, 0x27, 0x28, 0x29, 0x48, 0x2f, 0x52, 0xa2, 0x3e, 0x40, 0x24, 0xf1, 0x62, 0xa7, 0x29,
	0x31, 0x17, 0x5f, 0xce, 0x67, 0x13, 0xd3, 0xd4, 0xa0, 0x7b, 0xcb, 0x37, 0x5f, 0xf8, 0x12, 0x06,
	0x6e, 0x84, 0x25, 0x52, 0xa3, 0xae, 0x27, 0x67, 0xbd, 0xe5, 0x6b, 0xd9, 0x8b, 0xf2, 0x8b, 0x01,
	0x08, 0xd2, 0xd1, 0xf0, 0x5b, 0xc8, 0x5e, 0x3e, 0x8c, 0xeb, 0x5a, 0x64, 0x05, 0xab, 0x76, 0x32,
	0x12, 0xbf, 0x83, 0xdd, 0xcf, 0xe9, 0xd9, 0xbe, 0x9b, 0x11, 0x57, 0x41, 0xa5, 0x01, 0x3f, 0xc2,
	0x65, 0x3f, 0x28, 0xff, 0x66, 0x1d, 0xe1, 0x3a, 0xc8, 0x5f, 0xe6, 0x07, 0xd8, 0x0a, 0xd2, 0x9d,
	0xc6, 0x4d, 0x10, 0x33, 0xc4, 0xe9, 0xd3, 0x88, 0xdb, 0x82, 0x55, 0x1b, 0x39, 0x03, 0x47, 0xc8,
	0x05, 0xe9, 0x47, 0x45, 0x06, 0xb3, 0x70, 0x7b, 0xc1, 0x68, 0x5a, 0xe5, 0x15, 0xe6, 0x05, 0xab,
	0xae, 0xe5, 0x82, 0x65, 0x03, 0xfb, 0xf4, 0x7f, 0x69, 0xa6, 0xe1, 0x33, 0xc5, 0xd9, 0x79, 0x3c,
	0x24, 0x56, 0x7f, 0x12, 0xa7, 0x06, 0xae, 0x1e, 0x14, 0x99, 0x46, 0xab, 0xc9, 0x1b, 0xc7, 0x4f,
	0x90, 0x2f, 0xeb, 0xd8, 0xd7, 0xb2, 0x17, 0x75, 0xea, 0x1f, 0x0f, 0xff, 0x06, 0xe1, 0xc1, 0xf2,
	0xe2, 0x35, 0x0b, 0x2b, 0xbd, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x89, 0x56, 0x40, 0xe3, 0x63,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GameAdapterClient is the client API for GameAdapter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GameAdapterClient interface {
	Command(ctx context.Context, in *CmdCommand, opts ...grpc.CallOption) (*CmdCommandReply, error)
}

type gameAdapterClient struct {
	cc *grpc.ClientConn
}

func NewGameAdapterClient(cc *grpc.ClientConn) GameAdapterClient {
	return &gameAdapterClient{cc}
}

func (c *gameAdapterClient) Command(ctx context.Context, in *CmdCommand, opts ...grpc.CallOption) (*CmdCommandReply, error) {
	out := new(CmdCommandReply)
	err := c.cc.Invoke(ctx, "/RPC.GameAdapter/Command", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameAdapterServer is the server API for GameAdapter service.
type GameAdapterServer interface {
	Command(context.Context, *CmdCommand) (*CmdCommandReply, error)
}

// UnimplementedGameAdapterServer can be embedded to have forward compatible implementations.
type UnimplementedGameAdapterServer struct {
}

func (*UnimplementedGameAdapterServer) Command(ctx context.Context, req *CmdCommand) (*CmdCommandReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Command not implemented")
}

func RegisterGameAdapterServer(s *grpc.Server, srv GameAdapterServer) {
	s.RegisterService(&_GameAdapter_serviceDesc, srv)
}

func _GameAdapter_Command_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CmdCommand)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameAdapterServer).Command(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RPC.GameAdapter/Command",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameAdapterServer).Command(ctx, req.(*CmdCommand))
	}
	return interceptor(ctx, in, info, handler)
}

var _GameAdapter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "RPC.GameAdapter",
	HandlerType: (*GameAdapterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Command",
			Handler:    _GameAdapter_Command_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Command.proto",
}