// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: core/contract/exchange_contract.proto

package pb

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

type ExchangeCreateContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress       []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	FirstTokenId       []byte `protobuf:"bytes,2,opt,name=first_token_id,json=firstTokenId,proto3" json:"first_token_id,omitempty"`
	FirstTokenBalance  int64  `protobuf:"varint,3,opt,name=first_token_balance,json=firstTokenBalance,proto3" json:"first_token_balance,omitempty"`
	SecondTokenId      []byte `protobuf:"bytes,4,opt,name=second_token_id,json=secondTokenId,proto3" json:"second_token_id,omitempty"`
	SecondTokenBalance int64  `protobuf:"varint,5,opt,name=second_token_balance,json=secondTokenBalance,proto3" json:"second_token_balance,omitempty"`
}

func (x *ExchangeCreateContract) Reset() {
	*x = ExchangeCreateContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_exchange_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeCreateContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeCreateContract) ProtoMessage() {}

func (x *ExchangeCreateContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_exchange_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeCreateContract.ProtoReflect.Descriptor instead.
func (*ExchangeCreateContract) Descriptor() ([]byte, []int) {
	return file_core_contract_exchange_contract_proto_rawDescGZIP(), []int{0}
}

func (x *ExchangeCreateContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *ExchangeCreateContract) GetFirstTokenId() []byte {
	if x != nil {
		return x.FirstTokenId
	}
	return nil
}

func (x *ExchangeCreateContract) GetFirstTokenBalance() int64 {
	if x != nil {
		return x.FirstTokenBalance
	}
	return 0
}

func (x *ExchangeCreateContract) GetSecondTokenId() []byte {
	if x != nil {
		return x.SecondTokenId
	}
	return nil
}

func (x *ExchangeCreateContract) GetSecondTokenBalance() int64 {
	if x != nil {
		return x.SecondTokenBalance
	}
	return 0
}

type ExchangeInjectContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ExchangeId   int64  `protobuf:"varint,2,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
	TokenId      []byte `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Quant        int64  `protobuf:"varint,4,opt,name=quant,proto3" json:"quant,omitempty"`
}

func (x *ExchangeInjectContract) Reset() {
	*x = ExchangeInjectContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_exchange_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeInjectContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeInjectContract) ProtoMessage() {}

func (x *ExchangeInjectContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_exchange_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeInjectContract.ProtoReflect.Descriptor instead.
func (*ExchangeInjectContract) Descriptor() ([]byte, []int) {
	return file_core_contract_exchange_contract_proto_rawDescGZIP(), []int{1}
}

func (x *ExchangeInjectContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *ExchangeInjectContract) GetExchangeId() int64 {
	if x != nil {
		return x.ExchangeId
	}
	return 0
}

func (x *ExchangeInjectContract) GetTokenId() []byte {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *ExchangeInjectContract) GetQuant() int64 {
	if x != nil {
		return x.Quant
	}
	return 0
}

type ExchangeWithdrawContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ExchangeId   int64  `protobuf:"varint,2,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
	TokenId      []byte `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Quant        int64  `protobuf:"varint,4,opt,name=quant,proto3" json:"quant,omitempty"`
}

func (x *ExchangeWithdrawContract) Reset() {
	*x = ExchangeWithdrawContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_exchange_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeWithdrawContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeWithdrawContract) ProtoMessage() {}

