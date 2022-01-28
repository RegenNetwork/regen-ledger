// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/data/v1alpha2/types.proto

package data

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

var MediaType_name = map[int32]string{
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

var MediaType_value = map[string]int32{
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

func (x MediaType) String() string {
	return proto.EnumName(MediaType_name, int32(x))
}

func (MediaType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e68eefb44eeab1df, []int{0}
}

// GraphCanonicalizationAlgorithm is the graph canonicalization algorithm
type GraphCanonicalizationAlgorithm int32

const (
	// unspecified and invalid
	GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED GraphCanonicalizationAlgorithm = 0
	// URDNA2015 graph hashing
	GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015 GraphCanonicalizationAlgorithm = 1
)

var GraphCanonicalizationAlgorithm_name = map[int32]string{
	0: "GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED",
	1: "GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015",
}

var GraphCanonicalizationAlgorithm_value = map[string]int32{
	"GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED": 0,
	"GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015":   1,
}

func (x GraphCanonicalizationAlgorithm) String() string {
	return proto.EnumName(GraphCanonicalizationAlgorithm_name, int32(x))
}

func (GraphCanonicalizationAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e68eefb44eeab1df, []int{1}
}

// GraphMerkleTree is the graph merkle tree type used for hashing, if any
type GraphMerkleTree int32

const (
	// no merkle tree
	GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED GraphMerkleTree = 0
)

var GraphMerkleTree_name = map[int32]string{
	0: "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED",
}

var GraphMerkleTree_value = map[string]int32{
	"GRAPH_MERKLE_TREE_NONE_UNSPECIFIED": 0,
}

func (x GraphMerkleTree) String() string {
	return proto.EnumName(GraphMerkleTree_name, int32(x))
}

func (GraphMerkleTree) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e68eefb44eeab1df, []int{2}
}

// DigestAlgorithm is the hash digest algorithm
type DigestAlgorithm int32

const (
	// unspecified and invalid
	DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED DigestAlgorithm = 0
	// BLAKE2b-256
	DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256 DigestAlgorithm = 1
)

var DigestAlgorithm_name = map[int32]string{
	0: "DIGEST_ALGORITHM_UNSPECIFIED",
	1: "DIGEST_ALGORITHM_BLAKE2B_256",
}

var DigestAlgorithm_value = map[string]int32{
	"DIGEST_ALGORITHM_UNSPECIFIED": 0,
	"DIGEST_ALGORITHM_BLAKE2B_256": 1,
}

func (x DigestAlgorithm) String() string {
	return proto.EnumName(DigestAlgorithm_name, int32(x))
}

func (DigestAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e68eefb44eeab1df, []int{3}
}

// ContentHash specifies a hash based content identifier for a piece of data
type ContentHash struct {
	// sum selects the type of content hash
	//
	// Types that are valid to be assigned to Sum:
	//	*ContentHash_Raw_
	//	*ContentHash_Graph_
	Sum isContentHash_Sum `protobuf_oneof:"sum"`
}

func (m *ContentHash) Reset()         { *m = ContentHash{} }
func (m *ContentHash) String() string { return proto.CompactTextString(m) }
func (*ContentHash) ProtoMessage()    {}
func (*ContentHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_e68eefb44eeab1df, []int{0}
}
func (m *ContentHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentHash.Merge(m, src)
}
func (m *ContentHash) XXX_Size() int {
	return m.Size()
}
func (m *ContentHash) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentHash.DiscardUnknown(m)
}

var xxx_messageInfo_ContentHash proto.InternalMessageInfo

type isContentHash_Sum interface {
	isContentHash_Sum()
	MarshalTo([]byte) (int, error)
	Size() int
}

type ContentHash_Raw_ struct {
	Raw *ContentHash_Raw `protobuf:"bytes,1,opt,name=raw,proto3,oneof" json:"raw,omitempty"`
}
type ContentHash_Graph_ struct {
	Graph *ContentHash_Graph `protobuf:"bytes,2,opt,name=graph,proto3,oneof" json:"graph,omitempty"`
}

