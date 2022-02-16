// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enci/oracle/v1beta1/oracle.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_tendermint_tendermint_libs_bytes "github.com/tendermint/tendermint/libs/bytes"
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

// Vote is a vote for a given claim by a validator
type Vote struct {
	RoundId     uint64                                               `protobuf:"varint,1,opt,name=roundId,proto3" json:"roundId,omitempty"`
	ClaimHash   github_com_tendermint_tendermint_libs_bytes.HexBytes `protobuf:"bytes,2,opt,name=claimHash,proto3,casttype=github.com/tendermint/tendermint/libs/bytes.HexBytes" json:"claimHash,omitempty"`
	ConsensusId string                                               `protobuf:"bytes,3,opt,name=consensusId,proto3" json:"consensusId,omitempty"`
	ClaimType   string                                               `protobuf:"bytes,4,opt,name=claimType,proto3" json:"claimType,omitempty"`
	Validator   github_com_cosmos_cosmos_sdk_types.ValAddress        `protobuf:"bytes,5,opt,name=validator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"validator,omitempty"`
}

func (m *Vote) Reset()         { *m = Vote{} }
func (m *Vote) String() string { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()    {}
func (*Vote) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5e6dcb96c8266e5, []int{0}
}
func (m *Vote) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Vote) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Vote.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Vote) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vote.Merge(m, src)
}
func (m *Vote) XXX_Size() int {
	return m.Size()
}
func (m *Vote) XXX_DiscardUnknown() {
	xxx_messageInfo_Vote.DiscardUnknown(m)
}

var xxx_messageInfo_Vote proto.InternalMessageInfo

func (m *Vote) GetRoundId() uint64 {
	if m != nil {
		return m.RoundId
	}
	return 0
}

func (m *Vote) GetClaimHash() github_com_tendermint_tendermint_libs_bytes.HexBytes {
	if m != nil {
		return m.ClaimHash
	}
	return nil
}

func (m *Vote) GetConsensusId() string {
	if m != nil {
		return m.ConsensusId
	}
	return ""
}

func (m *Vote) GetClaimType() string {
	if m != nil {
		return m.ClaimType
	}
	return ""
}

func (m *Vote) GetValidator() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.Validator
	}
	return nil
}

// Round contains all claim votes for a given round
type Round struct {
	RoundId   uint64 `protobuf:"varint,1,opt,name=roundId,proto3" json:"roundId,omitempty"`
	ClaimType string `protobuf:"bytes,2,opt,name=claimType,proto3" json:"claimType,omitempty"`
	Votes     []Vote `protobuf:"bytes,3,rep,name=votes,proto3" json:"votes"`
}

func (m *Round) Reset()         { *m = Round{} }
func (m *Round) String() string { return proto.CompactTextString(m) }
func (*Round) ProtoMessage()    {}
func (*Round) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5e6dcb96c8266e5, []int{1}
}
func (m *Round) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Round) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Round.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Round) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Round.Merge(m, src)
}
func (m *Round) XXX_Size() int {
	return m.Size()
}
func (m *Round) XXX_DiscardUnknown() {
	xxx_messageInfo_Round.DiscardUnknown(m)
}

var xxx_messageInfo_Round proto.InternalMessageInfo

func (m *Round) GetRoundId() uint64 {
	if m != nil {
		return m.RoundId
	}
	return 0
}

func (m *Round) GetClaimType() string {
	if m != nil {
		return m.ClaimType
	}
	return ""
}

func (m *Round) GetVotes() []Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

