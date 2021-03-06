// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/hailo-srv/proto/api/api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/hailo-srv/proto/api/api.proto

It has these top-level messages:
	StatusRequest
	StatusResponse
*/
package api

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

type StatusRequest struct {
}

func (m *StatusRequest) Reset()                    { *m = StatusRequest{} }
func (m *StatusRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusRequest) ProtoMessage()               {}
func (*StatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StatusResponse struct {
	Status string `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
}

func (m *StatusResponse) Reset()                    { *m = StatusResponse{} }
func (m *StatusResponse) String() string            { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()               {}
func (*StatusResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatusResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*StatusRequest)(nil), "StatusRequest")
	proto.RegisterType((*StatusResponse)(nil), "StatusResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Api service

type ApiClient interface {
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := grpc.Invoke(ctx, "/Api/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiServer interface {
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Api/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Api_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/micro/hailo-srv/proto/api/api.proto",
}

func init() { proto.RegisterFile("github.com/micro/hailo-srv/proto/api/api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0xcd, 0x4c, 0x2e, 0xca, 0xd7, 0xcf, 0x48, 0xcc,
	0xcc, 0xc9, 0xd7, 0x2d, 0x2e, 0x2a, 0xd3, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2c, 0xc8,
	0x04, 0x61, 0x3d, 0x30, 0x4f, 0x89, 0x9f, 0x8b, 0x37, 0xb8, 0x24, 0xb1, 0xa4, 0xb4, 0x38, 0x28,
	0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x49, 0x83, 0x8b, 0x0f, 0x26, 0x50, 0x5c, 0x90, 0x9f, 0x57,
	0x9c, 0x2a, 0x24, 0xc6, 0xc5, 0x56, 0x0c, 0x16, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82,
	0xf2, 0x8c, 0x8c, 0xb8, 0x98, 0x1d, 0x0b, 0x32, 0x85, 0xb4, 0xb9, 0xd8, 0x20, 0x1a, 0x84, 0xf8,
	0xf4, 0x50, 0x8c, 0x92, 0xe2, 0xd7, 0x43, 0x35, 0x49, 0x89, 0x21, 0x89, 0x0d, 0x6c, 0xab, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x20, 0x7c, 0xae, 0xa7, 0x00, 0x00, 0x00,
}
