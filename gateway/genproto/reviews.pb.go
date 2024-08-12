// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/reviews.proto

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

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	BookingId  *Booking  `protobuf:"bytes,2,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	UserId     string    `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProviderId *Provider `protobuf:"bytes,4,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	Rating     int32     `protobuf:"varint,5,opt,name=rating,proto3" json:"rating,omitempty"`
	Comment    string    `protobuf:"bytes,6,opt,name=comment,proto3" json:"comment,omitempty"`
	CreatedAt  string    `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string    `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt  string    `protobuf:"bytes,9,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reviews_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reviews_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_protos_reviews_proto_rawDescGZIP(), []int{0}
}

func (x *Review) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Review) GetBookingId() *Booking {
	if x != nil {
		return x.BookingId
	}
	return nil
}

func (x *Review) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Review) GetProviderId() *Provider {
	if x != nil {
		return x.ProviderId
	}
	return nil
}

func (x *Review) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Review) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *Review) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Review) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Review) GetDeletedAt() string {
	if x != nil {
		return x.DeletedAt
	}
	return ""
}

type ListAllReviewsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId  string  `protobuf:"bytes,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	UserId     string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProviderId string  `protobuf:"bytes,3,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	Filter     *Filter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListAllReviewsReq) Reset() {
	*x = ListAllReviewsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reviews_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllReviewsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllReviewsReq) ProtoMessage() {}

func (x *ListAllReviewsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reviews_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllReviewsReq.ProtoReflect.Descriptor instead.
func (*ListAllReviewsReq) Descriptor() ([]byte, []int) {
	return file_protos_reviews_proto_rawDescGZIP(), []int{1}
}

func (x *ListAllReviewsReq) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

func (x *ListAllReviewsReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListAllReviewsReq) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *ListAllReviewsReq) GetFilter() *Filter {
	if x != nil {
		return x.Filter
	}
	return nil
}

type ListAllReviewsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reviews []*Review `protobuf:"bytes,1,rep,name=reviews,proto3" json:"reviews,omitempty"`
}

func (x *ListAllReviewsResp) Reset() {
	*x = ListAllReviewsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reviews_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllReviewsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllReviewsResp) ProtoMessage() {}

func (x *ListAllReviewsResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reviews_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllReviewsResp.ProtoReflect.Descriptor instead.
func (*ListAllReviewsResp) Descriptor() ([]byte, []int) {
	return file_protos_reviews_proto_rawDescGZIP(), []int{2}
}

func (x *ListAllReviewsResp) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type AddReviewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId  string `protobuf:"bytes,1,opt,name=booking_id,json=bookingId,proto3" json:"booking_id,omitempty"`
	Rating     int32  `protobuf:"varint,2,opt,name=rating,proto3" json:"rating,omitempty"`
	Comment    string `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProviderId string `protobuf:"bytes,5,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
}

func (x *AddReviewReq) Reset() {
	*x = AddReviewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reviews_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReviewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReviewReq) ProtoMessage() {}

func (x *AddReviewReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reviews_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReviewReq.ProtoReflect.Descriptor instead.
func (*AddReviewReq) Descriptor() ([]byte, []int) {
	return file_protos_reviews_proto_rawDescGZIP(), []int{3}
}

func (x *AddReviewReq) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

func (x *AddReviewReq) GetRating() int32 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *AddReviewReq) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

func (x *AddReviewReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddReviewReq) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

type UpdateReviewsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Review *Review `protobuf:"bytes,2,opt,name=review,proto3" json:"review,omitempty"`
}

func (x *UpdateReviewsReq) Reset() {
	*x = UpdateReviewsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reviews_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReviewsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewsReq) ProtoMessage() {}

func (x *UpdateReviewsReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reviews_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewsReq.ProtoReflect.Descriptor instead.
func (*UpdateReviewsReq) Descriptor() ([]byte, []int) {
	return file_protos_reviews_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateReviewsReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateReviewsReq) GetReview() *Review {
	if x != nil {
		return x.Review
	}
	return nil
}

type UpdateReviewsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateReviewsResp) Reset() {
	*x = UpdateReviewsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reviews_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReviewsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReviewsResp) ProtoMessage() {}

