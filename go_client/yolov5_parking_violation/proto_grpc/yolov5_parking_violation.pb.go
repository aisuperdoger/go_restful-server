// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: yolov5_parking_violation.proto

package yolov5_parking_violation

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

type VideoPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *VideoPath) Reset() {
	*x = VideoPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yolov5_parking_violation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPath) ProtoMessage() {}

func (x *VideoPath) ProtoReflect() protoreflect.Message {
	mi := &file_yolov5_parking_violation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPath.ProtoReflect.Descriptor instead.
func (*VideoPath) Descriptor() ([]byte, []int) {
	return file_yolov5_parking_violation_proto_rawDescGZIP(), []int{0}
}

func (x *VideoPath) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

var File_yolov5_parking_violation_proto protoreflect.FileDescriptor

var file_yolov5_parking_violation_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x79, 0x6f, 0x6c, 0x6f, 0x76, 0x35, 0x5f, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x18, 0x79, 0x6f, 0x6c, 0x6f, 0x76, 0x35, 0x5f, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1f, 0x0a, 0x09, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x32, 0x83, 0x01, 0x0a, 0x1d,
	0x59, 0x6f, 0x6c, 0x6f, 0x76, 0x35, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x56, 0x69, 0x6f,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x62, 0x0a,
	0x16, 0x79, 0x6f, 0x6c, 0x6f, 0x76, 0x35, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x56, 0x69,
	0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x79, 0x6f, 0x6c, 0x6f, 0x76, 0x35,
	0x5f, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x23, 0x2e, 0x79,
	0x6f, 0x6c, 0x6f, 0x76, 0x35, 0x5f, 0x70, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x69,
	0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x61, 0x74,
	0x68, 0x42, 0x1c, 0x5a, 0x1a, 0x2e, 0x2f, 0x79, 0x6f, 0x6c, 0x6f, 0x76, 0x35, 0x5f, 0x70, 0x61,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x69, 0x6f, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yolov5_parking_violation_proto_rawDescOnce sync.Once
	file_yolov5_parking_violation_proto_rawDescData = file_yolov5_parking_violation_proto_rawDesc
)

func file_yolov5_parking_violation_proto_rawDescGZIP() []byte {
	file_yolov5_parking_violation_proto_rawDescOnce.Do(func() {
		file_yolov5_parking_violation_proto_rawDescData = protoimpl.X.CompressGZIP(file_yolov5_parking_violation_proto_rawDescData)
	})
	return file_yolov5_parking_violation_proto_rawDescData
}

var file_yolov5_parking_violation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yolov5_parking_violation_proto_goTypes = []interface{}{
	(*VideoPath)(nil), // 0: yolov5_parking_violation.VideoPath
}
var file_yolov5_parking_violation_proto_depIdxs = []int32{
	0, // 0: yolov5_parking_violation.Yolov5ParkingViolationService.yolov5ParkingViolation:input_type -> yolov5_parking_violation.VideoPath
	0, // 1: yolov5_parking_violation.Yolov5ParkingViolationService.yolov5ParkingViolation:output_type -> yolov5_parking_violation.VideoPath
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yolov5_parking_violation_proto_init() }
func file_yolov5_parking_violation_proto_init() {
	if File_yolov5_parking_violation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yolov5_parking_violation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoPath); i {
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
			RawDescriptor: file_yolov5_parking_violation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yolov5_parking_violation_proto_goTypes,
		DependencyIndexes: file_yolov5_parking_violation_proto_depIdxs,
		MessageInfos:      file_yolov5_parking_violation_proto_msgTypes,
	}.Build()
	File_yolov5_parking_violation_proto = out.File
	file_yolov5_parking_violation_proto_rawDesc = nil
	file_yolov5_parking_violation_proto_goTypes = nil
	file_yolov5_parking_violation_proto_depIdxs = nil
}
