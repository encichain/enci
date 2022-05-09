// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enci/enciprice/v1beta1/enciprice.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type EnciUsd struct {
	// Price in USD per 1 ENCI (1_000_000uenci)
	Price       github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=price,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price" yaml:"price"`
	Blockheight int64                                  `protobuf:"varint,2,opt,name=blockheight,proto3" json:"blockheight,omitempty"`
}

func (m *EnciUsd) Reset()         { *m = EnciUsd{} }
func (m *EnciUsd) String() string { return proto.CompactTextString(m) }
func (*EnciUsd) ProtoMessage()    {}
func (*EnciUsd) Descriptor() ([]byte, []int) {
	return fileDescriptor_43d8b417ef9b0725, []int{0}
}
func (m *EnciUsd) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EnciUsd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EnciUsd.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EnciUsd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnciUsd.Merge(m, src)
}
func (m *EnciUsd) XXX_Size() int {
	return m.Size()
}
func (m *EnciUsd) XXX_DiscardUnknown() {
	xxx_messageInfo_EnciUsd.DiscardUnknown(m)
}

var xxx_messageInfo_EnciUsd proto.InternalMessageInfo

// Params defines the parameters for the enciprice module.
type Params struct {
	SlashEnabled bool `protobuf:"varint,1,opt,name=slash_enabled,json=slashEnabled,proto3" json:"slash_enabled,omitempty" yaml:"slash_enabled"`
	// The window of blocks where oracle misses are counted towards slashing
	SlashWindow   uint64 `protobuf:"varint,2,opt,name=slash_window,json=slashWindow,proto3" json:"slash_window,omitempty" yaml:"slash_window"`
	MissThreshold int64  `protobuf:"varint,3,opt,name=miss_threshold,json=missThreshold,proto3" json:"miss_threshold,omitempty" yaml:"miss_threshold"`
	// Fraction of stake to be slashed for exceeding miss threshold
	SlashFraction github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=slash_fraction,json=slashFraction,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"slash_fraction" yaml:"slash_fraction"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_43d8b417ef9b0725, []int{1}
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

func (m *Params) GetSlashEnabled() bool {
	if m != nil {
		return m.SlashEnabled
	}
	return false
}

func (m *Params) GetSlashWindow() uint64 {
	if m != nil {
		return m.SlashWindow
	}
	return 0
}

func (m *Params) GetMissThreshold() int64 {
	if m != nil {
		return m.MissThreshold
	}
	return 0
}

func init() {
	proto.RegisterType((*EnciUsd)(nil), "enci.enciprice.v1beta1.EnciUsd")
	proto.RegisterType((*Params)(nil), "enci.enciprice.v1beta1.Params")
}

func init() {
	proto.RegisterFile("enci/enciprice/v1beta1/enciprice.proto", fileDescriptor_43d8b417ef9b0725)
}

var fileDescriptor_43d8b417ef9b0725 = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x4d, 0xde, 0x8b, 0x4f, 0x9d, 0xd7, 0x76, 0x11, 0x5b, 0x8d, 0x2e, 0x32, 0x65, 0x16, 0xa5,
	0x20, 0x26, 0x14, 0x77, 0x01, 0x45, 0x82, 0xd5, 0xad, 0x84, 0x8a, 0xe0, 0xa6, 0x4c, 0x26, 0x63,
	0x32, 0x34, 0xc9, 0x94, 0x4c, 0xb4, 0xf6, 0x0b, 0x74, 0xe9, 0xd2, 0x65, 0xff, 0xc3, 0x1f, 0xe8,
	0xb2, 0x4b, 0x71, 0x11, 0xa4, 0xfd, 0x83, 0x7c, 0x81, 0x64, 0xa6, 0xa5, 0xe9, 0xd2, 0xcd, 0xcc,
	0x3d, 0xe7, 0xde, 0x7b, 0xce, 0xe5, 0x72, 0xc1, 0x88, 0xe6, 0x84, 0xb9, 0xcd, 0xb3, 0x2c, 0x18,
	0xa1, 0xee, 0x97, 0x49, 0x48, 0x4b, 0x3c, 0x39, 0x33, 0xce, 0xb2, 0xe0, 0x25, 0x37, 0x1f, 0x36,
	0x84, 0x73, 0x66, 0x8f, 0x75, 0x4f, 0xfa, 0x31, 0x8f, 0xb9, 0x2c, 0x71, 0x9b, 0x48, 0x55, 0xa3,
	0x6f, 0x3a, 0xb8, 0x3b, 0xcd, 0x09, 0x7b, 0x2f, 0x22, 0x73, 0x06, 0xee, 0xc8, 0x16, 0x4b, 0x1f,
	0xea, 0xe3, 0xfb, 0xfe, 0xcb, 0x6d, 0x05, 0xb5, 0x3f, 0x15, 0x1c, 0xc5, 0xac, 0x4c, 0x3e, 0x87,
	0x0e, 0xe1, 0x99, 0x4b, 0xb8, 0xc8, 0xb8, 0x38, 0x7e, 0xcf, 0x44, 0xb4, 0x70, 0xcb, 0xf5, 0x92,
	0x0a, 0xe7, 0x35, 0x25, 0x75, 0x05, 0x3b, 0x6b, 0x9c, 0xa5, 0x1e, 0x92, 0x22, 0x28, 0x50, 0x62,
	0xe6, 0x10, 0xdc, 0x86, 0x29, 0x27, 0x8b, 0x84, 0xb2, 0x38, 0x29, 0xad, 0xab, 0xa1, 0x3e, 0xbe,
	0x0e, 0xda, 0x94, 0x67, 0x7c, 0xdf, 0x40, 0x0d, 0xfd, 0xba, 0x02, 0x37, 0xef, 0x70, 0x81, 0x33,
	0x61, 0xbe, 0x00, 0x5d, 0x91, 0x62, 0x91, 0xcc, 0x69, 0x8e, 0xc3, 0x94, 0x46, 0x72, 0xa0, 0x7b,
	0xbe, 0x55, 0x57, 0xb0, 0xaf, 0x2c, 0x2e, 0xd2, 0x28, 0xe8, 0x48, 0x3c, 0x55, 0xd0, 0xf4, 0x80,
	0xc2, 0xf3, 0x15, 0xcb, 0x23, 0xbe, 0x92, 0x96, 0x86, 0xff, 0xa8, 0xae, 0xe0, 0x83, 0x76, 0xb7,
	0xca, 0xa2, 0xe0, 0x56, 0xc2, 0x0f, 0x12, 0x99, 0xaf, 0x40, 0x2f, 0x63, 0x42, 0xcc, 0xcb, 0xa4,
	0xa0, 0x22, 0xe1, 0x69, 0x64, 0x5d, 0x37, 0x03, 0xfb, 0x8f, 0xeb, 0x0a, 0x0e, 0x54, 0xf7, 0x65,
	0x1e, 0x05, 0xdd, 0x86, 0x98, 0x9d, 0xb0, 0x99, 0x83, 0x9e, 0xd2, 0xff, 0x54, 0x60, 0x52, 0x32,
	0x9e, 0x5b, 0x86, 0x5c, 0xe7, 0xdb, 0xff, 0x5e, 0xe7, 0xa0, 0x3d, 0xed, 0x49, 0x0d, 0x05, 0x6a,
	0x37, 0x6f, 0x8e, 0xd8, 0x33, 0x7e, 0x6e, 0xa0, 0xe6, 0x4f, 0xb7, 0x7b, 0x5b, 0xdf, 0xed, 0x6d,
	0xfd, 0xef, 0xde, 0xd6, 0x7f, 0x1c, 0x6c, 0x6d, 0x77, 0xb0, 0xb5, 0xdf, 0x07, 0x5b, 0xfb, 0xf8,
	0xb4, 0xe5, 0xd7, 0x5c, 0x05, 0x49, 0x30, 0xcb, 0x65, 0xe4, 0x7e, 0x6d, 0x9d, 0x93, 0x34, 0x0e,
	0x6f, 0xe4, 0x55, 0x3c, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x95, 0x4e, 0x83, 0x9b, 0x6d, 0x02,
	0x00, 0x00,
}