func (*ContentHash_Raw_) isContentHash_Sum()   {}
func (*ContentHash_Graph_) isContentHash_Sum() {}

func (m *ContentHash) GetSum() isContentHash_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *ContentHash) GetRaw() *ContentHash_Raw {
	if x, ok := m.GetSum().(*ContentHash_Raw_); ok {
		return x.Raw
	}
	return nil
}

func (m *ContentHash) GetGraph() *ContentHash_Graph {
	if x, ok := m.GetSum().(*ContentHash_Graph_); ok {
		return x.Graph
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ContentHash) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ContentHash_Raw_)(nil),
		(*ContentHash_Graph_)(nil),
	}
}

// Raw is the content hash type used for raw data
type ContentHash_Raw struct {
	// hash represents the hash of the data based on the specified digest_algorithm
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// digest_algorithm represents the hash digest algorithm.
	DigestAlgorithm DigestAlgorithm `protobuf:"varint,2,opt,name=digest_algorithm,json=digestAlgorithm,proto3,enum=regen.data.v1alpha2.DigestAlgorithm" json:"digest_algorithm,omitempty"`
	// media_type represents the MediaType for raw data.
	MediaType MediaType `protobuf:"varint,3,opt,name=media_type,json=mediaType,proto3,enum=regen.data.v1alpha2.MediaType" json:"media_type,omitempty"`
}

func (m *ContentHash_Raw) Reset()         { *m = ContentHash_Raw{} }
func (m *ContentHash_Raw) String() string { return proto.CompactTextString(m) }
func (*ContentHash_Raw) ProtoMessage()    {}
func (*ContentHash_Raw) Descriptor() ([]byte, []int) {
	return fileDescriptor_e68eefb44eeab1df, []int{0, 0}
}
func (m *ContentHash_Raw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentHash_Raw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentHash_Raw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentHash_Raw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentHash_Raw.Merge(m, src)
}
func (m *ContentHash_Raw) XXX_Size() int {
	return m.Size()
}
func (m *ContentHash_Raw) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentHash_Raw.DiscardUnknown(m)
}

var xxx_messageInfo_ContentHash_Raw proto.InternalMessageInfo

func (m *ContentHash_Raw) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ContentHash_Raw) GetDigestAlgorithm() DigestAlgorithm {
	if m != nil {
		return m.DigestAlgorithm
	}
	return DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED
}

func (m *ContentHash_Raw) GetMediaType() MediaType {
	if m != nil {
		return m.MediaType
	}
	return MediaType_MEDIA_TYPE_UNSPECIFIED
}

// Graph is the content hash type used for RDF graph data
type ContentHash_Graph struct {
	// hash represents the hash of the data based on the specified digest_algorithm
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// digest_algorithm represents the hash digest algorithm.
	DigestAlgorithm DigestAlgorithm `protobuf:"varint,2,opt,name=digest_algorithm,json=digestAlgorithm,proto3,enum=regen.data.v1alpha2.DigestAlgorithm" json:"digest_algorithm,omitempty"`
	// graph_canonicalization_algorithm represents the RDF graph canonicalization algorithm.
	CanonicalizationAlgorithm GraphCanonicalizationAlgorithm `protobuf:"varint,3,opt,name=canonicalization_algorithm,json=canonicalizationAlgorithm,proto3,enum=regen.data.v1alpha2.GraphCanonicalizationAlgorithm" json:"canonicalization_algorithm,omitempty"`
	// merkle_tree is the merkle tree type used for the graph hash, if any
	MerkleTree GraphMerkleTree `protobuf:"varint,4,opt,name=merkle_tree,json=merkleTree,proto3,enum=regen.data.v1alpha2.GraphMerkleTree" json:"merkle_tree,omitempty"`
}

func (m *ContentHash_Graph) Reset()         { *m = ContentHash_Graph{} }
func (m *ContentHash_Graph) String() string { return proto.CompactTextString(m) }
func (*ContentHash_Graph) ProtoMessage()    {}
func (*ContentHash_Graph) Descriptor() ([]byte, []int) {
	return fileDescriptor_e68eefb44eeab1df, []int{0, 1}
}
func (m *ContentHash_Graph) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentHash_Graph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentHash_Graph.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentHash_Graph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentHash_Graph.Merge(m, src)
}
func (m *ContentHash_Graph) XXX_Size() int {
	return m.Size()
}
func (m *ContentHash_Graph) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentHash_Graph.DiscardUnknown(m)
}

