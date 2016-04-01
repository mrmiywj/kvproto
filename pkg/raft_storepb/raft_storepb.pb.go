// Code generated by protoc-gen-go.
// source: raft_storepb.proto
// DO NOT EDIT!

/*
Package raft_storepb is a generated protocol buffer package.

It is generated from these files:
	raft_storepb.proto

It has these top-level messages:
	RaftMessage
	RaftTruncatedState
	KeyValue
	RaftSnapshotData
	StoreIdent
*/
package raft_storepb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import raftpb "github.com/pingcap/kvproto/pkg/raftpb"
import metapb "github.com/pingcap/kvproto/pkg/metapb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type RaftMessage struct {
	RegionId         *uint64             `protobuf:"varint,1,opt,name=region_id" json:"region_id,omitempty"`
	FromPeer         *metapb.Peer        `protobuf:"bytes,2,opt,name=from_peer" json:"from_peer,omitempty"`
	ToPeer           *metapb.Peer        `protobuf:"bytes,3,opt,name=to_peer" json:"to_peer,omitempty"`
	Message          *raftpb.Message     `protobuf:"bytes,4,opt,name=message" json:"message,omitempty"`
	RegionEpoch      *metapb.RegionEpoch `protobuf:"bytes,5,opt,name=region_epoch" json:"region_epoch,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *RaftMessage) Reset()                    { *m = RaftMessage{} }
func (m *RaftMessage) String() string            { return proto.CompactTextString(m) }
func (*RaftMessage) ProtoMessage()               {}
func (*RaftMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RaftMessage) GetRegionId() uint64 {
	if m != nil && m.RegionId != nil {
		return *m.RegionId
	}
	return 0
}

func (m *RaftMessage) GetFromPeer() *metapb.Peer {
	if m != nil {
		return m.FromPeer
	}
	return nil
}

func (m *RaftMessage) GetToPeer() *metapb.Peer {
	if m != nil {
		return m.ToPeer
	}
	return nil
}

func (m *RaftMessage) GetMessage() *raftpb.Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *RaftMessage) GetRegionEpoch() *metapb.RegionEpoch {
	if m != nil {
		return m.RegionEpoch
	}
	return nil
}

type RaftTruncatedState struct {
	Index            *uint64 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	Term             *uint64 `protobuf:"varint,2,opt,name=term" json:"term,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RaftTruncatedState) Reset()                    { *m = RaftTruncatedState{} }
func (m *RaftTruncatedState) String() string            { return proto.CompactTextString(m) }
func (*RaftTruncatedState) ProtoMessage()               {}
func (*RaftTruncatedState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RaftTruncatedState) GetIndex() uint64 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

func (m *RaftTruncatedState) GetTerm() uint64 {
	if m != nil && m.Term != nil {
		return *m.Term
	}
	return 0
}

type KeyValue struct {
	Key              []byte `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *KeyValue) Reset()                    { *m = KeyValue{} }
func (m *KeyValue) String() string            { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()               {}
func (*KeyValue) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RaftSnapshotData struct {
	Region           *metapb.Region `protobuf:"bytes,1,opt,name=region" json:"region,omitempty"`
	Data             []*KeyValue    `protobuf:"bytes,2,rep,name=data" json:"data,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *RaftSnapshotData) Reset()                    { *m = RaftSnapshotData{} }
func (m *RaftSnapshotData) String() string            { return proto.CompactTextString(m) }
func (*RaftSnapshotData) ProtoMessage()               {}
func (*RaftSnapshotData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RaftSnapshotData) GetRegion() *metapb.Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *RaftSnapshotData) GetData() []*KeyValue {
	if m != nil {
		return m.Data
	}
	return nil
}

