// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: orderer/configuration.proto

package orderer

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

// State defines the orderer mode of operation, typically for consensus-type migration.
// NORMAL is during normal operation, when consensus-type migration is not, and can not, take place.
// MAINTENANCE is when the consensus-type can be changed.
type ConsensusType_State int32

const (
	ConsensusType_STATE_NORMAL      ConsensusType_State = 0
	ConsensusType_STATE_MAINTENANCE ConsensusType_State = 1
)

// Enum value maps for ConsensusType_State.
var (
	ConsensusType_State_name = map[int32]string{
		0: "STATE_NORMAL",
		1: "STATE_MAINTENANCE",
	}
	ConsensusType_State_value = map[string]int32{
		"STATE_NORMAL":      0,
		"STATE_MAINTENANCE": 1,
	}
)

func (x ConsensusType_State) Enum() *ConsensusType_State {
	p := new(ConsensusType_State)
	*p = x
	return p
}

func (x ConsensusType_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConsensusType_State) Descriptor() protoreflect.EnumDescriptor {
	return file_orderer_configuration_proto_enumTypes[0].Descriptor()
}

func (ConsensusType_State) Type() protoreflect.EnumType {
	return &file_orderer_configuration_proto_enumTypes[0]
}

func (x ConsensusType_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConsensusType_State.Descriptor instead.
func (ConsensusType_State) EnumDescriptor() ([]byte, []int) {
	return file_orderer_configuration_proto_rawDescGZIP(), []int{0, 0}
}

type ConsensusType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The consensus type: "solo", "kafka" or "etcdraft".
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Opaque metadata, dependent on the consensus type.
	Metadata []byte `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The state signals the ordering service to go into maintenance mode, typically for consensus-type migration.
	State ConsensusType_State `protobuf:"varint,3,opt,name=state,proto3,enum=orderer.ConsensusType_State" json:"state,omitempty"`
}

func (x *ConsensusType) Reset() {
	*x = ConsensusType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusType) ProtoMessage() {}

func (x *ConsensusType) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusType.ProtoReflect.Descriptor instead.
func (*ConsensusType) Descriptor() ([]byte, []int) {
	return file_orderer_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *ConsensusType) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ConsensusType) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ConsensusType) GetState() ConsensusType_State {
	if x != nil {
		return x.State
	}
	return ConsensusType_STATE_NORMAL
}

type BatchSize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Simply specified as number of messages for now, in the future
	// we may want to allow this to be specified by size in bytes
	MaxMessageCount uint32 `protobuf:"varint,1,opt,name=max_message_count,json=maxMessageCount,proto3" json:"max_message_count,omitempty"`
	// The byte count of the serialized messages in a batch cannot
	// exceed this value.
	AbsoluteMaxBytes uint32 `protobuf:"varint,2,opt,name=absolute_max_bytes,json=absoluteMaxBytes,proto3" json:"absolute_max_bytes,omitempty"`
	// The byte count of the serialized messages in a batch should not
	// exceed this value.
	PreferredMaxBytes uint32 `protobuf:"varint,3,opt,name=preferred_max_bytes,json=preferredMaxBytes,proto3" json:"preferred_max_bytes,omitempty"`
}

func (x *BatchSize) Reset() {
	*x = BatchSize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchSize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchSize) ProtoMessage() {}

func (x *BatchSize) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchSize.ProtoReflect.Descriptor instead.
func (*BatchSize) Descriptor() ([]byte, []int) {
	return file_orderer_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *BatchSize) GetMaxMessageCount() uint32 {
	if x != nil {
		return x.MaxMessageCount
	}
	return 0
}

func (x *BatchSize) GetAbsoluteMaxBytes() uint32 {
	if x != nil {
		return x.AbsoluteMaxBytes
	}
	return 0
}

func (x *BatchSize) GetPreferredMaxBytes() uint32 {
	if x != nil {
		return x.PreferredMaxBytes
	}
	return 0
}

type BatchTimeout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Any duration string parseable by ParseDuration():
	// https://golang.org/pkg/time/#ParseDuration
	Timeout string `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *BatchTimeout) Reset() {
	*x = BatchTimeout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchTimeout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchTimeout) ProtoMessage() {}

func (x *BatchTimeout) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchTimeout.ProtoReflect.Descriptor instead.
func (*BatchTimeout) Descriptor() ([]byte, []int) {
	return file_orderer_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *BatchTimeout) GetTimeout() string {
	if x != nil {
		return x.Timeout
	}
	return ""
}

