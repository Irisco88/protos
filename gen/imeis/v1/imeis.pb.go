// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: imeis/v1/imeis.proto

package imeisv1

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

type Imeis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ImeisName      string `protobuf:"bytes,2,opt,name=imeis_name,json=imeisName,proto3" json:"imeis_name,omitempty"`
	DeviceName     string `protobuf:"bytes,3,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	TestVehicle    string `protobuf:"bytes,4,opt,name=test_vehicle,json=testVehicle,proto3" json:"test_vehicle,omitempty"`
	BodyNumber     string `protobuf:"bytes,5,opt,name=body_number,json=bodyNumber,proto3" json:"body_number,omitempty"`
	ProjectRelated string `protobuf:"bytes,6,opt,name=project_related,json=projectRelated,proto3" json:"project_related,omitempty"`
	DriverName     string `protobuf:"bytes,7,opt,name=driver_name,json=driverName,proto3" json:"driver_name,omitempty"`
	Laboratory     string `protobuf:"bytes,8,opt,name=laboratory,proto3" json:"laboratory,omitempty"`
	TestName       string `protobuf:"bytes,9,opt,name=test_name,json=testName,proto3" json:"test_name,omitempty"`
	TestLocation   string `protobuf:"bytes,10,opt,name=test_location,json=testLocation,proto3" json:"test_location,omitempty"`
	StartDate      string `protobuf:"bytes,11,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	EndDate        string `protobuf:"bytes,12,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty"`
	StandardNumber string `protobuf:"bytes,13,opt,name=standard_number,json=standardNumber,proto3" json:"standard_number,omitempty"`
}

func (x *Imeis) Reset() {
	*x = Imeis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imeis_v1_imeis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Imeis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Imeis) ProtoMessage() {}

func (x *Imeis) ProtoReflect() protoreflect.Message {
	mi := &file_imeis_v1_imeis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Imeis.ProtoReflect.Descriptor instead.
func (*Imeis) Descriptor() ([]byte, []int) {
	return file_imeis_v1_imeis_proto_rawDescGZIP(), []int{0}
}

func (x *Imeis) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Imeis) GetImeisName() string {
	if x != nil {
		return x.ImeisName
	}
	return ""
}

func (x *Imeis) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *Imeis) GetTestVehicle() string {
	if x != nil {
		return x.TestVehicle
	}
	return ""
}

func (x *Imeis) GetBodyNumber() string {
	if x != nil {
		return x.BodyNumber
	}
	return ""
}

func (x *Imeis) GetProjectRelated() string {
	if x != nil {
		return x.ProjectRelated
	}
	return ""
}

func (x *Imeis) GetDriverName() string {
	if x != nil {
		return x.DriverName
	}
	return ""
}

func (x *Imeis) GetLaboratory() string {
	if x != nil {
		return x.Laboratory
	}
	return ""
}

func (x *Imeis) GetTestName() string {
	if x != nil {
		return x.TestName
	}
	return ""
}

func (x *Imeis) GetTestLocation() string {
	if x != nil {
		return x.TestLocation
	}
	return ""
}

func (x *Imeis) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Imeis) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Imeis) GetStandardNumber() string {
	if x != nil {
		return x.StandardNumber
	}
	return ""
}

type ListImeissRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListImeissRequest) Reset() {
	*x = ListImeissRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imeis_v1_imeis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListImeissRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImeissRequest) ProtoMessage() {}

func (x *ListImeissRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imeis_v1_imeis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImeissRequest.ProtoReflect.Descriptor instead.
func (*ListImeissRequest) Descriptor() ([]byte, []int) {
	return file_imeis_v1_imeis_proto_rawDescGZIP(), []int{1}
}

type ListImeissResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imeiss []*Imeis `protobuf:"bytes,1,rep,name=imeiss,proto3" json:"imeiss,omitempty"`
}

func (x *ListImeissResponse) Reset() {
	*x = ListImeissResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imeis_v1_imeis_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListImeissResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListImeissResponse) ProtoMessage() {}

func (x *ListImeissResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imeis_v1_imeis_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListImeissResponse.ProtoReflect.Descriptor instead.
func (*ListImeissResponse) Descriptor() ([]byte, []int) {
	return file_imeis_v1_imeis_proto_rawDescGZIP(), []int{2}
}

func (x *ListImeissResponse) GetImeiss() []*Imeis {
	if x != nil {
		return x.Imeiss
	}
	return nil
}

type UpdateImeisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imeis *Imeis `protobuf:"bytes,1,opt,name=imeis,proto3" json:"imeis,omitempty"`
}

func (x *UpdateImeisRequest) Reset() {
	*x = UpdateImeisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imeis_v1_imeis_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateImeisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateImeisRequest) ProtoMessage() {}

func (x *UpdateImeisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imeis_v1_imeis_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateImeisRequest.ProtoReflect.Descriptor instead.
func (*UpdateImeisRequest) Descriptor() ([]byte, []int) {
	return file_imeis_v1_imeis_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateImeisRequest) GetImeis() *Imeis {
	if x != nil {
		return x.Imeis
	}
	return nil
}

type UpdateImeisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateImeisResponse) Reset() {
	*x = UpdateImeisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imeis_v1_imeis_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateImeisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateImeisResponse) ProtoMessage() {}

