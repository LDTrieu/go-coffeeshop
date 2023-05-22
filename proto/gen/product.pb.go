// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: product.proto

package gen

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type GetItemTypesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetItemTypesRequest) Reset() {
	*x = GetItemTypesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemTypesRequest) ProtoMessage() {}

func (x *GetItemTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemTypesRequest.ProtoReflect.Descriptor instead.
func (*GetItemTypesRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

type GetItemTypesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemTypes []*ItemTypeDto `protobuf:"bytes,1,rep,name=item_types,json=itemTypes,proto3" json:"item_types,omitempty"`
}

func (x *GetItemTypesResponse) Reset() {
	*x = GetItemTypesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemTypesResponse) ProtoMessage() {}

func (x *GetItemTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemTypesResponse.ProtoReflect.Descriptor instead.
func (*GetItemTypesResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *GetItemTypesResponse) GetItemTypes() []*ItemTypeDto {
	if x != nil {
		return x.ItemTypes
	}
	return nil
}

type GetItemsByTypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemTypes string `protobuf:"bytes,1,opt,name=item_types,json=itemTypes,proto3" json:"item_types,omitempty"`
}

func (x *GetItemsByTypeRequest) Reset() {
	*x = GetItemsByTypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsByTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsByTypeRequest) ProtoMessage() {}

func (x *GetItemsByTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsByTypeRequest.ProtoReflect.Descriptor instead.
func (*GetItemsByTypeRequest) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{2}
}

func (x *GetItemsByTypeRequest) GetItemTypes() string {
	if x != nil {
		return x.ItemTypes
	}
	return ""
}

type GetItemsByTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*ItemDto `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *GetItemsByTypeResponse) Reset() {
	*x = GetItemsByTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItemsByTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItemsByTypeResponse) ProtoMessage() {}

func (x *GetItemsByTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItemsByTypeResponse.ProtoReflect.Descriptor instead.
func (*GetItemsByTypeResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{3}
}

func (x *GetItemsByTypeResponse) GetItems() []*ItemDto {
	if x != nil {
		return x.Items
	}
	return nil
}

type ItemDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price float64 `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty"`
	Type  int32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ItemDto) Reset() {
	*x = ItemDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemDto) ProtoMessage() {}

func (x *ItemDto) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemDto.ProtoReflect.Descriptor instead.
func (*ItemDto) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{4}
}

func (x *ItemDto) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ItemDto) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type ItemTypeDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type  int32   `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Price float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *ItemTypeDto) Reset() {
	*x = ItemTypeDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ItemTypeDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ItemTypeDto) ProtoMessage() {}

func (x *ItemTypeDto) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ItemTypeDto.ProtoReflect.Descriptor instead.
func (*ItemTypeDto) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{5}
}

func (x *ItemTypeDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ItemTypeDto) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *ItemTypeDto) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1e, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x61, 0x70, 0x69, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x62, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0a,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x61, 0x70,
	0x69, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x44, 0x74, 0x6f, 0x52, 0x09, 0x69,
	0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x57, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x2e, 0x63,
	0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x44,
	0x74, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x33, 0x0a, 0x07, 0x49, 0x74, 0x65,
	0x6d, 0x44, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4b,
	0x0a, 0x0b, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0xdf, 0x03, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xd8,
	0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12,
	0x33, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x61, 0x70, 0x69,
	0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x5d, 0x92, 0x41, 0x40, 0x0a,
	0x09, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x0f, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x69, 0x74, 0x65, 0x6d, 0x20, 0x74, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x22, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x20, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69,
	0x74, 0x65, 0x6d, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x12, 0xf1, 0x01, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x35, 0x2e, 0x67,
	0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x67, 0x6f, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x70, 0x92, 0x41, 0x42,
	0x0a, 0x09, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x20, 0x62, 0x79, 0x20, 0x74, 0x79, 0x70, 0x65, 0x1a,
	0x21, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x20, 0x62, 0x79, 0x20, 0x74,
	0x79, 0x70, 0x65, 0x20, 0x6f, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x2d, 0x62, 0x79, 0x2d, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2f, 0x7b, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x7d, 0x42, 0x19, 0x5a,
	0x17, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_product_proto_goTypes = []interface{}{
	(*GetItemTypesRequest)(nil),    // 0: go.coffeeshop.proto.productapi.GetItemTypesRequest
	(*GetItemTypesResponse)(nil),   // 1: go.coffeeshop.proto.productapi.GetItemTypesResponse
	(*GetItemsByTypeRequest)(nil),  // 2: go.coffeeshop.proto.productapi.GetItemsByTypeRequest
	(*GetItemsByTypeResponse)(nil), // 3: go.coffeeshop.proto.productapi.GetItemsByTypeResponse
	(*ItemDto)(nil),                // 4: go.coffeeshop.proto.productapi.ItemDto
	(*ItemTypeDto)(nil),            // 5: go.coffeeshop.proto.productapi.ItemTypeDto
}
var file_product_proto_depIdxs = []int32{
	5, // 0: go.coffeeshop.proto.productapi.GetItemTypesResponse.item_types:type_name -> go.coffeeshop.proto.productapi.ItemTypeDto
	4, // 1: go.coffeeshop.proto.productapi.GetItemsByTypeResponse.items:type_name -> go.coffeeshop.proto.productapi.ItemDto
	0, // 2: go.coffeeshop.proto.productapi.ProductService.GetItemTypes:input_type -> go.coffeeshop.proto.productapi.GetItemTypesRequest
	2, // 3: go.coffeeshop.proto.productapi.ProductService.GetItemsByType:input_type -> go.coffeeshop.proto.productapi.GetItemsByTypeRequest
	1, // 4: go.coffeeshop.proto.productapi.ProductService.GetItemTypes:output_type -> go.coffeeshop.proto.productapi.GetItemTypesResponse
	3, // 5: go.coffeeshop.proto.productapi.ProductService.GetItemsByType:output_type -> go.coffeeshop.proto.productapi.GetItemsByTypeResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemTypesRequest); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemTypesResponse); i {
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
		file_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemsByTypeRequest); i {
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
		file_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItemsByTypeResponse); i {
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
		file_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemDto); i {
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
		file_product_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ItemTypeDto); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}
