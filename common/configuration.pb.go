// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: common/configuration.proto

package common

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

// HashingAlgorithm is encoded into the configuration transaction as a
// configuration item of type Chain with a Key of "HashingAlgorithm" and a
// Value of HashingAlgorithm as marshaled protobuf bytes
type HashingAlgorithm struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SHA256 is currently the only supported and tested algorithm.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *HashingAlgorithm) Reset() {
	*x = HashingAlgorithm{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashingAlgorithm) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashingAlgorithm) ProtoMessage() {}

func (x *HashingAlgorithm) ProtoReflect() protoreflect.Message {
	mi := &file_common_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashingAlgorithm.ProtoReflect.Descriptor instead.
func (*HashingAlgorithm) Descriptor() ([]byte, []int) {
	return file_common_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *HashingAlgorithm) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// BlockDataHashingStructure is encoded into the configuration transaction as a configuration item of
// type Chain with a Key of "BlockDataHashingStructure" and a Value of HashingAlgorithm as marshaled protobuf bytes
type BlockDataHashingStructure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// width specifies the width of the Merkle tree to use when computing the BlockDataHash
	// in order to replicate flat hashing, set this width to MAX_UINT32
	Width uint32 `protobuf:"varint,1,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *BlockDataHashingStructure) Reset() {
	*x = BlockDataHashingStructure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockDataHashingStructure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockDataHashingStructure) ProtoMessage() {}

func (x *BlockDataHashingStructure) ProtoReflect() protoreflect.Message {
	mi := &file_common_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockDataHashingStructure.ProtoReflect.Descriptor instead.
func (*BlockDataHashingStructure) Descriptor() ([]byte, []int) {
	return file_common_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *BlockDataHashingStructure) GetWidth() uint32 {
	if x != nil {
		return x.Width
	}
	return 0
}

// OrdererAddresses is encoded into the configuration transaction as a configuration item of type Chain
// with a Key of "OrdererAddresses" and a Value of OrdererAddresses as marshaled protobuf bytes
type OrdererAddresses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (x *OrdererAddresses) Reset() {
	*x = OrdererAddresses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdererAddresses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdererAddresses) ProtoMessage() {}

func (x *OrdererAddresses) ProtoReflect() protoreflect.Message {
	mi := &file_common_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdererAddresses.ProtoReflect.Descriptor instead.
func (*OrdererAddresses) Descriptor() ([]byte, []int) {
	return file_common_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *OrdererAddresses) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

// Consortium represents the consortium context in which the channel was created
type Consortium struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Consortium) Reset() {
	*x = Consortium{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_configuration_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consortium) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consortium) ProtoMessage() {}

func (x *Consortium) ProtoReflect() protoreflect.Message {
	mi := &file_common_configuration_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consortium.ProtoReflect.Descriptor instead.
func (*Consortium) Descriptor() ([]byte, []int) {
	return file_common_configuration_proto_rawDescGZIP(), []int{3}
}

func (x *Consortium) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Capabilities message defines the capabilities a particular binary must implement
// for that binary to be able to safely participate in the channel.  The capabilities
// message is defined at the /Channel level, the /Channel/Application level, and the
// /Channel/Orderer level.
//
// The /Channel level capabilties define capabilities which both the orderer and peer
// binaries must satisfy.  These capabilties might be things like a new MSP type,
// or a new policy type.
//
// The /Channel/Orderer level capabilties define capabilities which must be supported
// by the orderer, but which have no bearing on the behavior of the peer.  For instance
// if the orderer changes the logic for how it constructs new channels, only all orderers
// must agree on the new logic.  The peers do not need to be aware of this change as
// they only interact with the channel after it has been constructed.
//
// Finally, the /Channel/Application level capabilities define capabilities which the peer
// binary must satisfy, but which have no bearing on the orderer.  For instance, if the
// peer adds a new UTXO transaction type, or changes the chaincode lifecycle requirements,
// all peers must agree on the new logic.  However, orderers never inspect transactions
// this deeply, and therefore have no need to be aware of the change.
//
// The capabilities strings defined in these messages typically correspond to release
// binary versions (e.g. "V1.1"), and are used primarilly as a mechanism for a fully
// upgraded network to switch from one set of logic to a new one.
//
// Although for V1.1, the orderers must be upgraded to V1.1 prior to the rest of the
// network, going forward, because of the split between the /Channel, /Channel/Orderer
// and /Channel/Application capabilities.  It should be possible for the orderer and
// application networks to upgrade themselves independently (with the exception of any
// new capabilities defined at the /Channel level).
type Capabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capabilities map[string]*Capability `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Capabilities) Reset() {
	*x = Capabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_configuration_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capabilities) ProtoMessage() {}

