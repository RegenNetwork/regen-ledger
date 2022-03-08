// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/v1/types.proto

package v1

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
	// allowed_class_creators is an allowlist defining the addresses with
	// the required permissions to create credit classes
	AllowedClassCreators []string `protobuf:"bytes,2,rep,name=allowed_class_creators,json=allowedClassCreators,proto3" json:"allowed_class_creators,omitempty"`
	// allowlist_enabled is a param that enables/disables the allowlist for credit
	// creation
	AllowlistEnabled bool `protobuf:"varint,3,opt,name=allowlist_enabled,json=allowlistEnabled,proto3" json:"allowlist_enabled,omitempty"`
	// credit_types is a list of definitions for credit types
	CreditTypes []*CreditType `protobuf:"bytes,4,rep,name=credit_types,json=creditTypes,proto3" json:"credit_types,omitempty"`
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

func (m *Params) GetCreditTypes() []*CreditType {
	if m != nil {
		return m.CreditTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "regen.ecocredit.v1.Params")
}

func init() { proto.RegisterFile("regen/ecocredit/v1/types.proto", fileDescriptor_7b044b6b740b984f) }

var fileDescriptor_7b044b6b740b984f = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4e, 0xeb, 0x30,
	0x10, 0x86, 0x93, 0xf6, 0xa9, 0x7a, 0x2f, 0x7d, 0x42, 0x25, 0xaa, 0x50, 0xe8, 0xc2, 0xad, 0x58,
	0x45, 0x42, 0xb5, 0x09, 0xf4, 0x02, 0x34, 0x82, 0x25, 0x42, 0x11, 0x2b, 0x36, 0x95, 0xe3, 0x0c,
	0x21, 0x6a, 0x1a, 0x57, 0xb6, 0xdb, 0xd2, 0x5b, 0x70, 0x09, 0x36, 0x9c, 0xa4, 0xcb, 0x2e, 0x59,
	0x01, 0x6a, 0x2f, 0x82, 0xe2, 0x98, 0x0a, 0x04, 0x2b, 0x8f, 0xe7, 0x9b, 0xf1, 0x3f, 0xff, 0xd8,
	0x41, 0x02, 0x52, 0x28, 0x08, 0x30, 0xce, 0x04, 0x24, 0x99, 0x22, 0xf3, 0x80, 0xa8, 0xe5, 0x14,
	0x24, 0x9e, 0x0a, 0xae, 0xb8, 0xeb, 0x6a, 0x8e, 0x77, 0x1c, 0xcf, 0x83, 0x4e, 0x3b, 0xe5, 0x29,
	0xd7, 0x98, 0x94, 0x51, 0x55, 0xd9, 0x41, 0x8c, 0xcb, 0x09, 0x97, 0x24, 0xa6, 0x12, 0xc8, 0x3c,
	0x88, 0x41, 0xd1, 0x80, 0x30, 0x9e, 0x15, 0x9f, 0xfc, 0x17, 0x25, 0xa9, 0xa8, 0x82, 0x8a, 0x1f,
	0x3d, 0xd5, 0x9c, 0xc6, 0x35, 0x15, 0x74, 0x22, 0xdd, 0x99, 0xd3, 0xaa, 0x6a, 0x46, 0x2c, 0xa7,
	0x52, 0x8e, 0xee, 0x00, 0x3c, 0xbb, 0x57, 0xf7, 0x9b, 0xa7, 0x87, 0xb8, 0x52, 0xc1, 0xa5, 0x0a,
	0x36, 0x2a, 0x38, 0xe4, 0x59, 0x31, 0x3c, 0x59, 0xbd, 0x76, 0xad, 0xe7, 0xb7, 0xae, 0x9f, 0x66,
	0xea, 0x7e, 0x16, 0x63, 0xc6, 0x27, 0xc4, 0x8c, 0x54, 0x1d, 0x7d, 0x99, 0x8c, 0x8d, 0xb7, 0xb2,
	0x41, 0x46, 0x7b, 0x95, 0x48, 0x58, 0x6a, 0x5c, 0x02, 0xb8, 0x03, 0xe7, 0x80, 0xe6, 0x39, 0x5f,
	0x40, 0x62, 0x74, 0x99, 0x00, 0xaa, 0xb8, 0x90, 0x5e, 0xad, 0x57, 0xf7, 0xff, 0x45, 0x6d, 0x43,
	0x75, 0x43, 0x68, 0x98, 0x7b, 0xec, 0xec, 0xeb, 0x7c, 0x9e, 0x49, 0x35, 0x82, 0x82, 0xc6, 0x39,
	0x24, 0x5e, 0xbd, 0x67, 0xfb, 0x7f, 0xa3, 0xd6, 0x0e, 0x5c, 0x54, 0x79, 0xf7, 0xdc, 0xf9, 0x6f,
	0x9c, 0xe9, 0x41, 0xbc, 0x3f, 0xda, 0x15, 0xc2, 0x3f, 0xb7, 0x8c, 0x43, 0x1d, 0xdd, 0x2c, 0xa7,
	0x10, 0x35, 0xd9, 0x2e, 0x96, 0xc3, 0xab, 0xd5, 0x06, 0xd9, 0xeb, 0x0d, 0xb2, 0xdf, 0x37, 0xc8,
	0x7e, 0xdc, 0x22, 0x6b, 0xbd, 0x45, 0xd6, 0xcb, 0x16, 0x59, 0xb7, 0x83, 0x2f, 0xce, 0xf5, 0x83,
	0xfd, 0x02, 0xd4, 0x82, 0x8b, 0xb1, 0xb9, 0xe5, 0x90, 0xa4, 0x20, 0xc8, 0xc3, 0xb7, 0x3f, 0x88,
	0x1b, 0x7a, 0xfd, 0x67, 0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xe6, 0x9d, 0xd8, 0x0a, 0x02,
	0x00, 0x00,
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
	if len(m.CreditTypes) > 0 {
		for iNdEx := len(m.CreditTypes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CreditTypes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.AllowlistEnabled {
		i--
		if m.AllowlistEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.AllowedClassCreators) > 0 {
		for iNdEx := len(m.AllowedClassCreators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowedClassCreators[iNdEx])
			copy(dAtA[i:], m.AllowedClassCreators[iNdEx])
			i = encodeVarintTypes(dAtA, i, uint64(len(m.AllowedClassCreators[iNdEx])))
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
	if len(m.AllowedClassCreators) > 0 {
		for _, s := range m.AllowedClassCreators {
			l = len(s)
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.AllowlistEnabled {
		n += 2
	}
	if len(m.CreditTypes) > 0 {
		for _, e := range m.CreditTypes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
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
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditTypes", wireType)
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
			m.CreditTypes = append(m.CreditTypes, &CreditType{})
			if err := m.CreditTypes[len(m.CreditTypes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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