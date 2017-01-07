// Code generated by protoc-gen-go.
// source: metdata.proto
// DO NOT EDIT!

/*
Package v1beta1 is a generated protocol buffer package.

It is generated from these files:
	metdata.proto

It has these top-level messages:
	MetadataGetRequest
	MetadataGetResponse
	GitMetadata
	DockerMetadata
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

type MetadataGetRequest struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Type     string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Registry string `protobuf:"bytes,3,opt,name=registry" json:"registry,omitempty"`
}

func (m *MetadataGetRequest) Reset()                    { *m = MetadataGetRequest{} }
func (m *MetadataGetRequest) String() string            { return proto.CompactTextString(m) }
func (*MetadataGetRequest) ProtoMessage()               {}
func (*MetadataGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *MetadataGetRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetadataGetRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *MetadataGetRequest) GetRegistry() string {
	if m != nil {
		return m.Registry
	}
	return ""
}

type MetadataGetResponse struct {
	Status *appscode_dtypes.Status `protobuf:"bytes,1,opt,name=status" json:"status,omitempty"`
	// Types that are valid to be assigned to Metdata:
	//	*MetadataGetResponse_Git
	//	*MetadataGetResponse_Docker
	Metdata isMetadataGetResponse_Metdata `protobuf_oneof:"metdata"`
}

func (m *MetadataGetResponse) Reset()                    { *m = MetadataGetResponse{} }
func (m *MetadataGetResponse) String() string            { return proto.CompactTextString(m) }
func (*MetadataGetResponse) ProtoMessage()               {}
func (*MetadataGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isMetadataGetResponse_Metdata interface {
	isMetadataGetResponse_Metdata()
}

type MetadataGetResponse_Git struct {
	Git *GitMetadata `protobuf:"bytes,2,opt,name=git,oneof"`
}
type MetadataGetResponse_Docker struct {
	Docker *DockerMetadata `protobuf:"bytes,3,opt,name=docker,oneof"`
}

func (*MetadataGetResponse_Git) isMetadataGetResponse_Metdata()    {}
func (*MetadataGetResponse_Docker) isMetadataGetResponse_Metdata() {}

func (m *MetadataGetResponse) GetMetdata() isMetadataGetResponse_Metdata {
	if m != nil {
		return m.Metdata
	}
	return nil
}

func (m *MetadataGetResponse) GetStatus() *appscode_dtypes.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *MetadataGetResponse) GetGit() *GitMetadata {
	if x, ok := m.GetMetdata().(*MetadataGetResponse_Git); ok {
		return x.Git
	}
	return nil
}

func (m *MetadataGetResponse) GetDocker() *DockerMetadata {
	if x, ok := m.GetMetdata().(*MetadataGetResponse_Docker); ok {
		return x.Docker
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*MetadataGetResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _MetadataGetResponse_OneofMarshaler, _MetadataGetResponse_OneofUnmarshaler, _MetadataGetResponse_OneofSizer, []interface{}{
		(*MetadataGetResponse_Git)(nil),
		(*MetadataGetResponse_Docker)(nil),
	}
}

func _MetadataGetResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*MetadataGetResponse)
	// metdata
	switch x := m.Metdata.(type) {
	case *MetadataGetResponse_Git:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Git); err != nil {
			return err
		}
	case *MetadataGetResponse_Docker:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Docker); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("MetadataGetResponse.Metdata has unexpected type %T", x)
	}
	return nil
}

func _MetadataGetResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*MetadataGetResponse)
	switch tag {
	case 2: // metdata.git
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GitMetadata)
		err := b.DecodeMessage(msg)
		m.Metdata = &MetadataGetResponse_Git{msg}
		return true, err
	case 3: // metdata.docker
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DockerMetadata)
		err := b.DecodeMessage(msg)
		m.Metdata = &MetadataGetResponse_Docker{msg}
		return true, err
	default:
		return false, nil
	}
}

func _MetadataGetResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*MetadataGetResponse)
	// metdata
	switch x := m.Metdata.(type) {
	case *MetadataGetResponse_Git:
		s := proto.Size(x.Git)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *MetadataGetResponse_Docker:
		s := proto.Size(x.Docker)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GitMetadata struct {
	Name     string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Branches []string `protobuf:"bytes,2,rep,name=branches" json:"branches,omitempty"`
	Notes    []string `protobuf:"bytes,3,rep,name=notes" json:"notes,omitempty"`
	Tags     []string `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
}

func (m *GitMetadata) Reset()                    { *m = GitMetadata{} }
func (m *GitMetadata) String() string            { return proto.CompactTextString(m) }
func (*GitMetadata) ProtoMessage()               {}
func (*GitMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GitMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GitMetadata) GetBranches() []string {
	if m != nil {
		return m.Branches
	}
	return nil
}

func (m *GitMetadata) GetNotes() []string {
	if m != nil {
		return m.Notes
	}
	return nil
}

func (m *GitMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type DockerMetadata struct {
	Name string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags []string `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty"`
}

func (m *DockerMetadata) Reset()                    { *m = DockerMetadata{} }
func (m *DockerMetadata) String() string            { return proto.CompactTextString(m) }
func (*DockerMetadata) ProtoMessage()               {}
func (*DockerMetadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DockerMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DockerMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*MetadataGetRequest)(nil), "appscode.apps.v1beta1.MetadataGetRequest")
	proto.RegisterType((*MetadataGetResponse)(nil), "appscode.apps.v1beta1.MetadataGetResponse")
	proto.RegisterType((*GitMetadata)(nil), "appscode.apps.v1beta1.GitMetadata")
	proto.RegisterType((*DockerMetadata)(nil), "appscode.apps.v1beta1.DockerMetadata")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Metadata service

type MetadataClient interface {
	Get(ctx context.Context, in *MetadataGetRequest, opts ...grpc.CallOption) (*MetadataGetResponse, error)
}

type metadataClient struct {
	cc *grpc.ClientConn
}

func NewMetadataClient(cc *grpc.ClientConn) MetadataClient {
	return &metadataClient{cc}
}

func (c *metadataClient) Get(ctx context.Context, in *MetadataGetRequest, opts ...grpc.CallOption) (*MetadataGetResponse, error) {
	out := new(MetadataGetResponse)
	err := grpc.Invoke(ctx, "/appscode.apps.v1beta1.Metadata/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Metadata service

type MetadataServer interface {
	Get(context.Context, *MetadataGetRequest) (*MetadataGetResponse, error)
}

func RegisterMetadataServer(s *grpc.Server, srv MetadataServer) {
	s.RegisterService(&_Metadata_serviceDesc, srv)
}

func _Metadata_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetadataGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetadataServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/appscode.apps.v1beta1.Metadata/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetadataServer).Get(ctx, req.(*MetadataGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metadata_serviceDesc = grpc.ServiceDesc{
	ServiceName: "appscode.apps.v1beta1.Metadata",
	HandlerType: (*MetadataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Metadata_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metdata.proto",
}

func init() { proto.RegisterFile("metdata.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0xcd, 0x8a, 0xd4, 0x40,
	0x10, 0xc7, 0x4d, 0xb2, 0xce, 0x66, 0x2b, 0xac, 0x87, 0x56, 0x31, 0x04, 0x3f, 0xd6, 0x88, 0xa0,
	0x1e, 0xd2, 0x4c, 0x04, 0xd1, 0x93, 0x30, 0x08, 0xe3, 0x45, 0x18, 0xe2, 0x45, 0xbc, 0xf5, 0x24,
	0x45, 0x0c, 0x9a, 0x74, 0x9b, 0xae, 0x11, 0x06, 0xf1, 0x32, 0x77, 0x4f, 0xde, 0x7c, 0x0f, 0x5f,
	0xc2, 0xab, 0xaf, 0xe0, 0x83, 0x48, 0x77, 0x92, 0x71, 0x82, 0x13, 0xf0, 0x12, 0xea, 0xeb, 0x97,
	0x7f, 0x75, 0x55, 0xc1, 0x79, 0x8d, 0x54, 0x08, 0x12, 0x89, 0x6a, 0x25, 0x49, 0x76, 0x5d, 0x28,
	0xa5, 0x73, 0x59, 0x60, 0x62, 0x8c, 0xe4, 0xd3, 0x7c, 0x8d, 0x24, 0xe6, 0xd1, 0xcd, 0x52, 0xca,
	0xf2, 0x03, 0x72, 0xa1, 0x2a, 0x2e, 0x9a, 0x46, 0x92, 0xa0, 0x4a, 0x36, 0xba, 0x83, 0xa2, 0xdb,
	0x03, 0x34, 0x91, 0xbf, 0x33, 0xca, 0x17, 0xb4, 0x55, 0xa8, 0xb9, 0xfd, 0x76, 0x05, 0xf1, 0x1b,
	0x60, 0xaf, 0x90, 0x84, 0xe9, 0x63, 0x89, 0x94, 0xe1, 0xc7, 0x0d, 0x6a, 0x62, 0x0c, 0x4e, 0x1a,
	0x51, 0x63, 0xe8, 0x5c, 0x38, 0x0f, 0xce, 0x32, 0x6b, 0x9b, 0x98, 0x01, 0x43, 0xb7, 0x8b, 0x19,
	0x9b, 0x45, 0xe0, 0xb7, 0x58, 0x56, 0x9a, 0xda, 0x6d, 0xe8, 0xd9, 0xf8, 0xde, 0x8f, 0x7f, 0x3a,
	0x70, 0x75, 0xf4, 0x6b, 0xad, 0x64, 0xa3, 0x91, 0x71, 0x98, 0x69, 0x12, 0xb4, 0xd1, 0xf6, 0xef,
	0x41, 0x7a, 0x23, 0xd9, 0x3f, 0xbc, 0xeb, 0x2f, 0x79, 0x6d, 0xd3, 0x59, 0x5f, 0xc6, 0x9e, 0x80,
	0x57, 0x56, 0x64, 0x75, 0x83, 0x34, 0x4e, 0x8e, 0x8e, 0x29, 0x59, 0x56, 0x34, 0x88, 0xbd, 0xbc,
	0x94, 0x19, 0x80, 0x3d, 0x87, 0x59, 0x21, 0xf3, 0xf7, 0xd8, 0xda, 0xd6, 0x82, 0xf4, 0xfe, 0x04,
	0xfa, 0xc2, 0x16, 0x1d, 0xd0, 0x3d, 0xb6, 0x38, 0x83, 0xd3, 0x7e, 0x45, 0x71, 0x09, 0xc1, 0x81,
	0xc2, 0xd1, 0xf9, 0x44, 0xe0, 0xaf, 0x5b, 0xd1, 0xe4, 0xef, 0x50, 0x87, 0xee, 0x85, 0x67, 0x66,
	0x31, 0xf8, 0xec, 0x1a, 0x5c, 0x6e, 0x24, 0xa1, 0x0e, 0x3d, 0x9b, 0xe8, 0x1c, 0x3b, 0x51, 0x51,
	0xea, 0xf0, 0xc4, 0x06, 0xad, 0x1d, 0x3f, 0x85, 0x2b, 0xe3, 0x7e, 0x26, 0x77, 0x61, 0x48, 0xf7,
	0x2f, 0x99, 0x7e, 0x77, 0xc0, 0xdf, 0x43, 0x5f, 0x1d, 0xf0, 0x96, 0x48, 0xec, 0xe1, 0xc4, 0x9b,
	0xff, 0xdd, 0x79, 0xf4, 0xe8, 0x7f, 0x4a, 0xbb, 0x1d, 0xc6, 0x7c, 0xf7, 0x23, 0x74, 0x7d, 0x67,
	0xf7, 0xeb, 0xf7, 0x37, 0xf7, 0x1e, 0xbb, 0xdb, 0x1d, 0x9f, 0x52, 0x9a, 0xf7, 0x10, 0xaf, 0x7b,
	0x88, 0x7f, 0x36, 0x2b, 0xfd, 0xb2, 0x78, 0x06, 0xb7, 0x72, 0x59, 0x1f, 0x2a, 0x54, 0x23, 0x95,
	0xc5, 0xf9, 0x20, 0xb3, 0x32, 0x67, 0xb9, 0x72, 0xde, 0x9e, 0xf6, 0x99, 0xf5, 0xcc, 0x1e, 0xea,
	0xe3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x76, 0xd0, 0xc1, 0xc6, 0x2f, 0x03, 0x00, 0x00,
}
