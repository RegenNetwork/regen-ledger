// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/v1/types.proto

package core

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Params defines the updatable global parameters of the ecocredit module for
// use with the x/params module.
type Params struct {
	// credit_class_fee is the fixed fee charged on creation of a new credit class
	CreditClassFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=credit_class_fee,json=creditClassFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"credit_class_fee"`
	// basket_fee is the fixed fee charged on creation of a new basket
	BasketFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=basket_fee,json=basketFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"basket_fee"`
	// allowed_class_creators is an allowlist defining the addresses with
	// the required permissions to create credit classes
	AllowedClassCreators []string `protobuf:"bytes,3,rep,name=allowed_class_creators,json=allowedClassCreators,proto3" json:"allowed_class_creators,omitempty"`
	// allowlist_enabled is a param that enables/disables the allowlist for credit
	// creation
	AllowlistEnabled bool `protobuf:"varint,4,opt,name=allowlist_enabled,json=allowlistEnabled,proto3" json:"allowlist_enabled,omitempty"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b044b6b740b984f, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetCreditClassFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.CreditClassFee
	}
	return nil
}

func (m *Params) GetBasketFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.BasketFee
	}
	return nil
}

func (m *Params) GetAllowedClassCreators() []string {
	if m != nil {
		return m.AllowedClassCreators
	}
	return nil
}

func (m *Params) GetAllowlistEnabled() bool {
	if m != nil {
		return m.AllowlistEnabled
	}
	return false
}

// OriginTx is the transaction from another chain or registry that triggered
// the minting of credits.
type OriginTx struct {
	// id is the transaction ID of an originating transaction or operation based
	// on a type (i.e. transaction ID, serial number).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// source is the source chain or registry of the transaction originating the
	// mint process (e.g. polygon, ethereum, verra).
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (m *OriginTx) Reset()         { *m = OriginTx{} }
func (m *OriginTx) String() string { return proto.CompactTextString(m) }
func (*OriginTx) ProtoMessage()    {}
func (*OriginTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b044b6b740b984f, []int{1}
}
func (m *OriginTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *OriginTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_OriginTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *OriginTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OriginTx.Merge(m, src)
}
func (m *OriginTx) XXX_Size() int {
	return m.Size()
}
func (m *OriginTx) XXX_DiscardUnknown() {
	xxx_messageInfo_OriginTx.DiscardUnknown(m)
}

var xxx_messageInfo_OriginTx proto.InternalMessageInfo

func (m *OriginTx) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *OriginTx) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

// CreditTypeProposal is a gov Content type for adding a credit type.
type CreditTypeProposal struct {
	// title is the title of the proposal.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// description is the description of the proposal.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// credit_type is the credit type to be added to the network if the proposal
	// passes.
	CreditType *CreditType `protobuf:"bytes,3,opt,name=credit_type,json=creditType,proto3" json:"credit_type,omitempty"`
}

func (m *CreditTypeProposal) Reset()      { *m = CreditTypeProposal{} }
func (*CreditTypeProposal) ProtoMessage() {}
func (*CreditTypeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b044b6b740b984f, []int{2}
}
func (m *CreditTypeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreditTypeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreditTypeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CreditTypeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreditTypeProposal.Merge(m, src)
}
func (m *CreditTypeProposal) XXX_Size() int {
	return m.Size()
}
func (m *CreditTypeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_CreditTypeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_CreditTypeProposal proto.InternalMessageInfo

func (m *CreditTypeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *CreditTypeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreditTypeProposal) GetCreditType() *CreditType {
	if m != nil {
		return m.CreditType
	}
	return nil
}

// Credits represents a simple structure for credits.
type Credits struct {
	// batch_denom is the denom of the credit batch
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// amount is the amount of credits
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Credits) Reset()         { *m = Credits{} }
func (m *Credits) String() string { return proto.CompactTextString(m) }
func (*Credits) ProtoMessage()    {}
func (*Credits) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b044b6b740b984f, []int{3}
}
func (m *Credits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Credits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Credits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Credits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Credits.Merge(m, src)
}
func (m *Credits) XXX_Size() int {
	return m.Size()
}
func (m *Credits) XXX_DiscardUnknown() {
	xxx_messageInfo_Credits.DiscardUnknown(m)
}

var xxx_messageInfo_Credits proto.InternalMessageInfo

func (m *Credits) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *Credits) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*Params)(nil), "regen.ecocredit.v1.Params")
	proto.RegisterType((*OriginTx)(nil), "regen.ecocredit.v1.OriginTx")
	proto.RegisterType((*CreditTypeProposal)(nil), "regen.ecocredit.v1.CreditTypeProposal")
	proto.RegisterType((*Credits)(nil), "regen.ecocredit.v1.Credits")
}

func init() { proto.RegisterFile("regen/ecocredit/v1/types.proto", fileDescriptor_7b044b6b740b984f) }

var fileDescriptor_7b044b6b740b984f = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0x73, 0x49, 0x09, 0x8d, 0x23, 0x55, 0xc5, 0x8a, 0xaa, 0xd0, 0xe1, 0x72, 0xca, 0x74,
	0x12, 0xaa, 0x8f, 0x04, 0xc4, 0xc0, 0x82, 0x94, 0x00, 0x2b, 0x51, 0xd4, 0x89, 0x25, 0xf2, 0xf9,
	0x1e, 0x57, 0x93, 0xcb, 0xbd, 0xc8, 0x76, 0xd2, 0xf6, 0x5b, 0x20, 0xb1, 0x30, 0x32, 0xf3, 0x49,
	0x3a, 0xa1, 0x8e, 0x4c, 0x80, 0x92, 0x2f, 0x82, 0x7c, 0x76, 0x43, 0x24, 0x18, 0x3b, 0x9d, 0xdf,
	0xff, 0xef, 0xf3, 0xef, 0xfd, 0xed, 0x47, 0x42, 0x05, 0x39, 0x94, 0x09, 0x08, 0x14, 0x0a, 0x32,
	0x69, 0x92, 0xf5, 0x20, 0x31, 0xd7, 0x4b, 0xd0, 0x6c, 0xa9, 0xd0, 0x20, 0xa5, 0x95, 0xcf, 0x76,
	0x3e, 0x5b, 0x0f, 0x4e, 0x3b, 0x39, 0xe6, 0x58, 0xd9, 0x89, 0x5d, 0xb9, 0x9d, 0xa7, 0xa1, 0x40,
	0xbd, 0x40, 0x9d, 0xa4, 0x5c, 0x43, 0xb2, 0x1e, 0xa4, 0x60, 0xf8, 0x20, 0x11, 0x28, 0xcb, 0x3b,
	0xff, 0x3f, 0x24, 0x6d, 0xb8, 0x01, 0xe7, 0xf7, 0xbf, 0xd7, 0x49, 0x73, 0xc2, 0x15, 0x5f, 0x68,
	0xba, 0x22, 0xc7, 0x6e, 0xcf, 0x4c, 0x14, 0x5c, 0xeb, 0xd9, 0x07, 0x80, 0x6e, 0x10, 0x35, 0xe2,
	0xf6, 0xf0, 0x31, 0x73, 0x14, 0x66, 0x29, 0xcc, 0x53, 0xd8, 0x18, 0x65, 0x39, 0x7a, 0x7a, 0xf3,
	0xb3, 0x57, 0xfb, 0xf6, 0xab, 0x17, 0xe7, 0xd2, 0x5c, 0xac, 0x52, 0x26, 0x70, 0x91, 0xf8, 0x96,
	0xdc, 0xe7, 0x4c, 0x67, 0x73, 0x9f, 0xcd, 0xfe, 0xa0, 0xa7, 0x47, 0x0e, 0x32, 0xb6, 0x8c, 0xb7,
	0x00, 0xf4, 0x23, 0x21, 0x29, 0xd7, 0x73, 0x30, 0x15, 0xb0, 0x7e, 0xff, 0xc0, 0x96, 0x3b, 0xde,
	0xb2, 0x9e, 0x93, 0x13, 0x5e, 0x14, 0x78, 0x09, 0x99, 0xcf, 0x28, 0x14, 0x70, 0x83, 0x4a, 0x77,
	0x1b, 0x51, 0x23, 0x6e, 0x4d, 0x3b, 0xde, 0xad, 0x9a, 0x1b, 0x7b, 0x8f, 0x3e, 0x21, 0x8f, 0x2a,
	0xbd, 0x90, 0xda, 0xcc, 0xa0, 0xe4, 0x69, 0x01, 0x59, 0xf7, 0x20, 0x0a, 0xe2, 0xc3, 0xe9, 0xf1,
	0xce, 0x78, 0xe3, 0xf4, 0xfe, 0x90, 0x1c, 0xbe, 0x53, 0x32, 0x97, 0xe5, 0xf9, 0x15, 0x3d, 0x22,
	0x75, 0x99, 0x75, 0x83, 0x28, 0x88, 0x5b, 0xd3, 0xba, 0xcc, 0xe8, 0x09, 0x69, 0x6a, 0x5c, 0x29,
	0x61, 0x63, 0x5a, 0xcd, 0x57, 0xfd, 0xcf, 0x01, 0xa1, 0xe3, 0xea, 0x56, 0xce, 0xaf, 0x97, 0x30,
	0x51, 0xb8, 0x44, 0xcd, 0x0b, 0xda, 0x21, 0x0f, 0x8c, 0x34, 0x05, 0xf8, 0x13, 0x5c, 0x41, 0x23,
	0xd2, 0xce, 0x40, 0x0b, 0x25, 0x97, 0x46, 0x62, 0xe9, 0x4f, 0xda, 0x97, 0xe8, 0x2b, 0xd2, 0xf6,
	0x0f, 0x69, 0xaf, 0xa1, 0xdb, 0x88, 0x82, 0xb8, 0x3d, 0x0c, 0xd9, 0xbf, 0x33, 0xc5, 0xfe, 0x42,
	0xa7, 0x44, 0xec, 0xd6, 0x2f, 0x0f, 0xbe, 0x7c, 0xed, 0xd5, 0xfa, 0x23, 0xf2, 0xd0, 0xf9, 0x9a,
	0xf6, 0x48, 0x3b, 0xe5, 0x46, 0x5c, 0xcc, 0x32, 0x28, 0x71, 0xe1, 0xfb, 0x21, 0x95, 0xf4, 0xda,
	0x2a, 0x36, 0x19, 0x5f, 0xe0, 0xaa, 0x34, 0x77, 0xc9, 0x5c, 0x35, 0x9a, 0xdc, 0x6c, 0xc2, 0xe0,
	0x76, 0x13, 0x06, 0xbf, 0x37, 0x61, 0xf0, 0x69, 0x1b, 0xd6, 0x6e, 0xb7, 0x61, 0xed, 0xc7, 0x36,
	0xac, 0xbd, 0x7f, 0xb1, 0xf7, 0x7e, 0x55, 0x67, 0x67, 0x25, 0x98, 0x4b, 0x54, 0x73, 0x5f, 0x15,
	0x90, 0xe5, 0xa0, 0x92, 0xab, 0xbd, 0xd1, 0x15, 0xa8, 0x20, 0x6d, 0x56, 0x73, 0xfb, 0xec, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0x8c, 0xd1, 0xfe, 0x43, 0x03, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AllowlistEnabled {
		i--
		if m.AllowlistEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.AllowedClassCreators) > 0 {
		for iNdEx := len(m.AllowedClassCreators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedClassCreators[iNdEx])
			copy(dAtA[i:], m.AllowedClassCreators[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.AllowedClassCreators[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.BasketFee) > 0 {
		for iNdEx := len(m.BasketFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BasketFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.CreditClassFee) > 0 {
		for iNdEx := len(m.CreditClassFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreditClassFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *OriginTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *OriginTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *OriginTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Source) > 0 {
		i -= len(m.Source)
		copy(dAtA[i:], m.Source)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Source)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CreditTypeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreditTypeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CreditTypeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CreditType != nil {
		{
			size, err := m.CreditType.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Credits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Credits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Credits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CreditClassFee) > 0 {
		for _, e := range m.CreditClassFee {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.BasketFee) > 0 {
		for _, e := range m.BasketFee {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.AllowedClassCreators) > 0 {
		for _, s := range m.AllowedClassCreators {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.AllowlistEnabled {
		n += 2
	}
	return n
}

func (m *OriginTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Source)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *CreditTypeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.CreditType != nil {
		l = m.CreditType.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Credits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditClassFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreditClassFee = append(m.CreditClassFee, types.Coin{})
			if err := m.CreditClassFee[len(m.CreditClassFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BasketFee = append(m.BasketFee, types.Coin{})
			if err := m.BasketFee[len(m.BasketFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowedClassCreators", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowedClassCreators = append(m.AllowedClassCreators, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowlistEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AllowlistEnabled = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *OriginTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: OriginTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: OriginTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Source = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreditTypeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreditTypeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreditTypeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditType", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CreditType == nil {
				m.CreditType = &CreditType{}
			}
			if err := m.CreditType.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Credits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Credits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Credits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
