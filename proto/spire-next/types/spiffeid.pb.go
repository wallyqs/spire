// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spiffeid.proto

package types

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

// A SPIFFE ID, consisting of the trust domain name and a path portions of
// the SPIFFE ID URI.
type SPIFFEID struct {
	// Trust domain portion the SPIFFE ID (e.g. "example.org")
	TrustDomain string `protobuf:"bytes,1,opt,name=trust_domain,json=trustDomain,proto3" json:"trust_domain,omitempty"`
	// The path component of the SPIFFE ID (e.g. "/foo/bar/baz"). The path
	// SHOULD have a leading slash. Consumers MUST normalize the path before
	// making any sort of comparison between IDs.
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SPIFFEID) Reset()         { *m = SPIFFEID{} }
func (m *SPIFFEID) String() string { return proto.CompactTextString(m) }
func (*SPIFFEID) ProtoMessage()    {}
func (*SPIFFEID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6e9c857bfc8eae8, []int{0}
}

func (m *SPIFFEID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SPIFFEID.Unmarshal(m, b)
}
func (m *SPIFFEID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SPIFFEID.Marshal(b, m, deterministic)
}
func (m *SPIFFEID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SPIFFEID.Merge(m, src)
}
func (m *SPIFFEID) XXX_Size() int {
	return xxx_messageInfo_SPIFFEID.Size(m)
}
func (m *SPIFFEID) XXX_DiscardUnknown() {
	xxx_messageInfo_SPIFFEID.DiscardUnknown(m)
}

var xxx_messageInfo_SPIFFEID proto.InternalMessageInfo

func (m *SPIFFEID) GetTrustDomain() string {
	if m != nil {
		return m.TrustDomain
	}
	return ""
}

func (m *SPIFFEID) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*SPIFFEID)(nil), "spire.types.SPIFFEID")
}

func init() { proto.RegisterFile("spiffeid.proto", fileDescriptor_f6e9c857bfc8eae8) }

var fileDescriptor_f6e9c857bfc8eae8 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2e, 0xc8, 0x4c,
	0x4b, 0x4b, 0xcd, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x2e, 0xc8, 0x2c,
	0x4a, 0xd5, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x72, 0xe4, 0xe2, 0x08, 0x0e, 0xf0, 0x74, 0x73,
	0x73, 0xf5, 0x74, 0x11, 0x52, 0xe4, 0xe2, 0x29, 0x29, 0x2a, 0x2d, 0x2e, 0x89, 0x4f, 0xc9, 0xcf,
	0x4d, 0xcc, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x06, 0x8b, 0xb9, 0x80, 0x85,
	0x84, 0x84, 0xb8, 0x58, 0x0a, 0x12, 0x4b, 0x32, 0x24, 0x98, 0xc0, 0x52, 0x60, 0xb6, 0x93, 0x41,
	0x94, 0x5e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc4, 0x32, 0x7d,
	0xb0, 0x1d, 0xfa, 0x60, 0x0b, 0x21, 0x6c, 0xdd, 0xbc, 0xd4, 0x8a, 0x12, 0x7d, 0xb0, 0xa5, 0x49,
	0x6c, 0x60, 0x71, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x9b, 0x8b, 0x26, 0x9a, 0x00,
	0x00, 0x00,
}