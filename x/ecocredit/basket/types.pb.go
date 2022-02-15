// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/basket/v1/types.proto

package basket

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// BasketCredit represents the information for a credit batch inside a basket.
type BasketCredit struct {
	// batch_denom is the unique ID of the credit batch.
	BatchDenom string `protobuf:"bytes,1,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// amount is the number of credits being put into or taken out of the basket.
	// Decimal values are acceptable within the precision of the corresponding
	//  credit type for this batch.
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *BasketCredit) Reset()         { *m = BasketCredit{} }
func (m *BasketCredit) String() string { return proto.CompactTextString(m) }
func (*BasketCredit) ProtoMessage()    {}
func (*BasketCredit) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c256e957c69c4d, []int{0}
}
func (m *BasketCredit) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasketCredit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasketCredit.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasketCredit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasketCredit.Merge(m, src)
}
func (m *BasketCredit) XXX_Size() int {
	return m.Size()
}
func (m *BasketCredit) XXX_DiscardUnknown() {
	xxx_messageInfo_BasketCredit.DiscardUnknown(m)
}

var xxx_messageInfo_BasketCredit proto.InternalMessageInfo

func (m *BasketCredit) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *BasketCredit) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

// DateCriteria represents a basket credit batch date criteria.
type DateCriteria struct {
	// sum is the oneof representing the date criteria.
	//
	// Types that are valid to be assigned to Sum:
	//	*DateCriteria_MinStartDate
	//	*DateCriteria_StartDateWindow
	Sum isDateCriteria_Sum `protobuf_oneof:"sum"`
}

func (m *DateCriteria) Reset()         { *m = DateCriteria{} }
func (m *DateCriteria) String() string { return proto.CompactTextString(m) }
func (*DateCriteria) ProtoMessage()    {}
func (*DateCriteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_e6c256e957c69c4d, []int{1}
}
func (m *DateCriteria) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DateCriteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DateCriteria.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DateCriteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateCriteria.Merge(m, src)
}
func (m *DateCriteria) XXX_Size() int {
	return m.Size()
}
func (m *DateCriteria) XXX_DiscardUnknown() {
	xxx_messageInfo_DateCriteria.DiscardUnknown(m)
}

var xxx_messageInfo_DateCriteria proto.InternalMessageInfo

type isDateCriteria_Sum interface {
	isDateCriteria_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type DateCriteria_MinStartDate struct {
	MinStartDate *types.Timestamp `protobuf:"bytes,1,opt,name=min_start_date,json=minStartDate,proto3,oneof" json:"min_start_date,omitempty"`
}
type DateCriteria_StartDateWindow struct {
	StartDateWindow *types.Duration `protobuf:"bytes,2,opt,name=start_date_window,json=startDateWindow,proto3,oneof" json:"start_date_window,omitempty"`
}

func (*DateCriteria_MinStartDate) isDateCriteria_Sum()    {}
func (*DateCriteria_StartDateWindow) isDateCriteria_Sum() {}

func (m *DateCriteria) GetSum() isDateCriteria_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *DateCriteria) GetMinStartDate() *types.Timestamp {
	if x, ok := m.GetSum().(*DateCriteria_MinStartDate); ok {
		return x.MinStartDate
	}
	return nil
}

func (m *DateCriteria) GetStartDateWindow() *types.Duration {
	if x, ok := m.GetSum().(*DateCriteria_StartDateWindow); ok {
		return x.StartDateWindow
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DateCriteria) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DateCriteria_MinStartDate)(nil),
		(*DateCriteria_StartDateWindow)(nil),
	}
}

func init() {
	proto.RegisterType((*BasketCredit)(nil), "regen.ecocredit.basket.v1.BasketCredit")
	proto.RegisterType((*DateCriteria)(nil), "regen.ecocredit.basket.v1.DateCriteria")
}

func init() {
	proto.RegisterFile("regen/ecocredit/basket/v1/types.proto", fileDescriptor_e6c256e957c69c4d)
}

var fileDescriptor_e6c256e957c69c4d = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0x41, 0x4f, 0xc2, 0x30,
	0x18, 0x86, 0x37, 0x8d, 0x24, 0x16, 0xa2, 0x71, 0x07, 0x03, 0x1c, 0x8a, 0x21, 0x31, 0xf1, 0x62,
	0x1b, 0xf0, 0xe2, 0x79, 0x90, 0xc0, 0x79, 0x9a, 0x98, 0x78, 0x59, 0xba, 0xed, 0x73, 0x34, 0xd0,
	0x96, 0x74, 0xdf, 0x40, 0xff, 0x85, 0xbf, 0xc1, 0x5f, 0xe3, 0x91, 0xa3, 0x47, 0x03, 0x7f, 0xc4,
	0xd0, 0x81, 0x26, 0x72, 0xfc, 0xbe, 0xf7, 0x7d, 0x9f, 0xbe, 0x6d, 0xc9, 0xb5, 0x85, 0x1c, 0x34,
	0x87, 0xd4, 0xa4, 0x16, 0x32, 0x89, 0x3c, 0x11, 0xc5, 0x14, 0x90, 0x2f, 0x7a, 0x1c, 0xdf, 0xe6,
	0x50, 0xb0, 0xb9, 0x35, 0x68, 0x82, 0x96, 0xb3, 0xb1, 0x5f, 0x1b, 0xab, 0x6c, 0x6c, 0xd1, 0x6b,
	0x77, 0x72, 0x63, 0xf2, 0x19, 0x70, 0x67, 0x4c, 0xca, 0x17, 0x8e, 0x52, 0x41, 0x81, 0x42, 0xcd,
	0xab, 0x6c, 0x9b, 0xfe, 0x37, 0x64, 0xa5, 0x15, 0x28, 0x8d, 0xae, 0xf4, 0xee, 0x88, 0x34, 0x42,
	0x47, 0x1b, 0x38, 0x74, 0xd0, 0x21, 0xf5, 0x44, 0x60, 0x3a, 0x89, 0x33, 0xd0, 0x46, 0x35, 0xfd,
	0x2b, 0xff, 0xe6, 0x34, 0x22, 0x6e, 0x35, 0xdc, 0x6e, 0x82, 0x4b, 0x52, 0x13, 0xca, 0x94, 0x1a,
	0x9b, 0x47, 0x4e, 0xdb, 0x4d, 0xdd, 0x0f, 0x9f, 0x34, 0x86, 0x02, 0x61, 0x60, 0x25, 0x82, 0x95,
	0x22, 0x08, 0xc9, 0x99, 0x92, 0x3a, 0x2e, 0x50, 0x58, 0x8c, 0x33, 0x81, 0xe0, 0x60, 0xf5, 0x7e,
	0x9b, 0x55, 0x95, 0xd8, 0xbe, 0x12, 0x7b, 0xdc, 0x77, 0x1e, 0x7b, 0x51, 0x43, 0x49, 0xfd, 0xb0,
	0x8d, 0x6c, 0x59, 0xc1, 0x88, 0x5c, 0xfc, 0xe5, 0xe3, 0xa5, 0xd4, 0x99, 0x59, 0xba, 0x73, 0xeb,
	0xfd, 0xd6, 0x01, 0x66, 0xb8, 0xbb, 0xd9, 0xd8, 0x8b, 0xce, 0x8b, 0x3d, 0xe2, 0xc9, 0x65, 0xc2,
	0x13, 0x72, 0x5c, 0x94, 0x2a, 0x8c, 0x3e, 0xd7, 0xd4, 0x5f, 0xad, 0xa9, 0xff, 0xbd, 0xa6, 0xfe,
	0xfb, 0x86, 0x7a, 0xab, 0x0d, 0xf5, 0xbe, 0x36, 0xd4, 0x7b, 0xbe, 0xcf, 0x25, 0x4e, 0xca, 0x84,
	0xa5, 0x46, 0x71, 0xf7, 0xdc, 0xb7, 0x1a, 0x70, 0x69, 0xec, 0x74, 0x37, 0xcd, 0x20, 0xcb, 0xc1,
	0xf2, 0xd7, 0x83, 0xcf, 0x4a, 0x6a, 0xae, 0xc0, 0xdd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4,
	0x16, 0xb2, 0x08, 0xcd, 0x01, 0x00, 0x00,
}

func (m *BasketCredit) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasketCredit) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasketCredit) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *DateCriteria) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DateCriteria) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DateCriteria) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *DateCriteria_MinStartDate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DateCriteria_MinStartDate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.MinStartDate != nil {
		{
			size, err := m.MinStartDate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *DateCriteria_StartDateWindow) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DateCriteria_StartDateWindow) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.StartDateWindow != nil {
		{
			size, err := m.StartDateWindow.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
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
func (m *BasketCredit) Size() (n int) {
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

func (m *DateCriteria) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *DateCriteria_MinStartDate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinStartDate != nil {
		l = m.MinStartDate.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *DateCriteria_StartDateWindow) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartDateWindow != nil {
		l = m.StartDateWindow.Size()
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
func (m *BasketCredit) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: BasketCredit: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasketCredit: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *DateCriteria) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DateCriteria: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DateCriteria: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinStartDate", wireType)
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
			v := &types.Timestamp{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &DateCriteria_MinStartDate{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDateWindow", wireType)
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
			v := &types.Duration{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &DateCriteria_StartDateWindow{v}
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
