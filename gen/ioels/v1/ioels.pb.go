// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: ioels/v1/ioels.proto

package ioelsv1

import (
	_ "github.com/irisco88/protos/gen/common/v1"
	_ "github.com/irisco88/protos/gen/google/api"
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

type Ioels struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ElementName  string `protobuf:"bytes,2,opt,name=element_name,json=elementName,proto3" json:"element_name,omitempty"`
	ElementLable string `protobuf:"bytes,3,opt,name=element_lable,json=elementLable,proto3" json:"element_lable,omitempty"`
	ElementColor string `protobuf:"bytes,4,opt,name=element_color,json=elementColor,proto3" json:"element_color,omitempty"`
}

func (x *Ioels) Reset() {
	*x = Ioels{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioels_v1_ioels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ioels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ioels) ProtoMessage() {}

func (x *Ioels) ProtoReflect() protoreflect.Message {
	mi := &file_ioels_v1_ioels_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ioels.ProtoReflect.Descriptor instead.
func (*Ioels) Descriptor() ([]byte, []int) {
	return file_ioels_v1_ioels_proto_rawDescGZIP(), []int{0}
}

func (x *Ioels) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ioels) GetElementName() string {
	if x != nil {
		return x.ElementName
	}
	return ""
}

func (x *Ioels) GetElementLable() string {
	if x != nil {
		return x.ElementLable
	}
	return ""
}

func (x *Ioels) GetElementColor() string {
	if x != nil {
		return x.ElementColor
	}
	return ""
}

type ListIoelssRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListIoelssRequest) Reset() {
	*x = ListIoelssRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioels_v1_ioels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIoelssRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIoelssRequest) ProtoMessage() {}

func (x *ListIoelssRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ioels_v1_ioels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIoelssRequest.ProtoReflect.Descriptor instead.
func (*ListIoelssRequest) Descriptor() ([]byte, []int) {
	return file_ioels_v1_ioels_proto_rawDescGZIP(), []int{1}
}

type ListIoelssResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ioelss []*Ioels `protobuf:"bytes,1,rep,name=ioelss,proto3" json:"ioelss,omitempty"`
}

func (x *ListIoelssResponse) Reset() {
	*x = ListIoelssResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioels_v1_ioels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIoelssResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIoelssResponse) ProtoMessage() {}

func (x *ListIoelssResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ioels_v1_ioels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIoelssResponse.ProtoReflect.Descriptor instead.
func (*ListIoelssResponse) Descriptor() ([]byte, []int) {
	return file_ioels_v1_ioels_proto_rawDescGZIP(), []int{2}
}

func (x *ListIoelssResponse) GetIoelss() []*Ioels {
	if x != nil {
		return x.Ioelss
	}
	return nil
}

type UpdateIoelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ioels *Ioels `protobuf:"bytes,1,opt,name=ioels,proto3" json:"ioels,omitempty"`
}

func (x *UpdateIoelsRequest) Reset() {
	*x = UpdateIoelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioels_v1_ioels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIoelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIoelsRequest) ProtoMessage() {}

func (x *UpdateIoelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ioels_v1_ioels_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIoelsRequest.ProtoReflect.Descriptor instead.
func (*UpdateIoelsRequest) Descriptor() ([]byte, []int) {
	return file_ioels_v1_ioels_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateIoelsRequest) GetIoels() *Ioels {
	if x != nil {
		return x.Ioels
	}
	return nil
}

type UpdateIoelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateIoelsResponse) Reset() {
	*x = UpdateIoelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioels_v1_ioels_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIoelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIoelsResponse) ProtoMessage() {}

func (x *UpdateIoelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ioels_v1_ioels_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIoelsResponse.ProtoReflect.Descriptor instead.
func (*UpdateIoelsResponse) Descriptor() ([]byte, []int) {
	return file_ioels_v1_ioels_proto_rawDescGZIP(), []int{4}
}

type CreateIoelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ioels *Ioels `protobuf:"bytes,1,opt,name=ioels,proto3" json:"ioels,omitempty"`
}

func (x *CreateIoelsRequest) Reset() {
	*x = CreateIoelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioels_v1_ioels_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIoelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIoelsRequest) ProtoMessage() {}

