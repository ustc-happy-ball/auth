// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.11.4
// source: grpc_req_res.proto

package databaseGrpc

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_grpc_req_res_proto protoreflect.FileDescriptor

var file_grpc_req_res_proto_rawDesc = []byte{
	0x0a, 0x12, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x72, 0x65, 0x71, 0x5f, 0x72, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72,
	0x70, 0x63, 0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0xce, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x69, 0x0a, 0x12, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46,
	0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x27, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72,
	0x70, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x51, 0x0a, 0x0a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x12, 0x1f, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x32, 0xd0, 0x01, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6f, 0x0a, 0x14, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69,
	0x6e, 0x64, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x29, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61,
	0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x12, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x47, 0x72, 0x70,
	0x63, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x47, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_grpc_req_res_proto_goTypes = []interface{}{
	(*AccountFindByPhoneRequest)(nil),    // 0: databaseGrpc.AccountFindByPhoneRequest
	(*AccountAddRequest)(nil),            // 1: databaseGrpc.AccountAddRequest
	(*PlayerFindByPlayerIdRequest)(nil),  // 2: databaseGrpc.PlayerFindByPlayerIdRequest
	(*PlayerAddRequest)(nil),             // 3: databaseGrpc.PlayerAddRequest
	(*AccountFindByPhoneResponse)(nil),   // 4: databaseGrpc.AccountFindByPhoneResponse
	(*AccountAddResponse)(nil),           // 5: databaseGrpc.AccountAddResponse
	(*PlayerFindByPlayerIdResponse)(nil), // 6: databaseGrpc.PlayerFindByPlayerIdResponse
	(*PlayerAddResponse)(nil),            // 7: databaseGrpc.PlayerAddResponse
}
var file_grpc_req_res_proto_depIdxs = []int32{
	0, // 0: databaseGrpc.AccountService.AccountFindByPhone:input_type -> databaseGrpc.AccountFindByPhoneRequest
	1, // 1: databaseGrpc.AccountService.AccountAdd:input_type -> databaseGrpc.AccountAddRequest
	2, // 2: databaseGrpc.PlayerService.PlayerFindByPlayerId:input_type -> databaseGrpc.PlayerFindByPlayerIdRequest
	3, // 3: databaseGrpc.PlayerService.PlayerAdd:input_type -> databaseGrpc.PlayerAddRequest
	4, // 4: databaseGrpc.AccountService.AccountFindByPhone:output_type -> databaseGrpc.AccountFindByPhoneResponse
	5, // 5: databaseGrpc.AccountService.AccountAdd:output_type -> databaseGrpc.AccountAddResponse
	6, // 6: databaseGrpc.PlayerService.PlayerFindByPlayerId:output_type -> databaseGrpc.PlayerFindByPlayerIdResponse
	7, // 7: databaseGrpc.PlayerService.PlayerAdd:output_type -> databaseGrpc.PlayerAddResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_req_res_proto_init() }
func file_grpc_req_res_proto_init() {
	if File_grpc_req_res_proto != nil {
		return
	}
	file_request_proto_init()
	file_response_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_req_res_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_grpc_req_res_proto_goTypes,
		DependencyIndexes: file_grpc_req_res_proto_depIdxs,
	}.Build()
	File_grpc_req_res_proto = out.File
	file_grpc_req_res_proto_rawDesc = nil
	file_grpc_req_res_proto_goTypes = nil
	file_grpc_req_res_proto_depIdxs = nil
}