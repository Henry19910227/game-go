// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: proto_data/TestAssign.proto

package req

import (
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

// 定義 TestAssign 結構
type TestAssign struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 小遊戲 id
	MiniGameId int32 `protobuf:"varint,1,opt,name=miniGameId,proto3" json:"miniGameId,omitempty"`
	// 牌面，使用 repeated 表示數組
	ElementsArray []int32 `protobuf:"varint,2,rep,packed,name=elementsArray,proto3" json:"elementsArray,omitempty"`
}

func (x *TestAssign) Reset() {
	*x = TestAssign{}
	mi := &file_proto_data_TestAssign_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestAssign) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestAssign) ProtoMessage() {}

func (x *TestAssign) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_TestAssign_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestAssign.ProtoReflect.Descriptor instead.
func (*TestAssign) Descriptor() ([]byte, []int) {
	return file_proto_data_TestAssign_proto_rawDescGZIP(), []int{0}
}

func (x *TestAssign) GetMiniGameId() int32 {
	if x != nil {
		return x.MiniGameId
	}
	return 0
}

func (x *TestAssign) GetElementsArray() []int32 {
	if x != nil {
		return x.ElementsArray
	}
	return nil
}

var File_proto_data_TestAssign_proto protoreflect.FileDescriptor

var file_proto_data_TestAssign_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x54, 0x65, 0x73,
	0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x52, 0x0a,
	0x0a, 0x54, 0x65, 0x73, 0x74, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x6d,
	0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x6d, 0x69, 0x6e, 0x69, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x0d, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x41, 0x72, 0x72, 0x61,
	0x79, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x72, 0x65,
	0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_data_TestAssign_proto_rawDescOnce sync.Once
	file_proto_data_TestAssign_proto_rawDescData = file_proto_data_TestAssign_proto_rawDesc
)

func file_proto_data_TestAssign_proto_rawDescGZIP() []byte {
	file_proto_data_TestAssign_proto_rawDescOnce.Do(func() {
		file_proto_data_TestAssign_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_data_TestAssign_proto_rawDescData)
	})
	return file_proto_data_TestAssign_proto_rawDescData
}

var file_proto_data_TestAssign_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_data_TestAssign_proto_goTypes = []any{
	(*TestAssign)(nil), // 0: TestAssign
}
var file_proto_data_TestAssign_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_data_TestAssign_proto_init() }
func file_proto_data_TestAssign_proto_init() {
	if File_proto_data_TestAssign_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_data_TestAssign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_data_TestAssign_proto_goTypes,
		DependencyIndexes: file_proto_data_TestAssign_proto_depIdxs,
		MessageInfos:      file_proto_data_TestAssign_proto_msgTypes,
	}.Build()
	File_proto_data_TestAssign_proto = out.File
	file_proto_data_TestAssign_proto_rawDesc = nil
	file_proto_data_TestAssign_proto_goTypes = nil
	file_proto_data_TestAssign_proto_depIdxs = nil
}
