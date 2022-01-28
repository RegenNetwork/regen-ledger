package datav1alpha2

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_MsgAnchorData        protoreflect.MessageDescriptor
	fd_MsgAnchorData_sender protoreflect.FieldDescriptor
	fd_MsgAnchorData_hash   protoreflect.FieldDescriptor
)

func init() {
	file_regen_data_v1alpha2_tx_proto_init()
	md_MsgAnchorData = File_regen_data_v1alpha2_tx_proto.Messages().ByName("MsgAnchorData")
	fd_MsgAnchorData_sender = md_MsgAnchorData.Fields().ByName("sender")
	fd_MsgAnchorData_hash = md_MsgAnchorData.Fields().ByName("hash")
}

var _ protoreflect.Message = (*fastReflection_MsgAnchorData)(nil)

type fastReflection_MsgAnchorData MsgAnchorData

func (x *MsgAnchorData) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgAnchorData)(x)
}

func (x *MsgAnchorData) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_data_v1alpha2_tx_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgAnchorData_messageType fastReflection_MsgAnchorData_messageType
var _ protoreflect.MessageType = fastReflection_MsgAnchorData_messageType{}

type fastReflection_MsgAnchorData_messageType struct{}

func (x fastReflection_MsgAnchorData_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgAnchorData)(nil)
}
func (x fastReflection_MsgAnchorData_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgAnchorData)
}
func (x fastReflection_MsgAnchorData_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAnchorData
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgAnchorData) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAnchorData
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgAnchorData) Type() protoreflect.MessageType {
	return _fastReflection_MsgAnchorData_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgAnchorData) New() protoreflect.Message {
	return new(fastReflection_MsgAnchorData)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgAnchorData) Interface() protoreflect.ProtoMessage {
	return (*MsgAnchorData)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgAnchorData) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sender != "" {
		value := protoreflect.ValueOfString(x.Sender)
		if !f(fd_MsgAnchorData_sender, value) {
			return
		}
	}
	if x.Hash != nil {
		value := protoreflect.ValueOfMessage(x.Hash.ProtoReflect())
		if !f(fd_MsgAnchorData_hash, value) {
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
func (x *fastReflection_MsgAnchorData) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgAnchorData.sender":
		return x.Sender != ""
	case "regen.data.v1alpha2.MsgAnchorData.hash":
		return x.Hash != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorData does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAnchorData) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgAnchorData.sender":
		x.Sender = ""
	case "regen.data.v1alpha2.MsgAnchorData.hash":
		x.Hash = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorData does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgAnchorData) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.data.v1alpha2.MsgAnchorData.sender":
		value := x.Sender
		return protoreflect.ValueOfString(value)
	case "regen.data.v1alpha2.MsgAnchorData.hash":
		value := x.Hash
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorData does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgAnchorData) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgAnchorData.sender":
		x.Sender = value.Interface().(string)
	case "regen.data.v1alpha2.MsgAnchorData.hash":
		x.Hash = value.Message().Interface().(*ContentHash)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorData does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgAnchorData) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgAnchorData.hash":
		if x.Hash == nil {
			x.Hash = new(ContentHash)
		}
		return protoreflect.ValueOfMessage(x.Hash.ProtoReflect())
	case "regen.data.v1alpha2.MsgAnchorData.sender":
		panic(fmt.Errorf("field sender of message regen.data.v1alpha2.MsgAnchorData is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorData does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgAnchorData) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgAnchorData.sender":
		return protoreflect.ValueOfString("")
	case "regen.data.v1alpha2.MsgAnchorData.hash":
		m := new(ContentHash)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorData does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgAnchorData) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.data.v1alpha2.MsgAnchorData", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgAnchorData) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAnchorData) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgAnchorData) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgAnchorData) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgAnchorData)
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
		l = len(x.Sender)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Hash != nil {
			l = options.Size(x.Hash)
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
		x := input.Message.Interface().(*MsgAnchorData)
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
		if x.Hash != nil {
			encoded, err := options.Marshal(x.Hash)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Sender) > 0 {
			i -= len(x.Sender)
			copy(dAtA[i:], x.Sender)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Sender)))
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
		x := input.Message.Interface().(*MsgAnchorData)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAnchorData: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAnchorData: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
				x.Sender = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Hash == nil {
					x.Hash = &ContentHash{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Hash); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
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

var (
	md_MsgAnchorDataResponse           protoreflect.MessageDescriptor
	fd_MsgAnchorDataResponse_timestamp protoreflect.FieldDescriptor
	fd_MsgAnchorDataResponse_iri       protoreflect.FieldDescriptor
)

func init() {
	file_regen_data_v1alpha2_tx_proto_init()
	md_MsgAnchorDataResponse = File_regen_data_v1alpha2_tx_proto.Messages().ByName("MsgAnchorDataResponse")
	fd_MsgAnchorDataResponse_timestamp = md_MsgAnchorDataResponse.Fields().ByName("timestamp")
	fd_MsgAnchorDataResponse_iri = md_MsgAnchorDataResponse.Fields().ByName("iri")
}

var _ protoreflect.Message = (*fastReflection_MsgAnchorDataResponse)(nil)

type fastReflection_MsgAnchorDataResponse MsgAnchorDataResponse

func (x *MsgAnchorDataResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgAnchorDataResponse)(x)
}

