// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/services.proto

package genproto

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

type Services struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Duration    int32   `protobuf:"varint,5,opt,name=duration,proto3" json:"duration,omitempty"`
	CreatedAt   string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt   string  `protobuf:"bytes,8,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Services) Reset() {
	*x = Services{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Services) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Services) ProtoMessage() {}

func (x *Services) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Services.ProtoReflect.Descriptor instead.
func (*Services) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{0}
}

func (x *Services) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Services) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Services) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Services) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Services) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Services) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Services) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Services) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type AddServiceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Duration    int32   `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *AddServiceReq) Reset() {
	*x = AddServiceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddServiceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddServiceReq) ProtoMessage() {}

func (x *AddServiceReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddServiceReq.ProtoReflect.Descriptor instead.
func (*AddServiceReq) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{1}
}

func (x *AddServiceReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddServiceReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddServiceReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AddServiceReq) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

type GetServicesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services *Services `protobuf:"bytes,1,opt,name=services,proto3" json:"services,omitempty"`
}

func (x *GetServicesResp) Reset() {
	*x = GetServicesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesResp) ProtoMessage() {}

func (x *GetServicesResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesResp.ProtoReflect.Descriptor instead.
func (*GetServicesResp) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{2}
}

func (x *GetServicesResp) GetServices() *Services {
	if x != nil {
		return x.Services
	}
	return nil
}

type ListAllServicesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Filter      *Filter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListAllServicesReq) Reset() {
	*x = ListAllServicesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllServicesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllServicesReq) ProtoMessage() {}

func (x *ListAllServicesReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllServicesReq.ProtoReflect.Descriptor instead.
func (*ListAllServicesReq) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{3}
}

func (x *ListAllServicesReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListAllServicesReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ListAllServicesReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListAllServicesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*Services `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *ListAllServicesResp) Reset() {
	*x = ListAllServicesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllServicesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllServicesResp) ProtoMessage() {}

func (x *ListAllServicesResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllServicesResp.ProtoReflect.Descriptor instead.
func (*ListAllServicesResp) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{4}
}

func (x *ListAllServicesResp) GetServices() []*Services {
	if x != nil {
		return x.Services
	}
	return nil
}

type UpdateServiceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Services *Services `protobuf:"bytes,2,opt,name=services,proto3" json:"services,omitempty"`
}

func (x *UpdateServiceReq) Reset() {
	*x = UpdateServiceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServiceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServiceReq) ProtoMessage() {}

func (x *UpdateServiceReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServiceReq.ProtoReflect.Descriptor instead.
func (*UpdateServiceReq) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateServiceReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateServiceReq) GetServices() *Services {
	if x != nil {
		return x.Services
	}
	return nil
}

type UpdateServiceResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services *Services `protobuf:"bytes,1,opt,name=services,proto3" json:"services,omitempty"`
}

func (x *UpdateServiceResp) Reset() {
	*x = UpdateServiceResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateServiceResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateServiceResp) ProtoMessage() {}

func (x *UpdateServiceResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateServiceResp.ProtoReflect.Descriptor instead.
func (*UpdateServiceResp) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateServiceResp) GetServices() *Services {
	if x != nil {
		return x.Services
	}
	return nil
}

type DeleteServiesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteServiesReq) Reset() {
	*x = DeleteServiesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServiesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServiesReq) ProtoMessage() {}

func (x *DeleteServiesReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServiesReq.ProtoReflect.Descriptor instead.
func (*DeleteServiesReq) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteServiesReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteServiesResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteServiesResp) Reset() {
	*x = DeleteServiesResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServiesResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServiesResp) ProtoMessage() {}

func (x *DeleteServiesResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServiesResp.ProtoReflect.Descriptor instead.
func (*DeleteServiesResp) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteServiesResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteServiesResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SearchServicessReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Price       float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *SearchServicessReq) Reset() {
	*x = SearchServicessReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchServicessReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchServicessReq) ProtoMessage() {}

func (x *SearchServicessReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchServicessReq.ProtoReflect.Descriptor instead.
func (*SearchServicessReq) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{9}
}

func (x *SearchServicessReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SearchServicessReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SearchServicessReq) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type SearchServicessResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*Services `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *SearchServicessResp) Reset() {
	*x = SearchServicessResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchServicessResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchServicessResp) ProtoMessage() {}

