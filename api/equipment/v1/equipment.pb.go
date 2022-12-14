// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: api/equipment/v1/equipment.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateEquipmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateEquipmentRequest) Reset() {
	*x = CreateEquipmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_equipment_v1_equipment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEquipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEquipmentRequest) ProtoMessage() {}

func (x *CreateEquipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_equipment_v1_equipment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEquipmentRequest.ProtoReflect.Descriptor instead.
func (*CreateEquipmentRequest) Descriptor() ([]byte, []int) {
	return file_api_equipment_v1_equipment_proto_rawDescGZIP(), []int{0}
}

type CreateEquipmentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateEquipmentReply) Reset() {
	*x = CreateEquipmentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_equipment_v1_equipment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEquipmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEquipmentReply) ProtoMessage() {}

func (x *CreateEquipmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_equipment_v1_equipment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEquipmentReply.ProtoReflect.Descriptor instead.
func (*CreateEquipmentReply) Descriptor() ([]byte, []int) {
	return file_api_equipment_v1_equipment_proto_rawDescGZIP(), []int{1}
}

type UpdateEquipmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateEquipmentRequest) Reset() {
	*x = UpdateEquipmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_equipment_v1_equipment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEquipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEquipmentRequest) ProtoMessage() {}

func (x *UpdateEquipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_equipment_v1_equipment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEquipmentRequest.ProtoReflect.Descriptor instead.
func (*UpdateEquipmentRequest) Descriptor() ([]byte, []int) {
	return file_api_equipment_v1_equipment_proto_rawDescGZIP(), []int{2}
}

type UpdateEquipmentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateEquipmentReply) Reset() {
	*x = UpdateEquipmentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_equipment_v1_equipment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEquipmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEquipmentReply) ProtoMessage() {}

func (x *UpdateEquipmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_equipment_v1_equipment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEquipmentReply.ProtoReflect.Descriptor instead.
func (*UpdateEquipmentReply) Descriptor() ([]byte, []int) {
	return file_api_equipment_v1_equipment_proto_rawDescGZIP(), []int{3}
}

type DeleteEquipmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEquipmentRequest) Reset() {
	*x = DeleteEquipmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_equipment_v1_equipment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEquipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEquipmentRequest) ProtoMessage() {}

func (x *DeleteEquipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_equipment_v1_equipment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEquipmentRequest.ProtoReflect.Descriptor instead.
func (*DeleteEquipmentRequest) Descriptor() ([]byte, []int) {
	return file_api_equipment_v1_equipment_proto_rawDescGZIP(), []int{4}
}

type DeleteEquipmentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEquipmentReply) Reset() {
	*x = DeleteEquipmentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_equipment_v1_equipment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEquipmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEquipmentReply) ProtoMessage() {}

func (x *DeleteEquipmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_equipment_v1_equipment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEquipmentReply.ProtoReflect.Descriptor instead.
func (*DeleteEquipmentReply) Descriptor() ([]byte, []int) {
	return file_api_equipment_v1_equipment_proto_rawDescGZIP(), []int{5}
}

type GetEquipmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EquipId int64 `protobuf:"varint,1,opt,name=equip_id,json=equipId,proto3" json:"equip_id,omitempty"`
}

func (x *GetEquipmentRequest) Reset() {
	*x = GetEquipmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_equipment_v1_equipment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEquipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEquipmentRequest) ProtoMessage() {}

func (x *GetEquipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_equipment_v1_equipment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEquipmentRequest.ProtoReflect.Descriptor instead.
func (*GetEquipmentRequest) Descriptor() ([]byte, []int) {
	return file_api_equipment_v1_equipment_proto_rawDescGZIP(), []int{6}
}

func (x *GetEquipmentRequest) GetEquipId() int64 {
	if x != nil {
		return x.EquipId
	}
	return 0
}

type GetEquipmentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EquipId   int64  `protobuf:"varint,1,opt,name=equip_id,json=equipId,proto3" json:"equip_id,omitempty"`
	EquipName string `protobuf:"bytes,2,opt,name=equip_name,json=equipName,proto3" json:"equip_name,omitempty"`
}

func (x *GetEquipmentReply) Reset() {
	*x = GetEquipmentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_equipment_v1_equipment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEquipmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEquipmentReply) ProtoMessage() {}

func (x *GetEquipmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_equipment_v1_equipment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEquipmentReply.ProtoReflect.Descriptor instead.
func (*GetEquipmentReply) Descriptor() ([]byte, []int) {
	return file_api_equipment_v1_equipment_proto_rawDescGZIP(), []int{7}
}

func (x *GetEquipmentReply) GetEquipId() int64 {
	if x != nil {
		return x.EquipId
	}
	return 0
}

func (x *GetEquipmentReply) GetEquipName() string {
	if x != nil {
		return x.EquipName
	}
	return ""
}

type ListEquipmentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListEquipmentRequest) Reset() {
	*x = ListEquipmentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_equipment_v1_equipment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEquipmentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEquipmentRequest) ProtoMessage() {}

func (x *ListEquipmentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_equipment_v1_equipment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEquipmentRequest.ProtoReflect.Descriptor instead.
func (*ListEquipmentRequest) Descriptor() ([]byte, []int) {
	return file_api_equipment_v1_equipment_proto_rawDescGZIP(), []int{8}
}

type ListEquipmentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListEquipmentReply) Reset() {
	*x = ListEquipmentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_equipment_v1_equipment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEquipmentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEquipmentReply) ProtoMessage() {}

