package groupv1alpha1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/gogo/protobuf/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_GenesisState_2_list)(nil)

type _GenesisState_2_list struct {
	list *[]*GroupInfo
}

func (x *_GenesisState_2_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_2_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_2_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*GroupInfo)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_2_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*GroupInfo)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_2_list) AppendMutable() protoreflect.Value {
	v := new(GroupInfo)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_2_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_2_list) NewElement() protoreflect.Value {
	v := new(GroupInfo)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_2_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_3_list)(nil)

type _GenesisState_3_list struct {
	list *[]*GroupMember
}

func (x *_GenesisState_3_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_3_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_3_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*GroupMember)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_3_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*GroupMember)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_3_list) AppendMutable() protoreflect.Value {
	v := new(GroupMember)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_3_list) NewElement() protoreflect.Value {
	v := new(GroupMember)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_3_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_5_list)(nil)

type _GenesisState_5_list struct {
	list *[]*GroupAccountInfo
}

func (x *_GenesisState_5_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_5_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_5_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*GroupAccountInfo)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_5_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*GroupAccountInfo)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_5_list) AppendMutable() protoreflect.Value {
	v := new(GroupAccountInfo)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_5_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_5_list) NewElement() protoreflect.Value {
	v := new(GroupAccountInfo)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_5_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_7_list)(nil)

type _GenesisState_7_list struct {
	list *[]*Proposal
}

func (x *_GenesisState_7_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_7_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_7_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Proposal)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_7_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Proposal)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_7_list) AppendMutable() protoreflect.Value {
	v := new(Proposal)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_7_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_7_list) NewElement() protoreflect.Value {
	v := new(Proposal)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_7_list) IsValid() bool {
	return x.list != nil
}

var _ protoreflect.List = (*_GenesisState_8_list)(nil)

type _GenesisState_8_list struct {
	list *[]*Vote
}

func (x *_GenesisState_8_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_GenesisState_8_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_GenesisState_8_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Vote)
	(*x.list)[i] = concreteValue
}

func (x *_GenesisState_8_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Vote)
	*x.list = append(*x.list, concreteValue)
}

func (x *_GenesisState_8_list) AppendMutable() protoreflect.Value {
	v := new(Vote)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_8_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_GenesisState_8_list) NewElement() protoreflect.Value {
	v := new(Vote)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_GenesisState_8_list) IsValid() bool {
	return x.list != nil
}

var (
	md_GenesisState                   protoreflect.MessageDescriptor
	fd_GenesisState_group_seq         protoreflect.FieldDescriptor
	fd_GenesisState_groups            protoreflect.FieldDescriptor
	fd_GenesisState_group_members     protoreflect.FieldDescriptor
	fd_GenesisState_group_account_seq protoreflect.FieldDescriptor
	fd_GenesisState_group_accounts    protoreflect.FieldDescriptor
	fd_GenesisState_proposal_seq      protoreflect.FieldDescriptor
	fd_GenesisState_proposals         protoreflect.FieldDescriptor
	fd_GenesisState_votes             protoreflect.FieldDescriptor
)

func init() {
	file_regen_group_v1alpha1_genesis_proto_init()
	md_GenesisState = File_regen_group_v1alpha1_genesis_proto.Messages().ByName("GenesisState")
	fd_GenesisState_group_seq = md_GenesisState.Fields().ByName("group_seq")
	fd_GenesisState_groups = md_GenesisState.Fields().ByName("groups")
	fd_GenesisState_group_members = md_GenesisState.Fields().ByName("group_members")
	fd_GenesisState_group_account_seq = md_GenesisState.Fields().ByName("group_account_seq")
	fd_GenesisState_group_accounts = md_GenesisState.Fields().ByName("group_accounts")
	fd_GenesisState_proposal_seq = md_GenesisState.Fields().ByName("proposal_seq")
	fd_GenesisState_proposals = md_GenesisState.Fields().ByName("proposals")
	fd_GenesisState_votes = md_GenesisState.Fields().ByName("votes")
}

