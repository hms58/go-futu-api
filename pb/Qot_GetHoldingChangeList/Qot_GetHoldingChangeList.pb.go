// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Qot_GetHoldingChangeList.proto

package Qot_GetHoldingChangeList

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
	Security       *Security `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	HolderCategory *int32    `protobuf:"varint,2,req,name=holderCategory" json:"holderCategory,omitempty"`
	//以下是发布时间筛选，不传返回所有数据，传了返回发布时间属于开始时间到结束时间段内的数据
	BeginTime            *string  `protobuf:"bytes,3,opt,name=beginTime" json:"beginTime,omitempty"`
	EndTime              *string  `protobuf:"bytes,4,opt,name=endTime" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2S) Reset()         { *m = C2S{} }
func (m *C2S) String() string { return proto.CompactTextString(m) }
func (*C2S) ProtoMessage()    {}
func (*C2S) Descriptor() ([]byte, []int) {
	return fileDescriptor_02515541775ad748, []int{0}
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

func (m *C2S) GetHolderCategory() int32 {
	if m != nil && m.HolderCategory != nil {
		return *m.HolderCategory
	}
	return 0
}

func (m *C2S) GetBeginTime() string {
	if m != nil && m.BeginTime != nil {
		return *m.BeginTime
	}
	return ""
}

func (m *C2S) GetEndTime() string {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return ""
}

type S2C struct {
	Security             *Security             `protobuf:"bytes,1,req,name=security" json:"security,omitempty"`
	HoldingChangeList    []*ShareHoldingChange `protobuf:"bytes,2,rep,name=holdingChangeList" json:"holdingChangeList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *S2C) Reset()         { *m = S2C{} }
func (m *S2C) String() string { return proto.CompactTextString(m) }
func (*S2C) ProtoMessage()    {}
func (*S2C) Descriptor() ([]byte, []int) {
	return fileDescriptor_02515541775ad748, []int{1}
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

func (m *S2C) GetHoldingChangeList() []*ShareHoldingChange {
	if m != nil {
		return m.HoldingChangeList
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
	return fileDescriptor_02515541775ad748, []int{2}
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
	return fileDescriptor_02515541775ad748, []int{3}
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
	proto.RegisterType((*C2S)(nil), "Qot_GetHoldingChangeList.C2S")
	proto.RegisterType((*S2C)(nil), "Qot_GetHoldingChangeList.S2C")
	proto.RegisterType((*Request)(nil), "Qot_GetHoldingChangeList.Request")
	proto.RegisterType((*Response)(nil), "Qot_GetHoldingChangeList.Response")
}

func init() { proto.RegisterFile("Qot_GetHoldingChangeList.proto", fileDescriptor_02515541775ad748) }

var fileDescriptor_02515541775ad748 = []byte{
	// 312 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x3f, 0x4f, 0xfa, 0x40,
	0x18, 0xc7, 0xd3, 0x16, 0x7e, 0xc0, 0xc3, 0x2f, 0x46, 0x2f, 0xc6, 0x5c, 0x88, 0x92, 0xa6, 0x83,
	0xe9, 0x22, 0x92, 0x8b, 0x13, 0xeb, 0x0d, 0x3a, 0xe0, 0xe0, 0x95, 0xdd, 0x20, 0x7d, 0xd2, 0x36,
	0x81, 0x3b, 0xbc, 0x3b, 0x86, 0xbe, 0x00, 0x47, 0x67, 0xdf, 0xae, 0xe9, 0xd1, 0x2a, 0x60, 0x70,
	0x70, 0xfc, 0xfe, 0x79, 0x2e, 0xdf, 0x7c, 0x0e, 0x86, 0x4f, 0xca, 0x3e, 0xdf, 0xa3, 0x7d, 0x50,
	0xcb, 0xb4, 0x90, 0x19, 0xcf, 0xe7, 0x32, 0xc3, 0x69, 0x61, 0xec, 0x68, 0xad, 0x95, 0x55, 0x84,
	0x1e, 0xcb, 0x07, 0xff, 0xb9, 0x5a, 0xad, 0x94, 0xdc, 0xf6, 0x06, 0xa7, 0x55, 0x6f, 0xd7, 0x89,
	0x3e, 0x3c, 0x08, 0x38, 0x4b, 0xc8, 0x18, 0xba, 0x06, 0x17, 0x1b, 0x5d, 0xd8, 0x92, 0x7a, 0xa1,
	0x1f, 0xf7, 0xd9, 0xf9, 0x68, 0xa7, 0x9c, 0xd4, 0x99, 0xf8, 0x6a, 0x91, 0x6b, 0x38, 0xc9, 0xd5,
	0x32, 0x45, 0xcd, 0xe7, 0x16, 0x33, 0xa5, 0x4b, 0xea, 0x87, 0x7e, 0xdc, 0x16, 0x07, 0x2e, 0xb9,
	0x84, 0xde, 0x0b, 0x66, 0x85, 0x9c, 0x15, 0x2b, 0xa4, 0x41, 0xe8, 0xc5, 0x3d, 0xf1, 0x6d, 0x10,
	0x0a, 0x1d, 0x94, 0xa9, 0xcb, 0x5a, 0x2e, 0x6b, 0x64, 0xf4, 0xe6, 0x41, 0x90, 0x30, 0xfe, 0x87,
	0x65, 0x53, 0x38, 0xcb, 0x0f, 0x41, 0x50, 0x3f, 0x0c, 0xe2, 0x3e, 0x1b, 0xee, 0x9d, 0xe6, 0x73,
	0x8d, 0x7b, 0xc8, 0xc4, 0xcf, 0xc3, 0x68, 0x02, 0x1d, 0x81, 0xaf, 0x1b, 0x34, 0x96, 0xdc, 0x42,
	0xb0, 0x60, 0xa6, 0x5e, 0x71, 0x35, 0x3a, 0xfa, 0x29, 0x9c, 0x25, 0xa2, 0x6a, 0x46, 0xef, 0x1e,
	0x74, 0x05, 0x9a, 0xb5, 0x92, 0x06, 0xc9, 0x10, 0x3a, 0x1a, 0xed, 0xac, 0x5c, 0xa3, 0x7b, 0xa1,
	0x3d, 0x69, 0xdd, 0xdc, 0x8d, 0xc7, 0xa2, 0x31, 0xc9, 0x05, 0xfc, 0xd3, 0x68, 0x1f, 0x4d, 0x46,
	0x7d, 0x47, 0xa2, 0x56, 0x0e, 0x91, 0xd6, 0x5c, 0xa5, 0x5b, 0x7c, 0x6d, 0xd1, 0xc8, 0x6a, 0x8f,
	0x61, 0x0b, 0x07, 0xee, 0xd7, 0x3d, 0x09, 0xe3, 0xa2, 0x6a, 0x7e, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x2e, 0x10, 0x11, 0xc3, 0x48, 0x02, 0x00, 0x00,
}
