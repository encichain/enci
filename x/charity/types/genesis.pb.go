// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: charity/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
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
	Params        Params        `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	TaxRateLimits TaxRateLimits `protobuf:"bytes,2,opt,name=tax_rate_limits,json=taxRateLimits,proto3" json:"tax_rate_limits"`
	TaxCaps       []TaxCap      `protobuf:"bytes,3,rep,name=tax_caps,json=taxCaps,proto3" json:"tax_caps"`
	// Current amount of tax collected during current collection period
	CollectionPeriods []CollectionPeriod `protobuf:"bytes,4,rep,name=collection_periods,json=collectionPeriods,proto3" json:"collection_periods"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a8c5ac700ef3705, []int{0}
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

func (m *GenesisState) GetCollectionPeriods() []CollectionPeriod {
	if m != nil {
		return m.CollectionPeriods
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "user.encichain.charity.GenesisState")
}

func init() { proto.RegisterFile("charity/genesis.proto", fileDescriptor_7a8c5ac700ef3705) }

var fileDescriptor_7a8c5ac700ef3705 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0x33, 0x31,
	0x14, 0x85, 0x67, 0xda, 0xd2, 0xff, 0x67, 0x54, 0xc4, 0x41, 0x65, 0xe8, 0x22, 0x16, 0x41, 0xa8,
	0x9b, 0x09, 0xd5, 0xad, 0x20, 0xb4, 0x0b, 0x37, 0x2e, 0x4a, 0xdb, 0x95, 0x20, 0xe5, 0x4e, 0x0c,
	0xd3, 0x40, 0x67, 0x32, 0xe4, 0xde, 0xca, 0xf4, 0x2d, 0x7c, 0x25, 0x77, 0x5d, 0x76, 0xe9, 0x4a,
	0xa4, 0x7d, 0x11, 0x69, 0x1a, 0x8b, 0x88, 0x75, 0x95, 0xe4, 0xe4, 0x9c, 0xef, 0x5c, 0xb8, 0xc1,
	0x89, 0x18, 0x83, 0x51, 0x34, 0xe3, 0xa9, 0xcc, 0x25, 0x2a, 0x8c, 0x0b, 0xa3, 0x49, 0x87, 0xa7,
	0x53, 0x94, 0x26, 0x96, 0xb9, 0x50, 0x62, 0x0c, 0x2a, 0x8f, 0x9d, 0xab, 0x71, 0x9c, 0xea, 0x54,
	0x5b, 0x0b, 0x5f, 0xdf, 0x36, 0xee, 0xc6, 0x16, 0xe2, 0x4e, 0x27, 0x33, 0xa1, 0x31, 0xd3, 0xc8,
	0x13, 0x40, 0xc9, 0x9f, 0xdb, 0x89, 0x24, 0x68, 0x73, 0xa1, 0x55, 0xbe, 0xf9, 0x3f, 0x7f, 0xad,
	0x04, 0xfb, 0x77, 0x9b, 0xda, 0x01, 0x01, 0xc9, 0xf0, 0x26, 0xa8, 0x17, 0x60, 0x20, 0xc3, 0xc8,
	0x6f, 0xfa, 0xad, 0xbd, 0x2b, 0x16, 0xff, 0x3e, 0x46, 0xdc, 0xb3, 0xae, 0x4e, 0x6d, 0xfe, 0x7e,
	0xe6, 0xf5, 0x5d, 0x26, 0x1c, 0x04, 0x87, 0x04, 0xe5, 0xc8, 0x00, 0xc9, 0xd1, 0x44, 0x65, 0x8a,
	0x30, 0xaa, 0x58, 0xcc, 0xc5, 0x2e, 0xcc, 0x10, 0xca, 0x3e, 0x90, 0xbc, 0xb7, 0x66, 0x47, 0x3b,
	0xa0, 0xef, 0x62, 0x78, 0x1b, 0xfc, 0x5f, 0x43, 0x05, 0x14, 0x18, 0x55, 0x9b, 0xd5, 0xbf, 0x86,
	0x1a, 0x42, 0xd9, 0x85, 0xc2, 0x61, 0xfe, 0x91, 0x7d, 0x61, 0xf8, 0x18, 0x84, 0x42, 0x4f, 0x26,
	0x52, 0x90, 0xd2, 0xf9, 0xa8, 0x90, 0x46, 0xe9, 0x27, 0x8c, 0x6a, 0x16, 0xd5, 0xda, 0x85, 0xea,
	0x6e, 0x13, 0x3d, 0x1b, 0x70, 0xd0, 0x23, 0xf1, 0x43, 0xc7, 0x4e, 0x77, 0xbe, 0x64, 0xfe, 0x62,
	0xc9, 0xfc, 0x8f, 0x25, 0xf3, 0x5f, 0x56, 0xcc, 0x5b, 0xac, 0x98, 0xf7, 0xb6, 0x62, 0xde, 0xc3,
	0x65, 0xaa, 0x68, 0x3c, 0x4d, 0x62, 0xa1, 0x33, 0xbe, 0xae, 0xe1, 0xdb, 0x1a, 0x5e, 0x7e, 0x2d,
	0x8a, 0xd3, 0xac, 0x90, 0x98, 0xd4, 0xed, 0x3e, 0xae, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x71,
	0x59, 0x2b, 0x7e, 0x0d, 0x02, 0x00, 0x00,
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
	if len(m.CollectionPeriods) > 0 {
		for iNdEx := len(m.CollectionPeriods) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CollectionPeriods[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.CollectionPeriods) > 0 {
		for _, e := range m.CollectionPeriods {
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
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionPeriods", wireType)
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
			m.CollectionPeriods = append(m.CollectionPeriods, CollectionPeriod{})
			if err := m.CollectionPeriods[len(m.CollectionPeriods)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
