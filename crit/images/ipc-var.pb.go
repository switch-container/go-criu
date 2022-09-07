// SPDX-License-Identifier: MIT

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: ipc-var.proto

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

type IpcVarEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SemCtls          []uint32 `protobuf:"varint,1,rep,name=sem_ctls,json=semCtls" json:"sem_ctls,omitempty"`
	MsgCtlmax        *uint32  `protobuf:"varint,2,req,name=msg_ctlmax,json=msgCtlmax" json:"msg_ctlmax,omitempty"`
	MsgCtlmnb        *uint32  `protobuf:"varint,3,req,name=msg_ctlmnb,json=msgCtlmnb" json:"msg_ctlmnb,omitempty"`
	MsgCtlmni        *uint32  `protobuf:"varint,4,req,name=msg_ctlmni,json=msgCtlmni" json:"msg_ctlmni,omitempty"`
	AutoMsgmni       *uint32  `protobuf:"varint,5,req,name=auto_msgmni,json=autoMsgmni" json:"auto_msgmni,omitempty"`
	ShmCtlmax        *uint64  `protobuf:"varint,6,req,name=shm_ctlmax,json=shmCtlmax" json:"shm_ctlmax,omitempty"`
	ShmCtlall        *uint64  `protobuf:"varint,7,req,name=shm_ctlall,json=shmCtlall" json:"shm_ctlall,omitempty"`
	ShmCtlmni        *uint32  `protobuf:"varint,8,req,name=shm_ctlmni,json=shmCtlmni" json:"shm_ctlmni,omitempty"`
	ShmRmidForced    *uint32  `protobuf:"varint,9,req,name=shm_rmid_forced,json=shmRmidForced" json:"shm_rmid_forced,omitempty"`
	MqQueuesMax      *uint32  `protobuf:"varint,10,req,name=mq_queues_max,json=mqQueuesMax" json:"mq_queues_max,omitempty"`
	MqMsgMax         *uint32  `protobuf:"varint,11,req,name=mq_msg_max,json=mqMsgMax" json:"mq_msg_max,omitempty"`
	MqMsgsizeMax     *uint32  `protobuf:"varint,12,req,name=mq_msgsize_max,json=mqMsgsizeMax" json:"mq_msgsize_max,omitempty"`
	MqMsgDefault     *uint32  `protobuf:"varint,13,opt,name=mq_msg_default,json=mqMsgDefault" json:"mq_msg_default,omitempty"`
	MqMsgsizeDefault *uint32  `protobuf:"varint,14,opt,name=mq_msgsize_default,json=mqMsgsizeDefault" json:"mq_msgsize_default,omitempty"`
	MsgNextId        *uint32  `protobuf:"varint,15,opt,name=msg_next_id,json=msgNextId" json:"msg_next_id,omitempty"`
	SemNextId        *uint32  `protobuf:"varint,16,opt,name=sem_next_id,json=semNextId" json:"sem_next_id,omitempty"`
	ShmNextId        *uint32  `protobuf:"varint,17,opt,name=shm_next_id,json=shmNextId" json:"shm_next_id,omitempty"`
}

func (x *IpcVarEntry) Reset() {
	*x = IpcVarEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ipc_var_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpcVarEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpcVarEntry) ProtoMessage() {}

