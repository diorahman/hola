// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hola/hola.proto

/*
Package hola is a generated protocol buffer package.

It is generated from these files:
	hola/hola.proto

It has these top-level messages:
	GetRequest
	GetResponse
*/
package hola

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

type GetRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResponse struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "hola.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "hola.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Hola service

type HolaClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type holaClient struct {
	cc *grpc.ClientConn
}

func NewHolaClient(cc *grpc.ClientConn) HolaClient {
	return &holaClient{cc}
}

func (c *holaClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/hola.Hola/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hola service

type HolaServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterHolaServer(s *grpc.Server, srv HolaServer) {
	s.RegisterService(&_Hola_serviceDesc, srv)
}

func _Hola_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HolaServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hola.Hola/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HolaServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Hola_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hola.Hola",
	HandlerType: (*HolaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Hola_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hola/hola.proto",
}

func init() { proto.RegisterFile("hola/hola.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xc8, 0xcf, 0x49,
	0xd4, 0x07, 0x11, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0xb6, 0x92, 0x1c, 0x17,
	0x97, 0x7b, 0x6a, 0x49, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x00, 0x17, 0x73, 0x76,
	0x6a, 0xa5, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x88, 0xa9, 0xa4, 0xcc, 0xc5, 0x0d, 0x96,
	0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x85,
	0x2a, 0x81, 0x70, 0x8c, 0x4c, 0xb8, 0x58, 0x3c, 0xf2, 0x73, 0x12, 0x85, 0x74, 0xb8, 0x98, 0xdd,
	0x53, 0x4b, 0x84, 0x04, 0xf4, 0xc0, 0xd6, 0x20, 0xcc, 0x95, 0x12, 0x44, 0x12, 0x81, 0x98, 0xa4,
	0xc4, 0x90, 0xc4, 0x06, 0x76, 0x87, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x2f, 0xb9, 0xb7,
	0x9a, 0x00, 0x00, 0x00,
}
