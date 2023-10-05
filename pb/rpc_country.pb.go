// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.24.2
// source: rpc_country.proto

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

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CountryCode string                 `protobuf:"bytes,3,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{0}
}

func (x *Country) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Country) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Country) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type CreateCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CountryCode string `protobuf:"bytes,2,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
}

func (x *CreateCountryRequest) Reset() {
	*x = CreateCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountryRequest) ProtoMessage() {}

func (x *CreateCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountryRequest.ProtoReflect.Descriptor instead.
func (*CreateCountryRequest) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCountryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCountryRequest) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

type CreateCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country *Country `protobuf:"bytes,1,opt,name=Country,proto3" json:"Country,omitempty"`
}

func (x *CreateCountryResponse) Reset() {
	*x = CreateCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCountryResponse) ProtoMessage() {}

func (x *CreateCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCountryResponse.ProtoReflect.Descriptor instead.
func (*CreateCountryResponse) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCountryResponse) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

type GetCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   *int64  `protobuf:"varint,1,opt,name=ID,proto3,oneof" json:"ID,omitempty"`
	Name *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *GetCountryRequest) Reset() {
	*x = GetCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountryRequest) ProtoMessage() {}

func (x *GetCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountryRequest.ProtoReflect.Descriptor instead.
func (*GetCountryRequest) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{3}
}

func (x *GetCountryRequest) GetID() int64 {
	if x != nil && x.ID != nil {
		return *x.ID
	}
	return 0
}

func (x *GetCountryRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

type GetCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country *Country `protobuf:"bytes,1,opt,name=Country,proto3" json:"Country,omitempty"`
}

func (x *GetCountryResponse) Reset() {
	*x = GetCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCountryResponse) ProtoMessage() {}

func (x *GetCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCountryResponse.ProtoReflect.Descriptor instead.
func (*GetCountryResponse) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{4}
}

func (x *GetCountryResponse) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

type ListCountriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListCountriesRequest) Reset() {
	*x = ListCountriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCountriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountriesRequest) ProtoMessage() {}

func (x *ListCountriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountriesRequest.ProtoReflect.Descriptor instead.
func (*ListCountriesRequest) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{5}
}

func (x *ListCountriesRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListCountriesRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListCountriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country []*Country `protobuf:"bytes,1,rep,name=Country,proto3" json:"Country,omitempty"`
	Limit   int32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset  int32      `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Total   int64      `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ListCountriesResponse) Reset() {
	*x = ListCountriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCountriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCountriesResponse) ProtoMessage() {}

func (x *ListCountriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCountriesResponse.ProtoReflect.Descriptor instead.
func (*ListCountriesResponse) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{6}
}

func (x *ListCountriesResponse) GetCountry() []*Country {
	if x != nil {
		return x.Country
	}
	return nil
}

func (x *ListCountriesResponse) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListCountriesResponse) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListCountriesResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type UpdateCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64   `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name        *string `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	CountryCode *string `protobuf:"bytes,3,opt,name=country_code,json=countryCode,proto3,oneof" json:"country_code,omitempty"`
}

func (x *UpdateCountryRequest) Reset() {
	*x = UpdateCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCountryRequest) ProtoMessage() {}

func (x *UpdateCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCountryRequest.ProtoReflect.Descriptor instead.
func (*UpdateCountryRequest) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCountryRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateCountryRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateCountryRequest) GetCountryCode() string {
	if x != nil && x.CountryCode != nil {
		return *x.CountryCode
	}
	return ""
}

type UpdateCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country *Country `protobuf:"bytes,1,opt,name=Country,proto3" json:"Country,omitempty"`
}

func (x *UpdateCountryResponse) Reset() {
	*x = UpdateCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCountryResponse) ProtoMessage() {}

func (x *UpdateCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCountryResponse.ProtoReflect.Descriptor instead.
func (*UpdateCountryResponse) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCountryResponse) GetCountry() *Country {
	if x != nil {
		return x.Country
	}
	return nil
}

type DeleteCountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteCountryRequest) Reset() {
	*x = DeleteCountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCountryRequest) ProtoMessage() {}

