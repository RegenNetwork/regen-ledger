package datav1alpha2

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_ContentHash       protoreflect.MessageDescriptor
	fd_ContentHash_raw   protoreflect.FieldDescriptor
	fd_ContentHash_graph protoreflect.FieldDescriptor
)

func init() {
	file_regen_data_v1alpha2_types_proto_init()
	md_ContentHash = File_regen_data_v1alpha2_types_proto.Messages().ByName("ContentHash")
	fd_ContentHash_raw = md_ContentHash.Fields().ByName("raw")
	fd_ContentHash_graph = md_ContentHash.Fields().ByName("graph")
}

var _ protoreflect.Message = (*fastReflection_ContentHash)(nil)

type fastReflection_ContentHash ContentHash

func (x *ContentHash) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ContentHash)(x)
}

func (x *ContentHash) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_data_v1alpha2_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ContentHash_messageType fastReflection_ContentHash_messageType
var _ protoreflect.MessageType = fastReflection_ContentHash_messageType{}

type fastReflection_ContentHash_messageType struct{}

func (x fastReflection_ContentHash_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ContentHash)(nil)
}
func (x fastReflection_ContentHash_messageType) New() protoreflect.Message {
	return new(fastReflection_ContentHash)
}
func (x fastReflection_ContentHash_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ContentHash
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ContentHash) Descriptor() protoreflect.MessageDescriptor {
	return md_ContentHash
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ContentHash) Type() protoreflect.MessageType {
	return _fastReflection_ContentHash_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ContentHash) New() protoreflect.Message {
	return new(fastReflection_ContentHash)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ContentHash) Interface() protoreflect.ProtoMessage {
	return (*ContentHash)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ContentHash) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Sum != nil {
		switch o := x.Sum.(type) {
		case *ContentHash_Raw_:
			v := o.Raw
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_ContentHash_raw, value) {
				return
			}
		case *ContentHash_Graph_:
			v := o.Graph
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_ContentHash_graph, value) {
				return
			}
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
func (x *fastReflection_ContentHash) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.raw":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*ContentHash_Raw_); ok {
			return true
		} else {
			return false
		}
	case "regen.data.v1alpha2.ContentHash.graph":
		if x.Sum == nil {
			return false
		} else if _, ok := x.Sum.(*ContentHash_Graph_); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ContentHash) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.raw":
		x.Sum = nil
	case "regen.data.v1alpha2.ContentHash.graph":
		x.Sum = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ContentHash) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.data.v1alpha2.ContentHash.raw":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*ContentHash_Raw)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*ContentHash_Raw_); ok {
			return protoreflect.ValueOfMessage(v.Raw.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*ContentHash_Raw)(nil).ProtoReflect())
		}
	case "regen.data.v1alpha2.ContentHash.graph":
		if x.Sum == nil {
			return protoreflect.ValueOfMessage((*ContentHash_Graph)(nil).ProtoReflect())
		} else if v, ok := x.Sum.(*ContentHash_Graph_); ok {
			return protoreflect.ValueOfMessage(v.Graph.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*ContentHash_Graph)(nil).ProtoReflect())
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ContentHash) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.raw":
		cv := value.Message().Interface().(*ContentHash_Raw)
		x.Sum = &ContentHash_Raw_{Raw: cv}
	case "regen.data.v1alpha2.ContentHash.graph":
		cv := value.Message().Interface().(*ContentHash_Graph)
		x.Sum = &ContentHash_Graph_{Graph: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ContentHash) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.raw":
		if x.Sum == nil {
			value := &ContentHash_Raw{}
			oneofValue := &ContentHash_Raw_{Raw: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *ContentHash_Raw_:
			return protoreflect.ValueOfMessage(m.Raw.ProtoReflect())
		default:
			value := &ContentHash_Raw{}
			oneofValue := &ContentHash_Raw_{Raw: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "regen.data.v1alpha2.ContentHash.graph":
		if x.Sum == nil {
			value := &ContentHash_Graph{}
			oneofValue := &ContentHash_Graph_{Graph: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Sum.(type) {
		case *ContentHash_Graph_:
			return protoreflect.ValueOfMessage(m.Graph.ProtoReflect())
		default:
			value := &ContentHash_Graph{}
			oneofValue := &ContentHash_Graph_{Graph: value}
			x.Sum = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ContentHash) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.raw":
		value := &ContentHash_Raw{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.data.v1alpha2.ContentHash.graph":
		value := &ContentHash_Graph{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ContentHash) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "regen.data.v1alpha2.ContentHash.sum":
		if x.Sum == nil {
			return nil
		}
		switch x.Sum.(type) {
		case *ContentHash_Raw_:
			return x.Descriptor().Fields().ByName("raw")
		case *ContentHash_Graph_:
			return x.Descriptor().Fields().ByName("graph")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.data.v1alpha2.ContentHash", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ContentHash) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ContentHash) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ContentHash) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ContentHash) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ContentHash)
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
		switch x := x.Sum.(type) {
		case *ContentHash_Raw_:
			if x == nil {
				break
			}
			l = options.Size(x.Raw)
			n += 1 + l + runtime.Sov(uint64(l))
		case *ContentHash_Graph_:
			if x == nil {
				break
			}
			l = options.Size(x.Graph)
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
		x := input.Message.Interface().(*ContentHash)
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
		switch x := x.Sum.(type) {
		case *ContentHash_Raw_:
			encoded, err := options.Marshal(x.Raw)
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
		case *ContentHash_Graph_:
			encoded, err := options.Marshal(x.Graph)
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
		x := input.Message.Interface().(*ContentHash)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ContentHash: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ContentHash: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
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
				v := &ContentHash_Raw{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &ContentHash_Raw_{v}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Graph", wireType)
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
				v := &ContentHash_Graph{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Sum = &ContentHash_Graph_{v}
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
	md_ContentHash_Raw                  protoreflect.MessageDescriptor
	fd_ContentHash_Raw_hash             protoreflect.FieldDescriptor
	fd_ContentHash_Raw_digest_algorithm protoreflect.FieldDescriptor
	fd_ContentHash_Raw_media_type       protoreflect.FieldDescriptor
)

func init() {
	file_regen_data_v1alpha2_types_proto_init()
	md_ContentHash_Raw = File_regen_data_v1alpha2_types_proto.Messages().ByName("ContentHash").Messages().ByName("Raw")
	fd_ContentHash_Raw_hash = md_ContentHash_Raw.Fields().ByName("hash")
	fd_ContentHash_Raw_digest_algorithm = md_ContentHash_Raw.Fields().ByName("digest_algorithm")
	fd_ContentHash_Raw_media_type = md_ContentHash_Raw.Fields().ByName("media_type")
}

var _ protoreflect.Message = (*fastReflection_ContentHash_Raw)(nil)

type fastReflection_ContentHash_Raw ContentHash_Raw

func (x *ContentHash_Raw) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ContentHash_Raw)(x)
}

func (x *ContentHash_Raw) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_data_v1alpha2_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ContentHash_Raw_messageType fastReflection_ContentHash_Raw_messageType
var _ protoreflect.MessageType = fastReflection_ContentHash_Raw_messageType{}

type fastReflection_ContentHash_Raw_messageType struct{}

func (x fastReflection_ContentHash_Raw_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ContentHash_Raw)(nil)
}
func (x fastReflection_ContentHash_Raw_messageType) New() protoreflect.Message {
	return new(fastReflection_ContentHash_Raw)
}
func (x fastReflection_ContentHash_Raw_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ContentHash_Raw
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ContentHash_Raw) Descriptor() protoreflect.MessageDescriptor {
	return md_ContentHash_Raw
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ContentHash_Raw) Type() protoreflect.MessageType {
	return _fastReflection_ContentHash_Raw_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ContentHash_Raw) New() protoreflect.Message {
	return new(fastReflection_ContentHash_Raw)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ContentHash_Raw) Interface() protoreflect.ProtoMessage {
	return (*ContentHash_Raw)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ContentHash_Raw) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Hash) != 0 {
		value := protoreflect.ValueOfBytes(x.Hash)
		if !f(fd_ContentHash_Raw_hash, value) {
			return
		}
	}
	if x.DigestAlgorithm != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.DigestAlgorithm))
		if !f(fd_ContentHash_Raw_digest_algorithm, value) {
			return
		}
	}
	if x.MediaType != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.MediaType))
		if !f(fd_ContentHash_Raw_media_type, value) {
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
func (x *fastReflection_ContentHash_Raw) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.Raw.hash":
		return len(x.Hash) != 0
	case "regen.data.v1alpha2.ContentHash.Raw.digest_algorithm":
		return x.DigestAlgorithm != 0
	case "regen.data.v1alpha2.ContentHash.Raw.media_type":
		return x.MediaType != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Raw"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Raw does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ContentHash_Raw) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.Raw.hash":
		x.Hash = nil
	case "regen.data.v1alpha2.ContentHash.Raw.digest_algorithm":
		x.DigestAlgorithm = 0
	case "regen.data.v1alpha2.ContentHash.Raw.media_type":
		x.MediaType = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Raw"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Raw does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ContentHash_Raw) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.data.v1alpha2.ContentHash.Raw.hash":
		value := x.Hash
		return protoreflect.ValueOfBytes(value)
	case "regen.data.v1alpha2.ContentHash.Raw.digest_algorithm":
		value := x.DigestAlgorithm
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "regen.data.v1alpha2.ContentHash.Raw.media_type":
		value := x.MediaType
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Raw"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Raw does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ContentHash_Raw) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.Raw.hash":
		x.Hash = value.Bytes()
	case "regen.data.v1alpha2.ContentHash.Raw.digest_algorithm":
		x.DigestAlgorithm = (DigestAlgorithm)(value.Enum())
	case "regen.data.v1alpha2.ContentHash.Raw.media_type":
		x.MediaType = (MediaType)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Raw"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Raw does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ContentHash_Raw) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.Raw.hash":
		panic(fmt.Errorf("field hash of message regen.data.v1alpha2.ContentHash.Raw is not mutable"))
	case "regen.data.v1alpha2.ContentHash.Raw.digest_algorithm":
		panic(fmt.Errorf("field digest_algorithm of message regen.data.v1alpha2.ContentHash.Raw is not mutable"))
	case "regen.data.v1alpha2.ContentHash.Raw.media_type":
		panic(fmt.Errorf("field media_type of message regen.data.v1alpha2.ContentHash.Raw is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Raw"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Raw does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ContentHash_Raw) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.Raw.hash":
		return protoreflect.ValueOfBytes(nil)
	case "regen.data.v1alpha2.ContentHash.Raw.digest_algorithm":
		return protoreflect.ValueOfEnum(0)
	case "regen.data.v1alpha2.ContentHash.Raw.media_type":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Raw"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Raw does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ContentHash_Raw) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.data.v1alpha2.ContentHash.Raw", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ContentHash_Raw) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ContentHash_Raw) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ContentHash_Raw) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ContentHash_Raw) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ContentHash_Raw)
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
		l = len(x.Hash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DigestAlgorithm != 0 {
			n += 1 + runtime.Sov(uint64(x.DigestAlgorithm))
		}
		if x.MediaType != 0 {
			n += 1 + runtime.Sov(uint64(x.MediaType))
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
		x := input.Message.Interface().(*ContentHash_Raw)
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
		if x.MediaType != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MediaType))
			i--
			dAtA[i] = 0x18
		}
		if x.DigestAlgorithm != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.DigestAlgorithm))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Hash) > 0 {
			i -= len(x.Hash)
			copy(dAtA[i:], x.Hash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Hash)))
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
		x := input.Message.Interface().(*ContentHash_Raw)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ContentHash_Raw: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ContentHash_Raw: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Hash = append(x.Hash[:0], dAtA[iNdEx:postIndex]...)
				if x.Hash == nil {
					x.Hash = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DigestAlgorithm", wireType)
				}
				x.DigestAlgorithm = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.DigestAlgorithm |= DigestAlgorithm(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MediaType", wireType)
				}
				x.MediaType = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MediaType |= MediaType(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
	md_ContentHash_Graph                            protoreflect.MessageDescriptor
	fd_ContentHash_Graph_hash                       protoreflect.FieldDescriptor
	fd_ContentHash_Graph_digest_algorithm           protoreflect.FieldDescriptor
	fd_ContentHash_Graph_canonicalization_algorithm protoreflect.FieldDescriptor
	fd_ContentHash_Graph_merkle_tree                protoreflect.FieldDescriptor
)