func (x *MsgAnchorDataResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_data_v1alpha2_tx_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgAnchorDataResponse_messageType fastReflection_MsgAnchorDataResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgAnchorDataResponse_messageType{}

type fastReflection_MsgAnchorDataResponse_messageType struct{}

func (x fastReflection_MsgAnchorDataResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgAnchorDataResponse)(nil)
}
func (x fastReflection_MsgAnchorDataResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgAnchorDataResponse)
}
func (x fastReflection_MsgAnchorDataResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAnchorDataResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgAnchorDataResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgAnchorDataResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgAnchorDataResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgAnchorDataResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgAnchorDataResponse) New() protoreflect.Message {
	return new(fastReflection_MsgAnchorDataResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgAnchorDataResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgAnchorDataResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgAnchorDataResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Timestamp != nil {
		value := protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
		if !f(fd_MsgAnchorDataResponse_timestamp, value) {
			return
		}
	}
	if x.Iri != "" {
		value := protoreflect.ValueOfString(x.Iri)
		if !f(fd_MsgAnchorDataResponse_iri, value) {
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
func (x *fastReflection_MsgAnchorDataResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgAnchorDataResponse.timestamp":
		return x.Timestamp != nil
	case "regen.data.v1alpha2.MsgAnchorDataResponse.iri":
		return x.Iri != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorDataResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAnchorDataResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgAnchorDataResponse.timestamp":
		x.Timestamp = nil
	case "regen.data.v1alpha2.MsgAnchorDataResponse.iri":
		x.Iri = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorDataResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgAnchorDataResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.data.v1alpha2.MsgAnchorDataResponse.timestamp":
		value := x.Timestamp
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.data.v1alpha2.MsgAnchorDataResponse.iri":
		value := x.Iri
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorDataResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgAnchorDataResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgAnchorDataResponse.timestamp":
		x.Timestamp = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.data.v1alpha2.MsgAnchorDataResponse.iri":
		x.Iri = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorDataResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgAnchorDataResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgAnchorDataResponse.timestamp":
		if x.Timestamp == nil {
			x.Timestamp = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
	case "regen.data.v1alpha2.MsgAnchorDataResponse.iri":
		panic(fmt.Errorf("field iri of message regen.data.v1alpha2.MsgAnchorDataResponse is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorDataResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgAnchorDataResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgAnchorDataResponse.timestamp":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.data.v1alpha2.MsgAnchorDataResponse.iri":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgAnchorDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgAnchorDataResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgAnchorDataResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.data.v1alpha2.MsgAnchorDataResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgAnchorDataResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgAnchorDataResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgAnchorDataResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgAnchorDataResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgAnchorDataResponse)
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
		if x.Timestamp != nil {
			l = options.Size(x.Timestamp)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Iri)
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
		x := input.Message.Interface().(*MsgAnchorDataResponse)
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
		if len(x.Iri) > 0 {
			i -= len(x.Iri)
			copy(dAtA[i:], x.Iri)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Iri)))
			i--
			dAtA[i] = 0x12
		}
		if x.Timestamp != nil {
			encoded, err := options.Marshal(x.Timestamp)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
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
		x := input.Message.Interface().(*MsgAnchorDataResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAnchorDataResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgAnchorDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Timestamp == nil {
					x.Timestamp = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Timestamp); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Iri", wireType)
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
				x.Iri = string(dAtA[iNdEx:postIndex])
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

var _ protoreflect.List = (*_MsgSignData_1_list)(nil)

type _MsgSignData_1_list struct {
	list *[]string
}

func (x *_MsgSignData_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_MsgSignData_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfString((*x.list)[i])
}

func (x *_MsgSignData_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	(*x.list)[i] = concreteValue
}

func (x *_MsgSignData_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.String()
	concreteValue := valueUnwrapped
	*x.list = append(*x.list, concreteValue)
}

func (x *_MsgSignData_1_list) AppendMutable() protoreflect.Value {
	panic(fmt.Errorf("AppendMutable can not be called on message MsgSignData at list field Signers as it is not of Message kind"))
}

func (x *_MsgSignData_1_list) Truncate(n int) {
	*x.list = (*x.list)[:n]
}

func (x *_MsgSignData_1_list) NewElement() protoreflect.Value {
	v := ""
	return protoreflect.ValueOfString(v)
}

func (x *_MsgSignData_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_MsgSignData         protoreflect.MessageDescriptor
	fd_MsgSignData_signers protoreflect.FieldDescriptor
	fd_MsgSignData_hash    protoreflect.FieldDescriptor
)

func init() {
	file_regen_data_v1alpha2_tx_proto_init()
	md_MsgSignData = File_regen_data_v1alpha2_tx_proto.Messages().ByName("MsgSignData")
	fd_MsgSignData_signers = md_MsgSignData.Fields().ByName("signers")
	fd_MsgSignData_hash = md_MsgSignData.Fields().ByName("hash")
}

var _ protoreflect.Message = (*fastReflection_MsgSignData)(nil)

type fastReflection_MsgSignData MsgSignData

func (x *MsgSignData) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSignData)(x)
}

func (x *MsgSignData) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_data_v1alpha2_tx_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSignData_messageType fastReflection_MsgSignData_messageType
var _ protoreflect.MessageType = fastReflection_MsgSignData_messageType{}

type fastReflection_MsgSignData_messageType struct{}

func (x fastReflection_MsgSignData_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSignData)(nil)
}
func (x fastReflection_MsgSignData_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSignData)
}
func (x fastReflection_MsgSignData_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSignData
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSignData) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSignData
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSignData) Type() protoreflect.MessageType {
	return _fastReflection_MsgSignData_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSignData) New() protoreflect.Message {
	return new(fastReflection_MsgSignData)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSignData) Interface() protoreflect.ProtoMessage {
	return (*MsgSignData)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSignData) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Signers) != 0 {
		value := protoreflect.ValueOfList(&_MsgSignData_1_list{list: &x.Signers})
		if !f(fd_MsgSignData_signers, value) {
			return
		}
	}
	if x.Hash != nil {
		value := protoreflect.ValueOfMessage(x.Hash.ProtoReflect())
		if !f(fd_MsgSignData_hash, value) {
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
func (x *fastReflection_MsgSignData) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgSignData.signers":
		return len(x.Signers) != 0
	case "regen.data.v1alpha2.MsgSignData.hash":
		return x.Hash != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignData does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSignData) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgSignData.signers":
		x.Signers = nil
	case "regen.data.v1alpha2.MsgSignData.hash":
		x.Hash = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignData does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSignData) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.data.v1alpha2.MsgSignData.signers":
		if len(x.Signers) == 0 {
			return protoreflect.ValueOfList(&_MsgSignData_1_list{})
		}
		listValue := &_MsgSignData_1_list{list: &x.Signers}
		return protoreflect.ValueOfList(listValue)
	case "regen.data.v1alpha2.MsgSignData.hash":
		value := x.Hash
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignData does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSignData) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgSignData.signers":
		lv := value.List()
		clv := lv.(*_MsgSignData_1_list)
		x.Signers = *clv.list
	case "regen.data.v1alpha2.MsgSignData.hash":
		x.Hash = value.Message().Interface().(*ContentHash_Graph)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignData does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSignData) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgSignData.signers":
		if x.Signers == nil {
			x.Signers = []string{}
		}
		value := &_MsgSignData_1_list{list: &x.Signers}
		return protoreflect.ValueOfList(value)
	case "regen.data.v1alpha2.MsgSignData.hash":
		if x.Hash == nil {
			x.Hash = new(ContentHash_Graph)
		}
		return protoreflect.ValueOfMessage(x.Hash.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignData does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSignData) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.MsgSignData.signers":
		list := []string{}
		return protoreflect.ValueOfList(&_MsgSignData_1_list{list: &list})
	case "regen.data.v1alpha2.MsgSignData.hash":
		m := new(ContentHash_Graph)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignData"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignData does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSignData) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.data.v1alpha2.MsgSignData", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSignData) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSignData) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSignData) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSignData) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSignData)
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
		if len(x.Signers) > 0 {
			for _, s := range x.Signers {
				l = len(s)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.Hash != nil {
			l = options.Size(x.Hash)
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
		x := input.Message.Interface().(*MsgSignData)
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
		if x.Hash != nil {
			encoded, err := options.Marshal(x.Hash)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Signers) > 0 {
			for iNdEx := len(x.Signers) - 1; iNdEx >= 0; iNdEx-- {
				i -= len(x.Signers[iNdEx])
				copy(dAtA[i:], x.Signers[iNdEx])
				i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signers[iNdEx])))
				i--
				dAtA[i] = 0xa
			}
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
		x := input.Message.Interface().(*MsgSignData)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSignData: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSignData: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signers", wireType)
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
				x.Signers = append(x.Signers, string(dAtA[iNdEx:postIndex]))
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Hash == nil {
					x.Hash = &ContentHash_Graph{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Hash); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
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

var (
	md_MsgSignDataResponse protoreflect.MessageDescriptor
)

func init() {
	file_regen_data_v1alpha2_tx_proto_init()
	md_MsgSignDataResponse = File_regen_data_v1alpha2_tx_proto.Messages().ByName("MsgSignDataResponse")
}

var _ protoreflect.Message = (*fastReflection_MsgSignDataResponse)(nil)

type fastReflection_MsgSignDataResponse MsgSignDataResponse

func (x *MsgSignDataResponse) ProtoReflect() protoreflect.Message {
	return (*fastReflection_MsgSignDataResponse)(x)
}

func (x *MsgSignDataResponse) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_data_v1alpha2_tx_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_MsgSignDataResponse_messageType fastReflection_MsgSignDataResponse_messageType
var _ protoreflect.MessageType = fastReflection_MsgSignDataResponse_messageType{}

type fastReflection_MsgSignDataResponse_messageType struct{}

func (x fastReflection_MsgSignDataResponse_messageType) Zero() protoreflect.Message {
	return (*fastReflection_MsgSignDataResponse)(nil)
}
func (x fastReflection_MsgSignDataResponse_messageType) New() protoreflect.Message {
	return new(fastReflection_MsgSignDataResponse)
}
func (x fastReflection_MsgSignDataResponse_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSignDataResponse
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_MsgSignDataResponse) Descriptor() protoreflect.MessageDescriptor {
	return md_MsgSignDataResponse
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_MsgSignDataResponse) Type() protoreflect.MessageType {
	return _fastReflection_MsgSignDataResponse_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_MsgSignDataResponse) New() protoreflect.Message {
	return new(fastReflection_MsgSignDataResponse)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_MsgSignDataResponse) Interface() protoreflect.ProtoMessage {
	return (*MsgSignDataResponse)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_MsgSignDataResponse) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
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
func (x *fastReflection_MsgSignDataResponse) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignDataResponse does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSignDataResponse) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignDataResponse does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_MsgSignDataResponse) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignDataResponse does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_MsgSignDataResponse) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignDataResponse does not contain field %s", fd.FullName()))
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
func (x *fastReflection_MsgSignDataResponse) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignDataResponse does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_MsgSignDataResponse) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.MsgSignDataResponse"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.MsgSignDataResponse does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_MsgSignDataResponse) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.data.v1alpha2.MsgSignDataResponse", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_MsgSignDataResponse) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_MsgSignDataResponse) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_MsgSignDataResponse) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_MsgSignDataResponse) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*MsgSignDataResponse)
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
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*MsgSignDataResponse)
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
		x := input.Message.Interface().(*MsgSignDataResponse)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSignDataResponse: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: MsgSignDataResponse: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
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
// 	protoc        (unknown)
// source: regen/data/v1alpha2/tx.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MsgAnchorData is the Msg/AnchorData request type.
type MsgAnchorData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sender is the address of the sender of the transaction.
	// The sender in StoreData is not attesting to the veracity of the underlying
	// data. They can simply be a intermediary providing services.
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// hash is the hash-based identifier for the anchored content.
	Hash *ContentHash `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *MsgAnchorData) Reset() {
	*x = MsgAnchorData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_data_v1alpha2_tx_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAnchorData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAnchorData) ProtoMessage() {}

// Deprecated: Use MsgAnchorData.ProtoReflect.Descriptor instead.
func (*MsgAnchorData) Descriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_tx_proto_rawDescGZIP(), []int{0}
}

func (x *MsgAnchorData) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *MsgAnchorData) GetHash() *ContentHash {
	if x != nil {
		return x.Hash
	}
	return nil
}

// MsgAnchorData is the Msg/AnchorData response type.
type MsgAnchorDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// timestamp is the timestamp of the block at which the data was anchored.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// iri is the IRI of the data that was anchored.
	Iri string `protobuf:"bytes,2,opt,name=iri,proto3" json:"iri,omitempty"`
}

func (x *MsgAnchorDataResponse) Reset() {
	*x = MsgAnchorDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_data_v1alpha2_tx_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgAnchorDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgAnchorDataResponse) ProtoMessage() {}

// Deprecated: Use MsgAnchorDataResponse.ProtoReflect.Descriptor instead.
func (*MsgAnchorDataResponse) Descriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_tx_proto_rawDescGZIP(), []int{1}
}

func (x *MsgAnchorDataResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *MsgAnchorDataResponse) GetIri() string {
	if x != nil {
		return x.Iri
	}
	return ""
}

// MsgSignData is the Msg/SignData request type.
type MsgSignData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signers are the addresses of the accounts signing the data.
	// By making a SignData request, the signers are attesting to the veracity
	// of the data referenced by the cid. The precise meaning of this may vary
	// depending on the underlying data.
	Signers []string `protobuf:"bytes,1,rep,name=signers,proto3" json:"signers,omitempty"`
	// hash is the hash-based identifier for the anchored content. Only RDF graph
	// data can be signed as its data model is intended to specifically convey semantic meaning.
	Hash *ContentHash_Graph `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *MsgSignData) Reset() {
	*x = MsgSignData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_data_v1alpha2_tx_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSignData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSignData) ProtoMessage() {}

