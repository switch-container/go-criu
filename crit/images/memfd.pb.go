// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: memfd.proto

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

type MemfdFileEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *uint32    `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Flags   *uint32    `protobuf:"varint,2,req,name=flags" json:"flags,omitempty"`
	Pos     *uint64    `protobuf:"varint,3,req,name=pos" json:"pos,omitempty"`
	Fown    *FownEntry `protobuf:"bytes,4,req,name=fown" json:"fown,omitempty"`
	InodeId *uint32    `protobuf:"varint,5,req,name=inode_id,json=inodeId" json:"inode_id,omitempty"`
}

func (x *MemfdFileEntry) Reset() {
	*x = MemfdFileEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memfd_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemfdFileEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemfdFileEntry) ProtoMessage() {}

func (x *MemfdFileEntry) ProtoReflect() protoreflect.Message {
	mi := &file_memfd_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemfdFileEntry.ProtoReflect.Descriptor instead.
func (*MemfdFileEntry) Descriptor() ([]byte, []int) {
	return file_memfd_proto_rawDescGZIP(), []int{0}
}

func (x *MemfdFileEntry) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *MemfdFileEntry) GetFlags() uint32 {
	if x != nil && x.Flags != nil {
		return *x.Flags
	}
	return 0
}

func (x *MemfdFileEntry) GetPos() uint64 {
	if x != nil && x.Pos != nil {
		return *x.Pos
	}
	return 0
}

func (x *MemfdFileEntry) GetFown() *FownEntry {
	if x != nil {
		return x.Fown
	}
	return nil
}

func (x *MemfdFileEntry) GetInodeId() uint32 {
	if x != nil && x.InodeId != nil {
		return *x.InodeId
	}
	return 0
}

type MemfdInodeEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Uid         *uint32 `protobuf:"varint,2,req,name=uid" json:"uid,omitempty"`
	Gid         *uint32 `protobuf:"varint,3,req,name=gid" json:"gid,omitempty"`
	Size        *uint64 `protobuf:"varint,4,req,name=size" json:"size,omitempty"`
	Shmid       *uint32 `protobuf:"varint,5,req,name=shmid" json:"shmid,omitempty"`
	Seals       *uint32 `protobuf:"varint,6,req,name=seals" json:"seals,omitempty"`
	InodeId     *uint64 `protobuf:"varint,7,req,name=inode_id,json=inodeId" json:"inode_id,omitempty"`
	HugetlbFlag *uint32 `protobuf:"varint,8,opt,name=hugetlb_flag,json=hugetlbFlag" json:"hugetlb_flag,omitempty"`
}

func (x *MemfdInodeEntry) Reset() {
	*x = MemfdInodeEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_memfd_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemfdInodeEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemfdInodeEntry) ProtoMessage() {}

func (x *MemfdInodeEntry) ProtoReflect() protoreflect.Message {
	mi := &file_memfd_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemfdInodeEntry.ProtoReflect.Descriptor instead.
func (*MemfdInodeEntry) Descriptor() ([]byte, []int) {
	return file_memfd_proto_rawDescGZIP(), []int{1}
}

func (x *MemfdInodeEntry) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *MemfdInodeEntry) GetUid() uint32 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *MemfdInodeEntry) GetGid() uint32 {
	if x != nil && x.Gid != nil {
		return *x.Gid
	}
	return 0
}

func (x *MemfdInodeEntry) GetSize() uint64 {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return 0
}

func (x *MemfdInodeEntry) GetShmid() uint32 {
	if x != nil && x.Shmid != nil {
		return *x.Shmid
	}
	return 0
}

func (x *MemfdInodeEntry) GetSeals() uint32 {
	if x != nil && x.Seals != nil {
		return *x.Seals
	}
	return 0
}

func (x *MemfdInodeEntry) GetInodeId() uint64 {
	if x != nil && x.InodeId != nil {
		return *x.InodeId
	}
	return 0
}

func (x *MemfdInodeEntry) GetHugetlbFlag() uint32 {
	if x != nil && x.HugetlbFlag != nil {
		return *x.HugetlbFlag
	}
	return 0
}

var File_memfd_proto protoreflect.FileDescriptor

var file_memfd_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x66, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x6f,
	0x70, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x66, 0x6f, 0x77, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x10, 0x6d, 0x65, 0x6d, 0x66, 0x64, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x05, 0x66, 0x6c,
	0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x42, 0x10, 0xd2, 0x3f, 0x0d, 0x1a, 0x0b,
	0x72, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x05, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x02, 0x28, 0x04, 0x52,
	0x03, 0x70, 0x6f, 0x73, 0x12, 0x1f, 0x0a, 0x04, 0x66, 0x6f, 0x77, 0x6e, 0x18, 0x04, 0x20, 0x02,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x6f, 0x77, 0x6e, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x04, 0x66, 0x6f, 0x77, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x07, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x22, 0xdb, 0x01, 0x0a, 0x11, 0x6d, 0x65, 0x6d, 0x66, 0x64, 0x5f, 0x69, 0x6e, 0x6f, 0x64, 0x65,
	0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x67, 0x69, 0x64, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28, 0x04, 0x52, 0x04, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x68, 0x6d, 0x69, 0x64, 0x18, 0x05, 0x20, 0x02, 0x28,
	0x0d, 0x52, 0x05, 0x73, 0x68, 0x6d, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x05, 0x73, 0x65, 0x61, 0x6c,
	0x73, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x42, 0x10, 0xd2, 0x3f, 0x0d, 0x1a, 0x0b, 0x73, 0x65,
	0x61, 0x6c, 0x73, 0x2e, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x05, 0x73, 0x65, 0x61, 0x6c, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x02,
	0x28, 0x04, 0x52, 0x07, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x68,
	0x75, 0x67, 0x65, 0x74, 0x6c, 0x62, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x68, 0x75, 0x67, 0x65, 0x74, 0x6c, 0x62, 0x46, 0x6c, 0x61, 0x67,
}

var (
	file_memfd_proto_rawDescOnce sync.Once
	file_memfd_proto_rawDescData = file_memfd_proto_rawDesc
)

func file_memfd_proto_rawDescGZIP() []byte {
	file_memfd_proto_rawDescOnce.Do(func() {
		file_memfd_proto_rawDescData = protoimpl.X.CompressGZIP(file_memfd_proto_rawDescData)
	})
	return file_memfd_proto_rawDescData
}

var file_memfd_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_memfd_proto_goTypes = []interface{}{
	(*MemfdFileEntry)(nil),  // 0: memfd_file_entry
	(*MemfdInodeEntry)(nil), // 1: memfd_inode_entry
	(*FownEntry)(nil),       // 2: fown_entry
}
var file_memfd_proto_depIdxs = []int32{
	2, // 0: memfd_file_entry.fown:type_name -> fown_entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_memfd_proto_init() }
func file_memfd_proto_init() {
	if File_memfd_proto != nil {
		return
	}
	file_opts_proto_init()
	file_fown_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_memfd_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemfdFileEntry); i {
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
		file_memfd_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemfdInodeEntry); i {
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
			RawDescriptor: file_memfd_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_memfd_proto_goTypes,
		DependencyIndexes: file_memfd_proto_depIdxs,
		MessageInfos:      file_memfd_proto_msgTypes,
	}.Build()
	File_memfd_proto = out.File
	file_memfd_proto_rawDesc = nil
	file_memfd_proto_goTypes = nil
	file_memfd_proto_depIdxs = nil
}
