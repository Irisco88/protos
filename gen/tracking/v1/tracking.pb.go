// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: tracking/v1/tracking.proto

package trackingv1

import (
	v1 "github.com/openfms/protos/gen/device/v1"
	_ "github.com/openfms/protos/gen/google/api"
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

type LiveDevicesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ImeiList []string `protobuf:"bytes,1,rep,name=imei_list,json=imeiList,proto3" json:"imei_list,omitempty"`
}

func (x *LiveDevicesRequest) Reset() {
	*x = LiveDevicesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_v1_tracking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveDevicesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveDevicesRequest) ProtoMessage() {}

func (x *LiveDevicesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_v1_tracking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveDevicesRequest.ProtoReflect.Descriptor instead.
func (*LiveDevicesRequest) Descriptor() ([]byte, []int) {
	return file_tracking_v1_tracking_proto_rawDescGZIP(), []int{0}
}

func (x *LiveDevicesRequest) GetImeiList() []string {
	if x != nil {
		return x.ImeiList
	}
	return nil
}

type LiveDevicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Point *v1.AVLData `protobuf:"bytes,1,opt,name=point,proto3" json:"point,omitempty"`
}

func (x *LiveDevicesResponse) Reset() {
	*x = LiveDevicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_v1_tracking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LiveDevicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LiveDevicesResponse) ProtoMessage() {}

func (x *LiveDevicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_v1_tracking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LiveDevicesResponse.ProtoReflect.Descriptor instead.
func (*LiveDevicesResponse) Descriptor() ([]byte, []int) {
	return file_tracking_v1_tracking_proto_rawDescGZIP(), []int{1}
}

func (x *LiveDevicesResponse) GetPoint() *v1.AVLData {
	if x != nil {
		return x.Point
	}
	return nil
}

var File_tracking_v1_tracking_proto protoreflect.FileDescriptor

var file_tracking_v1_tracking_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x31, 0x0a, 0x12, 0x4c, 0x69, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x65, 0x69, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x65, 0x69, 0x4c, 0x69,
	0x73, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x4c, 0x69, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x56, 0x4c, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x32, 0x84, 0x01, 0x0a, 0x0f, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x71, 0x0a, 0x0b, 0x4c, 0x69, 0x76, 0x65, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x76, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x15, 0x12, 0x13, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x28, 0x01, 0x30, 0x01, 0x42, 0xa3, 0x01, 0x0a, 0x0f, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x42, 0x0d,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e,
	0x66, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74,
	0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x69, 0x6e, 0x67, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0b, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x17, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69,
	0x6e, 0x67, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x0c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracking_v1_tracking_proto_rawDescOnce sync.Once
	file_tracking_v1_tracking_proto_rawDescData = file_tracking_v1_tracking_proto_rawDesc
)

func file_tracking_v1_tracking_proto_rawDescGZIP() []byte {
	file_tracking_v1_tracking_proto_rawDescOnce.Do(func() {
		file_tracking_v1_tracking_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracking_v1_tracking_proto_rawDescData)
	})
	return file_tracking_v1_tracking_proto_rawDescData
}

var file_tracking_v1_tracking_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tracking_v1_tracking_proto_goTypes = []interface{}{
	(*LiveDevicesRequest)(nil),  // 0: tracking.v1.LiveDevicesRequest
	(*LiveDevicesResponse)(nil), // 1: tracking.v1.LiveDevicesResponse
	(*v1.AVLData)(nil),          // 2: device.v1.AVLData
}
var file_tracking_v1_tracking_proto_depIdxs = []int32{
	2, // 0: tracking.v1.LiveDevicesResponse.point:type_name -> device.v1.AVLData
	0, // 1: tracking.v1.TrackingService.LiveDevices:input_type -> tracking.v1.LiveDevicesRequest
	1, // 2: tracking.v1.TrackingService.LiveDevices:output_type -> tracking.v1.LiveDevicesResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tracking_v1_tracking_proto_init() }
func file_tracking_v1_tracking_proto_init() {
	if File_tracking_v1_tracking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracking_v1_tracking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveDevicesRequest); i {
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
		file_tracking_v1_tracking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LiveDevicesResponse); i {
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
			RawDescriptor: file_tracking_v1_tracking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tracking_v1_tracking_proto_goTypes,
		DependencyIndexes: file_tracking_v1_tracking_proto_depIdxs,
		MessageInfos:      file_tracking_v1_tracking_proto_msgTypes,
	}.Build()
	File_tracking_v1_tracking_proto = out.File
	file_tracking_v1_tracking_proto_rawDesc = nil
	file_tracking_v1_tracking_proto_goTypes = nil
	file_tracking_v1_tracking_proto_depIdxs = nil
}