func (x *Capabilities) ProtoReflect() protoreflect.Message {
	mi := &file_common_configuration_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capabilities.ProtoReflect.Descriptor instead.
func (*Capabilities) Descriptor() ([]byte, []int) {
	return file_common_configuration_proto_rawDescGZIP(), []int{4}
}

func (x *Capabilities) GetCapabilities() map[string]*Capability {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

// Capability is an empty message for the time being.  It is defined as a protobuf
// message rather than a constant, so that we may extend capabilities with other fields
// if the need arises in the future.  For the time being, a capability being in the
// capabilities map requires that that capability be supported.
type Capability struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Capability) Reset() {
	*x = Capability{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_configuration_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Capability) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Capability) ProtoMessage() {}

func (x *Capability) ProtoReflect() protoreflect.Message {
	mi := &file_common_configuration_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Capability.ProtoReflect.Descriptor instead.
func (*Capability) Descriptor() ([]byte, []int) {
	return file_common_configuration_proto_rawDescGZIP(), []int{5}
}

var File_common_configuration_proto protoreflect.FileDescriptor

var file_common_configuration_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x22, 0x26, 0x0a, 0x10, 0x48, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x31, 0x0a, 0x19,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x48, 0x61, 0x73, 0x68, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x69, 0x64,
	0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x22,
	0x30, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65,
	0x73, 0x22, 0x20, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x72, 0x74, 0x69, 0x75, 0x6d, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xaf, 0x01, 0x0a, 0x0c, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x1a, 0x53, 0x0a, 0x11, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x0c, 0x0a, 0x0a, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x42, 0xa8, 0x01, 0x0a, 0x24, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65,
	0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x12, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x61, 0x62, 0x72, 0x69,
	0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58, 0xaa, 0x02,
	0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0xca, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0xe2, 0x02, 0x12, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_configuration_proto_rawDescOnce sync.Once
	file_common_configuration_proto_rawDescData = file_common_configuration_proto_rawDesc
)

func file_common_configuration_proto_rawDescGZIP() []byte {
	file_common_configuration_proto_rawDescOnce.Do(func() {
		file_common_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_configuration_proto_rawDescData)
	})
	return file_common_configuration_proto_rawDescData
}

var file_common_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_common_configuration_proto_goTypes = []interface{}{
	(*HashingAlgorithm)(nil),          // 0: common.HashingAlgorithm
	(*BlockDataHashingStructure)(nil), // 1: common.BlockDataHashingStructure
	(*OrdererAddresses)(nil),          // 2: common.OrdererAddresses
	(*Consortium)(nil),                // 3: common.Consortium
	(*Capabilities)(nil),              // 4: common.Capabilities
	(*Capability)(nil),                // 5: common.Capability
	nil,                               // 6: common.Capabilities.CapabilitiesEntry
}
var file_common_configuration_proto_depIdxs = []int32{
	6, // 0: common.Capabilities.capabilities:type_name -> common.Capabilities.CapabilitiesEntry
	5, // 1: common.Capabilities.CapabilitiesEntry.value:type_name -> common.Capability
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_common_configuration_proto_init() }
func file_common_configuration_proto_init() {
	if File_common_configuration_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_configuration_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashingAlgorithm); i {
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
		file_common_configuration_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockDataHashingStructure); i {
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
		file_common_configuration_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdererAddresses); i {
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
		file_common_configuration_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consortium); i {
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
		file_common_configuration_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capabilities); i {
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
		file_common_configuration_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Capability); i {
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
			RawDescriptor: file_common_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_configuration_proto_goTypes,
		DependencyIndexes: file_common_configuration_proto_depIdxs,
		MessageInfos:      file_common_configuration_proto_msgTypes,
	}.Build()
	File_common_configuration_proto = out.File
	file_common_configuration_proto_rawDesc = nil
	file_common_configuration_proto_goTypes = nil
	file_common_configuration_proto_depIdxs = nil
}
