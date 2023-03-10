// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/feature.proto

package pb

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

type UserProfile struct {
	Age                  int32    `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Gender               string   `protobuf:"bytes,2,opt,name=gender,proto3" json:"gender,omitempty"`
	Province             string   `protobuf:"bytes,3,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Network              string   `protobuf:"bytes,5,opt,name=network,proto3" json:"network,omitempty"`
	LastLoginIp          []string `protobuf:"bytes,6,rep,name=last_login_ip,json=lastLoginIp,proto3" json:"last_login_ip,omitempty"`
	TagWeigt             []*Tag   `protobuf:"bytes,7,rep,name=tag_weigt,json=tagWeigt,proto3" json:"tag_weigt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserProfile) Reset()         { *m = UserProfile{} }
func (m *UserProfile) String() string { return proto.CompactTextString(m) }
func (*UserProfile) ProtoMessage()    {}
func (*UserProfile) Descriptor() ([]byte, []int) {
	return fileDescriptor_769667aecd789a5e, []int{0}
}

func (m *UserProfile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserProfile.Unmarshal(m, b)
}
func (m *UserProfile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserProfile.Marshal(b, m, deterministic)
}
func (m *UserProfile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserProfile.Merge(m, src)
}
func (m *UserProfile) XXX_Size() int {
	return xxx_messageInfo_UserProfile.Size(m)
}
func (m *UserProfile) XXX_DiscardUnknown() {
	xxx_messageInfo_UserProfile.DiscardUnknown(m)
}

var xxx_messageInfo_UserProfile proto.InternalMessageInfo

func (m *UserProfile) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *UserProfile) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserProfile) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *UserProfile) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UserProfile) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *UserProfile) GetLastLoginIp() []string {
	if m != nil {
		return m.LastLoginIp
	}
	return nil
}

func (m *UserProfile) GetTagWeigt() []*Tag {
	if m != nil {
		return m.TagWeigt
	}
	return nil
}

type Tag struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Weight               float64  `protobuf:"fixed64,2,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_769667aecd789a5e, []int{1}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetWeight() float64 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func init() {
	proto.RegisterType((*UserProfile)(nil), "UserProfile")
	proto.RegisterType((*Tag)(nil), "Tag")
}

func init() { proto.RegisterFile("pb/feature.proto", fileDescriptor_769667aecd789a5e) }

var fileDescriptor_769667aecd789a5e = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xa9, 0xed, 0x76, 0xb7, 0x29, 0xc2, 0x92, 0x83, 0x04, 0x4f, 0xb5, 0x22, 0xf4, 0xd4,
	0xa2, 0xbe, 0x81, 0xe0, 0x41, 0xf0, 0x20, 0x61, 0x45, 0xf4, 0x52, 0xa6, 0x75, 0x36, 0x0d, 0xdb,
	0x26, 0x21, 0xcd, 0xba, 0xf8, 0x86, 0x3e, 0x96, 0x24, 0xae, 0x7b, 0xfb, 0xbf, 0x6f, 0x18, 0x86,
	0x7f, 0xc8, 0xda, 0x74, 0xcd, 0x16, 0xc1, 0xed, 0x2d, 0xd6, 0xc6, 0x6a, 0xa7, 0xcb, 0x9f, 0x88,
	0xe4, 0xaf, 0x33, 0xda, 0x17, 0xab, 0xb7, 0x72, 0x44, 0xba, 0x26, 0x31, 0x08, 0x64, 0x51, 0x11,
	0x55, 0x0b, 0xee, 0x23, 0xbd, 0x20, 0xa9, 0x40, 0xf5, 0x89, 0x96, 0x9d, 0x15, 0x51, 0x95, 0xf1,
	0x23, 0xd1, 0x4b, 0xb2, 0x32, 0x56, 0x7f, 0x49, 0xd5, 0x23, 0x8b, 0xc3, 0xe4, 0xc4, 0x94, 0x92,
	0xa4, 0x97, 0xee, 0x9b, 0x25, 0xc1, 0x87, 0x4c, 0x19, 0x59, 0x2a, 0x74, 0x07, 0x6d, 0x77, 0x6c,
	0x11, 0xf4, 0x3f, 0xd2, 0x92, 0x9c, 0x8f, 0x30, 0xbb, 0x76, 0xd4, 0x42, 0xaa, 0x56, 0x1a, 0x96,
	0x16, 0x71, 0x95, 0xf1, 0xdc, 0xcb, 0x67, 0xef, 0x9e, 0x0c, 0xbd, 0x22, 0x99, 0x03, 0xd1, 0x1e,
	0x50, 0x0a, 0xc7, 0x96, 0x45, 0x5c, 0xe5, 0x77, 0x49, 0xbd, 0x01, 0xc1, 0x57, 0x0e, 0xc4, 0x9b,
	0xb7, 0xe5, 0x2d, 0x89, 0x37, 0x20, 0xfc, 0x6d, 0x05, 0xd3, 0x5f, 0x85, 0x8c, 0x87, 0xec, 0x3b,
	0xf8, 0xcd, 0xc1, 0x85, 0x0e, 0x11, 0x3f, 0xd2, 0xc3, 0xcd, 0xc7, 0xb5, 0x90, 0x6e, 0xd8, 0x77,
	0x75, 0xaf, 0xa7, 0xe6, 0x11, 0x66, 0xad, 0xde, 0x41, 0x35, 0x42, 0xb7, 0x1d, 0xaa, 0x7e, 0x98,
	0xc0, 0xee, 0x1a, 0xd3, 0x75, 0x69, 0xf8, 0xd5, 0xfd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46,
	0x8b, 0x65, 0xeb, 0x3f, 0x01, 0x00, 0x00,
}
