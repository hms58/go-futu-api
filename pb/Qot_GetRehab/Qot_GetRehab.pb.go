// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetRehab.proto

package Qot_GetRehab

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

type C2S struct {
	SecurityList         []*Security `protobuf:"bytes,1,rep,name=securityList" json:"securityList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b325efa9ba713c, []int{0}
}

func (m *C2S) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2S.Unmarshal(m, b)
}
func (m *C2S) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2S.Marshal(b, m, deterministic)
}
func (m *C2S) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2S.Merge(m, src)
}
func (m *C2S) XXX_Size() int {
	return xxx_messageInfo_C2S.Size(m)
}
func (m *C2S) XXX_DiscardUnknown() {
	xxx_messageInfo_C2S.DiscardUnknown(m)
}

var xxx_messageInfo_C2S proto.InternalMessageInfo

func (m *C2S) GetSecurityList() []*Security {
	if m != nil {
		return m.SecurityList
	}
	return nil
}

type SecurityRehab struct {
	Security             *Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	RehabList            []*Rehab  `protobuf:"bytes,2,rep,name=rehabList" json:"rehabList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SecurityRehab) Reset()         { *m = SecurityRehab{} }
func (m *SecurityRehab) String() string { return proto.CompactTextString(m) }
func (*SecurityRehab) ProtoMessage()    {}
func (*SecurityRehab) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b325efa9ba713c, []int{1}
}

func (m *SecurityRehab) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityRehab.Unmarshal(m, b)
}
func (m *SecurityRehab) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityRehab.Marshal(b, m, deterministic)
}
func (m *SecurityRehab) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityRehab.Merge(m, src)
}
func (m *SecurityRehab) XXX_Size() int {
	return xxx_messageInfo_SecurityRehab.Size(m)
}
func (m *SecurityRehab) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityRehab.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityRehab proto.InternalMessageInfo

func (m *SecurityRehab) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

func (m *SecurityRehab) GetRehabList() []*Rehab {
	if m != nil {
		return m.RehabList
	}
	return nil
}

type S2C struct {
	SecurityRehabList    []*SecurityRehab `protobuf:"bytes,1,rep,name=securityRehabList" json:"securityRehabList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b325efa9ba713c, []int{2}
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

func (m *S2C) GetSecurityRehabList() []*SecurityRehab {
	if m != nil {
		return m.SecurityRehabList
	}
	return nil
}

type Request struct {
	C2S                  *C2S     `protobuf:"bytes,1,req,name=c2s" json:"c2s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b325efa9ba713c, []int{3}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetC2S() *C2S {
	if m != nil {
		return m.C2S
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
	return fileDescriptor_f4b325efa9ba713c, []int{4}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetRehab.C2S")
	proto.RegisterType((*SecurityRehab)(nil), "Qot_GetRehab.SecurityRehab")
	proto.RegisterType((*S2C)(nil), "Qot_GetRehab.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetRehab.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetRehab.Response")
}

func init() { proto.RegisterFile("Qot_GetRehab.proto", fileDescriptor_f4b325efa9ba713c) }

var fileDescriptor_f4b325efa9ba713c = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xc9, 0xba, 0xb9, 0xed, 0x39, 0xc1, 0x06, 0x91, 0x30, 0x41, 0x4a, 0xbd, 0xf4, 0x62,
	0x2d, 0xc1, 0x83, 0x78, 0xf1, 0x90, 0x83, 0x08, 0x0a, 0xfa, 0xea, 0x5d, 0xb4, 0x3e, 0x74, 0x87,
	0x2d, 0x35, 0xc9, 0x0e, 0x3b, 0xfa, 0x9f, 0x4b, 0x6a, 0x3b, 0x53, 0x86, 0xc7, 0xbc, 0xcf, 0xf7,
	0x17, 0x01, 0xfe, 0xa4, 0xdd, 0xcb, 0x2d, 0x39, 0xa4, 0xcf, 0xd7, 0xb7, 0xbc, 0x36, 0xda, 0x69,
	0x3e, 0x0b, 0x6f, 0xf3, 0x99, 0xd2, 0xcb, 0xa5, 0x5e, 0xfd, 0xb2, 0xf9, 0xa1, 0x67, 0xe1, 0x25,
	0xbd, 0x81, 0x48, 0xc9, 0x92, 0x5f, 0xc1, 0xcc, 0x52, 0xb5, 0x36, 0x0b, 0xb7, 0xb9, 0x5f, 0x58,
	0x27, 0x58, 0x12, 0x65, 0xfb, 0xf2, 0x28, 0x0f, 0xf4, 0x65, 0xcb, 0xb1, 0xa7, 0x4c, 0x0d, 0x1c,
	0x6c, 0x89, 0x6f, 0xe4, 0x05, 0x4c, 0x3a, 0x81, 0x60, 0xc9, 0xe0, 0xdf, 0x98, 0xad, 0x8a, 0x5f,
	0xc0, 0xd4, 0x78, 0x6b, 0xd3, 0x3c, 0x68, 0x9a, 0xe3, 0xd0, 0xd2, 0xe4, 0xe2, 0x9f, 0x26, 0x7d,
	0x84, 0xa8, 0x94, 0x8a, 0xdf, 0x41, 0x6c, 0xc3, 0xea, 0x60, 0xf9, 0x49, 0xde, 0xfb, 0x99, 0xde,
	0x42, 0xdc, 0x75, 0xa5, 0x39, 0x8c, 0x91, 0xbe, 0xd6, 0x64, 0x1d, 0x3f, 0x83, 0xa8, 0x92, 0xb6,
	0x9d, 0x1e, 0xf7, 0x73, 0x94, 0x2c, 0xd1, 0xd3, 0xf4, 0x9b, 0xc1, 0x04, 0xc9, 0xd6, 0x7a, 0x65,
	0x89, 0x9f, 0xc2, 0xd8, 0x90, 0x7b, 0xde, 0xd4, 0xd4, 0xb8, 0x46, 0xd7, 0xc3, 0xf3, 0xcb, 0xa2,
	0xc0, 0xee, 0xc8, 0x8f, 0x61, 0xcf, 0x90, 0x7b, 0xb0, 0x1f, 0x62, 0x90, 0xb0, 0x6c, 0x8a, 0xed,
	0x8b, 0x0b, 0x18, 0x93, 0x31, 0x4a, 0xbf, 0x93, 0x88, 0x12, 0x96, 0x8d, 0xb0, 0x7b, 0xfa, 0x0d,
	0x56, 0x56, 0x62, 0x98, 0xb0, 0xdd, 0x0d, 0xa5, 0x54, 0xe8, 0xe9, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x3a, 0x3c, 0x43, 0x33, 0xfd, 0x01, 0x00, 0x00,
}
