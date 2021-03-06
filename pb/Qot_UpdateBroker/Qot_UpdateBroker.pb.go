// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_UpdateBroker.proto

package Qot_UpdateBroker

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

type S2C struct {
	Security             *Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	BrokerAskList        []*Broker `protobuf:"bytes,2,rep,name=brokerAskList" json:"brokerAskList,omitempty"`
	BrokerBidList        []*Broker `protobuf:"bytes,3,rep,name=brokerBidList" json:"brokerBidList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b58d393d5c0052b, []int{0}
}

func (m *S2C) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S2C.Unmarshal(m, b)
}
func (m *S2C) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S2C.Marshal(b, m, deterministic)
}
func (m *S2C) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S2C.Merge(m, src)
}
func (m *S2C) XXX_Size() int {
	return xxx_messageInfo_S2C.Size(m)
}
func (m *S2C) XXX_DiscardUnknown() {
	xxx_messageInfo_S2C.DiscardUnknown(m)
}

var xxx_messageInfo_S2C proto.InternalMessageInfo

func (m *S2C) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *S2C) GetBrokerAskList() []*Broker {
	if m != nil {
		return m.BrokerAskList
	}
	return nil
}

func (m *S2C) GetBrokerBidList() []*Broker {
	if m != nil {
		return m.BrokerBidList
	}
	return nil
}

type Response struct {
	RetType              *int32   `protobuf:"varint,1,req,name=retType,def=-400" json:"retType,omitempty"`
	RetMsg               *string  `protobuf:"bytes,2,opt,name=retMsg" json:"retMsg,omitempty"`
	ErrCode              *int32   `protobuf:"varint,3,opt,name=errCode" json:"errCode,omitempty"`
	S2C                  *S2C     `protobuf:"bytes,4,opt,name=s2c" json:"s2c,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1b58d393d5c0052b, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

const Default_Response_RetType int32 = -400

func (m *Response) GetRetType() int32 {
	if m != nil && m.RetType != nil {
		return *m.RetType
	}
	return Default_Response_RetType
}

func (m *Response) GetRetMsg() string {
	if m != nil && m.RetMsg != nil {
		return *m.RetMsg
	}
	return ""
}

func (m *Response) GetErrCode() int32 {
	if m != nil && m.ErrCode != nil {
		return *m.ErrCode
	}
	return 0
}

func (m *Response) GetS2C() *S2C {
	if m != nil {
		return m.S2C
	}
	return nil
}

func init() {
	proto.RegisterType((*S2C)(nil), "Qot_UpdateBroker.S2C")
	proto.RegisterType((*Response)(nil), "Qot_UpdateBroker.Response")
}

func init() { proto.RegisterFile("Qot_UpdateBroker.proto", fileDescriptor_1b58d393d5c0052b) }

var fileDescriptor_1b58d393d5c0052b = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x0b, 0xcc, 0x2f, 0x89,
	0x0f, 0x2d, 0x48, 0x49, 0x2c, 0x49, 0x75, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x40, 0x17, 0x97, 0xe2, 0x71, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0xc8,
	0x4b, 0x81, 0xe5, 0x91, 0x45, 0x94, 0x56, 0x33, 0x72, 0x31, 0x07, 0x1b, 0x39, 0x0b, 0x19, 0x70,
	0x71, 0x14, 0xa7, 0x26, 0x97, 0x16, 0x65, 0x96, 0x54, 0x4a, 0x30, 0x2a, 0x30, 0x69, 0x70, 0x1b,
	0x89, 0xe8, 0x21, 0x29, 0x0e, 0x86, 0xca, 0x05, 0xc1, 0x55, 0x09, 0x59, 0x70, 0xf1, 0x26, 0x81,
	0xed, 0x70, 0x2c, 0xce, 0xf6, 0xc9, 0x2c, 0x2e, 0x91, 0x60, 0x52, 0x60, 0xd6, 0xe0, 0x36, 0x12,
	0x42, 0xd6, 0x06, 0x71, 0x44, 0x10, 0xaa, 0x42, 0x84, 0x4e, 0xa7, 0xcc, 0x14, 0xb0, 0x4e, 0x66,
	0x42, 0x3a, 0xa1, 0x0a, 0x95, 0x5a, 0x19, 0xb9, 0x38, 0x82, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x85, 0xe4, 0xb8, 0xd8, 0x8b, 0x52, 0x4b, 0x42, 0x2a, 0x0b, 0x52, 0xc1, 0x2e, 0x66, 0xb5,
	0x62, 0xd1, 0x35, 0x31, 0x30, 0x08, 0x82, 0x09, 0x0a, 0x89, 0x71, 0xb1, 0x15, 0xa5, 0x96, 0xf8,
	0x16, 0xa7, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x12, 0x5c, 0xec, 0xa9,
	0x45, 0x45, 0xce, 0xf9, 0x29, 0xa9, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x30, 0xae, 0x90,
	0x3a, 0x17, 0x73, 0xb1, 0x51, 0xb2, 0x04, 0x8b, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xa8, 0x1e, 0x46,
	0x20, 0x07, 0x1b, 0x39, 0x07, 0x81, 0x54, 0x00, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x59, 0x7a,
	0xdb, 0x80, 0x01, 0x00, 0x00,
}