func (x *UpdateImeisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imeis_v1_imeis_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateImeisResponse.ProtoReflect.Descriptor instead.
func (*UpdateImeisResponse) Descriptor() ([]byte, []int) {
	return file_imeis_v1_imeis_proto_rawDescGZIP(), []int{4}
}

type CreateImeisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imeis *Imeis `protobuf:"bytes,1,opt,name=imeis,proto3" json:"imeis,omitempty"`
}

func (x *CreateImeisRequest) Reset() {
	*x = CreateImeisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imeis_v1_imeis_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImeisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImeisRequest) ProtoMessage() {}

func (x *CreateImeisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imeis_v1_imeis_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImeisRequest.ProtoReflect.Descriptor instead.
func (*CreateImeisRequest) Descriptor() ([]byte, []int) {
	return file_imeis_v1_imeis_proto_rawDescGZIP(), []int{5}
}

func (x *CreateImeisRequest) GetImeis() *Imeis {
	if x != nil {
		return x.Imeis
	}
	return nil
}

type CreateImeisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImeisId uint32 `protobuf:"varint,1,opt,name=imeis_id,json=imeisId,proto3" json:"imeis_id,omitempty"`
}

func (x *CreateImeisResponse) Reset() {
	*x = CreateImeisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imeis_v1_imeis_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateImeisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateImeisResponse) ProtoMessage() {}

func (x *CreateImeisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imeis_v1_imeis_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateImeisResponse.ProtoReflect.Descriptor instead.
func (*CreateImeisResponse) Descriptor() ([]byte, []int) {
	return file_imeis_v1_imeis_proto_rawDescGZIP(), []int{6}
}

func (x *CreateImeisResponse) GetImeisId() uint32 {
	if x != nil {
		return x.ImeisId
	}
	return 0
}

type DeleteImeisRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImeisId uint32 `protobuf:"varint,1,opt,name=imeis_id,json=imeisId,proto3" json:"imeis_id,omitempty"`
}

func (x *DeleteImeisRequest) Reset() {
	*x = DeleteImeisRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imeis_v1_imeis_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImeisRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImeisRequest) ProtoMessage() {}

func (x *DeleteImeisRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imeis_v1_imeis_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImeisRequest.ProtoReflect.Descriptor instead.
func (*DeleteImeisRequest) Descriptor() ([]byte, []int) {
	return file_imeis_v1_imeis_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteImeisRequest) GetImeisId() uint32 {
	if x != nil {
		return x.ImeisId
	}
	return 0
}

type DeleteImeisResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteImeisResponse) Reset() {
	*x = DeleteImeisResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imeis_v1_imeis_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteImeisResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteImeisResponse) ProtoMessage() {}