func init() {
	file_regen_data_v1alpha2_types_proto_init()
	md_ContentHash_Graph = File_regen_data_v1alpha2_types_proto.Messages().ByName("ContentHash").Messages().ByName("Graph")
	fd_ContentHash_Graph_hash = md_ContentHash_Graph.Fields().ByName("hash")
	fd_ContentHash_Graph_digest_algorithm = md_ContentHash_Graph.Fields().ByName("digest_algorithm")
	fd_ContentHash_Graph_canonicalization_algorithm = md_ContentHash_Graph.Fields().ByName("canonicalization_algorithm")
	fd_ContentHash_Graph_merkle_tree = md_ContentHash_Graph.Fields().ByName("merkle_tree")
}

var _ protoreflect.Message = (*fastReflection_ContentHash_Graph)(nil)

type fastReflection_ContentHash_Graph ContentHash_Graph

func (x *ContentHash_Graph) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ContentHash_Graph)(x)
}

func (x *ContentHash_Graph) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_data_v1alpha2_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ContentHash_Graph_messageType fastReflection_ContentHash_Graph_messageType
var _ protoreflect.MessageType = fastReflection_ContentHash_Graph_messageType{}

type fastReflection_ContentHash_Graph_messageType struct{}

func (x fastReflection_ContentHash_Graph_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ContentHash_Graph)(nil)
}
func (x fastReflection_ContentHash_Graph_messageType) New() protoreflect.Message {
	return new(fastReflection_ContentHash_Graph)
}
func (x fastReflection_ContentHash_Graph_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ContentHash_Graph
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ContentHash_Graph) Descriptor() protoreflect.MessageDescriptor {
	return md_ContentHash_Graph
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ContentHash_Graph) Type() protoreflect.MessageType {
	return _fastReflection_ContentHash_Graph_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ContentHash_Graph) New() protoreflect.Message {
	return new(fastReflection_ContentHash_Graph)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ContentHash_Graph) Interface() protoreflect.ProtoMessage {
	return (*ContentHash_Graph)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ContentHash_Graph) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Hash) != 0 {
		value := protoreflect.ValueOfBytes(x.Hash)
		if !f(fd_ContentHash_Graph_hash, value) {
			return
		}
	}
	if x.DigestAlgorithm != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.DigestAlgorithm))
		if !f(fd_ContentHash_Graph_digest_algorithm, value) {
			return
		}
	}
	if x.CanonicalizationAlgorithm != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.CanonicalizationAlgorithm))
		if !f(fd_ContentHash_Graph_canonicalization_algorithm, value) {
			return
		}
	}
	if x.MerkleTree != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.MerkleTree))
		if !f(fd_ContentHash_Graph_merkle_tree, value) {
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
func (x *fastReflection_ContentHash_Graph) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.Graph.hash":
		return len(x.Hash) != 0
	case "regen.data.v1alpha2.ContentHash.Graph.digest_algorithm":
		return x.DigestAlgorithm != 0
	case "regen.data.v1alpha2.ContentHash.Graph.canonicalization_algorithm":
		return x.CanonicalizationAlgorithm != 0
	case "regen.data.v1alpha2.ContentHash.Graph.merkle_tree":
		return x.MerkleTree != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Graph"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Graph does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ContentHash_Graph) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.Graph.hash":
		x.Hash = nil
	case "regen.data.v1alpha2.ContentHash.Graph.digest_algorithm":
		x.DigestAlgorithm = 0
	case "regen.data.v1alpha2.ContentHash.Graph.canonicalization_algorithm":
		x.CanonicalizationAlgorithm = 0
	case "regen.data.v1alpha2.ContentHash.Graph.merkle_tree":
		x.MerkleTree = 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Graph"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Graph does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ContentHash_Graph) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.data.v1alpha2.ContentHash.Graph.hash":
		value := x.Hash
		return protoreflect.ValueOfBytes(value)
	case "regen.data.v1alpha2.ContentHash.Graph.digest_algorithm":
		value := x.DigestAlgorithm
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "regen.data.v1alpha2.ContentHash.Graph.canonicalization_algorithm":
		value := x.CanonicalizationAlgorithm
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "regen.data.v1alpha2.ContentHash.Graph.merkle_tree":
		value := x.MerkleTree
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Graph"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Graph does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ContentHash_Graph) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.Graph.hash":
		x.Hash = value.Bytes()
	case "regen.data.v1alpha2.ContentHash.Graph.digest_algorithm":
		x.DigestAlgorithm = (DigestAlgorithm)(value.Enum())
	case "regen.data.v1alpha2.ContentHash.Graph.canonicalization_algorithm":
		x.CanonicalizationAlgorithm = (GraphCanonicalizationAlgorithm)(value.Enum())
	case "regen.data.v1alpha2.ContentHash.Graph.merkle_tree":
		x.MerkleTree = (GraphMerkleTree)(value.Enum())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Graph"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Graph does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ContentHash_Graph) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.Graph.hash":
		panic(fmt.Errorf("field hash of message regen.data.v1alpha2.ContentHash.Graph is not mutable"))
	case "regen.data.v1alpha2.ContentHash.Graph.digest_algorithm":
		panic(fmt.Errorf("field digest_algorithm of message regen.data.v1alpha2.ContentHash.Graph is not mutable"))
	case "regen.data.v1alpha2.ContentHash.Graph.canonicalization_algorithm":
		panic(fmt.Errorf("field canonicalization_algorithm of message regen.data.v1alpha2.ContentHash.Graph is not mutable"))
	case "regen.data.v1alpha2.ContentHash.Graph.merkle_tree":
		panic(fmt.Errorf("field merkle_tree of message regen.data.v1alpha2.ContentHash.Graph is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Graph"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Graph does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ContentHash_Graph) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.ContentHash.Graph.hash":
		return protoreflect.ValueOfBytes(nil)
	case "regen.data.v1alpha2.ContentHash.Graph.digest_algorithm":
		return protoreflect.ValueOfEnum(0)
	case "regen.data.v1alpha2.ContentHash.Graph.canonicalization_algorithm":
		return protoreflect.ValueOfEnum(0)
	case "regen.data.v1alpha2.ContentHash.Graph.merkle_tree":
		return protoreflect.ValueOfEnum(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.ContentHash.Graph"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.ContentHash.Graph does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ContentHash_Graph) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.data.v1alpha2.ContentHash.Graph", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ContentHash_Graph) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ContentHash_Graph) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ContentHash_Graph) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ContentHash_Graph) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ContentHash_Graph)
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
		l = len(x.Hash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.DigestAlgorithm != 0 {
			n += 1 + runtime.Sov(uint64(x.DigestAlgorithm))
		}
		if x.CanonicalizationAlgorithm != 0 {
			n += 1 + runtime.Sov(uint64(x.CanonicalizationAlgorithm))
		}
		if x.MerkleTree != 0 {
			n += 1 + runtime.Sov(uint64(x.MerkleTree))
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
		x := input.Message.Interface().(*ContentHash_Graph)
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
		if x.MerkleTree != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.MerkleTree))
			i--
			dAtA[i] = 0x20
		}
		if x.CanonicalizationAlgorithm != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.CanonicalizationAlgorithm))
			i--
			dAtA[i] = 0x18
		}
		if x.DigestAlgorithm != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.DigestAlgorithm))
			i--
			dAtA[i] = 0x10
		}
		if len(x.Hash) > 0 {
			i -= len(x.Hash)
			copy(dAtA[i:], x.Hash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Hash)))
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
		x := input.Message.Interface().(*ContentHash_Graph)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ContentHash_Graph: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ContentHash_Graph: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Hash = append(x.Hash[:0], dAtA[iNdEx:postIndex]...)
				if x.Hash == nil {
					x.Hash = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DigestAlgorithm", wireType)
				}
				x.DigestAlgorithm = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.DigestAlgorithm |= DigestAlgorithm(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CanonicalizationAlgorithm", wireType)
				}
				x.CanonicalizationAlgorithm = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.CanonicalizationAlgorithm |= GraphCanonicalizationAlgorithm(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MerkleTree", wireType)
				}
				x.MerkleTree = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.MerkleTree |= GraphMerkleTree(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
	md_SignerEntry           protoreflect.MessageDescriptor
	fd_SignerEntry_signer    protoreflect.FieldDescriptor
	fd_SignerEntry_timestamp protoreflect.FieldDescriptor
)

