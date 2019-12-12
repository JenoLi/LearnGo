// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

package grpcTest

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

type TestRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{0}
}

func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestRequest.Unmarshal(m, b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return xxx_messageInfo_TestRequest.Size(m)
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TestReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestReply) Reset()         { *m = TestReply{} }
func (m *TestReply) String() string { return proto.CompactTextString(m) }
func (*TestReply) ProtoMessage()    {}
func (*TestReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{1}
}

func (m *TestReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestReply.Unmarshal(m, b)
}
func (m *TestReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestReply.Marshal(b, m, deterministic)
}
func (m *TestReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestReply.Merge(m, src)
}
func (m *TestReply) XXX_Size() int {
	return xxx_messageInfo_TestReply.Size(m)
}
func (m *TestReply) XXX_DiscardUnknown() {
	xxx_messageInfo_TestReply.DiscardUnknown(m)
}

var xxx_messageInfo_TestReply proto.InternalMessageInfo

func (m *TestReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AddTarget struct {
	V1                   int32    `protobuf:"varint,1,opt,name=v1,proto3" json:"v1,omitempty"`
	V2                   int32    `protobuf:"varint,2,opt,name=v2,proto3" json:"v2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTarget) Reset()         { *m = AddTarget{} }
func (m *AddTarget) String() string { return proto.CompactTextString(m) }
func (*AddTarget) ProtoMessage()    {}
func (*AddTarget) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{2}
}

func (m *AddTarget) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTarget.Unmarshal(m, b)
}
func (m *AddTarget) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTarget.Marshal(b, m, deterministic)
}
func (m *AddTarget) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTarget.Merge(m, src)
}
func (m *AddTarget) XXX_Size() int {
	return xxx_messageInfo_AddTarget.Size(m)
}
func (m *AddTarget) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTarget.DiscardUnknown(m)
}

var xxx_messageInfo_AddTarget proto.InternalMessageInfo

func (m *AddTarget) GetV1() int32 {
	if m != nil {
		return m.V1
	}
	return 0
}

func (m *AddTarget) GetV2() int32 {
	if m != nil {
		return m.V2
	}
	return 0
}

type SumValue struct {
	Sum                  int32    `protobuf:"varint,1,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumValue) Reset()         { *m = SumValue{} }
func (m *SumValue) String() string { return proto.CompactTextString(m) }
func (*SumValue) ProtoMessage()    {}
func (*SumValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_c161fcfdc0c3ff1e, []int{3}
}

func (m *SumValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumValue.Unmarshal(m, b)
}
func (m *SumValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumValue.Marshal(b, m, deterministic)
}
func (m *SumValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumValue.Merge(m, src)
}
func (m *SumValue) XXX_Size() int {
	return xxx_messageInfo_SumValue.Size(m)
}
func (m *SumValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SumValue.DiscardUnknown(m)
}

var xxx_messageInfo_SumValue proto.InternalMessageInfo

func (m *SumValue) GetSum() int32 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "grpcTest.TestRequest")
	proto.RegisterType((*TestReply)(nil), "grpcTest.TestReply")
	proto.RegisterType((*AddTarget)(nil), "grpcTest.AddTarget")
	proto.RegisterType((*SumValue)(nil), "grpcTest.SumValue")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor_c161fcfdc0c3ff1e) }

