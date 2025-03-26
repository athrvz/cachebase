// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.0--rc1
// source: internal/proto/cache.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_internal_proto_cache_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Found         bool                   `protobuf:"varint,2,opt,name=found,proto3" json:"found,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_internal_proto_cache_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *GetResponse) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

type SetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Ttl           *int32                 `protobuf:"varint,3,opt,name=ttl,proto3,oneof" json:"ttl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	mi := &file_internal_proto_cache_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{2}
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *SetRequest) GetTtl() int32 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

type SetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	mi := &file_internal_proto_cache_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{3}
}

func (x *SetResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_internal_proto_cache_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	mi := &file_internal_proto_cache_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pong          string                 `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PingResponse) Reset() {
	*x = PingResponse{}
	mi := &file_internal_proto_cache_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResponse) ProtoMessage() {}

func (x *PingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResponse.ProtoReflect.Descriptor instead.
func (*PingResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{6}
}

func (x *PingResponse) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

// (ds) List
type ListPushRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values        []string               `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPushRequest) Reset() {
	*x = ListPushRequest{}
	mi := &file_internal_proto_cache_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPushRequest) ProtoMessage() {}

func (x *ListPushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPushRequest.ProtoReflect.Descriptor instead.
func (*ListPushRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{7}
}

func (x *ListPushRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ListPushRequest) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type ListPushResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Length        int32                  `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPushResponse) Reset() {
	*x = ListPushResponse{}
	mi := &file_internal_proto_cache_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPushResponse) ProtoMessage() {}

func (x *ListPushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPushResponse.ProtoReflect.Descriptor instead.
func (*ListPushResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{8}
}

func (x *ListPushResponse) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type ListPopRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPopRequest) Reset() {
	*x = ListPopRequest{}
	mi := &file_internal_proto_cache_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPopRequest) ProtoMessage() {}

