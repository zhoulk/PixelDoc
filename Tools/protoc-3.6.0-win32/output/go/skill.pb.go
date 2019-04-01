// Code generated by protoc-gen-go. DO NOT EDIT.
// source: skill.proto

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

type SkillRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SkillRequest) Reset()         { *m = SkillRequest{} }
func (m *SkillRequest) String() string { return proto.CompactTextString(m) }
func (*SkillRequest) ProtoMessage()    {}
func (*SkillRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd709e691a520876, []int{0}
}

func (m *SkillRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkillRequest.Unmarshal(m, b)
}
func (m *SkillRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkillRequest.Marshal(b, m, deterministic)
}
func (m *SkillRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkillRequest.Merge(m, src)
}
func (m *SkillRequest) XXX_Size() int {
	return xxx_messageInfo_SkillRequest.Size(m)
}
func (m *SkillRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SkillRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SkillRequest proto.InternalMessageInfo

type SkillResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Skills               []*Skill     `protobuf:"bytes,3,rep,name=skills,proto3" json:"skills,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SkillResponse) Reset()         { *m = SkillResponse{} }
func (m *SkillResponse) String() string { return proto.CompactTextString(m) }
func (*SkillResponse) ProtoMessage()    {}
func (*SkillResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd709e691a520876, []int{1}
}

func (m *SkillResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkillResponse.Unmarshal(m, b)
}
func (m *SkillResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkillResponse.Marshal(b, m, deterministic)
}
func (m *SkillResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkillResponse.Merge(m, src)
}
func (m *SkillResponse) XXX_Size() int {
	return xxx_messageInfo_SkillResponse.Size(m)
}
func (m *SkillResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SkillResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SkillResponse proto.InternalMessageInfo

func (m *SkillResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *SkillResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *SkillResponse) GetSkills() []*Skill {
	if m != nil {
		return m.Skills
	}
	return nil
}

type SkillUpgradeRequest struct {
	SkillId              string   `protobuf:"bytes,1,opt,name=skillId,proto3" json:"skillId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SkillUpgradeRequest) Reset()         { *m = SkillUpgradeRequest{} }
func (m *SkillUpgradeRequest) String() string { return proto.CompactTextString(m) }
func (*SkillUpgradeRequest) ProtoMessage()    {}
func (*SkillUpgradeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd709e691a520876, []int{2}
}

func (m *SkillUpgradeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkillUpgradeRequest.Unmarshal(m, b)
}
func (m *SkillUpgradeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkillUpgradeRequest.Marshal(b, m, deterministic)
}
func (m *SkillUpgradeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkillUpgradeRequest.Merge(m, src)
}
func (m *SkillUpgradeRequest) XXX_Size() int {
	return xxx_messageInfo_SkillUpgradeRequest.Size(m)
}
func (m *SkillUpgradeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SkillUpgradeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SkillUpgradeRequest proto.InternalMessageInfo

func (m *SkillUpgradeRequest) GetSkillId() string {
	if m != nil {
		return m.SkillId
	}
	return ""
}

type SkillUpgradeResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Skill                *Skill       `protobuf:"bytes,3,opt,name=skill,proto3" json:"skill,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SkillUpgradeResponse) Reset()         { *m = SkillUpgradeResponse{} }
func (m *SkillUpgradeResponse) String() string { return proto.CompactTextString(m) }
func (*SkillUpgradeResponse) ProtoMessage()    {}
func (*SkillUpgradeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd709e691a520876, []int{3}
}

func (m *SkillUpgradeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkillUpgradeResponse.Unmarshal(m, b)
}
func (m *SkillUpgradeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkillUpgradeResponse.Marshal(b, m, deterministic)
}
func (m *SkillUpgradeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkillUpgradeResponse.Merge(m, src)
}
func (m *SkillUpgradeResponse) XXX_Size() int {
	return xxx_messageInfo_SkillUpgradeResponse.Size(m)
}
func (m *SkillUpgradeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SkillUpgradeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SkillUpgradeResponse proto.InternalMessageInfo

func (m *SkillUpgradeResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *SkillUpgradeResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *SkillUpgradeResponse) GetSkill() *Skill {
	if m != nil {
		return m.Skill
	}
	return nil
}

type Skill struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Level                int32    `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	Type                 int32    `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
	Desc                 string   `protobuf:"bytes,5,opt,name=Desc,proto3" json:"Desc,omitempty"`
	IsOpen               bool     `protobuf:"varint,6,opt,name=IsOpen,proto3" json:"IsOpen,omitempty"`
	SkillId              string   `protobuf:"bytes,7,opt,name=SkillId,proto3" json:"SkillId,omitempty"`
	HeroId               string   `protobuf:"bytes,8,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	LevelDesc            []string `protobuf:"bytes,9,rep,name=LevelDesc,proto3" json:"LevelDesc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Skill) Reset()         { *m = Skill{} }
func (m *Skill) String() string { return proto.CompactTextString(m) }
func (*Skill) ProtoMessage()    {}
func (*Skill) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd709e691a520876, []int{4}
}

func (m *Skill) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Skill.Unmarshal(m, b)
}
func (m *Skill) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Skill.Marshal(b, m, deterministic)
}
func (m *Skill) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Skill.Merge(m, src)
}
func (m *Skill) XXX_Size() int {
	return xxx_messageInfo_Skill.Size(m)
}
func (m *Skill) XXX_DiscardUnknown() {
	xxx_messageInfo_Skill.DiscardUnknown(m)
}

