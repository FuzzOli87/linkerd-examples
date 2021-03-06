// Code generated by protoc-gen-go.
// source: proto/word.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WordRequest struct {
}

func (m *WordRequest) Reset()                    { *m = WordRequest{} }
func (m *WordRequest) String() string            { return proto1.CompactTextString(m) }
func (*WordRequest) ProtoMessage()               {}
func (*WordRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type WordResponse struct {
	Word string `protobuf:"bytes,1,opt,name=word" json:"word,omitempty"`
}

func (m *WordResponse) Reset()                    { *m = WordResponse{} }
func (m *WordResponse) String() string            { return proto1.CompactTextString(m) }
func (*WordResponse) ProtoMessage()               {}
func (*WordResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *WordResponse) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func init() {
	proto1.RegisterType((*WordRequest)(nil), "proto.WordRequest")
	proto1.RegisterType((*WordResponse)(nil), "proto.WordResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for WordSvc service

type WordSvcClient interface {
	GetWord(ctx context.Context, in *WordRequest, opts ...grpc.CallOption) (*WordResponse, error)
}

type wordSvcClient struct {
	cc *grpc.ClientConn
}

func NewWordSvcClient(cc *grpc.ClientConn) WordSvcClient {
	return &wordSvcClient{cc}
}

func (c *wordSvcClient) GetWord(ctx context.Context, in *WordRequest, opts ...grpc.CallOption) (*WordResponse, error) {
	out := new(WordResponse)
	err := grpc.Invoke(ctx, "/proto.WordSvc/GetWord", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for WordSvc service

type WordSvcServer interface {
	GetWord(context.Context, *WordRequest) (*WordResponse, error)
}

func RegisterWordSvcServer(s *grpc.Server, srv WordSvcServer) {
	s.RegisterService(&_WordSvc_serviceDesc, srv)
}

func _WordSvc_GetWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WordSvcServer).GetWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.WordSvc/GetWord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WordSvcServer).GetWord(ctx, req.(*WordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WordSvc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.WordSvc",
	HandlerType: (*WordSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWord",
			Handler:    _WordSvc_GetWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/word.proto",
}

func init() { proto1.RegisterFile("proto/word.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xcf, 0x2f, 0x4a, 0xd1, 0x03, 0x33, 0x85, 0x58, 0xc1, 0x94, 0x12, 0x2f, 0x17,
	0x77, 0x78, 0x7e, 0x51, 0x4a, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x92, 0x12, 0x17, 0x0f,
	0x84, 0x5b, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x2a, 0x24, 0xc4, 0xc5, 0x02, 0xd2, 0x23, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x1b, 0xd9, 0x73, 0xb1, 0x83, 0xd4, 0x04, 0x97, 0x25, 0x0b,
	0x99, 0x70, 0xb1, 0xbb, 0xa7, 0x96, 0x80, 0x78, 0x42, 0x42, 0x10, 0x73, 0xf5, 0x90, 0x4c, 0x93,
	0x12, 0x46, 0x11, 0x83, 0x18, 0xa9, 0xc4, 0x90, 0xc4, 0x06, 0x16, 0x35, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x0c, 0x09, 0x5d, 0xd3, 0x95, 0x00, 0x00, 0x00,
}
