// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.3
// source: act_service.proto

package songcontestraterprotos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_act_service_proto protoreflect.FileDescriptor

var file_act_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x72, 0x1a, 0x09, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x91, 0x03,
	0x0a, 0x03, 0x41, 0x63, 0x74, 0x12, 0x48, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x67,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4a, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74, 0x12, 0x1f, 0x2e, 0x73, 0x6f, 0x6e, 0x67,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x6f, 0x6e,
	0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x41, 0x63,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73,
	0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e,
	0x41, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x73, 0x6f, 0x6e,
	0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x72, 0x2e, 0x41, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x50, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x73,
	0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x1a, 0x5a, 0x18, 0x2e, 0x3b, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_act_service_proto_goTypes = []any{
	(*emptypb.Empty)(nil),    // 0: google.protobuf.Empty
	(*GetActRequest)(nil),    // 1: songcontestrater.GetActRequest
	(*CreateActRequest)(nil), // 2: songcontestrater.CreateActRequest
	(*UpdateActRequest)(nil), // 3: songcontestrater.UpdateActRequest
	(*DeleteActRequest)(nil), // 4: songcontestrater.DeleteActRequest
	(*ListActsResponse)(nil), // 5: songcontestrater.ListActsResponse
	(*ActResponse)(nil),      // 6: songcontestrater.ActResponse
}
var file_act_service_proto_depIdxs = []int32{
	0, // 0: songcontestrater.Act.ListActs:input_type -> google.protobuf.Empty
	1, // 1: songcontestrater.Act.GetAct:input_type -> songcontestrater.GetActRequest
	2, // 2: songcontestrater.Act.CreateAct:input_type -> songcontestrater.CreateActRequest
	3, // 3: songcontestrater.Act.UpdateAct:input_type -> songcontestrater.UpdateActRequest
	4, // 4: songcontestrater.Act.DeleteAct:input_type -> songcontestrater.DeleteActRequest
	5, // 5: songcontestrater.Act.ListActs:output_type -> songcontestrater.ListActsResponse
	6, // 6: songcontestrater.Act.GetAct:output_type -> songcontestrater.ActResponse
	6, // 7: songcontestrater.Act.CreateAct:output_type -> songcontestrater.ActResponse
	6, // 8: songcontestrater.Act.UpdateAct:output_type -> songcontestrater.ActResponse
	6, // 9: songcontestrater.Act.DeleteAct:output_type -> songcontestrater.ActResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_act_service_proto_init() }
func file_act_service_proto_init() {
	if File_act_service_proto != nil {
		return
	}
	file_act_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_act_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_act_service_proto_goTypes,
		DependencyIndexes: file_act_service_proto_depIdxs,
	}.Build()
	File_act_service_proto = out.File
	file_act_service_proto_rawDesc = nil
	file_act_service_proto_goTypes = nil
	file_act_service_proto_depIdxs = nil
}
