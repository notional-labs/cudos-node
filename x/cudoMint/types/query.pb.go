// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cudoventures/cudosnode/cudoMint/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

func init() {
	proto.RegisterFile("cudoventures/cudosnode/cudoMint/query.proto", fileDescriptor_385d66ab754b82f2)
}

var fileDescriptor_385d66ab754b82f2 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0xcf, 0xb1, 0x4e, 0x85, 0x30,
	0x14, 0xc6, 0x71, 0x18, 0xd4, 0x84, 0xd1, 0x91, 0x98, 0xba, 0x6b, 0xec, 0x09, 0xf2, 0x06, 0x3a,
	0x33, 0xb8, 0x38, 0xb8, 0xb5, 0x70, 0x52, 0x9b, 0x48, 0x4f, 0xa5, 0xa7, 0x44, 0xde, 0xc2, 0xc7,
	0x72, 0x64, 0x74, 0x34, 0xf0, 0x22, 0x86, 0x12, 0xbd, 0xf7, 0x6e, 0x6d, 0xf2, 0x3b, 0x5f, 0xf2,
	0x2f, 0x6e, 0xdb, 0xd8, 0xd1, 0x88, 0x8e, 0xe3, 0x80, 0x01, 0xb6, 0x4f, 0x70, 0xd4, 0x61, 0x7a,
	0x35, 0xd6, 0x31, 0xbc, 0x47, 0x1c, 0x26, 0xe9, 0x07, 0x62, 0xba, 0xbc, 0x3e, 0xc6, 0xf2, 0x1f,
	0xcb, 0x3f, 0x5c, 0x5e, 0x19, 0x22, 0xf3, 0x86, 0xa0, 0xbc, 0x05, 0xe5, 0x1c, 0xb1, 0x62, 0x4b,
	0x2e, 0xec, 0xe7, 0xe5, 0x4d, 0x4b, 0xa1, 0xa7, 0x00, 0x5a, 0x05, 0xdc, 0x77, 0x61, 0xac, 0x34,
	0xb2, 0xaa, 0xc0, 0x2b, 0x63, 0x5d, 0xc2, 0xbb, 0xbd, 0xbf, 0x28, 0xce, 0x9e, 0x36, 0xf1, 0xd0,
	0x7c, 0x2d, 0x22, 0x9f, 0x17, 0x91, 0xff, 0x2c, 0x22, 0xff, 0x5c, 0x45, 0x36, 0xaf, 0x22, 0xfb,
	0x5e, 0x45, 0xf6, 0x52, 0x1b, 0xcb, 0xaf, 0x51, 0xcb, 0x96, 0x7a, 0x78, 0x8c, 0x1d, 0x3d, 0x9f,
	0x54, 0xdc, 0xa5, 0x8c, 0x8f, 0x43, 0x08, 0x4f, 0x1e, 0x83, 0x3e, 0x4f, 0xf3, 0xf5, 0x6f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x7d, 0x92, 0xb3, 0x4a, 0xf8, 0x00, 0x00, 0x00,
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
	ServiceName: "cudoventures.cudosnode.cudoMint.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "cudoventures/cudosnode/cudoMint/query.proto",
}
