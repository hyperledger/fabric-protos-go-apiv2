// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: orderer/etcdraft/metadata.proto

package etcdraft

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

// BlockMetadata stores data used by the Raft OSNs when
// coordinating with each other, to be serialized into
// block meta dta field and used after failres and restarts.
type BlockMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Maintains a mapping between the cluster's OSNs
	// and their Raft IDs.
	ConsenterIds []uint64 `protobuf:"varint,1,rep,packed,name=consenter_ids,json=consenterIds,proto3" json:"consenter_ids,omitempty"`
	// Carries the Raft ID value that will be assigned
	// to the next OSN that will join this cluster.
	NextConsenterId uint64 `protobuf:"varint,2,opt,name=next_consenter_id,json=nextConsenterId,proto3" json:"next_consenter_id,omitempty"`
	// Index of etcd/raft entry for current block.
	RaftIndex uint64 `protobuf:"varint,3,opt,name=raft_index,json=raftIndex,proto3" json:"raft_index,omitempty"`
}

func (x *BlockMetadata) Reset() {
	*x = BlockMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_etcdraft_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockMetadata) ProtoMessage() {}

func (x *BlockMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_etcdraft_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockMetadata.ProtoReflect.Descriptor instead.
func (*BlockMetadata) Descriptor() ([]byte, []int) {
	return file_orderer_etcdraft_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *BlockMetadata) GetConsenterIds() []uint64 {
	if x != nil {
		return x.ConsenterIds
	}
	return nil
}

func (x *BlockMetadata) GetNextConsenterId() uint64 {
	if x != nil {
		return x.NextConsenterId
	}
	return 0
}

func (x *BlockMetadata) GetRaftIndex() uint64 {
	if x != nil {
		return x.RaftIndex
	}
	return 0
}

// ClusterMetadata encapsulates metadata that is exchanged among cluster nodes
type ClusterMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Indicates active nodes in cluster that are reacheable by Raft leader
	ActiveNodes []uint64 `protobuf:"varint,1,rep,packed,name=active_nodes,json=activeNodes,proto3" json:"active_nodes,omitempty"`
}

func (x *ClusterMetadata) Reset() {
	*x = ClusterMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_etcdraft_metadata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterMetadata) ProtoMessage() {}

func (x *ClusterMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_etcdraft_metadata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterMetadata.ProtoReflect.Descriptor instead.
func (*ClusterMetadata) Descriptor() ([]byte, []int) {
	return file_orderer_etcdraft_metadata_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterMetadata) GetActiveNodes() []uint64 {
	if x != nil {
		return x.ActiveNodes
	}
	return nil
}

var File_orderer_etcdraft_metadata_proto protoreflect.FileDescriptor

var file_orderer_etcdraft_metadata_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x65, 0x74, 0x63, 0x64, 0x72, 0x61,
	0x66, 0x74, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x65, 0x74, 0x63, 0x64, 0x72, 0x61, 0x66, 0x74, 0x22, 0x7f, 0x0a, 0x0d, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x61, 0x66, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x09, 0x72, 0x61, 0x66, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x34, 0x0a, 0x0f,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x73, 0x42, 0xbf, 0x01, 0x0a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x65, 0x74, 0x63,
	0x64, 0x72, 0x61, 0x66, 0x74, 0x42, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66,
	0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x65, 0x74,
	0x63, 0x64, 0x72, 0x61, 0x66, 0x74, 0xa2, 0x02, 0x03, 0x45, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x45,
	0x74, 0x63, 0x64, 0x72, 0x61, 0x66, 0x74, 0xca, 0x02, 0x08, 0x45, 0x74, 0x63, 0x64, 0x72, 0x61,
	0x66, 0x74, 0xe2, 0x02, 0x14, 0x45, 0x74, 0x63, 0x64, 0x72, 0x61, 0x66, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x45, 0x74, 0x63, 0x64,
	0x72, 0x61, 0x66, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderer_etcdraft_metadata_proto_rawDescOnce sync.Once
	file_orderer_etcdraft_metadata_proto_rawDescData = file_orderer_etcdraft_metadata_proto_rawDesc
)

func file_orderer_etcdraft_metadata_proto_rawDescGZIP() []byte {
	file_orderer_etcdraft_metadata_proto_rawDescOnce.Do(func() {
		file_orderer_etcdraft_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderer_etcdraft_metadata_proto_rawDescData)
	})
	return file_orderer_etcdraft_metadata_proto_rawDescData
}

var file_orderer_etcdraft_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_orderer_etcdraft_metadata_proto_goTypes = []interface{}{
	(*BlockMetadata)(nil),   // 0: etcdraft.BlockMetadata
	(*ClusterMetadata)(nil), // 1: etcdraft.ClusterMetadata
}
var file_orderer_etcdraft_metadata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orderer_etcdraft_metadata_proto_init() }
func file_orderer_etcdraft_metadata_proto_init() {
	if File_orderer_etcdraft_metadata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderer_etcdraft_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockMetadata); i {
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
		file_orderer_etcdraft_metadata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterMetadata); i {
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
			RawDescriptor: file_orderer_etcdraft_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderer_etcdraft_metadata_proto_goTypes,
		DependencyIndexes: file_orderer_etcdraft_metadata_proto_depIdxs,
		MessageInfos:      file_orderer_etcdraft_metadata_proto_msgTypes,
	}.Build()
	File_orderer_etcdraft_metadata_proto = out.File
	file_orderer_etcdraft_metadata_proto_rawDesc = nil
	file_orderer_etcdraft_metadata_proto_goTypes = nil
	file_orderer_etcdraft_metadata_proto_depIdxs = nil
}
