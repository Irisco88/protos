// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: ioelss/v1/ioelss.proto

package ioelssv1

import (
	_ "github.com/irisco88/protos/gen/common/v1"
	_ "github.com/irisco88/protos/gen/google/api"
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

type Ioelss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ElementName  string                 `protobuf:"bytes,2,opt,name=element_name,json=elementName,proto3" json:"element_name,omitempty"`
	ElementLable string                 `protobuf:"bytes,3,opt,name=element_lable,json=elementLable,proto3" json:"element_lable,omitempty"`
	ElementColor string                 `protobuf:"bytes,4,opt,name=element_color,json=elementColor,proto3" json:"element_color,omitempty"`
	EditDate     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=edit_date,json=editDate,proto3" json:"edit_date,omitempty"`
}

func (x *Ioelss) Reset() {
	*x = Ioelss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioelss_v1_ioelss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ioelss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ioelss) ProtoMessage() {}

func (x *Ioelss) ProtoReflect() protoreflect.Message {
	mi := &file_ioelss_v1_ioelss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ioelss.ProtoReflect.Descriptor instead.
func (*Ioelss) Descriptor() ([]byte, []int) {
	return file_ioelss_v1_ioelss_proto_rawDescGZIP(), []int{0}
}

func (x *Ioelss) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ioelss) GetElementName() string {
	if x != nil {
		return x.ElementName
	}
	return ""
}

func (x *Ioelss) GetElementLable() string {
	if x != nil {
		return x.ElementLable
	}
	return ""
}

func (x *Ioelss) GetElementColor() string {
	if x != nil {
		return x.ElementColor
	}
	return ""
}

func (x *Ioelss) GetEditDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EditDate
	}
	return nil
}

type Ioelsss struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           uint32                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ElementName  string                 `protobuf:"bytes,2,opt,name=element_name,json=elementName,proto3" json:"element_name,omitempty"`
	ElementLable string                 `protobuf:"bytes,3,opt,name=element_lable,json=elementLable,proto3" json:"element_lable,omitempty"`
	ElementColor string                 `protobuf:"bytes,4,opt,name=element_color,json=elementColor,proto3" json:"element_color,omitempty"`
	EditDate     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=edit_date,json=editDate,proto3" json:"edit_date,omitempty"`
}

func (x *Ioelsss) Reset() {
	*x = Ioelsss{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioelss_v1_ioelss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ioelsss) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ioelsss) ProtoMessage() {}

func (x *Ioelsss) ProtoReflect() protoreflect.Message {
	mi := &file_ioelss_v1_ioelss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ioelsss.ProtoReflect.Descriptor instead.
func (*Ioelsss) Descriptor() ([]byte, []int) {
	return file_ioelss_v1_ioelss_proto_rawDescGZIP(), []int{1}
}

func (x *Ioelsss) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Ioelsss) GetElementName() string {
	if x != nil {
		return x.ElementName
	}
	return ""
}

func (x *Ioelsss) GetElementLable() string {
	if x != nil {
		return x.ElementLable
	}
	return ""
}

func (x *Ioelsss) GetElementColor() string {
	if x != nil {
		return x.ElementColor
	}
	return ""
}

func (x *Ioelsss) GetEditDate() *timestamppb.Timestamp {
	if x != nil {
		return x.EditDate
	}
	return nil
}

type ListIoelsssRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListIoelsssRequest) Reset() {
	*x = ListIoelsssRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioelss_v1_ioelss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIoelsssRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIoelsssRequest) ProtoMessage() {}

func (x *ListIoelsssRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ioelss_v1_ioelss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIoelsssRequest.ProtoReflect.Descriptor instead.
func (*ListIoelsssRequest) Descriptor() ([]byte, []int) {
	return file_ioelss_v1_ioelss_proto_rawDescGZIP(), []int{2}
}

type ListIoelsssResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ioelssss []*Ioelsss `protobuf:"bytes,1,rep,name=ioelssss,proto3" json:"ioelssss,omitempty"`
}

