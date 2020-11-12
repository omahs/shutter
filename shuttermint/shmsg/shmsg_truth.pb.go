// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: shmsg/shmsg_truth.proto

package shmsg

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type G1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	G1Bytes []byte `protobuf:"bytes,1,opt,name=g1bytes,proto3" json:"g1bytes,omitempty"` // Unmarshal with new(G1).Unmarshal(msg.g1bytes)
}

func (x *G1) Reset() {
	*x = G1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_shmsg_truth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *G1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*G1) ProtoMessage() {}

func (x *G1) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_shmsg_truth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use G1.ProtoReflect.Descriptor instead.
func (*G1) Descriptor() ([]byte, []int) {
	return file_shmsg_shmsg_truth_proto_rawDescGZIP(), []int{0}
}

func (x *G1) GetG1Bytes() []byte {
	if x != nil {
		return x.G1Bytes
	}
	return nil
}

type G2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	G2Bytes []byte `protobuf:"bytes,1,opt,name=g2bytes,proto3" json:"g2bytes,omitempty"` // Unmarshal with new(G2).Unmarshal(msg.g2bytes)
}

func (x *G2) Reset() {
	*x = G2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_shmsg_truth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *G2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*G2) ProtoMessage() {}

func (x *G2) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_shmsg_truth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use G2.ProtoReflect.Descriptor instead.
func (*G2) Descriptor() ([]byte, []int) {
	return file_shmsg_shmsg_truth_proto_rawDescGZIP(), []int{1}
}

func (x *G2) GetG2Bytes() []byte {
	if x != nil {
		return x.G2Bytes
	}
	return nil
}

type GT struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gtbytes []byte `protobuf:"bytes,1,opt,name=gtbytes,proto3" json:"gtbytes,omitempty"` // Unmarshal with new(GT).Unmarshal(msg.g2bytes)
}

func (x *GT) Reset() {
	*x = GT{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_shmsg_truth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GT) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GT) ProtoMessage() {}

func (x *GT) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_shmsg_truth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GT.ProtoReflect.Descriptor instead.
func (*GT) Descriptor() ([]byte, []int) {
	return file_shmsg_shmsg_truth_proto_rawDescGZIP(), []int{2}
}

func (x *GT) GetGtbytes() []byte {
	if x != nil {
		return x.Gtbytes
	}
	return nil
}

type PolyEvalMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eon           uint64 `protobuf:"varint,1,opt,name=eon,proto3" json:"eon,omitempty"`
	Receiver      []byte `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	EncryptedEval []byte `protobuf:"bytes,3,opt,name=encryptedEval,proto3" json:"encryptedEval,omitempty"`
}

func (x *PolyEvalMsg) Reset() {
	*x = PolyEvalMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_shmsg_truth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolyEvalMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolyEvalMsg) ProtoMessage() {}

func (x *PolyEvalMsg) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_shmsg_truth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolyEvalMsg.ProtoReflect.Descriptor instead.
func (*PolyEvalMsg) Descriptor() ([]byte, []int) {
	return file_shmsg_shmsg_truth_proto_rawDescGZIP(), []int{3}
}

func (x *PolyEvalMsg) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *PolyEvalMsg) GetReceiver() []byte {
	if x != nil {
		return x.Receiver
	}
	return nil
}

func (x *PolyEvalMsg) GetEncryptedEval() []byte {
	if x != nil {
		return x.EncryptedEval
	}
	return nil
}

type PolyCommitmentMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eon    uint64   `protobuf:"varint,1,opt,name=eon,proto3" json:"eon,omitempty"`
	Gammas [][]byte `protobuf:"bytes,2,rep,name=gammas,proto3" json:"gammas,omitempty"`
}

func (x *PolyCommitmentMsg) Reset() {
	*x = PolyCommitmentMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_shmsg_truth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolyCommitmentMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolyCommitmentMsg) ProtoMessage() {}

func (x *PolyCommitmentMsg) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_shmsg_truth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolyCommitmentMsg.ProtoReflect.Descriptor instead.
func (*PolyCommitmentMsg) Descriptor() ([]byte, []int) {
	return file_shmsg_shmsg_truth_proto_rawDescGZIP(), []int{4}
}

func (x *PolyCommitmentMsg) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *PolyCommitmentMsg) GetGammas() [][]byte {
	if x != nil {
		return x.Gammas
	}
	return nil
}

type AccusationMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eon     uint64 `protobuf:"varint,1,opt,name=eon,proto3" json:"eon,omitempty"`
	Accused []byte `protobuf:"bytes,2,opt,name=accused,proto3" json:"accused,omitempty"`
}

func (x *AccusationMsg) Reset() {
	*x = AccusationMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_shmsg_truth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccusationMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccusationMsg) ProtoMessage() {}

func (x *AccusationMsg) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_shmsg_truth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccusationMsg.ProtoReflect.Descriptor instead.
func (*AccusationMsg) Descriptor() ([]byte, []int) {
	return file_shmsg_shmsg_truth_proto_rawDescGZIP(), []int{5}
}

func (x *AccusationMsg) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *AccusationMsg) GetAccused() []byte {
	if x != nil {
		return x.Accused
	}
	return nil
}

type ApologyMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eon      uint64 `protobuf:"varint,1,opt,name=eon,proto3" json:"eon,omitempty"`
	Accuser  []byte `protobuf:"bytes,2,opt,name=accuser,proto3" json:"accuser,omitempty"`
	PolyEval []byte `protobuf:"bytes,3,opt,name=polyEval,proto3" json:"polyEval,omitempty"`
}

func (x *ApologyMsg) Reset() {
	*x = ApologyMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_shmsg_truth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApologyMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApologyMsg) ProtoMessage() {}

func (x *ApologyMsg) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_shmsg_truth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApologyMsg.ProtoReflect.Descriptor instead.
func (*ApologyMsg) Descriptor() ([]byte, []int) {
	return file_shmsg_shmsg_truth_proto_rawDescGZIP(), []int{6}
}

func (x *ApologyMsg) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *ApologyMsg) GetAccuser() []byte {
	if x != nil {
		return x.Accuser
	}
	return nil
}

func (x *ApologyMsg) GetPolyEval() []byte {
	if x != nil {
		return x.PolyEval
	}
	return nil
}

type EpochSKShareMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eon          uint64 `protobuf:"varint,1,opt,name=eon,proto3" json:"eon,omitempty"`
	Epoch        uint64 `protobuf:"varint,2,opt,name=epoch,proto3" json:"epoch,omitempty"`
	EpochSKShare []byte `protobuf:"bytes,3,opt,name=epochSKShare,proto3" json:"epochSKShare,omitempty"`
}

func (x *EpochSKShareMsg) Reset() {
	*x = EpochSKShareMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shmsg_shmsg_truth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpochSKShareMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpochSKShareMsg) ProtoMessage() {}

func (x *EpochSKShareMsg) ProtoReflect() protoreflect.Message {
	mi := &file_shmsg_shmsg_truth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpochSKShareMsg.ProtoReflect.Descriptor instead.
func (*EpochSKShareMsg) Descriptor() ([]byte, []int) {
	return file_shmsg_shmsg_truth_proto_rawDescGZIP(), []int{7}
}

func (x *EpochSKShareMsg) GetEon() uint64 {
	if x != nil {
		return x.Eon
	}
	return 0
}

func (x *EpochSKShareMsg) GetEpoch() uint64 {
	if x != nil {
		return x.Epoch
	}
	return 0
}

func (x *EpochSKShareMsg) GetEpochSKShare() []byte {
	if x != nil {
		return x.EpochSKShare
	}
	return nil
}

var File_shmsg_shmsg_truth_proto protoreflect.FileDescriptor

var file_shmsg_shmsg_truth_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x68, 0x6d, 0x73, 0x67, 0x2f, 0x73, 0x68, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x72,
	0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x68, 0x6d, 0x73, 0x67,
	0x22, 0x1e, 0x0a, 0x02, 0x47, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x31, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x31, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x22, 0x1e, 0x0a, 0x02, 0x47, 0x32, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x32, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x32, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x22, 0x1e, 0x0a, 0x02, 0x47, 0x54, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x74, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x67, 0x74, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x22, 0x61, 0x0a, 0x0b, 0x50, 0x6f, 0x6c, 0x79, 0x45, 0x76, 0x61, 0x6c, 0x4d, 0x73, 0x67, 0x12,
	0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x12, 0x24, 0x0a,
	0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x45, 0x76, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x45,
	0x76, 0x61, 0x6c, 0x22, 0x3d, 0x0a, 0x11, 0x50, 0x6f, 0x6c, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x61,
	0x6d, 0x6d, 0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x6d,
	0x61, 0x73, 0x22, 0x3b, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x75, 0x73, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x65, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x75, 0x73, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x63, 0x63, 0x75, 0x73, 0x65, 0x64, 0x22,
	0x54, 0x0a, 0x0a, 0x41, 0x70, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x61, 0x63, 0x63, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x6c,
	0x79, 0x45, 0x76, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x6f, 0x6c,
	0x79, 0x45, 0x76, 0x61, 0x6c, 0x22, 0x5d, 0x0a, 0x0f, 0x45, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x4b,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x65, 0x70, 0x6f, 0x63, 0x68,
	0x12, 0x22, 0x0a, 0x0c, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x4b, 0x53, 0x68, 0x61, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x53, 0x4b, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x73, 0x68, 0x6d, 0x73, 0x67, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shmsg_shmsg_truth_proto_rawDescOnce sync.Once
	file_shmsg_shmsg_truth_proto_rawDescData = file_shmsg_shmsg_truth_proto_rawDesc
)

func file_shmsg_shmsg_truth_proto_rawDescGZIP() []byte {
	file_shmsg_shmsg_truth_proto_rawDescOnce.Do(func() {
		file_shmsg_shmsg_truth_proto_rawDescData = protoimpl.X.CompressGZIP(file_shmsg_shmsg_truth_proto_rawDescData)
	})
	return file_shmsg_shmsg_truth_proto_rawDescData
}

var file_shmsg_shmsg_truth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_shmsg_shmsg_truth_proto_goTypes = []interface{}{
	(*G1)(nil),                // 0: shmsg.G1
	(*G2)(nil),                // 1: shmsg.G2
	(*GT)(nil),                // 2: shmsg.GT
	(*PolyEvalMsg)(nil),       // 3: shmsg.PolyEvalMsg
	(*PolyCommitmentMsg)(nil), // 4: shmsg.PolyCommitmentMsg
	(*AccusationMsg)(nil),     // 5: shmsg.AccusationMsg
	(*ApologyMsg)(nil),        // 6: shmsg.ApologyMsg
	(*EpochSKShareMsg)(nil),   // 7: shmsg.EpochSKShareMsg
}
var file_shmsg_shmsg_truth_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shmsg_shmsg_truth_proto_init() }
func file_shmsg_shmsg_truth_proto_init() {
	if File_shmsg_shmsg_truth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shmsg_shmsg_truth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*G1); i {
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
		file_shmsg_shmsg_truth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*G2); i {
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
		file_shmsg_shmsg_truth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GT); i {
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
		file_shmsg_shmsg_truth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolyEvalMsg); i {
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
		file_shmsg_shmsg_truth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolyCommitmentMsg); i {
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
		file_shmsg_shmsg_truth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccusationMsg); i {
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
		file_shmsg_shmsg_truth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApologyMsg); i {
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
		file_shmsg_shmsg_truth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpochSKShareMsg); i {
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
			RawDescriptor: file_shmsg_shmsg_truth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shmsg_shmsg_truth_proto_goTypes,
		DependencyIndexes: file_shmsg_shmsg_truth_proto_depIdxs,
		MessageInfos:      file_shmsg_shmsg_truth_proto_msgTypes,
	}.Build()
	File_shmsg_shmsg_truth_proto = out.File
	file_shmsg_shmsg_truth_proto_rawDesc = nil
	file_shmsg_shmsg_truth_proto_goTypes = nil
	file_shmsg_shmsg_truth_proto_depIdxs = nil
}