func init() {
	file_regen_data_v1alpha2_types_proto_init()
	md_SignerEntry = File_regen_data_v1alpha2_types_proto.Messages().ByName("SignerEntry")
	fd_SignerEntry_signer = md_SignerEntry.Fields().ByName("signer")
	fd_SignerEntry_timestamp = md_SignerEntry.Fields().ByName("timestamp")
}

var _ protoreflect.Message = (*fastReflection_SignerEntry)(nil)

type fastReflection_SignerEntry SignerEntry

func (x *SignerEntry) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SignerEntry)(x)
}

func (x *SignerEntry) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_data_v1alpha2_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SignerEntry_messageType fastReflection_SignerEntry_messageType
var _ protoreflect.MessageType = fastReflection_SignerEntry_messageType{}

type fastReflection_SignerEntry_messageType struct{}

func (x fastReflection_SignerEntry_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SignerEntry)(nil)
}
func (x fastReflection_SignerEntry_messageType) New() protoreflect.Message {
	return new(fastReflection_SignerEntry)
}
func (x fastReflection_SignerEntry_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SignerEntry
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SignerEntry) Descriptor() protoreflect.MessageDescriptor {
	return md_SignerEntry
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SignerEntry) Type() protoreflect.MessageType {
	return _fastReflection_SignerEntry_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SignerEntry) New() protoreflect.Message {
	return new(fastReflection_SignerEntry)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SignerEntry) Interface() protoreflect.ProtoMessage {
	return (*SignerEntry)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SignerEntry) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Signer != "" {
		value := protoreflect.ValueOfString(x.Signer)
		if !f(fd_SignerEntry_signer, value) {
			return
		}
	}
	if x.Timestamp != nil {
		value := protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
		if !f(fd_SignerEntry_timestamp, value) {
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
func (x *fastReflection_SignerEntry) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.data.v1alpha2.SignerEntry.signer":
		return x.Signer != ""
	case "regen.data.v1alpha2.SignerEntry.timestamp":
		return x.Timestamp != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.SignerEntry"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.SignerEntry does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignerEntry) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.SignerEntry.signer":
		x.Signer = ""
	case "regen.data.v1alpha2.SignerEntry.timestamp":
		x.Timestamp = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.SignerEntry"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.SignerEntry does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SignerEntry) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.data.v1alpha2.SignerEntry.signer":
		value := x.Signer
		return protoreflect.ValueOfString(value)
	case "regen.data.v1alpha2.SignerEntry.timestamp":
		value := x.Timestamp
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.SignerEntry"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.SignerEntry does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SignerEntry) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.data.v1alpha2.SignerEntry.signer":
		x.Signer = value.Interface().(string)
	case "regen.data.v1alpha2.SignerEntry.timestamp":
		x.Timestamp = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.SignerEntry"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.SignerEntry does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SignerEntry) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.SignerEntry.timestamp":
		if x.Timestamp == nil {
			x.Timestamp = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.Timestamp.ProtoReflect())
	case "regen.data.v1alpha2.SignerEntry.signer":
		panic(fmt.Errorf("field signer of message regen.data.v1alpha2.SignerEntry is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.SignerEntry"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.SignerEntry does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SignerEntry) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.data.v1alpha2.SignerEntry.signer":
		return protoreflect.ValueOfString("")
	case "regen.data.v1alpha2.SignerEntry.timestamp":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.data.v1alpha2.SignerEntry"))
		}
		panic(fmt.Errorf("message regen.data.v1alpha2.SignerEntry does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SignerEntry) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.data.v1alpha2.SignerEntry", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SignerEntry) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SignerEntry) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SignerEntry) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SignerEntry) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SignerEntry)
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
		l = len(x.Signer)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Timestamp != nil {
			l = options.Size(x.Timestamp)
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
		x := input.Message.Interface().(*SignerEntry)
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
			dAtA[i] = 0x12
		}
		if len(x.Signer) > 0 {
			i -= len(x.Signer)
			copy(dAtA[i:], x.Signer)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Signer)))
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
		x := input.Message.Interface().(*SignerEntry)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignerEntry: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SignerEntry: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
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
				x.Signer = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
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
// source: regen/data/v1alpha2/types.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// MediaType defines MIME media types to be used with a ContentHash.Raw hash.
type MediaType int32

