// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetOrderDetail.proto

package Qot_GetOrderDetail

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
	Security             *Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_9136f364d33d18c8, []int{0}
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

func (m *C2S) GetSecurity() *Security {
	if m != nil {
		return m.Security
	}
	return nil
}

type S2C struct {
	Security             *Security    `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	OrderDetailAsk       *OrderDetail `protobuf:"bytes,2,req,name=orderDetailAsk" json:"orderDetailAsk,omitempty"`
	OrderDetailBid       *OrderDetail `protobuf:"bytes,3,req,name=orderDetailBid" json:"orderDetailBid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_9136f364d33d18c8, []int{1}
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

func (m *S2C) GetOrderDetailAsk() *OrderDetail {
	if m != nil {
		return m.OrderDetailAsk
	}
	return nil
}

func (m *S2C) GetOrderDetailBid() *OrderDetail {
	if m != nil {
		return m.OrderDetailBid
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
	return fileDescriptor_9136f364d33d18c8, []int{2}
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
	return fileDescriptor_9136f364d33d18c8, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetOrderDetail.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetOrderDetail.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetOrderDetail.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetOrderDetail.Response")
}

func init() { proto.RegisterFile("Qot_GetOrderDetail.proto", fileDescriptor_9136f364d33d18c8) }

var fileDescriptor_9136f364d33d18c8 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xf3, 0x40,
	0x10, 0x86, 0xd9, 0x6c, 0xfb, 0xa5, 0xdf, 0x28, 0x22, 0x8b, 0xe8, 0xd2, 0x83, 0x84, 0x9c, 0xe2,
	0xc1, 0x12, 0x96, 0x82, 0xe0, 0x45, 0xec, 0x0a, 0x9e, 0x44, 0x9c, 0xf5, 0x2e, 0x92, 0x0c, 0x12,
	0xb4, 0xdd, 0xb8, 0xbb, 0x3d, 0xf4, 0x17, 0xf8, 0x9b, 0xfc, 0x77, 0x92, 0x36, 0xd5, 0x10, 0x11,
	0xf4, 0x38, 0x33, 0xcf, 0xfb, 0xf2, 0xec, 0x82, 0xbc, 0xb3, 0xe1, 0xe1, 0x9a, 0xc2, 0xad, 0x2b,
	0xc9, 0x5d, 0x51, 0x78, 0xac, 0x5e, 0x26, 0xb5, 0xb3, 0xc1, 0x0a, 0xf1, 0xfd, 0x32, 0xde, 0xd5,
	0x76, 0x3e, 0xb7, 0x8b, 0x0d, 0x31, 0xde, 0x6f, 0x88, 0xee, 0x26, 0x3d, 0x03, 0xae, 0x95, 0x11,
	0x39, 0x8c, 0x3c, 0x15, 0x4b, 0x57, 0x85, 0x95, 0x64, 0x49, 0x94, 0xed, 0xa8, 0x83, 0x49, 0x87,
	0x35, 0xed, 0x0d, 0x3f, 0xa9, 0xf4, 0x9d, 0x01, 0x37, 0x4a, 0xff, 0x3d, 0x29, 0x2e, 0x60, 0xcf,
	0x7e, 0x19, 0x5e, 0xfa, 0x67, 0x19, 0xad, 0x73, 0x47, 0xdd, 0x5c, 0xe7, 0x0d, 0xd8, 0xc3, 0x7b,
	0x05, 0xb3, 0xaa, 0x94, 0xfc, 0xf7, 0x05, 0xb3, 0xaa, 0x4c, 0xa7, 0x10, 0x23, 0xbd, 0x2e, 0xc9,
	0x07, 0x71, 0x02, 0xbc, 0x50, 0xbe, 0x35, 0xdf, 0x14, 0xf4, 0xfe, 0x56, 0x2b, 0x83, 0x0d, 0x93,
	0xbe, 0x31, 0x18, 0x21, 0xf9, 0xda, 0x2e, 0x3c, 0x89, 0x63, 0x88, 0x1d, 0x85, 0xfb, 0x55, 0x4d,
	0xeb, 0xec, 0xf0, 0x7c, 0x70, 0x3a, 0xcd, 0x73, 0xdc, 0x2e, 0xc5, 0x21, 0xfc, 0x73, 0x14, 0x6e,
	0xfc, 0x93, 0x8c, 0x12, 0x96, 0xfd, 0xc7, 0x76, 0x12, 0x12, 0x62, 0x72, 0x4e, 0xdb, 0x92, 0x24,
	0x4f, 0x58, 0x36, 0xc4, 0xed, 0xd8, 0x98, 0x78, 0x55, 0xc8, 0x41, 0xc2, 0x7e, 0x32, 0x31, 0x4a,
	0x63, 0xc3, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0x24, 0x6a, 0xf0, 0xb7, 0x03, 0x02, 0x00, 0x00,
}