func (x *ExchangeWithdrawContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_exchange_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeWithdrawContract.ProtoReflect.Descriptor instead.
func (*ExchangeWithdrawContract) Descriptor() ([]byte, []int) {
	return file_core_contract_exchange_contract_proto_rawDescGZIP(), []int{2}
}

func (x *ExchangeWithdrawContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *ExchangeWithdrawContract) GetExchangeId() int64 {
	if x != nil {
		return x.ExchangeId
	}
	return 0
}

func (x *ExchangeWithdrawContract) GetTokenId() []byte {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *ExchangeWithdrawContract) GetQuant() int64 {
	if x != nil {
		return x.Quant
	}
	return 0
}

type ExchangeTransactionContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ExchangeId   int64  `protobuf:"varint,2,opt,name=exchange_id,json=exchangeId,proto3" json:"exchange_id,omitempty"`
	TokenId      []byte `protobuf:"bytes,3,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Quant        int64  `protobuf:"varint,4,opt,name=quant,proto3" json:"quant,omitempty"`
	Expected     int64  `protobuf:"varint,5,opt,name=expected,proto3" json:"expected,omitempty"`
}

func (x *ExchangeTransactionContract) Reset() {
	*x = ExchangeTransactionContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_contract_exchange_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExchangeTransactionContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExchangeTransactionContract) ProtoMessage() {}

func (x *ExchangeTransactionContract) ProtoReflect() protoreflect.Message {
	mi := &file_core_contract_exchange_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExchangeTransactionContract.ProtoReflect.Descriptor instead.
func (*ExchangeTransactionContract) Descriptor() ([]byte, []int) {
	return file_core_contract_exchange_contract_proto_rawDescGZIP(), []int{3}
}

func (x *ExchangeTransactionContract) GetOwnerAddress() []byte {
	if x != nil {
		return x.OwnerAddress
	}
	return nil
}

func (x *ExchangeTransactionContract) GetExchangeId() int64 {
	if x != nil {
		return x.ExchangeId
	}
	return 0
}

func (x *ExchangeTransactionContract) GetTokenId() []byte {
	if x != nil {
		return x.TokenId
	}
	return nil
}

func (x *ExchangeTransactionContract) GetQuant() int64 {
	if x != nil {
		return x.Quant
	}
	return 0
}

func (x *ExchangeTransactionContract) GetExpected() int64 {
	if x != nil {
		return x.Expected
	}
	return 0
}

var File_core_contract_exchange_contract_proto protoreflect.FileDescriptor

var file_core_contract_exchange_contract_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2f,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x22, 0xed, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x24, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0d, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12,
	0x30, 0x0a, 0x14, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x22, 0x8f, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e,
	0x6a, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x18, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x22, 0xb0, 0x01, 0x0a, 0x1b, 0x45, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a,
	0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x42, 0x41, 0x0a, 0x18, 0x6f, 0x72,
	0x67, 0x2e, 0x74, 0x72, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x68, 0x61, 0x6e, 0x67, 0x30, 0x31, 0x32, 0x35, 0x2f, 0x67, 0x6f, 0x2d,
	0x74, 0x72, 0x6f, 0x6e, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_contract_exchange_contract_proto_rawDescOnce sync.Once
	file_core_contract_exchange_contract_proto_rawDescData = file_core_contract_exchange_contract_proto_rawDesc
)

func file_core_contract_exchange_contract_proto_rawDescGZIP() []byte {
	file_core_contract_exchange_contract_proto_rawDescOnce.Do(func() {
		file_core_contract_exchange_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_contract_exchange_contract_proto_rawDescData)
	})
	return file_core_contract_exchange_contract_proto_rawDescData
}

var file_core_contract_exchange_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_core_contract_exchange_contract_proto_goTypes = []interface{}{
	(*ExchangeCreateContract)(nil),      // 0: protocol.ExchangeCreateContract
	(*ExchangeInjectContract)(nil),      // 1: protocol.ExchangeInjectContract
	(*ExchangeWithdrawContract)(nil),    // 2: protocol.ExchangeWithdrawContract
	(*ExchangeTransactionContract)(nil), // 3: protocol.ExchangeTransactionContract
}
var file_core_contract_exchange_contract_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_core_contract_exchange_contract_proto_init() }
func file_core_contract_exchange_contract_proto_init() {
	if File_core_contract_exchange_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_contract_exchange_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeCreateContract); i {
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
		file_core_contract_exchange_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeInjectContract); i {
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
		file_core_contract_exchange_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeWithdrawContract); i {
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
		file_core_contract_exchange_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExchangeTransactionContract); i {
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
			RawDescriptor: file_core_contract_exchange_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_contract_exchange_contract_proto_goTypes,
		DependencyIndexes: file_core_contract_exchange_contract_proto_depIdxs,
		MessageInfos:      file_core_contract_exchange_contract_proto_msgTypes,
	}.Build()
	File_core_contract_exchange_contract_proto = out.File
	file_core_contract_exchange_contract_proto_rawDesc = nil
	file_core_contract_exchange_contract_proto_goTypes = nil
	file_core_contract_exchange_contract_proto_depIdxs = nil
}