func (x *CreateIoelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ioels_v1_ioels_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIoelsRequest.ProtoReflect.Descriptor instead.
func (*CreateIoelsRequest) Descriptor() ([]byte, []int) {
	return file_ioels_v1_ioels_proto_rawDescGZIP(), []int{5}
}

func (x *CreateIoelsRequest) GetIoels() *Ioels {
	if x != nil {
		return x.Ioels
	}
	return nil
}

type CreateIoelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IoelsId uint32 `protobuf:"varint,1,opt,name=ioels_id,json=ioelsId,proto3" json:"ioels_id,omitempty"`
}

func (x *CreateIoelsResponse) Reset() {
	*x = CreateIoelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioels_v1_ioels_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIoelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIoelsResponse) ProtoMessage() {}

func (x *CreateIoelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ioels_v1_ioels_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIoelsResponse.ProtoReflect.Descriptor instead.
func (*CreateIoelsResponse) Descriptor() ([]byte, []int) {
	return file_ioels_v1_ioels_proto_rawDescGZIP(), []int{6}
}

func (x *CreateIoelsResponse) GetIoelsId() uint32 {
	if x != nil {
		return x.IoelsId
	}
	return 0
}

type DeleteIoelsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IoelsId uint32 `protobuf:"varint,1,opt,name=ioels_id,json=ioelsId,proto3" json:"ioels_id,omitempty"`
}

func (x *DeleteIoelsRequest) Reset() {
	*x = DeleteIoelsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioels_v1_ioels_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIoelsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIoelsRequest) ProtoMessage() {}

func (x *DeleteIoelsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ioels_v1_ioels_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIoelsRequest.ProtoReflect.Descriptor instead.
func (*DeleteIoelsRequest) Descriptor() ([]byte, []int) {
	return file_ioels_v1_ioels_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteIoelsRequest) GetIoelsId() uint32 {
	if x != nil {
		return x.IoelsId
	}
	return 0
}

type DeleteIoelsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteIoelsResponse) Reset() {
	*x = DeleteIoelsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioels_v1_ioels_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIoelsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIoelsResponse) ProtoMessage() {}

func (x *DeleteIoelsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ioels_v1_ioels_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIoelsResponse.ProtoReflect.Descriptor instead.
func (*DeleteIoelsResponse) Descriptor() ([]byte, []int) {
	return file_ioels_v1_ioels_proto_rawDescGZIP(), []int{8}
}

var File_ioels_v1_ioels_proto protoreflect.FileDescriptor

var file_ioels_v1_ioels_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x01, 0x0a, 0x05, 0x49,
	0x6f, 0x65, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f,
	0x72, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6f,
	0x65, 0x6c, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06,
	0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69,
	0x6f, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52, 0x06, 0x69,
	0x6f, 0x65, 0x6c, 0x73, 0x73, 0x22, 0x3b, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6f, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x69,
	0x6f, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x6f, 0x65,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52, 0x05, 0x69, 0x6f, 0x65,
	0x6c, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x05, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52,
	0x05, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x22, 0x30, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x07, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x07, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xee, 0x03, 0x0a, 0x0c, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x79, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73,
	0x12, 0x1c, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x69, 0x6f,
	0x65, 0x6c, 0x73, 0x92, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x0a, 0x01, 0x02, 0x12, 0x79, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x12, 0x1c, 0x2e, 0x69, 0x6f,
	0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6f, 0x65, 0x6c,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x3a, 0x01, 0x2a, 0x1a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x65,
	0x6c, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x69, 0x6f, 0x65, 0x6c, 0x92, 0xd3, 0xe4,
	0x93, 0x02, 0x04, 0x0a, 0x02, 0x02, 0x00, 0x12, 0x82, 0x01, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x12, 0x1c, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x2a, 0x24, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2f, 0x7b, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x5f, 0x69,
	0x64, 0x7d, 0x92, 0xd3, 0xe4, 0x93, 0x02, 0x04, 0x0a, 0x02, 0x02, 0x00, 0x12, 0x63, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x12, 0x1b, 0x2e, 0x69, 0x6f, 0x65,
	0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x42, 0x8c, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x0a, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x72, 0x69,
	0x73, 0x63, 0x6f, 0x38, 0x38, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x08, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14,
	0x49, 0x6f, 0x65, 0x6c, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ioels_v1_ioels_proto_rawDescOnce sync.Once
	file_ioels_v1_ioels_proto_rawDescData = file_ioels_v1_ioels_proto_rawDesc
)