func (x *ListPopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPopRequest.ProtoReflect.Descriptor instead.
func (*ListPopRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{9}
}

func (x *ListPopRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ListPopResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         string                 `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListPopResponse) Reset() {
	*x = ListPopResponse{}
	mi := &file_internal_proto_cache_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListPopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPopResponse) ProtoMessage() {}

func (x *ListPopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPopResponse.ProtoReflect.Descriptor instead.
func (*ListPopResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{10}
}

func (x *ListPopResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// (ds) Set
type SetAddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values        []string               `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	Ttl           *int32                 `protobuf:"varint,3,opt,name=ttl,proto3,oneof" json:"ttl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetAddRequest) Reset() {
	*x = SetAddRequest{}
	mi := &file_internal_proto_cache_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAddRequest) ProtoMessage() {}

func (x *SetAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAddRequest.ProtoReflect.Descriptor instead.
func (*SetAddRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{11}
}

func (x *SetAddRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetAddRequest) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *SetAddRequest) GetTtl() int32 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

type SetAddResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Added         int32                  `protobuf:"varint,1,opt,name=added,proto3" json:"added,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetAddResponse) Reset() {
	*x = SetAddResponse{}
	mi := &file_internal_proto_cache_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAddResponse) ProtoMessage() {}

func (x *SetAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAddResponse.ProtoReflect.Descriptor instead.
func (*SetAddResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{12}
}

func (x *SetAddResponse) GetAdded() int32 {
	if x != nil {
		return x.Added
	}
	return 0
}

type SetRemoveRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values        []string               `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetRemoveRequest) Reset() {
	*x = SetRemoveRequest{}
	mi := &file_internal_proto_cache_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRemoveRequest) ProtoMessage() {}

func (x *SetRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRemoveRequest.ProtoReflect.Descriptor instead.
func (*SetRemoveRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{13}
}

func (x *SetRemoveRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRemoveRequest) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

type SetRemoveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Removed       int32                  `protobuf:"varint,1,opt,name=removed,proto3" json:"removed,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetRemoveResponse) Reset() {
	*x = SetRemoveResponse{}
	mi := &file_internal_proto_cache_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetRemoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRemoveResponse) ProtoMessage() {}

func (x *SetRemoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRemoveResponse.ProtoReflect.Descriptor instead.
func (*SetRemoveResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{14}
}

func (x *SetRemoveResponse) GetRemoved() int32 {
	if x != nil {
		return x.Removed
	}
	return 0
}

type SetIsMemberRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetIsMemberRequest) Reset() {
	*x = SetIsMemberRequest{}
	mi := &file_internal_proto_cache_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetIsMemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIsMemberRequest) ProtoMessage() {}

func (x *SetIsMemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIsMemberRequest.ProtoReflect.Descriptor instead.
func (*SetIsMemberRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{15}
}

func (x *SetIsMemberRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetIsMemberRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type SetIsMemberResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsMember      bool                   `protobuf:"varint,1,opt,name=is_member,json=isMember,proto3" json:"is_member,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetIsMemberResponse) Reset() {
	*x = SetIsMemberResponse{}
	mi := &file_internal_proto_cache_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetIsMemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetIsMemberResponse) ProtoMessage() {}

func (x *SetIsMemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetIsMemberResponse.ProtoReflect.Descriptor instead.
func (*SetIsMemberResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{16}
}

func (x *SetIsMemberResponse) GetIsMember() bool {
	if x != nil {
		return x.IsMember
	}
	return false
}

type SetMembersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Key           string                 `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetMembersRequest) Reset() {
	*x = SetMembersRequest{}
	mi := &file_internal_proto_cache_proto_msgTypes[17]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMembersRequest) ProtoMessage() {}

func (x *SetMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[17]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMembersRequest.ProtoReflect.Descriptor instead.
func (*SetMembersRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{17}
}

func (x *SetMembersRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type SetMembersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Members       []string               `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetMembersResponse) Reset() {
	*x = SetMembersResponse{}
	mi := &file_internal_proto_cache_proto_msgTypes[18]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMembersResponse) ProtoMessage() {}

func (x *SetMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_cache_proto_msgTypes[18]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMembersResponse.ProtoReflect.Descriptor instead.
func (*SetMembersResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_cache_proto_rawDescGZIP(), []int{18}
}

func (x *SetMembersResponse) GetMembers() []string {
	if x != nil {
		return x.Members
	}
	return nil
}

var File_internal_proto_cache_proto protoreflect.FileDescriptor

var file_internal_proto_cache_proto_rawDesc = string([]byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x1e, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x22, 0x39, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x22, 0x53, 0x0a, 0x0a, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x15, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00,
	0x52, 0x03, 0x74, 0x74, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x74, 0x6c,
	0x22, 0x27, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x21, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2a, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x22, 0x0a, 0x0c, 0x50, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x3b, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x2a, 0x0a, 0x10, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x22, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x27, 0x0a, 0x0f, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x58, 0x0a, 0x0d, 0x53, 0x65, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x15, 0x0a,
	0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x74, 0x74,
	0x6c, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x74, 0x6c, 0x22, 0x26, 0x0a, 0x0e,
	0x53, 0x65, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61,
	0x64, 0x64, 0x65, 0x64, 0x22, 0x3c, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0x2d, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x64, 0x22, 0x3c, 0x0a, 0x12, 0x53, 0x65, 0x74, 0x49, 0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x32, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x49, 0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x25, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x65,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x32, 0xcd, 0x04, 0x0a, 0x0c, 0x43,
	0x61, 0x63, 0x68, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x53, 0x65, 0x74,
	0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x73, 0x68, 0x12,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x38, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x70, 0x12, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x06, 0x53, 0x65,
	0x74, 0x41, 0x64, 0x64, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3e, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x65, 0x74, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x44, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x49, 0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x73, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x49, 0x73, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_internal_proto_cache_proto_rawDescOnce sync.Once
	file_internal_proto_cache_proto_rawDescData []byte
)

func file_internal_proto_cache_proto_rawDescGZIP() []byte {
	file_internal_proto_cache_proto_rawDescOnce.Do(func() {
		file_internal_proto_cache_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_proto_cache_proto_rawDesc), len(file_internal_proto_cache_proto_rawDesc)))
	})
	return file_internal_proto_cache_proto_rawDescData
}

var file_internal_proto_cache_proto_msgTypes = make([]protoimpl.MessageInfo, 19)
var file_internal_proto_cache_proto_goTypes = []any{
	(*GetRequest)(nil),          // 0: proto.GetRequest
	(*GetResponse)(nil),         // 1: proto.GetResponse
	(*SetRequest)(nil),          // 2: proto.SetRequest
	(*SetResponse)(nil),         // 3: proto.SetResponse
	(*DeleteRequest)(nil),       // 4: proto.DeleteRequest
	(*DeleteResponse)(nil),      // 5: proto.DeleteResponse
	(*PingResponse)(nil),        // 6: proto.PingResponse
	(*ListPushRequest)(nil),     // 7: proto.ListPushRequest
	(*ListPushResponse)(nil),    // 8: proto.ListPushResponse
	(*ListPopRequest)(nil),      // 9: proto.ListPopRequest
	(*ListPopResponse)(nil),     // 10: proto.ListPopResponse
	(*SetAddRequest)(nil),       // 11: proto.SetAddRequest
	(*SetAddResponse)(nil),      // 12: proto.SetAddResponse
	(*SetRemoveRequest)(nil),    // 13: proto.SetRemoveRequest
	(*SetRemoveResponse)(nil),   // 14: proto.SetRemoveResponse
	(*SetIsMemberRequest)(nil),  // 15: proto.SetIsMemberRequest
	(*SetIsMemberResponse)(nil), // 16: proto.SetIsMemberResponse
	(*SetMembersRequest)(nil),   // 17: proto.SetMembersRequest
	(*SetMembersResponse)(nil),  // 18: proto.SetMembersResponse
	(*emptypb.Empty)(nil),       // 19: google.protobuf.Empty
}
var file_internal_proto_cache_proto_depIdxs = []int32{
	0,  // 0: proto.CacheService.Get:input_type -> proto.GetRequest
	2,  // 1: proto.CacheService.Set:input_type -> proto.SetRequest
	4,  // 2: proto.CacheService.Delete:input_type -> proto.DeleteRequest
	19, // 3: proto.CacheService.Ping:input_type -> google.protobuf.Empty
	7,  // 4: proto.CacheService.ListPush:input_type -> proto.ListPushRequest
	9,  // 5: proto.CacheService.ListPop:input_type -> proto.ListPopRequest
	11, // 6: proto.CacheService.SetAdd:input_type -> proto.SetAddRequest
	13, // 7: proto.CacheService.SetRemove:input_type -> proto.SetRemoveRequest
	15, // 8: proto.CacheService.SetIsMember:input_type -> proto.SetIsMemberRequest
	17, // 9: proto.CacheService.SetMembers:input_type -> proto.SetMembersRequest
	1,  // 10: proto.CacheService.Get:output_type -> proto.GetResponse
	3,  // 11: proto.CacheService.Set:output_type -> proto.SetResponse
	5,  // 12: proto.CacheService.Delete:output_type -> proto.DeleteResponse
	6,  // 13: proto.CacheService.Ping:output_type -> proto.PingResponse
	8,  // 14: proto.CacheService.ListPush:output_type -> proto.ListPushResponse
	10, // 15: proto.CacheService.ListPop:output_type -> proto.ListPopResponse
	12, // 16: proto.CacheService.SetAdd:output_type -> proto.SetAddResponse
	14, // 17: proto.CacheService.SetRemove:output_type -> proto.SetRemoveResponse
	16, // 18: proto.CacheService.SetIsMember:output_type -> proto.SetIsMemberResponse
	18, // 19: proto.CacheService.SetMembers:output_type -> proto.SetMembersResponse
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_cache_proto_init() }
func file_internal_proto_cache_proto_init() {
	if File_internal_proto_cache_proto != nil {
		return
	}
	file_internal_proto_cache_proto_msgTypes[2].OneofWrappers = []any{}
	file_internal_proto_cache_proto_msgTypes[11].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_proto_cache_proto_rawDesc), len(file_internal_proto_cache_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   19,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_cache_proto_goTypes,
		DependencyIndexes: file_internal_proto_cache_proto_depIdxs,
		MessageInfos:      file_internal_proto_cache_proto_msgTypes,
	}.Build()
	File_internal_proto_cache_proto = out.File
	file_internal_proto_cache_proto_goTypes = nil
	file_internal_proto_cache_proto_depIdxs = nil
}