func (m *EnciUsd) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EnciUsd) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EnciUsd) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Blockheight != 0 {
		i = encodeVarintEnciprice(dAtA, i, uint64(m.Blockheight))
		i--
		dAtA[i] = 0x10
	}
	{
		size := m.Price.Size()
		i -= size
		if _, err := m.Price.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEnciprice(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	{
		size := m.SlashFraction.Size()
		i -= size
		if _, err := m.SlashFraction.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintEnciprice(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.MissThreshold != 0 {
		i = encodeVarintEnciprice(dAtA, i, uint64(m.MissThreshold))
		i--
		dAtA[i] = 0x18
	}
	if m.SlashWindow != 0 {
		i = encodeVarintEnciprice(dAtA, i, uint64(m.SlashWindow))
		i--
		dAtA[i] = 0x10
	}
	if m.SlashEnabled {
		i--
		if m.SlashEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEnciprice(dAtA []byte, offset int, v uint64) int {
	offset -= sovEnciprice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EnciUsd) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Price.Size()
	n += 1 + l + sovEnciprice(uint64(l))
	if m.Blockheight != 0 {
		n += 1 + sovEnciprice(uint64(m.Blockheight))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SlashEnabled {
		n += 2
	}
	if m.SlashWindow != 0 {
		n += 1 + sovEnciprice(uint64(m.SlashWindow))
	}
	if m.MissThreshold != 0 {
		n += 1 + sovEnciprice(uint64(m.MissThreshold))
	}
	l = m.SlashFraction.Size()
	n += 1 + l + sovEnciprice(uint64(l))
	return n
}

func sovEnciprice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEnciprice(x uint64) (n int) {
	return sovEnciprice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EnciUsd) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnciprice
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
			return fmt.Errorf("proto: EnciUsd: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EnciUsd: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnciprice
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
				return ErrInvalidLengthEnciprice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnciprice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Blockheight", wireType)
			}
			m.Blockheight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnciprice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Blockheight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEnciprice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEnciprice
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnciprice
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnciprice
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
			m.SlashEnabled = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashWindow", wireType)
			}
			m.SlashWindow = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnciprice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SlashWindow |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissThreshold", wireType)
			}
			m.MissThreshold = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnciprice
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissThreshold |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlashFraction", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnciprice
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
				return ErrInvalidLengthEnciprice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEnciprice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SlashFraction.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnciprice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEnciprice
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
func skipEnciprice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnciprice
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
					return 0, ErrIntOverflowEnciprice
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
					return 0, ErrIntOverflowEnciprice
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
				return 0, ErrInvalidLengthEnciprice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEnciprice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEnciprice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEnciprice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnciprice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEnciprice = fmt.Errorf("proto: unexpected end of group")
)
