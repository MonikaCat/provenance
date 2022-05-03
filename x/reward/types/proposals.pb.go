// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/reward/v1/proposals.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/provenance-io/provenance/x/epoch/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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
	proto.RegisterFile("provenance/reward/v1/proposals.proto", fileDescriptor_081a3605de2e1dba)
}

var fileDescriptor_081a3605de2e1dba = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0x3f, 0x4e, 0xc3, 0x30,
	0x14, 0xc6, 0x93, 0x85, 0xa1, 0x23, 0xaa, 0x84, 0xd4, 0xc1, 0x08, 0x89, 0x95, 0x3c, 0x05, 0x6e,
	0xc0, 0x09, 0x98, 0xd9, 0x6c, 0xf7, 0xe1, 0x5a, 0x6a, 0xf2, 0x59, 0x7e, 0x4e, 0x80, 0x5b, 0x70,
	0x2c, 0xc6, 0x8e, 0x8c, 0x28, 0xb9, 0x08, 0xaa, 0xdb, 0x28, 0x11, 0xb0, 0xbd, 0xa7, 0xef, 0xa7,
	0xef, 0xcf, 0xea, 0x36, 0x44, 0xf4, 0xdc, 0xea, 0xd6, 0x32, 0x45, 0x7e, 0xd5, 0x71, 0x4b, 0x7d,
	0x4d, 0x21, 0x22, 0x40, 0xf4, 0x5e, 0xaa, 0x10, 0x91, 0x70, 0xb9, 0x9e, 0xa9, 0xea, 0x44, 0x55,
	0x7d, 0xbd, 0x59, 0x3b, 0x38, 0x64, 0x80, 0x8e, 0xd7, 0x89, 0xdd, 0x28, 0x07, 0xb8, 0x3d, 0x53,
	0xfe, 0x4c, 0xf7, 0x42, 0xdb, 0x2e, 0xea, 0xe4, 0xd1, 0x9e, 0xf5, 0xeb, 0xdf, 0x7a, 0xf2, 0x0d,
	0x4b, 0xd2, 0x4d, 0x98, 0x0c, 0x2c, 0xa4, 0x81, 0x90, 0xd1, 0xc2, 0xd4, 0xd7, 0x86, 0x93, 0xae,
	0xc9, 0xc2, 0x4f, 0x06, 0x37, 0x8b, 0xca, 0x1c, 0x60, 0x77, 0xc7, 0xc6, 0x8e, 0x5b, 0x16, 0x2f,
	0xff, 0x20, 0xf3, 0xaa, 0x73, 0xf3, 0x8c, 0x3c, 0xba, 0xcf, 0x41, 0x95, 0x87, 0x41, 0x95, 0xdf,
	0x83, 0x2a, 0x3f, 0x46, 0x55, 0x1c, 0x46, 0x55, 0x7c, 0x8d, 0xaa, 0x58, 0x5d, 0xf9, 0x3c, 0xe5,
	0xcf, 0xde, 0xa7, 0xf2, 0xf9, 0xde, 0xf9, 0xb4, 0xeb, 0x4c, 0x65, 0xd1, 0xd0, 0x8c, 0xdc, 0x79,
	0x2c, 0x3e, 0x7a, 0x9b, 0x22, 0xd3, 0x7b, 0x60, 0x31, 0x17, 0x39, 0xef, 0xe1, 0x27, 0x00, 0x00,
	0xff, 0xff, 0xb8, 0x5c, 0xe1, 0x28, 0x6a, 0x01, 0x00, 0x00,
}
