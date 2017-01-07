// Code generated by protoc-gen-go.
// source: version.proto
// DO NOT EDIT!

package v1beta1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/appscodeapis/appscode/api"
import appscode_dtypes "github.com/appscode/api/dtypes"
import appscode_version "github.com/appscode/api/version"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type VersionGetResponse struct {
	Status  *appscode_dtypes.Status   `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	Version *appscode_version.Version `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *VersionGetResponse) Reset()                    { *m = VersionGetResponse{} }
func (m *VersionGetResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionGetResponse) ProtoMessage()               {}
func (*VersionGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *VersionGetResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VersionGetResponse) GetVersion() *appscode_version.Version {
	if m != nil {
		return m.Version
	}
	return nil
}

func init() {
	proto.RegisterType((*VersionGetResponse)(nil), "appscode.apps.v1beta1.VersionGetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Version service

type VersionClient interface {
	Get(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*VersionGetResponse, error)
}

type versionClient struct {
	cc *grpc.ClientConn
}

func NewVersionClient(cc *grpc.ClientConn) VersionClient {
	return &versionClient{cc}
}

func (c *versionClient) Get(ctx context.Context, in *appscode_dtypes.VoidRequest, opts ...grpc.CallOption) (*VersionGetResponse, error) {
	out := new(VersionGetResponse)
	err := grpc.Invoke(ctx, "/appscode.apps.v1beta1.Version/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Version service

type VersionServer interface {
	Get(context.Context, *appscode_dtypes.VoidRequest) (*VersionGetResponse, error)
}

func RegisterVersionServer(s *grpc.Server, srv VersionServer) {
	s.RegisterService(&_Version_serviceDesc, srv)
}

func _Version_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(appscode_dtypes.VoidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VersionServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.apps.v1beta1.Version/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VersionServer).Get(ctx, req.(*appscode_dtypes.VoidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Version_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.apps.v1beta1.Version",
	HandlerType: (*VersionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Version_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "version.proto",
}

func init() { proto.RegisterFile("version.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x50, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x56, 0x8a, 0xd4, 0x20, 0x03, 0x8b, 0x25, 0x44, 0x1a, 0xca, 0x8f, 0xb2, 0x00, 0x8b, 0xad,
	0xb6, 0x0b, 0x73, 0x97, 0xae, 0x55, 0x90, 0x3a, 0xb0, 0xb9, 0xcd, 0xa9, 0xb2, 0x04, 0x3e, 0xd3,
	0xbb, 0x16, 0xc1, 0x98, 0x57, 0xe0, 0x59, 0x78, 0x12, 0x5e, 0x81, 0x07, 0x41, 0x4d, 0x1c, 0x20,
	0x2a, 0x2c, 0x96, 0x75, 0xdf, 0xdf, 0xdd, 0x27, 0x8e, 0x36, 0xb0, 0x22, 0x8b, 0x4e, 0xf9, 0x15,
	0x32, 0xca, 0x63, 0xe3, 0x3d, 0x2d, 0xb0, 0x00, 0xb5, 0xfd, 0xa8, 0xcd, 0x60, 0x0e, 0x6c, 0x06,
	0x69, 0x7f, 0x89, 0xb8, 0x7c, 0x00, 0x6d, 0xbc, 0xd5, 0xc6, 0x39, 0x64, 0xc3, 0x16, 0x1d, 0xd5,
	0xa2, 0xf4, 0xbc, 0x11, 0xfd, 0x83, 0x5f, 0xb4, 0xf0, 0x82, 0x5f, 0x3c, 0x90, 0xae, 0xde, 0x40,
	0xc8, 0x5a, 0x84, 0xb0, 0x91, 0x6e, 0x6d, 0x96, 0xbd, 0x0a, 0x39, 0xab, 0x07, 0x13, 0xe0, 0x1c,
	0xc8, 0xa3, 0x23, 0x90, 0x5a, 0x74, 0x89, 0x0d, 0xaf, 0x29, 0x89, 0x2e, 0xa3, 0xeb, 0x83, 0xe1,
	0x89, 0xfa, 0x3e, 0xa0, 0xce, 0x51, 0x77, 0x15, 0x9c, 0x07, 0x9a, 0x1c, 0x89, 0x38, 0xf8, 0x26,
	0x9d, 0x4a, 0xd1, 0xfb, 0x51, 0x34, 0x81, 0x21, 0x27, 0x6f, 0x98, 0xc3, 0x32, 0x12, 0x71, 0x18,
	0xca, 0x67, 0xb1, 0x37, 0x01, 0x96, 0xfd, 0x9d, 0xa0, 0x19, 0xda, 0x22, 0x87, 0xa7, 0x35, 0x10,
	0xa7, 0x37, 0xea, 0xcf, 0x1e, 0xd5, 0xee, 0x05, 0xd9, 0x55, 0xf9, 0x9e, 0x74, 0xf6, 0xa3, 0xf2,
	0xe3, 0xf3, 0xad, 0x73, 0x2a, 0x7b, 0x75, 0x85, 0xde, 0x93, 0x0e, 0x9a, 0xa6, 0x86, 0xf1, 0xad,
	0x38, 0x5b, 0xe0, 0xe3, 0x6f, 0x63, 0xdb, 0x32, 0x1f, 0x1f, 0x06, 0xf7, 0xe9, 0xb6, 0xaf, 0x69,
	0x74, 0x1f, 0x07, 0x60, 0xde, 0xad, 0x1a, 0x1c, 0x7d, 0x05, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x3b,
	0x16, 0xd8, 0xec, 0x01, 0x00, 0x00,
}
