// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: kv.proto

package gen

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// values from https://github.com/ledgerwatch/interfaces/blob/master/remote/kv.proto#L68
type Op int32

const (
	Op_FIRST      Op = 0
	Op_SEEK       Op = 1
	Op_CURRENT    Op = 4
	Op_NEXT       Op = 8
	Op_SEEK_EXACT Op = 15
	Op_OPEN       Op = 30
	Op_CLOSE      Op = 31
	Op_GET        Op = 64
)

// Enum value maps for Op.
var (
	Op_name = map[int32]string{
		0:  "FIRST",
		1:  "SEEK",
		4:  "CURRENT",
		8:  "NEXT",
		15: "SEEK_EXACT",
		30: "OPEN",
		31: "CLOSE",
		64: "GET",
	}
	Op_value = map[string]int32{
		"FIRST":      0,
		"SEEK":       1,
		"CURRENT":    4,
		"NEXT":       8,
		"SEEK_EXACT": 15,
		"OPEN":       30,
		"CLOSE":      31,
		"GET":        64,
	}
)

func (x Op) Enum() *Op {
	p := new(Op)
	*p = x
	return p
}

func (x Op) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Op) Descriptor() protoreflect.EnumDescriptor {
	return file_kv_proto_enumTypes[0].Descriptor()
}

func (Op) Type() protoreflect.EnumType {
	return &file_kv_proto_enumTypes[0]
}

func (x Op) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Op.Descriptor instead.
func (Op) EnumDescriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{0}
}

type Cursor struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op         Op     `protobuf:"varint,1,opt,name=op,proto3,enum=database.Op" json:"op,omitempty"`
	BucketName []byte `protobuf:"bytes,2,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	Cursor     uint32 `protobuf:"varint,3,opt,name=cursor,proto3" json:"cursor,omitempty"`
	K          []byte `protobuf:"bytes,4,opt,name=k,proto3" json:"k,omitempty"`
	V          []byte `protobuf:"bytes,5,opt,name=v,proto3" json:"v,omitempty"` // not used
}

func (x *Cursor) Reset() {
	*x = Cursor{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cursor) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cursor) ProtoMessage() {}

func (x *Cursor) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cursor.ProtoReflect.Descriptor instead.
func (*Cursor) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{0}
}

func (x *Cursor) GetOp() Op {
	if x != nil {
		return x.Op
	}
	return Op_FIRST
}

func (x *Cursor) GetBucketName() []byte {
	if x != nil {
		return x.BucketName
	}
	return nil
}

func (x *Cursor) GetCursor() uint32 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *Cursor) GetK() []byte {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *Cursor) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	K        []byte `protobuf:"bytes,1,opt,name=k,proto3" json:"k,omitempty"`
	V        []byte `protobuf:"bytes,2,opt,name=v,proto3" json:"v,omitempty"`
	CursorId uint32 `protobuf:"varint,3,opt,name=cursor_id,json=cursorId,proto3" json:"cursor_id,omitempty"`
	ViewId   uint64 `protobuf:"varint,4,opt,name=view_id,json=viewId,proto3" json:"view_id,omitempty"` // not used
	TxId     uint64 `protobuf:"varint,5,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`       // not used
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{1}
}

func (x *Pair) GetK() []byte {
	if x != nil {
		return x.K
	}
	return nil
}

func (x *Pair) GetV() []byte {
	if x != nil {
		return x.V
	}
	return nil
}

func (x *Pair) GetCursorId() uint32 {
	if x != nil {
		return x.CursorId
	}
	return 0
}

func (x *Pair) GetViewId() uint64 {
	if x != nil {
		return x.ViewId
	}
	return 0
}

func (x *Pair) GetTxId() uint64 {
	if x != nil {
		return x.TxId
	}
	return 0
}

type VersionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major uint32 `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor uint32 `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch uint32 `protobuf:"varint,3,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (x *VersionReply) Reset() {
	*x = VersionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_kv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VersionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionReply) ProtoMessage() {}

func (x *VersionReply) ProtoReflect() protoreflect.Message {
	mi := &file_kv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionReply.ProtoReflect.Descriptor instead.
func (*VersionReply) Descriptor() ([]byte, []int) {
	return file_kv_proto_rawDescGZIP(), []int{2}
}

func (x *VersionReply) GetMajor() uint32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *VersionReply) GetMinor() uint32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *VersionReply) GetPatch() uint32 {
	if x != nil {
		return x.Patch
	}
	return 0
}

var File_kv_proto protoreflect.FileDescriptor

var file_kv_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6b, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7b, 0x0a, 0x06, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x02, 0x6f,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x4f, 0x70, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75, 0x63,
	0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x75,
	0x72, 0x73, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x6b,
	0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x76, 0x22, 0x6d,
	0x0a, 0x04, 0x50, 0x61, 0x69, 0x72, 0x12, 0x0c, 0x0a, 0x01, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x01, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x01, 0x76, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x76, 0x69, 0x65, 0x77, 0x49, 0x64, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x78, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x74, 0x78, 0x49, 0x64, 0x22, 0x50, 0x0a,
	0x0c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6d, 0x61,
	0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2a,
	0x5e, 0x0a, 0x02, 0x4f, 0x70, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x49, 0x52, 0x53, 0x54, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x53, 0x45, 0x45, 0x4b, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x55,
	0x52, 0x52, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45, 0x58, 0x54, 0x10,
	0x08, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x45, 0x4b, 0x5f, 0x45, 0x58, 0x41, 0x43, 0x54, 0x10,
	0x0f, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x50, 0x45, 0x4e, 0x10, 0x1e, 0x12, 0x09, 0x0a, 0x05, 0x43,
	0x4c, 0x4f, 0x53, 0x45, 0x10, 0x1f, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x40, 0x32,
	0x6b, 0x0a, 0x02, 0x4b, 0x56, 0x12, 0x39, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x2a, 0x0a, 0x02, 0x54, 0x78, 0x12, 0x10, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x1a, 0x0e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x28, 0x01, 0x30, 0x01, 0x42, 0x1a, 0x5a, 0x18,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x75, 0x6e, 0x6f, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_kv_proto_rawDescOnce sync.Once
	file_kv_proto_rawDescData = file_kv_proto_rawDesc
)

func file_kv_proto_rawDescGZIP() []byte {
	file_kv_proto_rawDescOnce.Do(func() {
		file_kv_proto_rawDescData = protoimpl.X.CompressGZIP(file_kv_proto_rawDescData)
	})
	return file_kv_proto_rawDescData
}

var file_kv_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_kv_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_kv_proto_goTypes = []interface{}{
	(Op)(0),              // 0: database.Op
	(*Cursor)(nil),       // 1: database.Cursor
	(*Pair)(nil),         // 2: database.Pair
	(*VersionReply)(nil), // 3: database.VersionReply
	(*empty.Empty)(nil),  // 4: google.protobuf.Empty
}
var file_kv_proto_depIdxs = []int32{
	0, // 0: database.Cursor.op:type_name -> database.Op
	4, // 1: database.KV.Version:input_type -> google.protobuf.Empty
	1, // 2: database.KV.Tx:input_type -> database.Cursor
	3, // 3: database.KV.Version:output_type -> database.VersionReply
	2, // 4: database.KV.Tx:output_type -> database.Pair
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_kv_proto_init() }
func file_kv_proto_init() {
	if File_kv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_kv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cursor); i {
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
		file_kv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_kv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VersionReply); i {
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
			RawDescriptor: file_kv_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_kv_proto_goTypes,
		DependencyIndexes: file_kv_proto_depIdxs,
		EnumInfos:         file_kv_proto_enumTypes,
		MessageInfos:      file_kv_proto_msgTypes,
	}.Build()
	File_kv_proto = out.File
	file_kv_proto_rawDesc = nil
	file_kv_proto_goTypes = nil
	file_kv_proto_depIdxs = nil
}