func (x *ListEquipmentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_equipment_v1_equipment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEquipmentReply.ProtoReflect.Descriptor instead.
func (*ListEquipmentReply) Descriptor() ([]byte, []int) {
	return file_api_equipment_v1_equipment_proto_rawDescGZIP(), []int{9}
}

var File_api_equipment_v1_equipment_proto protoreflect.FileDescriptor

var file_api_equipment_v1_equipment_proto_rawDesc = []byte{
	0x0a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69,
	0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x30, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x49, 0x64, 0x22, 0x4d, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x71, 0x75, 0x69,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x32, 0xec, 0x03, 0x0a, 0x09, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x5b, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5b,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x24, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x71, 0x75,
	0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5b, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x24,
	0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x71, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x45,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x71, 0x75, 0x69, 0x70,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x65, 0x71,
	0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x71,
	0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x7b, 0x65, 0x71, 0x75, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x55, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x22, 0x2e, 0x65,
	0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x20, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x31, 0x0a, 0x10, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x1b, 0x68, 0x61, 0x72, 0x76, 0x65, 0x73,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x71, 0x75, 0x69, 0x70, 0x6d, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_equipment_v1_equipment_proto_rawDescOnce sync.Once
	file_api_equipment_v1_equipment_proto_rawDescData = file_api_equipment_v1_equipment_proto_rawDesc
)

func file_api_equipment_v1_equipment_proto_rawDescGZIP() []byte {
	file_api_equipment_v1_equipment_proto_rawDescOnce.Do(func() {
		file_api_equipment_v1_equipment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_equipment_v1_equipment_proto_rawDescData)
	})
	return file_api_equipment_v1_equipment_proto_rawDescData
}

var file_api_equipment_v1_equipment_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_equipment_v1_equipment_proto_goTypes = []interface{}{
	(*CreateEquipmentRequest)(nil), // 0: equipment.v1.CreateEquipmentRequest
	(*CreateEquipmentReply)(nil),   // 1: equipment.v1.CreateEquipmentReply
	(*UpdateEquipmentRequest)(nil), // 2: equipment.v1.UpdateEquipmentRequest
	(*UpdateEquipmentReply)(nil),   // 3: equipment.v1.UpdateEquipmentReply
	(*DeleteEquipmentRequest)(nil), // 4: equipment.v1.DeleteEquipmentRequest
	(*DeleteEquipmentReply)(nil),   // 5: equipment.v1.DeleteEquipmentReply
	(*GetEquipmentRequest)(nil),    // 6: equipment.v1.GetEquipmentRequest
	(*GetEquipmentReply)(nil),      // 7: equipment.v1.GetEquipmentReply
	(*ListEquipmentRequest)(nil),   // 8: equipment.v1.ListEquipmentRequest
	(*ListEquipmentReply)(nil),     // 9: equipment.v1.ListEquipmentReply
}
var file_api_equipment_v1_equipment_proto_depIdxs = []int32{
	0, // 0: equipment.v1.Equipment.CreateEquipment:input_type -> equipment.v1.CreateEquipmentRequest
	2, // 1: equipment.v1.Equipment.UpdateEquipment:input_type -> equipment.v1.UpdateEquipmentRequest
	4, // 2: equipment.v1.Equipment.DeleteEquipment:input_type -> equipment.v1.DeleteEquipmentRequest
	6, // 3: equipment.v1.Equipment.GetEquipment:input_type -> equipment.v1.GetEquipmentRequest
	8, // 4: equipment.v1.Equipment.ListEquipment:input_type -> equipment.v1.ListEquipmentRequest
	1, // 5: equipment.v1.Equipment.CreateEquipment:output_type -> equipment.v1.CreateEquipmentReply
	3, // 6: equipment.v1.Equipment.UpdateEquipment:output_type -> equipment.v1.UpdateEquipmentReply
	5, // 7: equipment.v1.Equipment.DeleteEquipment:output_type -> equipment.v1.DeleteEquipmentReply
	7, // 8: equipment.v1.Equipment.GetEquipment:output_type -> equipment.v1.GetEquipmentReply
	9, // 9: equipment.v1.Equipment.ListEquipment:output_type -> equipment.v1.ListEquipmentReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_equipment_v1_equipment_proto_init() }
func file_api_equipment_v1_equipment_proto_init() {
	if File_api_equipment_v1_equipment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_equipment_v1_equipment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEquipmentRequest); i {
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
		file_api_equipment_v1_equipment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEquipmentReply); i {
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
		file_api_equipment_v1_equipment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEquipmentRequest); i {
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
		file_api_equipment_v1_equipment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEquipmentReply); i {
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
		file_api_equipment_v1_equipment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEquipmentRequest); i {
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
		file_api_equipment_v1_equipment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEquipmentReply); i {
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
		file_api_equipment_v1_equipment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEquipmentRequest); i {
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
		file_api_equipment_v1_equipment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEquipmentReply); i {
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
		file_api_equipment_v1_equipment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEquipmentRequest); i {
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
		file_api_equipment_v1_equipment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEquipmentReply); i {
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
			RawDescriptor: file_api_equipment_v1_equipment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_equipment_v1_equipment_proto_goTypes,
		DependencyIndexes: file_api_equipment_v1_equipment_proto_depIdxs,
		MessageInfos:      file_api_equipment_v1_equipment_proto_msgTypes,
	}.Build()
	File_api_equipment_v1_equipment_proto = out.File
	file_api_equipment_v1_equipment_proto_rawDesc = nil
	file_api_equipment_v1_equipment_proto_goTypes = nil
	file_api_equipment_v1_equipment_proto_depIdxs = nil
}