const (
	// MEDIA_TYPE_UNSPECIFIED can be used for raw binary data
	MediaType_MEDIA_TYPE_UNSPECIFIED MediaType = 0
	// plain text
	MediaType_MEDIA_TYPE_TEXT_PLAIN MediaType = 1
	// JSON
	MediaType_MEDIA_TYPE_JSON MediaType = 2
	// CSV
	MediaType_MEDIA_TYPE_CSV MediaType = 3
	// XML
	MediaType_MEDIA_TYPE_XML MediaType = 4
	// PDF
	MediaType_MEDIA_TYPE_PDF MediaType = 5
	// TIIF
	MediaType_MEDIA_TYPE_TIFF MediaType = 16
	// JPG
	MediaType_MEDIA_TYPE_JPG MediaType = 17
	// PNG
	MediaType_MEDIA_TYPE_PNG MediaType = 18
	// SVG
	MediaType_MEDIA_TYPE_SVG MediaType = 19
	// WEBP
	MediaType_MEDIA_TYPE_WEBP MediaType = 20
	// AVIF
	MediaType_MEDIA_TYPE_AVIF MediaType = 21
	// GIF
	MediaType_MEDIA_TYPE_GIF MediaType = 22
	// APNG
	MediaType_MEDIA_TYPE_APNG MediaType = 23
	// MPEG
	MediaType_MEDIA_TYPE_MPEG MediaType = 32
	// MP4
	MediaType_MEDIA_TYPE_MP4 MediaType = 33
	// WEBM
	MediaType_MEDIA_TYPE_WEBM MediaType = 34
	// OGG
	MediaType_MEDIA_TYPE_OGG MediaType = 35
)