// Deprecated: Use MsgSignData.ProtoReflect.Descriptor instead.
func (*MsgSignData) Descriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_tx_proto_rawDescGZIP(), []int{2}
}

func (x *MsgSignData) GetSigners() []string {
	if x != nil {
		return x.Signers
	}
	return nil
}

func (x *MsgSignData) GetHash() *ContentHash_Graph {
	if x != nil {
		return x.Hash
	}
	return nil
}

// MsgSignDataResponse is the Msg/SignData response type.
type MsgSignDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MsgSignDataResponse) Reset() {
	*x = MsgSignDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_data_v1alpha2_tx_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgSignDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgSignDataResponse) ProtoMessage() {}

// Deprecated: Use MsgSignDataResponse.ProtoReflect.Descriptor instead.
func (*MsgSignDataResponse) Descriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_tx_proto_rawDescGZIP(), []int{3}
}

var File_regen_data_v1alpha2_tx_proto protoreflect.FileDescriptor

var file_regen_data_v1alpha2_tx_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x74, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x32, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0d, 0x4d,
	0x73, 0x67, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x22, 0x63, 0x0a, 0x15, 0x4d, 0x73,
	0x67, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a,
	0x03, 0x69, 0x72, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x72, 0x69, 0x22,
	0x69, 0x0a, 0x0b, 0x4d, 0x73, 0x67, 0x53, 0x69, 0x67, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x3a, 0x04, 0x88, 0xa0, 0x1f, 0x00, 0x22, 0x15, 0x0a, 0x13, 0x4d, 0x73,
	0x67, 0x53, 0x69, 0x67, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xbb, 0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x5c, 0x0a, 0x0a, 0x41, 0x6e, 0x63,
	0x68, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73,
	0x67, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x2a, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x32, 0x2e, 0x4d, 0x73, 0x67, 0x41, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x08, 0x53, 0x69, 0x67, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x20, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x69, 0x67,
	0x6e, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x28, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x73, 0x67, 0x53,
	0x69, 0x67, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0xdc, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x07, 0x54, 0x78, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x32, 0xa2, 0x02, 0x03, 0x52, 0x44, 0x58, 0xaa, 0x02, 0x13, 0x52, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xca, 0x02,
	0x13, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0xe2, 0x02, 0x1f, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x44, 0x61, 0x74,
	0x61, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a,
	0x44, 0x61, 0x74, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_data_v1alpha2_tx_proto_rawDescOnce sync.Once
	file_regen_data_v1alpha2_tx_proto_rawDescData = file_regen_data_v1alpha2_tx_proto_rawDesc
)

