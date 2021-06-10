// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: api/billing.proto

package rpc

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

type IOMetricsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         *UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`                  // UUID for Customer to charge (bill)
	Reads      int64 `protobuf:"varint,4,opt,name=reads,proto3" json:"reads,omitempty"`           // Reads is a counter of number of reads
	Writes     int64 `protobuf:"varint,5,opt,name=writes,proto3" json:"writes,omitempty"`         // Writes is a counter of number of writes
	ReadBytes  int64 `protobuf:"varint,6,opt,name=readBytes,proto3" json:"readBytes,omitempty"`   // readBytes is total number of bytes read
	WriteBytes int64 `protobuf:"varint,7,opt,name=writeBytes,proto3" json:"writeBytes,omitempty"` // writeBytes is total number of bytes written
}

func (x *IOMetricsReq) Reset() {
	*x = IOMetricsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_billing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IOMetricsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IOMetricsReq) ProtoMessage() {}

func (x *IOMetricsReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_billing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IOMetricsReq.ProtoReflect.Descriptor instead.
func (*IOMetricsReq) Descriptor() ([]byte, []int) {
	return file_api_billing_proto_rawDescGZIP(), []int{0}
}

func (x *IOMetricsReq) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *IOMetricsReq) GetReads() int64 {
	if x != nil {
		return x.Reads
	}
	return 0
}

func (x *IOMetricsReq) GetWrites() int64 {
	if x != nil {
		return x.Writes
	}
	return 0
}

func (x *IOMetricsReq) GetReadBytes() int64 {
	if x != nil {
		return x.ReadBytes
	}
	return 0
}

func (x *IOMetricsReq) GetWriteBytes() int64 {
	if x != nil {
		return x.WriteBytes
	}
	return 0
}

type CustomerBilledResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  *UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // UUID for Customer to charge (bill)
	UnitsBilled         int64 `protobuf:"varint,2,opt,name=unitsBilled,proto3" json:"unitsBilled,omitempty"`
	PerUnitCost         int64 `protobuf:"varint,3,opt,name=perUnitCost,proto3" json:"perUnitCost,omitempty"`
	TotalMonthlyCharged int64 `protobuf:"varint,4,opt,name=totalMonthlyCharged,proto3" json:"totalMonthlyCharged,omitempty"`
}

func (x *CustomerBilledResp) Reset() {
	*x = CustomerBilledResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_billing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerBilledResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerBilledResp) ProtoMessage() {}

func (x *CustomerBilledResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_billing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerBilledResp.ProtoReflect.Descriptor instead.
func (*CustomerBilledResp) Descriptor() ([]byte, []int) {
	return file_api_billing_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerBilledResp) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *CustomerBilledResp) GetUnitsBilled() int64 {
	if x != nil {
		return x.UnitsBilled
	}
	return 0
}

func (x *CustomerBilledResp) GetPerUnitCost() int64 {
	if x != nil {
		return x.PerUnitCost
	}
	return 0
}

func (x *CustomerBilledResp) GetTotalMonthlyCharged() int64 {
	if x != nil {
		return x.TotalMonthlyCharged
	}
	return 0
}

type UUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *UUID) Reset() {
	*x = UUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_billing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UUID) ProtoMessage() {}

func (x *UUID) ProtoReflect() protoreflect.Message {
	mi := &file_api_billing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UUID.ProtoReflect.Descriptor instead.
func (*UUID) Descriptor() ([]byte, []int) {
	return file_api_billing_proto_rawDescGZIP(), []int{2}
}

func (x *UUID) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_api_billing_proto protoreflect.FileDescriptor

var file_api_billing_proto_rawDesc = []byte{
	0x0a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x72, 0x72, 0x75, 0x73, 0x6d,
	0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0xa6, 0x01, 0x0a, 0x0c, 0x49, 0x4f,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x72,
	0x72, 0x75, 0x73, 0x6d, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x42, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x77, 0x72, 0x69, 0x74, 0x65, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42,
	0x69, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x72, 0x72,
	0x75, 0x73, 0x6d, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x55, 0x49,
	0x44, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x42, 0x69,
	0x6c, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x75, 0x6e, 0x69, 0x74,
	0x73, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x55, 0x6e,
	0x69, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x55, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x13, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4d, 0x6f, 0x6e,
	0x74, 0x68, 0x6c, 0x79, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x64, 0x22, 0x1c, 0x0a, 0x04, 0x55,
	0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x67, 0x0a, 0x07, 0x42, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x12, 0x5c, 0x0a, 0x0c, 0x41, 0x64, 0x64, 0x49, 0x4f, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x69, 0x72, 0x72, 0x75,
	0x73, 0x6d, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x4f, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x28, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x69, 0x72, 0x72, 0x75, 0x73, 0x6d, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x09, 0x5a, 0x07, 0x70, 0x6b, 0x67, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_billing_proto_rawDescOnce sync.Once
	file_api_billing_proto_rawDescData = file_api_billing_proto_rawDesc
)

func file_api_billing_proto_rawDescGZIP() []byte {
	file_api_billing_proto_rawDescOnce.Do(func() {
		file_api_billing_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_billing_proto_rawDescData)
	})
	return file_api_billing_proto_rawDescData
}

var file_api_billing_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_billing_proto_goTypes = []interface{}{
	(*IOMetricsReq)(nil),       // 0: com.cirrusmd.billing.IOMetricsReq
	(*CustomerBilledResp)(nil), // 1: com.cirrusmd.billing.CustomerBilledResp
	(*UUID)(nil),               // 2: com.cirrusmd.billing.UUID
}
var file_api_billing_proto_depIdxs = []int32{
	2, // 0: com.cirrusmd.billing.IOMetricsReq.id:type_name -> com.cirrusmd.billing.UUID
	2, // 1: com.cirrusmd.billing.CustomerBilledResp.id:type_name -> com.cirrusmd.billing.UUID
	0, // 2: com.cirrusmd.billing.Billing.AddIOMetrics:input_type -> com.cirrusmd.billing.IOMetricsReq
	1, // 3: com.cirrusmd.billing.Billing.AddIOMetrics:output_type -> com.cirrusmd.billing.CustomerBilledResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_billing_proto_init() }
func file_api_billing_proto_init() {
	if File_api_billing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_billing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IOMetricsReq); i {
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
		file_api_billing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerBilledResp); i {
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
		file_api_billing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UUID); i {
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
			RawDescriptor: file_api_billing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_billing_proto_goTypes,
		DependencyIndexes: file_api_billing_proto_depIdxs,
		MessageInfos:      file_api_billing_proto_msgTypes,
	}.Build()
	File_api_billing_proto = out.File
	file_api_billing_proto_rawDesc = nil
	file_api_billing_proto_goTypes = nil
	file_api_billing_proto_depIdxs = nil
}