// Enum value maps for MediaType.
var (
	MediaType_name = map[int32]string{
		0:  "MEDIA_TYPE_UNSPECIFIED",
		1:  "MEDIA_TYPE_TEXT_PLAIN",
		2:  "MEDIA_TYPE_JSON",
		3:  "MEDIA_TYPE_CSV",
		4:  "MEDIA_TYPE_XML",
		5:  "MEDIA_TYPE_PDF",
		16: "MEDIA_TYPE_TIFF",
		17: "MEDIA_TYPE_JPG",
		18: "MEDIA_TYPE_PNG",
		19: "MEDIA_TYPE_SVG",
		20: "MEDIA_TYPE_WEBP",
		21: "MEDIA_TYPE_AVIF",
		22: "MEDIA_TYPE_GIF",
		23: "MEDIA_TYPE_APNG",
		32: "MEDIA_TYPE_MPEG",
		33: "MEDIA_TYPE_MP4",
		34: "MEDIA_TYPE_WEBM",
		35: "MEDIA_TYPE_OGG",
	}
	MediaType_value = map[string]int32{
		"MEDIA_TYPE_UNSPECIFIED": 0,
		"MEDIA_TYPE_TEXT_PLAIN":  1,
		"MEDIA_TYPE_JSON":        2,
		"MEDIA_TYPE_CSV":         3,
		"MEDIA_TYPE_XML":         4,
		"MEDIA_TYPE_PDF":         5,
		"MEDIA_TYPE_TIFF":        16,
		"MEDIA_TYPE_JPG":         17,
		"MEDIA_TYPE_PNG":         18,
		"MEDIA_TYPE_SVG":         19,
		"MEDIA_TYPE_WEBP":        20,
		"MEDIA_TYPE_AVIF":        21,
		"MEDIA_TYPE_GIF":         22,
		"MEDIA_TYPE_APNG":        23,
		"MEDIA_TYPE_MPEG":        32,
		"MEDIA_TYPE_MP4":         33,
		"MEDIA_TYPE_WEBM":        34,
		"MEDIA_TYPE_OGG":         35,
	}
)

func (x MediaType) Enum() *MediaType {
	p := new(MediaType)
	*p = x
	return p
}

func (x MediaType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MediaType) Descriptor() protoreflect.EnumDescriptor {
	return file_regen_data_v1alpha2_types_proto_enumTypes[0].Descriptor()
}

func (MediaType) Type() protoreflect.EnumType {
	return &file_regen_data_v1alpha2_types_proto_enumTypes[0]
}

func (x MediaType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MediaType.Descriptor instead.
func (MediaType) EnumDescriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_types_proto_rawDescGZIP(), []int{0}
}

// GraphCanonicalizationAlgorithm is the graph canonicalization algorithm
type GraphCanonicalizationAlgorithm int32

const (
	// unspecified and invalid
	GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED GraphCanonicalizationAlgorithm = 0
	// URDNA2015 graph hashing
	GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015 GraphCanonicalizationAlgorithm = 1
)

// Enum value maps for GraphCanonicalizationAlgorithm.
var (
	GraphCanonicalizationAlgorithm_name = map[int32]string{
		0: "GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED",
		1: "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
	}
	GraphCanonicalizationAlgorithm_value = map[string]int32{
		"GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED": 0,
		"GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015":   1,
	}
)

func (x GraphCanonicalizationAlgorithm) Enum() *GraphCanonicalizationAlgorithm {
	p := new(GraphCanonicalizationAlgorithm)
	*p = x
	return p
}

func (x GraphCanonicalizationAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GraphCanonicalizationAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_regen_data_v1alpha2_types_proto_enumTypes[1].Descriptor()
}

func (GraphCanonicalizationAlgorithm) Type() protoreflect.EnumType {
	return &file_regen_data_v1alpha2_types_proto_enumTypes[1]
}

