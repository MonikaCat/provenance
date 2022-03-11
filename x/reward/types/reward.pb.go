// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/reward/v1/reward.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/provenance-io/provenance/x/epoch/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Params holds parameters for the reward module
type Params struct {
	// distribution epoch identifier
	DistrEpochIdentifier string `protobuf:"bytes,1,opt,name=distr_epoch_identifier,json=distrEpochIdentifier,proto3" json:"distr_epoch_identifier,omitempty" yaml:"distr_epoch_identifier"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c3894741a216575, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDistrEpochIdentifier() string {
	if m != nil {
		return m.DistrEpochIdentifier
	}
	return ""
}

type Criteria struct {
	Id           uint64                                   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DistributeTo string                                   `protobuf:"bytes,2,opt,name=distribute_to,json=distributeTo,proto3" json:"distribute_to,omitempty"`
	Coins        github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	// distribution start time
	StartHeight      int64                                    `protobuf:"varint,4,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty"`
	EndHeight        int64                                    `protobuf:"varint,5,opt,name=end_height,json=endHeight,proto3" json:"end_height,omitempty"`
	DistributedCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,6,rep,name=distributed_coins,json=distributedCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"distributed_coins"`
}

func (m *Criteria) Reset()         { *m = Criteria{} }
func (m *Criteria) String() string { return proto.CompactTextString(m) }
func (*Criteria) ProtoMessage()    {}
func (*Criteria) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c3894741a216575, []int{1}
}
func (m *Criteria) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Criteria) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Criteria.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Criteria) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Criteria.Merge(m, src)
}
func (m *Criteria) XXX_Size() int {
	return m.Size()
}
func (m *Criteria) XXX_DiscardUnknown() {
	xxx_messageInfo_Criteria.DiscardUnknown(m)
}

var xxx_messageInfo_Criteria proto.InternalMessageInfo

func (m *Criteria) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Criteria) GetDistributeTo() string {
	if m != nil {
		return m.DistributeTo
	}
	return ""
}

func (m *Criteria) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *Criteria) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *Criteria) GetEndHeight() int64 {
	if m != nil {
		return m.EndHeight
	}
	return 0
}

func (m *Criteria) GetDistributedCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DistributedCoins
	}
	return nil
}