var xxx_messageInfo_ContentHash_Graph proto.InternalMessageInfo

func (m *ContentHash_Graph) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ContentHash_Graph) GetDigestAlgorithm() DigestAlgorithm {
	if m != nil {
		return m.DigestAlgorithm
	}
	return DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED
}

func (m *ContentHash_Graph) GetCanonicalizationAlgorithm() GraphCanonicalizationAlgorithm {
	if m != nil {
		return m.CanonicalizationAlgorithm
	}
	return GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED
}

func (m *ContentHash_Graph) GetMerkleTree() GraphMerkleTree {
	if m != nil {
		return m.MerkleTree
	}
	return GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED
}

// SignerEntry is a signer entry wrapping a signer address and timestamp
type SignerEntry struct {
	// signer is the address of the signer
	Signer string `protobuf:"bytes,1,opt,name=signer,proto3" json:"signer,omitempty"`
	// timestamp is the time at which the data was signed
	Timestamp *types.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *SignerEntry) Reset()         { *m = SignerEntry{} }
func (m *SignerEntry) String() string { return proto.CompactTextString(m) }
func (*SignerEntry) ProtoMessage()    {}
func (*SignerEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_e68eefb44eeab1df, []int{1}
}
func (m *SignerEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignerEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignerEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignerEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignerEntry.Merge(m, src)
}
func (m *SignerEntry) XXX_Size() int {
	return m.Size()
}
func (m *SignerEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_SignerEntry.DiscardUnknown(m)
}

var xxx_messageInfo_SignerEntry proto.InternalMessageInfo

func (m *SignerEntry) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

func (m *SignerEntry) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func init() {
	proto.RegisterEnum("regen.data.v1alpha2.MediaType", MediaType_name, MediaType_value)
	proto.RegisterEnum("regen.data.v1alpha2.GraphCanonicalizationAlgorithm", GraphCanonicalizationAlgorithm_name, GraphCanonicalizationAlgorithm_value)
	proto.RegisterEnum("regen.data.v1alpha2.GraphMerkleTree", GraphMerkleTree_name, GraphMerkleTree_value)
	proto.RegisterEnum("regen.data.v1alpha2.DigestAlgorithm", DigestAlgorithm_name, DigestAlgorithm_value)
	proto.RegisterType((*ContentHash)(nil), "regen.data.v1alpha2.ContentHash")
	proto.RegisterType((*ContentHash_Raw)(nil), "regen.data.v1alpha2.ContentHash.Raw")
	proto.RegisterType((*ContentHash_Graph)(nil), "regen.data.v1alpha2.ContentHash.Graph")
	proto.RegisterType((*SignerEntry)(nil), "regen.data.v1alpha2.SignerEntry")
}

func init() { proto.RegisterFile("regen/data/v1alpha2/types.proto", fileDescriptor_e68eefb44eeab1df) }