func (x GraphCanonicalizationAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GraphCanonicalizationAlgorithm.Descriptor instead.
func (GraphCanonicalizationAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_types_proto_rawDescGZIP(), []int{1}
}

// GraphMerkleTree is the graph merkle tree type used for hashing, if any
type GraphMerkleTree int32

const (
	// no merkle tree
	GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED GraphMerkleTree = 0
)

// Enum value maps for GraphMerkleTree.
var (
	GraphMerkleTree_name = map[int32]string{
		0: "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED",
	}
	GraphMerkleTree_value = map[string]int32{
		"GRAPH_MERKLE_TREE_NONE_UNSPECIFIED": 0,
	}
)

func (x GraphMerkleTree) Enum() *GraphMerkleTree {
	p := new(GraphMerkleTree)
	*p = x
	return p
}

func (x GraphMerkleTree) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GraphMerkleTree) Descriptor() protoreflect.EnumDescriptor {
	return file_regen_data_v1alpha2_types_proto_enumTypes[2].Descriptor()
}

func (GraphMerkleTree) Type() protoreflect.EnumType {
	return &file_regen_data_v1alpha2_types_proto_enumTypes[2]
}

func (x GraphMerkleTree) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GraphMerkleTree.Descriptor instead.
func (GraphMerkleTree) EnumDescriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_types_proto_rawDescGZIP(), []int{2}
}

// DigestAlgorithm is the hash digest algorithm
type DigestAlgorithm int32

const (
	// unspecified and invalid
	DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED DigestAlgorithm = 0
	// BLAKE2b-256
	DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256 DigestAlgorithm = 1
)

// Enum value maps for DigestAlgorithm.
var (
	DigestAlgorithm_name = map[int32]string{
		0: "DIGEST_ALGORITHM_UNSPECIFIED",
		1: "DIGEST_ALGORITHM_BLAKE2B_256",
	}
	DigestAlgorithm_value = map[string]int32{
		"DIGEST_ALGORITHM_UNSPECIFIED": 0,
		"DIGEST_ALGORITHM_BLAKE2B_256": 1,
	}
)

func (x DigestAlgorithm) Enum() *DigestAlgorithm {
	p := new(DigestAlgorithm)
	*p = x
	return p
}

func (x DigestAlgorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DigestAlgorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_regen_data_v1alpha2_types_proto_enumTypes[3].Descriptor()
}

func (DigestAlgorithm) Type() protoreflect.EnumType {
	return &file_regen_data_v1alpha2_types_proto_enumTypes[3]
}

func (x DigestAlgorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DigestAlgorithm.Descriptor instead.
func (DigestAlgorithm) EnumDescriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_types_proto_rawDescGZIP(), []int{3}
}

// ContentHash specifies a hash based content identifier for a piece of data
type ContentHash struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// sum selects the type of content hash
	//
	// Types that are assignable to Sum:
	//	*ContentHash_Raw_
	//	*ContentHash_Graph_
	Sum isContentHash_Sum `protobuf_oneof:"sum"`
}

func (x *ContentHash) Reset() {
	*x = ContentHash{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_data_v1alpha2_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentHash) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentHash) ProtoMessage() {}

// Deprecated: Use ContentHash.ProtoReflect.Descriptor instead.
func (*ContentHash) Descriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_types_proto_rawDescGZIP(), []int{0}
}

func (x *ContentHash) GetSum() isContentHash_Sum {
	if x != nil {
		return x.Sum
	}
	return nil
}

func (x *ContentHash) GetRaw() *ContentHash_Raw {
	if x, ok := x.GetSum().(*ContentHash_Raw_); ok {
		return x.Raw
	}
	return nil
}

func (x *ContentHash) GetGraph() *ContentHash_Graph {
	if x, ok := x.GetSum().(*ContentHash_Graph_); ok {
		return x.Graph
	}
	return nil
}

type isContentHash_Sum interface {
	isContentHash_Sum()
}

type ContentHash_Raw_ struct {
	// Raw specifies "raw" data which does not specify a deterministic, canonical encoding.
	// Users of these hashes MUST maintain a copy of the hashed data which is preserved bit by bit.
	// All other content encodings specify a deterministic, canonical encoding allowing implementations to
	// choose from a variety of alternative formats for transport and encoding while maintaining the guarantee
	// that the canonical hash will not change. The media type for "raw" data is defined by the MediaType enum.
	Raw *ContentHash_Raw `protobuf:"bytes,1,opt,name=raw,proto3,oneof"`
}

type ContentHash_Graph_ struct {
	// Graph specifies graph data that conforms to the RDF data model.
	// The canonicalization algorithm used for an RDF graph is specified by GraphCanonicalizationAlgorithm.
	Graph *ContentHash_Graph `protobuf:"bytes,2,opt,name=graph,proto3,oneof"`
}

func (*ContentHash_Raw_) isContentHash_Sum() {}

func (*ContentHash_Graph_) isContentHash_Sum() {}

// SignerEntry is a signer entry wrapping a signer address and timestamp
type SignerEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signer is the address of the signer
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// timestamp is the time at which the data was signed
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *SignerEntry) Reset() {
	*x = SignerEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_data_v1alpha2_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignerEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignerEntry) ProtoMessage() {}

// Deprecated: Use SignerEntry.ProtoReflect.Descriptor instead.
func (*SignerEntry) Descriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_types_proto_rawDescGZIP(), []int{1}
}

func (x *SignerEntry) GetSigner() string {
	if x != nil {
		return x.Signer
	}
	return ""
}

func (x *SignerEntry) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

// Raw is the content hash type used for raw data
type ContentHash_Raw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash represents the hash of the data based on the specified digest_algorithm
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// digest_algorithm represents the hash digest algorithm.
	DigestAlgorithm DigestAlgorithm `protobuf:"varint,2,opt,name=digest_algorithm,json=digestAlgorithm,proto3,enum=regen.data.v1alpha2.DigestAlgorithm" json:"digest_algorithm,omitempty"`
	// media_type represents the MediaType for raw data.
	MediaType MediaType `protobuf:"varint,3,opt,name=media_type,json=mediaType,proto3,enum=regen.data.v1alpha2.MediaType" json:"media_type,omitempty"`
}

