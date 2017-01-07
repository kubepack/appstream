// Code generated by protoc-gen-go.
// source: metdata.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	metdata.proto

It has these top-level messages:
	GitRequest
	GitResponse
	DockerRequest
	DockerResponse
*/
package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"

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

type GitRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *GitRequest) Reset()                    { *m = GitRequest{} }
func (m *GitRequest) String() string            { return proto.CompactTextString(m) }
func (*GitRequest) ProtoMessage()               {}
func (*GitRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GitRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GitResponse struct {
	Status    *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Greetings string                  `protobuf:"bytes,2,opt,name=greetings" json:"greetings,omitempty"`
}

func (m *GitResponse) Reset()                    { *m = GitResponse{} }
func (m *GitResponse) String() string            { return proto.CompactTextString(m) }
func (*GitResponse) ProtoMessage()               {}
func (*GitResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GitResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GitResponse) GetGreetings() string {
	if m != nil {
		return m.Greetings
	}
	return ""
}

type DockerRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DockerRequest) Reset()                    { *m = DockerRequest{} }
func (m *DockerRequest) String() string            { return proto.CompactTextString(m) }
func (*DockerRequest) ProtoMessage()               {}
func (*DockerRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DockerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DockerResponse struct {
	Status    *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Greetings string                  `protobuf:"bytes,2,opt,name=greetings" json:"greetings,omitempty"`
}

func (m *DockerResponse) Reset()                    { *m = DockerResponse{} }
func (m *DockerResponse) String() string            { return proto.CompactTextString(m) }
func (*DockerResponse) ProtoMessage()               {}
func (*DockerResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DockerResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DockerResponse) GetGreetings() string {
	if m != nil {
		return m.Greetings
	}
	return ""
}

func init() {
	proto.RegisterType((*GitRequest)(nil), "appscode.apps.v1beta1.GitRequest")
	proto.RegisterType((*GitResponse)(nil), "appscode.apps.v1beta1.GitResponse")
	proto.RegisterType((*DockerRequest)(nil), "appscode.apps.v1beta1.DockerRequest")
	proto.RegisterType((*DockerResponse)(nil), "appscode.apps.v1beta1.DockerResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metadata service

type MetadataClient interface {
	Git(ctx context.Context, in *GitRequest, opts ...grpc.CallOption) (*GitResponse, error)
	Docker(ctx context.Context, in *DockerRequest, opts ...grpc.CallOption) (*DockerResponse, error)
}

type metadataClient struct {
	cc *grpc.ClientConn
}

func NewMetadataClient(cc *grpc.ClientConn) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) Git(ctx context.Context, in *GitRequest, opts ...grpc.CallOption) (*GitResponse, error) {
	out := new(GitResponse)
	err := grpc.Invoke(ctx, "/appscode.apps.v1beta1.Metadata/Git", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metadataClient) Docker(ctx context.Context, in *DockerRequest, opts ...grpc.CallOption) (*DockerResponse, error) {
	out := new(DockerResponse)
	err := grpc.Invoke(ctx, "/appscode.apps.v1beta1.Metadata/Docker", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metadata service

type MetadataServer interface {
	Git(context.Context, *GitRequest) (*GitResponse, error)
	Docker(context.Context, *DockerRequest) (*DockerResponse, error)
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_Git_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).Git(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.apps.v1beta1.Metadata/Git",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).Git(ctx, req.(*GitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metadata_Docker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).Docker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.apps.v1beta1.Metadata/Docker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).Docker(ctx, req.(*DockerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.apps.v1beta1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Git",
			Handler:    _Metadata_Git_Handler,
		},
		{
			MethodName: "Docker",
			Handler:    _Metadata_Docker_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metdata.proto",
}

func init() { proto.RegisterFile("metdata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x92, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0x49, 0xfe, 0x7f, 0x6a, 0x3b, 0xa5, 0x1e, 0x16, 0xc4, 0x12, 0x6a, 0x6d, 0x53, 0x05,
	0x0f, 0xb2, 0x4b, 0xeb, 0xc9, 0x6b, 0x11, 0x7a, 0x12, 0x4a, 0xbd, 0x89, 0x20, 0xdb, 0x66, 0x08,
	0x8b, 0x26, 0xbb, 0x76, 0xa7, 0x82, 0x1e, 0x7b, 0xf0, 0xea, 0xc1, 0x67, 0xf1, 0x49, 0x7c, 0x05,
	0x1f, 0x44, 0xba, 0x49, 0x6d, 0x0b, 0x36, 0x9e, 0xbc, 0x84, 0x25, 0xdf, 0x6f, 0xbe, 0xfd, 0x76,
	0x66, 0xa0, 0x96, 0x20, 0x45, 0x92, 0x24, 0x37, 0x53, 0x4d, 0x9a, 0xed, 0x49, 0x63, 0xec, 0x44,
	0x47, 0xc8, 0x17, 0x07, 0xfe, 0xd8, 0x1d, 0x23, 0xc9, 0x6e, 0xd0, 0x88, 0xb5, 0x8e, 0xef, 0x51,
	0x48, 0xa3, 0x84, 0x4c, 0x53, 0x4d, 0x92, 0x94, 0x4e, 0x6d, 0x56, 0x14, 0x34, 0x97, 0x45, 0x5b,
	0xf4, 0xc3, 0x0d, 0x3d, 0xa2, 0x27, 0x83, 0x56, 0xb8, 0x6f, 0x06, 0x84, 0x2d, 0x80, 0x81, 0xa2,
	0x11, 0x3e, 0xcc, 0xd0, 0x12, 0x63, 0xf0, 0x3f, 0x95, 0x09, 0xd6, 0xbd, 0x96, 0x77, 0x52, 0x19,
	0xb9, 0x73, 0x78, 0x03, 0x55, 0x47, 0x58, 0xa3, 0x53, 0x8b, 0x4c, 0x40, 0xc9, 0x92, 0xa4, 0x99,
	0x75, 0x50, 0xb5, 0xb7, 0xcf, 0xbf, 0x73, 0x67, 0xf6, 0xfc, 0xca, 0xc9, 0xa3, 0x1c, 0x63, 0x0d,
	0xa8, 0xc4, 0x53, 0x44, 0x52, 0x69, 0x6c, 0xeb, 0xbe, 0x33, 0x5e, 0xfd, 0x08, 0x3b, 0x50, 0xbb,
	0xd0, 0x93, 0x3b, 0x9c, 0x16, 0x45, 0xb8, 0x85, 0xdd, 0x25, 0xf4, 0x27, 0x29, 0x7a, 0xaf, 0x3e,
	0x94, 0x2f, 0x91, 0xe4, 0x62, 0x1c, 0xec, 0x19, 0xfe, 0x0d, 0x14, 0xb1, 0x36, 0xff, 0x71, 0x20,
	0x7c, 0xd5, 0xae, 0x20, 0x2c, 0x42, 0xb2, 0xa4, 0xe1, 0xe9, 0xfc, 0xbd, 0xee, 0x97, 0xbd, 0xf9,
	0xc7, 0xe7, 0x9b, 0xdf, 0x62, 0xcd, 0x6c, 0x4e, 0xc6, 0x58, 0x91, 0xc3, 0x22, 0xc9, 0xef, 0x15,
	0xb1, 0x22, 0xf6, 0xe2, 0x41, 0x29, 0x7b, 0x2a, 0x3b, 0xda, 0x62, 0xbe, 0xd1, 0xae, 0xe0, 0xf8,
	0x17, 0x2a, 0x4f, 0x21, 0xd6, 0x52, 0x74, 0x58, 0xbb, 0x20, 0x45, 0xe4, 0x0a, 0xfb, 0xe7, 0x70,
	0x30, 0xd1, 0xc9, 0xba, 0xb9, 0xda, 0xb8, 0xa0, 0x5f, 0x5b, 0xf6, 0x6b, 0xb8, 0xd8, 0xa3, 0xa1,
	0x77, 0xbd, 0x93, 0x2b, 0xe3, 0x92, 0xdb, 0xac, 0xb3, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd,
	0x9a, 0xdb, 0x90, 0xe0, 0x02, 0x00, 0x00,
}