var fileDescriptor_e68eefb44eeab1df = []byte{
	// 717 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x41, 0x4f, 0xdb, 0x48,
	0x18, 0x86, 0x63, 0x92, 0x20, 0x65, 0xb2, 0x22, 0xb3, 0xc3, 0xc2, 0x86, 0x68, 0x65, 0xd8, 0xec,
	0x0a, 0x55, 0x11, 0x75, 0x20, 0x94, 0x8a, 0x1e, 0x5a, 0xc9, 0x49, 0x1c, 0xc7, 0x10, 0x3b, 0xd6,
	0xc4, 0x50, 0xca, 0xc5, 0x1a, 0x92, 0xa9, 0x63, 0x11, 0xdb, 0x91, 0x63, 0x9a, 0xd2, 0x63, 0x6f,
	0xbd, 0xf5, 0x5f, 0x54, 0xfd, 0x27, 0x3d, 0x72, 0xec, 0xb1, 0x82, 0xfe, 0x90, 0x2a, 0x13, 0x02,
	0xe9, 0x34, 0xc0, 0xad, 0xb7, 0x99, 0x6f, 0x9e, 0xf7, 0xfd, 0x5e, 0x79, 0xbe, 0x31, 0x58, 0x0d,
	0xa9, 0x43, 0xfd, 0x62, 0x87, 0x44, 0xa4, 0xf8, 0x66, 0x8b, 0xf4, 0xfa, 0x5d, 0x52, 0x2a, 0x46,
	0xe7, 0x7d, 0x3a, 0x90, 0xfa, 0x61, 0x10, 0x05, 0x68, 0x91, 0x01, 0xd2, 0x08, 0x90, 0x26, 0x40,
	0x6e, 0xd5, 0x09, 0x02, 0xa7, 0x47, 0x8b, 0x0c, 0x39, 0x39, 0x7b, 0x5d, 0x8c, 0x5c, 0x8f, 0x0e,
	0x22, 0xe2, 0xf5, 0xc7, 0xaa, 0xfc, 0xf7, 0x04, 0x48, 0x57, 0x02, 0x3f, 0xa2, 0x7e, 0x54, 0x27,
	0x83, 0x2e, 0xda, 0x05, 0xf1, 0x90, 0x0c, 0xb3, 0xc2, 0x9a, 0xf0, 0x28, 0x5d, 0xfa, 0x5f, 0x9a,
	0xe1, 0x29, 0x4d, 0xe1, 0x12, 0x26, 0xc3, 0x7a, 0x0c, 0x8f, 0x24, 0xe8, 0x05, 0x48, 0x3a, 0x21,
	0xe9, 0x77, 0xb3, 0x73, 0x4c, 0xbb, 0xfe, 0xa0, 0x56, 0x1d, 0xd1, 0xf5, 0x18, 0x1e, 0xcb, 0x72,
	0x9f, 0x05, 0x10, 0xc7, 0x64, 0x88, 0x10, 0x48, 0x74, 0xc9, 0xa0, 0xcb, 0x22, 0xfc, 0x81, 0xd9,
	0x1a, 0x35, 0x01, 0xec, 0xb8, 0x0e, 0x1d, 0x44, 0x36, 0xe9, 0x39, 0x41, 0xe8, 0x46, 0x5d, 0x8f,
	0xb5, 0x59, 0xb8, 0x23, 0x62, 0x95, 0xc1, 0xf2, 0x84, 0xc5, 0x99, 0xce, 0xcf, 0x05, 0xf4, 0x1c,
	0x00, 0x8f, 0x76, 0x5c, 0x62, 0x8f, 0xbe, 0x60, 0x36, 0xce, 0xac, 0xc4, 0x99, 0x56, 0xfa, 0x08,
	0xb3, 0xce, 0xfb, 0x14, 0xa7, 0xbc, 0xc9, 0x32, 0xf7, 0x69, 0x0e, 0x24, 0x59, 0xfc, 0xdf, 0x93,
	0x36, 0x04, 0xb9, 0x36, 0xf1, 0x03, 0xdf, 0x6d, 0x93, 0x9e, 0xfb, 0x8e, 0x44, 0x6e, 0xe0, 0x4f,
	0x59, 0x8f, 0xd3, 0x6f, 0xcf, 0xb4, 0x66, 0x21, 0x2b, 0x9c, 0xf6, 0xb6, 0xd3, 0x4a, 0xfb, 0xae,
	0x23, 0xa4, 0x80, 0xb4, 0x47, 0xc3, 0xd3, 0x1e, 0xb5, 0xa3, 0x90, 0xd2, 0x6c, 0xe2, 0x9e, 0xfc,
	0xac, 0x89, 0xce, 0x60, 0x2b, 0xa4, 0x14, 0x03, 0xef, 0x66, 0x5d, 0x4e, 0x82, 0xf8, 0xe0, 0xcc,
	0xcb, 0xdb, 0x20, 0xdd, 0x72, 0x1d, 0x9f, 0x86, 0x8a, 0x1f, 0x85, 0xe7, 0x68, 0x19, 0xcc, 0x0f,
	0xd8, 0x96, 0x7d, 0xb7, 0x14, 0xbe, 0xde, 0xa1, 0x5d, 0x90, 0xba, 0x19, 0xd0, 0xeb, 0x39, 0xca,
	0x49, 0xe3, 0x11, 0x96, 0x26, 0x23, 0x2c, 0x59, 0x13, 0x02, 0xdf, 0xc2, 0x85, 0x0f, 0x71, 0x90,
	0xba, 0xb9, 0x2a, 0x94, 0x03, 0xcb, 0xba, 0x52, 0xd5, 0x64, 0xdb, 0x7a, 0x65, 0x2a, 0xf6, 0x81,
	0xd1, 0x32, 0x95, 0x8a, 0x56, 0xd3, 0x94, 0x2a, 0x8c, 0xa1, 0x15, 0xb0, 0x34, 0x75, 0x66, 0x29,
	0x47, 0x96, 0x6d, 0x36, 0x64, 0xcd, 0x80, 0x02, 0x5a, 0x04, 0x99, 0xa9, 0xa3, 0xbd, 0x56, 0xd3,
	0x80, 0x73, 0x08, 0x81, 0x85, 0xa9, 0x62, 0xa5, 0x75, 0x08, 0xe3, 0x5c, 0xed, 0x48, 0x6f, 0xc0,
	0x04, 0x57, 0x33, 0xab, 0x35, 0x98, 0xe4, 0x0c, 0x2d, 0xad, 0x56, 0x83, 0x90, 0x03, 0xf7, 0x4c,
	0x15, 0xfe, 0xc9, 0x8b, 0x0d, 0x15, 0x22, 0xae, 0xd6, 0x3a, 0x54, 0xe1, 0x22, 0x67, 0xf8, 0x52,
	0x29, 0x9b, 0xf0, 0x2f, 0xae, 0x28, 0x1f, 0x6a, 0x35, 0xb8, 0xc4, 0xa9, 0x55, 0xad, 0x06, 0x97,
	0x79, 0x70, 0xd4, 0xe6, 0x6f, 0xae, 0xa8, 0x9b, 0x8a, 0x0a, 0xd7, 0x38, 0xb5, 0x6e, 0x3e, 0x81,
	0xff, 0xfe, 0xda, 0x5b, 0x87, 0x79, 0x0e, 0x6c, 0xaa, 0x2a, 0xfc, 0xaf, 0xf0, 0x5e, 0x00, 0xe2,
	0xfd, 0x83, 0x87, 0x36, 0xc1, 0x86, 0x8a, 0x65, 0xb3, 0x6e, 0x57, 0x64, 0xa3, 0x69, 0x68, 0x15,
	0xb9, 0xa1, 0x1d, 0xcb, 0x96, 0xd6, 0x34, 0x6c, 0xb9, 0xa1, 0x36, 0xb1, 0x66, 0xd5, 0x75, 0xee,
	0xda, 0x24, 0x50, 0x78, 0x58, 0x81, 0xab, 0x86, 0x5c, 0xda, 0xdc, 0xda, 0x81, 0x42, 0xe1, 0x19,
	0xc8, 0x70, 0x73, 0x89, 0xd6, 0x41, 0x7e, 0x6c, 0xa1, 0x2b, 0x78, 0xbf, 0xa1, 0xd8, 0x16, 0x56,
	0x14, 0xdb, 0x68, 0x1a, 0xdc, 0x84, 0x14, 0x0e, 0x40, 0x86, 0x7b, 0x92, 0x68, 0x0d, 0xfc, 0x53,
	0xd5, 0x54, 0xa5, 0x65, 0xdd, 0x99, 0x6f, 0x16, 0x51, 0x6e, 0xc8, 0xfb, 0x4a, 0xa9, 0x6c, 0x97,
	0x76, 0x9e, 0x42, 0xa1, 0x5c, 0xfb, 0x72, 0x29, 0x0a, 0x17, 0x97, 0xa2, 0xf0, 0xed, 0x52, 0x14,
	0x3e, 0x5e, 0x89, 0xb1, 0x8b, 0x2b, 0x31, 0xf6, 0xf5, 0x4a, 0x8c, 0x1d, 0x6f, 0x38, 0x6e, 0xd4,
	0x3d, 0x3b, 0x91, 0xda, 0x81, 0x57, 0x64, 0x0f, 0xec, 0xb1, 0x4f, 0xa3, 0x61, 0x10, 0x9e, 0x5e,
	0xef, 0x7a, 0xb4, 0xe3, 0xd0, 0xb0, 0xf8, 0x96, 0xfd, 0xfd, 0x4f, 0xe6, 0xd9, 0x4b, 0xd8, 0xfe,
	0x11, 0x00, 0x00, 0xff, 0xff, 0x17, 0xca, 0xb8, 0x37, 0x12, 0x06, 0x00, 0x00,
}

