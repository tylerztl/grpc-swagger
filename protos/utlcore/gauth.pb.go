// Code generated by protoc-gen-go. DO NOT EDIT.
// source: utlcore/gauth.proto

package utlcore // import "grpc-swagger/protos/utlcore"

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

type GASecretRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GASecretRequest) Reset()         { *m = GASecretRequest{} }
func (m *GASecretRequest) String() string { return proto.CompactTextString(m) }
func (*GASecretRequest) ProtoMessage()    {}
func (*GASecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gauth_c89b2beedb07b7d7, []int{0}
}
func (m *GASecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GASecretRequest.Unmarshal(m, b)
}
func (m *GASecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GASecretRequest.Marshal(b, m, deterministic)
}
func (dst *GASecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GASecretRequest.Merge(dst, src)
}
func (m *GASecretRequest) XXX_Size() int {
	return xxx_messageInfo_GASecretRequest.Size(m)
}
func (m *GASecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GASecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GASecretRequest proto.InternalMessageInfo

type GASecretResponse struct {
	Secret               string   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GASecretResponse) Reset()         { *m = GASecretResponse{} }
func (m *GASecretResponse) String() string { return proto.CompactTextString(m) }
func (*GASecretResponse) ProtoMessage()    {}
func (*GASecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_gauth_c89b2beedb07b7d7, []int{1}
}
func (m *GASecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GASecretResponse.Unmarshal(m, b)
}
func (m *GASecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GASecretResponse.Marshal(b, m, deterministic)
}
func (dst *GASecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GASecretResponse.Merge(dst, src)
}
func (m *GASecretResponse) XXX_Size() int {
	return xxx_messageInfo_GASecretResponse.Size(m)
}
func (m *GASecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GASecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GASecretResponse proto.InternalMessageInfo

func (m *GASecretResponse) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type GAVerifyRequest struct {
	Secret               string   `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	AuthCode             string   `protobuf:"bytes,2,opt,name=auth_code,json=authCode,proto3" json:"auth_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GAVerifyRequest) Reset()         { *m = GAVerifyRequest{} }
func (m *GAVerifyRequest) String() string { return proto.CompactTextString(m) }
func (*GAVerifyRequest) ProtoMessage()    {}
func (*GAVerifyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_gauth_c89b2beedb07b7d7, []int{2}
}
func (m *GAVerifyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GAVerifyRequest.Unmarshal(m, b)
}
func (m *GAVerifyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GAVerifyRequest.Marshal(b, m, deterministic)
}
func (dst *GAVerifyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GAVerifyRequest.Merge(dst, src)
}
func (m *GAVerifyRequest) XXX_Size() int {
	return xxx_messageInfo_GAVerifyRequest.Size(m)
}
func (m *GAVerifyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GAVerifyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GAVerifyRequest proto.InternalMessageInfo

func (m *GAVerifyRequest) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *GAVerifyRequest) GetAuthCode() string {
	if m != nil {
		return m.AuthCode
	}
	return ""
}

func init() {
	proto.RegisterType((*GASecretRequest)(nil), "protos.GASecretRequest")
	proto.RegisterType((*GASecretResponse)(nil), "protos.GASecretResponse")
	proto.RegisterType((*GAVerifyRequest)(nil), "protos.GAVerifyRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GAuthClient is the client API for GAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GAuthClient interface {
	GetGASecret(ctx context.Context, in *GASecretRequest, opts ...grpc.CallOption) (*GASecretResponse, error)
	VerifyGA(ctx context.Context, in *GAVerifyRequest, opts ...grpc.CallOption) (*ServerStatus, error)
}

type gAuthClient struct {
	cc *grpc.ClientConn
}

func NewGAuthClient(cc *grpc.ClientConn) GAuthClient {
	return &gAuthClient{cc}
}

func (c *gAuthClient) GetGASecret(ctx context.Context, in *GASecretRequest, opts ...grpc.CallOption) (*GASecretResponse, error) {
	out := new(GASecretResponse)
	err := c.cc.Invoke(ctx, "/protos.GAuth/GetGASecret", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gAuthClient) VerifyGA(ctx context.Context, in *GAVerifyRequest, opts ...grpc.CallOption) (*ServerStatus, error) {
	out := new(ServerStatus)
	err := c.cc.Invoke(ctx, "/protos.GAuth/VerifyGA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GAuthServer is the server API for GAuth service.
type GAuthServer interface {
	GetGASecret(context.Context, *GASecretRequest) (*GASecretResponse, error)
	VerifyGA(context.Context, *GAVerifyRequest) (*ServerStatus, error)
}

func RegisterGAuthServer(s *grpc.Server, srv GAuthServer) {
	s.RegisterService(&_GAuth_serviceDesc, srv)
}

func _GAuth_GetGASecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GASecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GAuthServer).GetGASecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GAuth/GetGASecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GAuthServer).GetGASecret(ctx, req.(*GASecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GAuth_VerifyGA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GAVerifyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GAuthServer).VerifyGA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.GAuth/VerifyGA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GAuthServer).VerifyGA(ctx, req.(*GAVerifyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GAuth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.GAuth",
	HandlerType: (*GAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGASecret",
			Handler:    _GAuth_GetGASecret_Handler,
		},
		{
			MethodName: "VerifyGA",
			Handler:    _GAuth_VerifyGA_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "utlcore/gauth.proto",
}

func init() { proto.RegisterFile("utlcore/gauth.proto", fileDescriptor_gauth_c89b2beedb07b7d7) }

var fileDescriptor_gauth_c89b2beedb07b7d7 = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0x9b, 0x17, 0xde, 0xd2, 0x8e, 0x07, 0x75, 0x2d, 0x1a, 0xd2, 0x8b, 0xe4, 0x24, 0x82,
	0x09, 0xe8, 0xd1, 0x53, 0x22, 0x36, 0x27, 0x41, 0x1a, 0xf0, 0xe0, 0x45, 0xb6, 0xdb, 0x71, 0x2b,
	0xda, 0xce, 0xba, 0x3b, 0xab, 0xf8, 0x0d, 0xfc, 0xd8, 0x92, 0x6c, 0x82, 0xff, 0xf0, 0xb4, 0xec,
	0xf3, 0x7b, 0x66, 0xf8, 0x31, 0xb0, 0xe7, 0xf9, 0x49, 0x91, 0xc5, 0x5c, 0x4b, 0xcf, 0xab, 0xcc,
	0x58, 0x62, 0x12, 0xc3, 0xf6, 0x71, 0xc9, 0xa4, 0x87, 0x8a, 0xd6, 0x6b, 0xda, 0x04, 0x9a, 0xee,
	0xc2, 0x76, 0x55, 0xd4, 0xa8, 0x2c, 0xf2, 0x1c, 0x9f, 0x3d, 0x3a, 0x4e, 0x8f, 0x61, 0xe7, 0x33,
	0x72, 0x86, 0x36, 0x0e, 0xc5, 0x3e, 0x0c, 0x5d, 0x9b, 0xc4, 0xd1, 0x61, 0x74, 0x34, 0x9e, 0x77,
	0xbf, 0x74, 0xd6, 0x8c, 0xdf, 0xa0, 0x7d, 0xb8, 0x7f, 0xeb, 0xc6, 0xff, 0xaa, 0x8a, 0x29, 0x8c,
	0x1b, 0xab, 0x3b, 0x45, 0x4b, 0x8c, 0xff, 0xb5, 0x68, 0xd4, 0x04, 0x17, 0xb4, 0xc4, 0xd3, 0xf7,
	0x08, 0xfe, 0x57, 0x85, 0xe7, 0x95, 0x28, 0x61, 0xab, 0x42, 0xee, 0x05, 0xc4, 0x41, 0xf0, 0x74,
	0xd9, 0x0f, 0xcb, 0x24, 0xfe, 0x0d, 0x82, 0x6b, 0x3a, 0x10, 0xe7, 0x30, 0x0a, 0x4e, 0x55, 0xf1,
	0x75, 0xc1, 0x37, 0xcf, 0x64, 0xd2, 0x83, 0x1a, 0xed, 0x0b, 0xda, 0x9a, 0x25, 0x7b, 0x97, 0x0e,
	0xca, 0x4b, 0x48, 0xc8, 0xea, 0x6c, 0x41, 0xc4, 0xd2, 0x98, 0xbe, 0xd4, 0x1d, 0xaf, 0x84, 0xab,
	0x59, 0x71, 0x2d, 0xd5, 0xa3, 0xd4, 0x78, 0x3b, 0xd5, 0xd6, 0xa8, 0x13, 0xf7, 0x2a, 0xb5, 0x46,
	0x9b, 0x87, 0x62, 0xde, 0x15, 0x17, 0xe1, 0xec, 0x67, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x95,
	0x59, 0xa6, 0x90, 0x94, 0x01, 0x00, 0x00,
}
