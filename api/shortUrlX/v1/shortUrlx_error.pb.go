// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.0--rc1
// source: shortUrlX/v1/shortUrlx_error.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ShortUrlErrorReason int32

const (
	// 为某个枚举单独设置错误码
	ShortUrlErrorReason_ShortUrl_NotFound ShortUrlErrorReason = 0
)

// Enum value maps for ShortUrlErrorReason.
var (
	ShortUrlErrorReason_name = map[int32]string{
		0: "ShortUrl_NotFound",
	}
	ShortUrlErrorReason_value = map[string]int32{
		"ShortUrl_NotFound": 0,
	}
)

func (x ShortUrlErrorReason) Enum() *ShortUrlErrorReason {
	p := new(ShortUrlErrorReason)
	*p = x
	return p
}

func (x ShortUrlErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShortUrlErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_shortUrlX_v1_shortUrlx_error_proto_enumTypes[0].Descriptor()
}

func (ShortUrlErrorReason) Type() protoreflect.EnumType {
	return &file_shortUrlX_v1_shortUrlx_error_proto_enumTypes[0]
}

func (x ShortUrlErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShortUrlErrorReason.Descriptor instead.
func (ShortUrlErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_shortUrlX_v1_shortUrlx_error_proto_rawDescGZIP(), []int{0}
}

var File_shortUrlX_v1_shortUrlx_error_proto protoreflect.FileDescriptor

var file_shortUrlX_v1_shortUrlx_error_proto_rawDesc = []byte{
	0x0a, 0x22, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x58, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x78, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x37, 0x0a, 0x13, 0x53, 0x68, 0x6f,
	0x72, 0x74, 0x55, 0x72, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x11, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x5f, 0x4e, 0x6f, 0x74,
	0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x00, 0x1a, 0x03, 0xa8, 0x45, 0x64, 0x1a, 0x04, 0xa0, 0x45,
	0xf4, 0x03, 0x42, 0x53, 0x0a, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x58, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x6d, 0x68, 0x75, 0x62,
	0x2f, 0x62, 0x69, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x6d, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55,
	0x72, 0x6c, 0x58, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c,
	0x58, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shortUrlX_v1_shortUrlx_error_proto_rawDescOnce sync.Once
	file_shortUrlX_v1_shortUrlx_error_proto_rawDescData = file_shortUrlX_v1_shortUrlx_error_proto_rawDesc
)

func file_shortUrlX_v1_shortUrlx_error_proto_rawDescGZIP() []byte {
	file_shortUrlX_v1_shortUrlx_error_proto_rawDescOnce.Do(func() {
		file_shortUrlX_v1_shortUrlx_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_shortUrlX_v1_shortUrlx_error_proto_rawDescData)
	})
	return file_shortUrlX_v1_shortUrlx_error_proto_rawDescData
}

var file_shortUrlX_v1_shortUrlx_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_shortUrlX_v1_shortUrlx_error_proto_goTypes = []interface{}{
	(ShortUrlErrorReason)(0), // 0: ShortUrlErrorReason
}
var file_shortUrlX_v1_shortUrlx_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shortUrlX_v1_shortUrlx_error_proto_init() }
func file_shortUrlX_v1_shortUrlx_error_proto_init() {
	if File_shortUrlX_v1_shortUrlx_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shortUrlX_v1_shortUrlx_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shortUrlX_v1_shortUrlx_error_proto_goTypes,
		DependencyIndexes: file_shortUrlX_v1_shortUrlx_error_proto_depIdxs,
		EnumInfos:         file_shortUrlX_v1_shortUrlx_error_proto_enumTypes,
	}.Build()
	File_shortUrlX_v1_shortUrlx_error_proto = out.File
	file_shortUrlX_v1_shortUrlx_error_proto_rawDesc = nil
	file_shortUrlX_v1_shortUrlx_error_proto_goTypes = nil
	file_shortUrlX_v1_shortUrlx_error_proto_depIdxs = nil
}