var xxx_messageInfo_Skill proto.InternalMessageInfo

func (m *Skill) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Skill) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Skill) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Skill) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Skill) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Skill) GetIsOpen() bool {
	if m != nil {
		return m.IsOpen
	}
	return false
}

func (m *Skill) GetSkillId() string {
	if m != nil {
		return m.SkillId
	}
	return ""
}

func (m *Skill) GetHeroId() string {
	if m != nil {
		return m.HeroId
	}
	return ""
}

func (m *Skill) GetLevelDesc() []string {
	if m != nil {
		return m.LevelDesc
	}
	return nil
}

func init() {
	proto.RegisterType((*SkillRequest)(nil), "msg.SkillRequest")
	proto.RegisterType((*SkillResponse)(nil), "msg.SkillResponse")
	proto.RegisterType((*SkillUpgradeRequest)(nil), "msg.SkillUpgradeRequest")
	proto.RegisterType((*SkillUpgradeResponse)(nil), "msg.SkillUpgradeResponse")
	proto.RegisterType((*Skill)(nil), "msg.Skill")
}

func init() { proto.RegisterFile("skill.proto", fileDescriptor_dd709e691a520876) }

var fileDescriptor_dd709e691a520876 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0x66, 0xbb, 0xcd, 0xb6, 0x99, 0x6a, 0xc1, 0x58, 0x24, 0x48, 0x0f, 0x21, 0x20, 0xec, 0xa9,
	0x42, 0x7d, 0x04, 0x15, 0x5c, 0x10, 0x85, 0x54, 0x1f, 0xa0, 0x36, 0x43, 0x11, 0xdb, 0x66, 0x4d,
	0xaa, 0xe8, 0xc1, 0x87, 0xf4, 0x8d, 0x24, 0x93, 0xd4, 0x9f, 0xbb, 0xb7, 0xf9, 0x7e, 0x26, 0xdf,
	0xc7, 0x10, 0x18, 0x84, 0xa7, 0xc7, 0xd5, 0x6a, 0xd2, 0x7a, 0xb7, 0x75, 0xa2, 0x5c, 0x87, 0xe5,
	0x31, 0x47, 0xef, 0x13, 0xd6, 0x43, 0xd8, 0x9b, 0x45, 0xd9, 0xe0, 0xf3, 0x0b, 0x86, 0xad, 0x7e,
	0x83, 0xfd, 0x8c, 0x43, 0xeb, 0x36, 0x01, 0xc5, 0x09, 0x74, 0x17, 0xce, 0xa2, 0x2c, 0x54, 0x51,
	0x0f, 0xa7, 0x07, 0x93, 0x75, 0x58, 0x4e, 0x76, 0xe2, 0xb9, 0xb3, 0x68, 0x48, 0x16, 0x63, 0x28,
	0xd1, 0x7b, 0xd9, 0x51, 0x45, 0x3d, 0x98, 0x02, 0xb9, 0x2e, 0xbd, 0x77, 0xde, 0x44, 0x5a, 0x68,
	0xa8, 0xa8, 0x44, 0x90, 0xa5, 0x2a, 0xbf, 0x0d, 0x29, 0x28, 0x2b, 0xfa, 0x14, 0x0e, 0x89, 0xb8,
	0x6f, 0x97, 0x7e, 0x6e, 0x31, 0x17, 0x12, 0x12, 0x7a, 0x64, 0x68, 0x2c, 0x55, 0xe0, 0x66, 0x07,
	0xf5, 0x07, 0x8c, 0xfe, 0x2e, 0xfc, 0x67, 0x63, 0x05, 0x8c, 0x72, 0x64, 0xf9, 0x4b, 0x4f, 0x85,
	0x93, 0xa0, 0x3f, 0x0b, 0x60, 0x44, 0x88, 0x21, 0x74, 0x72, 0x3b, 0x66, 0x3a, 0x8d, 0x15, 0x02,
	0xba, 0x37, 0xf3, 0x35, 0xd2, 0xd3, 0xdc, 0xd0, 0x2c, 0x46, 0xc0, 0xae, 0xf1, 0x15, 0xd3, 0x7b,
	0xcc, 0x24, 0x10, 0x9d, 0x77, 0xef, 0x2d, 0xca, 0x2e, 0x91, 0x34, 0x47, 0xee, 0x02, 0xc3, 0x42,
	0xb2, 0xb4, 0x1d, 0x67, 0x71, 0x04, 0x55, 0x13, 0x6e, 0x5b, 0xdc, 0xc8, 0x4a, 0x15, 0x75, 0xdf,
	0x64, 0x14, 0x8f, 0x33, 0xcb, 0xc7, 0xe9, 0xa5, 0xe3, 0x64, 0x18, 0x37, 0xae, 0xd0, 0xbb, 0xc6,
	0xca, 0x3e, 0x09, 0x19, 0x89, 0x31, 0x70, 0x8a, 0xa6, 0x08, 0xae, 0xca, 0x9a, 0x9b, 0x1f, 0xe2,
	0xa1, 0xa2, 0x4f, 0x71, 0xf6, 0x15, 0x00, 0x00, 0xff, 0xff, 0xd8, 0xd3, 0xdf, 0xd8, 0x33, 0x02,
	0x00, 0x00,
}
