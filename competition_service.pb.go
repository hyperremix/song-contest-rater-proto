// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.3
// source: competition_service.proto

package __

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

var File_competition_service_proto protoreflect.FileDescriptor

var file_competition_service_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x6f, 0x6e,
	0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x1a, 0x11, 0x63,
	0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x89, 0x04,
	0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2a, 0x2e, 0x73, 0x6f, 0x6e, 0x67,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x73, 0x6f, 0x6e, 0x67,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x2a, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73,
	0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x73, 0x6f, 0x6e,
	0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x68, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x73, 0x6f, 0x6e, 0x67, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_competition_service_proto_goTypes = []any{
	(*emptypb.Empty)(nil),            // 0: google.protobuf.Empty
	(*GetCompetitionRequest)(nil),    // 1: songcontestrater.GetCompetitionRequest
	(*CreateCompetitionRequest)(nil), // 2: songcontestrater.CreateCompetitionRequest
	(*UpdateCompetitionRequest)(nil), // 3: songcontestrater.UpdateCompetitionRequest
	(*DeleteCompetitionRequest)(nil), // 4: songcontestrater.DeleteCompetitionRequest
	(*ListCompetitionsResponse)(nil), // 5: songcontestrater.ListCompetitionsResponse
	(*CompetitionResponse)(nil),      // 6: songcontestrater.CompetitionResponse
}
var file_competition_service_proto_depIdxs = []int32{
	0, // 0: songcontestrater.Competition.ListCompetitions:input_type -> google.protobuf.Empty
	1, // 1: songcontestrater.Competition.GetCompetition:input_type -> songcontestrater.GetCompetitionRequest
	2, // 2: songcontestrater.Competition.CreateCompetition:input_type -> songcontestrater.CreateCompetitionRequest
	3, // 3: songcontestrater.Competition.UpdateCompetition:input_type -> songcontestrater.UpdateCompetitionRequest
	4, // 4: songcontestrater.Competition.DeleteCompetition:input_type -> songcontestrater.DeleteCompetitionRequest
	5, // 5: songcontestrater.Competition.ListCompetitions:output_type -> songcontestrater.ListCompetitionsResponse
	6, // 6: songcontestrater.Competition.GetCompetition:output_type -> songcontestrater.CompetitionResponse
	6, // 7: songcontestrater.Competition.CreateCompetition:output_type -> songcontestrater.CompetitionResponse
	6, // 8: songcontestrater.Competition.UpdateCompetition:output_type -> songcontestrater.CompetitionResponse
	6, // 9: songcontestrater.Competition.DeleteCompetition:output_type -> songcontestrater.CompetitionResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_competition_service_proto_init() }
func file_competition_service_proto_init() {
	if File_competition_service_proto != nil {
		return
	}
	file_competition_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_competition_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_competition_service_proto_goTypes,
		DependencyIndexes: file_competition_service_proto_depIdxs,
	}.Build()
	File_competition_service_proto = out.File
	file_competition_service_proto_rawDesc = nil
	file_competition_service_proto_goTypes = nil
	file_competition_service_proto_depIdxs = nil
}