func (x *ListIoelsssResponse) Reset() {
	*x = ListIoelsssResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioelss_v1_ioelss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListIoelsssResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListIoelsssResponse) ProtoMessage() {}

func (x *ListIoelsssResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ioelss_v1_ioelss_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListIoelsssResponse.ProtoReflect.Descriptor instead.
func (*ListIoelsssResponse) Descriptor() ([]byte, []int) {
	return file_ioelss_v1_ioelss_proto_rawDescGZIP(), []int{3}
}

func (x *ListIoelsssResponse) GetIoelssss() []*Ioelsss {
	if x != nil {
		return x.Ioelssss
	}
	return nil
}

type UpdateIoelssRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ioelss *Ioelss `protobuf:"bytes,1,opt,name=ioelss,proto3" json:"ioelss,omitempty"`
}

func (x *UpdateIoelssRequest) Reset() {
	*x = UpdateIoelssRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioelss_v1_ioelss_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIoelssRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIoelssRequest) ProtoMessage() {}

func (x *UpdateIoelssRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ioelss_v1_ioelss_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIoelssRequest.ProtoReflect.Descriptor instead.
func (*UpdateIoelssRequest) Descriptor() ([]byte, []int) {
	return file_ioelss_v1_ioelss_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateIoelssRequest) GetIoelss() *Ioelss {
	if x != nil {
		return x.Ioelss
	}
	return nil
}

type UpdateIoelssResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateIoelssResponse) Reset() {
	*x = UpdateIoelssResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioelss_v1_ioelss_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateIoelssResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateIoelssResponse) ProtoMessage() {}

func (x *UpdateIoelssResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ioelss_v1_ioelss_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateIoelssResponse.ProtoReflect.Descriptor instead.
func (*UpdateIoelssResponse) Descriptor() ([]byte, []int) {
	return file_ioelss_v1_ioelss_proto_rawDescGZIP(), []int{5}
}

type CreateIoelssRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ioelss *Ioelss `protobuf:"bytes,1,opt,name=ioelss,proto3" json:"ioelss,omitempty"`
}

func (x *CreateIoelssRequest) Reset() {
	*x = CreateIoelssRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioelss_v1_ioelss_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIoelssRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIoelssRequest) ProtoMessage() {}

func (x *CreateIoelssRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ioelss_v1_ioelss_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIoelssRequest.ProtoReflect.Descriptor instead.
func (*CreateIoelssRequest) Descriptor() ([]byte, []int) {
	return file_ioelss_v1_ioelss_proto_rawDescGZIP(), []int{6}
}

func (x *CreateIoelssRequest) GetIoelss() *Ioelss {
	if x != nil {
		return x.Ioelss
	}
	return nil
}

type CreateIoelssResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IoelssId uint32 `protobuf:"varint,1,opt,name=ioelss_id,json=ioelssId,proto3" json:"ioelss_id,omitempty"`
}

func (x *CreateIoelssResponse) Reset() {
	*x = CreateIoelssResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioelss_v1_ioelss_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateIoelssResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateIoelssResponse) ProtoMessage() {}

func (x *CreateIoelssResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ioelss_v1_ioelss_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateIoelssResponse.ProtoReflect.Descriptor instead.
func (*CreateIoelssResponse) Descriptor() ([]byte, []int) {
	return file_ioelss_v1_ioelss_proto_rawDescGZIP(), []int{7}
}

func (x *CreateIoelssResponse) GetIoelssId() uint32 {
	if x != nil {
		return x.IoelssId
	}
	return 0
}

type DeleteIoelssRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IoelssId uint32 `protobuf:"varint,1,opt,name=ioelss_id,json=ioelssId,proto3" json:"ioelss_id,omitempty"`
}

func (x *DeleteIoelssRequest) Reset() {
	*x = DeleteIoelssRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioelss_v1_ioelss_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIoelssRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIoelssRequest) ProtoMessage() {}

