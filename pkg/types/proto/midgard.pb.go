// Copyright 2020 The golang.design Initiative authors.
// All rights reserved. Use of this source code is governed by
// a GNU GPL-3.0 license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: midgard.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PingInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingInput) Reset() {
	*x = PingInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingInput) ProtoMessage() {}

func (x *PingInput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingInput.ProtoReflect.Descriptor instead.
func (*PingInput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{0}
}

type PingOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingOutput) Reset() {
	*x = PingOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingOutput) ProtoMessage() {}

func (x *PingOutput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingOutput.ProtoReflect.Descriptor instead.
func (*PingOutput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{1}
}

type GetClipboardInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetClipboardInput) Reset() {
	*x = GetClipboardInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClipboardInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClipboardInput) ProtoMessage() {}

func (x *GetClipboardInput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClipboardInput.ProtoReflect.Descriptor instead.
func (*GetClipboardInput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{2}
}

type GetClipboardOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *GetClipboardOutput) Reset() {
	*x = GetClipboardOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetClipboardOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClipboardOutput) ProtoMessage() {}

func (x *GetClipboardOutput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClipboardOutput.ProtoReflect.Descriptor instead.
func (*GetClipboardOutput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{3}
}

func (x *GetClipboardOutput) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PutClipboardInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (x *PutClipboardInput) Reset() {
	*x = PutClipboardInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutClipboardInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutClipboardInput) ProtoMessage() {}

func (x *PutClipboardInput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutClipboardInput.ProtoReflect.Descriptor instead.
func (*PutClipboardInput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{4}
}

func (x *PutClipboardInput) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type PutClipboardOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PutClipboardOutput) Reset() {
	*x = PutClipboardOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutClipboardOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutClipboardOutput) ProtoMessage() {}

func (x *PutClipboardOutput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutClipboardOutput.ProtoReflect.Descriptor instead.
func (*PutClipboardOutput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{5}
}

type GetURIInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetURIInput) Reset() {
	*x = GetURIInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURIInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURIInput) ProtoMessage() {}

func (x *GetURIInput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURIInput.ProtoReflect.Descriptor instead.
func (*GetURIInput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{6}
}

type GetURIOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *GetURIOutput) Reset() {
	*x = GetURIOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_midgard_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetURIOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetURIOutput) ProtoMessage() {}

func (x *GetURIOutput) ProtoReflect() protoreflect.Message {
	mi := &file_midgard_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetURIOutput.ProtoReflect.Descriptor instead.
func (*GetURIOutput) Descriptor() ([]byte, []int) {
	return file_midgard_proto_rawDescGZIP(), []int{7}
}

func (x *GetURIOutput) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_midgard_proto protoreflect.FileDescriptor

var file_midgard_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x69, 0x64, 0x67, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x50, 0x69, 0x6e, 0x67, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x50, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x28, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x27, 0x0a, 0x11, 0x50, 0x75, 0x74, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x22, 0x14, 0x0a, 0x12, 0x50, 0x75, 0x74,
	0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x0d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x52, 0x49, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x28,
	0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x52, 0x49, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0xfb, 0x01, 0x0a, 0x07, 0x4d, 0x69, 0x64,
	0x67, 0x61, 0x72, 0x64, 0x12, 0x2d, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x11,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f,
	0x61, 0x72, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x19, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x50, 0x75,
	0x74, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x75, 0x74, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x75, 0x74,
	0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22,
	0x00, 0x12, 0x33, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x55, 0x52, 0x49, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x49, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x52, 0x49, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_midgard_proto_rawDescOnce sync.Once
	file_midgard_proto_rawDescData = file_midgard_proto_rawDesc
)

func file_midgard_proto_rawDescGZIP() []byte {
	file_midgard_proto_rawDescOnce.Do(func() {
		file_midgard_proto_rawDescData = protoimpl.X.CompressGZIP(file_midgard_proto_rawDescData)
	})
	return file_midgard_proto_rawDescData
}

var file_midgard_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_midgard_proto_goTypes = []interface{}{
	(*PingInput)(nil),          // 0: proto.PingInput
	(*PingOutput)(nil),         // 1: proto.PingOutput
	(*GetClipboardInput)(nil),  // 2: proto.GetClipboardInput
	(*GetClipboardOutput)(nil), // 3: proto.GetClipboardOutput
	(*PutClipboardInput)(nil),  // 4: proto.PutClipboardInput
	(*PutClipboardOutput)(nil), // 5: proto.PutClipboardOutput
	(*GetURIInput)(nil),        // 6: proto.GetURIInput
	(*GetURIOutput)(nil),       // 7: proto.GetURIOutput
}
var file_midgard_proto_depIdxs = []int32{
	0, // 0: proto.Midgard.Ping:input_type -> proto.PingInput
	2, // 1: proto.Midgard.GetClipboard:input_type -> proto.GetClipboardInput
	4, // 2: proto.Midgard.PutClipboard:input_type -> proto.PutClipboardInput
	6, // 3: proto.Midgard.GetURI:input_type -> proto.GetURIInput
	1, // 4: proto.Midgard.Ping:output_type -> proto.PingOutput
	3, // 5: proto.Midgard.GetClipboard:output_type -> proto.GetClipboardOutput
	5, // 6: proto.Midgard.PutClipboard:output_type -> proto.PutClipboardOutput
	7, // 7: proto.Midgard.GetURI:output_type -> proto.GetURIOutput
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_midgard_proto_init() }
func file_midgard_proto_init() {
	if File_midgard_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_midgard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingInput); i {
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
		file_midgard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingOutput); i {
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
		file_midgard_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClipboardInput); i {
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
		file_midgard_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetClipboardOutput); i {
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
		file_midgard_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutClipboardInput); i {
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
		file_midgard_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutClipboardOutput); i {
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
		file_midgard_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURIInput); i {
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
		file_midgard_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetURIOutput); i {
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
			RawDescriptor: file_midgard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_midgard_proto_goTypes,
		DependencyIndexes: file_midgard_proto_depIdxs,
		MessageInfos:      file_midgard_proto_msgTypes,
	}.Build()
	File_midgard_proto = out.File
	file_midgard_proto_rawDesc = nil
	file_midgard_proto_goTypes = nil
	file_midgard_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MidgardClient is the client API for Midgard service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MidgardClient interface {
	Ping(ctx context.Context, in *PingInput, opts ...grpc.CallOption) (*PingOutput, error)
	GetClipboard(ctx context.Context, in *GetClipboardInput, opts ...grpc.CallOption) (*GetClipboardOutput, error)
	PutClipboard(ctx context.Context, in *PutClipboardInput, opts ...grpc.CallOption) (*PutClipboardOutput, error)
	GetURI(ctx context.Context, in *GetURIInput, opts ...grpc.CallOption) (*GetURIOutput, error)
}

type midgardClient struct {
	cc grpc.ClientConnInterface
}

func NewMidgardClient(cc grpc.ClientConnInterface) MidgardClient {
	return &midgardClient{cc}
}

func (c *midgardClient) Ping(ctx context.Context, in *PingInput, opts ...grpc.CallOption) (*PingOutput, error) {
	out := new(PingOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardClient) GetClipboard(ctx context.Context, in *GetClipboardInput, opts ...grpc.CallOption) (*GetClipboardOutput, error) {
	out := new(GetClipboardOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/GetClipboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardClient) PutClipboard(ctx context.Context, in *PutClipboardInput, opts ...grpc.CallOption) (*PutClipboardOutput, error) {
	out := new(PutClipboardOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/PutClipboard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *midgardClient) GetURI(ctx context.Context, in *GetURIInput, opts ...grpc.CallOption) (*GetURIOutput, error) {
	out := new(GetURIOutput)
	err := c.cc.Invoke(ctx, "/proto.Midgard/GetURI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MidgardServer is the server API for Midgard service.
type MidgardServer interface {
	Ping(context.Context, *PingInput) (*PingOutput, error)
	GetClipboard(context.Context, *GetClipboardInput) (*GetClipboardOutput, error)
	PutClipboard(context.Context, *PutClipboardInput) (*PutClipboardOutput, error)
	GetURI(context.Context, *GetURIInput) (*GetURIOutput, error)
}

// UnimplementedMidgardServer can be embedded to have forward compatible implementations.
type UnimplementedMidgardServer struct {
}

func (*UnimplementedMidgardServer) Ping(context.Context, *PingInput) (*PingOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedMidgardServer) GetClipboard(context.Context, *GetClipboardInput) (*GetClipboardOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClipboard not implemented")
}
func (*UnimplementedMidgardServer) PutClipboard(context.Context, *PutClipboardInput) (*PutClipboardOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutClipboard not implemented")
}
func (*UnimplementedMidgardServer) GetURI(context.Context, *GetURIInput) (*GetURIOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetURI not implemented")
}

func RegisterMidgardServer(s *grpc.Server, srv MidgardServer) {
	s.RegisterService(&_Midgard_serviceDesc, srv)
}

func _Midgard_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).Ping(ctx, req.(*PingInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Midgard_GetClipboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClipboardInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).GetClipboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/GetClipboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).GetClipboard(ctx, req.(*GetClipboardInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Midgard_PutClipboard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutClipboardInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).PutClipboard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/PutClipboard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).PutClipboard(ctx, req.(*PutClipboardInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Midgard_GetURI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetURIInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MidgardServer).GetURI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Midgard/GetURI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MidgardServer).GetURI(ctx, req.(*GetURIInput))
	}
	return interceptor(ctx, in, info, handler)
}

var _Midgard_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Midgard",
	HandlerType: (*MidgardServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Midgard_Ping_Handler,
		},
		{
			MethodName: "GetClipboard",
			Handler:    _Midgard_GetClipboard_Handler,
		},
		{
			MethodName: "PutClipboard",
			Handler:    _Midgard_PutClipboard_Handler,
		},
		{
			MethodName: "GetURI",
			Handler:    _Midgard_GetURI_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "midgard.proto",
}