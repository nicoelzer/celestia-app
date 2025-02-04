// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: qgb/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("qgb/query.proto", fileDescriptor_f3c1fd86445aad81) }

var fileDescriptor_f3c1fd86445aad81 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xce, 0x31, 0xaa, 0xc2, 0x40,
	0x10, 0x06, 0xe0, 0x84, 0xc7, 0x7b, 0x0f, 0xd2, 0x08, 0x96, 0x41, 0xf6, 0x00, 0x82, 0x19, 0xa2,
	0x37, 0xb0, 0xb3, 0xb4, 0xb5, 0x9b, 0x0d, 0xcb, 0xb8, 0x90, 0xec, 0x6c, 0xb2, 0x13, 0x31, 0xb7,
	0xf0, 0x58, 0x96, 0x29, 0x2d, 0x25, 0xb9, 0x88, 0x24, 0x01, 0xbb, 0xbf, 0xf8, 0xfe, 0x9f, 0x3f,
	0x59, 0xd5, 0xa4, 0xa1, 0x6e, 0x4d, 0xd3, 0x65, 0xbe, 0x61, 0xe1, 0xf5, 0x4f, 0x4d, 0x3a, 0xdd,
	0x10, 0x33, 0x95, 0x06, 0xd0, 0x5b, 0x40, 0xe7, 0x58, 0x50, 0x2c, 0xbb, 0xb0, 0x90, 0x74, 0x5b,
	0x70, 0xa8, 0x38, 0x80, 0xc6, 0x60, 0x96, 0x2e, 0xdc, 0x72, 0x6d, 0x04, 0x73, 0xf0, 0x48, 0xd6,
	0xcd, 0x78, 0xb1, 0xfb, 0xff, 0xe4, 0xf7, 0x3c, 0x89, 0xe3, 0xe9, 0x39, 0xa8, 0xb8, 0x1f, 0x54,
	0xfc, 0x1e, 0x54, 0xfc, 0x18, 0x55, 0xd4, 0x8f, 0x2a, 0x7a, 0x8d, 0x2a, 0xba, 0x00, 0x59, 0xb9,
	0xb6, 0x3a, 0x2b, 0xb8, 0x82, 0xc2, 0x94, 0x26, 0x88, 0x45, 0x6e, 0xe8, 0x9b, 0x77, 0xe8, 0x3d,
	0xdc, 0x61, 0x3a, 0x2a, 0x9d, 0x37, 0x41, 0xff, 0xcd, 0xd3, 0x87, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2e, 0xe6, 0x68, 0x3e, 0xbc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qgb.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "qgb/query.proto",
}
