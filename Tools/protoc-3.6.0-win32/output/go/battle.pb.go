// Code generated by protoc-gen-go. DO NOT EDIT.
// source: battle.proto

package msg

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

type BattleStartRequest struct {
	BattleId             string   `protobuf:"bytes,1,opt,name=battleId,proto3" json:"battleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BattleStartRequest) Reset()         { *m = BattleStartRequest{} }
func (m *BattleStartRequest) String() string { return proto.CompactTextString(m) }
func (*BattleStartRequest) ProtoMessage()    {}
func (*BattleStartRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{0}
}

func (m *BattleStartRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleStartRequest.Unmarshal(m, b)
}
func (m *BattleStartRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleStartRequest.Marshal(b, m, deterministic)
}
func (m *BattleStartRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleStartRequest.Merge(m, src)
}
func (m *BattleStartRequest) XXX_Size() int {
	return xxx_messageInfo_BattleStartRequest.Size(m)
}
func (m *BattleStartRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleStartRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BattleStartRequest proto.InternalMessageInfo

func (m *BattleStartRequest) GetBattleId() string {
	if m != nil {
		return m.BattleId
	}
	return ""
}

type BattleStartResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Heros                []*Hero      `protobuf:"bytes,3,rep,name=heros,proto3" json:"heros,omitempty"`
	Skills               []*Skill     `protobuf:"bytes,4,rep,name=skills,proto3" json:"skills,omitempty"`
	Items                []*Item      `protobuf:"bytes,5,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BattleStartResponse) Reset()         { *m = BattleStartResponse{} }
func (m *BattleStartResponse) String() string { return proto.CompactTextString(m) }
func (*BattleStartResponse) ProtoMessage()    {}
func (*BattleStartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{1}
}

func (m *BattleStartResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleStartResponse.Unmarshal(m, b)
}
func (m *BattleStartResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleStartResponse.Marshal(b, m, deterministic)
}
func (m *BattleStartResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleStartResponse.Merge(m, src)
}
func (m *BattleStartResponse) XXX_Size() int {
	return xxx_messageInfo_BattleStartResponse.Size(m)
}
func (m *BattleStartResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleStartResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BattleStartResponse proto.InternalMessageInfo

func (m *BattleStartResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *BattleStartResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *BattleStartResponse) GetHeros() []*Hero {
	if m != nil {
		return m.Heros
	}
	return nil
}

func (m *BattleStartResponse) GetSkills() []*Skill {
	if m != nil {
		return m.Skills
	}
	return nil
}

func (m *BattleStartResponse) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type BattleGuanKaRequest struct {
	GuanKaId             int32    `protobuf:"varint,1,opt,name=guanKaId,proto3" json:"guanKaId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BattleGuanKaRequest) Reset()         { *m = BattleGuanKaRequest{} }
func (m *BattleGuanKaRequest) String() string { return proto.CompactTextString(m) }
func (*BattleGuanKaRequest) ProtoMessage()    {}
func (*BattleGuanKaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{2}
}

func (m *BattleGuanKaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleGuanKaRequest.Unmarshal(m, b)
}
func (m *BattleGuanKaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleGuanKaRequest.Marshal(b, m, deterministic)
}
func (m *BattleGuanKaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleGuanKaRequest.Merge(m, src)
}
func (m *BattleGuanKaRequest) XXX_Size() int {
	return xxx_messageInfo_BattleGuanKaRequest.Size(m)
}
func (m *BattleGuanKaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleGuanKaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BattleGuanKaRequest proto.InternalMessageInfo

func (m *BattleGuanKaRequest) GetGuanKaId() int32 {
	if m != nil {
		return m.GuanKaId
	}
	return 0
}

type BattleGuanKaResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Result               int32        `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
	Guanka               *GuanKa      `protobuf:"bytes,4,opt,name=guanka,proto3" json:"guanka,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BattleGuanKaResponse) Reset()         { *m = BattleGuanKaResponse{} }
func (m *BattleGuanKaResponse) String() string { return proto.CompactTextString(m) }
func (*BattleGuanKaResponse) ProtoMessage()    {}
func (*BattleGuanKaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{3}
}

func (m *BattleGuanKaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleGuanKaResponse.Unmarshal(m, b)
}
func (m *BattleGuanKaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleGuanKaResponse.Marshal(b, m, deterministic)
}
func (m *BattleGuanKaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleGuanKaResponse.Merge(m, src)
}
func (m *BattleGuanKaResponse) XXX_Size() int {
	return xxx_messageInfo_BattleGuanKaResponse.Size(m)
}
func (m *BattleGuanKaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleGuanKaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BattleGuanKaResponse proto.InternalMessageInfo

func (m *BattleGuanKaResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *BattleGuanKaResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *BattleGuanKaResponse) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *BattleGuanKaResponse) GetGuanka() *GuanKa {
	if m != nil {
		return m.Guanka
	}
	return nil
}

func init() {
	proto.RegisterType((*BattleStartRequest)(nil), "msg.BattleStartRequest")
	proto.RegisterType((*BattleStartResponse)(nil), "msg.BattleStartResponse")
	proto.RegisterType((*BattleGuanKaRequest)(nil), "msg.BattleGuanKaRequest")
	proto.RegisterType((*BattleGuanKaResponse)(nil), "msg.BattleGuanKaResponse")
}

func init() { proto.RegisterFile("battle.proto", fileDescriptor_88d4c72b5869c382) }

var fileDescriptor_88d4c72b5869c382 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x51, 0xd1, 0x4a, 0xc3, 0x40,
	0x10, 0x24, 0xa6, 0x0d, 0x66, 0x23, 0x82, 0xa7, 0xc8, 0x51, 0x04, 0x4b, 0x44, 0xe8, 0x53, 0xd1,
	0xfa, 0x07, 0x8a, 0x68, 0xf1, 0xed, 0xfa, 0x05, 0x57, 0xb3, 0xc4, 0xd2, 0xa4, 0x17, 0xf7, 0x2e,
	0xff, 0xe2, 0xe7, 0xf8, 0x69, 0x72, 0x7b, 0x97, 0xd2, 0xbe, 0xfb, 0xb6, 0x33, 0xb3, 0x33, 0xb7,
	0xc3, 0xc1, 0xd9, 0x5a, 0x3b, 0xd7, 0xe0, 0xbc, 0x23, 0xe3, 0x8c, 0x48, 0x5b, 0x5b, 0x4f, 0x72,
	0x24, 0x0a, 0x78, 0x92, 0xb7, 0xba, 0x8b, 0x23, 0x7c, 0x21, 0x99, 0x38, 0x17, 0x76, 0xbb, 0x69,
	0x9a, 0x41, 0xd8, 0x38, 0x6c, 0xc3, 0x5c, 0x3e, 0x80, 0x78, 0xe6, 0xbc, 0x95, 0xd3, 0xe4, 0x14,
	0x7e, 0xf7, 0x68, 0x9d, 0x98, 0xc0, 0x69, 0x78, 0x65, 0x59, 0xc9, 0x64, 0x9a, 0xcc, 0x72, 0xb5,
	0xc7, 0xe5, 0x6f, 0x02, 0x97, 0x47, 0x16, 0xdb, 0x99, 0x9d, 0x45, 0x71, 0x0f, 0xa3, 0x4f, 0x53,
	0x21, 0xef, 0x9f, 0x2f, 0x2e, 0xe6, 0xad, 0xad, 0xe7, 0x83, 0xf8, 0x62, 0x2a, 0x54, 0x2c, 0x8b,
	0x1b, 0x48, 0x91, 0x48, 0x9e, 0x4c, 0x93, 0x59, 0xb1, 0x00, 0xde, 0x7a, 0x25, 0x32, 0xa4, 0x3c,
	0x2d, 0x6e, 0x61, 0xec, 0xaf, 0xb6, 0x32, 0x9d, 0xa6, 0xb3, 0x62, 0x91, 0xb3, 0xfe, 0x8e, 0x64,
	0x54, 0xe0, 0x45, 0x09, 0x19, 0x57, 0xb1, 0x72, 0xc4, 0x1b, 0x21, 0x61, 0xe5, 0x29, 0x15, 0x15,
	0x1f, 0xe2, 0x1b, 0x5a, 0x39, 0x3e, 0x08, 0x59, 0x3a, 0x6c, 0x55, 0xe0, 0xcb, 0xc7, 0xa1, 0xc1,
	0x5b, 0xaf, 0x77, 0x1f, 0xfa, 0xa0, 0x75, 0xcd, 0x44, 0x6c, 0x3d, 0x56, 0x7b, 0x5c, 0xfe, 0x24,
	0x70, 0x75, 0xec, 0xf9, 0xcf, 0xda, 0xd7, 0x90, 0x11, 0xda, 0xbe, 0x71, 0x32, 0xe5, 0x77, 0x23,
	0x12, 0x77, 0x90, 0xf9, 0x0b, 0xb6, 0x5a, 0x8e, 0xd8, 0x58, 0xb0, 0x31, 0x5e, 0x10, 0xa5, 0x75,
	0xc6, 0x3f, 0xf9, 0xf4, 0x17, 0x00, 0x00, 0xff, 0xff, 0x87, 0x98, 0x6a, 0x2d, 0x19, 0x02, 0x00,
	0x00,
}