func file_regen_data_v1alpha2_tx_proto_rawDescGZIP() []byte {
	file_regen_data_v1alpha2_tx_proto_rawDescOnce.Do(func() {
		file_regen_data_v1alpha2_tx_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_data_v1alpha2_tx_proto_rawDescData)
	})
	return file_regen_data_v1alpha2_tx_proto_rawDescData
}

var file_regen_data_v1alpha2_tx_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_regen_data_v1alpha2_tx_proto_goTypes = []interface{}{
	(*MsgAnchorData)(nil),         // 0: regen.data.v1alpha2.MsgAnchorData
	(*MsgAnchorDataResponse)(nil), // 1: regen.data.v1alpha2.MsgAnchorDataResponse
	(*MsgSignData)(nil),           // 2: regen.data.v1alpha2.MsgSignData
	(*MsgSignDataResponse)(nil),   // 3: regen.data.v1alpha2.MsgSignDataResponse
	(*ContentHash)(nil),           // 4: regen.data.v1alpha2.ContentHash
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*ContentHash_Graph)(nil),     // 6: regen.data.v1alpha2.ContentHash.Graph
}
var file_regen_data_v1alpha2_tx_proto_depIdxs = []int32{
	4, // 0: regen.data.v1alpha2.MsgAnchorData.hash:type_name -> regen.data.v1alpha2.ContentHash
	5, // 1: regen.data.v1alpha2.MsgAnchorDataResponse.timestamp:type_name -> google.protobuf.Timestamp
	6, // 2: regen.data.v1alpha2.MsgSignData.hash:type_name -> regen.data.v1alpha2.ContentHash.Graph
	0, // 3: regen.data.v1alpha2.Msg.AnchorData:input_type -> regen.data.v1alpha2.MsgAnchorData
	2, // 4: regen.data.v1alpha2.Msg.SignData:input_type -> regen.data.v1alpha2.MsgSignData
	1, // 5: regen.data.v1alpha2.Msg.AnchorData:output_type -> regen.data.v1alpha2.MsgAnchorDataResponse
	3, // 6: regen.data.v1alpha2.Msg.SignData:output_type -> regen.data.v1alpha2.MsgSignDataResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_regen_data_v1alpha2_tx_proto_init() }
func file_regen_data_v1alpha2_tx_proto_init() {
	if File_regen_data_v1alpha2_tx_proto != nil {
		return
	}
	file_regen_data_v1alpha2_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_regen_data_v1alpha2_tx_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAnchorData); i {
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
		file_regen_data_v1alpha2_tx_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgAnchorDataResponse); i {
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
		file_regen_data_v1alpha2_tx_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSignData); i {
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
		file_regen_data_v1alpha2_tx_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgSignDataResponse); i {
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
			RawDescriptor: file_regen_data_v1alpha2_tx_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_regen_data_v1alpha2_tx_proto_goTypes,
		DependencyIndexes: file_regen_data_v1alpha2_tx_proto_depIdxs,
		MessageInfos:      file_regen_data_v1alpha2_tx_proto_msgTypes,
	}.Build()
	File_regen_data_v1alpha2_tx_proto = out.File
	file_regen_data_v1alpha2_tx_proto_rawDesc = nil
	file_regen_data_v1alpha2_tx_proto_goTypes = nil
	file_regen_data_v1alpha2_tx_proto_depIdxs = nil
}