func (m *ContentHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentHash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentHash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *ContentHash_Raw_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentHash_Raw_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Raw != nil {
		{
			size, err := m.Raw.MarshalToSizedBuffer(dAtA[:i])
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
func (m *ContentHash_Graph_) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentHash_Graph_) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.Graph != nil {
		{
			size, err := m.Graph.MarshalToSizedBuffer(dAtA[:i])
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
func (m *ContentHash_Raw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentHash_Raw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentHash_Raw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MediaType != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MediaType))
		i--
		dAtA[i] = 0x18
	}
	if m.DigestAlgorithm != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.DigestAlgorithm))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContentHash_Graph) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentHash_Graph) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentHash_Graph) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MerkleTree != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MerkleTree))
		i--
		dAtA[i] = 0x20
	}
	if m.CanonicalizationAlgorithm != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.CanonicalizationAlgorithm))
		i--
		dAtA[i] = 0x18
	}
	if m.DigestAlgorithm != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.DigestAlgorithm))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *SignerEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignerEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignerEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Signer)))
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
func (m *ContentHash) Size() (n int) {
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

func (m *ContentHash_Raw_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Raw != nil {
		l = m.Raw.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *ContentHash_Graph_) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Graph != nil {
		l = m.Graph.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *ContentHash_Raw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.DigestAlgorithm != 0 {
		n += 1 + sovTypes(uint64(m.DigestAlgorithm))
	}
	if m.MediaType != 0 {
		n += 1 + sovTypes(uint64(m.MediaType))
	}
	return n
}

func (m *ContentHash_Graph) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.DigestAlgorithm != 0 {
		n += 1 + sovTypes(uint64(m.DigestAlgorithm))
	}
	if m.CanonicalizationAlgorithm != 0 {
		n += 1 + sovTypes(uint64(m.CanonicalizationAlgorithm))
	}
	if m.MerkleTree != 0 {
		n += 1 + sovTypes(uint64(m.MerkleTree))
	}
	return n
}

func (m *SignerEntry) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
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
func (m *ContentHash) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ContentHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
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
			v := &ContentHash_Raw{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ContentHash_Raw_{v}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Graph", wireType)
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
			v := &ContentHash_Graph{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &ContentHash_Graph_{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
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
func (m *ContentHash_Raw) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Raw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Raw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DigestAlgorithm", wireType)
			}
			m.DigestAlgorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DigestAlgorithm |= DigestAlgorithm(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MediaType", wireType)
			}
			m.MediaType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MediaType |= MediaType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
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
func (m *ContentHash_Graph) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Graph: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Graph: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DigestAlgorithm", wireType)
			}
			m.DigestAlgorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DigestAlgorithm |= DigestAlgorithm(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanonicalizationAlgorithm", wireType)
			}
			m.CanonicalizationAlgorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CanonicalizationAlgorithm |= GraphCanonicalizationAlgorithm(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkleTree", wireType)
			}
			m.MerkleTree = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MerkleTree |= GraphMerkleTree(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
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
func (m *SignerEntry) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: SignerEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignerEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
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
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
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
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
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
