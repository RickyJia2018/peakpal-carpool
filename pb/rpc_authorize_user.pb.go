// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.2
// source: rpc_authorize_user.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthorizeUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *AuthorizeUserRequest) Reset() {
	*x = AuthorizeUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_authorize_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeUserRequest) ProtoMessage() {}

func (x *AuthorizeUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_authorize_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeUserRequest.ProtoReflect.Descriptor instead.
func (*AuthorizeUserRequest) Descriptor() ([]byte, []int) {
	return file_rpc_authorize_user_proto_rawDescGZIP(), []int{0}
}

func (x *AuthorizeUserRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type AuthorizeUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        string                 `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserId    int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IssuedAt  *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=issuedAt,proto3" json:"issuedAt,omitempty"`
	ExpiredAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
}

func (x *AuthorizeUserResponse) Reset() {
	*x = AuthorizeUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_authorize_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthorizeUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthorizeUserResponse) ProtoMessage() {}

func (x *AuthorizeUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_authorize_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthorizeUserResponse.ProtoReflect.Descriptor instead.
func (*AuthorizeUserResponse) Descriptor() ([]byte, []int) {
	return file_rpc_authorize_user_proto_rawDescGZIP(), []int{1}
}

func (x *AuthorizeUserResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *AuthorizeUserResponse) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthorizeUserResponse) GetIssuedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.IssuedAt
	}
	return nil
}

func (x *AuthorizeUserResponse) GetExpiredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiredAt
	}
	return nil
}

var File_rpc_authorize_user_proto protoreflect.FileDescriptor

var file_rpc_authorize_user_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x39, 0x0a, 0x14, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb2, 0x01, 0x0a, 0x15, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x36, 0x0a,
	0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x64, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x64, 0x41, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69,
	0x63, 0x6b, 0x79, 0x6a, 0x69, 0x61, 0x32, 0x30, 0x31, 0x38, 0x2f, 0x70, 0x65, 0x61, 0x6b, 0x2d,
	0x70, 0x61, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_authorize_user_proto_rawDescOnce sync.Once
	file_rpc_authorize_user_proto_rawDescData = file_rpc_authorize_user_proto_rawDesc
)

func file_rpc_authorize_user_proto_rawDescGZIP() []byte {
	file_rpc_authorize_user_proto_rawDescOnce.Do(func() {
		file_rpc_authorize_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_authorize_user_proto_rawDescData)
	})
	return file_rpc_authorize_user_proto_rawDescData
}

var file_rpc_authorize_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rpc_authorize_user_proto_goTypes = []interface{}{
	(*AuthorizeUserRequest)(nil),  // 0: pb.AuthorizeUserRequest
	(*AuthorizeUserResponse)(nil), // 1: pb.AuthorizeUserResponse
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_rpc_authorize_user_proto_depIdxs = []int32{
	2, // 0: pb.AuthorizeUserResponse.issuedAt:type_name -> google.protobuf.Timestamp
	2, // 1: pb.AuthorizeUserResponse.expiredAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_rpc_authorize_user_proto_init() }
func file_rpc_authorize_user_proto_init() {
	if File_rpc_authorize_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_authorize_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeUserRequest); i {
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
		file_rpc_authorize_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthorizeUserResponse); i {
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
			RawDescriptor: file_rpc_authorize_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_authorize_user_proto_goTypes,
		DependencyIndexes: file_rpc_authorize_user_proto_depIdxs,
		MessageInfos:      file_rpc_authorize_user_proto_msgTypes,
	}.Build()
	File_rpc_authorize_user_proto = out.File
	file_rpc_authorize_user_proto_rawDesc = nil
	file_rpc_authorize_user_proto_goTypes = nil
	file_rpc_authorize_user_proto_depIdxs = nil
}