func file_ioels_v1_ioels_proto_rawDescGZIP() []byte {
	file_ioels_v1_ioels_proto_rawDescOnce.Do(func() {
		file_ioels_v1_ioels_proto_rawDescData = protoimpl.X.CompressGZIP(file_ioels_v1_ioels_proto_rawDescData)
	})
	return file_ioels_v1_ioels_proto_rawDescData
}

var file_ioels_v1_ioels_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_ioels_v1_ioels_proto_goTypes = []interface{}{
	(*Ioels)(nil),               // 0: ioels.v1.Ioels
	(*ListIoelssRequest)(nil),   // 1: ioels.v1.ListIoelssRequest
	(*ListIoelssResponse)(nil),  // 2: ioels.v1.ListIoelssResponse
	(*UpdateIoelsRequest)(nil),  // 3: ioels.v1.UpdateIoelsRequest
	(*UpdateIoelsResponse)(nil), // 4: ioels.v1.UpdateIoelsResponse
	(*CreateIoelsRequest)(nil),  // 5: ioels.v1.CreateIoelsRequest
	(*CreateIoelsResponse)(nil), // 6: ioels.v1.CreateIoelsResponse
	(*DeleteIoelsRequest)(nil),  // 7: ioels.v1.DeleteIoelsRequest
	(*DeleteIoelsResponse)(nil), // 8: ioels.v1.DeleteIoelsResponse
}
var file_ioels_v1_ioels_proto_depIdxs = []int32{
	0, // 0: ioels.v1.ListIoelssResponse.ioelss:type_name -> ioels.v1.Ioels
	0, // 1: ioels.v1.UpdateIoelsRequest.ioels:type_name -> ioels.v1.Ioels
	0, // 2: ioels.v1.CreateIoelsRequest.ioels:type_name -> ioels.v1.Ioels
	5, // 3: ioels.v1.IoelsService.CreateIoels:input_type -> ioels.v1.CreateIoelsRequest
	3, // 4: ioels.v1.IoelsService.UpdateIoels:input_type -> ioels.v1.UpdateIoelsRequest
	7, // 5: ioels.v1.IoelsService.DeleteIoels:input_type -> ioels.v1.DeleteIoelsRequest
	1, // 6: ioels.v1.IoelsService.ListIoelss:input_type -> ioels.v1.ListIoelssRequest
	6, // 7: ioels.v1.IoelsService.CreateIoels:output_type -> ioels.v1.CreateIoelsResponse
	4, // 8: ioels.v1.IoelsService.UpdateIoels:output_type -> ioels.v1.UpdateIoelsResponse
	8, // 9: ioels.v1.IoelsService.DeleteIoels:output_type -> ioels.v1.DeleteIoelsResponse
	2, // 10: ioels.v1.IoelsService.ListIoelss:output_type -> ioels.v1.ListIoelssResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ioels_v1_ioels_proto_init() }
func file_ioels_v1_ioels_proto_init() {
	if File_ioels_v1_ioels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ioels_v1_ioels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ioels); i {
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
		file_ioels_v1_ioels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIoelssRequest); i {
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
		file_ioels_v1_ioels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIoelssResponse); i {
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
		file_ioels_v1_ioels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIoelsRequest); i {
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
		file_ioels_v1_ioels_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIoelsResponse); i {
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
		file_ioels_v1_ioels_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIoelsRequest); i {
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
		file_ioels_v1_ioels_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIoelsResponse); i {
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
		file_ioels_v1_ioels_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIoelsRequest); i {
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
		file_ioels_v1_ioels_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIoelsResponse); i {
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
			RawDescriptor: file_ioels_v1_ioels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ioels_v1_ioels_proto_goTypes,
		DependencyIndexes: file_ioels_v1_ioels_proto_depIdxs,
		MessageInfos:      file_ioels_v1_ioels_proto_msgTypes,
	}.Build()
	File_ioels_v1_ioels_proto = out.File
	file_ioels_v1_ioels_proto_rawDesc = nil
	file_ioels_v1_ioels_proto_goTypes = nil
	file_ioels_v1_ioels_proto_depIdxs = nil
}