func (x *UpdateReviewsResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reviews_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReviewsResp.ProtoReflect.Descriptor instead.
func (*UpdateReviewsResp) Descriptor() ([]byte, []int) {
	return file_protos_reviews_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateReviewsResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *UpdateReviewsResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteReviewReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteReviewReq) Reset() {
	*x = DeleteReviewReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reviews_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReviewReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReviewReq) ProtoMessage() {}

func (x *DeleteReviewReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reviews_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReviewReq.ProtoReflect.Descriptor instead.
func (*DeleteReviewReq) Descriptor() ([]byte, []int) {
	return file_protos_reviews_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteReviewReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteReviewResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteReviewResp) Reset() {
	*x = DeleteReviewResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_reviews_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReviewResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReviewResp) ProtoMessage() {}

func (x *DeleteReviewResp) ProtoReflect() protoreflect.Message {
	mi := &file_protos_reviews_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReviewResp.ProtoReflect.Descriptor instead.
func (*DeleteReviewResp) Descriptor() ([]byte, []int) {
	return file_protos_reviews_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteReviewResp) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *DeleteReviewResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_protos_reviews_proto protoreflect.FileDescriptor

var file_protos_reviews_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68,
	0x1a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30,
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x33, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x96, 0x01,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63,
	0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x06,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x40, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c,
	0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x07, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x0c, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x61, 0x74, 0x69,
	0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a, 0x06, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77,
	0x61, 0x73, 0x68, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x06, 0x72, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x22, 0x47, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xd4, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68,
	0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e,
	0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x30,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x11, 0x2e, 0x63, 0x61,
	0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x10,
	0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x47, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x12, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x63,
	0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x45, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x5f,
	0x77, 0x61, 0x73, 0x68, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x4b, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x12, 0x1b, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x1a,
	0x1c, 0x2e, 0x63, 0x61, 0x72, 0x5f, 0x77, 0x61, 0x73, 0x68, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x0b, 0x5a,
	0x09, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_reviews_proto_rawDescOnce sync.Once
	file_protos_reviews_proto_rawDescData = file_protos_reviews_proto_rawDesc
)

func file_protos_reviews_proto_rawDescGZIP() []byte {
	file_protos_reviews_proto_rawDescOnce.Do(func() {
		file_protos_reviews_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_reviews_proto_rawDescData)
	})
	return file_protos_reviews_proto_rawDescData
}

var file_protos_reviews_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_reviews_proto_goTypes = []any{
	(*Review)(nil),             // 0: car_wash.Review
	(*ListAllReviewsReq)(nil),  // 1: car_wash.ListAllReviewsReq
	(*ListAllReviewsResp)(nil), // 2: car_wash.ListAllReviewsResp
	(*AddReviewReq)(nil),       // 3: car_wash.AddReviewReq
	(*UpdateReviewsReq)(nil),   // 4: car_wash.UpdateReviewsReq
	(*UpdateReviewsResp)(nil),  // 5: car_wash.UpdateReviewsResp
	(*DeleteReviewReq)(nil),    // 6: car_wash.DeleteReviewReq
	(*DeleteReviewResp)(nil),   // 7: car_wash.DeleteReviewResp
	(*Booking)(nil),            // 8: car_wash.Booking
	(*Provider)(nil),           // 9: car_wash.Provider
	(*Filter)(nil),             // 10: car_wash.Filter
	(*GetById)(nil),            // 11: car_wash.GetById
	(*Empty)(nil),              // 12: car_wash.Empty
}
var file_protos_reviews_proto_depIdxs = []int32{
	8,  // 0: car_wash.Review.booking_id:type_name -> car_wash.Booking
	9,  // 1: car_wash.Review.provider_id:type_name -> car_wash.Provider
	10, // 2: car_wash.ListAllReviewsReq.filter:type_name -> car_wash.Filter
	0,  // 3: car_wash.ListAllReviewsResp.reviews:type_name -> car_wash.Review
	0,  // 4: car_wash.UpdateReviewsReq.review:type_name -> car_wash.Review
	3,  // 5: car_wash.ReviewService.AddReview:input_type -> car_wash.AddReviewReq
	11, // 6: car_wash.ReviewService.GetReview:input_type -> car_wash.GetById
	4,  // 7: car_wash.ReviewService.UpdateReview:input_type -> car_wash.UpdateReviewsReq
	6,  // 8: car_wash.ReviewService.DeleteReview:input_type -> car_wash.DeleteReviewReq
	1,  // 9: car_wash.ReviewService.ListAllReviews:input_type -> car_wash.ListAllReviewsReq
	12, // 10: car_wash.ReviewService.AddReview:output_type -> car_wash.Empty
	0,  // 11: car_wash.ReviewService.GetReview:output_type -> car_wash.Review
	5,  // 12: car_wash.ReviewService.UpdateReview:output_type -> car_wash.UpdateReviewsResp
	7,  // 13: car_wash.ReviewService.DeleteReview:output_type -> car_wash.DeleteReviewResp
	2,  // 14: car_wash.ReviewService.ListAllReviews:output_type -> car_wash.ListAllReviewsResp
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_protos_reviews_proto_init() }
func file_protos_reviews_proto_init() {
	if File_protos_reviews_proto != nil {
		return
	}
	file_protos_common_proto_init()
	file_protos_booking_proto_init()
	file_protos_providers_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_reviews_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Review); i {
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
		file_protos_reviews_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListAllReviewsReq); i {
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
		file_protos_reviews_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListAllReviewsResp); i {
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
		file_protos_reviews_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*AddReviewReq); i {
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
		file_protos_reviews_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateReviewsReq); i {
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
		file_protos_reviews_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateReviewsResp); i {
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
		file_protos_reviews_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteReviewReq); i {
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
		file_protos_reviews_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteReviewResp); i {
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
			RawDescriptor: file_protos_reviews_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_reviews_proto_goTypes,
		DependencyIndexes: file_protos_reviews_proto_depIdxs,
		MessageInfos:      file_protos_reviews_proto_msgTypes,
	}.Build()
	File_protos_reviews_proto = out.File
	file_protos_reviews_proto_rawDesc = nil
	file_protos_reviews_proto_goTypes = nil
	file_protos_reviews_proto_depIdxs = nil
}
