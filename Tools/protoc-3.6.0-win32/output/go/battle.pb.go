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

type BattleCreateRequest struct {
	Type                 int32    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Args                 []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BattleCreateRequest) Reset()         { *m = BattleCreateRequest{} }
func (m *BattleCreateRequest) String() string { return proto.CompactTextString(m) }
func (*BattleCreateRequest) ProtoMessage()    {}
func (*BattleCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{0}
}

func (m *BattleCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleCreateRequest.Unmarshal(m, b)
}
func (m *BattleCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleCreateRequest.Marshal(b, m, deterministic)
}
func (m *BattleCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleCreateRequest.Merge(m, src)
}
func (m *BattleCreateRequest) XXX_Size() int {
	return xxx_messageInfo_BattleCreateRequest.Size(m)
}
func (m *BattleCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BattleCreateRequest proto.InternalMessageInfo

func (m *BattleCreateRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *BattleCreateRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

type BattleCreateResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	BattleId             string       `protobuf:"bytes,3,opt,name=battleId,proto3" json:"battleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BattleCreateResponse) Reset()         { *m = BattleCreateResponse{} }
func (m *BattleCreateResponse) String() string { return proto.CompactTextString(m) }
func (*BattleCreateResponse) ProtoMessage()    {}
func (*BattleCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{1}
}

func (m *BattleCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleCreateResponse.Unmarshal(m, b)
}
func (m *BattleCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleCreateResponse.Marshal(b, m, deterministic)
}
func (m *BattleCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleCreateResponse.Merge(m, src)
}
func (m *BattleCreateResponse) XXX_Size() int {
	return xxx_messageInfo_BattleCreateResponse.Size(m)
}
func (m *BattleCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BattleCreateResponse proto.InternalMessageInfo

func (m *BattleCreateResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *BattleCreateResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *BattleCreateResponse) GetBattleId() string {
	if m != nil {
		return m.BattleId
	}
	return ""
}

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
	return fileDescriptor_88d4c72b5869c382, []int{2}
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
	Equips               []*Equip     `protobuf:"bytes,5,rep,name=equips,proto3" json:"equips,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BattleStartResponse) Reset()         { *m = BattleStartResponse{} }
func (m *BattleStartResponse) String() string { return proto.CompactTextString(m) }
func (*BattleStartResponse) ProtoMessage()    {}
func (*BattleStartResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{3}
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

func (m *BattleStartResponse) GetEquips() []*Equip {
	if m != nil {
		return m.Equips
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
	return fileDescriptor_88d4c72b5869c382, []int{4}
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
	return fileDescriptor_88d4c72b5869c382, []int{5}
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

type BattleResultRequest struct {
	BattleId             string   `protobuf:"bytes,1,opt,name=battleId,proto3" json:"battleId,omitempty"`
	Result               int32    `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BattleResultRequest) Reset()         { *m = BattleResultRequest{} }
func (m *BattleResultRequest) String() string { return proto.CompactTextString(m) }
func (*BattleResultRequest) ProtoMessage()    {}
func (*BattleResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{6}
}

func (m *BattleResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleResultRequest.Unmarshal(m, b)
}
func (m *BattleResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleResultRequest.Marshal(b, m, deterministic)
}
func (m *BattleResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleResultRequest.Merge(m, src)
}
func (m *BattleResultRequest) XXX_Size() int {
	return xxx_messageInfo_BattleResultRequest.Size(m)
}
func (m *BattleResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BattleResultRequest proto.InternalMessageInfo

func (m *BattleResultRequest) GetBattleId() string {
	if m != nil {
		return m.BattleId
	}
	return ""
}

func (m *BattleResultRequest) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

type BattleResultResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Earn                 *Earn        `protobuf:"bytes,3,opt,name=earn,proto3" json:"earn,omitempty"`
	Exp                  int32        `protobuf:"varint,4,opt,name=Exp,proto3" json:"Exp,omitempty"`
	Level                int32        `protobuf:"varint,5,opt,name=Level,proto3" json:"Level,omitempty"`
	LevelUpExp           int32        `protobuf:"varint,6,opt,name=LevelUpExp,proto3" json:"LevelUpExp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BattleResultResponse) Reset()         { *m = BattleResultResponse{} }
func (m *BattleResultResponse) String() string { return proto.CompactTextString(m) }
func (*BattleResultResponse) ProtoMessage()    {}
func (*BattleResultResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_88d4c72b5869c382, []int{7}
}

func (m *BattleResultResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BattleResultResponse.Unmarshal(m, b)
}
func (m *BattleResultResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BattleResultResponse.Marshal(b, m, deterministic)
}
func (m *BattleResultResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BattleResultResponse.Merge(m, src)
}
func (m *BattleResultResponse) XXX_Size() int {
	return xxx_messageInfo_BattleResultResponse.Size(m)
}
func (m *BattleResultResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BattleResultResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BattleResultResponse proto.InternalMessageInfo

func (m *BattleResultResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *BattleResultResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *BattleResultResponse) GetEarn() *Earn {
	if m != nil {
		return m.Earn
	}
	return nil
}

func (m *BattleResultResponse) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *BattleResultResponse) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *BattleResultResponse) GetLevelUpExp() int32 {
	if m != nil {
		return m.LevelUpExp
	}
	return 0
}

func init() {
	proto.RegisterType((*BattleCreateRequest)(nil), "msg.BattleCreateRequest")
	proto.RegisterType((*BattleCreateResponse)(nil), "msg.BattleCreateResponse")
	proto.RegisterType((*BattleStartRequest)(nil), "msg.BattleStartRequest")
	proto.RegisterType((*BattleStartResponse)(nil), "msg.BattleStartResponse")
	proto.RegisterType((*BattleGuanKaRequest)(nil), "msg.BattleGuanKaRequest")
	proto.RegisterType((*BattleGuanKaResponse)(nil), "msg.BattleGuanKaResponse")
	proto.RegisterType((*BattleResultRequest)(nil), "msg.BattleResultRequest")
	proto.RegisterType((*BattleResultResponse)(nil), "msg.BattleResultResponse")
}

func init() { proto.RegisterFile("battle.proto", fileDescriptor_88d4c72b5869c382) }

var fileDescriptor_88d4c72b5869c382 = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcb, 0xae, 0xd3, 0x30,
	0x10, 0x55, 0x9a, 0x87, 0xc8, 0x14, 0x21, 0x30, 0x57, 0xc8, 0xaa, 0x78, 0x44, 0x41, 0x48, 0x59,
	0x55, 0x50, 0xd6, 0x6c, 0xb8, 0xaa, 0xe0, 0x0a, 0x56, 0xbe, 0xe2, 0x03, 0x7c, 0xe9, 0xa8, 0x54,
	0x37, 0xad, 0xd3, 0xb1, 0xc3, 0xe3, 0x4f, 0xf8, 0x1e, 0xf8, 0x31, 0xe4, 0xb1, 0x5d, 0x05, 0x56,
	0x2c, 0xba, 0x3b, 0x73, 0xce, 0x1c, 0xcf, 0x9c, 0x8c, 0x02, 0x77, 0x6f, 0xb4, 0x73, 0x3d, 0x2e,
	0x07, 0x32, 0xce, 0x88, 0x7c, 0x6f, 0xb7, 0x8b, 0x1a, 0x89, 0x42, 0xbd, 0xa8, 0xf7, 0x7a, 0x88,
	0x10, 0xbe, 0x20, 0x99, 0x88, 0xe7, 0xf6, 0x76, 0xd7, 0xf7, 0x49, 0xd8, 0x39, 0xdc, 0x07, 0xdc,
	0xbe, 0x81, 0x87, 0x6f, 0xf9, 0xbd, 0x4b, 0x42, 0xed, 0x50, 0xe1, 0x71, 0x44, 0xeb, 0x84, 0x80,
	0xc2, 0xfd, 0x18, 0x50, 0x66, 0x4d, 0xd6, 0x95, 0x8a, 0xb1, 0xe7, 0x34, 0x6d, 0xad, 0x9c, 0x35,
	0x79, 0x57, 0x2b, 0xc6, 0xed, 0x37, 0xb8, 0xf8, 0xdb, 0x6e, 0x07, 0x73, 0xb0, 0x28, 0x5e, 0x40,
	0xf1, 0xd9, 0x6c, 0x82, 0xff, 0xde, 0xea, 0xc1, 0x72, 0x6f, 0xb7, 0xcb, 0x24, 0x5e, 0x9a, 0x0d,
	0x2a, 0x96, 0xc5, 0x63, 0xc8, 0x91, 0x48, 0xce, 0x9a, 0xac, 0x9b, 0xaf, 0x80, 0xbb, 0xd6, 0x44,
	0x86, 0x94, 0xa7, 0xc5, 0x02, 0xee, 0x84, 0xac, 0x57, 0x1b, 0x99, 0x37, 0x59, 0x57, 0xab, 0x53,
	0xdd, 0xbe, 0x04, 0x11, 0x06, 0x5f, 0x3b, 0x4d, 0x2e, 0xad, 0x3d, 0x75, 0x64, 0xff, 0x38, 0x7e,
	0x67, 0x29, 0x6a, 0xb4, 0x9c, 0x73, 0xd5, 0x67, 0x50, 0xfa, 0xaf, 0x6d, 0x65, 0xde, 0xe4, 0xdd,
	0x7c, 0x55, 0xb3, 0xfe, 0x1e, 0xc9, 0xa8, 0xc0, 0x8b, 0x16, 0x2a, 0x3e, 0x81, 0x95, 0x05, 0x77,
	0x84, 0x17, 0xae, 0x3d, 0xa5, 0xa2, 0xe2, 0x7b, 0xf0, 0x38, 0xee, 0x06, 0x2b, 0xcb, 0x49, 0xcf,
	0xda, 0x53, 0x2a, 0x2a, 0xed, 0xab, 0x14, 0xe2, 0xdd, 0xa8, 0x0f, 0x1f, 0xf4, 0x24, 0xf8, 0x96,
	0x89, 0x18, 0xbc, 0x54, 0xa7, 0xba, 0xfd, 0x99, 0xa5, 0x23, 0x25, 0xcf, 0x39, 0x93, 0x3f, 0x82,
	0x8a, 0xd0, 0x8e, 0xbd, 0xe3, 0x13, 0x95, 0x2a, 0x56, 0xe2, 0x39, 0x54, 0x7e, 0x83, 0x5b, 0x2d,
	0x0b, 0x36, 0xce, 0xd9, 0x18, 0x37, 0x88, 0x52, 0x7b, 0x95, 0xd2, 0x28, 0x36, 0xfd, 0xc7, 0x19,
	0x27, 0xf3, 0x66, 0xd3, 0x79, 0xed, 0xaf, 0x53, 0xca, 0xf4, 0xd6, 0x39, 0x53, 0x3e, 0x81, 0x02,
	0x35, 0x1d, 0x38, 0x63, 0x3a, 0xef, 0x5a, 0xd3, 0x41, 0x31, 0x2d, 0xee, 0x43, 0xbe, 0xfe, 0x3e,
	0x70, 0xd2, 0x52, 0x79, 0x28, 0x2e, 0xa0, 0xfc, 0x88, 0x5f, 0xb1, 0x97, 0x25, 0x73, 0xa1, 0x10,
	0x4f, 0x01, 0x18, 0x7c, 0x1a, 0x7c, 0x7b, 0xc5, 0xd2, 0x84, 0xb9, 0xa9, 0xf8, 0xa7, 0x7c, 0xfd,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x10, 0x9b, 0xb1, 0x22, 0xe4, 0x03, 0x00, 0x00,
}
