// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: srv/broker/proto/message.proto

package micro_arch_srv_broker

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Event int32

const (
	Event_OK       Event = 0
	Event_Register Event = 1
)

var Event_name = map[int32]string{
	0: "OK",
	1: "Register",
}

var Event_value = map[string]int32{
	"OK":       0,
	"Register": 1,
}

func (x Event) String() string {
	return proto.EnumName(Event_name, int32(x))
}

func (Event) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3a462ab7a12951cb, []int{0}
}

type BaseMessage struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseMessage) Reset()         { *m = BaseMessage{} }
func (m *BaseMessage) String() string { return proto.CompactTextString(m) }
func (*BaseMessage) ProtoMessage()    {}
func (*BaseMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a462ab7a12951cb, []int{0}
}
func (m *BaseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseMessage.Unmarshal(m, b)
}
func (m *BaseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseMessage.Marshal(b, m, deterministic)
}
func (m *BaseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseMessage.Merge(m, src)
}
func (m *BaseMessage) XXX_Size() int {
	return xxx_messageInfo_BaseMessage.Size(m)
}
func (m *BaseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BaseMessage proto.InternalMessageInfo

func (m *BaseMessage) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BaseMessage) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *BaseMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

//Registrant
type ConfirmEmail struct {
	Msg                  *BaseMessage `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Username             string       `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	AuthId               string       `protobuf:"bytes,3,opt,name=authId,proto3" json:"authId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConfirmEmail) Reset()         { *m = ConfirmEmail{} }
func (m *ConfirmEmail) String() string { return proto.CompactTextString(m) }
func (*ConfirmEmail) ProtoMessage()    {}
func (*ConfirmEmail) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a462ab7a12951cb, []int{1}
}
func (m *ConfirmEmail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmEmail.Unmarshal(m, b)
}
func (m *ConfirmEmail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmEmail.Marshal(b, m, deterministic)
}
func (m *ConfirmEmail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmEmail.Merge(m, src)
}
func (m *ConfirmEmail) XXX_Size() int {
	return xxx_messageInfo_ConfirmEmail.Size(m)
}
func (m *ConfirmEmail) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmEmail.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmEmail proto.InternalMessageInfo

func (m *ConfirmEmail) GetMsg() *BaseMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ConfirmEmail) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ConfirmEmail) GetAuthId() string {
	if m != nil {
		return m.AuthId
	}
	return ""
}

type ConfirmMobile struct {
	Msg                  *BaseMessage `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	OneTimeToken         string       `protobuf:"bytes,2,opt,name=oneTimeToken,proto3" json:"oneTimeToken,omitempty"`
	Mobile               string       `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConfirmMobile) Reset()         { *m = ConfirmMobile{} }
func (m *ConfirmMobile) String() string { return proto.CompactTextString(m) }
func (*ConfirmMobile) ProtoMessage()    {}
func (*ConfirmMobile) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a462ab7a12951cb, []int{2}
}
func (m *ConfirmMobile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmMobile.Unmarshal(m, b)
}
func (m *ConfirmMobile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmMobile.Marshal(b, m, deterministic)
}
func (m *ConfirmMobile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmMobile.Merge(m, src)
}
func (m *ConfirmMobile) XXX_Size() int {
	return xxx_messageInfo_ConfirmMobile.Size(m)
}
func (m *ConfirmMobile) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmMobile.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmMobile proto.InternalMessageInfo

func (m *ConfirmMobile) GetMsg() *BaseMessage {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ConfirmMobile) GetOneTimeToken() string {
	if m != nil {
		return m.OneTimeToken
	}
	return ""
}

func (m *ConfirmMobile) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func init() {
	proto.RegisterEnum("micro.arch.srv.broker.Event", Event_name, Event_value)
	proto.RegisterType((*BaseMessage)(nil), "micro.arch.srv.broker.BaseMessage")
	proto.RegisterType((*ConfirmEmail)(nil), "micro.arch.srv.broker.ConfirmEmail")
	proto.RegisterType((*ConfirmMobile)(nil), "micro.arch.srv.broker.ConfirmMobile")
}

func init() { proto.RegisterFile("srv/broker/proto/message.proto", fileDescriptor_3a462ab7a12951cb) }

var fileDescriptor_3a462ab7a12951cb = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0xa2, 0xb1, 0x9d, 0x46, 0x29, 0x03, 0x4a, 0x10, 0x94, 0x92, 0x53, 0xf1, 0xb0,
	0x01, 0xf5, 0x09, 0x94, 0x1e, 0xa4, 0x14, 0x21, 0xf4, 0x05, 0x36, 0xcd, 0x98, 0x2e, 0xed, 0x66,
	0x65, 0x36, 0x0d, 0x5e, 0x7d, 0x73, 0xe9, 0x66, 0x2b, 0x0a, 0x9e, 0x7a, 0xdb, 0x6f, 0xf9, 0xf9,
	0xbf, 0x1f, 0x06, 0xee, 0x2c, 0x77, 0x79, 0xc9, 0x66, 0x43, 0x9c, 0x7f, 0xb0, 0x69, 0x4d, 0xae,
	0xc9, 0x5a, 0x59, 0x93, 0x70, 0x84, 0x57, 0x5a, 0xad, 0xd8, 0x08, 0xc9, 0xab, 0xb5, 0xb0, 0xdc,
	0x89, 0x3e, 0x9a, 0xcd, 0x61, 0xf4, 0x2c, 0x2d, 0x2d, 0xfa, 0x2c, 0x5e, 0x42, 0xa8, 0xaa, 0x34,
	0x98, 0x04, 0xd3, 0x61, 0x11, 0xaa, 0x0a, 0x11, 0x4e, 0x5b, 0xa5, 0x29, 0x0d, 0x27, 0xc1, 0x34,
	0x2a, 0xdc, 0x1b, 0x53, 0x38, 0xf7, 0xd5, 0x69, 0xe4, 0x82, 0x07, 0xcc, 0x3e, 0x21, 0x79, 0x31,
	0xcd, 0xbb, 0x62, 0x3d, 0xd3, 0x52, 0x6d, 0xf1, 0x09, 0x22, 0x6d, 0x6b, 0x57, 0x37, 0x7a, 0xc8,
	0xc4, 0xbf, 0x0b, 0xc4, 0x2f, 0x7d, 0xb1, 0x8f, 0xe3, 0x0d, 0x0c, 0x76, 0x96, 0xb8, 0x91, 0xde,
	0x3b, 0x2c, 0x7e, 0x18, 0xaf, 0x21, 0x96, 0xbb, 0x76, 0xfd, 0x5a, 0x79, 0xb5, 0xa7, 0xec, 0x2b,
	0x80, 0x0b, 0xaf, 0x5e, 0x98, 0x52, 0x6d, 0xe9, 0x48, 0x77, 0x06, 0x89, 0x69, 0x68, 0xa9, 0x34,
	0x2d, 0xcd, 0x86, 0x1a, 0xef, 0xff, 0xf3, 0xb7, 0xdf, 0xa0, 0x9d, 0xe3, 0xb0, 0xa1, 0xa7, 0xfb,
	0x5b, 0x38, 0x9b, 0x75, 0xd4, 0xb4, 0x18, 0x43, 0xf8, 0x36, 0x1f, 0x9f, 0x60, 0x02, 0x83, 0x82,
	0x6a, 0x65, 0x5b, 0xe2, 0x71, 0x50, 0xc6, 0xee, 0x0e, 0x8f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x92, 0x46, 0xa2, 0xc8, 0xa9, 0x01, 0x00, 0x00,
}