// A Reward is the metadata of reward data per address
type Reward struct {
	// address of user reward
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty" yaml:"address"`
	// total initial reward amount for the user
	InitialRewardAmount github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=initial_reward_amount,json=initialRewardAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"initial_reward_amount" yaml:"initial_reward_amount"`
	// true if action is completed
	// index of bool in array refers to action enum #
	ActionCompleted []bool `protobuf:"varint,3,rep,packed,name=action_completed,json=actionCompleted,proto3" json:"action_completed,omitempty" yaml:"action_completed"`
}

func (m *Reward) Reset()         { *m = Reward{} }
func (m *Reward) String() string { return proto.CompactTextString(m) }
func (*Reward) ProtoMessage()    {}
func (*Reward) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c3894741a216575, []int{2}
}
func (m *Reward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Reward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Reward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Reward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reward.Merge(m, src)
}
func (m *Reward) XXX_Size() int {
	return m.Size()
}
func (m *Reward) XXX_DiscardUnknown() {
	xxx_messageInfo_Reward.DiscardUnknown(m)
}

var xxx_messageInfo_Reward proto.InternalMessageInfo

func (m *Reward) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Reward) GetInitialRewardAmount() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.InitialRewardAmount
	}
	return nil
}

func (m *Reward) GetActionCompleted() []bool {
	if m != nil {
		return m.ActionCompleted
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "provenance.reward.v1.Params")
	proto.RegisterType((*Criteria)(nil), "provenance.reward.v1.Criteria")
	proto.RegisterType((*Reward)(nil), "provenance.reward.v1.Reward")
}

func init() { proto.RegisterFile("provenance/reward/v1/reward.proto", fileDescriptor_0c3894741a216575) }

var fileDescriptor_0c3894741a216575 = []byte{
	// 554 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xc1, 0x6e, 0xd3, 0x3e,
	0x1c, 0x6e, 0xd2, 0xad, 0xff, 0xcd, 0xdb, 0x7f, 0x0c, 0x53, 0x58, 0x19, 0x2c, 0x69, 0xc3, 0xa5,
	0x07, 0x16, 0xd3, 0x71, 0xe3, 0x46, 0x2b, 0x10, 0xdc, 0xaa, 0x08, 0x09, 0x89, 0x4b, 0xe4, 0xc4,
	0x5e, 0x6a, 0xd1, 0xc4, 0x91, 0xed, 0x96, 0xed, 0x2d, 0x78, 0x03, 0x38, 0xf3, 0x24, 0x3b, 0xee,
	0xc8, 0x29, 0xa0, 0xf6, 0x0d, 0xfa, 0x04, 0x28, 0x76, 0xa2, 0x56, 0x63, 0x12, 0x42, 0xe2, 0x54,
	0xfb, 0xfb, 0x7d, 0xfe, 0xfc, 0x7d, 0xf5, 0xef, 0x17, 0xd0, 0xcb, 0x05, 0x9f, 0xd3, 0x0c, 0x67,
	0x31, 0x45, 0x82, 0x7e, 0xc2, 0x82, 0xa0, 0xf9, 0xa0, 0x5a, 0xf9, 0xb9, 0xe0, 0x8a, 0xc3, 0xf6,
	0x9a, 0xe2, 0x57, 0x85, 0xf9, 0xe0, 0xb8, 0x9d, 0xf0, 0x84, 0x6b, 0x02, 0x2a, 0x57, 0x86, 0x7b,
	0xec, 0x24, 0x9c, 0x27, 0x53, 0x8a, 0xf4, 0x2e, 0x9a, 0x9d, 0x23, 0x32, 0x13, 0x58, 0x31, 0x9e,
	0x55, 0x75, 0xf7, 0x66, 0x5d, 0xb1, 0x94, 0x4a, 0x85, 0xd3, 0xbc, 0x16, 0x88, 0xb9, 0x4c, 0xb9,
	0x44, 0x11, 0x96, 0x14, 0xcd, 0x07, 0x11, 0x55, 0x78, 0x80, 0x62, 0xce, 0x6a, 0x81, 0x4d, 0xbf,
	0x34, 0xe7, 0xf1, 0xa4, 0xb4, 0x9b, 0xd0, 0x8c, 0x4a, 0x26, 0x0d, 0xc5, 0xc3, 0xa0, 0x35, 0xc6,
	0x02, 0xa7, 0x12, 0xbe, 0x07, 0x0f, 0x08, 0x93, 0x4a, 0x84, 0x9a, 0x19, 0x32, 0x42, 0x33, 0xc5,
	0xce, 0x19, 0x15, 0x1d, 0xab, 0x6b, 0xf5, 0x77, 0x87, 0xbd, 0x55, 0xe1, 0x9e, 0x5c, 0xe2, 0x74,
	0xfa, 0xc2, 0xbb, 0x9d, 0xe7, 0x05, 0x6d, 0x5d, 0x78, 0x55, 0xe2, 0x6f, 0xd7, 0x70, 0x61, 0x83,
	0x9d, 0x91, 0x60, 0x8a, 0x0a, 0x86, 0xe1, 0x01, 0xb0, 0x19, 0xd1, 0x8a, 0x5b, 0x81, 0xcd, 0x08,
	0x7c, 0x02, 0xfe, 0xd7, 0x87, 0x58, 0x34, 0x53, 0x34, 0x54, 0xbc, 0x63, 0x97, 0x97, 0x05, 0xfb,
	0x6b, 0xf0, 0x1d, 0x87, 0x18, 0x6c, 0x97, 0xa9, 0x64, 0xa7, 0xd9, 0x6d, 0xf6, 0xf7, 0xce, 0x1e,
	0xfa, 0x26, 0xb7, 0x5f, 0xe6, 0xf6, 0xab, 0xdc, 0xfe, 0x88, 0xb3, 0x6c, 0xf8, 0xec, 0xaa, 0x70,
	0x1b, 0xdf, 0x7e, 0xb8, 0xfd, 0x84, 0xa9, 0xc9, 0x2c, 0xf2, 0x63, 0x9e, 0xa2, 0xea, 0x4f, 0x32,
	0x3f, 0xa7, 0x92, 0x7c, 0x44, 0xea, 0x32, 0xa7, 0x52, 0x1f, 0x90, 0x81, 0x51, 0x86, 0x3d, 0xb0,
	0x2f, 0x15, 0x16, 0x2a, 0x9c, 0x50, 0x96, 0x4c, 0x54, 0x67, 0xab, 0x6b, 0xf5, 0x9b, 0xc1, 0x9e,
	0xc6, 0xde, 0x68, 0x08, 0x9e, 0x00, 0x40, 0x33, 0x52, 0x13, 0xb6, 0x35, 0x61, 0x97, 0x66, 0xa4,
	0x2a, 0x5f, 0x80, 0xbb, 0x6b, 0xd3, 0x24, 0x34, 0x86, 0x5b, 0xff, 0xde, 0xf0, 0xe1, 0xc6, 0x2d,
	0x1a, 0xf1, 0xbe, 0xda, 0xa0, 0x15, 0xe8, 0x5e, 0x83, 0x4f, 0xc1, 0x7f, 0x98, 0x10, 0x41, 0xa5,
	0xac, 0x5e, 0x0d, 0xae, 0x0a, 0xf7, 0xc0, 0xbc, 0x5a, 0x55, 0xf0, 0x82, 0x9a, 0x02, 0xbf, 0x58,
	0xe0, 0x3e, 0xcb, 0x98, 0x62, 0x78, 0x1a, 0x9a, 0x66, 0x0d, 0x71, 0xca, 0x67, 0x99, 0xea, 0xd8,
	0x7f, 0xf2, 0x3d, 0x2e, 0x7d, 0xaf, 0x0a, 0xf7, 0xb1, 0xd1, 0xbe, 0x55, 0xc5, 0xfb, 0xab, 0x5c,
	0xf7, 0x2a, 0x0d, 0x93, 0xe4, 0xa5, 0x56, 0x80, 0xaf, 0xc1, 0x21, 0x8e, 0xcb, 0x91, 0x08, 0x63,
	0x9e, 0xe6, 0x53, 0xaa, 0x28, 0xd1, 0x4d, 0xb0, 0x33, 0x7c, 0xb4, 0x2a, 0xdc, 0xa3, 0x2a, 0xd8,
	0x0d, 0x86, 0x17, 0xdc, 0x31, 0xd0, 0xa8, 0x46, 0x86, 0xc9, 0xd5, 0xc2, 0xb1, 0xae, 0x17, 0x8e,
	0xf5, 0x73, 0xe1, 0x58, 0x9f, 0x97, 0x4e, 0xe3, 0x7a, 0xe9, 0x34, 0xbe, 0x2f, 0x9d, 0x06, 0x38,
	0x62, 0x7a, 0x1c, 0x7f, 0x9b, 0xd9, 0xb1, 0xf5, 0xe1, 0x6c, 0xc3, 0xfb, 0x9a, 0x72, 0xca, 0xf8,
	0xc6, 0x0e, 0x5d, 0xd4, 0x5f, 0x02, 0x9d, 0x25, 0x6a, 0xe9, 0xb1, 0x7a, 0xfe, 0x2b, 0x00, 0x00,
	0xff, 0xff, 0xda, 0xa0, 0x85, 0x56, 0x2b, 0x04, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DistrEpochIdentifier) > 0 {
		i -= len(m.DistrEpochIdentifier)
		copy(dAtA[i:], m.DistrEpochIdentifier)
		i = encodeVarintReward(dAtA, i, uint64(len(m.DistrEpochIdentifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Criteria) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Criteria) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Criteria) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DistributedCoins) > 0 {
		for iNdEx := len(m.DistributedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DistributedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReward(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.EndHeight != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.EndHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.StartHeight != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReward(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.DistributeTo) > 0 {
		i -= len(m.DistributeTo)
		copy(dAtA[i:], m.DistributeTo)
		i = encodeVarintReward(dAtA, i, uint64(len(m.DistributeTo)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintReward(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Reward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Reward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ActionCompleted) > 0 {
		for iNdEx := len(m.ActionCompleted) - 1; iNdEx >= 0; iNdEx-- {
			i--
			if m.ActionCompleted[iNdEx] {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
		}
		i = encodeVarintReward(dAtA, i, uint64(len(m.ActionCompleted)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.InitialRewardAmount) > 0 {
		for iNdEx := len(m.InitialRewardAmount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InitialRewardAmount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintReward(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintReward(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintReward(dAtA []byte, offset int, v uint64) int {
	offset -= sovReward(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DistrEpochIdentifier)
	if l > 0 {
		n += 1 + l + sovReward(uint64(l))
	}
	return n
}

func (m *Criteria) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovReward(uint64(m.Id))
	}
	l = len(m.DistributeTo)
	if l > 0 {
		n += 1 + l + sovReward(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovReward(uint64(l))
		}
	}
	if m.StartHeight != 0 {
		n += 1 + sovReward(uint64(m.StartHeight))
	}
	if m.EndHeight != 0 {
		n += 1 + sovReward(uint64(m.EndHeight))
	}
	if len(m.DistributedCoins) > 0 {
		for _, e := range m.DistributedCoins {
			l = e.Size()
			n += 1 + l + sovReward(uint64(l))
		}
	}
	return n
}

func (m *Reward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovReward(uint64(l))
	}
	if len(m.InitialRewardAmount) > 0 {
		for _, e := range m.InitialRewardAmount {
			l = e.Size()
			n += 1 + l + sovReward(uint64(l))
		}
	}
	if len(m.ActionCompleted) > 0 {
		n += 1 + sovReward(uint64(len(m.ActionCompleted))) + len(m.ActionCompleted)*1
	}
	return n
}

func sovReward(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozReward(x uint64) (n int) {
	return sovReward(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReward
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistrEpochIdentifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistrEpochIdentifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReward
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Criteria) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReward
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Criteria: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Criteria: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributeTo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributeTo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndHeight", wireType)
			}
			m.EndHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributedCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributedCoins = append(m.DistributedCoins, types.Coin{})
			if err := m.DistributedCoins[len(m.DistributedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReward
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Reward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowReward
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Reward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialRewardAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowReward
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthReward
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthReward
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InitialRewardAmount = append(m.InitialRewardAmount, types.Coin{})
			if err := m.InitialRewardAmount[len(m.InitialRewardAmount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType == 0 {
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowReward
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowReward
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthReward
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthReward
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				elementCount = packedLen
				if elementCount != 0 && len(m.ActionCompleted) == 0 {
					m.ActionCompleted = make([]bool, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowReward
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.ActionCompleted = append(m.ActionCompleted, bool(v != 0))
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field ActionCompleted", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipReward(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthReward
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipReward(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowReward
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReward
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowReward
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthReward
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupReward
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthReward
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthReward        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowReward          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupReward = fmt.Errorf("proto: unexpected end of group")
)