func (x *IpcVarEntry) ProtoReflect() protoreflect.Message {
	mi := &file_ipc_var_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpcVarEntry.ProtoReflect.Descriptor instead.
func (*IpcVarEntry) Descriptor() ([]byte, []int) {
	return file_ipc_var_proto_rawDescGZIP(), []int{0}
}

func (x *IpcVarEntry) GetSemCtls() []uint32 {
	if x != nil {
		return x.SemCtls
	}
	return nil
}

func (x *IpcVarEntry) GetMsgCtlmax() uint32 {
	if x != nil && x.MsgCtlmax != nil {
		return *x.MsgCtlmax
	}
	return 0
}

func (x *IpcVarEntry) GetMsgCtlmnb() uint32 {
	if x != nil && x.MsgCtlmnb != nil {
		return *x.MsgCtlmnb
	}
	return 0
}

func (x *IpcVarEntry) GetMsgCtlmni() uint32 {
	if x != nil && x.MsgCtlmni != nil {
		return *x.MsgCtlmni
	}
	return 0
}

func (x *IpcVarEntry) GetAutoMsgmni() uint32 {
	if x != nil && x.AutoMsgmni != nil {
		return *x.AutoMsgmni
	}
	return 0
}

func (x *IpcVarEntry) GetShmCtlmax() uint64 {
	if x != nil && x.ShmCtlmax != nil {
		return *x.ShmCtlmax
	}
	return 0
}

func (x *IpcVarEntry) GetShmCtlall() uint64 {
	if x != nil && x.ShmCtlall != nil {
		return *x.ShmCtlall
	}
	return 0
}

func (x *IpcVarEntry) GetShmCtlmni() uint32 {
	if x != nil && x.ShmCtlmni != nil {
		return *x.ShmCtlmni
	}
	return 0
}

func (x *IpcVarEntry) GetShmRmidForced() uint32 {
	if x != nil && x.ShmRmidForced != nil {
		return *x.ShmRmidForced
	}
	return 0
}

func (x *IpcVarEntry) GetMqQueuesMax() uint32 {
	if x != nil && x.MqQueuesMax != nil {
		return *x.MqQueuesMax
	}
	return 0
}

func (x *IpcVarEntry) GetMqMsgMax() uint32 {
	if x != nil && x.MqMsgMax != nil {
		return *x.MqMsgMax
	}
	return 0
}

func (x *IpcVarEntry) GetMqMsgsizeMax() uint32 {
	if x != nil && x.MqMsgsizeMax != nil {
		return *x.MqMsgsizeMax
	}
	return 0
}

func (x *IpcVarEntry) GetMqMsgDefault() uint32 {
	if x != nil && x.MqMsgDefault != nil {
		return *x.MqMsgDefault
	}
	return 0
}

func (x *IpcVarEntry) GetMqMsgsizeDefault() uint32 {
	if x != nil && x.MqMsgsizeDefault != nil {
		return *x.MqMsgsizeDefault
	}
	return 0
}

func (x *IpcVarEntry) GetMsgNextId() uint32 {
	if x != nil && x.MsgNextId != nil {
		return *x.MsgNextId
	}
	return 0
}

func (x *IpcVarEntry) GetSemNextId() uint32 {
	if x != nil && x.SemNextId != nil {
		return *x.SemNextId
	}
	return 0
}

func (x *IpcVarEntry) GetShmNextId() uint32 {
	if x != nil && x.ShmNextId != nil {
		return *x.ShmNextId
	}
	return 0
}

var File_ipc_var_proto protoreflect.FileDescriptor

var file_ipc_var_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x69, 0x70, 0x63, 0x2d, 0x76, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc9, 0x04, 0x0a, 0x0d, 0x69, 0x70, 0x63, 0x5f, 0x76, 0x61, 0x72, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6d, 0x5f, 0x63, 0x74, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x65, 0x6d, 0x43, 0x74, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x73, 0x67, 0x5f, 0x63, 0x74, 0x6c, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x09, 0x6d, 0x73, 0x67, 0x43, 0x74, 0x6c, 0x6d, 0x61, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x6d,
	0x73, 0x67, 0x5f, 0x63, 0x74, 0x6c, 0x6d, 0x6e, 0x62, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52,
	0x09, 0x6d, 0x73, 0x67, 0x43, 0x74, 0x6c, 0x6d, 0x6e, 0x62, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x73,
	0x67, 0x5f, 0x63, 0x74, 0x6c, 0x6d, 0x6e, 0x69, 0x18, 0x04, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09,
	0x6d, 0x73, 0x67, 0x43, 0x74, 0x6c, 0x6d, 0x6e, 0x69, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74,
	0x6f, 0x5f, 0x6d, 0x73, 0x67, 0x6d, 0x6e, 0x69, 0x18, 0x05, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0a,
	0x61, 0x75, 0x74, 0x6f, 0x4d, 0x73, 0x67, 0x6d, 0x6e, 0x69, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68,
	0x6d, 0x5f, 0x63, 0x74, 0x6c, 0x6d, 0x61, 0x78, 0x18, 0x06, 0x20, 0x02, 0x28, 0x04, 0x52, 0x09,
	0x73, 0x68, 0x6d, 0x43, 0x74, 0x6c, 0x6d, 0x61, 0x78, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6d,
	0x5f, 0x63, 0x74, 0x6c, 0x61, 0x6c, 0x6c, 0x18, 0x07, 0x20, 0x02, 0x28, 0x04, 0x52, 0x09, 0x73,
	0x68, 0x6d, 0x43, 0x74, 0x6c, 0x61, 0x6c, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68, 0x6d, 0x5f,
	0x63, 0x74, 0x6c, 0x6d, 0x6e, 0x69, 0x18, 0x08, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x68,
	0x6d, 0x43, 0x74, 0x6c, 0x6d, 0x6e, 0x69, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x68, 0x6d, 0x5f, 0x72,
	0x6d, 0x69, 0x64, 0x5f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x18, 0x09, 0x20, 0x02, 0x28, 0x0d,
	0x52, 0x0d, 0x73, 0x68, 0x6d, 0x52, 0x6d, 0x69, 0x64, 0x46, 0x6f, 0x72, 0x63, 0x65, 0x64, 0x12,
	0x22, 0x0a, 0x0d, 0x6d, 0x71, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x73, 0x5f, 0x6d, 0x61, 0x78,
	0x18, 0x0a, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x71, 0x51, 0x75, 0x65, 0x75, 0x65, 0x73,
	0x4d, 0x61, 0x78, 0x12, 0x1c, 0x0a, 0x0a, 0x6d, 0x71, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x6d, 0x61,
	0x78, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x71, 0x4d, 0x73, 0x67, 0x4d, 0x61,
	0x78, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x71, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x69, 0x7a, 0x65, 0x5f,
	0x6d, 0x61, 0x78, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0c, 0x6d, 0x71, 0x4d, 0x73, 0x67,
	0x73, 0x69, 0x7a, 0x65, 0x4d, 0x61, 0x78, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x71, 0x5f, 0x6d, 0x73,
	0x67, 0x5f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0c, 0x6d, 0x71, 0x4d, 0x73, 0x67, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x2c, 0x0a,
	0x12, 0x6d, 0x71, 0x5f, 0x6d, 0x73, 0x67, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6d, 0x71, 0x4d, 0x73, 0x67,
	0x73, 0x69, 0x7a, 0x65, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x1e, 0x0a, 0x0b, 0x6d,
	0x73, 0x67, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x6d, 0x73, 0x67, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x73,
	0x65, 0x6d, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x73, 0x65, 0x6d, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0b, 0x73,
	0x68, 0x6d, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x73, 0x68, 0x6d, 0x4e, 0x65, 0x78, 0x74, 0x49, 0x64,
}

var (
	file_ipc_var_proto_rawDescOnce sync.Once
	file_ipc_var_proto_rawDescData = file_ipc_var_proto_rawDesc
)

func file_ipc_var_proto_rawDescGZIP() []byte {
	file_ipc_var_proto_rawDescOnce.Do(func() {
		file_ipc_var_proto_rawDescData = protoimpl.X.CompressGZIP(file_ipc_var_proto_rawDescData)
	})
	return file_ipc_var_proto_rawDescData
}

var file_ipc_var_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ipc_var_proto_goTypes = []interface{}{
	(*IpcVarEntry)(nil), // 0: ipc_var_entry
}
var file_ipc_var_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ipc_var_proto_init() }
func file_ipc_var_proto_init() {
	if File_ipc_var_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ipc_var_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpcVarEntry); i {
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
			RawDescriptor: file_ipc_var_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ipc_var_proto_goTypes,
		DependencyIndexes: file_ipc_var_proto_depIdxs,
		MessageInfos:      file_ipc_var_proto_msgTypes,
	}.Build()
	File_ipc_var_proto = out.File
	file_ipc_var_proto_rawDesc = nil
	file_ipc_var_proto_goTypes = nil
	file_ipc_var_proto_depIdxs = nil
}
