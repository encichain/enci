// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enci/charity/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the charity module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	// this line is used by starport scaffolding # ibc/genesis/proto
	Params           Params                                   `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TaxRateLimits    TaxRateLimits                            `protobuf:"bytes,2,opt,name=tax_rate_limits,json=taxRateLimits,proto3" json:"tax_rate_limits"`
	TaxCaps          []TaxCap                                 `protobuf:"bytes,3,rep,name=tax_caps,json=taxCaps,proto3" json:"tax_caps"`
	TaxProceeds      github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=tax_proceeds,json=taxProceeds,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"tax_proceeds"`
	CollectionEpochs []CollectionEpoch                        `protobuf:"bytes,5,rep,name=collection_epochs,json=collectionEpochs,proto3" json:"collection_epochs"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_038f4082053f175f, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTaxRateLimits() TaxRateLimits {
	if m != nil {
		return m.TaxRateLimits
	}
	return TaxRateLimits{}
}

func (m *GenesisState) GetTaxCaps() []TaxCap {
	if m != nil {
		return m.TaxCaps
	}
	return nil
}

func (m *GenesisState) GetTaxProceeds() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.TaxProceeds
	}
	return nil
}

func (m *GenesisState) GetCollectionEpochs() []CollectionEpoch {
	if m != nil {
		return m.CollectionEpochs
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "enci.charity.v1beta1.GenesisState")
}

func init() {
	proto.RegisterFile("enci/charity/v1beta1/genesis.proto", fileDescriptor_038f4082053f175f)
}

var fileDescriptor_038f4082053f175f = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xc1, 0x4e, 0xfa, 0x30,
	0x1c, 0xc7, 0xb7, 0x3f, 0xfc, 0xd1, 0x0c, 0x8c, 0xba, 0x70, 0x98, 0xc4, 0x0c, 0x82, 0x31, 0xc1,
	0x83, 0xad, 0xe8, 0xcd, 0xc4, 0x0b, 0x8b, 0xf1, 0xe2, 0x01, 0x91, 0x83, 0xf1, 0x42, 0xba, 0xd2,
	0x6c, 0x8d, 0x6c, 0x5d, 0xd6, 0x6a, 0xc6, 0x5b, 0xf8, 0x14, 0x1e, 0x7c, 0x12, 0x8e, 0x1c, 0x3d,
	0xa9, 0x81, 0x17, 0x31, 0xed, 0x2a, 0x62, 0x32, 0x4f, 0xeb, 0xba, 0xcf, 0xf7, 0xd3, 0xef, 0x7e,
	0x9b, 0xd5, 0x26, 0x31, 0xa6, 0x10, 0x87, 0x28, 0xa5, 0x62, 0x0a, 0x9f, 0xba, 0x3e, 0x11, 0xa8,
	0x0b, 0x03, 0x12, 0x13, 0x4e, 0x39, 0x48, 0x52, 0x26, 0x98, 0x5d, 0x97, 0x0c, 0xd0, 0x0c, 0xd0,
	0x4c, 0xa3, 0x1e, 0xb0, 0x80, 0x29, 0x00, 0xca, 0x55, 0xce, 0x36, 0x8a, 0x7d, 0xdf, 0xd9, 0x9c,
	0x71, 0x31, 0xe3, 0x11, 0xe3, 0xd0, 0x47, 0x9c, 0xfc, 0x20, 0x8c, 0xc6, 0xf9, 0xf3, 0xf6, 0x4b,
	0xc9, 0xaa, 0x5d, 0xe5, 0x0d, 0x6e, 0x05, 0x12, 0xc4, 0x3e, 0xb7, 0x2a, 0x09, 0x4a, 0x51, 0xc4,
	0x1d, 0xb3, 0x65, 0x76, 0xaa, 0xa7, 0xfb, 0xa0, 0xa8, 0x11, 0xe8, 0x2b, 0xa6, 0x57, 0x9e, 0xbd,
	0x37, 0x8d, 0x81, 0x4e, 0xd8, 0x37, 0xd6, 0xb6, 0x40, 0xd9, 0x28, 0x45, 0x82, 0x8c, 0x26, 0x34,
	0xa2, 0x82, 0x3b, 0xff, 0x94, 0xe4, 0xa0, 0x58, 0x32, 0x44, 0xd9, 0x00, 0x09, 0x72, 0xad, 0x50,
	0xed, 0xda, 0x12, 0xeb, 0x9b, 0xf6, 0x85, 0xb5, 0x29, 0x95, 0x18, 0x25, 0xdc, 0x29, 0xb5, 0x4a,
	0x7f, 0x17, 0x1a, 0xa2, 0xcc, 0x43, 0x89, 0x96, 0x6c, 0x08, 0x75, 0xc7, 0xed, 0xd8, 0xaa, 0xc9,
	0x78, 0x92, 0x32, 0x4c, 0xc8, 0x98, 0x3b, 0x65, 0xa5, 0xd8, 0x03, 0xf9, 0x54, 0x80, 0x9c, 0xca,
	0xca, 0xe0, 0x31, 0x1a, 0xf7, 0x4e, 0x64, 0xfe, 0xf5, 0xa3, 0xd9, 0x09, 0xa8, 0x08, 0x1f, 0x7d,
	0x80, 0x59, 0x04, 0xf5, 0x08, 0xf3, 0xcb, 0x31, 0x1f, 0x3f, 0x40, 0x31, 0x4d, 0x08, 0x57, 0x01,
	0x3e, 0xa8, 0x0a, 0x94, 0xf5, 0xb5, 0xdf, 0xbe, 0xb3, 0x76, 0x31, 0x9b, 0x4c, 0x08, 0x16, 0x94,
	0xc5, 0x23, 0x92, 0x30, 0x1c, 0x72, 0xe7, 0xbf, 0x3a, 0xf4, 0xb0, 0xb8, 0xb7, 0xb7, 0xc2, 0x2f,
	0x25, 0xad, 0x5f, 0x60, 0x07, 0xff, 0xde, 0xe6, 0x3d, 0x6f, 0xb6, 0x70, 0xcd, 0xf9, 0xc2, 0x35,
	0x3f, 0x17, 0xae, 0xf9, 0xbc, 0x74, 0x8d, 0xf9, 0xd2, 0x35, 0xde, 0x96, 0xae, 0x71, 0x7f, 0xb4,
	0x56, 0x55, 0x1e, 0x81, 0x43, 0x44, 0x63, 0xb5, 0x82, 0xd9, 0xea, 0xef, 0x50, 0x8d, 0xfd, 0x8a,
	0xfa, 0xe8, 0x67, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0x31, 0xd3, 0xab, 0x8a, 0x02, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CollectionEpochs) > 0 {
		for iNdEx := len(m.CollectionEpochs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollectionEpochs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.TaxProceeds) > 0 {
		for iNdEx := len(m.TaxProceeds) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TaxProceeds[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.TaxCaps) > 0 {
		for iNdEx := len(m.TaxCaps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TaxCaps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size, err := m.TaxRateLimits.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TaxRateLimits.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.TaxCaps) > 0 {
		for _, e := range m.TaxCaps {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.TaxProceeds) > 0 {
		for _, e := range m.TaxProceeds {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CollectionEpochs) > 0 {
		for _, e := range m.CollectionEpochs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxRateLimits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TaxRateLimits.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxCaps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaxCaps = append(m.TaxCaps, TaxCap{})
			if err := m.TaxCaps[len(m.TaxCaps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaxProceeds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaxProceeds = append(m.TaxProceeds, types.Coin{})
			if err := m.TaxProceeds[len(m.TaxProceeds)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionEpochs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionEpochs = append(m.CollectionEpochs, CollectionEpoch{})
			if err := m.CollectionEpochs[len(m.CollectionEpochs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
