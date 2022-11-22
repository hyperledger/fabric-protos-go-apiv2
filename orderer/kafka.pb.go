// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: orderer/kafka.proto

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

type KafkaMessageRegular_Class int32

const (
	KafkaMessageRegular_UNKNOWN KafkaMessageRegular_Class = 0
	KafkaMessageRegular_NORMAL  KafkaMessageRegular_Class = 1
	KafkaMessageRegular_CONFIG  KafkaMessageRegular_Class = 2
)

// Enum value maps for KafkaMessageRegular_Class.
var (
	KafkaMessageRegular_Class_name = map[int32]string{
		0: "UNKNOWN",
		1: "NORMAL",
		2: "CONFIG",
	}
	KafkaMessageRegular_Class_value = map[string]int32{
		"UNKNOWN": 0,
		"NORMAL":  1,
		"CONFIG":  2,
	}
)

func (x KafkaMessageRegular_Class) Enum() *KafkaMessageRegular_Class {
	p := new(KafkaMessageRegular_Class)
	*p = x
	return p
}

func (x KafkaMessageRegular_Class) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (KafkaMessageRegular_Class) Descriptor() protoreflect.EnumDescriptor {
	return file_orderer_kafka_proto_enumTypes[0].Descriptor()
}

func (KafkaMessageRegular_Class) Type() protoreflect.EnumType {
	return &file_orderer_kafka_proto_enumTypes[0]
}

func (x KafkaMessageRegular_Class) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use KafkaMessageRegular_Class.Descriptor instead.
func (KafkaMessageRegular_Class) EnumDescriptor() ([]byte, []int) {
	return file_orderer_kafka_proto_rawDescGZIP(), []int{1, 0}
}

// KafkaMessage is a wrapper type for the messages
// that the Kafka-based orderer deals with.
type KafkaMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*KafkaMessage_Regular
	//	*KafkaMessage_TimeToCut
	//	*KafkaMessage_Connect
	Type isKafkaMessage_Type `protobuf_oneof:"Type"`
}

func (x *KafkaMessage) Reset() {
	*x = KafkaMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_kafka_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaMessage) ProtoMessage() {}

func (x *KafkaMessage) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_kafka_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaMessage.ProtoReflect.Descriptor instead.
func (*KafkaMessage) Descriptor() ([]byte, []int) {
	return file_orderer_kafka_proto_rawDescGZIP(), []int{0}
}

func (m *KafkaMessage) GetType() isKafkaMessage_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *KafkaMessage) GetRegular() *KafkaMessageRegular {
	if x, ok := x.GetType().(*KafkaMessage_Regular); ok {
		return x.Regular
	}
	return nil
}

func (x *KafkaMessage) GetTimeToCut() *KafkaMessageTimeToCut {
	if x, ok := x.GetType().(*KafkaMessage_TimeToCut); ok {
		return x.TimeToCut
	}
	return nil
}

func (x *KafkaMessage) GetConnect() *KafkaMessageConnect {
	if x, ok := x.GetType().(*KafkaMessage_Connect); ok {
		return x.Connect
	}
	return nil
}

type isKafkaMessage_Type interface {
	isKafkaMessage_Type()
}

type KafkaMessage_Regular struct {
	Regular *KafkaMessageRegular `protobuf:"bytes,1,opt,name=regular,proto3,oneof"`
}

type KafkaMessage_TimeToCut struct {
	TimeToCut *KafkaMessageTimeToCut `protobuf:"bytes,2,opt,name=time_to_cut,json=timeToCut,proto3,oneof"`
}

type KafkaMessage_Connect struct {
	Connect *KafkaMessageConnect `protobuf:"bytes,3,opt,name=connect,proto3,oneof"`
}

func (*KafkaMessage_Regular) isKafkaMessage_Type() {}

func (*KafkaMessage_TimeToCut) isKafkaMessage_Type() {}

func (*KafkaMessage_Connect) isKafkaMessage_Type() {}

