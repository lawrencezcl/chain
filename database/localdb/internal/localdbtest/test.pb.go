// Code generated by protoc-gen-go.
// source: test.proto
// DO NOT EDIT!

/*
Package localdbtest is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	TestItem
*/
package localdbtest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestItem struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *TestItem) Reset()                    { *m = TestItem{} }
func (m *TestItem) String() string            { return proto.CompactTextString(m) }
func (*TestItem) ProtoMessage()               {}
func (*TestItem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestItem) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*TestItem)(nil), "localdbtest.TestItem")
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 82 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xc9, 0x4f, 0x4e, 0xcc, 0x49, 0x49, 0x02,
	0x09, 0x29, 0x29, 0x70, 0x71, 0x84, 0xa4, 0x16, 0x97, 0x78, 0x96, 0xa4, 0xe6, 0x0a, 0x89, 0x70,
	0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x38, 0x49,
	0x6c, 0x60, 0x5d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xd5, 0xdb, 0xd6, 0x43, 0x00,
	0x00, 0x00,
}
