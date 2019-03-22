// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hero.proto

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

type HeroRandomLevel int32

const (
	HeroRandomLevel_GOOD   HeroRandomLevel = 0
	HeroRandomLevel_BETTER HeroRandomLevel = 1
	HeroRandomLevel_BEST   HeroRandomLevel = 2
)

var HeroRandomLevel_name = map[int32]string{
	0: "GOOD",
	1: "BETTER",
	2: "BEST",
}

var HeroRandomLevel_value = map[string]int32{
	"GOOD":   0,
	"BETTER": 1,
	"BEST":   2,
}

func (x HeroRandomLevel) String() string {
	return proto.EnumName(HeroRandomLevel_name, int32(x))
}

func (HeroRandomLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{0}
}

type HeroRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeroRequest) Reset()         { *m = HeroRequest{} }
func (m *HeroRequest) String() string { return proto.CompactTextString(m) }
func (*HeroRequest) ProtoMessage()    {}
func (*HeroRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{0}
}

func (m *HeroRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroRequest.Unmarshal(m, b)
}
func (m *HeroRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroRequest.Marshal(b, m, deterministic)
}
func (m *HeroRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroRequest.Merge(m, src)
}
func (m *HeroRequest) XXX_Size() int {
	return xxx_messageInfo_HeroRequest.Size(m)
}
func (m *HeroRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeroRequest proto.InternalMessageInfo

type HeroResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Heros                []*Hero      `protobuf:"bytes,3,rep,name=heros,proto3" json:"heros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HeroResponse) Reset()         { *m = HeroResponse{} }
func (m *HeroResponse) String() string { return proto.CompactTextString(m) }
func (*HeroResponse) ProtoMessage()    {}
func (*HeroResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{1}
}

func (m *HeroResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroResponse.Unmarshal(m, b)
}
func (m *HeroResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroResponse.Marshal(b, m, deterministic)
}
func (m *HeroResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroResponse.Merge(m, src)
}
func (m *HeroResponse) XXX_Size() int {
	return xxx_messageInfo_HeroResponse.Size(m)
}
func (m *HeroResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeroResponse proto.InternalMessageInfo

func (m *HeroResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *HeroResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *HeroResponse) GetHeros() []*Hero {
	if m != nil {
		return m.Heros
	}
	return nil
}

type HeroRandomRequest struct {
	Level                HeroRandomLevel `protobuf:"varint,1,opt,name=Level,proto3,enum=msg.HeroRandomLevel" json:"Level,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *HeroRandomRequest) Reset()         { *m = HeroRandomRequest{} }
func (m *HeroRandomRequest) String() string { return proto.CompactTextString(m) }
func (*HeroRandomRequest) ProtoMessage()    {}
func (*HeroRandomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{2}
}

func (m *HeroRandomRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroRandomRequest.Unmarshal(m, b)
}
func (m *HeroRandomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroRandomRequest.Marshal(b, m, deterministic)
}
func (m *HeroRandomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroRandomRequest.Merge(m, src)
}
func (m *HeroRandomRequest) XXX_Size() int {
	return xxx_messageInfo_HeroRandomRequest.Size(m)
}
func (m *HeroRandomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroRandomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeroRandomRequest proto.InternalMessageInfo

func (m *HeroRandomRequest) GetLevel() HeroRandomLevel {
	if m != nil {
		return m.Level
	}
	return HeroRandomLevel_GOOD
}

type HeroRandomResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HeroRandomResponse) Reset()         { *m = HeroRandomResponse{} }
func (m *HeroRandomResponse) String() string { return proto.CompactTextString(m) }
func (*HeroRandomResponse) ProtoMessage()    {}
func (*HeroRandomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{3}
}

func (m *HeroRandomResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroRandomResponse.Unmarshal(m, b)
}
func (m *HeroRandomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroRandomResponse.Marshal(b, m, deterministic)
}
func (m *HeroRandomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroRandomResponse.Merge(m, src)
}
func (m *HeroRandomResponse) XXX_Size() int {
	return xxx_messageInfo_HeroRandomResponse.Size(m)
}
func (m *HeroRandomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroRandomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeroRandomResponse proto.InternalMessageInfo

func (m *HeroRandomResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *HeroRandomResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

type HeroOwnRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeroOwnRequest) Reset()         { *m = HeroOwnRequest{} }
func (m *HeroOwnRequest) String() string { return proto.CompactTextString(m) }
func (*HeroOwnRequest) ProtoMessage()    {}
func (*HeroOwnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{4}
}

func (m *HeroOwnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroOwnRequest.Unmarshal(m, b)
}
func (m *HeroOwnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroOwnRequest.Marshal(b, m, deterministic)
}
func (m *HeroOwnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroOwnRequest.Merge(m, src)
}
func (m *HeroOwnRequest) XXX_Size() int {
	return xxx_messageInfo_HeroOwnRequest.Size(m)
}
func (m *HeroOwnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroOwnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeroOwnRequest proto.InternalMessageInfo

type HeroOwnResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	Heros                []*Hero      `protobuf:"bytes,3,rep,name=heros,proto3" json:"heros,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HeroOwnResponse) Reset()         { *m = HeroOwnResponse{} }
func (m *HeroOwnResponse) String() string { return proto.CompactTextString(m) }
func (*HeroOwnResponse) ProtoMessage()    {}
func (*HeroOwnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{5}
}

func (m *HeroOwnResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroOwnResponse.Unmarshal(m, b)
}
func (m *HeroOwnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroOwnResponse.Marshal(b, m, deterministic)
}
func (m *HeroOwnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroOwnResponse.Merge(m, src)
}
func (m *HeroOwnResponse) XXX_Size() int {
	return xxx_messageInfo_HeroOwnResponse.Size(m)
}
func (m *HeroOwnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroOwnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeroOwnResponse proto.InternalMessageInfo

func (m *HeroOwnResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *HeroOwnResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *HeroOwnResponse) GetHeros() []*Hero {
	if m != nil {
		return m.Heros
	}
	return nil
}

type HeroSelectRequest struct {
	HeroId               string   `protobuf:"bytes,1,opt,name=heroId,proto3" json:"heroId,omitempty"`
	Pos                  int32    `protobuf:"varint,2,opt,name=pos,proto3" json:"pos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeroSelectRequest) Reset()         { *m = HeroSelectRequest{} }
func (m *HeroSelectRequest) String() string { return proto.CompactTextString(m) }
func (*HeroSelectRequest) ProtoMessage()    {}
func (*HeroSelectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{6}
}

func (m *HeroSelectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroSelectRequest.Unmarshal(m, b)
}
func (m *HeroSelectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroSelectRequest.Marshal(b, m, deterministic)
}
func (m *HeroSelectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroSelectRequest.Merge(m, src)
}
func (m *HeroSelectRequest) XXX_Size() int {
	return xxx_messageInfo_HeroSelectRequest.Size(m)
}
func (m *HeroSelectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroSelectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeroSelectRequest proto.InternalMessageInfo

func (m *HeroSelectRequest) GetHeroId() string {
	if m != nil {
		return m.HeroId
	}
	return ""
}

func (m *HeroSelectRequest) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

type HeroSelectResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	HeroIds              []string     `protobuf:"bytes,3,rep,name=heroIds,proto3" json:"heroIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HeroSelectResponse) Reset()         { *m = HeroSelectResponse{} }
func (m *HeroSelectResponse) String() string { return proto.CompactTextString(m) }
func (*HeroSelectResponse) ProtoMessage()    {}
func (*HeroSelectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{7}
}

func (m *HeroSelectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroSelectResponse.Unmarshal(m, b)
}
func (m *HeroSelectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroSelectResponse.Marshal(b, m, deterministic)
}
func (m *HeroSelectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroSelectResponse.Merge(m, src)
}
func (m *HeroSelectResponse) XXX_Size() int {
	return xxx_messageInfo_HeroSelectResponse.Size(m)
}
func (m *HeroSelectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroSelectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeroSelectResponse proto.InternalMessageInfo

func (m *HeroSelectResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *HeroSelectResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *HeroSelectResponse) GetHeroIds() []string {
	if m != nil {
		return m.HeroIds
	}
	return nil
}

type HeroUnSelectRequest struct {
	HeroId               string   `protobuf:"bytes,1,opt,name=heroId,proto3" json:"heroId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeroUnSelectRequest) Reset()         { *m = HeroUnSelectRequest{} }
func (m *HeroUnSelectRequest) String() string { return proto.CompactTextString(m) }
func (*HeroUnSelectRequest) ProtoMessage()    {}
func (*HeroUnSelectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{8}
}

func (m *HeroUnSelectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroUnSelectRequest.Unmarshal(m, b)
}
func (m *HeroUnSelectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroUnSelectRequest.Marshal(b, m, deterministic)
}
func (m *HeroUnSelectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroUnSelectRequest.Merge(m, src)
}
func (m *HeroUnSelectRequest) XXX_Size() int {
	return xxx_messageInfo_HeroUnSelectRequest.Size(m)
}
func (m *HeroUnSelectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroUnSelectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeroUnSelectRequest proto.InternalMessageInfo

func (m *HeroUnSelectRequest) GetHeroId() string {
	if m != nil {
		return m.HeroId
	}
	return ""
}

type HeroUnSelectResponse struct {
	Code                 ResponseCode `protobuf:"varint,1,opt,name=code,proto3,enum=msg.ResponseCode" json:"code,omitempty"`
	Err                  *Error       `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	HeroIds              []string     `protobuf:"bytes,3,rep,name=heroIds,proto3" json:"heroIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *HeroUnSelectResponse) Reset()         { *m = HeroUnSelectResponse{} }
func (m *HeroUnSelectResponse) String() string { return proto.CompactTextString(m) }
func (*HeroUnSelectResponse) ProtoMessage()    {}
func (*HeroUnSelectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{9}
}

func (m *HeroUnSelectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeroUnSelectResponse.Unmarshal(m, b)
}
func (m *HeroUnSelectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeroUnSelectResponse.Marshal(b, m, deterministic)
}
func (m *HeroUnSelectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeroUnSelectResponse.Merge(m, src)
}
func (m *HeroUnSelectResponse) XXX_Size() int {
	return xxx_messageInfo_HeroUnSelectResponse.Size(m)
}
func (m *HeroUnSelectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeroUnSelectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeroUnSelectResponse proto.InternalMessageInfo

func (m *HeroUnSelectResponse) GetCode() ResponseCode {
	if m != nil {
		return m.Code
	}
	return ResponseCode_FAIL
}

func (m *HeroUnSelectResponse) GetErr() *Error {
	if m != nil {
		return m.Err
	}
	return nil
}

func (m *HeroUnSelectResponse) GetHeroIds() []string {
	if m != nil {
		return m.HeroIds
	}
	return nil
}

type Hero struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Level                int32    `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	Strength             int32    `protobuf:"varint,4,opt,name=Strength,proto3" json:"Strength,omitempty"`
	Agility              int32    `protobuf:"varint,5,opt,name=Agility,proto3" json:"Agility,omitempty"`
	Intelligence         int32    `protobuf:"varint,6,opt,name=Intelligence,proto3" json:"Intelligence,omitempty"`
	Armor                int32    `protobuf:"varint,7,opt,name=Armor,proto3" json:"Armor,omitempty"`
	AttackMin            int32    `protobuf:"varint,8,opt,name=AttackMin,proto3" json:"AttackMin,omitempty"`
	AttackMax            int32    `protobuf:"varint,9,opt,name=AttackMax,proto3" json:"AttackMax,omitempty"`
	Blood                int32    `protobuf:"varint,10,opt,name=Blood,proto3" json:"Blood,omitempty"`
	Type                 int32    `protobuf:"varint,11,opt,name=Type,proto3" json:"Type,omitempty"`
	StrengthStep         int32    `protobuf:"varint,12,opt,name=StrengthStep,proto3" json:"StrengthStep,omitempty"`
	AgilityStep          int32    `protobuf:"varint,13,opt,name=AgilityStep,proto3" json:"AgilityStep,omitempty"`
	IntelligenceStep     int32    `protobuf:"varint,14,opt,name=IntelligenceStep,proto3" json:"IntelligenceStep,omitempty"`
	SkillIds             []string `protobuf:"bytes,15,rep,name=SkillIds,proto3" json:"SkillIds,omitempty"`
	HeroId               string   `protobuf:"bytes,16,opt,name=HeroId,proto3" json:"HeroId,omitempty"`
	PlayerId             string   `protobuf:"bytes,17,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	IsSelect             bool     `protobuf:"varint,18,opt,name=IsSelect,proto3" json:"IsSelect,omitempty"`
	Pos                  int32    `protobuf:"varint,19,opt,name=Pos,proto3" json:"Pos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hero) Reset()         { *m = Hero{} }
func (m *Hero) String() string { return proto.CompactTextString(m) }
func (*Hero) ProtoMessage()    {}
func (*Hero) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5c1be3f86b99613, []int{10}
}

func (m *Hero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hero.Unmarshal(m, b)
}
func (m *Hero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hero.Marshal(b, m, deterministic)
}
func (m *Hero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hero.Merge(m, src)
}
func (m *Hero) XXX_Size() int {
	return xxx_messageInfo_Hero.Size(m)
}
func (m *Hero) XXX_DiscardUnknown() {
	xxx_messageInfo_Hero.DiscardUnknown(m)
}

var xxx_messageInfo_Hero proto.InternalMessageInfo

func (m *Hero) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Hero) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hero) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Hero) GetStrength() int32 {
	if m != nil {
		return m.Strength
	}
	return 0
}

func (m *Hero) GetAgility() int32 {
	if m != nil {
		return m.Agility
	}
	return 0
}

func (m *Hero) GetIntelligence() int32 {
	if m != nil {
		return m.Intelligence
	}
	return 0
}

func (m *Hero) GetArmor() int32 {
	if m != nil {
		return m.Armor
	}
	return 0
}

func (m *Hero) GetAttackMin() int32 {
	if m != nil {
		return m.AttackMin
	}
	return 0
}

func (m *Hero) GetAttackMax() int32 {
	if m != nil {
		return m.AttackMax
	}
	return 0
}

func (m *Hero) GetBlood() int32 {
	if m != nil {
		return m.Blood
	}
	return 0
}

func (m *Hero) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Hero) GetStrengthStep() int32 {
	if m != nil {
		return m.StrengthStep
	}
	return 0
}

func (m *Hero) GetAgilityStep() int32 {
	if m != nil {
		return m.AgilityStep
	}
	return 0
}

func (m *Hero) GetIntelligenceStep() int32 {
	if m != nil {
		return m.IntelligenceStep
	}
	return 0
}

func (m *Hero) GetSkillIds() []string {
	if m != nil {
		return m.SkillIds
	}
	return nil
}

func (m *Hero) GetHeroId() string {
	if m != nil {
		return m.HeroId
	}
	return ""
}

func (m *Hero) GetPlayerId() string {
	if m != nil {
		return m.PlayerId
	}
	return ""
}

func (m *Hero) GetIsSelect() bool {
	if m != nil {
		return m.IsSelect
	}
	return false
}

func (m *Hero) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func init() {
	proto.RegisterEnum("msg.HeroRandomLevel", HeroRandomLevel_name, HeroRandomLevel_value)
	proto.RegisterType((*HeroRequest)(nil), "msg.HeroRequest")
	proto.RegisterType((*HeroResponse)(nil), "msg.HeroResponse")
	proto.RegisterType((*HeroRandomRequest)(nil), "msg.HeroRandomRequest")
	proto.RegisterType((*HeroRandomResponse)(nil), "msg.HeroRandomResponse")
	proto.RegisterType((*HeroOwnRequest)(nil), "msg.HeroOwnRequest")
	proto.RegisterType((*HeroOwnResponse)(nil), "msg.HeroOwnResponse")
	proto.RegisterType((*HeroSelectRequest)(nil), "msg.HeroSelectRequest")
	proto.RegisterType((*HeroSelectResponse)(nil), "msg.HeroSelectResponse")
	proto.RegisterType((*HeroUnSelectRequest)(nil), "msg.HeroUnSelectRequest")
	proto.RegisterType((*HeroUnSelectResponse)(nil), "msg.HeroUnSelectResponse")
	proto.RegisterType((*Hero)(nil), "msg.Hero")
}

func init() { proto.RegisterFile("hero.proto", fileDescriptor_d5c1be3f86b99613) }

var fileDescriptor_d5c1be3f86b99613 = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0xd1, 0x6e, 0xda, 0x30,
	0x14, 0x1d, 0x24, 0x50, 0x72, 0x69, 0x69, 0x70, 0xd1, 0x64, 0x55, 0x95, 0x86, 0x22, 0x4d, 0x42,
	0x48, 0x43, 0x1a, 0x7b, 0x9e, 0x26, 0xd8, 0xd0, 0x8a, 0xb4, 0x0d, 0x64, 0xd8, 0xc3, 0x1e, 0x33,
	0x72, 0x45, 0x51, 0x43, 0xcc, 0x6c, 0x77, 0x2d, 0x3f, 0xb0, 0xef, 0x9e, 0x7c, 0x4d, 0xd2, 0xa0,
	0xbd, 0xec, 0xa1, 0xea, 0x9b, 0xcf, 0x39, 0xd7, 0x3e, 0xc7, 0x37, 0xd7, 0x01, 0xb8, 0x41, 0x25,
	0x07, 0x3b, 0x25, 0x8d, 0x64, 0xde, 0x56, 0xaf, 0x2f, 0x03, 0x54, 0xca, 0xe1, 0xe8, 0x0c, 0x9a,
	0xd7, 0xa8, 0xa4, 0xc0, 0x5f, 0x77, 0xa8, 0x4d, 0x64, 0xe0, 0xd4, 0x41, 0xbd, 0x93, 0x99, 0x46,
	0xf6, 0x1a, 0xfc, 0x95, 0x4c, 0x90, 0x57, 0xba, 0x95, 0x5e, 0x6b, 0xd8, 0x1e, 0x6c, 0xf5, 0x7a,
	0x90, 0x8b, 0x1f, 0x65, 0x82, 0x82, 0x64, 0x76, 0x05, 0x1e, 0x2a, 0xc5, 0xab, 0xdd, 0x4a, 0xaf,
	0x39, 0x04, 0xaa, 0x9a, 0x28, 0x25, 0x95, 0xb0, 0x34, 0x7b, 0x05, 0x35, 0x9b, 0x40, 0x73, 0xaf,
	0xeb, 0xf5, 0x9a, 0xc3, 0x80, 0x74, 0xb2, 0x71, 0x7c, 0xf4, 0x01, 0xda, 0x04, 0xe3, 0x2c, 0x91,
	0xdb, 0x43, 0x14, 0xd6, 0x87, 0xda, 0x17, 0xfc, 0x8d, 0xe9, 0xc1, 0xbb, 0xf3, 0xb8, 0x8b, 0xca,
	0x48, 0x13, 0xae, 0x24, 0xfa, 0x01, 0xac, 0x7c, 0xc0, 0x13, 0x86, 0x8f, 0x42, 0x68, 0xd9, 0xa3,
	0x67, 0xf7, 0x59, 0xde, 0xa3, 0x7b, 0x38, 0x2f, 0x98, 0x67, 0x6d, 0xd3, 0x7b, 0xd7, 0xa6, 0x05,
	0xa6, 0xb8, 0x32, 0x79, 0x9b, 0x5e, 0x42, 0xdd, 0xaa, 0xd3, 0x84, 0xcc, 0x03, 0x71, 0x40, 0x2c,
	0x04, 0x6f, 0x27, 0x35, 0x79, 0xd5, 0x84, 0x5d, 0x46, 0xda, 0x35, 0x29, 0xdf, 0xfe, 0x94, 0xd1,
	0x39, 0x9c, 0x38, 0x5b, 0x17, 0x3e, 0x10, 0x39, 0x8c, 0xde, 0xc0, 0x85, 0x35, 0xfd, 0x9e, 0xfd,
	0x57, 0xea, 0xe8, 0x0e, 0x3a, 0xc7, 0xe5, 0xcf, 0x93, 0xf2, 0x8f, 0x0f, 0xbe, 0xf5, 0x65, 0x2d,
	0xa8, 0x1e, 0x32, 0xd5, 0x44, 0x75, 0x9a, 0x30, 0x06, 0xfe, 0xb7, 0x78, 0x8b, 0x74, 0x62, 0x20,
	0x68, 0xcd, 0x3a, 0xf9, 0x60, 0x7a, 0x54, 0xe6, 0x00, 0xbb, 0x84, 0xc6, 0xc2, 0x28, 0xcc, 0xd6,
	0xe6, 0x86, 0xfb, 0x24, 0x14, 0xd8, 0x1a, 0x8f, 0xd6, 0x9b, 0x74, 0x63, 0xf6, 0xbc, 0x46, 0x52,
	0x0e, 0x59, 0x04, 0xa7, 0xd3, 0xcc, 0x60, 0x9a, 0x6e, 0xd6, 0x98, 0xad, 0x90, 0xd7, 0x49, 0x3e,
	0xe2, 0xac, 0xdf, 0x48, 0x6d, 0xa5, 0xe2, 0x27, 0xce, 0x8f, 0x00, 0xbb, 0x82, 0x60, 0x64, 0x4c,
	0xbc, 0xba, 0xfd, 0xba, 0xc9, 0x78, 0x83, 0x94, 0x47, 0xa2, 0xa4, 0xc6, 0x0f, 0x3c, 0x38, 0x52,
	0xe3, 0x07, 0x7b, 0xe2, 0x38, 0x95, 0x32, 0xe1, 0xe0, 0x4e, 0x24, 0x60, 0xef, 0xba, 0xdc, 0xef,
	0x90, 0x37, 0x89, 0xa4, 0xb5, 0xcd, 0x97, 0xdf, 0x62, 0x61, 0x70, 0xc7, 0x4f, 0x5d, 0xbe, 0x32,
	0xc7, 0xba, 0xd0, 0x3c, 0x5c, 0x87, 0x4a, 0xce, 0xa8, 0xa4, 0x4c, 0xb1, 0x3e, 0x84, 0xe5, 0x1b,
	0x51, 0x59, 0x8b, 0xca, 0xfe, 0xe1, 0xa9, 0x8f, 0xb7, 0x9b, 0x34, 0xb5, 0x5f, 0xe9, 0x9c, 0xbe,
	0x52, 0x81, 0xed, 0xd4, 0x5c, 0xbb, 0xa9, 0x09, 0xdd, 0xd4, 0x38, 0x64, 0xf7, 0xcc, 0xd3, 0x78,
	0x8f, 0x6a, 0x9a, 0xf0, 0x36, 0x29, 0x05, 0xb6, 0xda, 0x54, 0xbb, 0x69, 0xe2, 0xac, 0x5b, 0xe9,
	0x35, 0x44, 0x81, 0xed, 0x1b, 0x99, 0x4b, 0xcd, 0x2f, 0xdc, 0x1b, 0x99, 0x4b, 0xdd, 0x7f, 0xeb,
	0xde, 0x76, 0xe9, 0x17, 0xc3, 0x1a, 0xe0, 0x7f, 0x9e, 0xcd, 0x3e, 0x85, 0x2f, 0x18, 0x40, 0x7d,
	0x3c, 0x59, 0x2e, 0x27, 0x22, 0xac, 0x58, 0x76, 0x3c, 0x59, 0x2c, 0xc3, 0xea, 0xcf, 0x3a, 0xfd,
	0x48, 0xdf, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x84, 0x9a, 0xe0, 0x7f, 0x66, 0x05, 0x00, 0x00,
}
