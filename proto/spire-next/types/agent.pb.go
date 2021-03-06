// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

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

type Agent struct {
	// Output only. SPIFFE ID of the agent.
	Id *SPIFFEID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Output only. The method by which the agent attested.
	AttestationType string `protobuf:"bytes,2,opt,name=attestation_type,json=attestationType,proto3" json:"attestation_type,omitempty"`
	// Output only. The X509-SVID serial number.
	X509SvidSerialNumber string `protobuf:"bytes,3,opt,name=x509svid_serial_number,json=x509svidSerialNumber,proto3" json:"x509svid_serial_number,omitempty"`
	// Output only. The X509-SVID expiration (seconds since Unix epoch).
	X509SvidExpiresAt int64 `protobuf:"varint,4,opt,name=x509svid_expires_at,json=x509svidExpiresAt,proto3" json:"x509svid_expires_at,omitempty"`
	// Output only. The selectors attributed to the agent during attestation.
	Selectors []*Selector `protobuf:"bytes,5,rep,name=selectors,proto3" json:"selectors,omitempty"`
	// Output only. Whether or not the agent is banned.
	Banned               bool     `protobuf:"varint,6,opt,name=banned,proto3" json:"banned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Agent) Reset()         { *m = Agent{} }
func (m *Agent) String() string { return proto.CompactTextString(m) }
func (*Agent) ProtoMessage()    {}
func (*Agent) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}

func (m *Agent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Agent.Unmarshal(m, b)
}
func (m *Agent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Agent.Marshal(b, m, deterministic)
}
func (m *Agent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Agent.Merge(m, src)
}
func (m *Agent) XXX_Size() int {
	return xxx_messageInfo_Agent.Size(m)
}
func (m *Agent) XXX_DiscardUnknown() {
	xxx_messageInfo_Agent.DiscardUnknown(m)
}

var xxx_messageInfo_Agent proto.InternalMessageInfo

func (m *Agent) GetId() *SPIFFEID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Agent) GetAttestationType() string {
	if m != nil {
		return m.AttestationType
	}
	return ""
}

func (m *Agent) GetX509SvidSerialNumber() string {
	if m != nil {
		return m.X509SvidSerialNumber
	}
	return ""
}

func (m *Agent) GetX509SvidExpiresAt() int64 {
	if m != nil {
		return m.X509SvidExpiresAt
	}
	return 0
}

func (m *Agent) GetSelectors() []*Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func (m *Agent) GetBanned() bool {
	if m != nil {
		return m.Banned
	}
	return false
}

type AgentMask struct {
	Id                   bool     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AttestationType      bool     `protobuf:"varint,2,opt,name=attestation_type,json=attestationType,proto3" json:"attestation_type,omitempty"`
	X509SvidSerialNumber bool     `protobuf:"varint,3,opt,name=x509svid_serial_number,json=x509svidSerialNumber,proto3" json:"x509svid_serial_number,omitempty"`
	X509SvidExpiresAt    bool     `protobuf:"varint,4,opt,name=x509svid_expires_at,json=x509svidExpiresAt,proto3" json:"x509svid_expires_at,omitempty"`
	Selectors            bool     `protobuf:"varint,5,opt,name=selectors,proto3" json:"selectors,omitempty"`
	Banned               bool     `protobuf:"varint,6,opt,name=banned,proto3" json:"banned,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AgentMask) Reset()         { *m = AgentMask{} }
func (m *AgentMask) String() string { return proto.CompactTextString(m) }
func (*AgentMask) ProtoMessage()    {}
func (*AgentMask) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{1}
}

func (m *AgentMask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AgentMask.Unmarshal(m, b)
}
func (m *AgentMask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AgentMask.Marshal(b, m, deterministic)
}
func (m *AgentMask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AgentMask.Merge(m, src)
}
func (m *AgentMask) XXX_Size() int {
	return xxx_messageInfo_AgentMask.Size(m)
}
func (m *AgentMask) XXX_DiscardUnknown() {
	xxx_messageInfo_AgentMask.DiscardUnknown(m)
}

var xxx_messageInfo_AgentMask proto.InternalMessageInfo

func (m *AgentMask) GetId() bool {
	if m != nil {
		return m.Id
	}
	return false
}

func (m *AgentMask) GetAttestationType() bool {
	if m != nil {
		return m.AttestationType
	}
	return false
}

func (m *AgentMask) GetX509SvidSerialNumber() bool {
	if m != nil {
		return m.X509SvidSerialNumber
	}
	return false
}

func (m *AgentMask) GetX509SvidExpiresAt() bool {
	if m != nil {
		return m.X509SvidExpiresAt
	}
	return false
}

func (m *AgentMask) GetSelectors() bool {
	if m != nil {
		return m.Selectors
	}
	return false
}

