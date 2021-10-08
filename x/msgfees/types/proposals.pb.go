// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provenance/msgfees/v1/proposals.proto

package types

import (
	fmt "fmt"
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx"
	_ "github.com/cosmos/cosmos-sdk/x/bank/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/provenance-io/provenance/x/marker/types"
	_ "github.com/regen-network/cosmos-proto"
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

// AddMsgBasedFeesProposal defines a governance proposal to add additional msg fees
type AddMsgBasedFeesProposal struct {
	Title       string                                  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
	Msg         *types1.Any                             `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	MinFee      types.Coin                              `protobuf:"bytes,5,opt,name=min_fee,json=minFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_fee" yaml:"min_fee"`
	//  Fee rate, based on Gas used.
	FeeRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=fee_rate,json=feeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_rate,omitempty"`
}

func (m *AddMsgBasedFeesProposal) Reset()      { *m = AddMsgBasedFeesProposal{} }
func (*AddMsgBasedFeesProposal) ProtoMessage() {}
func (*AddMsgBasedFeesProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e168825d6c34a4, []int{0}
}
func (m *AddMsgBasedFeesProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AddMsgBasedFeesProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AddMsgBasedFeesProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AddMsgBasedFeesProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMsgBasedFeesProposal.Merge(m, src)
}
func (m *AddMsgBasedFeesProposal) XXX_Size() int {
	return m.Size()
}
func (m *AddMsgBasedFeesProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMsgBasedFeesProposal.DiscardUnknown(m)
}

var xxx_messageInfo_AddMsgBasedFeesProposal proto.InternalMessageInfo

func (m *AddMsgBasedFeesProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddMsgBasedFeesProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *AddMsgBasedFeesProposal) GetMsg() *types1.Any {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *AddMsgBasedFeesProposal) GetMinFee() types.Coin {
	if m != nil {
		return m.MinFee
	}
	return types.Coin{}
}

// UpdateMsgBasedFeesProposal defines a governance proposal to update a current msg based fees
type UpdateMsgBasedFeesProposal struct {
	Title       string                                  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Amount      github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
	Msg         *types1.Any                             `protobuf:"bytes,4,opt,name=msg,proto3" json:"msg,omitempty"`
	MinFee      types.Coin                              `protobuf:"bytes,5,opt,name=min_fee,json=minFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"min_fee" yaml:"min_fee"`
	//  Fee rate, based on Gas used.
	FeeRate github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=fee_rate,json=feeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_rate,omitempty"`
}

func (m *UpdateMsgBasedFeesProposal) Reset()      { *m = UpdateMsgBasedFeesProposal{} }
func (*UpdateMsgBasedFeesProposal) ProtoMessage() {}
func (*UpdateMsgBasedFeesProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e168825d6c34a4, []int{1}
}
func (m *UpdateMsgBasedFeesProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UpdateMsgBasedFeesProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UpdateMsgBasedFeesProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UpdateMsgBasedFeesProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateMsgBasedFeesProposal.Merge(m, src)
}
func (m *UpdateMsgBasedFeesProposal) XXX_Size() int {
	return m.Size()
}
func (m *UpdateMsgBasedFeesProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateMsgBasedFeesProposal.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateMsgBasedFeesProposal proto.InternalMessageInfo

func (m *UpdateMsgBasedFeesProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateMsgBasedFeesProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateMsgBasedFeesProposal) GetMsg() *types1.Any {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *UpdateMsgBasedFeesProposal) GetMinFee() types.Coin {
	if m != nil {
		return m.MinFee
	}
	return types.Coin{}
}

// RemoveMsgBasedFeesProposal defines a governance proposal to delete a current msg based fees
type RemoveMsgBasedFeesProposal struct {
	Title       string      `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Msg         *types1.Any `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *RemoveMsgBasedFeesProposal) Reset()      { *m = RemoveMsgBasedFeesProposal{} }
func (*RemoveMsgBasedFeesProposal) ProtoMessage() {}
func (*RemoveMsgBasedFeesProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2e168825d6c34a4, []int{2}
}
func (m *RemoveMsgBasedFeesProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RemoveMsgBasedFeesProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveMsgBasedFeesProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RemoveMsgBasedFeesProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveMsgBasedFeesProposal.Merge(m, src)
}
func (m *RemoveMsgBasedFeesProposal) XXX_Size() int {
	return m.Size()
}
func (m *RemoveMsgBasedFeesProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveMsgBasedFeesProposal.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveMsgBasedFeesProposal proto.InternalMessageInfo

func (m *RemoveMsgBasedFeesProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *RemoveMsgBasedFeesProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RemoveMsgBasedFeesProposal) GetMsg() *types1.Any {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*AddMsgBasedFeesProposal)(nil), "provenance.msgfees.v1.AddMsgBasedFeesProposal")
	proto.RegisterType((*UpdateMsgBasedFeesProposal)(nil), "provenance.msgfees.v1.UpdateMsgBasedFeesProposal")
	proto.RegisterType((*RemoveMsgBasedFeesProposal)(nil), "provenance.msgfees.v1.RemoveMsgBasedFeesProposal")
}

func init() {
	proto.RegisterFile("provenance/msgfees/v1/proposals.proto", fileDescriptor_a2e168825d6c34a4)
}

var fileDescriptor_a2e168825d6c34a4 = []byte{
	// 543 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xec, 0x94, 0x3f, 0x8b, 0x13, 0x4f,
	0x18, 0xc7, 0xb3, 0xbf, 0xfc, 0x2e, 0x77, 0xce, 0x89, 0xc5, 0x12, 0x71, 0x2f, 0xc5, 0x6e, 0x0c,
	0x78, 0xa6, 0x30, 0x33, 0xc4, 0x74, 0xd7, 0xdd, 0x7a, 0x5c, 0x77, 0x70, 0x2c, 0xd8, 0xd8, 0x84,
	0xd9, 0xdd, 0x27, 0xeb, 0x90, 0xcc, 0xcc, 0xb2, 0x33, 0x59, 0x92, 0x77, 0x21, 0xd8, 0x58, 0x5e,
	0x6d, 0xed, 0x8b, 0x38, 0x04, 0x21, 0xa5, 0x58, 0x44, 0x49, 0x1a, 0xb1, 0xf4, 0x15, 0x48, 0x76,
	0x27, 0x7f, 0xe0, 0x04, 0x4f, 0xb0, 0x12, 0xab, 0x9d, 0x67, 0xbf, 0xdf, 0x99, 0xf9, 0xcc, 0xf7,
	0x81, 0x07, 0x3d, 0x4a, 0x33, 0x99, 0x83, 0xa0, 0x22, 0x02, 0xc2, 0x55, 0x32, 0x00, 0x50, 0x24,
	0xef, 0x92, 0x34, 0x93, 0xa9, 0x54, 0x74, 0xa4, 0x70, 0x9a, 0x49, 0x2d, 0xed, 0xfb, 0x5b, 0x1b,
	0x36, 0x36, 0x9c, 0x77, 0x1b, 0xf5, 0x44, 0x26, 0xb2, 0x70, 0x90, 0xd5, 0xaa, 0x34, 0x37, 0xdc,
	0x48, 0x2a, 0x2e, 0x15, 0x09, 0xa9, 0x18, 0x92, 0xbc, 0x1b, 0x82, 0xa6, 0xdd, 0xa2, 0xb8, 0xa1,
	0x2b, 0xd8, 0xe8, 0x91, 0x64, 0xc2, 0xe8, 0x0f, 0x77, 0x99, 0x68, 0x36, 0x84, 0x6c, 0x85, 0x54,
	0xae, 0x8c, 0xe5, 0xf8, 0xa7, 0x16, 0x1a, 0x45, 0xa0, 0x54, 0x92, 0x51, 0xa1, 0x8d, 0xef, 0x28,
	0x91, 0x32, 0x19, 0x01, 0x29, 0xaa, 0x70, 0x3c, 0x20, 0x54, 0x4c, 0x8d, 0xd4, 0x30, 0x14, 0x7a,
	0xb2, 0x61, 0xd0, 0x93, 0xf5, 0xb6, 0x52, 0xeb, 0x97, 0x4f, 0x2b, 0x8b, 0x52, 0x6a, 0x7d, 0xa8,
	0xa2, 0x07, 0xa7, 0x71, 0x7c, 0xa1, 0x12, 0x9f, 0x2a, 0x88, 0xcf, 0x01, 0xd4, 0xa5, 0x09, 0xcb,
	0xae, 0xa3, 0x3d, 0xcd, 0xf4, 0x08, 0x1c, 0xab, 0x69, 0xb5, 0xef, 0x04, 0x65, 0x61, 0x37, 0xd1,
	0x61, 0x0c, 0x2a, 0xca, 0x58, 0xaa, 0x99, 0x14, 0xce, 0x7f, 0x85, 0xb6, 0xfb, 0xcb, 0x0e, 0x51,
	0x8d, 0x72, 0x39, 0x16, 0xda, 0xa9, 0x36, 0xad, 0xf6, 0xe1, 0xd3, 0x23, 0x6c, 0xae, 0x5c, 0x25,
	0x84, 0x0d, 0x1d, 0x7e, 0x26, 0x99, 0xf0, 0xc9, 0xf5, 0xdc, 0xab, 0x7c, 0x9a, 0x7b, 0x8f, 0x13,
	0xa6, 0x5f, 0x8e, 0x43, 0x1c, 0x49, 0x6e, 0xf8, 0xcc, 0xa7, 0xa3, 0xe2, 0x21, 0xd1, 0xd3, 0x14,
	0x54, 0xb1, 0x21, 0x30, 0x27, 0xdb, 0x3d, 0x54, 0xe5, 0x2a, 0x71, 0xfe, 0x2f, 0x2e, 0xa8, 0xe3,
	0x32, 0x17, 0xbc, 0xce, 0x05, 0x9f, 0x8a, 0xa9, 0x7f, 0xf8, 0xfe, 0x5d, 0x67, 0x5f, 0xc5, 0x43,
	0x7c, 0xa1, 0x92, 0x60, 0xe5, 0xb6, 0x73, 0xb4, 0xcf, 0x99, 0xe8, 0x0f, 0x00, 0x9c, 0xbd, 0x5f,
	0x91, 0xf9, 0x2b, 0xb2, 0xef, 0x73, 0xef, 0xde, 0x94, 0xf2, 0xd1, 0x49, 0xcb, 0xec, 0x6b, 0xbd,
	0xfd, 0xec, 0xb5, 0x6f, 0xc9, 0xaa, 0x82, 0x1a, 0x67, 0xe2, 0x1c, 0xc0, 0xee, 0xa3, 0x83, 0x01,
	0x40, 0x3f, 0xa3, 0x1a, 0x9c, 0x5a, 0xd3, 0x6a, 0xdf, 0xf5, 0xcf, 0xcc, 0xbb, 0x8f, 0x6f, 0x71,
	0xd6, 0x19, 0x44, 0xdf, 0xe6, 0x9e, 0xbd, 0x3e, 0xe1, 0x89, 0xe4, 0x4c, 0x03, 0x4f, 0xf5, 0x34,
	0xd8, 0x1f, 0x00, 0x04, 0x54, 0xc3, 0xc9, 0xc1, 0x9b, 0x2b, 0xaf, 0xf2, 0xf5, 0xca, 0xb3, 0x5a,
	0xb3, 0x2a, 0x6a, 0x3c, 0x4f, 0x63, 0xaa, 0xe1, 0x5f, 0x4b, 0xff, 0x92, 0x96, 0xbe, 0xb6, 0x50,
	0x23, 0x00, 0x2e, 0xf3, 0x3f, 0xdb, 0x52, 0x13, 0x77, 0xf5, 0x77, 0xe2, 0xde, 0x52, 0xf9, 0xec,
	0x7a, 0xe1, 0x5a, 0xb3, 0x85, 0x6b, 0x7d, 0x59, 0xb8, 0xd6, 0xab, 0xa5, 0x5b, 0x99, 0x2d, 0xdd,
	0xca, 0xc7, 0xa5, 0x5b, 0x41, 0x0e, 0x2b, 0x26, 0xe7, 0xcd, 0xf9, 0x7a, 0x69, 0xbd, 0xe8, 0xed,
	0x84, 0xb3, 0xf5, 0x74, 0x98, 0xdc, 0xa9, 0xc8, 0x64, 0x33, 0xba, 0x8b, 0xb4, 0xc2, 0x5a, 0x01,
	0xd5, 0xfb, 0x11, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x3c, 0x6e, 0xac, 0xdd, 0x05, 0x00, 0x00,
}

func (this *AddMsgBasedFeesProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AddMsgBasedFeesProposal)
	if !ok {
		that2, ok := that.(AddMsgBasedFeesProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if !this.Amount.Equal(that1.Amount) {
		return false
	}
	if !this.Msg.Equal(that1.Msg) {
		return false
	}
	if !this.MinFee.Equal(&that1.MinFee) {
		return false
	}
	if !this.FeeRate.Equal(that1.FeeRate) {
		return false
	}
	return true
}
func (this *UpdateMsgBasedFeesProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateMsgBasedFeesProposal)
	if !ok {
		that2, ok := that.(UpdateMsgBasedFeesProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if !this.Amount.Equal(that1.Amount) {
		return false
	}
	if !this.Msg.Equal(that1.Msg) {
		return false
	}
	if !this.MinFee.Equal(&that1.MinFee) {
		return false
	}
	if !this.FeeRate.Equal(that1.FeeRate) {
		return false
	}
	return true
}
func (this *RemoveMsgBasedFeesProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RemoveMsgBasedFeesProposal)
	if !ok {
		that2, ok := that.(RemoveMsgBasedFeesProposal)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if !this.Msg.Equal(that1.Msg) {
		return false
	}
	return true
}
func (m *AddMsgBasedFeesProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AddMsgBasedFeesProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AddMsgBasedFeesProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeeRate.Size()
		i -= size
		if _, err := m.FeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.MinFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProposals(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UpdateMsgBasedFeesProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UpdateMsgBasedFeesProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UpdateMsgBasedFeesProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeeRate.Size()
		i -= size
		if _, err := m.FeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.MinFee.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProposals(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintProposals(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RemoveMsgBasedFeesProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveMsgBasedFeesProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveMsgBasedFeesProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintProposals(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposals(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposals(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposals(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AddMsgBasedFeesProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovProposals(uint64(l))
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovProposals(uint64(l))
	}
	l = m.MinFee.Size()
	n += 1 + l + sovProposals(uint64(l))
	l = m.FeeRate.Size()
	n += 1 + l + sovProposals(uint64(l))
	return n
}

func (m *UpdateMsgBasedFeesProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovProposals(uint64(l))
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovProposals(uint64(l))
	}
	l = m.MinFee.Size()
	n += 1 + l + sovProposals(uint64(l))
	l = m.FeeRate.Size()
	n += 1 + l + sovProposals(uint64(l))
	return n
}

func (m *RemoveMsgBasedFeesProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposals(uint64(l))
	}
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovProposals(uint64(l))
	}
	return n
}

func sovProposals(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposals(x uint64) (n int) {
	return sovProposals(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AddMsgBasedFeesProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: AddMsgBasedFeesProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AddMsgBasedFeesProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &types1.Any{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func (m *UpdateMsgBasedFeesProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: UpdateMsgBasedFeesProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UpdateMsgBasedFeesProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &types1.Any{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MinFee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func (m *RemoveMsgBasedFeesProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposals
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
			return fmt.Errorf("proto: RemoveMsgBasedFeesProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveMsgBasedFeesProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposals
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
				return ErrInvalidLengthProposals
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposals
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &types1.Any{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposals(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposals
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
func skipProposals(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
					return 0, ErrIntOverflowProposals
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
				return 0, ErrInvalidLengthProposals
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposals
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposals
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposals        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposals          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposals = fmt.Errorf("proto: unexpected end of group")
)
