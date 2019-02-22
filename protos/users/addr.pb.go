// Code generated by protoc-gen-go. DO NOT EDIT.
// source: addr.proto

package users

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Addr struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=Mobile,proto3" json:"Mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Addr) Reset()         { *m = Addr{} }
func (m *Addr) String() string { return proto.CompactTextString(m) }
func (*Addr) ProtoMessage()    {}
func (*Addr) Descriptor() ([]byte, []int) {
	return fileDescriptor_eaa67da8ec6ef8e1, []int{0}
}

func (m *Addr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addr.Unmarshal(m, b)
}
func (m *Addr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addr.Marshal(b, m, deterministic)
}
func (m *Addr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addr.Merge(m, src)
}
func (m *Addr) XXX_Size() int {
	return xxx_messageInfo_Addr.Size(m)
}
func (m *Addr) XXX_DiscardUnknown() {
	xxx_messageInfo_Addr.DiscardUnknown(m)
}

var xxx_messageInfo_Addr proto.InternalMessageInfo

func (m *Addr) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Addr) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Addr) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func init() {
	proto.RegisterType((*Addr)(nil), "users.Addr")
}

func init() { proto.RegisterFile("addr.proto", fileDescriptor_eaa67da8ec6ef8e1) }

var fileDescriptor_eaa67da8ec6ef8e1 = []byte{
	// 105 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4c, 0x49, 0x29,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x2d, 0x4e, 0x2d, 0x2a, 0x56, 0x72, 0xe2,
	0x62, 0x71, 0x4c, 0x49, 0x29, 0x12, 0xe2, 0xe3, 0x62, 0xf2, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4,
	0x60, 0x0d, 0x62, 0xf2, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0xf1, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x52,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0xc4, 0xb8, 0xd8, 0x7c, 0xf3, 0x93, 0x32, 0x73, 0x52,
	0x25, 0x98, 0xc1, 0xa2, 0x50, 0x5e, 0x12, 0x1b, 0xd8, 0x44, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xda, 0x4f, 0x7b, 0x99, 0x5f, 0x00, 0x00, 0x00,
}