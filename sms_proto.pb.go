// Code generated by protoc-gen-go.
// source: sms.proto
// DO NOT EDIT!

/*
Package sms is a generated protocol buffer package.

It is generated from these files:
	sms.proto

It has these top-level messages:
	Sms
*/
package main

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

type Sms_MessageType int32

const (
	Sms_MESSAGE_TYPE_ALL    Sms_MessageType = 0
	Sms_MESSAGE_TYPE_INBOX  Sms_MessageType = 1
	Sms_MESSAGE_TYPE_SENT   Sms_MessageType = 2
	Sms_MESSAGE_TYPE_DRAFT  Sms_MessageType = 3
	Sms_MESSAGE_TYPE_OUTBOX Sms_MessageType = 4
	Sms_MESSAGE_TYPE_FAILED Sms_MessageType = 5
	Sms_MESSAGE_TYPE_QUEUED Sms_MessageType = 6
)

var Sms_MessageType_name = map[int32]string{
	0: "MESSAGE_TYPE_ALL",
	1: "MESSAGE_TYPE_INBOX",
	2: "MESSAGE_TYPE_SENT",
	3: "MESSAGE_TYPE_DRAFT",
	4: "MESSAGE_TYPE_OUTBOX",
	5: "MESSAGE_TYPE_FAILED",
	6: "MESSAGE_TYPE_QUEUED",
}
var Sms_MessageType_value = map[string]int32{
	"MESSAGE_TYPE_ALL":    0,
	"MESSAGE_TYPE_INBOX":  1,
	"MESSAGE_TYPE_SENT":   2,
	"MESSAGE_TYPE_DRAFT":  3,
	"MESSAGE_TYPE_OUTBOX": 4,
	"MESSAGE_TYPE_FAILED": 5,
	"MESSAGE_TYPE_QUEUED": 6,
}

func (x Sms_MessageType) String() string {
	return proto.EnumName(Sms_MessageType_name, int32(x))
}
func (Sms_MessageType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Sms struct {
	Address      string `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Body         string `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	DateReceived uint64 `protobuf:"varint,3,opt,name=date_received,json=dateReceived" json:"date_received,omitempty"`
	DateSent     uint64 `protobuf:"varint,4,opt,name=date_sent,json=dateSent" json:"date_sent,omitempty"`
	Person       uint32 `protobuf:"varint,5,opt,name=person" json:"person,omitempty"`
	Type         uint32 `protobuf:"varint,6,opt,name=type" json:"type,omitempty"`
	Protocol     uint32 `protobuf:"varint,7,opt,name=protocol" json:"protocol,omitempty"`
}

func (m *Sms) Reset()                    { *m = Sms{} }
func (m *Sms) String() string            { return proto.CompactTextString(m) }
func (*Sms) ProtoMessage()               {}
func (*Sms) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Sms)(nil), "sms.Sms")
	proto.RegisterEnum("sms.Sms_MessageType", Sms_MessageType_name, Sms_MessageType_value)
}

func init() { proto.RegisterFile("sms.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 275 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0xcf, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x06, 0x60, 0xa7, 0x49, 0xd3, 0xe6, 0x68, 0x61, 0x3c, 0x6a, 0x1d, 0x74, 0x13, 0xea, 0x26,
	0x2b, 0x37, 0x3e, 0x41, 0x24, 0x53, 0x29, 0xa4, 0xad, 0xe6, 0x02, 0xba, 0x0a, 0x69, 0x33, 0x88,
	0x60, 0x2e, 0xe4, 0x04, 0xa1, 0x0f, 0xe6, 0xa3, 0xb9, 0x97, 0x0c, 0x55, 0x08, 0xd9, 0x9d, 0xff,
	0x3b, 0x3f, 0x73, 0x01, 0x9b, 0x0a, 0xba, 0xaf, 0x9b, 0xaa, 0xad, 0xd0, 0xa0, 0x82, 0x16, 0x3f,
	0x23, 0x30, 0xa2, 0x82, 0x50, 0xc0, 0x24, 0xcb, 0xf3, 0x46, 0x11, 0x09, 0xe6, 0x30, 0xd7, 0x0e,
	0xff, 0x22, 0x22, 0x98, 0xbb, 0x2a, 0x3f, 0x88, 0x91, 0x66, 0x3d, 0xe3, 0x1d, 0xcc, 0xf2, 0xac,
	0x55, 0x69, 0xa3, 0xf6, 0xea, 0xe3, 0x4b, 0xe5, 0xc2, 0x70, 0x98, 0x6b, 0x86, 0x67, 0x1d, 0x86,
	0x47, 0xc3, 0x5b, 0xb0, 0x75, 0x89, 0x54, 0xd9, 0x0a, 0x53, 0x17, 0xa6, 0x1d, 0x44, 0xaa, 0x6c,
	0x71, 0x0e, 0x56, 0xad, 0x1a, 0xaa, 0x4a, 0x31, 0x76, 0x98, 0x3b, 0x0b, 0x8f, 0xa9, 0xbb, 0xad,
	0x3d, 0xd4, 0x4a, 0x58, 0x5a, 0xf5, 0x8c, 0x37, 0x30, 0xd5, 0x2f, 0xde, 0x57, 0x9f, 0x62, 0xa2,
	0xfd, 0x3f, 0x2f, 0xbe, 0x19, 0x9c, 0xae, 0x15, 0x51, 0xf6, 0xae, 0xe2, 0xae, 0x7b, 0x09, 0x7c,
	0x2d, 0xa3, 0xc8, 0x7b, 0x92, 0x69, 0xfc, 0xf6, 0x2c, 0x53, 0x2f, 0x08, 0xf8, 0x09, 0xce, 0x01,
	0x7b, 0xba, 0xda, 0x3c, 0x6e, 0x5f, 0x39, 0xc3, 0x2b, 0x38, 0xef, 0x79, 0x24, 0x37, 0x31, 0x1f,
	0x0d, 0xea, 0x7e, 0xe8, 0x2d, 0x63, 0x6e, 0xe0, 0x35, 0x5c, 0xf4, 0x7c, 0x9b, 0xc4, 0xdd, 0x39,
	0xe6, 0x60, 0xb1, 0xf4, 0x56, 0x81, 0xf4, 0xf9, 0x78, 0xb0, 0x78, 0x49, 0x64, 0x22, 0x7d, 0x6e,
	0xed, 0x2c, 0xfd, 0x83, 0x87, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xff, 0x3f, 0x16, 0x82, 0x90,
	0x01, 0x00, 0x00,
}