func (x *DeleteImeisResponse) ProtoReflect() protoreflect.Message {
	mi := &file_imeis_v1_imeis_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteImeisResponse.ProtoReflect.Descriptor instead.
func (*DeleteImeisResponse) Descriptor() ([]byte, []int) {
	return file_imeis_v1_imeis_proto_rawDescGZIP(), []int{8}
}

var File_imeis_v1_imeis_proto protoreflect.FileDescriptor

var file_imeis_v1_imeis_proto_rawDesc = []byte{
	0x0a, 0x14, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x65, 0x69, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x03, 0x0a, 0x05, 0x49,
	0x6d, 0x65, 0x69, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x68,
	0x69, 0x63, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x65, 0x73, 0x74,
	0x56, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x6f, 0x64, 0x79, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6f,
	0x64, 0x79, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x61, 0x62, 0x6f, 0x72, 0x61, 0x74, 0x6f,
	0x72, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61,
	0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x27,
	0x0a, 0x0f, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72,
	0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6d, 0x65, 0x69, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d,
	0x65, 0x69, 0x73, 0x52, 0x06, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x73, 0x22, 0x3b, 0x0a, 0x12, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x65, 0x69,
	0x73, 0x52, 0x05, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x3b, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x05, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x22, 0x30, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x49, 0x64, 0x22, 0x2f,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x49, 0x64, 0x22,
	0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xed, 0x03, 0x0a, 0x0c, 0x49, 0x6d, 0x65, 0x69, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x12, 0x1c, 0x2e, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x22, 0x18,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x69, 0x6d, 0x65, 0x69, 0x92, 0xd3, 0xe4, 0x93, 0x02, 0x03, 0x0a, 0x01,
	0x02, 0x12, 0x79, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73,
	0x12, 0x1c, 0x2e, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6d, 0x65, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x3a, 0x01, 0x2a, 0x1a, 0x18, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x69, 0x6d,
	0x65, 0x69, 0x92, 0xd3, 0xe4, 0x93, 0x02, 0x04, 0x0a, 0x02, 0x02, 0x00, 0x12, 0x82, 0x01, 0x0a,
	0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x12, 0x1c, 0x2e, 0x69,
	0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d,
	0x65, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6d, 0x65,
	0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6d, 0x65, 0x69,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x26, 0x2a, 0x24, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x65, 0x69, 0x73,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2f, 0x7b, 0x69, 0x6d,
	0x65, 0x69, 0x73, 0x5f, 0x69, 0x64, 0x7d, 0x92, 0xd3, 0xe4, 0x93, 0x02, 0x04, 0x0a, 0x02, 0x02,
	0x00, 0x12, 0x63, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x73, 0x12,
	0x1b, 0x2e, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6d, 0x65, 0x69, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69,
	0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6d, 0x65, 0x69,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x65, 0x69,
	0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x8c, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x6d, 0x65, 0x69, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x72, 0x69, 0x73, 0x63, 0x6f, 0x38, 0x38, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x6d, 0x65, 0x69, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x69,
	0x6d, 0x65, 0x69, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x49,
	0x6d, 0x65, 0x69, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x14, 0x49, 0x6d, 0x65, 0x69, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x49, 0x6d, 0x65, 0x69,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imeis_v1_imeis_proto_rawDescOnce sync.Once
	file_imeis_v1_imeis_proto_rawDescData = file_imeis_v1_imeis_proto_rawDesc
)

func file_imeis_v1_imeis_proto_rawDescGZIP() []byte {
	file_imeis_v1_imeis_proto_rawDescOnce.Do(func() {
		file_imeis_v1_imeis_proto_rawDescData = protoimpl.X.CompressGZIP(file_imeis_v1_imeis_proto_rawDescData)
	})
	return file_imeis_v1_imeis_proto_rawDescData
}

var file_imeis_v1_imeis_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_imeis_v1_imeis_proto_goTypes = []interface{}{
	(*Imeis)(nil),               // 0: imeis.v1.Imeis
	(*ListImeissRequest)(nil),   // 1: imeis.v1.ListImeissRequest
	(*ListImeissResponse)(nil),  // 2: imeis.v1.ListImeissResponse
	(*UpdateImeisRequest)(nil),  // 3: imeis.v1.UpdateImeisRequest
	(*UpdateImeisResponse)(nil), // 4: imeis.v1.UpdateImeisResponse
	(*CreateImeisRequest)(nil),  // 5: imeis.v1.CreateImeisRequest
	(*CreateImeisResponse)(nil), // 6: imeis.v1.CreateImeisResponse
	(*DeleteImeisRequest)(nil),  // 7: imeis.v1.DeleteImeisRequest
	(*DeleteImeisResponse)(nil), // 8: imeis.v1.DeleteImeisResponse
}
var file_imeis_v1_imeis_proto_depIdxs = []int32{
	0, // 0: imeis.v1.ListImeissResponse.imeiss:type_name -> imeis.v1.Imeis
	0, // 1: imeis.v1.UpdateImeisRequest.imeis:type_name -> imeis.v1.Imeis
	0, // 2: imeis.v1.CreateImeisRequest.imeis:type_name -> imeis.v1.Imeis
	5, // 3: imeis.v1.ImeisService.CreateImeis:input_type -> imeis.v1.CreateImeisRequest
	3, // 4: imeis.v1.ImeisService.UpdateImeis:input_type -> imeis.v1.UpdateImeisRequest
	7, // 5: imeis.v1.ImeisService.DeleteImeis:input_type -> imeis.v1.DeleteImeisRequest
	1, // 6: imeis.v1.ImeisService.ListImeiss:input_type -> imeis.v1.ListImeissRequest
	6, // 7: imeis.v1.ImeisService.CreateImeis:output_type -> imeis.v1.CreateImeisResponse
	4, // 8: imeis.v1.ImeisService.UpdateImeis:output_type -> imeis.v1.UpdateImeisResponse
	8, // 9: imeis.v1.ImeisService.DeleteImeis:output_type -> imeis.v1.DeleteImeisResponse
	2, // 10: imeis.v1.ImeisService.ListImeiss:output_type -> imeis.v1.ListImeissResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_imeis_v1_imeis_proto_init() }
func file_imeis_v1_imeis_proto_init() {
	if File_imeis_v1_imeis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imeis_v1_imeis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Imeis); i {
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
		file_imeis_v1_imeis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListImeissRequest); i {
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
		file_imeis_v1_imeis_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListImeissResponse); i {
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
		file_imeis_v1_imeis_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateImeisRequest); i {
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
		file_imeis_v1_imeis_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateImeisResponse); i {
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
		file_imeis_v1_imeis_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImeisRequest); i {
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
		file_imeis_v1_imeis_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateImeisResponse); i {
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
		file_imeis_v1_imeis_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImeisRequest); i {
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
		file_imeis_v1_imeis_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteImeisResponse); i {
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
			RawDescriptor: file_imeis_v1_imeis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imeis_v1_imeis_proto_goTypes,
		DependencyIndexes: file_imeis_v1_imeis_proto_depIdxs,
		MessageInfos:      file_imeis_v1_imeis_proto_msgTypes,
	}.Build()
	File_imeis_v1_imeis_proto = out.File
	file_imeis_v1_imeis_proto_rawDesc = nil
	file_imeis_v1_imeis_proto_goTypes = nil
	file_imeis_v1_imeis_proto_depIdxs = nil
}