// KafkaMessageRegular wraps a marshalled envelope.
type KafkaMessageRegular struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload        []byte                    `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	ConfigSeq      uint64                    `protobuf:"varint,2,opt,name=config_seq,json=configSeq,proto3" json:"config_seq,omitempty"`
	Class          KafkaMessageRegular_Class `protobuf:"varint,3,opt,name=class,proto3,enum=orderer.KafkaMessageRegular_Class" json:"class,omitempty"`
	OriginalOffset int64                     `protobuf:"varint,4,opt,name=original_offset,json=originalOffset,proto3" json:"original_offset,omitempty"`
}

func (x *KafkaMessageRegular) Reset() {
	*x = KafkaMessageRegular{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_kafka_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaMessageRegular) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaMessageRegular) ProtoMessage() {}

func (x *KafkaMessageRegular) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_kafka_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaMessageRegular.ProtoReflect.Descriptor instead.
func (*KafkaMessageRegular) Descriptor() ([]byte, []int) {
	return file_orderer_kafka_proto_rawDescGZIP(), []int{1}
}

func (x *KafkaMessageRegular) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *KafkaMessageRegular) GetConfigSeq() uint64 {
	if x != nil {
		return x.ConfigSeq
	}
	return 0
}

func (x *KafkaMessageRegular) GetClass() KafkaMessageRegular_Class {
	if x != nil {
		return x.Class
	}
	return KafkaMessageRegular_UNKNOWN
}

func (x *KafkaMessageRegular) GetOriginalOffset() int64 {
	if x != nil {
		return x.OriginalOffset
	}
	return 0
}

// KafkaMessageTimeToCut is used to signal to the orderers
// that it is time to cut block <block_number>.
type KafkaMessageTimeToCut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=block_number,json=blockNumber,proto3" json:"block_number,omitempty"`
}

func (x *KafkaMessageTimeToCut) Reset() {
	*x = KafkaMessageTimeToCut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_kafka_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaMessageTimeToCut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaMessageTimeToCut) ProtoMessage() {}

func (x *KafkaMessageTimeToCut) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_kafka_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaMessageTimeToCut.ProtoReflect.Descriptor instead.
func (*KafkaMessageTimeToCut) Descriptor() ([]byte, []int) {
	return file_orderer_kafka_proto_rawDescGZIP(), []int{2}
}

func (x *KafkaMessageTimeToCut) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

