// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: pagemap.proto

package images

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

type PagemapHead struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PagesId *uint32 `protobuf:"varint,1,req,name=pages_id,json=pagesId" json:"pages_id,omitempty"`
}

func (x *PagemapHead) Reset() {
	*x = PagemapHead{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagemap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagemapHead) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagemapHead) ProtoMessage() {}

func (x *PagemapHead) ProtoReflect() protoreflect.Message {
	mi := &file_pagemap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagemapHead.ProtoReflect.Descriptor instead.
func (*PagemapHead) Descriptor() ([]byte, []int) {
	return file_pagemap_proto_rawDescGZIP(), []int{0}
}

func (x *PagemapHead) GetPagesId() uint32 {
	if x != nil && x.PagesId != nil {
		return *x.PagesId
	}
	return 0
}

type PagemapEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vaddr    *uint64 `protobuf:"varint,1,req,name=vaddr" json:"vaddr,omitempty"`
	NrPages  *uint32 `protobuf:"varint,2,req,name=nr_pages,json=nrPages" json:"nr_pages,omitempty"`
	InParent *bool   `protobuf:"varint,3,opt,name=in_parent,json=inParent" json:"in_parent,omitempty"`
	Flags    *uint32 `protobuf:"varint,4,opt,name=flags" json:"flags,omitempty"`
}

func (x *PagemapEntry) Reset() {
	*x = PagemapEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pagemap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagemapEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagemapEntry) ProtoMessage() {}

func (x *PagemapEntry) ProtoReflect() protoreflect.Message {
	mi := &file_pagemap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagemapEntry.ProtoReflect.Descriptor instead.
func (*PagemapEntry) Descriptor() ([]byte, []int) {
	return file_pagemap_proto_rawDescGZIP(), []int{1}
}

func (x *PagemapEntry) GetVaddr() uint64 {
	if x != nil && x.Vaddr != nil {
		return *x.Vaddr
	}
	return 0
}

func (x *PagemapEntry) GetNrPages() uint32 {
	if x != nil && x.NrPages != nil {
		return *x.NrPages
	}
	return 0
}

func (x *PagemapEntry) GetInParent() bool {
	if x != nil && x.InParent != nil {
		return *x.InParent
	}
	return false
}

func (x *PagemapEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

var File_pagemap_proto protoreflect.FileDescriptor

var file_pagemap_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x67, 0x65, 0x6d, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0a, 0x6f, 0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x0c, 0x70,
	0x61, 0x67, 0x65, 0x6d, 0x61, 0x70, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x70,
	0x61, 0x67, 0x65, 0x73, 0x49, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x0d, 0x70, 0x61, 0x67, 0x65, 0x6d,
	0x61, 0x70, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1b, 0x0a, 0x05, 0x76, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x42, 0x05, 0xd2, 0x3f, 0x02, 0x08, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x64, 0x64, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x6e, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x6e, 0x72, 0x50, 0x61, 0x67, 0x65, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a,
	0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0f, 0xd2, 0x3f,
	0x0c, 0x1a, 0x0a, 0x70, 0x6d, 0x61, 0x70, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x05, 0x66,
	0x6c, 0x61, 0x67, 0x73,
}

var (
	file_pagemap_proto_rawDescOnce sync.Once
	file_pagemap_proto_rawDescData = file_pagemap_proto_rawDesc
)

func file_pagemap_proto_rawDescGZIP() []byte {
	file_pagemap_proto_rawDescOnce.Do(func() {
		file_pagemap_proto_rawDescData = protoimpl.X.CompressGZIP(file_pagemap_proto_rawDescData)
	})
	return file_pagemap_proto_rawDescData
}

var file_pagemap_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pagemap_proto_goTypes = []interface{}{
	(*PagemapHead)(nil),  // 0: pagemap_head
	(*PagemapEntry)(nil), // 1: pagemap_entry
}
var file_pagemap_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pagemap_proto_init() }
func file_pagemap_proto_init() {
	if File_pagemap_proto != nil {
		return
	}
	file_opts_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pagemap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagemapHead); i {
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
		file_pagemap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagemapEntry); i {
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
			RawDescriptor: file_pagemap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pagemap_proto_goTypes,
		DependencyIndexes: file_pagemap_proto_depIdxs,
		MessageInfos:      file_pagemap_proto_msgTypes,
	}.Build()
	File_pagemap_proto = out.File
	file_pagemap_proto_rawDesc = nil
	file_pagemap_proto_goTypes = nil
	file_pagemap_proto_depIdxs = nil
}
