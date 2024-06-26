// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: zklend/starknet/transform/v1/transforms.proto

package pbtransform

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

// Stream block headers only. The `transactions` field is always empty.
type BlockHeaderOnly struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlockHeaderOnly) Reset() {
	*x = BlockHeaderOnly{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zklend_starknet_transform_v1_transforms_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockHeaderOnly) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockHeaderOnly) ProtoMessage() {}

func (x *BlockHeaderOnly) ProtoReflect() protoreflect.Message {
	mi := &file_zklend_starknet_transform_v1_transforms_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockHeaderOnly.ProtoReflect.Descriptor instead.
func (*BlockHeaderOnly) Descriptor() ([]byte, []int) {
	return file_zklend_starknet_transform_v1_transforms_proto_rawDescGZIP(), []int{0}
}

// Stream every single block, but each block will only contain transactions that match with `event_filters`.
// A TransactionEventFilter message with an empty `event_filters` is invalid. Do not send any filter instead
// if you wish to receive full blocks.
type TransactionEventFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventFilters []*ContractEventFilter `protobuf:"bytes,1,rep,name=event_filters,json=eventFilters,proto3" json:"event_filters,omitempty"`
}

func (x *TransactionEventFilter) Reset() {
	*x = TransactionEventFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zklend_starknet_transform_v1_transforms_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionEventFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionEventFilter) ProtoMessage() {}

func (x *TransactionEventFilter) ProtoReflect() protoreflect.Message {
	mi := &file_zklend_starknet_transform_v1_transforms_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionEventFilter.ProtoReflect.Descriptor instead.
func (*TransactionEventFilter) Descriptor() ([]byte, []int) {
	return file_zklend_starknet_transform_v1_transforms_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionEventFilter) GetEventFilters() []*ContractEventFilter {
	if x != nil {
		return x.EventFilters
	}
	return nil
}

// Only include transactions which emit at least one event that *BOTH*
// * is emitted by `contract_address`
// * matches with at least one topic in `topics`
type ContractEventFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractAddress []byte             `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
	Topics          []*TopicWithRanges `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *ContractEventFilter) Reset() {
	*x = ContractEventFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zklend_starknet_transform_v1_transforms_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractEventFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractEventFilter) ProtoMessage() {}

func (x *ContractEventFilter) ProtoReflect() protoreflect.Message {
	mi := &file_zklend_starknet_transform_v1_transforms_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractEventFilter.ProtoReflect.Descriptor instead.
func (*ContractEventFilter) Descriptor() ([]byte, []int) {
	return file_zklend_starknet_transform_v1_transforms_proto_rawDescGZIP(), []int{2}
}

func (x *ContractEventFilter) GetContractAddress() []byte {
	if x != nil {
		return x.ContractAddress
	}
	return nil
}

func (x *ContractEventFilter) GetTopics() []*TopicWithRanges {
	if x != nil {
		return x.Topics
	}
	return nil
}

// Matches events whose `keys[0]` equals `topic`, *AND* in any of the `block_ranges`.
type TopicWithRanges struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic       []byte        `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	BlockRanges []*BlockRange `protobuf:"bytes,2,rep,name=block_ranges,json=blockRanges,proto3" json:"block_ranges,omitempty"`
}

func (x *TopicWithRanges) Reset() {
	*x = TopicWithRanges{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zklend_starknet_transform_v1_transforms_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicWithRanges) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicWithRanges) ProtoMessage() {}

func (x *TopicWithRanges) ProtoReflect() protoreflect.Message {
	mi := &file_zklend_starknet_transform_v1_transforms_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicWithRanges.ProtoReflect.Descriptor instead.
func (*TopicWithRanges) Descriptor() ([]byte, []int) {
	return file_zklend_starknet_transform_v1_transforms_proto_rawDescGZIP(), []int{3}
}

func (x *TopicWithRanges) GetTopic() []byte {
	if x != nil {
		return x.Topic
	}
	return nil
}

func (x *TopicWithRanges) GetBlockRanges() []*BlockRange {
	if x != nil {
		return x.BlockRanges
	}
	return nil
}

// A range of blocks. `start_block` is inclusive, and `end_block` is exclusive. When `end_block` is `0`, it means
// that any block height >= `start_block` is matched.
type BlockRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartBlock uint64 `protobuf:"varint,1,opt,name=start_block,json=startBlock,proto3" json:"start_block,omitempty"`
	EndBlock   uint64 `protobuf:"varint,2,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
}

func (x *BlockRange) Reset() {
	*x = BlockRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zklend_starknet_transform_v1_transforms_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRange) ProtoMessage() {}