func (x *ContentHash_Raw) Reset() {
	*x = ContentHash_Raw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_data_v1alpha2_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentHash_Raw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentHash_Raw) ProtoMessage() {}

// Deprecated: Use ContentHash_Raw.ProtoReflect.Descriptor instead.
func (*ContentHash_Raw) Descriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_types_proto_rawDescGZIP(), []int{0, 0}
}

func (x *ContentHash_Raw) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *ContentHash_Raw) GetDigestAlgorithm() DigestAlgorithm {
	if x != nil {
		return x.DigestAlgorithm
	}
	return DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED
}

func (x *ContentHash_Raw) GetMediaType() MediaType {
	if x != nil {
		return x.MediaType
	}
	return MediaType_MEDIA_TYPE_UNSPECIFIED
}

// Graph is the content hash type used for RDF graph data
type ContentHash_Graph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// hash represents the hash of the data based on the specified digest_algorithm
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// digest_algorithm represents the hash digest algorithm.
	DigestAlgorithm DigestAlgorithm `protobuf:"varint,2,opt,name=digest_algorithm,json=digestAlgorithm,proto3,enum=regen.data.v1alpha2.DigestAlgorithm" json:"digest_algorithm,omitempty"`
	// graph_canonicalization_algorithm represents the RDF graph canonicalization algorithm.
	CanonicalizationAlgorithm GraphCanonicalizationAlgorithm `protobuf:"varint,3,opt,name=canonicalization_algorithm,json=canonicalizationAlgorithm,proto3,enum=regen.data.v1alpha2.GraphCanonicalizationAlgorithm" json:"canonicalization_algorithm,omitempty"`
	// merkle_tree is the merkle tree type used for the graph hash, if any
	MerkleTree GraphMerkleTree `protobuf:"varint,4,opt,name=merkle_tree,json=merkleTree,proto3,enum=regen.data.v1alpha2.GraphMerkleTree" json:"merkle_tree,omitempty"`
}

func (x *ContentHash_Graph) Reset() {
	*x = ContentHash_Graph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_data_v1alpha2_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentHash_Graph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentHash_Graph) ProtoMessage() {}

// Deprecated: Use ContentHash_Graph.ProtoReflect.Descriptor instead.
func (*ContentHash_Graph) Descriptor() ([]byte, []int) {
	return file_regen_data_v1alpha2_types_proto_rawDescGZIP(), []int{0, 1}
}

func (x *ContentHash_Graph) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *ContentHash_Graph) GetDigestAlgorithm() DigestAlgorithm {
	if x != nil {
		return x.DigestAlgorithm
	}
	return DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED
}

func (x *ContentHash_Graph) GetCanonicalizationAlgorithm() GraphCanonicalizationAlgorithm {
	if x != nil {
		return x.CanonicalizationAlgorithm
	}
	return GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED
}

func (x *ContentHash_Graph) GetMerkleTree() GraphMerkleTree {
	if x != nil {
		return x.MerkleTree
	}
	return GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED
}

var File_regen_data_v1alpha2_types_proto protoreflect.FileDescriptor

var file_regen_data_v1alpha2_types_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x04, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x38, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x2e, 0x52, 0x61, 0x77, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61,
	0x77, 0x12, 0x3e, 0x0a, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x48, 0x00, 0x52, 0x05, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x1a, 0xa9, 0x01, 0x0a, 0x03, 0x52, 0x61, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x4f, 0x0a,
	0x10, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x44, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x0f, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x3d,
	0x0a, 0x0a, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x1a, 0xa7, 0x02,
	0x0a, 0x05, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x4f, 0x0a, 0x10, 0x64,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x0f, 0x64, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x72, 0x0a, 0x1a,
	0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x33, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x43, 0x61, 0x6e, 0x6f,
	0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x19, 0x63, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x12, 0x45, 0x0a, 0x0b, 0x6d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x65, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2e, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x52, 0x0a, 0x6d, 0x65, 0x72,
	0x6b, 0x6c, 0x65, 0x54, 0x72, 0x65, 0x65, 0x42, 0x05, 0x0a, 0x03, 0x73, 0x75, 0x6d, 0x22, 0x5f,
	0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2a,
	0x89, 0x03, 0x0a, 0x09, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x16, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x45, 0x44,
	0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x58, 0x54, 0x5f, 0x50, 0x4c, 0x41,
	0x49, 0x4e, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x44,
	0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x53, 0x56, 0x10, 0x03, 0x12, 0x12, 0x0a,
	0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x58, 0x4d, 0x4c, 0x10,
	0x04, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x50, 0x44, 0x46, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x54, 0x49, 0x46, 0x46, 0x10, 0x10, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45,
	0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x50, 0x47, 0x10, 0x11, 0x12, 0x12,
	0x0a, 0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x4e, 0x47,
	0x10, 0x12, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x53, 0x56, 0x47, 0x10, 0x13, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45, 0x42, 0x50, 0x10, 0x14, 0x12, 0x13, 0x0a, 0x0f, 0x4d,
	0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x56, 0x49, 0x46, 0x10, 0x15,
	0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x47,
	0x49, 0x46, 0x10, 0x16, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x50, 0x4e, 0x47, 0x10, 0x17, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x44,
	0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x50, 0x45, 0x47, 0x10, 0x20, 0x12, 0x12,
	0x0a, 0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x50, 0x34,
	0x10, 0x21, 0x12, 0x13, 0x0a, 0x0f, 0x4d, 0x45, 0x44, 0x49, 0x41, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x57, 0x45, 0x42, 0x4d, 0x10, 0x22, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x45, 0x44, 0x49, 0x41,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x47, 0x47, 0x10, 0x23, 0x2a, 0x82, 0x01, 0x0a, 0x1e,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x43, 0x61, 0x6e, 0x6f, 0x6e, 0x69, 0x63, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x30,
	0x0a, 0x2c, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f, 0x43, 0x41, 0x4e, 0x4f, 0x4e, 0x49, 0x43, 0x41,
	0x4c, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54,
	0x48, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x2e, 0x0a, 0x2a, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f, 0x43, 0x41, 0x4e, 0x4f, 0x4e, 0x49,
	0x43, 0x41, 0x4c, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52,
	0x49, 0x54, 0x48, 0x4d, 0x5f, 0x55, 0x52, 0x44, 0x4e, 0x41, 0x32, 0x30, 0x31, 0x35, 0x10, 0x01,
	0x2a, 0x39, 0x0a, 0x0f, 0x47, 0x72, 0x61, 0x70, 0x68, 0x4d, 0x65, 0x72, 0x6b, 0x6c, 0x65, 0x54,
	0x72, 0x65, 0x65, 0x12, 0x26, 0x0a, 0x22, 0x47, 0x52, 0x41, 0x50, 0x48, 0x5f, 0x4d, 0x45, 0x52,
	0x4b, 0x4c, 0x45, 0x5f, 0x54, 0x52, 0x45, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x2a, 0x55, 0x0a, 0x0f, 0x44,
	0x69, 0x67, 0x65, 0x73, 0x74, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x20,
	0x0a, 0x1c, 0x44, 0x49, 0x47, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52, 0x49, 0x54,
	0x48, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x20, 0x0a, 0x1c, 0x44, 0x49, 0x47, 0x45, 0x53, 0x54, 0x5f, 0x41, 0x4c, 0x47, 0x4f, 0x52,
	0x49, 0x54, 0x48, 0x4d, 0x5f, 0x42, 0x4c, 0x41, 0x4b, 0x45, 0x32, 0x42, 0x5f, 0x32, 0x35, 0x36,
	0x10, 0x01, 0x42, 0xdf, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x42, 0x0a,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x3b, 0x64, 0x61, 0x74, 0x61,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xa2, 0x02, 0x03, 0x52, 0x44, 0x58, 0xaa, 0x02,
	0x13, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0xca, 0x02, 0x13, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x44, 0x61, 0x74,
	0x61, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0xe2, 0x02, 0x1f, 0x52, 0x65, 0x67,
	0x65, 0x6e, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x52,
	0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_data_v1alpha2_types_proto_rawDescOnce sync.Once
	file_regen_data_v1alpha2_types_proto_rawDescData = file_regen_data_v1alpha2_types_proto_rawDesc
)