// TestClaim is a concrete Claim type we use for testing
type TestClaim struct {
	BlockHeight int64  `protobuf:"varint,1,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	ClaimType   string `protobuf:"bytes,2,opt,name=claimType,proto3" json:"claimType,omitempty"`
	Content     string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *TestClaim) Reset()         { *m = TestClaim{} }
func (m *TestClaim) String() string { return proto.CompactTextString(m) }
func (*TestClaim) ProtoMessage()    {}
func (*TestClaim) Descriptor() ([]byte, []int) {
	return fileDescriptor_f5e6dcb96c8266e5, []int{2}
}
func (m *TestClaim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestClaim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestClaim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestClaim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestClaim.Merge(m, src)
}
func (m *TestClaim) XXX_Size() int {
	return m.Size()
}
func (m *TestClaim) XXX_DiscardUnknown() {
	xxx_messageInfo_TestClaim.DiscardUnknown(m)
}

var xxx_messageInfo_TestClaim proto.InternalMessageInfo

func (m *TestClaim) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *TestClaim) GetClaimType() string {
	if m != nil {
		return m.ClaimType
	}
	return ""
}

func (m *TestClaim) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*Vote)(nil), "enci.oracle.v1beta1.Vote")
	proto.RegisterType((*Round)(nil), "enci.oracle.v1beta1.Round")
	proto.RegisterType((*TestClaim)(nil), "enci.oracle.v1beta1.TestClaim")
}

func init() { proto.RegisterFile("enci/oracle/v1beta1/oracle.proto", fileDescriptor_f5e6dcb96c8266e5) }

var fileDescriptor_f5e6dcb96c8266e5 = []byte{
	// 394 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x3d, 0xaf, 0x9b, 0x30,
	0x14, 0x85, 0x40, 0xfa, 0x84, 0x5f, 0x27, 0xda, 0xc1, 0xad, 0x2a, 0x82, 0x32, 0xb1, 0x3c, 0x50,
	0xfa, 0x21, 0x75, 0x2d, 0x5d, 0xf2, 0xa6, 0x4a, 0xe8, 0x29, 0x43, 0x37, 0x63, 0x5f, 0x81, 0x15,
	0xb0, 0x23, 0xec, 0xa0, 0xe4, 0x27, 0x74, 0xeb, 0xcf, 0xca, 0x98, 0xb1, 0x53, 0x54, 0x25, 0xff,
	0x22, 0x53, 0x65, 0x42, 0x14, 0x5a, 0x55, 0xed, 0xc4, 0xb9, 0xe7, 0x1e, 0xdd, 0x7b, 0x38, 0xbe,
	0x28, 0x04, 0x41, 0x79, 0x22, 0x1b, 0x42, 0x2b, 0x48, 0xda, 0x59, 0x0e, 0x9a, 0xcc, 0xfa, 0x32,
	0x5e, 0x35, 0x52, 0x4b, 0xff, 0x85, 0x51, 0xc4, 0x3d, 0xd5, 0x2b, 0x5e, 0xbf, 0x2c, 0x64, 0x21,
	0xbb, 0x7e, 0x62, 0xd0, 0x45, 0x3a, 0xfd, 0x36, 0x42, 0xee, 0x42, 0x6a, 0xf0, 0x31, 0xba, 0x6b,
	0xe4, 0x5a, 0xb0, 0x47, 0x86, 0xed, 0xd0, 0x8e, 0xdc, 0xec, 0x5a, 0xfa, 0x0b, 0xe4, 0xd1, 0x8a,
	0xf0, 0x7a, 0x4e, 0x54, 0x89, 0x47, 0xa1, 0x1d, 0x3d, 0x4f, 0x3f, 0x9e, 0x0f, 0x93, 0xf7, 0x05,
	0xd7, 0xe5, 0x3a, 0x8f, 0xa9, 0xac, 0x13, 0x0d, 0x82, 0x41, 0x53, 0x73, 0xa1, 0x87, 0xb0, 0xe2,
	0xb9, 0x4a, 0xf2, 0xad, 0x06, 0x15, 0xcf, 0x61, 0x93, 0x1a, 0x90, 0xdd, 0x46, 0xf9, 0x21, 0xba,
	0xa7, 0x52, 0x28, 0x10, 0x6a, 0xad, 0x1e, 0x19, 0x76, 0x42, 0x3b, 0xf2, 0xb2, 0x21, 0xe5, 0xbf,
	0xe9, 0x37, 0x3f, 0x6d, 0x57, 0x80, 0xdd, 0xae, 0x7f, 0x23, 0xfc, 0x2f, 0xc8, 0x6b, 0x49, 0xc5,
	0x19, 0xd1, 0xb2, 0xc1, 0xe3, 0xce, 0xd7, 0xec, 0x7c, 0x98, 0x3c, 0x0c, 0x7c, 0x51, 0xa9, 0x6a,
	0xa9, 0xfa, 0xcf, 0x83, 0x62, 0xcb, 0x44, 0x6f, 0x57, 0xa0, 0xe2, 0x05, 0xa9, 0x3e, 0x31, 0xd6,
	0x80, 0x52, 0xd9, 0x6d, 0xc6, 0xb4, 0x45, 0xe3, 0xcc, 0xfc, 0xf3, 0x3f, 0xb2, 0xf8, 0xcd, 0xd1,
	0xe8, 0x4f, 0x47, 0x1f, 0xd0, 0xb8, 0x95, 0x1a, 0x14, 0x76, 0x42, 0x27, 0xba, 0x7f, 0xfb, 0x2a,
	0xfe, 0xcb, 0x3b, 0xc4, 0x26, 0xed, 0xd4, 0xdd, 0x1d, 0x26, 0x56, 0x76, 0x51, 0x4f, 0x01, 0x79,
	0x4f, 0xa0, 0xf4, 0x67, 0x33, 0xc7, 0xa4, 0x92, 0x57, 0x92, 0x2e, 0xe7, 0xc0, 0x8b, 0x52, 0x77,
	0xfb, 0x9d, 0x6c, 0x48, 0xfd, 0xc7, 0x03, 0x46, 0x77, 0x54, 0x0a, 0x0d, 0x42, 0xf7, 0x89, 0x5e,
	0xcb, 0x34, 0xdd, 0x1d, 0x03, 0x7b, 0x7f, 0x0c, 0xec, 0x9f, 0xc7, 0xc0, 0xfe, 0x7e, 0x0a, 0xac,
	0xfd, 0x29, 0xb0, 0x7e, 0x9c, 0x02, 0xeb, 0x6b, 0x34, 0x88, 0xcc, 0x58, 0xa6, 0x25, 0xe1, 0xa2,
	0x43, 0xc9, 0xe6, 0x7a, 0x68, 0x5d, 0x70, 0xf9, 0xb3, 0xee, 0x6a, 0xde, 0xfd, 0x0a, 0x00, 0x00,
	0xff, 0xff, 0xd0, 0xb7, 0x08, 0x93, 0x84, 0x02, 0x00, 0x00,
}

func (m *Vote) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Vote) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Vote) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ClaimType) > 0 {
		i -= len(m.ClaimType)
		copy(dAtA[i:], m.ClaimType)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ClaimType)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ConsensusId) > 0 {
		i -= len(m.ConsensusId)
		copy(dAtA[i:], m.ConsensusId)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ConsensusId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClaimHash) > 0 {
		i -= len(m.ClaimHash)
		copy(dAtA[i:], m.ClaimHash)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ClaimHash)))
		i--
		dAtA[i] = 0x12
	}
	if m.RoundId != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.RoundId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Round) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Round) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Round) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintOracle(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ClaimType) > 0 {
		i -= len(m.ClaimType)
		copy(dAtA[i:], m.ClaimType)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ClaimType)))
		i--
		dAtA[i] = 0x12
	}
	if m.RoundId != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.RoundId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *TestClaim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestClaim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TestClaim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ClaimType) > 0 {
		i -= len(m.ClaimType)
		copy(dAtA[i:], m.ClaimType)
		i = encodeVarintOracle(dAtA, i, uint64(len(m.ClaimType)))
		i--
		dAtA[i] = 0x12
	}
	if m.BlockHeight != 0 {
		i = encodeVarintOracle(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintOracle(dAtA []byte, offset int, v uint64) int {
	offset -= sovOracle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Vote) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoundId != 0 {
		n += 1 + sovOracle(uint64(m.RoundId))
	}
	l = len(m.ClaimHash)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.ConsensusId)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.ClaimType)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func (m *Round) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RoundId != 0 {
		n += 1 + sovOracle(uint64(m.RoundId))
	}
	l = len(m.ClaimType)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovOracle(uint64(l))
		}
	}
	return n
}

func (m *TestClaim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovOracle(uint64(m.BlockHeight))
	}
	l = len(m.ClaimType)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovOracle(uint64(l))
	}
	return n
}

func sovOracle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOracle(x uint64) (n int) {
	return sovOracle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Vote) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Vote: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Vote: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoundId", wireType)
			}
			m.RoundId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoundId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimHash = append(m.ClaimHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ClaimHash == nil {
				m.ClaimHash = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = append(m.Validator[:0], dAtA[iNdEx:postIndex]...)
			if m.Validator == nil {
				m.Validator = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *Round) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: Round: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Round: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoundId", wireType)
			}
			m.RoundId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoundId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func (m *TestClaim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOracle
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
			return fmt.Errorf("proto: TestClaim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestClaim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOracle
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
				return ErrInvalidLengthOracle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOracle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOracle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOracle
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
func skipOracle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
					return 0, ErrIntOverflowOracle
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
				return 0, ErrInvalidLengthOracle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOracle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOracle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOracle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOracle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOracle = fmt.Errorf("proto: unexpected end of group")
)
