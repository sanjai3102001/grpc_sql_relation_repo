// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: usermgmt.proto

package go_usermgmt_grpc

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

type NewUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId   int32  `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Member    int32  `protobuf:"varint,3,opt,name=member,proto3" json:"member,omitempty"`
	Admin     int32  `protobuf:"varint,4,opt,name=admin,proto3" json:"admin,omitempty"`
	ServiceId int32  `protobuf:"varint,5,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
}

func (x *NewUser) Reset() {
	*x = NewUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUser) ProtoMessage() {}

func (x *NewUser) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUser.ProtoReflect.Descriptor instead.
func (*NewUser) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{0}
}

func (x *NewUser) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *NewUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewUser) GetMember() int32 {
	if x != nil {
		return x.Member
	}
	return 0
}

func (x *NewUser) GetAdmin() int32 {
	if x != nil {
		return x.Admin
	}
	return 0
}

func (x *NewUser) GetServiceId() int32 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupId   int32  `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Member    int32  `protobuf:"varint,3,opt,name=member,proto3" json:"member,omitempty"`
	Admin     int32  `protobuf:"varint,4,opt,name=admin,proto3" json:"admin,omitempty"`
	ServiceId int32  `protobuf:"varint,5,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetGroupId() int32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetMember() int32 {
	if x != nil {
		return x.Member
	}
	return 0
}

func (x *User) GetAdmin() int32 {
	if x != nil {
		return x.Admin
	}
	return 0
}

func (x *User) GetServiceId() int32 {
	if x != nil {
		return x.ServiceId
	}
	return 0
}

type GetUsersParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUsersParams) Reset() {
	*x = GetUsersParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUsersParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersParams) ProtoMessage() {}

func (x *GetUsersParams) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersParams.ProtoReflect.Descriptor instead.
func (*GetUsersParams) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{2}
}

type UsersList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersList) Reset() {
	*x = UsersList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_usermgmt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersList) ProtoMessage() {}

func (x *UsersList) ProtoReflect() protoreflect.Message {
	mi := &file_usermgmt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersList.ProtoReflect.Descriptor instead.
func (*UsersList) Descriptor() ([]byte, []int) {
	return file_usermgmt_proto_rawDescGZIP(), []int{3}
}

func (x *UsersList) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_usermgmt_proto protoreflect.FileDescriptor

var file_usermgmt_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x22, 0x85, 0x01, 0x0a, 0x07, 0x4e,
	0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x31, 0x0a, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0x83, 0x01, 0x0a,
	0x0e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x34, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x13, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x65, 0x63, 0x68, 0x2d, 0x77, 0x69, 0x74, 0x68, 0x2d, 0x6d, 0x6f, 0x73, 0x73, 0x2f,
	0x67, 0x6f, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x3b, 0x67, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x6d, 0x67, 0x6d, 0x74, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_usermgmt_proto_rawDescOnce sync.Once
	file_usermgmt_proto_rawDescData = file_usermgmt_proto_rawDesc
)

func file_usermgmt_proto_rawDescGZIP() []byte {
	file_usermgmt_proto_rawDescOnce.Do(func() {
		file_usermgmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_usermgmt_proto_rawDescData)
	})
	return file_usermgmt_proto_rawDescData
}

var file_usermgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_usermgmt_proto_goTypes = []interface{}{
	(*NewUser)(nil),        // 0: usermgmt.NewUser
	(*User)(nil),           // 1: usermgmt.User
	(*GetUsersParams)(nil), // 2: usermgmt.GetUsersParams
	(*UsersList)(nil),      // 3: usermgmt.UsersList
}
var file_usermgmt_proto_depIdxs = []int32{
	1, // 0: usermgmt.UsersList.users:type_name -> usermgmt.User
	0, // 1: usermgmt.UserManagement.CreateNewUser:input_type -> usermgmt.NewUser
	2, // 2: usermgmt.UserManagement.GetUsers:input_type -> usermgmt.GetUsersParams
	1, // 3: usermgmt.UserManagement.CreateNewUser:output_type -> usermgmt.User
	3, // 4: usermgmt.UserManagement.GetUsers:output_type -> usermgmt.UsersList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_usermgmt_proto_init() }
func file_usermgmt_proto_init() {
	if File_usermgmt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_usermgmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUser); i {
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
		file_usermgmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_usermgmt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUsersParams); i {
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
		file_usermgmt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersList); i {
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
			RawDescriptor: file_usermgmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_usermgmt_proto_goTypes,
		DependencyIndexes: file_usermgmt_proto_depIdxs,
		MessageInfos:      file_usermgmt_proto_msgTypes,
	}.Build()
	File_usermgmt_proto = out.File
	file_usermgmt_proto_rawDesc = nil
	file_usermgmt_proto_goTypes = nil
	file_usermgmt_proto_depIdxs = nil
}