type StoreIdent struct {
	ClusterId        *uint64 `protobuf:"varint,1,opt,name=cluster_id" json:"cluster_id,omitempty"`
	StoreId          *uint64 `protobuf:"varint,2,opt,name=store_id" json:"store_id,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *StoreIdent) Reset()                    { *m = StoreIdent{} }
func (m *StoreIdent) String() string            { return proto.CompactTextString(m) }
func (*StoreIdent) ProtoMessage()               {}
func (*StoreIdent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StoreIdent) GetClusterId() uint64 {
	if m != nil && m.ClusterId != nil {
		return *m.ClusterId
	}
	return 0
}

func (m *StoreIdent) GetStoreId() uint64 {
	if m != nil && m.StoreId != nil {
		return *m.StoreId
	}
	return 0
}

func init() {
	proto.RegisterType((*RaftMessage)(nil), "raft_storepb.RaftMessage")
	proto.RegisterType((*RaftTruncatedState)(nil), "raft_storepb.RaftTruncatedState")
	proto.RegisterType((*KeyValue)(nil), "raft_storepb.KeyValue")
	proto.RegisterType((*RaftSnapshotData)(nil), "raft_storepb.RaftSnapshotData")
	proto.RegisterType((*StoreIdent)(nil), "raft_storepb.StoreIdent")
}

var fileDescriptor0 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xa9, 0x49, 0x6d, 0x3b, 0x59, 0xb5, 0x8e, 0x20, 0xa1, 0xa0, 0x96, 0x20, 0xa2, 0x97,
	0x80, 0xf9, 0x0d, 0x7a, 0x10, 0x11, 0xa4, 0x11, 0xf1, 0x16, 0xd6, 0x64, 0xda, 0x06, 0x9b, 0x6c,
	0xd8, 0x6c, 0xc4, 0xfe, 0x28, 0xff, 0xa3, 0xbb, 0x9b, 0x04, 0x2d, 0x78, 0x09, 0xcc, 0x7b, 0x6f,
	0xde, 0x7c, 0x64, 0x01, 0x25, 0x5f, 0xaa, 0xa4, 0x56, 0x42, 0x52, 0xf5, 0x1e, 0x56, 0x52, 0x28,
	0x81, 0xec, 0xaf, 0x36, 0xb3, 0x53, 0xef, 0xcd, 0x58, 0x41, 0x8a, 0xf7, 0x53, 0xf0, 0x3d, 0x00,
	0x6f, 0xa1, 0xed, 0x27, 0xaa, 0x6b, 0xbe, 0x22, 0x3c, 0x86, 0x89, 0xa4, 0x55, 0x2e, 0xca, 0x24,
	0xcf, 0xfc, 0xc1, 0x7c, 0x70, 0xed, 0xe2, 0x05, 0x4c, 0x96, 0x52, 0x14, 0x49, 0x45, 0x24, 0xfd,
	0x3d, 0x2d, 0x79, 0x11, 0x0b, 0xbb, 0x92, 0x67, 0xad, 0xe1, 0x19, 0x8c, 0x94, 0x68, 0x6d, 0xe7,
	0x1f, 0x7b, 0x0e, 0xa3, 0xa2, 0x6d, 0xf7, 0x5d, 0x6b, 0x1f, 0x85, 0x1d, 0x50, 0x7f, 0xf4, 0x06,
	0x58, 0x77, 0x94, 0x2a, 0x91, 0xae, 0xfd, 0xa1, 0x8d, 0x9d, 0xf4, 0x2d, 0x0b, 0xeb, 0xdd, 0x1b,
	0x2b, 0xb8, 0x05, 0x34, 0xb8, 0x2f, 0xb2, 0x29, 0x53, 0xae, 0x28, 0x8b, 0x95, 0xfe, 0xe2, 0x01,
	0x0c, 0xf3, 0x32, 0xa3, 0xaf, 0x8e, 0x98, 0x81, 0xab, 0x48, 0x16, 0x16, 0xd6, 0x0d, 0xae, 0x60,
	0xfc, 0x48, 0xdb, 0x57, 0xbe, 0x69, 0x08, 0x3d, 0x70, 0x3e, 0x68, 0x6b, 0x63, 0xcc, 0x6c, 0x7d,
	0x1a, 0xd5, 0xe6, 0x58, 0xf0, 0x06, 0x53, 0x53, 0x1d, 0x97, 0xbc, 0xaa, 0xd7, 0x42, 0xdd, 0x71,
	0xc5, 0xf1, 0x1c, 0xf6, 0x5b, 0x32, 0xbb, 0xe2, 0x45, 0x87, 0xbb, 0x4c, 0x78, 0x09, 0x6e, 0xa6,
	0x73, 0xba, 0xc1, 0xd1, 0xee, 0x69, 0xb8, 0xf3, 0x16, 0xfd, 0xd5, 0x20, 0x02, 0x88, 0x8d, 0xf6,
	0x90, 0x51, 0xa9, 0x10, 0x01, 0xd2, 0x4d, 0x53, 0x6b, 0xc2, 0xdf, 0x7f, 0x3c, 0x85, 0xb1, 0xdd,
	0x32, 0x8a, 0xa5, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xce, 0xc8, 0x5b, 0xd7, 0x01, 0x00,
	0x00,
}