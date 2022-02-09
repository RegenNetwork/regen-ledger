// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package basketv1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_BasketBalance              protoreflect.MessageDescriptor
	fd_BasketBalance_basket_denom protoreflect.FieldDescriptor
	fd_BasketBalance_batch_id     protoreflect.FieldDescriptor
	fd_BasketBalance_balance      protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_basket_v1beta1_state_proto_init()
	md_BasketBalance = File_regen_ecocredit_basket_v1beta1_state_proto.Messages().ByName("BasketBalance")
	fd_BasketBalance_basket_denom = md_BasketBalance.Fields().ByName("basket_denom")
	fd_BasketBalance_batch_id = md_BasketBalance.Fields().ByName("batch_id")
	fd_BasketBalance_balance = md_BasketBalance.Fields().ByName("balance")
}

var _ protoreflect.Message = (*fastReflection_BasketBalance)(nil)

type fastReflection_BasketBalance BasketBalance

func (x *BasketBalance) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BasketBalance)(x)
}

func (x *BasketBalance) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_basket_v1beta1_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BasketBalance_messageType fastReflection_BasketBalance_messageType
var _ protoreflect.MessageType = fastReflection_BasketBalance_messageType{}

type fastReflection_BasketBalance_messageType struct{}

func (x fastReflection_BasketBalance_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BasketBalance)(nil)
}
func (x fastReflection_BasketBalance_messageType) New() protoreflect.Message {
	return new(fastReflection_BasketBalance)
}
func (x fastReflection_BasketBalance_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BasketBalance
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BasketBalance) Descriptor() protoreflect.MessageDescriptor {
	return md_BasketBalance
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BasketBalance) Type() protoreflect.MessageType {
	return _fastReflection_BasketBalance_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BasketBalance) New() protoreflect.Message {
	return new(fastReflection_BasketBalance)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BasketBalance) Interface() protoreflect.ProtoMessage {
	return (*BasketBalance)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BasketBalance) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BasketDenom != "" {
		value := protoreflect.ValueOfString(x.BasketDenom)
		if !f(fd_BasketBalance_basket_denom, value) {
			return
		}
	}
	if x.BatchId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BatchId)
		if !f(fd_BasketBalance_batch_id, value) {
			return
		}
	}
	if x.Balance != "" {
		value := protoreflect.ValueOfString(x.Balance)
		if !f(fd_BasketBalance_balance, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BasketBalance) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.basket.v1beta1.BasketBalance.basket_denom":
		return x.BasketDenom != ""
	case "regen.ecocredit.basket.v1beta1.BasketBalance.batch_id":
		return x.BatchId != uint64(0)
	case "regen.ecocredit.basket.v1beta1.BasketBalance.balance":
		return x.Balance != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.basket.v1beta1.BasketBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.basket.v1beta1.BasketBalance does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketBalance) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.basket.v1beta1.BasketBalance.basket_denom":
		x.BasketDenom = ""
	case "regen.ecocredit.basket.v1beta1.BasketBalance.batch_id":
		x.BatchId = uint64(0)
	case "regen.ecocredit.basket.v1beta1.BasketBalance.balance":
		x.Balance = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.basket.v1beta1.BasketBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.basket.v1beta1.BasketBalance does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BasketBalance) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.basket.v1beta1.BasketBalance.basket_denom":
		value := x.BasketDenom
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.basket.v1beta1.BasketBalance.batch_id":
		value := x.BatchId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.basket.v1beta1.BasketBalance.balance":
		value := x.Balance
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.basket.v1beta1.BasketBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.basket.v1beta1.BasketBalance does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketBalance) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.basket.v1beta1.BasketBalance.basket_denom":
		x.BasketDenom = value.Interface().(string)
	case "regen.ecocredit.basket.v1beta1.BasketBalance.batch_id":
		x.BatchId = value.Uint()
	case "regen.ecocredit.basket.v1beta1.BasketBalance.balance":
		x.Balance = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.basket.v1beta1.BasketBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.basket.v1beta1.BasketBalance does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketBalance) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.basket.v1beta1.BasketBalance.basket_denom":
		panic(fmt.Errorf("field basket_denom of message regen.ecocredit.basket.v1beta1.BasketBalance is not mutable"))
	case "regen.ecocredit.basket.v1beta1.BasketBalance.batch_id":
		panic(fmt.Errorf("field batch_id of message regen.ecocredit.basket.v1beta1.BasketBalance is not mutable"))
	case "regen.ecocredit.basket.v1beta1.BasketBalance.balance":
		panic(fmt.Errorf("field balance of message regen.ecocredit.basket.v1beta1.BasketBalance is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.basket.v1beta1.BasketBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.basket.v1beta1.BasketBalance does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BasketBalance) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.basket.v1beta1.BasketBalance.basket_denom":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.basket.v1beta1.BasketBalance.batch_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.basket.v1beta1.BasketBalance.balance":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.basket.v1beta1.BasketBalance"))
		}
		panic(fmt.Errorf("message regen.ecocredit.basket.v1beta1.BasketBalance does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BasketBalance) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.basket.v1beta1.BasketBalance", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BasketBalance) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BasketBalance) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BasketBalance) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BasketBalance) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BasketBalance)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.BasketDenom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.BatchId != 0 {
			n += 1 + runtime.Sov(uint64(x.BatchId))
		}
		l = len(x.Balance)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BasketBalance)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Balance) > 0 {
			i -= len(x.Balance)
			copy(dAtA[i:], x.Balance)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Balance)))
			i--
			dAtA[i] = 0x1a
		}
		if x.BatchId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BatchId))
			i--
			dAtA[i] = 0x10
		}
		if len(x.BasketDenom) > 0 {
			i -= len(x.BasketDenom)
			copy(dAtA[i:], x.BasketDenom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.BasketDenom)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BasketBalance)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BasketBalance: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BasketBalance: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.BasketDenom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchId", wireType)
				}
				x.BatchId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BatchId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Balance = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        v3.19.1
