// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/protobuf/api.proto

package grpc_server

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

// Message
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_protobuf_api_proto_rawDescGZIP(), []int{0}
}

// User request
type UserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *UserListRequest) Reset() {
	*x = UserListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListRequest) ProtoMessage() {}

func (x *UserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListRequest.ProtoReflect.Descriptor instead.
func (*UserListRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_api_proto_rawDescGZIP(), []int{1}
}

func (x *UserListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *UserListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type UserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserIDRequest) Reset() {
	*x = UserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserIDRequest) ProtoMessage() {}

func (x *UserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserIDRequest.ProtoReflect.Descriptor instead.
func (*UserIDRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_api_proto_rawDescGZIP(), []int{2}
}

func (x *UserIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UserCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginId  string `protobuf:"bytes,1,opt,name=loginId,proto3" json:"loginId,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserCreateRequest) Reset() {
	*x = UserCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCreateRequest) ProtoMessage() {}

func (x *UserCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCreateRequest.ProtoReflect.Descriptor instead.
func (*UserCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_api_proto_rawDescGZIP(), []int{3}
}

func (x *UserCreateRequest) GetLoginId() string {
	if x != nil {
		return x.LoginId
	}
	return ""
}

func (x *UserCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserCreateRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserCreateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type UserUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserUpdateRequest) Reset() {
	*x = UserUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserUpdateRequest) ProtoMessage() {}

func (x *UserUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserUpdateRequest.ProtoReflect.Descriptor instead.
func (*UserUpdateRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_api_proto_rawDescGZIP(), []int{4}
}

func (x *UserUpdateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserUpdateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserUpdateRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserUpdateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// User response
type UserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uesrs []*UserInfoResponse `protobuf:"bytes,1,rep,name=uesrs,proto3" json:"uesrs,omitempty"`
}

func (x *UserListResponse) Reset() {
	*x = UserListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListResponse) ProtoMessage() {}

func (x *UserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListResponse.ProtoReflect.Descriptor instead.
func (*UserListResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_api_proto_rawDescGZIP(), []int{5}
}

func (x *UserListResponse) GetUesrs() []*UserInfoResponse {
	if x != nil {
		return x.Uesrs
	}
	return nil
}

type UserInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LoginId string `protobuf:"bytes,2,opt,name=loginId,proto3" json:"loginId,omitempty"`
	Phone   string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Email   string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *UserInfoResponse) Reset() {
	*x = UserInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoResponse) ProtoMessage() {}

func (x *UserInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoResponse.ProtoReflect.Descriptor instead.
func (*UserInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_api_proto_rawDescGZIP(), []int{6}
}

func (x *UserInfoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfoResponse) GetLoginId() string {
	if x != nil {
		return x.LoginId
	}
	return ""
}

func (x *UserInfoResponse) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserInfoResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Token request
type TokenCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginId  string `protobuf:"bytes,1,opt,name=loginId,proto3" json:"loginId,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *TokenCreateRequest) Reset() {
	*x = TokenCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenCreateRequest) ProtoMessage() {}

func (x *TokenCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenCreateRequest.ProtoReflect.Descriptor instead.
func (*TokenCreateRequest) Descriptor() ([]byte, []int) {
	return file_api_protobuf_api_proto_rawDescGZIP(), []int{7}
}

func (x *TokenCreateRequest) GetLoginId() string {
	if x != nil {
		return x.LoginId
	}
	return ""
}

func (x *TokenCreateRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Token response
type TokenInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *TokenInfoResponse) Reset() {
	*x = TokenInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobuf_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenInfoResponse) ProtoMessage() {}

func (x *TokenInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobuf_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenInfoResponse.ProtoReflect.Descriptor instead.
func (*TokenInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_protobuf_api_proto_rawDescGZIP(), []int{8}
}

func (x *TokenInfoResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_api_protobuf_api_proto protoreflect.FileDescriptor

var file_api_protobuf_api_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x3f, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x75, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x6b, 0x0a, 0x11, 0x55, 0x73,
	0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3b, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x75,
	0x65, 0x73, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x75,
	0x65, 0x73, 0x72, 0x73, 0x22, 0x68, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x4a,
	0x0a, 0x12, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x29, 0x0a, 0x11, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xf4, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x31,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x12, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x32, 0x41, 0x0a, 0x05,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x38, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x13, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x1d, 0x5a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protobuf_api_proto_rawDescOnce sync.Once
	file_api_protobuf_api_proto_rawDescData = file_api_protobuf_api_proto_rawDesc
)

func file_api_protobuf_api_proto_rawDescGZIP() []byte {
	file_api_protobuf_api_proto_rawDescOnce.Do(func() {
		file_api_protobuf_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protobuf_api_proto_rawDescData)
	})
	return file_api_protobuf_api_proto_rawDescData
}

var file_api_protobuf_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_api_protobuf_api_proto_goTypes = []interface{}{
	(*Empty)(nil),              // 0: Empty
	(*UserListRequest)(nil),    // 1: UserListRequest
	(*UserIDRequest)(nil),      // 2: UserIDRequest
	(*UserCreateRequest)(nil),  // 3: UserCreateRequest
	(*UserUpdateRequest)(nil),  // 4: UserUpdateRequest
	(*UserListResponse)(nil),   // 5: UserListResponse
	(*UserInfoResponse)(nil),   // 6: UserInfoResponse
	(*TokenCreateRequest)(nil), // 7: TokenCreateRequest
	(*TokenInfoResponse)(nil),  // 8: TokenInfoResponse
}
var file_api_protobuf_api_proto_depIdxs = []int32{
	6, // 0: UserListResponse.uesrs:type_name -> UserInfoResponse
	1, // 1: User.ListUser:input_type -> UserListRequest
	3, // 2: User.CreateUser:input_type -> UserCreateRequest
	2, // 3: User.GetUser:input_type -> UserIDRequest
	4, // 4: User.UpdateUser:input_type -> UserUpdateRequest
	2, // 5: User.DeleteUser:input_type -> UserIDRequest
	7, // 6: Token.CreateToken:input_type -> TokenCreateRequest
	5, // 7: User.ListUser:output_type -> UserListResponse
	6, // 8: User.CreateUser:output_type -> UserInfoResponse
	6, // 9: User.GetUser:output_type -> UserInfoResponse
	0, // 10: User.UpdateUser:output_type -> Empty
	0, // 11: User.DeleteUser:output_type -> Empty
	8, // 12: Token.CreateToken:output_type -> TokenInfoResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_protobuf_api_proto_init() }
func file_api_protobuf_api_proto_init() {
	if File_api_protobuf_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protobuf_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_api_protobuf_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListRequest); i {
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
		file_api_protobuf_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserIDRequest); i {
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
		file_api_protobuf_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCreateRequest); i {
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
		file_api_protobuf_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserUpdateRequest); i {
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
		file_api_protobuf_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserListResponse); i {
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
		file_api_protobuf_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoResponse); i {
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
		file_api_protobuf_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenCreateRequest); i {
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
		file_api_protobuf_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenInfoResponse); i {
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
			RawDescriptor: file_api_protobuf_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_api_protobuf_api_proto_goTypes,
		DependencyIndexes: file_api_protobuf_api_proto_depIdxs,
		MessageInfos:      file_api_protobuf_api_proto_msgTypes,
	}.Build()
	File_api_protobuf_api_proto = out.File
	file_api_protobuf_api_proto_rawDesc = nil
	file_api_protobuf_api_proto_goTypes = nil
	file_api_protobuf_api_proto_depIdxs = nil
}
