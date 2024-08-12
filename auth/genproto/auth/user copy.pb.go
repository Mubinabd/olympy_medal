// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: submodule/auth_service/user copy.proto

package auth

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

type GetById struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetById) Reset() {
	*x = GetById{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_auth_service_user_copy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetById) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetById) ProtoMessage() {}

func (x *GetById) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_auth_service_user_copy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetById.ProtoReflect.Descriptor instead.
func (*GetById) Descriptor() ([]byte, []int) {
	return file_submodule_auth_service_user_copy_proto_rawDescGZIP(), []int{0}
}

func (x *GetById) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_auth_service_user_copy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_auth_service_user_copy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_submodule_auth_service_user_copy_proto_rawDescGZIP(), []int{1}
}

type UserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	LastName    string `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Role        string `protobuf:"bytes,6,opt,name=Role,proto3" json:"Role,omitempty"`
}

func (x *UserRes) Reset() {
	*x = UserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_auth_service_user_copy_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRes) ProtoMessage() {}

func (x *UserRes) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_auth_service_user_copy_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRes.ProtoReflect.Descriptor instead.
func (*UserRes) Descriptor() ([]byte, []int) {
	return file_submodule_auth_service_user_copy_proto_rawDescGZIP(), []int{2}
}

func (x *UserRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserRes) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRes) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UserRes) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UserRes) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type EditProfileReqBpdy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	Email       string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
}

func (x *EditProfileReqBpdy) Reset() {
	*x = EditProfileReqBpdy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_auth_service_user_copy_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EditProfileReqBpdy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EditProfileReqBpdy) ProtoMessage() {}

func (x *EditProfileReqBpdy) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_auth_service_user_copy_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EditProfileReqBpdy.ProtoReflect.Descriptor instead.
func (*EditProfileReqBpdy) Descriptor() ([]byte, []int) {
	return file_submodule_auth_service_user_copy_proto_rawDescGZIP(), []int{3}
}

func (x *EditProfileReqBpdy) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *EditProfileReqBpdy) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *EditProfileReqBpdy) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *EditProfileReqBpdy) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

type ChangePasswordReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CurrentPassword string `protobuf:"bytes,2,opt,name=CurrentPassword,proto3" json:"CurrentPassword,omitempty"`
	NewPassword     string `protobuf:"bytes,3,opt,name=NewPassword,proto3" json:"NewPassword,omitempty"`
}

func (x *ChangePasswordReq) Reset() {
	*x = ChangePasswordReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_auth_service_user_copy_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordReq) ProtoMessage() {}

func (x *ChangePasswordReq) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_auth_service_user_copy_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordReq.ProtoReflect.Descriptor instead.
func (*ChangePasswordReq) Descriptor() ([]byte, []int) {
	return file_submodule_auth_service_user_copy_proto_rawDescGZIP(), []int{4}
}

func (x *ChangePasswordReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangePasswordReq) GetCurrentPassword() string {
	if x != nil {
		return x.CurrentPassword
	}
	return ""
}

func (x *ChangePasswordReq) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type ChangePasswordReqBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPassword string `protobuf:"bytes,1,opt,name=CurrentPassword,proto3" json:"CurrentPassword,omitempty"`
	NewPassword     string `protobuf:"bytes,2,opt,name=NewPassword,proto3" json:"NewPassword,omitempty"`
}

func (x *ChangePasswordReqBody) Reset() {
	*x = ChangePasswordReqBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_auth_service_user_copy_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePasswordReqBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePasswordReqBody) ProtoMessage() {}

func (x *ChangePasswordReqBody) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_auth_service_user_copy_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePasswordReqBody.ProtoReflect.Descriptor instead.
func (*ChangePasswordReqBody) Descriptor() ([]byte, []int) {
	return file_submodule_auth_service_user_copy_proto_rawDescGZIP(), []int{5}
}

func (x *ChangePasswordReqBody) GetCurrentPassword() string {
	if x != nil {
		return x.CurrentPassword
	}
	return ""
}

func (x *ChangePasswordReqBody) GetNewPassword() string {
	if x != nil {
		return x.NewPassword
	}
	return ""
}

type SettingReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	PrivacyLevel string `protobuf:"bytes,2,opt,name=PrivacyLevel,proto3" json:"PrivacyLevel,omitempty"`
	Notification string `protobuf:"bytes,3,opt,name=Notification,proto3" json:"Notification,omitempty"`
	Language     string `protobuf:"bytes,4,opt,name=Language,proto3" json:"Language,omitempty"`
	Theme        string `protobuf:"bytes,5,opt,name=Theme,proto3" json:"Theme,omitempty"`
}

func (x *SettingReq) Reset() {
	*x = SettingReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_auth_service_user_copy_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingReq) ProtoMessage() {}

func (x *SettingReq) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_auth_service_user_copy_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingReq.ProtoReflect.Descriptor instead.
func (*SettingReq) Descriptor() ([]byte, []int) {
	return file_submodule_auth_service_user_copy_proto_rawDescGZIP(), []int{6}
}

func (x *SettingReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SettingReq) GetPrivacyLevel() string {
	if x != nil {
		return x.PrivacyLevel
	}
	return ""
}

func (x *SettingReq) GetNotification() string {
	if x != nil {
		return x.Notification
	}
	return ""
}

func (x *SettingReq) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *SettingReq) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

type Setting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivacyLevel string `protobuf:"bytes,1,opt,name=PrivacyLevel,proto3" json:"PrivacyLevel,omitempty"`
	Notification string `protobuf:"bytes,2,opt,name=Notification,proto3" json:"Notification,omitempty"`
	Language     string `protobuf:"bytes,3,opt,name=Language,proto3" json:"Language,omitempty"`
	Theme        string `protobuf:"bytes,4,opt,name=Theme,proto3" json:"Theme,omitempty"`
}

func (x *Setting) Reset() {
	*x = Setting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_submodule_auth_service_user_copy_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Setting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Setting) ProtoMessage() {}

func (x *Setting) ProtoReflect() protoreflect.Message {
	mi := &file_submodule_auth_service_user_copy_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Setting.ProtoReflect.Descriptor instead.
func (*Setting) Descriptor() ([]byte, []int) {
	return file_submodule_auth_service_user_copy_proto_rawDescGZIP(), []int{7}
}

func (x *Setting) GetPrivacyLevel() string {
	if x != nil {
		return x.PrivacyLevel
	}
	return ""
}

func (x *Setting) GetNotification() string {
	if x != nil {
		return x.Notification
	}
	return ""
}

func (x *Setting) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *Setting) GetTheme() string {
	if x != nil {
		return x.Theme
	}
	return ""
}

var File_submodule_auth_service_user_copy_proto protoreflect.FileDescriptor

var file_submodule_auth_service_user_copy_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x75, 0x62, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x20, 0x63, 0x6f,
	0x70, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x19,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0xa2, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x22, 0x89, 0x01, 0x0a, 0x12, 0x45, 0x64, 0x69, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x42, 0x70, 0x64, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x6f, 0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x63, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x28, 0x0a,
	0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x65, 0x77, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x65,
	0x77, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x96, 0x01, 0x0a, 0x0a, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x68, 0x65,
	0x6d, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x22,
	0x0a, 0x0c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x72, 0x69, 0x76, 0x61, 0x63, 0x79, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x32, 0xa2, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x12, 0x36, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x17, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x1a, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x2c, 0x0a, 0x0b, 0x45, 0x64, 0x69, 0x74, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0f, 0x5a,
	0x0d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_submodule_auth_service_user_copy_proto_rawDescOnce sync.Once
	file_submodule_auth_service_user_copy_proto_rawDescData = file_submodule_auth_service_user_copy_proto_rawDesc
)

func file_submodule_auth_service_user_copy_proto_rawDescGZIP() []byte {
	file_submodule_auth_service_user_copy_proto_rawDescOnce.Do(func() {
		file_submodule_auth_service_user_copy_proto_rawDescData = protoimpl.X.CompressGZIP(file_submodule_auth_service_user_copy_proto_rawDescData)
	})
	return file_submodule_auth_service_user_copy_proto_rawDescData
}

var file_submodule_auth_service_user_copy_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_submodule_auth_service_user_copy_proto_goTypes = []any{
	(*GetById)(nil),               // 0: auth.GetById
	(*Empty)(nil),                 // 1: auth.Empty
	(*UserRes)(nil),               // 2: auth.UserRes
	(*EditProfileReqBpdy)(nil),    // 3: auth.EditProfileReqBpdy
	(*ChangePasswordReq)(nil),     // 4: auth.ChangePasswordReq
	(*ChangePasswordReqBody)(nil), // 5: auth.ChangePasswordReqBody
	(*SettingReq)(nil),            // 6: auth.SettingReq
	(*Setting)(nil),               // 7: auth.Setting
}
var file_submodule_auth_service_user_copy_proto_depIdxs = []int32{
	0, // 0: auth.UserService.GetProfile:input_type -> auth.GetById
	2, // 1: auth.UserService.EditProfile:input_type -> auth.UserRes
	4, // 2: auth.UserService.ChangePassword:input_type -> auth.ChangePasswordReq
	0, // 3: auth.UserService.GetSetting:input_type -> auth.GetById
	6, // 4: auth.UserService.EditSetting:input_type -> auth.SettingReq
	0, // 5: auth.UserService.DeleteUser:input_type -> auth.GetById
	2, // 6: auth.UserService.GetProfile:output_type -> auth.UserRes
	2, // 7: auth.UserService.EditProfile:output_type -> auth.UserRes
	1, // 8: auth.UserService.ChangePassword:output_type -> auth.Empty
	7, // 9: auth.UserService.GetSetting:output_type -> auth.Setting
	1, // 10: auth.UserService.EditSetting:output_type -> auth.Empty
	1, // 11: auth.UserService.DeleteUser:output_type -> auth.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_submodule_auth_service_user_copy_proto_init() }
func file_submodule_auth_service_user_copy_proto_init() {
	if File_submodule_auth_service_user_copy_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_submodule_auth_service_user_copy_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetById); i {
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
		file_submodule_auth_service_user_copy_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_submodule_auth_service_user_copy_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UserRes); i {
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
		file_submodule_auth_service_user_copy_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EditProfileReqBpdy); i {
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
		file_submodule_auth_service_user_copy_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ChangePasswordReq); i {
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
		file_submodule_auth_service_user_copy_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ChangePasswordReqBody); i {
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
		file_submodule_auth_service_user_copy_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SettingReq); i {
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
		file_submodule_auth_service_user_copy_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Setting); i {
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
			RawDescriptor: file_submodule_auth_service_user_copy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_submodule_auth_service_user_copy_proto_goTypes,
		DependencyIndexes: file_submodule_auth_service_user_copy_proto_depIdxs,
		MessageInfos:      file_submodule_auth_service_user_copy_proto_msgTypes,
	}.Build()
	File_submodule_auth_service_user_copy_proto = out.File
	file_submodule_auth_service_user_copy_proto_rawDesc = nil
	file_submodule_auth_service_user_copy_proto_goTypes = nil
	file_submodule_auth_service_user_copy_proto_depIdxs = nil
}