// source: regen/ecocredit/basket/v1beta1/state.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// BasketBalance stores the amount of credits from a batch in a basket
type BasketBalance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// basket_denom is the denom of the basket
	BasketDenom string `protobuf:"bytes,1,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
	// batch_id is the id of the credit batch
	BatchId uint64 `protobuf:"varint,2,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	// balance is the amount of ecocredits held in the basket
	Balance string `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *BasketBalance) Reset() {
	*x = BasketBalance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_basket_v1beta1_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasketBalance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasketBalance) ProtoMessage() {}

// Deprecated: Use BasketBalance.ProtoReflect.Descriptor instead.
func (*BasketBalance) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_basket_v1beta1_state_proto_rawDescGZIP(), []int{0}
}

func (x *BasketBalance) GetBasketDenom() string {
	if x != nil {
		return x.BasketDenom
	}
	return ""
}

func (x *BasketBalance) GetBatchId() uint64 {
	if x != nil {
		return x.BatchId
	}
	return 0
}

func (x *BasketBalance) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

var File_regen_ecocredit_basket_v1beta1_state_proto protoreflect.FileDescriptor

var file_regen_ecocredit_basket_v1beta1_state_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x62, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1d, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x01, 0x0a, 0x0d,
	0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x21, 0xf2, 0x9e, 0xd3, 0x8e, 0x03, 0x1b, 0x0a, 0x17, 0x0a,
	0x15, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x2c, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x42, 0xa3, 0x02, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42,
	0x0a, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x56, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65,
	0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x2f,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x62, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x45, 0x42, 0xaa, 0x02, 0x1e, 0x52, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x42, 0x61,
	0x73, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1e, 0x52,
	0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x2a,
	0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c,
	0x42, 0x61, 0x73, 0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x52, 0x65, 0x67,
	0x65, 0x6e, 0x3a, 0x3a, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x42,
	0x61, 0x73, 0x6b, 0x65, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_basket_v1beta1_state_proto_rawDescOnce sync.Once
	file_regen_ecocredit_basket_v1beta1_state_proto_rawDescData = file_regen_ecocredit_basket_v1beta1_state_proto_rawDesc
)

func file_regen_ecocredit_basket_v1beta1_state_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_basket_v1beta1_state_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_basket_v1beta1_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_basket_v1beta1_state_proto_rawDescData)
	})
	return file_regen_ecocredit_basket_v1beta1_state_proto_rawDescData
}

var file_regen_ecocredit_basket_v1beta1_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_regen_ecocredit_basket_v1beta1_state_proto_goTypes = []interface{}{
	(*BasketBalance)(nil), // 0: regen.ecocredit.basket.v1beta1.BasketBalance
}
var file_regen_ecocredit_basket_v1beta1_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_basket_v1beta1_state_proto_init() }
func file_regen_ecocredit_basket_v1beta1_state_proto_init() {
	if File_regen_ecocredit_basket_v1beta1_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_basket_v1beta1_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasketBalance); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_basket_v1beta1_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_basket_v1beta1_state_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_basket_v1beta1_state_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_basket_v1beta1_state_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_basket_v1beta1_state_proto = out.File
	file_regen_ecocredit_basket_v1beta1_state_proto_rawDesc = nil
	file_regen_ecocredit_basket_v1beta1_state_proto_goTypes = nil
	file_regen_ecocredit_basket_v1beta1_state_proto_depIdxs = nil
}
