// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s2s.proto

package s2s

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

// client  <-> server  10000 - 19999
type MessageCmd int32

const (
	MessageCmd_NULL       MessageCmd = 0
	MessageCmd_GetUID     MessageCmd = 1
	MessageCmd_GetUIDList MessageCmd = 2
)

var MessageCmd_name = map[int32]string{
	0: "NULL",
	1: "GetUID",
	2: "GetUIDList",
}

var MessageCmd_value = map[string]int32{
	"NULL":       0,
	"GetUID":     1,
	"GetUIDList": 2,
}

func (x MessageCmd) String() string {
	return proto.EnumName(MessageCmd_name, int32(x))
}

func (MessageCmd) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_953aa047daafed82, []int{0}
}

type ErrorCode int32

const (
	ErrorCode_OK ErrorCode = 0
)

var ErrorCode_name = map[int32]string{
	0: "OK",
}

var ErrorCode_value = map[string]int32{
	"OK": 0,
}

func (x ErrorCode) String() string {
	return proto.EnumName(ErrorCode_name, int32(x))
}

func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_953aa047daafed82, []int{1}
}

type Result struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_953aa047daafed82, []int{0}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Result) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type GenUIDReq struct {
	Channel              int32    `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenUIDReq) Reset()         { *m = GenUIDReq{} }
func (m *GenUIDReq) String() string { return proto.CompactTextString(m) }
func (*GenUIDReq) ProtoMessage()    {}
func (*GenUIDReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_953aa047daafed82, []int{1}
}

func (m *GenUIDReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenUIDReq.Unmarshal(m, b)
}
func (m *GenUIDReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenUIDReq.Marshal(b, m, deterministic)
}
func (m *GenUIDReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenUIDReq.Merge(m, src)
}
func (m *GenUIDReq) XXX_Size() int {
	return xxx_messageInfo_GenUIDReq.Size(m)
}
func (m *GenUIDReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GenUIDReq.DiscardUnknown(m)
}

var xxx_messageInfo_GenUIDReq proto.InternalMessageInfo

func (m *GenUIDReq) GetChannel() int32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

type GenUIDResp struct {
	Result               *Result  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Uid                  uint64   `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenUIDResp) Reset()         { *m = GenUIDResp{} }
func (m *GenUIDResp) String() string { return proto.CompactTextString(m) }
func (*GenUIDResp) ProtoMessage()    {}
func (*GenUIDResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_953aa047daafed82, []int{2}
}

func (m *GenUIDResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenUIDResp.Unmarshal(m, b)
}
func (m *GenUIDResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenUIDResp.Marshal(b, m, deterministic)
}
func (m *GenUIDResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenUIDResp.Merge(m, src)
}
func (m *GenUIDResp) XXX_Size() int {
	return xxx_messageInfo_GenUIDResp.Size(m)
}
func (m *GenUIDResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GenUIDResp.DiscardUnknown(m)
}

var xxx_messageInfo_GenUIDResp proto.InternalMessageInfo

func (m *GenUIDResp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GenUIDResp) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

