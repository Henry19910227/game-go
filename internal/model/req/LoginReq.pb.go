// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.0
// source: proto_data/LoginReq.proto

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

// 定義 LoginReq 消息
type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 代理名稱
	AgentName string `protobuf:"bytes,1,opt,name=agentName,proto3" json:"agentName,omitempty"`
	// 登錄 token，client 從大廳/商戶登錄獲取
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// 客戶端版本
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	// 熱更版本。僅供 cocos 連 logon 使用，其他不用傳
	Base string `protobuf:"bytes,4,opt,name=base,proto3" json:"base,omitempty"`
	// 小遊戲版本。僅供 cocos 連 logon 使用，其他不用傳
	GameVersion string `protobuf:"bytes,5,opt,name=gameVersion,proto3" json:"gameVersion,omitempty"`
	// 賬號平台：2=2n1（默認）、4=試玩、5=h5。cocos 可不傳
	Platform int32 `protobuf:"varint,6,opt,name=platform,proto3" json:"platform,omitempty"`
	// serverID。遊戲服合併端口後，cocos 用此字段識別進哪個服
	Server int32 `protobuf:"varint,7,opt,name=server,proto3" json:"server,omitempty"`
	// 用於識別重複請求的唯一 id。僅供 cocos 識別登錄重發使用，h5 不用傳
	RequestId int32 `protobuf:"varint,8,opt,name=requestId,proto3" json:"requestId,omitempty"`
	// 暱稱
	Nickname string `protobuf:"bytes,9,opt,name=nickname,proto3" json:"nickname,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	mi := &file_proto_data_LoginReq_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_data_LoginReq_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_proto_data_LoginReq_proto_rawDescGZIP(), []int{0}
}

func (x *LoginReq) GetAgentName() string {
	if x != nil {
		return x.AgentName
	}
	return ""
}

func (x *LoginReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginReq) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *LoginReq) GetBase() string {
	if x != nil {
		return x.Base
	}
	return ""
}

func (x *LoginReq) GetGameVersion() string {
	if x != nil {
		return x.GameVersion
	}
	return ""
}

func (x *LoginReq) GetPlatform() int32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *LoginReq) GetServer() int32 {
	if x != nil {
		return x.Server
	}
	return 0
}

func (x *LoginReq) GetRequestId() int32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *LoginReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

var File_proto_data_LoginReq_proto protoreflect.FileDescriptor

var file_proto_data_LoginReq_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x73, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x61,
	0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x67, 0x61, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x72,
	0x65, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_data_LoginReq_proto_rawDescOnce sync.Once
	file_proto_data_LoginReq_proto_rawDescData = file_proto_data_LoginReq_proto_rawDesc
)

func file_proto_data_LoginReq_proto_rawDescGZIP() []byte {
	file_proto_data_LoginReq_proto_rawDescOnce.Do(func() {
		file_proto_data_LoginReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_data_LoginReq_proto_rawDescData)
	})
	return file_proto_data_LoginReq_proto_rawDescData
}

var file_proto_data_LoginReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_data_LoginReq_proto_goTypes = []any{
	(*LoginReq)(nil), // 0: LoginReq
}
var file_proto_data_LoginReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_data_LoginReq_proto_init() }
func file_proto_data_LoginReq_proto_init() {
	if File_proto_data_LoginReq_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_data_LoginReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_data_LoginReq_proto_goTypes,
		DependencyIndexes: file_proto_data_LoginReq_proto_depIdxs,
		MessageInfos:      file_proto_data_LoginReq_proto_msgTypes,
	}.Build()
	File_proto_data_LoginReq_proto = out.File
	file_proto_data_LoginReq_proto_rawDesc = nil
	file_proto_data_LoginReq_proto_goTypes = nil
	file_proto_data_LoginReq_proto_depIdxs = nil
}