// Carries a list of bootstrap brokers, i.e. this is not the exclusive set of
// brokers an ordering service
type KafkaBrokers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Each broker here should be identified using the (IP|host):port notation,
	// e.g. 127.0.0.1:7050, or localhost:7050 are valid entries
	Brokers []string `protobuf:"bytes,1,rep,name=brokers,proto3" json:"brokers,omitempty"`
}

func (x *KafkaBrokers) Reset() {
	*x = KafkaBrokers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_configuration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaBrokers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaBrokers) ProtoMessage() {}

func (x *KafkaBrokers) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_configuration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaBrokers.ProtoReflect.Descriptor instead.
func (*KafkaBrokers) Descriptor() ([]byte, []int) {
	return file_orderer_configuration_proto_rawDescGZIP(), []int{3}
}

func (x *KafkaBrokers) GetBrokers() []string {
	if x != nil {
		return x.Brokers
	}
	return nil
}

// ChannelRestrictions is the mssage which conveys restrictions on channel creation for an orderer
type ChannelRestrictions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxCount uint64 `protobuf:"varint,1,opt,name=max_count,json=maxCount,proto3" json:"max_count,omitempty"` // The max count of channels to allow to be created, a value of 0 indicates no limit
}

func (x *ChannelRestrictions) Reset() {
	*x = ChannelRestrictions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_configuration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelRestrictions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelRestrictions) ProtoMessage() {}

func (x *ChannelRestrictions) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_configuration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelRestrictions.ProtoReflect.Descriptor instead.
func (*ChannelRestrictions) Descriptor() ([]byte, []int) {
	return file_orderer_configuration_proto_rawDescGZIP(), []int{4}
}

func (x *ChannelRestrictions) GetMaxCount() uint64 {
	if x != nil {
		return x.MaxCount
	}
	return 0
}

var File_orderer_configuration_proto protoreflect.FileDescriptor

var file_orderer_configuration_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x22, 0xa5, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65,
	0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x30, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e,
	0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x01, 0x22, 0x95,
	0x01, 0x0a, 0x09, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x62, 0x73, 0x6f,
	0x6c, 0x75, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x61, 0x62, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x65, 0x4d, 0x61,
	0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x72, 0x65, 0x64, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x4d, 0x61,
	0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x28, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x22, 0x28, 0x0a, 0x0c, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x73, 0x22, 0x32, 0x0a, 0x13, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xae,
	0x01, 0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67,
	0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x42, 0x12, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72,
	0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x4f, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x72, 0xca, 0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0xe2,
	0x02, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderer_configuration_proto_rawDescOnce sync.Once
	file_orderer_configuration_proto_rawDescData = file_orderer_configuration_proto_rawDesc
)

func file_orderer_configuration_proto_rawDescGZIP() []byte {
	file_orderer_configuration_proto_rawDescOnce.Do(func() {
		file_orderer_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderer_configuration_proto_rawDescData)
	})
	return file_orderer_configuration_proto_rawDescData
}

var file_orderer_configuration_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orderer_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_orderer_configuration_proto_goTypes = []interface{}{
	(ConsensusType_State)(0),    // 0: orderer.ConsensusType.State
	(*ConsensusType)(nil),       // 1: orderer.ConsensusType
	(*BatchSize)(nil),           // 2: orderer.BatchSize
	(*BatchTimeout)(nil),        // 3: orderer.BatchTimeout
	(*KafkaBrokers)(nil),        // 4: orderer.KafkaBrokers
	(*ChannelRestrictions)(nil), // 5: orderer.ChannelRestrictions
}
var file_orderer_configuration_proto_depIdxs = []int32{
	0, // 0: orderer.ConsensusType.state:type_name -> orderer.ConsensusType.State
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_orderer_configuration_proto_init() }
func file_orderer_configuration_proto_init() {
	if File_orderer_configuration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderer_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusType); i {
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
		file_orderer_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchSize); i {
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
		file_orderer_configuration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchTimeout); i {
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
		file_orderer_configuration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaBrokers); i {
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
		file_orderer_configuration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelRestrictions); i {
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
			RawDescriptor: file_orderer_configuration_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderer_configuration_proto_goTypes,
		DependencyIndexes: file_orderer_configuration_proto_depIdxs,
		EnumInfos:         file_orderer_configuration_proto_enumTypes,
		MessageInfos:      file_orderer_configuration_proto_msgTypes,
	}.Build()
	File_orderer_configuration_proto = out.File
	file_orderer_configuration_proto_rawDesc = nil
	file_orderer_configuration_proto_goTypes = nil
	file_orderer_configuration_proto_depIdxs = nil
}