var _ protoreflect.Message = (*fastReflection_GenesisState)(nil)

type fastReflection_GenesisState GenesisState

func (x *GenesisState) ProtoReflect() protoreflect.Message {
	return (*fastReflection_GenesisState)(x)
}

func (x *GenesisState) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_group_v1alpha1_genesis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_GenesisState_messageType fastReflection_GenesisState_messageType
var _ protoreflect.MessageType = fastReflection_GenesisState_messageType{}

type fastReflection_GenesisState_messageType struct{}

func (x fastReflection_GenesisState_messageType) Zero() protoreflect.Message {
	return (*fastReflection_GenesisState)(nil)
}
func (x fastReflection_GenesisState_messageType) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}
func (x fastReflection_GenesisState_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_GenesisState) Descriptor() protoreflect.MessageDescriptor {
	return md_GenesisState
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_GenesisState) Type() protoreflect.MessageType {
	return _fastReflection_GenesisState_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_GenesisState) New() protoreflect.Message {
	return new(fastReflection_GenesisState)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_GenesisState) Interface() protoreflect.ProtoMessage {
	return (*GenesisState)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_GenesisState) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.GroupSeq != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupSeq)
		if !f(fd_GenesisState_group_seq, value) {
			return
		}
	}
	if len(x.Groups) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_2_list{list: &x.Groups})
		if !f(fd_GenesisState_groups, value) {
			return
		}
	}
	if len(x.GroupMembers) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_3_list{list: &x.GroupMembers})
		if !f(fd_GenesisState_group_members, value) {
			return
		}
	}
	if x.GroupAccountSeq != uint64(0) {
		value := protoreflect.ValueOfUint64(x.GroupAccountSeq)
		if !f(fd_GenesisState_group_account_seq, value) {
			return
		}
	}
	if len(x.GroupAccounts) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_5_list{list: &x.GroupAccounts})
		if !f(fd_GenesisState_group_accounts, value) {
			return
		}
	}
	if x.ProposalSeq != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProposalSeq)
		if !f(fd_GenesisState_proposal_seq, value) {
			return
		}
	}
	if len(x.Proposals) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_7_list{list: &x.Proposals})
		if !f(fd_GenesisState_proposals, value) {
			return
		}
	}
	if len(x.Votes) != 0 {
		value := protoreflect.ValueOfList(&_GenesisState_8_list{list: &x.Votes})
		if !f(fd_GenesisState_votes, value) {
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
func (x *fastReflection_GenesisState) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GenesisState.group_seq":
		return x.GroupSeq != uint64(0)
	case "regen.group.v1alpha1.GenesisState.groups":
		return len(x.Groups) != 0
	case "regen.group.v1alpha1.GenesisState.group_members":
		return len(x.GroupMembers) != 0
	case "regen.group.v1alpha1.GenesisState.group_account_seq":
		return x.GroupAccountSeq != uint64(0)
	case "regen.group.v1alpha1.GenesisState.group_accounts":
		return len(x.GroupAccounts) != 0
	case "regen.group.v1alpha1.GenesisState.proposal_seq":
		return x.ProposalSeq != uint64(0)
	case "regen.group.v1alpha1.GenesisState.proposals":
		return len(x.Proposals) != 0
	case "regen.group.v1alpha1.GenesisState.votes":
		return len(x.Votes) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GenesisState"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GenesisState.group_seq":
		x.GroupSeq = uint64(0)
	case "regen.group.v1alpha1.GenesisState.groups":
		x.Groups = nil
	case "regen.group.v1alpha1.GenesisState.group_members":
		x.GroupMembers = nil
	case "regen.group.v1alpha1.GenesisState.group_account_seq":
		x.GroupAccountSeq = uint64(0)
	case "regen.group.v1alpha1.GenesisState.group_accounts":
		x.GroupAccounts = nil
	case "regen.group.v1alpha1.GenesisState.proposal_seq":
		x.ProposalSeq = uint64(0)
	case "regen.group.v1alpha1.GenesisState.proposals":
		x.Proposals = nil
	case "regen.group.v1alpha1.GenesisState.votes":
		x.Votes = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GenesisState"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_GenesisState) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.group.v1alpha1.GenesisState.group_seq":
		value := x.GroupSeq
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.GenesisState.groups":
		if len(x.Groups) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_2_list{})
		}
		listValue := &_GenesisState_2_list{list: &x.Groups}
		return protoreflect.ValueOfList(listValue)
	case "regen.group.v1alpha1.GenesisState.group_members":
		if len(x.GroupMembers) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_3_list{})
		}
		listValue := &_GenesisState_3_list{list: &x.GroupMembers}
		return protoreflect.ValueOfList(listValue)
	case "regen.group.v1alpha1.GenesisState.group_account_seq":
		value := x.GroupAccountSeq
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.GenesisState.group_accounts":
		if len(x.GroupAccounts) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_5_list{})
		}
		listValue := &_GenesisState_5_list{list: &x.GroupAccounts}
		return protoreflect.ValueOfList(listValue)
	case "regen.group.v1alpha1.GenesisState.proposal_seq":
		value := x.ProposalSeq
		return protoreflect.ValueOfUint64(value)
	case "regen.group.v1alpha1.GenesisState.proposals":
		if len(x.Proposals) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_7_list{})
		}
		listValue := &_GenesisState_7_list{list: &x.Proposals}
		return protoreflect.ValueOfList(listValue)
	case "regen.group.v1alpha1.GenesisState.votes":
		if len(x.Votes) == 0 {
			return protoreflect.ValueOfList(&_GenesisState_8_list{})
		}
		listValue := &_GenesisState_8_list{list: &x.Votes}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GenesisState"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GenesisState does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_GenesisState) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GenesisState.group_seq":
		x.GroupSeq = value.Uint()
	case "regen.group.v1alpha1.GenesisState.groups":
		lv := value.List()
		clv := lv.(*_GenesisState_2_list)
		x.Groups = *clv.list
	case "regen.group.v1alpha1.GenesisState.group_members":
		lv := value.List()
		clv := lv.(*_GenesisState_3_list)
		x.GroupMembers = *clv.list
	case "regen.group.v1alpha1.GenesisState.group_account_seq":
		x.GroupAccountSeq = value.Uint()
	case "regen.group.v1alpha1.GenesisState.group_accounts":
		lv := value.List()
		clv := lv.(*_GenesisState_5_list)
		x.GroupAccounts = *clv.list
	case "regen.group.v1alpha1.GenesisState.proposal_seq":
		x.ProposalSeq = value.Uint()
	case "regen.group.v1alpha1.GenesisState.proposals":
		lv := value.List()
		clv := lv.(*_GenesisState_7_list)
		x.Proposals = *clv.list
	case "regen.group.v1alpha1.GenesisState.votes":
		lv := value.List()
		clv := lv.(*_GenesisState_8_list)
		x.Votes = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GenesisState"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GenesisState does not contain field %s", fd.FullName()))
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
func (x *fastReflection_GenesisState) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GenesisState.groups":
		if x.Groups == nil {
			x.Groups = []*GroupInfo{}
		}
		value := &_GenesisState_2_list{list: &x.Groups}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.GenesisState.group_members":
		if x.GroupMembers == nil {
			x.GroupMembers = []*GroupMember{}
		}
		value := &_GenesisState_3_list{list: &x.GroupMembers}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.GenesisState.group_accounts":
		if x.GroupAccounts == nil {
			x.GroupAccounts = []*GroupAccountInfo{}
		}
		value := &_GenesisState_5_list{list: &x.GroupAccounts}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.GenesisState.proposals":
		if x.Proposals == nil {
			x.Proposals = []*Proposal{}
		}
		value := &_GenesisState_7_list{list: &x.Proposals}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.GenesisState.votes":
		if x.Votes == nil {
			x.Votes = []*Vote{}
		}
		value := &_GenesisState_8_list{list: &x.Votes}
		return protoreflect.ValueOfList(value)
	case "regen.group.v1alpha1.GenesisState.group_seq":
		panic(fmt.Errorf("field group_seq of message regen.group.v1alpha1.GenesisState is not mutable"))
	case "regen.group.v1alpha1.GenesisState.group_account_seq":
		panic(fmt.Errorf("field group_account_seq of message regen.group.v1alpha1.GenesisState is not mutable"))
	case "regen.group.v1alpha1.GenesisState.proposal_seq":
		panic(fmt.Errorf("field proposal_seq of message regen.group.v1alpha1.GenesisState is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GenesisState"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_GenesisState) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.group.v1alpha1.GenesisState.group_seq":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.GenesisState.groups":
		list := []*GroupInfo{}
		return protoreflect.ValueOfList(&_GenesisState_2_list{list: &list})
	case "regen.group.v1alpha1.GenesisState.group_members":
		list := []*GroupMember{}
		return protoreflect.ValueOfList(&_GenesisState_3_list{list: &list})
	case "regen.group.v1alpha1.GenesisState.group_account_seq":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.GenesisState.group_accounts":
		list := []*GroupAccountInfo{}
		return protoreflect.ValueOfList(&_GenesisState_5_list{list: &list})
	case "regen.group.v1alpha1.GenesisState.proposal_seq":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.group.v1alpha1.GenesisState.proposals":
		list := []*Proposal{}
		return protoreflect.ValueOfList(&_GenesisState_7_list{list: &list})
	case "regen.group.v1alpha1.GenesisState.votes":
		list := []*Vote{}
		return protoreflect.ValueOfList(&_GenesisState_8_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.group.v1alpha1.GenesisState"))
		}
		panic(fmt.Errorf("message regen.group.v1alpha1.GenesisState does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_GenesisState) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.group.v1alpha1.GenesisState", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_GenesisState) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_GenesisState) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_GenesisState) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_GenesisState) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*GenesisState)
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
		if x.GroupSeq != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupSeq))
		}
		if len(x.Groups) > 0 {
			for _, e := range x.Groups {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.GroupMembers) > 0 {
			for _, e := range x.GroupMembers {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.GroupAccountSeq != 0 {
			n += 1 + runtime.Sov(uint64(x.GroupAccountSeq))
		}
		if len(x.GroupAccounts) > 0 {
			for _, e := range x.GroupAccounts {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.ProposalSeq != 0 {
			n += 1 + runtime.Sov(uint64(x.ProposalSeq))
		}
		if len(x.Proposals) > 0 {
			for _, e := range x.Proposals {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if len(x.Votes) > 0 {
			for _, e := range x.Votes {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*GenesisState)
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
		if len(x.Votes) > 0 {
			for iNdEx := len(x.Votes) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Votes[iNdEx])
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
				dAtA[i] = 0x42
			}
		}
		if len(x.Proposals) > 0 {
			for iNdEx := len(x.Proposals) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Proposals[iNdEx])
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
				dAtA[i] = 0x3a
			}
		}
		if x.ProposalSeq != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProposalSeq))
			i--
			dAtA[i] = 0x30
		}
		if len(x.GroupAccounts) > 0 {
			for iNdEx := len(x.GroupAccounts) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.GroupAccounts[iNdEx])
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
				dAtA[i] = 0x2a
			}
		}
		if x.GroupAccountSeq != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupAccountSeq))
			i--
			dAtA[i] = 0x20
		}
		if len(x.GroupMembers) > 0 {
			for iNdEx := len(x.GroupMembers) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.GroupMembers[iNdEx])
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
				dAtA[i] = 0x1a
			}
		}
		if len(x.Groups) > 0 {
			for iNdEx := len(x.Groups) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Groups[iNdEx])
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
		}
		if x.GroupSeq != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.GroupSeq))
			i--
			dAtA[i] = 0x8
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
		x := input.Message.Interface().(*GenesisState)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupSeq", wireType)
				}
				x.GroupSeq = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupSeq |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Groups", wireType)
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
				x.Groups = append(x.Groups, &GroupInfo{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Groups[len(x.Groups)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupMembers", wireType)
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
				x.GroupMembers = append(x.GroupMembers, &GroupMember{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.GroupMembers[len(x.GroupMembers)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupAccountSeq", wireType)
				}
				x.GroupAccountSeq = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.GroupAccountSeq |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field GroupAccounts", wireType)
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
				x.GroupAccounts = append(x.GroupAccounts, &GroupAccountInfo{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.GroupAccounts[len(x.GroupAccounts)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProposalSeq", wireType)
				}
				x.ProposalSeq = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProposalSeq |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
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
				x.Proposals = append(x.Proposals, &Proposal{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Proposals[len(x.Proposals)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 8:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
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
				x.Votes = append(x.Votes, &Vote{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Votes[len(x.Votes)-1]); err != nil {
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
// source: regen/group/v1alpha1/genesis.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GenesisState defines the group module's genesis state.
type GenesisState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// group_seq is the group table orm.Sequence,
	// it is used to get the next group ID.
	GroupSeq uint64 `protobuf:"varint,1,opt,name=group_seq,json=groupSeq,proto3" json:"group_seq,omitempty"`
	// groups is the list of groups info.
	Groups []*GroupInfo `protobuf:"bytes,2,rep,name=groups,proto3" json:"groups,omitempty"`
	// group_members is the list of groups members.
	GroupMembers []*GroupMember `protobuf:"bytes,3,rep,name=group_members,json=groupMembers,proto3" json:"group_members,omitempty"`
	// group_account_seq is the group account table orm.Sequence,
	// it is used to generate the next group account address.
	GroupAccountSeq uint64 `protobuf:"varint,4,opt,name=group_account_seq,json=groupAccountSeq,proto3" json:"group_account_seq,omitempty"`
	// group_accounts is the list of group accounts info.
	GroupAccounts []*GroupAccountInfo `protobuf:"bytes,5,rep,name=group_accounts,json=groupAccounts,proto3" json:"group_accounts,omitempty"`
	// proposal_seq is the proposal table orm.Sequence,
	// it is used to get the next proposal ID.
	ProposalSeq uint64 `protobuf:"varint,6,opt,name=proposal_seq,json=proposalSeq,proto3" json:"proposal_seq,omitempty"`
	// proposals is the list of proposals.
	Proposals []*Proposal `protobuf:"bytes,7,rep,name=proposals,proto3" json:"proposals,omitempty"`
	// votes is the list of votes.
	Votes []*Vote `protobuf:"bytes,8,rep,name=votes,proto3" json:"votes,omitempty"`
}

func (x *GenesisState) Reset() {
	*x = GenesisState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_group_v1alpha1_genesis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenesisState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenesisState) ProtoMessage() {}

// Deprecated: Use GenesisState.ProtoReflect.Descriptor instead.
func (*GenesisState) Descriptor() ([]byte, []int) {
	return file_regen_group_v1alpha1_genesis_proto_rawDescGZIP(), []int{0}
}

func (x *GenesisState) GetGroupSeq() uint64 {
	if x != nil {
		return x.GroupSeq
	}
	return 0
}

func (x *GenesisState) GetGroups() []*GroupInfo {
	if x != nil {
		return x.Groups
	}
	return nil
}

func (x *GenesisState) GetGroupMembers() []*GroupMember {
	if x != nil {
		return x.GroupMembers
	}
	return nil
}

func (x *GenesisState) GetGroupAccountSeq() uint64 {
	if x != nil {
		return x.GroupAccountSeq
	}
	return 0
}

func (x *GenesisState) GetGroupAccounts() []*GroupAccountInfo {
	if x != nil {
		return x.GroupAccounts
	}
	return nil
}

func (x *GenesisState) GetProposalSeq() uint64 {
	if x != nil {
		return x.ProposalSeq
	}
	return 0
}

func (x *GenesisState) GetProposals() []*Proposal {
	if x != nil {
		return x.Proposals
	}
	return nil
}

func (x *GenesisState) GetVotes() []*Vote {
	if x != nil {
		return x.Votes
	}
	return nil
}

var File_regen_group_v1alpha1_genesis_proto protoreflect.FileDescriptor

var file_regen_group_v1alpha1_genesis_proto_rawDesc = []byte{
	0x0a, 0x22, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x20, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xba, 0x03, 0x0a, 0x0c, 0x47, 0x65, 0x6e, 0x65, 0x73, 0x69, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x71,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x71,
	0x12, 0x37, 0x0a, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x12, 0x46, 0x0a, 0x0d, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x52, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x12, 0x2a, 0x0a, 0x11, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x71, 0x12, 0x4d, 0x0a,
	0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x65, 0x71, 0x12,
	0x3c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x73, 0x12, 0x30, 0x0a,
	0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x42,
	0xe8, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x0c, 0x47, 0x65,
	0x6e, 0x65, 0x73, 0x69, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6c,
	0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x52, 0x47, 0x58,
	0xaa, 0x02, 0x14, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x14, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02,
	0x20, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x16, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_regen_group_v1alpha1_genesis_proto_rawDescOnce sync.Once
	file_regen_group_v1alpha1_genesis_proto_rawDescData = file_regen_group_v1alpha1_genesis_proto_rawDesc
)

func file_regen_group_v1alpha1_genesis_proto_rawDescGZIP() []byte {
	file_regen_group_v1alpha1_genesis_proto_rawDescOnce.Do(func() {
		file_regen_group_v1alpha1_genesis_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_group_v1alpha1_genesis_proto_rawDescData)
	})
	return file_regen_group_v1alpha1_genesis_proto_rawDescData
}

var file_regen_group_v1alpha1_genesis_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_regen_group_v1alpha1_genesis_proto_goTypes = []interface{}{
	(*GenesisState)(nil),     // 0: regen.group.v1alpha1.GenesisState
	(*GroupInfo)(nil),        // 1: regen.group.v1alpha1.GroupInfo
	(*GroupMember)(nil),      // 2: regen.group.v1alpha1.GroupMember
	(*GroupAccountInfo)(nil), // 3: regen.group.v1alpha1.GroupAccountInfo
	(*Proposal)(nil),         // 4: regen.group.v1alpha1.Proposal
	(*Vote)(nil),             // 5: regen.group.v1alpha1.Vote
}
var file_regen_group_v1alpha1_genesis_proto_depIdxs = []int32{
	1, // 0: regen.group.v1alpha1.GenesisState.groups:type_name -> regen.group.v1alpha1.GroupInfo
	2, // 1: regen.group.v1alpha1.GenesisState.group_members:type_name -> regen.group.v1alpha1.GroupMember
	3, // 2: regen.group.v1alpha1.GenesisState.group_accounts:type_name -> regen.group.v1alpha1.GroupAccountInfo
	4, // 3: regen.group.v1alpha1.GenesisState.proposals:type_name -> regen.group.v1alpha1.Proposal
	5, // 4: regen.group.v1alpha1.GenesisState.votes:type_name -> regen.group.v1alpha1.Vote
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_regen_group_v1alpha1_genesis_proto_init() }
func file_regen_group_v1alpha1_genesis_proto_init() {
	if File_regen_group_v1alpha1_genesis_proto != nil {
		return
	}
	file_regen_group_v1alpha1_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_regen_group_v1alpha1_genesis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenesisState); i {
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
			RawDescriptor: file_regen_group_v1alpha1_genesis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_group_v1alpha1_genesis_proto_goTypes,
		DependencyIndexes: file_regen_group_v1alpha1_genesis_proto_depIdxs,
		MessageInfos:      file_regen_group_v1alpha1_genesis_proto_msgTypes,
	}.Build()
	File_regen_group_v1alpha1_genesis_proto = out.File
	file_regen_group_v1alpha1_genesis_proto_rawDesc = nil
	file_regen_group_v1alpha1_genesis_proto_goTypes = nil
	file_regen_group_v1alpha1_genesis_proto_depIdxs = nil
}
