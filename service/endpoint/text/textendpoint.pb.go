// Code generated by protoc-gen-go.
// source: textendpoint.proto
// DO NOT EDIT!

/*
Package textendpoint is a generated protocol buffer package.

It is generated from these files:
	textendpoint.proto

It has these top-level messages:
	StreamTextRequest
	StreamTextResponse
	StreamTextResponseData
*/
package textendpoint

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

type StreamTextRequest struct {
	Echo      bool   `protobuf:"varint,1,opt,name=Echo,json=echo" json:"Echo,omitempty"`
	Input     string `protobuf:"bytes,2,opt,name=Input,json=input" json:"Input,omitempty"`
	SessionID string `protobuf:"bytes,3,opt,name=SessionID,json=sessionID" json:"SessionID,omitempty"`
}

func (m *StreamTextRequest) Reset()                    { *m = StreamTextRequest{} }
func (m *StreamTextRequest) String() string            { return proto.CompactTextString(m) }
func (*StreamTextRequest) ProtoMessage()               {}
func (*StreamTextRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StreamTextResponse struct {
	Code string                  `protobuf:"bytes,1,opt,name=Code,json=code" json:"Code,omitempty"`
	Data *StreamTextResponseData `protobuf:"bytes,2,opt,name=Data,json=data" json:"Data,omitempty"`
	Text string                  `protobuf:"bytes,3,opt,name=Text,json=text" json:"Text,omitempty"`
}

func (m *StreamTextResponse) Reset()                    { *m = StreamTextResponse{} }
func (m *StreamTextResponse) String() string            { return proto.CompactTextString(m) }
func (*StreamTextResponse) ProtoMessage()               {}
func (*StreamTextResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StreamTextResponse) GetData() *StreamTextResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type StreamTextResponseData struct {
	Output string `protobuf:"bytes,1,opt,name=Output,json=output" json:"Output,omitempty"`
}

func (m *StreamTextResponseData) Reset()                    { *m = StreamTextResponseData{} }
func (m *StreamTextResponseData) String() string            { return proto.CompactTextString(m) }
func (*StreamTextResponseData) ProtoMessage()               {}
func (*StreamTextResponseData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*StreamTextRequest)(nil), "api.StreamTextRequest")
	proto.RegisterType((*StreamTextResponse)(nil), "api.StreamTextResponse")
	proto.RegisterType((*StreamTextResponseData)(nil), "api.StreamTextResponseData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TextInterface service

type TextInterfaceClient interface {
	StreamText(ctx context.Context, opts ...grpc.CallOption) (TextInterface_StreamTextClient, error)
}

type textInterfaceClient struct {
	cc *grpc.ClientConn
}

func NewTextInterfaceClient(cc *grpc.ClientConn) TextInterfaceClient {
	return &textInterfaceClient{cc}
}

func (c *textInterfaceClient) StreamText(ctx context.Context, opts ...grpc.CallOption) (TextInterface_StreamTextClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TextInterface_serviceDesc.Streams[0], c.cc, "/api.TextInterface/StreamText", opts...)
	if err != nil {
		return nil, err
	}
	x := &textInterfaceStreamTextClient{stream}
	return x, nil
}

type TextInterface_StreamTextClient interface {
	Send(*StreamTextRequest) error
	Recv() (*StreamTextResponse, error)
	grpc.ClientStream
}

type textInterfaceStreamTextClient struct {
	grpc.ClientStream
}

func (x *textInterfaceStreamTextClient) Send(m *StreamTextRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *textInterfaceStreamTextClient) Recv() (*StreamTextResponse, error) {
	m := new(StreamTextResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TextInterface service

type TextInterfaceServer interface {
	StreamText(TextInterface_StreamTextServer) error
}

func RegisterTextInterfaceServer(s *grpc.Server, srv TextInterfaceServer) {
	s.RegisterService(&_TextInterface_serviceDesc, srv)
}

func _TextInterface_StreamText_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TextInterfaceServer).StreamText(&textInterfaceStreamTextServer{stream})
}

