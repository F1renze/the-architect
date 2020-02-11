// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: srv/auth/proto/message.proto

package micro_arch_srv_auth

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
	return fileDescriptor_178cf5ed681a0331, []int{0}
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
	return fileDescriptor_178cf5ed681a0331, []int{1}
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
	Mobile               string       `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ConfirmMobile) Reset()         { *m = ConfirmMobile{} }
func (m *ConfirmMobile) String() string { return proto.CompactTextString(m) }
func (*ConfirmMobile) ProtoMessage()    {}
func (*ConfirmMobile) Descriptor() ([]byte, []int) {
	return fileDescriptor_178cf5ed681a0331, []int{2}
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

func (m *ConfirmMobile) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func init() {
	proto.RegisterType((*BaseMessage)(nil), "micro.arch.srv.auth.BaseMessage")
	proto.RegisterType((*ConfirmEmail)(nil), "micro.arch.srv.auth.ConfirmEmail")
	proto.RegisterType((*ConfirmMobile)(nil), "micro.arch.srv.auth.ConfirmMobile")
}

func init() { proto.RegisterFile("srv/auth/proto/message.proto", fileDescriptor_178cf5ed681a0331) }

var fileDescriptor_178cf5ed681a0331 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0x41, 0x4b, 0x87, 0x30,
	0x18, 0x87, 0x51, 0xc3, 0xf2, 0xb5, 0x3a, 0xbc, 0x81, 0x8c, 0xe8, 0x20, 0x9e, 0x3c, 0x4d, 0xb0,
	0x6f, 0x50, 0x74, 0x88, 0xf0, 0xb2, 0x6b, 0xa7, 0xa9, 0x4b, 0x07, 0xce, 0xc5, 0xa6, 0x7e, 0xfe,
	0x70, 0xce, 0xe8, 0xd0, 0xe9, 0x7f, 0xdb, 0x03, 0xcf, 0xfb, 0x7b, 0x60, 0xf0, 0x64, 0xcd, 0x56,
	0xf1, 0x75, 0x19, 0xab, 0x6f, 0xa3, 0x17, 0x5d, 0x29, 0x61, 0x2d, 0x1f, 0x04, 0x75, 0x84, 0x0f,
	0x4a, 0x76, 0x46, 0x53, 0x6e, 0xba, 0x91, 0x5a, 0xb3, 0xd1, 0x5d, 0x2c, 0x3e, 0x20, 0x7d, 0xe1,
	0x56, 0x34, 0x87, 0x89, 0xf7, 0x10, 0xca, 0x9e, 0x04, 0x79, 0x50, 0x26, 0x2c, 0x94, 0x3d, 0x22,
	0x5c, 0x2d, 0x52, 0x09, 0x12, 0xe6, 0x41, 0x19, 0x31, 0xf7, 0x46, 0x02, 0xd7, 0x7e, 0x98, 0x44,
	0x4e, 0x3c, 0xb1, 0xd8, 0xe0, 0xf6, 0x55, 0xcf, 0x5f, 0xd2, 0xa8, 0x37, 0xc5, 0xe5, 0x84, 0x35,
	0x44, 0xca, 0x0e, 0x6e, 0x2e, 0xad, 0x73, 0xfa, 0x4f, 0x9f, 0xfe, 0x89, 0xb3, 0x5d, 0xc6, 0x47,
	0xb8, 0x59, 0xad, 0x30, 0x33, 0xf7, 0xd5, 0x84, 0xfd, 0x32, 0x66, 0x10, 0xef, 0x47, 0xef, 0xbd,
	0x0f, 0x7b, 0x2a, 0x3e, 0xe1, 0xce, 0x77, 0x1b, 0xdd, 0xca, 0x49, 0x5c, 0x14, 0xce, 0x20, 0x56,
	0xee, 0xfa, 0x1c, 0x3f, 0xa8, 0x8d, 0xdd, 0xef, 0x3d, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x8b,
	0x3f, 0x55, 0x8b, 0x5d, 0x01, 0x00, 0x00,
}