type GenUIDsReq struct {
	Channel              int32    `protobuf:"varint,1,opt,name=channel,proto3" json:"channel,omitempty"`
	Num                  int32    `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenUIDsReq) Reset()         { *m = GenUIDsReq{} }
func (m *GenUIDsReq) String() string { return proto.CompactTextString(m) }
func (*GenUIDsReq) ProtoMessage()    {}
func (*GenUIDsReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_953aa047daafed82, []int{3}
}

func (m *GenUIDsReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenUIDsReq.Unmarshal(m, b)
}
func (m *GenUIDsReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenUIDsReq.Marshal(b, m, deterministic)
}
func (m *GenUIDsReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenUIDsReq.Merge(m, src)
}
func (m *GenUIDsReq) XXX_Size() int {
	return xxx_messageInfo_GenUIDsReq.Size(m)
}
func (m *GenUIDsReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GenUIDsReq.DiscardUnknown(m)
}

var xxx_messageInfo_GenUIDsReq proto.InternalMessageInfo

func (m *GenUIDsReq) GetChannel() int32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *GenUIDsReq) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type GenUIDsResp struct {
	Result               *Result  `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Uids                 []uint64 `protobuf:"varint,2,rep,packed,name=uids,proto3" json:"uids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenUIDsResp) Reset()         { *m = GenUIDsResp{} }
func (m *GenUIDsResp) String() string { return proto.CompactTextString(m) }
func (*GenUIDsResp) ProtoMessage()    {}
func (*GenUIDsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_953aa047daafed82, []int{4}
}

func (m *GenUIDsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenUIDsResp.Unmarshal(m, b)
}
func (m *GenUIDsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenUIDsResp.Marshal(b, m, deterministic)
}
func (m *GenUIDsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenUIDsResp.Merge(m, src)
}
func (m *GenUIDsResp) XXX_Size() int {
	return xxx_messageInfo_GenUIDsResp.Size(m)
}
func (m *GenUIDsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GenUIDsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GenUIDsResp proto.InternalMessageInfo

func (m *GenUIDsResp) GetResult() *Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GenUIDsResp) GetUids() []uint64 {
	if m != nil {
		return m.Uids
	}
	return nil
}

func init() {
	proto.RegisterEnum("s2s.MessageCmd", MessageCmd_name, MessageCmd_value)
	proto.RegisterEnum("s2s.ErrorCode", ErrorCode_name, ErrorCode_value)
	proto.RegisterType((*Result)(nil), "s2s.Result")
	proto.RegisterType((*GenUIDReq)(nil), "s2s.GenUIDReq")
	proto.RegisterType((*GenUIDResp)(nil), "s2s.GenUIDResp")
	proto.RegisterType((*GenUIDsReq)(nil), "s2s.GenUIDsReq")
	proto.RegisterType((*GenUIDsResp)(nil), "s2s.GenUIDsResp")
}

func init() { proto.RegisterFile("s2s.proto", fileDescriptor_953aa047daafed82) }

var fileDescriptor_953aa047daafed82 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xcf, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x9b, 0x5f, 0x5b, 0x33, 0x01, 0x59, 0xc6, 0x4b, 0x8e, 0x61, 0x45, 0x08, 0x3d, 0xe4,
	0xb0, 0x5e, 0x04, 0x6f, 0xa6, 0x5a, 0x8a, 0x51, 0x61, 0x21, 0x17, 0x6f, 0xb5, 0x59, 0x6a, 0xa0,
	0x49, 0xea, 0x4e, 0xf2, 0xff, 0xcb, 0xae, 0x8d, 0x47, 0xf1, 0xf6, 0x2d, 0x3b, 0x6f, 0xbe, 0xe1,
	0x41, 0x4c, 0x92, 0x8a, 0x93, 0x19, 0xc6, 0x01, 0x03, 0x92, 0x24, 0x0a, 0x60, 0x4a, 0xd3, 0x74,
	0x1c, 0x11, 0x21, 0xdc, 0x0f, 0x8d, 0x4e, 0xbd, 0xcc, 0xcb, 0x23, 0xe5, 0x18, 0x39, 0x04, 0x1d,
	0x1d, 0x52, 0x3f, 0xf3, 0xf2, 0x58, 0x59, 0x14, 0x37, 0x10, 0x6f, 0x74, 0x5f, 0x6f, 0xd7, 0x4a,
	0x7f, 0x61, 0x0a, 0xcb, 0xfd, 0xe7, 0xae, 0xef, 0xf5, 0xf1, 0x9c, 0x9a, 0x9f, 0xa2, 0x04, 0x98,
	0xc7, 0xe8, 0x84, 0xd7, 0xc0, 0x8c, 0x93, 0xb8, 0xb1, 0x44, 0x26, 0x85, 0xbd, 0xe2, 0xc7, 0xab,
	0xce, 0x5f, 0xd6, 0x35, 0xb5, 0x8d, 0x73, 0x85, 0xca, 0xa2, 0xb8, 0x9b, 0x97, 0xd0, 0x9f, 0x32,
	0x9b, 0xec, 0xa7, 0xce, 0x25, 0x23, 0x65, 0x51, 0x3c, 0x41, 0xf2, 0x9b, 0xfc, 0xaf, 0x1f, 0x21,
	0x9c, 0xda, 0x86, 0x52, 0x3f, 0x0b, 0xf2, 0x50, 0x39, 0x5e, 0x49, 0x80, 0x17, 0x4d, 0xb4, 0x3b,
	0xe8, 0xb2, 0x6b, 0xf0, 0x02, 0xc2, 0xd7, 0xba, 0xaa, 0xf8, 0x02, 0x01, 0xd8, 0x46, 0x8f, 0xf5,
	0x76, 0xcd, 0x3d, 0xbc, 0xb4, 0x57, 0x5a, 0xae, 0x5a, 0x1a, 0xb9, 0xbf, 0xba, 0x82, 0xf8, 0xd1,
	0x98, 0xc1, 0x94, 0xb6, 0x40, 0x06, 0xfe, 0xdb, 0x33, 0x5f, 0x3c, 0x2c, 0xdf, 0xa3, 0xe2, 0x9e,
	0x24, 0x7d, 0x30, 0xd7, 0xfd, 0xed, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x70, 0x61, 0x0b,
	0x88, 0x01, 0x00, 0x00,
}