func (x *DeleteIoelssRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ioelss_v1_ioelss_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIoelssRequest.ProtoReflect.Descriptor instead.
func (*DeleteIoelssRequest) Descriptor() ([]byte, []int) {
	return file_ioelss_v1_ioelss_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteIoelssRequest) GetIoelssId() uint32 {
	if x != nil {
		return x.IoelssId
	}
	return 0
}

type DeleteIoelssResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteIoelssResponse) Reset() {
	*x = DeleteIoelssResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ioelss_v1_ioelss_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIoelssResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIoelssResponse) ProtoMessage() {}

func (x *DeleteIoelssResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ioelss_v1_ioelss_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIoelssResponse.ProtoReflect.Descriptor instead.
func (*DeleteIoelssResponse) Descriptor() ([]byte, []int) {
	return file_ioelss_v1_ioelss_proto_rawDescGZIP(), []int{9}
}

var File_ioelss_v1_ioelss_proto protoreflect.FileDescriptor

var file_ioelss_v1_ioelss_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x65, 0x6c,
	0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe,
	0x01, 0x0a, 0x06, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x64, 0x69, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74, 0x65, 0x22,
	0xbf, 0x01, 0x0a, 0x07, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6c, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x37, 0x0a, 0x09, 0x65, 0x64, 0x69, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x65, 0x64, 0x69, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6f, 0x65, 0x6c, 0x73, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x08, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6f, 0x65,
	0x6c, 0x73, 0x73, 0x73, 0x52, 0x08, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x73, 0x73, 0x22, 0x40,
	0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x06, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x06, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73,
	0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x06, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6f, 0x65, 0x6c,
	0x73, 0x73, 0x52, 0x06, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x22, 0x33, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x49, 0x64, 0x22,
	0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x6f, 0x65, 0x6c, 0x73,
	0x73, 0x49, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6f, 0x65,
	0x6c, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x8d, 0x04, 0x0a, 0x0d,
	0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x80, 0x01,
	0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x12, 0x1e,
	0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x92, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x0a, 0x01, 0x02,
	0x12, 0x80, 0x01, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73,
	0x73, 0x12, 0x1e, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x1a, 0x1a, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x92, 0xd3, 0xe4, 0x93, 0x02, 0x04, 0x0a,
	0x02, 0x02, 0x00, 0x12, 0x8a, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6f,
	0x65, 0x6c, 0x73, 0x73, 0x12, 0x1e, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x2a, 0x27, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2f, 0x7b, 0x69, 0x6f, 0x65, 0x6c,
	0x73, 0x73, 0x5f, 0x69, 0x64, 0x7d, 0x92, 0xd3, 0xe4, 0x93, 0x02, 0x04, 0x0a, 0x02, 0x02, 0x00,
	0x12, 0x69, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x73, 0x12,
	0x1d, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6f, 0x65, 0x6c, 0x73, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x94, 0x01, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x49,
	0x6f, 0x65, 0x6c, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x31, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x72, 0x69, 0x73, 0x63, 0x6f, 0x38,
	0x38, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6f, 0x65,
	0x6c, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x09, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x15,
	0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x49, 0x6f, 0x65, 0x6c, 0x73, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ioelss_v1_ioelss_proto_rawDescOnce sync.Once
	file_ioelss_v1_ioelss_proto_rawDescData = file_ioelss_v1_ioelss_proto_rawDesc
)

func file_ioelss_v1_ioelss_proto_rawDescGZIP() []byte {
	file_ioelss_v1_ioelss_proto_rawDescOnce.Do(func() {
		file_ioelss_v1_ioelss_proto_rawDescData = protoimpl.X.CompressGZIP(file_ioelss_v1_ioelss_proto_rawDescData)
	})
	return file_ioelss_v1_ioelss_proto_rawDescData
}