func (x *DeleteCountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCountryRequest.ProtoReflect.Descriptor instead.
func (*DeleteCountryRequest) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCountryRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteCountryResponse) Reset() {
	*x = DeleteCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_country_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCountryResponse) ProtoMessage() {}

func (x *DeleteCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_country_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCountryResponse.ProtoReflect.Descriptor instead.
func (*DeleteCountryResponse) Descriptor() ([]byte, []int) {
	return file_rpc_country_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteCountryResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_rpc_country_proto protoreflect.FileDescriptor

var file_rpc_country_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4d, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x3e, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x51, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x13, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x02, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12,
	0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x49, 0x44, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x44, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x15,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x22, 0x81, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x3e, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x31, 0x0a, 0x15,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x42,
	0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69,
	0x63, 0x6b, 0x79, 0x6a, 0x69, 0x61, 0x32, 0x30, 0x31, 0x38, 0x2f, 0x70, 0x65, 0x61, 0x6b, 0x2d,
	0x70, 0x61, 0x6c, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpc_country_proto_rawDescOnce sync.Once
	file_rpc_country_proto_rawDescData = file_rpc_country_proto_rawDesc
)

func file_rpc_country_proto_rawDescGZIP() []byte {
	file_rpc_country_proto_rawDescOnce.Do(func() {
		file_rpc_country_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_country_proto_rawDescData)
	})
	return file_rpc_country_proto_rawDescData
}

var file_rpc_country_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_rpc_country_proto_goTypes = []interface{}{
	(*Country)(nil),               // 0: pb.Country
	(*CreateCountryRequest)(nil),  // 1: pb.CreateCountryRequest
	(*CreateCountryResponse)(nil), // 2: pb.CreateCountryResponse
	(*GetCountryRequest)(nil),     // 3: pb.GetCountryRequest
	(*GetCountryResponse)(nil),    // 4: pb.GetCountryResponse
	(*ListCountriesRequest)(nil),  // 5: pb.ListCountriesRequest
	(*ListCountriesResponse)(nil), // 6: pb.ListCountriesResponse
	(*UpdateCountryRequest)(nil),  // 7: pb.UpdateCountryRequest
	(*UpdateCountryResponse)(nil), // 8: pb.UpdateCountryResponse
	(*DeleteCountryRequest)(nil),  // 9: pb.DeleteCountryRequest
	(*DeleteCountryResponse)(nil), // 10: pb.DeleteCountryResponse
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_rpc_country_proto_depIdxs = []int32{
	11, // 0: pb.Country.created_at:type_name -> google.protobuf.Timestamp
	0,  // 1: pb.CreateCountryResponse.Country:type_name -> pb.Country
	0,  // 2: pb.GetCountryResponse.Country:type_name -> pb.Country
	0,  // 3: pb.ListCountriesResponse.Country:type_name -> pb.Country
	0,  // 4: pb.UpdateCountryResponse.Country:type_name -> pb.Country
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_rpc_country_proto_init() }
func file_rpc_country_proto_init() {
	if File_rpc_country_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_country_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Country); i {
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
		file_rpc_country_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountryRequest); i {
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
		file_rpc_country_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCountryResponse); i {
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
		file_rpc_country_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountryRequest); i {
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
		file_rpc_country_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCountryResponse); i {
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
		file_rpc_country_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCountriesRequest); i {
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
		file_rpc_country_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCountriesResponse); i {
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
		file_rpc_country_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCountryRequest); i {
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
		file_rpc_country_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCountryResponse); i {
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
		file_rpc_country_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCountryRequest); i {
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
		file_rpc_country_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCountryResponse); i {
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
	file_rpc_country_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_rpc_country_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_rpc_country_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rpc_country_proto_goTypes,
		DependencyIndexes: file_rpc_country_proto_depIdxs,
		MessageInfos:      file_rpc_country_proto_msgTypes,
	}.Build()
	File_rpc_country_proto = out.File
	file_rpc_country_proto_rawDesc = nil
	file_rpc_country_proto_goTypes = nil
	file_rpc_country_proto_depIdxs = nil
}
