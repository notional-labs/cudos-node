// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/msg/v1/msg.proto

package msgservice

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

var E_Signer = &proto.ExtensionDesc{
	ExtendedType:  (*descriptorpb.MessageOptions)(nil),
	ExtensionType: ([]string)(nil),
	Field:         11110000,
	Name:          "cosmos.msg.v1.signer",
	Tag:           "bytes,11110000,rep,name=signer",
	Filename:      "cosmos/msg/v1/msg.proto",
}

func init() {
	proto.RegisterExtension(E_Signer)
}

func init() { proto.RegisterFile("cosmos/msg/v1/msg.proto", fileDescriptor_5c08b83ea858d203) }

var fileDescriptor_5c08b83ea858d203 = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0xcf, 0x2d, 0x4e, 0xd7, 0x2f, 0x33, 0x04, 0x51, 0x7a, 0x05, 0x45, 0xf9, 0x25,
	0xf9, 0x42, 0xbc, 0x10, 0x09, 0x3d, 0x90, 0x48, 0x99, 0xa1, 0x94, 0x42, 0x7a, 0x7e, 0x7e, 0x7a,
	0x4e, 0xaa, 0x3e, 0x58, 0x32, 0xa9, 0x34, 0x4d, 0x3f, 0x25, 0xb5, 0x38, 0xb9, 0x28, 0xb3, 0xa0,
	0x24, 0xbf, 0x08, 0xa2, 0xc1, 0xca, 0x8a, 0x8b, 0xad, 0x38, 0x33, 0x3d, 0x2f, 0xb5, 0x48, 0x48,
	0x5e, 0x0f, 0xa2, 0x58, 0x0f, 0xa6, 0x58, 0xcf, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x3d, 0xd5, 0xbf,
	0xa0, 0x24, 0x33, 0x3f, 0xaf, 0x58, 0xe2, 0x43, 0xcf, 0x32, 0x56, 0x05, 0x66, 0x0d, 0xce, 0x20,
	0xa8, 0x0e, 0xa7, 0xac, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x0a, 0x48,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77, 0x2e, 0x4d, 0xc9, 0x0f, 0x4b,
	0xcd, 0x2b, 0x29, 0x2d, 0x4a, 0x2d, 0xd6, 0x4f, 0x2e, 0x4d, 0xc9, 0x2f, 0xd6, 0xcd, 0xcb, 0x4f,
	0x49, 0xd5, 0x2f, 0xc9, 0xc8, 0x2c, 0x4a, 0x89, 0x2f, 0x48, 0x2c, 0x2a, 0xa9, 0x84, 0x0a, 0x43,
	0xdc, 0xae, 0x5b, 0x9c, 0x92, 0xad, 0x5f, 0x52, 0x59, 0x90, 0x0a, 0xf6, 0x5e, 0x71, 0x6a, 0x51,
	0x59, 0x66, 0x72, 0x6a, 0x12, 0x1b, 0xd8, 0x55, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17,
	0x30, 0x9f, 0xe8, 0xfa, 0x00, 0x00, 0x00,
}
