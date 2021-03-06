// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/hailo-srv/proto/auth/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/hailo-srv/proto/auth/auth.proto

It has these top-level messages:
	TestRequest
	TestResponse
*/
package auth

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

type TestRequest struct {
}

func (m *TestRequest) Reset()                    { *m = TestRequest{} }
func (m *TestRequest) String() string            { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()               {}
func (*TestRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type TestResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *TestResponse) Reset()                    { *m = TestResponse{} }
func (m *TestResponse) String() string            { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()               {}
func (*TestResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TestResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "TestRequest")
	proto.RegisterType((*TestResponse)(nil), "TestResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) Test(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := grpc.Invoke(ctx, "/Auth/Test", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	Test(context.Context, *TestRequest) (*TestResponse, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/Test",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).Test(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _Auth_Test_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/hailo-srv/proto/auth/auth.proto",
}

func init() { proto.RegisterFile("github.com/micro/hailo-srv/proto/auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0xcf, 0x48, 0xcc,
	0xcc, 0xc9, 0xd7, 0x2d, 0x2e, 0x2a, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2c, 0x2d,
	0xc9, 0x00, 0x13, 0x7a, 0x60, 0xbe, 0x12, 0x2f, 0x17, 0x77, 0x48, 0x6a, 0x71, 0x49, 0x50, 0x6a,
	0x61, 0x69, 0x6a, 0x71, 0x89, 0x92, 0x1a, 0x17, 0x0f, 0x84, 0x5b, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
	0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1,
	0x19, 0x04, 0xe5, 0x19, 0xe9, 0x72, 0xb1, 0x38, 0x96, 0x96, 0x64, 0x08, 0xa9, 0x72, 0xb1, 0x80,
	0xd4, 0x0b, 0xf1, 0xe8, 0x21, 0x99, 0x22, 0xc5, 0xab, 0x87, 0x6c, 0x88, 0x12, 0x43, 0x12, 0x1b,
	0xd8, 0x32, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0xf9, 0xcc, 0x79, 0xa0, 0x00, 0x00,
	0x00,
}
