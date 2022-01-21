// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// yext extends the Protobuf FieldOptions to add allow new options to be
// set storing characteristics of the YANG schema from which a protobuf
// is generated.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: yext.proto

package yext

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_yext_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1040,
		Name:          "yext.schemapath",
		Tag:           "bytes,1040,opt,name=schemapath",
		Filename:      "yext.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         1040,
		Name:          "yext.yang_name",
		Tag:           "bytes,1040,opt,name=yang_name",
		Filename:      "yext.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// schemapath stores the schema path to the field within the YANG schema.
	// The path stored is absolute if the entity is at the root of the schema
	// tree. In the case that it is not (i.e., it is an entity that has a parent
	// which is not a module, the path supplied is the relative path to the
	// parent of the entity). The field number for this extension is reserved
	// in the global protobuf registry.
	//
	// optional string schemapath = 1040;
	E_Schemapath = &file_yext_proto_extTypes[0]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// yang_name stores the original YANG name of the enumerated value, for
	// serialisation to a string. The field number for this extension is
	// reserved in the global protobuf registry.
	//
	// optional string yang_name = 1040;
	E_YangName = &file_yext_proto_extTypes[1]
)

var File_yext_proto protoreflect.FileDescriptor

var file_yext_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x79, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x79, 0x65,
	0x78, 0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x3e, 0x0a, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x90, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x70, 0x61, 0x74, 0x68, 0x3a, 0x3f, 0x0a, 0x09, 0x79, 0x61, 0x6e, 0x67, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x90, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x79, 0x61, 0x6e,
	0x67, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x79,
	0x67, 0x6f, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x65, 0x78, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_yext_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil),     // 0: google.protobuf.FieldOptions
	(*descriptorpb.EnumValueOptions)(nil), // 1: google.protobuf.EnumValueOptions
}
var file_yext_proto_depIdxs = []int32{
	0, // 0: yext.schemapath:extendee -> google.protobuf.FieldOptions
	1, // 1: yext.yang_name:extendee -> google.protobuf.EnumValueOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	0, // [0:2] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yext_proto_init() }
func file_yext_proto_init() {
	if File_yext_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_yext_proto_goTypes,
		DependencyIndexes: file_yext_proto_depIdxs,
		ExtensionInfos:    file_yext_proto_extTypes,
	}.Build()
	File_yext_proto = out.File
	file_yext_proto_rawDesc = nil
	file_yext_proto_goTypes = nil
	file_yext_proto_depIdxs = nil
}