func (m *AgentMask) GetBanned() bool {
	if m != nil {
		return m.Banned
	}
	return false
}

func init() {
	proto.RegisterType((*Agent)(nil), "spire.types.Agent")
	proto.RegisterType((*AgentMask)(nil), "spire.types.AgentMask")
}

func init() { proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77) }

var fileDescriptor_56ede974c0020f77 = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4e, 0xfa, 0x40,
	0x10, 0xc6, 0xd3, 0xf2, 0x87, 0x94, 0x25, 0xf9, 0xab, 0xab, 0x92, 0x8d, 0x31, 0xb1, 0x21, 0x31,
	0xa9, 0x07, 0xb7, 0x44, 0xf4, 0xe0, 0x11, 0x23, 0x24, 0x1c, 0x34, 0xa6, 0x78, 0xf2, 0xd2, 0x6c,
	0xe9, 0x80, 0x1b, 0x61, 0xdb, 0x74, 0x07, 0x03, 0xef, 0xe0, 0x53, 0xfa, 0x24, 0xa6, 0x0b, 0x45,
	0x54, 0xa2, 0xe1, 0x38, 0xf3, 0x7d, 0xb3, 0x3b, 0xbf, 0xc9, 0x47, 0x6a, 0x62, 0x04, 0x0a, 0x79,
	0x9a, 0x25, 0x98, 0xd0, 0x9a, 0x4e, 0x65, 0x06, 0x1c, 0xe7, 0x29, 0xe8, 0xa3, 0x13, 0x53, 0x9c,
	0x2b, 0x98, 0xa1, 0x6f, 0x3a, 0xbe, 0x86, 0x31, 0x0c, 0x30, 0xc9, 0x16, 0xee, 0x4d, 0x86, 0x54,
	0x0e, 0x87, 0x20, 0xe3, 0x85, 0xa1, 0xf1, 0x66, 0x93, 0x72, 0x3b, 0x7f, 0x9e, 0x9e, 0x12, 0x5b,
	0xc6, 0xcc, 0x72, 0x2d, 0xaf, 0x76, 0x71, 0xc8, 0xd7, 0x7e, 0xe1, 0xfd, 0x87, 0x5e, 0xb7, 0xdb,
	0xe9, 0xdd, 0x06, 0xb6, 0x8c, 0xe9, 0x19, 0xd9, 0x15, 0x88, 0xa0, 0x51, 0xa0, 0x4c, 0x54, 0x98,
	0x3b, 0x98, 0xed, 0x5a, 0x5e, 0x35, 0xd8, 0x59, 0xeb, 0x3f, 0xce, 0x53, 0xa0, 0x97, 0xa4, 0x3e,
	0xbb, 0x6a, 0x5e, 0xeb, 0x57, 0x19, 0x87, 0x1a, 0x32, 0x29, 0xc6, 0xa1, 0x9a, 0x4e, 0x22, 0xc8,
	0x58, 0xc9, 0x0c, 0x1c, 0x14, 0x6a, 0xdf, 0x88, 0xf7, 0x46, 0xa3, 0x9c, 0xec, 0xaf, 0xa6, 0x60,
	0x96, 0xaf, 0xa1, 0x43, 0x81, 0xec, 0x9f, 0x6b, 0x79, 0xa5, 0x60, 0xaf, 0x90, 0x3a, 0x0b, 0xa5,
	0x8d, 0xb4, 0x45, 0xaa, 0x05, 0xb4, 0x66, 0x65, 0xb7, 0xf4, 0x73, 0xfd, 0xa5, 0x1a, 0x7c, 0xfa,
	0x68, 0x9d, 0x54, 0x22, 0xa1, 0x14, 0xc4, 0xac, 0xe2, 0x5a, 0x9e, 0x13, 0x2c, 0xab, 0xc6, 0xbb,
	0x45, 0xaa, 0xe6, 0x1c, 0x77, 0x42, 0xbf, 0xd0, 0xff, 0xab, 0x93, 0x38, 0xbf, 0xb2, 0x3b, 0xdb,
	0xb2, 0x3b, 0xdb, 0xb3, 0x3b, 0x9b, 0xd8, 0x8f, 0xbf, 0xb2, 0xe7, 0xae, 0xbf, 0x21, 0x6f, 0x9a,
	0x4f, 0x7c, 0x24, 0xf1, 0x79, 0x1a, 0xf1, 0x41, 0x32, 0x59, 0x06, 0xc2, 0x37, 0x17, 0xf3, 0x4d,
	0x28, 0xfc, 0xef, 0xa1, 0x89, 0x2a, 0xa6, 0xdf, 0xfa, 0x08, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x3b,
	0x43, 0x43, 0x8a, 0x02, 0x00, 0x00,
}