func file_regen_data_v1alpha2_types_proto_rawDescGZIP() []byte {
	file_regen_data_v1alpha2_types_proto_rawDescOnce.Do(func() {
		file_regen_data_v1alpha2_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_data_v1alpha2_types_proto_rawDescData)
	})
	return file_regen_data_v1alpha2_types_proto_rawDescData
}

var file_regen_data_v1alpha2_types_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_regen_data_v1alpha2_types_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_regen_data_v1alpha2_types_proto_goTypes = []interface{}{
	(MediaType)(0),                      // 0: regen.data.v1alpha2.MediaType
	(GraphCanonicalizationAlgorithm)(0), // 1: regen.data.v1alpha2.GraphCanonicalizationAlgorithm
	(GraphMerkleTree)(0),                // 2: regen.data.v1alpha2.GraphMerkleTree
	(DigestAlgorithm)(0),                // 3: regen.data.v1alpha2.DigestAlgorithm
	(*ContentHash)(nil),                 // 4: regen.data.v1alpha2.ContentHash
	(*SignerEntry)(nil),                 // 5: regen.data.v1alpha2.SignerEntry
	(*ContentHash_Raw)(nil),             // 6: regen.data.v1alpha2.ContentHash.Raw
	(*ContentHash_Graph)(nil),           // 7: regen.data.v1alpha2.ContentHash.Graph
	(*timestamppb.Timestamp)(nil),       // 8: google.protobuf.Timestamp
}
var file_regen_data_v1alpha2_types_proto_depIdxs = []int32{
	6, // 0: regen.data.v1alpha2.ContentHash.raw:type_name -> regen.data.v1alpha2.ContentHash.Raw
	7, // 1: regen.data.v1alpha2.ContentHash.graph:type_name -> regen.data.v1alpha2.ContentHash.Graph
	8, // 2: regen.data.v1alpha2.SignerEntry.timestamp:type_name -> google.protobuf.Timestamp
	3, // 3: regen.data.v1alpha2.ContentHash.Raw.digest_algorithm:type_name -> regen.data.v1alpha2.DigestAlgorithm
	0, // 4: regen.data.v1alpha2.ContentHash.Raw.media_type:type_name -> regen.data.v1alpha2.MediaType
	3, // 5: regen.data.v1alpha2.ContentHash.Graph.digest_algorithm:type_name -> regen.data.v1alpha2.DigestAlgorithm
	1, // 6: regen.data.v1alpha2.ContentHash.Graph.canonicalization_algorithm:type_name -> regen.data.v1alpha2.GraphCanonicalizationAlgorithm
	2, // 7: regen.data.v1alpha2.ContentHash.Graph.merkle_tree:type_name -> regen.data.v1alpha2.GraphMerkleTree
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_regen_data_v1alpha2_types_proto_init() }
func file_regen_data_v1alpha2_types_proto_init() {
	if File_regen_data_v1alpha2_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_data_v1alpha2_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentHash); i {
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
		file_regen_data_v1alpha2_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignerEntry); i {
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
		file_regen_data_v1alpha2_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentHash_Raw); i {
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
		file_regen_data_v1alpha2_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContentHash_Graph); i {
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
	file_regen_data_v1alpha2_types_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ContentHash_Raw_)(nil),
		(*ContentHash_Graph_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_data_v1alpha2_types_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_data_v1alpha2_types_proto_goTypes,
		DependencyIndexes: file_regen_data_v1alpha2_types_proto_depIdxs,
		EnumInfos:         file_regen_data_v1alpha2_types_proto_enumTypes,
		MessageInfos:      file_regen_data_v1alpha2_types_proto_msgTypes,
	}.Build()
	File_regen_data_v1alpha2_types_proto = out.File
	file_regen_data_v1alpha2_types_proto_rawDesc = nil
	file_regen_data_v1alpha2_types_proto_goTypes = nil
	file_regen_data_v1alpha2_types_proto_depIdxs = nil
}
