// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: svc.biz.room/room.proto

package room

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InitDBResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result bool `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *InitDBResp) Reset() {
	*x = InitDBResp{}
	mi := &file_svc_biz_room_room_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitDBResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitDBResp) ProtoMessage() {}

func (x *InitDBResp) ProtoReflect() protoreflect.Message {
	mi := &file_svc_biz_room_room_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitDBResp.ProtoReflect.Descriptor instead.
func (*InitDBResp) Descriptor() ([]byte, []int) {
	return file_svc_biz_room_room_proto_rawDescGZIP(), []int{0}
}

func (x *InitDBResp) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_svc_biz_room_room_proto protoreflect.FileDescriptor

var file_svc_biz_room_room_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x2f, 0x72,
	0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x76, 0x63, 0x2e, 0x62,
	0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x24, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x42, 0x0a, 0x04, 0x52, 0x6f,
	0x6f, 0x6d, 0x12, 0x3a, 0x0a, 0x06, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72,
	0x6f, 0x6f, 0x6d, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x42, 0x52, 0x65, 0x73, 0x70, 0x42, 0x15,
	0x5a, 0x13, 0x2e, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x62, 0x69, 0x7a, 0x2e, 0x72, 0x6f, 0x6f, 0x6d,
	0x3b, 0x72, 0x6f, 0x6f, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_biz_room_room_proto_rawDescOnce sync.Once
	file_svc_biz_room_room_proto_rawDescData = file_svc_biz_room_room_proto_rawDesc
)

func file_svc_biz_room_room_proto_rawDescGZIP() []byte {
	file_svc_biz_room_room_proto_rawDescOnce.Do(func() {
		file_svc_biz_room_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_biz_room_room_proto_rawDescData)
	})
	return file_svc_biz_room_room_proto_rawDescData
}

var file_svc_biz_room_room_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_svc_biz_room_room_proto_goTypes = []any{
	(*InitDBResp)(nil),    // 0: svc.biz.room.InitDBResp
	(*emptypb.Empty)(nil), // 1: google.protobuf.Empty
}
var file_svc_biz_room_room_proto_depIdxs = []int32{
	1, // 0: svc.biz.room.Room.InitDB:input_type -> google.protobuf.Empty
	0, // 1: svc.biz.room.Room.InitDB:output_type -> svc.biz.room.InitDBResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_svc_biz_room_room_proto_init() }
func file_svc_biz_room_room_proto_init() {
	if File_svc_biz_room_room_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_biz_room_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_biz_room_room_proto_goTypes,
		DependencyIndexes: file_svc_biz_room_room_proto_depIdxs,
		MessageInfos:      file_svc_biz_room_room_proto_msgTypes,
	}.Build()
	File_svc_biz_room_room_proto = out.File
	file_svc_biz_room_room_proto_rawDesc = nil
	file_svc_biz_room_room_proto_goTypes = nil
	file_svc_biz_room_room_proto_depIdxs = nil
}