func (x *SearchServicessResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchServicessResp.ProtoReflect.Descriptor instead.
func (*SearchServicessResp) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{10}
}

func (x *SearchServicessResp) GetServices() []*Services {
	if x != nil {
		return x.Services
	}
	return nil
}

type GetServicesByPriceRangeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinPrice float32 `protobuf:"fixed32,1,opt,name=min_price,json=minPrice,proto3" json:"min_price,omitempty"`
	MaxPrice float32 `protobuf:"fixed32,2,opt,name=max_price,json=maxPrice,proto3" json:"max_price,omitempty"`
}

func (x *GetServicesByPriceRangeReq) Reset() {
	*x = GetServicesByPriceRangeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicesByPriceRangeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesByPriceRangeReq) ProtoMessage() {}

func (x *GetServicesByPriceRangeReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesByPriceRangeReq.ProtoReflect.Descriptor instead.
func (*GetServicesByPriceRangeReq) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{11}
}

func (x *GetServicesByPriceRangeReq) GetMinPrice() float32 {
	if x != nil {
		return x.MinPrice
	}
	return 0
}

func (x *GetServicesByPriceRangeReq) GetMaxPrice() float32 {
	if x != nil {
		return x.MaxPrice
	}
	return 0
}

type GetServicesByPriceRangeResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*Services `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
}

func (x *GetServicesByPriceRangeResp) Reset() {
	*x = GetServicesByPriceRangeResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_services_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicesByPriceRangeResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesByPriceRangeResp) ProtoMessage() {}

func (x *GetServicesByPriceRangeResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_services_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesByPriceRangeResp.ProtoReflect.Descriptor instead.
func (*GetServicesByPriceRangeResp) Descriptor() ([]byte, []int) {
	return file_protos_services_proto_rawDescGZIP(), []int{12}
}

func (x *GetServicesByPriceRangeResp) GetServices() []*Services {
	if x != nil {
		return x.Services
	}
	return nil
}

var File_protos_services_proto protoreflect.FileDescriptor

var file_protos_services_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73,
	0x68, 0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x77, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x41, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73,
	0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x74, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x28, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x45, 0x0a, 0x13, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x22, 0x52, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61,
	0x73, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63,
	0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47,
	0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x60, 0x0a, 0x12, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x45, 0x0a, 0x13, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x22, 0x56, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42,
	0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b,
	0x0a, 0x09, 0x6d, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x61, 0x78, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08,
	0x6d, 0x61, 0x78, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x4d, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x5f,
	0x77, 0x61, 0x73, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0xa1, 0x04, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x0a, 0x41,
	0x64, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x61, 0x72, 0x5f,
	0x77, 0x61, 0x73, 0x68, 0x2e, 0x41, 0x64, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x4e, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x6c, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x48, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x48, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x2e, 0x63, 0x61,
	0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61,
	0x73, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x4d, 0x0a, 0x0e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73,
	0x68, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x66, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x24,
	0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0b, 0x5a, 0x09, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_services_proto_rawDescOnce sync.Once
	file_protos_services_proto_rawDescData = file_protos_services_proto_rawDesc
)

func file_protos_services_proto_rawDescGZIP() []byte {
	file_protos_services_proto_rawDescOnce.Do(func() {
		file_protos_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_services_proto_rawDescData)
	})
	return file_protos_services_proto_rawDescData
}

var file_protos_services_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_protos_services_proto_goTypes = []any{
	(*Services)(nil),                    // 0: car_wash.Services
	(*AddServiceReq)(nil),               // 1: car_wash.AddServiceReq
	(*GetServicesResp)(nil),             // 2: car_wash.GetServicesResp
	(*ListAllServicesReq)(nil),          // 3: car_wash.ListAllServicesReq
	(*ListAllServicesResp)(nil),         // 4: car_wash.ListAllServicesResp
	(*UpdateServiceReq)(nil),            // 5: car_wash.UpdateServiceReq
	(*UpdateServiceResp)(nil),           // 6: car_wash.UpdateServiceResp
	(*DeleteServiesReq)(nil),            // 7: car_wash.DeleteServiesReq
	(*DeleteServiesResp)(nil),           // 8: car_wash.DeleteServiesResp
	(*SearchServicessReq)(nil),          // 9: car_wash.SearchServicessReq
	(*SearchServicessResp)(nil),         // 10: car_wash.SearchServicessResp
	(*GetServicesByPriceRangeReq)(nil),  // 11: car_wash.GetServicesByPriceRangeReq
	(*GetServicesByPriceRangeResp)(nil), // 12: car_wash.GetServicesByPriceRangeResp
	(*Filter)(nil),                      // 13: car_wash.Filter
	(*GetById)(nil),                     // 14: car_wash.GetById
	(*Empty)(nil),                       // 15: car_wash.Empty
}
var file_protos_services_proto_depIdxs = []int32{
	0,  // 0: car_wash.GetServicesResp.services:type_name -> car_wash.Services
	13, // 1: car_wash.ListAllServicesReq.filter:type_name -> car_wash.Filter
	0,  // 2: car_wash.ListAllServicesResp.services:type_name -> car_wash.Services
	0,  // 3: car_wash.UpdateServiceReq.services:type_name -> car_wash.Services
	0,  // 4: car_wash.UpdateServiceResp.services:type_name -> car_wash.Services
	0,  // 5: car_wash.SearchServicessResp.services:type_name -> car_wash.Services
	0,  // 6: car_wash.GetServicesByPriceRangeResp.services:type_name -> car_wash.Services
	1,  // 7: car_wash.ServicesService.AddService:input_type -> car_wash.AddServiceReq
	14, // 8: car_wash.ServicesService.GetServices:input_type -> car_wash.GetById
	3,  // 9: car_wash.ServicesService.ListAllServices:input_type -> car_wash.ListAllServicesReq
	5,  // 10: car_wash.ServicesService.UpdateService:input_type -> car_wash.UpdateServiceReq
	7,  // 11: car_wash.ServicesService.DeleteService:input_type -> car_wash.DeleteServiesReq
	9,  // 12: car_wash.ServicesService.SearchServices:input_type -> car_wash.SearchServicessReq
	11, // 13: car_wash.ServicesService.GetServicesByPriceRange:input_type -> car_wash.GetServicesByPriceRangeReq
	15, // 14: car_wash.ServicesService.AddService:output_type -> car_wash.Empty
	2,  // 15: car_wash.ServicesService.GetServices:output_type -> car_wash.GetServicesResp
	4,  // 16: car_wash.ServicesService.ListAllServices:output_type -> car_wash.ListAllServicesResp
	6,  // 17: car_wash.ServicesService.UpdateService:output_type -> car_wash.UpdateServiceResp
	8,  // 18: car_wash.ServicesService.DeleteService:output_type -> car_wash.DeleteServiesResp
	10, // 19: car_wash.ServicesService.SearchServices:output_type -> car_wash.SearchServicessResp
	12, // 20: car_wash.ServicesService.GetServicesByPriceRange:output_type -> car_wash.GetServicesByPriceRangeResp
	14, // [14:21] is the sub-list for method output_type
	7,  // [7:14] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_protos_services_proto_init() }
func file_protos_services_proto_init() {
	if File_protos_services_proto != nil {
		return
	}
	file_protos_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_services_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Services); i {
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
		file_protos_services_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddServiceReq); i {
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
		file_protos_services_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetServicesResp); i {
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
		file_protos_services_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListAllServicesReq); i {
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
		file_protos_services_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListAllServicesResp); i {
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
		file_protos_services_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateServiceReq); i {
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
		file_protos_services_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateServiceResp); i {
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
		file_protos_services_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteServiesReq); i {
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
		file_protos_services_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteServiesResp); i {
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
		file_protos_services_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*SearchServicessReq); i {
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
		file_protos_services_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*SearchServicessResp); i {
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
		file_protos_services_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetServicesByPriceRangeReq); i {
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
		file_protos_services_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*GetServicesByPriceRangeResp); i {
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
			RawDescriptor: file_protos_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_services_proto_goTypes,
		DependencyIndexes: file_protos_services_proto_depIdxs,
		MessageInfos:      file_protos_services_proto_msgTypes,
	}.Build()
	File_protos_services_proto = out.File
	file_protos_services_proto_rawDesc = nil
	file_protos_services_proto_goTypes = nil
	file_protos_services_proto_depIdxs = nil
}