func (x *BlockRange) ProtoReflect() protoreflect.Message {
	mi := &file_zklend_starknet_transform_v1_transforms_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRange.ProtoReflect.Descriptor instead.
func (*BlockRange) Descriptor() ([]byte, []int) {
	return file_zklend_starknet_transform_v1_transforms_proto_rawDescGZIP(), []int{4}
}

func (x *BlockRange) GetStartBlock() uint64 {
	if x != nil {
		return x.StartBlock
	}
	return 0
}

func (x *BlockRange) GetEndBlock() uint64 {
	if x != nil {
		return x.EndBlock
	}
	return 0
}

var File_zklend_starknet_transform_v1_transforms_proto protoreflect.FileDescriptor

var file_zklend_starknet_transform_v1_transforms_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x7a, 0x6b, 0x6c, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65,
	0x74, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1c, 0x7a, 0x6b, 0x6c, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x22, 0x11, 0x0a,
	0x0f, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4f, 0x6e, 0x6c, 0x79,
	0x22, 0x70, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x56, 0x0a, 0x0d, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x31, 0x2e, 0x7a, 0x6b, 0x6c, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b,
	0x6e, 0x65, 0x74, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x73, 0x22, 0x87, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x45, 0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x7a, 0x6b, 0x6c, 0x65, 0x6e, 0x64, 0x2e, 0x73,
	0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x57, 0x69, 0x74, 0x68, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x73, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x22, 0x74, 0x0a, 0x0f,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x57, 0x69, 0x74, 0x68, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x4b, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x7a, 0x6b,
	0x6c, 0x65, 0x6e, 0x64, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x22, 0x4a, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x64, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x42, 0x5a,
	0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x61,
	0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2d, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x66, 0x69, 0x72, 0x65,
	0x68, 0x6f, 0x73, 0x65, 0x2d, 0x73, 0x74, 0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x7a, 0x6b, 0x6c, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x74,
	0x61, 0x72, 0x6b, 0x6e, 0x65, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x70,
	0x62, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_zklend_starknet_transform_v1_transforms_proto_rawDescOnce sync.Once
	file_zklend_starknet_transform_v1_transforms_proto_rawDescData = file_zklend_starknet_transform_v1_transforms_proto_rawDesc
)

func file_zklend_starknet_transform_v1_transforms_proto_rawDescGZIP() []byte {
	file_zklend_starknet_transform_v1_transforms_proto_rawDescOnce.Do(func() {
		file_zklend_starknet_transform_v1_transforms_proto_rawDescData = protoimpl.X.CompressGZIP(file_zklend_starknet_transform_v1_transforms_proto_rawDescData)
	})
	return file_zklend_starknet_transform_v1_transforms_proto_rawDescData
}

var file_zklend_starknet_transform_v1_transforms_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_zklend_starknet_transform_v1_transforms_proto_goTypes = []interface{}{
	(*BlockHeaderOnly)(nil),        // 0: zklend.starknet.transform.v1.BlockHeaderOnly
	(*TransactionEventFilter)(nil), // 1: zklend.starknet.transform.v1.TransactionEventFilter
	(*ContractEventFilter)(nil),    // 2: zklend.starknet.transform.v1.ContractEventFilter
	(*TopicWithRanges)(nil),        // 3: zklend.starknet.transform.v1.TopicWithRanges
	(*BlockRange)(nil),             // 4: zklend.starknet.transform.v1.BlockRange
}
var file_zklend_starknet_transform_v1_transforms_proto_depIdxs = []int32{
	2, // 0: zklend.starknet.transform.v1.TransactionEventFilter.event_filters:type_name -> zklend.starknet.transform.v1.ContractEventFilter
	3, // 1: zklend.starknet.transform.v1.ContractEventFilter.topics:type_name -> zklend.starknet.transform.v1.TopicWithRanges
	4, // 2: zklend.starknet.transform.v1.TopicWithRanges.block_ranges:type_name -> zklend.starknet.transform.v1.BlockRange
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_zklend_starknet_transform_v1_transforms_proto_init() }
func file_zklend_starknet_transform_v1_transforms_proto_init() {
	if File_zklend_starknet_transform_v1_transforms_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zklend_starknet_transform_v1_transforms_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockHeaderOnly); i {
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
		file_zklend_starknet_transform_v1_transforms_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionEventFilter); i {
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
		file_zklend_starknet_transform_v1_transforms_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractEventFilter); i {
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
		file_zklend_starknet_transform_v1_transforms_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TopicWithRanges); i {
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
		file_zklend_starknet_transform_v1_transforms_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRange); i {
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
			RawDescriptor: file_zklend_starknet_transform_v1_transforms_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zklend_starknet_transform_v1_transforms_proto_goTypes,
		DependencyIndexes: file_zklend_starknet_transform_v1_transforms_proto_depIdxs,
		MessageInfos:      file_zklend_starknet_transform_v1_transforms_proto_msgTypes,
	}.Build()
	File_zklend_starknet_transform_v1_transforms_proto = out.File
	file_zklend_starknet_transform_v1_transforms_proto_rawDesc = nil
	file_zklend_starknet_transform_v1_transforms_proto_goTypes = nil
	file_zklend_starknet_transform_v1_transforms_proto_depIdxs = nil
}
