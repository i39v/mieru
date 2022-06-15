// Copyright (C) 2021  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: servercfg.proto

package appctlpb

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

type ServerAdvancedSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Allow using socks5 to access resources served in localhost.
	// This option should be set to false unless required due to testing purpose.
	AllowLocalDestination bool `protobuf:"varint,1,opt,name=allowLocalDestination,proto3" json:"allowLocalDestination,omitempty"`
}

func (x *ServerAdvancedSettings) Reset() {
	*x = ServerAdvancedSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servercfg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerAdvancedSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerAdvancedSettings) ProtoMessage() {}

func (x *ServerAdvancedSettings) ProtoReflect() protoreflect.Message {
	mi := &file_servercfg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerAdvancedSettings.ProtoReflect.Descriptor instead.
func (*ServerAdvancedSettings) Descriptor() ([]byte, []int) {
	return file_servercfg_proto_rawDescGZIP(), []int{0}
}

func (x *ServerAdvancedSettings) GetAllowLocalDestination() bool {
	if x != nil {
		return x.AllowLocalDestination
	}
	return false
}

type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Server's port-protocol bindings.
	PortBindings []*PortBinding `protobuf:"bytes,1,rep,name=portBindings,proto3" json:"portBindings,omitempty"`
	// A list of registered users.
	Users            []*User                 `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
	AdvancedSettings *ServerAdvancedSettings `protobuf:"bytes,3,opt,name=advancedSettings,proto3" json:"advancedSettings,omitempty"`
	// Server logging level.
	LoggingLevel LoggingLevel `protobuf:"varint,4,opt,name=loggingLevel,proto3,enum=appctl.LoggingLevel" json:"loggingLevel,omitempty"`
	// Maximum transmission unit of L2 payload.
	// This setting only applies to UDP protocol.
	Mtu int32 `protobuf:"varint,5,opt,name=mtu,proto3" json:"mtu,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servercfg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_servercfg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_servercfg_proto_rawDescGZIP(), []int{1}
}

func (x *ServerConfig) GetPortBindings() []*PortBinding {
	if x != nil {
		return x.PortBindings
	}
	return nil
}

func (x *ServerConfig) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ServerConfig) GetAdvancedSettings() *ServerAdvancedSettings {
	if x != nil {
		return x.AdvancedSettings
	}
	return nil
}

func (x *ServerConfig) GetLoggingLevel() LoggingLevel {
	if x != nil {
		return x.LoggingLevel
	}
	return LoggingLevel_DEFAULT
}

func (x *ServerConfig) GetMtu() int32 {
	if x != nil {
		return x.Mtu
	}
	return 0
}

var File_servercfg_proto protoreflect.FileDescriptor

var file_servercfg_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x63, 0x66, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x1a, 0x0b, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4e, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x34, 0x0a, 0x15, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x83, 0x02, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x37, 0x0a, 0x0c, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74,
	0x6c, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x0c, 0x70,
	0x6f, 0x72, 0x74, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x22, 0x0a, 0x05, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x70,
	0x63, 0x74, 0x6c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x4a, 0x0a, 0x10, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x70, 0x63,
	0x74, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65,
	0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x10, 0x61, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x38, 0x0a, 0x0c, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x4c, 0x6f, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x74, 0x75, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x6d, 0x74, 0x75, 0x32, 0x80, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x30, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x0d, 0x2e, 0x61,
	0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x61, 0x70,
	0x70, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x37, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14,
	0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x66, 0x65, 0x69, 0x6e, 0x2f,
	0x6d, 0x69, 0x65, 0x72, 0x75, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c,
	0x2f, 0x61, 0x70, 0x70, 0x63, 0x74, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_servercfg_proto_rawDescOnce sync.Once
	file_servercfg_proto_rawDescData = file_servercfg_proto_rawDesc
)

func file_servercfg_proto_rawDescGZIP() []byte {
	file_servercfg_proto_rawDescOnce.Do(func() {
		file_servercfg_proto_rawDescData = protoimpl.X.CompressGZIP(file_servercfg_proto_rawDescData)
	})
	return file_servercfg_proto_rawDescData
}

var file_servercfg_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_servercfg_proto_goTypes = []interface{}{
	(*ServerAdvancedSettings)(nil), // 0: appctl.ServerAdvancedSettings
	(*ServerConfig)(nil),           // 1: appctl.ServerConfig
	(*PortBinding)(nil),            // 2: appctl.PortBinding
	(*User)(nil),                   // 3: appctl.User
	(LoggingLevel)(0),              // 4: appctl.LoggingLevel
	(*Empty)(nil),                  // 5: appctl.Empty
}
var file_servercfg_proto_depIdxs = []int32{
	2, // 0: appctl.ServerConfig.portBindings:type_name -> appctl.PortBinding
	3, // 1: appctl.ServerConfig.users:type_name -> appctl.User
	0, // 2: appctl.ServerConfig.advancedSettings:type_name -> appctl.ServerAdvancedSettings
	4, // 3: appctl.ServerConfig.loggingLevel:type_name -> appctl.LoggingLevel
	5, // 4: appctl.ServerConfigService.GetConfig:input_type -> appctl.Empty
	1, // 5: appctl.ServerConfigService.SetConfig:input_type -> appctl.ServerConfig
	1, // 6: appctl.ServerConfigService.GetConfig:output_type -> appctl.ServerConfig
	1, // 7: appctl.ServerConfigService.SetConfig:output_type -> appctl.ServerConfig
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_servercfg_proto_init() }
func file_servercfg_proto_init() {
	if File_servercfg_proto != nil {
		return
	}
	file_empty_proto_init()
	file_endpoint_proto_init()
	file_logging_proto_init()
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_servercfg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerAdvancedSettings); i {
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
		file_servercfg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
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
			RawDescriptor: file_servercfg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_servercfg_proto_goTypes,
		DependencyIndexes: file_servercfg_proto_depIdxs,
		MessageInfos:      file_servercfg_proto_msgTypes,
	}.Build()
	File_servercfg_proto = out.File
	file_servercfg_proto_rawDesc = nil
	file_servercfg_proto_goTypes = nil
	file_servercfg_proto_depIdxs = nil
}
