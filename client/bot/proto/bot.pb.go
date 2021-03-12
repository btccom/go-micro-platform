// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/btccom/go-micro-platform/v2/bot/proto/bot.proto

/*
Package go_micro_bot is a generated protocol buffer package.

It is generated from these files:
	github.com/btccom/go-micro-platform/v2/bot/proto/bot.proto

It has these top-level messages:
	HelpRequest
	HelpResponse
	ExecRequest
	ExecResponse
*/
package go_micro_bot

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelpRequest struct {
}

func (m *HelpRequest) Reset()                    { *m = HelpRequest{} }
func (m *HelpRequest) String() string            { return proto.CompactTextString(m) }
func (*HelpRequest) ProtoMessage()               {}
func (*HelpRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type HelpResponse struct {
	Usage       string `protobuf:"bytes,1,opt,name=usage" json:"usage,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
}

func (m *HelpResponse) Reset()                    { *m = HelpResponse{} }
func (m *HelpResponse) String() string            { return proto.CompactTextString(m) }
func (*HelpResponse) ProtoMessage()               {}
func (*HelpResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelpResponse) GetUsage() string {
	if m != nil {
		return m.Usage
	}
	return ""
}

func (m *HelpResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type ExecRequest struct {
	Args []string `protobuf:"bytes,1,rep,name=args" json:"args,omitempty"`
}

func (m *ExecRequest) Reset()                    { *m = ExecRequest{} }
func (m *ExecRequest) String() string            { return proto.CompactTextString(m) }
func (*ExecRequest) ProtoMessage()               {}
func (*ExecRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ExecRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type ExecResponse struct {
	Result []byte `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Error  string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *ExecResponse) Reset()                    { *m = ExecResponse{} }
func (m *ExecResponse) String() string            { return proto.CompactTextString(m) }
func (*ExecResponse) ProtoMessage()               {}
func (*ExecResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ExecResponse) GetResult() []byte {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ExecResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*HelpRequest)(nil), "go.micro.bot.HelpRequest")
	proto.RegisterType((*HelpResponse)(nil), "go.micro.bot.HelpResponse")
	proto.RegisterType((*ExecRequest)(nil), "go.micro.bot.ExecRequest")
	proto.RegisterType((*ExecResponse)(nil), "go.micro.bot.ExecResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Command service

type CommandClient interface {
	Help(ctx context.Context, in *HelpRequest, opts ...grpc.CallOption) (*HelpResponse, error)
	Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error)
}

type commandClient struct {
	cc *grpc.ClientConn
}

func NewCommandClient(cc *grpc.ClientConn) CommandClient {
	return &commandClient{cc}
}

func (c *commandClient) Help(ctx context.Context, in *HelpRequest, opts ...grpc.CallOption) (*HelpResponse, error) {
	out := new(HelpResponse)
	err := grpc.Invoke(ctx, "/go.micro.bot.Command/Help", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandClient) Exec(ctx context.Context, in *ExecRequest, opts ...grpc.CallOption) (*ExecResponse, error) {
	out := new(ExecResponse)
	err := grpc.Invoke(ctx, "/go.micro.bot.Command/Exec", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Command service

type CommandServer interface {
	Help(context.Context, *HelpRequest) (*HelpResponse, error)
	Exec(context.Context, *ExecRequest) (*ExecResponse, error)
}

func RegisterCommandServer(s *grpc.Server, srv CommandServer) {
	s.RegisterService(&_Command_serviceDesc, srv)
}

func _Command_Help_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Help(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.bot.Command/Help",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Help(ctx, req.(*HelpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Command_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.bot.Command/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServer).Exec(ctx, req.(*ExecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Command_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.bot.Command",
	HandlerType: (*CommandServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Help",
			Handler:    _Command_Help_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _Command_Exec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/btccom/go-micro-platform/v2/bot/proto/bot.proto",
}

func init() { proto.RegisterFile("github.com/btccom/go-micro-platform/v2/bot/proto/bot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x50, 0xcb, 0x4a, 0xc4, 0x30,
	0x14, 0x9d, 0xea, 0x38, 0x32, 0xb7, 0x75, 0x13, 0x44, 0xea, 0xac, 0x6a, 0x56, 0x83, 0x8b, 0x0c,
	0xe8, 0x56, 0x70, 0x21, 0x8a, 0xeb, 0xfe, 0x41, 0xdb, 0xb9, 0xd4, 0xc2, 0xb4, 0x37, 0xde, 0x24,
	0xe0, 0x3f, 0xf8, 0xd3, 0x92, 0xc7, 0x22, 0x0c, 0x6e, 0xc2, 0x39, 0x39, 0xe1, 0x3c, 0x02, 0x8f,
	0xe3, 0x64, 0xbf, 0x5c, 0xaf, 0x06, 0x9a, 0x0f, 0xf3, 0x34, 0x30, 0xa5, 0xb3, 0x27, 0x7b, 0xd0,
	0x4c, 0x36, 0x20, 0x15, 0x90, 0xa8, 0x46, 0x52, 0x41, 0x55, 0x3d, 0x59, 0x79, 0x03, 0xe5, 0x27,
	0x9e, 0x74, 0x8b, 0xdf, 0x0e, 0x8d, 0x95, 0x1f, 0x50, 0x45, 0x6a, 0x34, 0x2d, 0x06, 0xc5, 0x2d,
	0x5c, 0x39, 0xd3, 0x8d, 0x58, 0x17, 0x4d, 0xb1, 0xdf, 0xb6, 0x91, 0x88, 0x06, 0xca, 0x23, 0x9a,
	0x81, 0x27, 0x6d, 0x27, 0x5a, 0xea, 0x8b, 0xa0, 0xe5, 0x57, 0xf2, 0x01, 0xca, 0xf7, 0x1f, 0x1c,
	0x92, 0xad, 0x10, 0xb0, 0xee, 0x78, 0x34, 0x75, 0xd1, 0x5c, 0xee, 0xb7, 0x6d, 0xc0, 0xf2, 0x05,
	0xaa, 0xf8, 0x24, 0x45, 0xdd, 0xc1, 0x86, 0xd1, 0xb8, 0x93, 0x0d, 0x59, 0x55, 0x9b, 0x98, 0xaf,
	0x80, 0xcc, 0xc4, 0x29, 0x26, 0x92, 0xa7, 0xdf, 0x02, 0xae, 0xdf, 0x68, 0x9e, 0xbb, 0xe5, 0x28,
	0x5e, 0x61, 0xed, 0x4b, 0x8b, 0x7b, 0x95, 0x4f, 0x53, 0xd9, 0xae, 0xdd, 0xee, 0x3f, 0x29, 0x06,
	0xcb, 0x95, 0x37, 0xf0, 0x55, 0xce, 0x0d, 0xb2, 0x05, 0xe7, 0x06, 0x79, 0x73, 0xb9, 0xea, 0x37,
	0xe1, 0x6b, 0x9f, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x15, 0x59, 0x33, 0x88, 0x01, 0x00,
	0x00,
}