// KafkaMessageConnect is posted by an orderer upon booting up.
// It is used to prevent the panic that would be caused if we
// were to consume an empty partition. It is ignored by all
// orderers when processing the partition.
type KafkaMessageConnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload []byte `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *KafkaMessageConnect) Reset() {
	*x = KafkaMessageConnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_kafka_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaMessageConnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaMessageConnect) ProtoMessage() {}

func (x *KafkaMessageConnect) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_kafka_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaMessageConnect.ProtoReflect.Descriptor instead.
func (*KafkaMessageConnect) Descriptor() ([]byte, []int) {
	return file_orderer_kafka_proto_rawDescGZIP(), []int{3}
}

func (x *KafkaMessageConnect) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

// KafkaMetadata is encoded into the ORDERER block to keep track of
// Kafka-related metadata associated with this block.
type KafkaMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LastOffsetPersisted is the encoded value for the Metadata message
	// which is encoded in the ORDERER block metadata index for the case
	// of the Kafka-based orderer.
	LastOffsetPersisted int64 `protobuf:"varint,1,opt,name=last_offset_persisted,json=lastOffsetPersisted,proto3" json:"last_offset_persisted,omitempty"`
	// LastOriginalOffsetProcessed is used to keep track of the newest
	// offset processed if a message is re-validated and re-ordered.
	// This value is used to deduplicate re-submitted messages from
	// multiple orderer so that we don't bother re-processing it again.
	LastOriginalOffsetProcessed int64 `protobuf:"varint,2,opt,name=last_original_offset_processed,json=lastOriginalOffsetProcessed,proto3" json:"last_original_offset_processed,omitempty"`
	// LastResubmittedConfigOffset is used to capture the newest offset of
	// CONFIG kafka message, which is revalidated and resubmitted. By comparing
	// this with LastOriginalOffsetProcessed, we could detemine whether there
	// are still CONFIG messages that have been resubmitted but NOT processed
	// yet. It's used as condition to block ingress messages, so we could reduce
	// the overhead of repeatedly resubmitting messages as config seq keeps
	// advancing.
	LastResubmittedConfigOffset int64 `protobuf:"varint,3,opt,name=last_resubmitted_config_offset,json=lastResubmittedConfigOffset,proto3" json:"last_resubmitted_config_offset,omitempty"`
}

func (x *KafkaMetadata) Reset() {
	*x = KafkaMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orderer_kafka_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaMetadata) ProtoMessage() {}

func (x *KafkaMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_orderer_kafka_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaMetadata.ProtoReflect.Descriptor instead.
func (*KafkaMetadata) Descriptor() ([]byte, []int) {
	return file_orderer_kafka_proto_rawDescGZIP(), []int{4}
}

func (x *KafkaMetadata) GetLastOffsetPersisted() int64 {
	if x != nil {
		return x.LastOffsetPersisted
	}
	return 0
}

func (x *KafkaMetadata) GetLastOriginalOffsetProcessed() int64 {
	if x != nil {
		return x.LastOriginalOffsetProcessed
	}
	return 0
}

func (x *KafkaMetadata) GetLastResubmittedConfigOffset() int64 {
	if x != nil {
		return x.LastResubmittedConfigOffset
	}
	return 0
}

var File_orderer_kafka_proto protoreflect.FileDescriptor

var file_orderer_kafka_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2f, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x22, 0xcc,
	0x01, 0x0a, 0x0c, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x38, 0x0a, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x48, 0x00,
	0x52, 0x07, 0x72, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x40, 0x0a, 0x0b, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x43, 0x75, 0x74, 0x48, 0x00,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x54, 0x6f, 0x43, 0x75, 0x74, 0x12, 0x38, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0xdf, 0x01,
	0x0a, 0x13, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x67, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x71, 0x12, 0x38,
	0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x2e, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x22, 0x2c, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41,
	0x4c, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4f, 0x4e, 0x46, 0x49, 0x47, 0x10, 0x02, 0x22,
	0x3a, 0x0a, 0x15, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x54, 0x6f, 0x43, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x13, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xcd, 0x01, 0x0a,
	0x0d, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x32,
	0x0a, 0x15, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x65,
	0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6c,
	0x61, 0x73, 0x74, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x64, 0x12, 0x43, 0x0a, 0x1e, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x61, 0x6c, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1b, 0x6c, 0x61, 0x73, 0x74,
	0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x43, 0x0a, 0x1e, 0x6c, 0x61, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x1b, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x42, 0xa6, 0x01, 0x0a,
	0x25, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72,
	0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x42, 0x0a, 0x4b, 0x61, 0x66, 0x6b, 0x61, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62,
	0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x4f, 0x58,
	0x58, 0xaa, 0x02, 0x07, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0xca, 0x02, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x72, 0xe2, 0x02, 0x13, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderer_kafka_proto_rawDescOnce sync.Once
	file_orderer_kafka_proto_rawDescData = file_orderer_kafka_proto_rawDesc
)

func file_orderer_kafka_proto_rawDescGZIP() []byte {
	file_orderer_kafka_proto_rawDescOnce.Do(func() {
		file_orderer_kafka_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderer_kafka_proto_rawDescData)
	})
	return file_orderer_kafka_proto_rawDescData
}

var file_orderer_kafka_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orderer_kafka_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_orderer_kafka_proto_goTypes = []interface{}{
	(KafkaMessageRegular_Class)(0), // 0: orderer.KafkaMessageRegular.Class
	(*KafkaMessage)(nil),           // 1: orderer.KafkaMessage
	(*KafkaMessageRegular)(nil),    // 2: orderer.KafkaMessageRegular
	(*KafkaMessageTimeToCut)(nil),  // 3: orderer.KafkaMessageTimeToCut
	(*KafkaMessageConnect)(nil),    // 4: orderer.KafkaMessageConnect
	(*KafkaMetadata)(nil),          // 5: orderer.KafkaMetadata
}
var file_orderer_kafka_proto_depIdxs = []int32{
	2, // 0: orderer.KafkaMessage.regular:type_name -> orderer.KafkaMessageRegular
	3, // 1: orderer.KafkaMessage.time_to_cut:type_name -> orderer.KafkaMessageTimeToCut
	4, // 2: orderer.KafkaMessage.connect:type_name -> orderer.KafkaMessageConnect
	0, // 3: orderer.KafkaMessageRegular.class:type_name -> orderer.KafkaMessageRegular.Class
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_orderer_kafka_proto_init() }
func file_orderer_kafka_proto_init() {
	if File_orderer_kafka_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orderer_kafka_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaMessage); i {
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
		file_orderer_kafka_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaMessageRegular); i {
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
		file_orderer_kafka_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaMessageTimeToCut); i {
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
		file_orderer_kafka_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaMessageConnect); i {
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
		file_orderer_kafka_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaMetadata); i {
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
	file_orderer_kafka_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*KafkaMessage_Regular)(nil),
		(*KafkaMessage_TimeToCut)(nil),
		(*KafkaMessage_Connect)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orderer_kafka_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderer_kafka_proto_goTypes,
		DependencyIndexes: file_orderer_kafka_proto_depIdxs,
		EnumInfos:         file_orderer_kafka_proto_enumTypes,
		MessageInfos:      file_orderer_kafka_proto_msgTypes,
	}.Build()
	File_orderer_kafka_proto = out.File
	file_orderer_kafka_proto_rawDesc = nil
	file_orderer_kafka_proto_goTypes = nil
	file_orderer_kafka_proto_depIdxs = nil
}
