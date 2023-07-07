// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: paddle_ocr.proto

package paddle_ocr

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

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"` // 需要注意的是，图片数据使用 bytes 类型进行传输。在返回图片数据时，通常需要将其转换为 Base64 编码后再传输。
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paddle_ocr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_paddle_ocr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_paddle_ocr_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

type ImageText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Image []byte `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
	Text  string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ImageText) Reset() {
	*x = ImageText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paddle_ocr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageText) ProtoMessage() {}

func (x *ImageText) ProtoReflect() protoreflect.Message {
	mi := &file_paddle_ocr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageText.ProtoReflect.Descriptor instead.
func (*ImageText) Descriptor() ([]byte, []int) {
	return file_paddle_ocr_proto_rawDescGZIP(), []int{1}
}

func (x *ImageText) GetImage() []byte {
	if x != nil {
		return x.Image
	}
	return nil
}

func (x *ImageText) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_paddle_ocr_proto protoreflect.FileDescriptor

var file_paddle_ocr_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6f, 0x63, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6f, 0x63, 0x72, 0x22, 0x1d,
	0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x35, 0x0a,
	0x09, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x32, 0x49, 0x0a, 0x10, 0x50, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x4f, 0x63,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x09, 0x50, 0x61, 0x64, 0x64,
	0x6c, 0x65, 0x4f, 0x63, 0x72, 0x12, 0x11, 0x2e, 0x70, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6f,
	0x63, 0x72, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x1a, 0x15, 0x2e, 0x70, 0x61, 0x64, 0x64, 0x6c,
	0x65, 0x5f, 0x6f, 0x63, 0x72, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x65, 0x78, 0x74, 0x42,
	0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70, 0x61, 0x64, 0x64, 0x6c, 0x65, 0x5f, 0x6f, 0x63, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paddle_ocr_proto_rawDescOnce sync.Once
	file_paddle_ocr_proto_rawDescData = file_paddle_ocr_proto_rawDesc
)

func file_paddle_ocr_proto_rawDescGZIP() []byte {
	file_paddle_ocr_proto_rawDescOnce.Do(func() {
		file_paddle_ocr_proto_rawDescData = protoimpl.X.CompressGZIP(file_paddle_ocr_proto_rawDescData)
	})
	return file_paddle_ocr_proto_rawDescData
}

var file_paddle_ocr_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_paddle_ocr_proto_goTypes = []interface{}{
	(*Image)(nil),     // 0: paddle_ocr.Image
	(*ImageText)(nil), // 1: paddle_ocr.ImageText
}
var file_paddle_ocr_proto_depIdxs = []int32{
	0, // 0: paddle_ocr.PaddleOcrService.PaddleOcr:input_type -> paddle_ocr.Image
	1, // 1: paddle_ocr.PaddleOcrService.PaddleOcr:output_type -> paddle_ocr.ImageText
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_paddle_ocr_proto_init() }
func file_paddle_ocr_proto_init() {
	if File_paddle_ocr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paddle_ocr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_paddle_ocr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageText); i {
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
			RawDescriptor: file_paddle_ocr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paddle_ocr_proto_goTypes,
		DependencyIndexes: file_paddle_ocr_proto_depIdxs,
		MessageInfos:      file_paddle_ocr_proto_msgTypes,
	}.Build()
	File_paddle_ocr_proto = out.File
	file_paddle_ocr_proto_rawDesc = nil
	file_paddle_ocr_proto_goTypes = nil
	file_paddle_ocr_proto_depIdxs = nil
}
