// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: charity/event.proto

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

// EventPayout defines the Payout Event
type EventPayout struct {
	Period      uint64                                   `protobuf:"varint,1,opt,name=period,proto3" json:"period,omitempty"`
	Payouts     []Payout                                 `protobuf:"bytes,2,rep,name=payouts,proto3" json:"payouts"`
	BurnedCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=burned_coins,json=burnedCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"burned_coins"`
}

func (m *EventPayout) Reset()         { *m = EventPayout{} }
func (m *EventPayout) String() string { return proto.CompactTextString(m) }
func (*EventPayout) ProtoMessage()    {}
func (*EventPayout) Descriptor() ([]byte, []int) {
	return fileDescriptor_696184d6db66912c, []int{0}
}
func (m *EventPayout) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventPayout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventPayout.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventPayout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventPayout.Merge(m, src)
}
func (m *EventPayout) XXX_Size() int {
	return m.Size()
}
func (m *EventPayout) XXX_DiscardUnknown() {
	xxx_messageInfo_EventPayout.DiscardUnknown(m)
}

var xxx_messageInfo_EventPayout proto.InternalMessageInfo

func (m *EventPayout) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *EventPayout) GetPayouts() []Payout {
	if m != nil {
		return m.Payouts
	}
	return nil
}

func (m *EventPayout) GetBurnedCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.BurnedCoins
	}
	return nil
}

// EventFailedPayouts defines a Payout failure Event
type EventFailedPayouts struct {
	Period uint64   `protobuf:"varint,1,opt,name=period,proto3" json:"period,omitempty"`
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (m *EventFailedPayouts) Reset()         { *m = EventFailedPayouts{} }
func (m *EventFailedPayouts) String() string { return proto.CompactTextString(m) }
func (*EventFailedPayouts) ProtoMessage()    {}
func (*EventFailedPayouts) Descriptor() ([]byte, []int) {
	return fileDescriptor_696184d6db66912c, []int{1}
}
func (m *EventFailedPayouts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventFailedPayouts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventFailedPayouts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventFailedPayouts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventFailedPayouts.Merge(m, src)
}
func (m *EventFailedPayouts) XXX_Size() int {
	return m.Size()
}
func (m *EventFailedPayouts) XXX_DiscardUnknown() {
	xxx_messageInfo_EventFailedPayouts.DiscardUnknown(m)
}

var xxx_messageInfo_EventFailedPayouts proto.InternalMessageInfo

func (m *EventFailedPayouts) GetPeriod() uint64 {
	if m != nil {
		return m.Period
	}
	return 0
}

func (m *EventFailedPayouts) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*EventPayout)(nil), "enci.charity.EventPayout")
	proto.RegisterType((*EventFailedPayouts)(nil), "enci.charity.EventFailedPayouts")
}

func init() { proto.RegisterFile("charity/event.proto", fileDescriptor_696184d6db66912c) }

var fileDescriptor_696184d6db66912c = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcf, 0x4e, 0x32, 0x31,
	0x14, 0xc5, 0x67, 0x3e, 0x08, 0x5f, 0x2c, 0xac, 0x46, 0x24, 0xc8, 0xa2, 0x10, 0x56, 0xb8, 0xb0,
	0x15, 0xf5, 0x09, 0x40, 0x5d, 0x93, 0x59, 0xba, 0x31, 0x33, 0x9d, 0x06, 0x1a, 0xa5, 0x77, 0xd2,
	0x16, 0x22, 0x6f, 0xe1, 0x73, 0xf8, 0x24, 0x24, 0x6e, 0x58, 0xba, 0x52, 0x03, 0x2f, 0x62, 0xfa,
	0x67, 0x12, 0x36, 0xae, 0xee, 0xbd, 0x73, 0xcf, 0xc9, 0xfd, 0x9d, 0x29, 0x3a, 0x65, 0x8b, 0x4c,
	0x09, 0xb3, 0xa1, 0x7c, 0xcd, 0xa5, 0x21, 0xa5, 0x02, 0x03, 0x49, 0x8b, 0x4b, 0x26, 0x48, 0xd8,
	0xf4, 0xce, 0x2a, 0x49, 0xa8, 0x5e, 0xd4, 0x6b, 0xcf, 0x61, 0x0e, 0xae, 0xa5, 0xb6, 0x0b, 0x5f,
	0x31, 0x03, 0xbd, 0x04, 0x4d, 0xf3, 0x4c, 0x73, 0xba, 0x1e, 0xe7, 0xdc, 0x64, 0x63, 0xca, 0x40,
	0x48, 0xbf, 0x1f, 0x7e, 0xc4, 0xa8, 0x79, 0x6f, 0x4f, 0xcd, 0xb2, 0x0d, 0xac, 0x4c, 0xd2, 0x41,
	0x8d, 0x92, 0x2b, 0x01, 0x45, 0x37, 0x1e, 0xc4, 0xa3, 0x7a, 0x1a, 0xa6, 0xe4, 0x16, 0xfd, 0x2f,
	0x9d, 0x42, 0x77, 0xff, 0x0d, 0x6a, 0xa3, 0xe6, 0x75, 0x9b, 0x1c, 0x43, 0x11, 0x6f, 0x9f, 0xd4,
	0xb7, 0x5f, 0xfd, 0x28, 0xad, 0xa4, 0x89, 0x44, 0xad, 0x7c, 0xa5, 0x24, 0x2f, 0x9e, 0xec, 0x49,
	0xdd, 0xad, 0x39, 0xeb, 0x39, 0xf1, 0x50, 0xc4, 0x42, 0x91, 0x00, 0x45, 0xa6, 0x20, 0xe4, 0xe4,
	0xca, 0xfa, 0xdf, 0xbf, 0xfb, 0xa3, 0xb9, 0x30, 0x8b, 0x55, 0x4e, 0x18, 0x2c, 0x69, 0x48, 0xe0,
	0xcb, 0xa5, 0x2e, 0x9e, 0xa9, 0xd9, 0x94, 0x5c, 0x3b, 0x83, 0x4e, 0x9b, 0xfe, 0x80, 0x1b, 0x86,
	0x77, 0x28, 0x71, 0x61, 0x1e, 0x32, 0xf1, 0xc2, 0x8b, 0x59, 0xa0, 0xf8, 0x2b, 0x53, 0x07, 0x35,
	0xb8, 0x52, 0xa0, 0x7c, 0xa4, 0x93, 0x34, 0x4c, 0x93, 0xe9, 0x76, 0x8f, 0xe3, 0xdd, 0x1e, 0xc7,
	0x3f, 0x7b, 0x1c, 0xbf, 0x1d, 0x70, 0xb4, 0x3b, 0xe0, 0xe8, 0xf3, 0x80, 0xa3, 0xc7, 0x8b, 0x23,
	0x2c, 0x1b, 0x9f, 0x2d, 0x32, 0x21, 0x5d, 0x47, 0x5f, 0xab, 0xe7, 0xf0, 0x74, 0x79, 0xc3, 0xfd,
	0xdf, 0x9b, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x08, 0x7e, 0x5c, 0xd1, 0x01, 0x00, 0x00,
}

func (m *EventPayout) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventPayout) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventPayout) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BurnedCoins) > 0 {
		for iNdEx := len(m.BurnedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BurnedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Payouts) > 0 {
		for iNdEx := len(m.Payouts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Payouts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Period != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *EventFailedPayouts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventFailedPayouts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventFailedPayouts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Errors) > 0 {
		for iNdEx := len(m.Errors) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Errors[iNdEx])
			copy(dAtA[i:], m.Errors[iNdEx])
			i = encodeVarintEvent(dAtA, i, uint64(len(m.Errors[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Period != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Period))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventPayout) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Period != 0 {
		n += 1 + sovEvent(uint64(m.Period))
	}
	if len(m.Payouts) > 0 {
		for _, e := range m.Payouts {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	if len(m.BurnedCoins) > 0 {
		for _, e := range m.BurnedCoins {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func (m *EventFailedPayouts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Period != 0 {
		n += 1 + sovEvent(uint64(m.Period))
	}
	if len(m.Errors) > 0 {
		for _, s := range m.Errors {
			l = len(s)
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventPayout) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventPayout: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventPayout: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payouts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payouts = append(m.Payouts, Payout{})
			if err := m.Payouts[len(m.Payouts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BurnedCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BurnedCoins = append(m.BurnedCoins, types.Coin{})
			if err := m.BurnedCoins[len(m.BurnedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventFailedPayouts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventFailedPayouts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventFailedPayouts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Period", wireType)
			}
			m.Period = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Period |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Errors", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Errors = append(m.Errors, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