var file_ioelss_v1_ioelss_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_ioelss_v1_ioelss_proto_goTypes = []interface{}{
	(*Ioelss)(nil),                // 0: ioelss.v1.Ioelss
	(*Ioelsss)(nil),               // 1: ioelss.v1.Ioelsss
	(*ListIoelsssRequest)(nil),    // 2: ioelss.v1.ListIoelsssRequest
	(*ListIoelsssResponse)(nil),   // 3: ioelss.v1.ListIoelsssResponse
	(*UpdateIoelssRequest)(nil),   // 4: ioelss.v1.UpdateIoelssRequest
	(*UpdateIoelssResponse)(nil),  // 5: ioelss.v1.UpdateIoelssResponse
	(*CreateIoelssRequest)(nil),   // 6: ioelss.v1.CreateIoelssRequest
	(*CreateIoelssResponse)(nil),  // 7: ioelss.v1.CreateIoelssResponse
	(*DeleteIoelssRequest)(nil),   // 8: ioelss.v1.DeleteIoelssRequest
	(*DeleteIoelssResponse)(nil),  // 9: ioelss.v1.DeleteIoelssResponse
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_ioelss_v1_ioelss_proto_depIdxs = []int32{
	10, // 0: ioelss.v1.Ioelss.edit_date:type_name -> google.protobuf.Timestamp
	10, // 1: ioelss.v1.Ioelsss.edit_date:type_name -> google.protobuf.Timestamp
	1,  // 2: ioelss.v1.ListIoelsssResponse.ioelssss:type_name -> ioelss.v1.Ioelsss
	0,  // 3: ioelss.v1.UpdateIoelssRequest.ioelss:type_name -> ioelss.v1.Ioelss
	0,  // 4: ioelss.v1.CreateIoelssRequest.ioelss:type_name -> ioelss.v1.Ioelss
	6,  // 5: ioelss.v1.IoelssService.CreateIoelss:input_type -> ioelss.v1.CreateIoelssRequest
	4,  // 6: ioelss.v1.IoelssService.UpdateIoelss:input_type -> ioelss.v1.UpdateIoelssRequest
	8,  // 7: ioelss.v1.IoelssService.DeleteIoelss:input_type -> ioelss.v1.DeleteIoelssRequest
	2,  // 8: ioelss.v1.IoelssService.ListIoelsss:input_type -> ioelss.v1.ListIoelsssRequest
	7,  // 9: ioelss.v1.IoelssService.CreateIoelss:output_type -> ioelss.v1.CreateIoelssResponse
	5,  // 10: ioelss.v1.IoelssService.UpdateIoelss:output_type -> ioelss.v1.UpdateIoelssResponse
	9,  // 11: ioelss.v1.IoelssService.DeleteIoelss:output_type -> ioelss.v1.DeleteIoelssResponse
	3,  // 12: ioelss.v1.IoelssService.ListIoelsss:output_type -> ioelss.v1.ListIoelsssResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_ioelss_v1_ioelss_proto_init() }
func file_ioelss_v1_ioelss_proto_init() {
	if File_ioelss_v1_ioelss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ioelss_v1_ioelss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ioelss); i {
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
		file_ioelss_v1_ioelss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ioelsss); i {
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
		file_ioelss_v1_ioelss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIoelsssRequest); i {
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
		file_ioelss_v1_ioelss_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListIoelsssResponse); i {
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
		file_ioelss_v1_ioelss_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIoelssRequest); i {
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
		file_ioelss_v1_ioelss_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateIoelssResponse); i {
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
		file_ioelss_v1_ioelss_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIoelssRequest); i {
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
		file_ioelss_v1_ioelss_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateIoelssResponse); i {
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
		file_ioelss_v1_ioelss_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIoelssRequest); i {
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
		file_ioelss_v1_ioelss_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIoelssResponse); i {
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
			RawDescriptor: file_ioelss_v1_ioelss_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ioelss_v1_ioelss_proto_goTypes,
		DependencyIndexes: file_ioelss_v1_ioelss_proto_depIdxs,
		MessageInfos:      file_ioelss_v1_ioelss_proto_msgTypes,
	}.Build()
	File_ioelss_v1_ioelss_proto = out.File
	file_ioelss_v1_ioelss_proto_rawDesc = nil
	file_ioelss_v1_ioelss_proto_goTypes = nil
	file_ioelss_v1_ioelss_proto_depIdxs = nil
}