type TextInterface_StreamTextServer interface {
	Send(*StreamTextResponse) error
	Recv() (*StreamTextRequest, error)
	grpc.ServerStream
}

type textInterfaceStreamTextServer struct {
	grpc.ServerStream
}

func (x *textInterfaceStreamTextServer) Send(m *StreamTextResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *textInterfaceStreamTextServer) Recv() (*StreamTextRequest, error) {
	m := new(StreamTextRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TextInterface_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.TextInterface",
	HandlerType: (*TextInterfaceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamText",
			Handler:       _TextInterface_StreamText_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "textendpoint.proto",
}

func init() { proto.RegisterFile("textendpoint.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0x31, 0x4f, 0xfb, 0x30,
	0x10, 0xc5, 0xff, 0xf9, 0xd7, 0x8d, 0xc8, 0x21, 0x06, 0x4e, 0x28, 0x44, 0xc0, 0x50, 0x65, 0xca,
	0x14, 0xaa, 0xf2, 0x11, 0x5a, 0x86, 0x4c, 0x48, 0x6e, 0x37, 0x26, 0x93, 0x1c, 0xaa, 0x87, 0xfa,
	0x4c, 0x7c, 0x91, 0xfa, 0xf1, 0x91, 0x0d, 0x08, 0xa4, 0xc2, 0x78, 0xbe, 0x7b, 0xbf, 0xf7, 0xfc,
	0x00, 0x85, 0x8e, 0x42, 0x6e, 0xf0, 0x6c, 0x9d, 0xb4, 0x7e, 0x64, 0x61, 0x9c, 0x19, 0x6f, 0xeb,
	0x67, 0xb8, 0xdc, 0xca, 0x48, 0xe6, 0xb0, 0xa3, 0xa3, 0x68, 0x7a, 0x9b, 0x28, 0x08, 0x22, 0xa8,
	0xc7, 0x7e, 0xcf, 0x55, 0xb6, 0xc8, 0x9a, 0x33, 0xad, 0xa8, 0xdf, 0x33, 0x5e, 0xc1, 0xbc, 0x73,
	0x7e, 0x92, 0xea, 0xff, 0x22, 0x6b, 0x0a, 0x3d, 0xb7, 0x71, 0xc0, 0x3b, 0x28, 0xb6, 0x14, 0x82,
	0x65, 0xd7, 0x6d, 0xaa, 0x59, 0xda, 0x14, 0xe1, 0xeb, 0xa1, 0x3e, 0x00, 0xfe, 0x84, 0x07, 0xcf,
	0x2e, 0x50, 0xa4, 0xaf, 0x79, 0xa0, 0x44, 0x2f, 0xb4, 0xea, 0x79, 0x20, 0xbc, 0x07, 0xb5, 0x31,
	0x62, 0x12, 0xfc, 0x7c, 0x75, 0xdb, 0x1a, 0x6f, 0xdb, 0x53, 0x69, 0x3c, 0xd1, 0x6a, 0x30, 0x62,
	0x22, 0x24, 0x6e, 0x3e, 0x3d, 0x55, 0xfc, 0x5e, 0xbd, 0x84, 0xf2, 0x77, 0x0d, 0x96, 0x90, 0x3f,
	0x4d, 0x12, 0xd3, 0x7f, 0x98, 0xe6, 0x9c, 0xa6, 0xd5, 0x0e, 0x2e, 0xe2, 0x6d, 0xe7, 0x84, 0xc6,
	0x57, 0xd3, 0x13, 0xae, 0x01, 0xbe, 0x11, 0x58, 0x9e, 0xe4, 0x48, 0xfd, 0xdc, 0x5c, 0xff, 0x91,
	0xaf, 0xfe, 0xd7, 0x64, 0xcb, 0xec, 0x25, 0x4f, 0xfd, 0x3e, 0xbc, 0x07, 0x00, 0x00, 0xff, 0xff,
	0x04, 0x13, 0x99, 0x09, 0x75, 0x01, 0x00, 0x00,
}