var fileDescriptor_c161fcfdc0c3ff1e = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xcf, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x77, 0xbb, 0xab, 0xb6, 0x23, 0x88, 0x8c, 0x08, 0x65, 0xf1, 0xa0, 0x01, 0x41, 0x10,
	0x8a, 0x5b, 0x2f, 0x5e, 0x17, 0x2f, 0x1e, 0x25, 0xbb, 0x78, 0x8f, 0x66, 0xc8, 0x25, 0xb1, 0x31,
	0x3f, 0x0a, 0xfd, 0xef, 0xa5, 0xb1, 0xd1, 0xe2, 0x29, 0xef, 0xbd, 0x7c, 0x87, 0x8f, 0x01, 0x08,
	0xe4, 0x43, 0x63, 0x5d, 0x17, 0x3a, 0x2c, 0x95, 0xb3, 0x1f, 0x07, 0xf2, 0x81, 0xdd, 0xc0, 0xe9,
	0xf8, 0x72, 0xfa, 0x8a, 0xe4, 0x03, 0x22, 0xac, 0x3f, 0x85, 0xa1, 0x7a, 0x79, 0xbd, 0xbc, 0xab,
	0x78, 0xca, 0xec, 0x16, 0xaa, 0x1f, 0xc4, 0xea, 0x01, 0x6b, 0x38, 0x31, 0xe4, 0xbd, 0x50, 0x99,
	0xc9, 0x95, 0xdd, 0x43, 0xb5, 0x93, 0xf2, 0x20, 0x9c, 0xa2, 0x80, 0x67, 0x50, 0xf4, 0xdb, 0x44,
	0x1c, 0xf1, 0xa2, 0xdf, 0xa6, 0xde, 0xd6, 0xc5, 0xd4, 0x5b, 0x76, 0x05, 0xe5, 0x3e, 0x9a, 0x37,
	0xa1, 0x23, 0xe1, 0x39, 0xac, 0x7c, 0x34, 0x13, 0x3c, 0xc6, 0xd6, 0xc1, 0x5a, 0xf1, 0xd7, 0x67,
	0x7c, 0x82, 0x72, 0x2f, 0x86, 0x17, 0xd2, 0xba, 0xc3, 0xcb, 0x26, 0x3b, 0x37, 0x33, 0xe1, 0xcd,
	0xc5, 0xff, 0xd9, 0xea, 0x81, 0x2d, 0xf0, 0x01, 0x56, 0x3b, 0x29, 0x71, 0xf6, 0xfb, 0xeb, 0xb6,
	0xc1, 0xbf, 0x31, 0x3b, 0xb0, 0xc5, 0xfb, 0x71, 0xba, 0xcc, 0xe3, 0x77, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x56, 0xe4, 0x98, 0xac, 0x27, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GRPCClient is the client API for GRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GRPCClient interface {
	SayHello(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestReply, error)
	Add(ctx context.Context, in *AddTarget, opts ...grpc.CallOption) (*SumValue, error)
}

type gRPCClient struct {
	cc *grpc.ClientConn
}

func NewGRPCClient(cc *grpc.ClientConn) GRPCClient {
	return &gRPCClient{cc}
}

func (c *gRPCClient) SayHello(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestReply, error) {
	out := new(TestReply)
	err := c.cc.Invoke(ctx, "/grpcTest.gRPC/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCClient) Add(ctx context.Context, in *AddTarget, opts ...grpc.CallOption) (*SumValue, error) {
	out := new(SumValue)
	err := c.cc.Invoke(ctx, "/grpcTest.gRPC/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCServer is the server API for GRPC service.
type GRPCServer interface {
	SayHello(context.Context, *TestRequest) (*TestReply, error)
	Add(context.Context, *AddTarget) (*SumValue, error)
}

// UnimplementedGRPCServer can be embedded to have forward compatible implementations.
type UnimplementedGRPCServer struct {
}

func (*UnimplementedGRPCServer) SayHello(ctx context.Context, req *TestRequest) (*TestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedGRPCServer) Add(ctx context.Context, req *AddTarget) (*SumValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}

func RegisterGRPCServer(s *grpc.Server, srv GRPCServer) {
	s.RegisterService(&_GRPC_serviceDesc, srv)
}

func _GRPC_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcTest.gRPC/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServer).SayHello(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPC_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTarget)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcTest.gRPC/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCServer).Add(ctx, req.(*AddTarget))
	}
	return interceptor(ctx, in, info, handler)
}

var _GRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcTest.gRPC",
	HandlerType: (*GRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _GRPC_SayHello_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _GRPC_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}