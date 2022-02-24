// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enci/oracle/v1beta1/params.proto

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

// Params represents the parameters for the module.
type Params struct {
	ClaimParams map[string]ClaimParams `protobuf:"bytes,1,rep,name=claim_params,json=claimParams,proto3" json:"claim_params" yaml:"claim_params" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_754d571249ec7e6b, []int{0}
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

func (m *Params) GetClaimParams() map[string]ClaimParams {
	if m != nil {
		return m.ClaimParams
	}
	return nil
}

// ClaimParams is the parameters set for each oracle claim type
type ClaimParams struct {
	VotePeriod    uint64                                 `protobuf:"varint,1,opt,name=vote_period,json=votePeriod,proto3" json:"vote_period,omitempty" yaml:"vote_period"`
	ClaimType     string                                 `protobuf:"bytes,2,opt,name=claim_type,json=claimType,proto3" json:"claim_type,omitempty" yaml:"claim_type"`
	Prevote       bool                                   `protobuf:"varint,3,opt,name=prevote,proto3" json:"prevote,omitempty" yaml:"prevote"`
	VoteThreshold github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=vote_threshold,json=voteThreshold,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"vote_threshold" yaml:"vote_threshold"`
}

func (m *ClaimParams) Reset()         { *m = ClaimParams{} }
func (m *ClaimParams) String() string { return proto.CompactTextString(m) }
func (*ClaimParams) ProtoMessage()    {}
func (*ClaimParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_754d571249ec7e6b, []int{1}
}
func (m *ClaimParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimParams.Merge(m, src)
}
func (m *ClaimParams) XXX_Size() int {
	return m.Size()
}
func (m *ClaimParams) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimParams.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimParams proto.InternalMessageInfo

func (m *ClaimParams) GetVotePeriod() uint64 {
	if m != nil {
		return m.VotePeriod
	}
	return 0
}

func (m *ClaimParams) GetClaimType() string {
	if m != nil {
		return m.ClaimType
	}
	return ""
}

func (m *ClaimParams) GetPrevote() bool {
	if m != nil {
		return m.Prevote
	}
	return false
}

func init() {
	proto.RegisterType((*Params)(nil), "enci.oracle.v1beta1.Params")
	proto.RegisterMapType((map[string]ClaimParams)(nil), "enci.oracle.v1beta1.Params.ClaimParamsEntry")
	proto.RegisterType((*ClaimParams)(nil), "enci.oracle.v1beta1.ClaimParams")
}

func init() { proto.RegisterFile("enci/oracle/v1beta1/params.proto", fileDescriptor_754d571249ec7e6b) }

var fileDescriptor_754d571249ec7e6b = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcf, 0x6e, 0xd3, 0x30,
	0x18, 0x8f, 0xdb, 0x31, 0xa8, 0x33, 0xa6, 0xe1, 0x31, 0x14, 0x0d, 0x29, 0x89, 0x7c, 0x40, 0x39,
	0x0c, 0x47, 0x1b, 0x08, 0xd0, 0x8e, 0xe1, 0xdf, 0x75, 0x8a, 0x76, 0xe2, 0x52, 0x5c, 0xd7, 0x6a,
	0xa2, 0x26, 0x71, 0x94, 0xa4, 0x15, 0x79, 0x0b, 0xce, 0x3c, 0x01, 0x8f, 0xd2, 0x63, 0x8f, 0x15,
	0x87, 0x08, 0xa5, 0x6f, 0x90, 0x27, 0x40, 0xb6, 0x5b, 0x88, 0x50, 0x4f, 0xf9, 0xbe, 0xfc, 0xfe,
	0x7d, 0x9f, 0xf5, 0x41, 0x97, 0x67, 0x2c, 0xf6, 0x45, 0x41, 0x59, 0xc2, 0xfd, 0xe5, 0xf5, 0x84,
	0x57, 0xf4, 0xda, 0xcf, 0x69, 0x41, 0xd3, 0x92, 0xe4, 0x85, 0xa8, 0x04, 0x3a, 0x97, 0x0c, 0xa2,
	0x19, 0x64, 0xc7, 0xb8, 0x7c, 0x3a, 0x13, 0x33, 0xa1, 0x70, 0x5f, 0x56, 0x9a, 0x8a, 0x37, 0x00,
	0x1e, 0xdf, 0x29, 0x2d, 0x4a, 0xe0, 0x09, 0x4b, 0x68, 0x9c, 0x8e, 0xb5, 0x97, 0x05, 0xdc, 0xa1,
	0x67, 0xde, 0x5c, 0x91, 0x03, 0x66, 0x44, 0x4b, 0xc8, 0x7b, 0xc9, 0xd7, 0xf5, 0xc7, 0xac, 0x2a,
	0xea, 0xe0, 0xf9, 0xaa, 0x71, 0x8c, 0xae, 0x71, 0xce, 0x6b, 0x9a, 0x26, 0xb7, 0xb8, 0xef, 0x87,
	0x43, 0x93, 0xfd, 0xa3, 0x5f, 0x7e, 0x85, 0x67, 0xff, 0xab, 0xd1, 0x19, 0x1c, 0xce, 0x79, 0x6d,
	0x01, 0x17, 0x78, 0xa3, 0x50, 0x96, 0xe8, 0x0d, 0x7c, 0xb0, 0xa4, 0xc9, 0x82, 0x5b, 0x03, 0x17,
	0x78, 0xe6, 0x8d, 0x7b, 0x70, 0x98, 0x9e, 0x4f, 0xa8, 0xe9, 0xb7, 0x83, 0x77, 0x00, 0xff, 0x18,
	0x40, 0xb3, 0x07, 0xa1, 0xb7, 0xd0, 0x5c, 0x8a, 0x8a, 0x8f, 0x73, 0x5e, 0xc4, 0x62, 0xaa, 0x52,
	0x8e, 0x82, 0x67, 0x5d, 0xe3, 0x20, 0x3d, 0x6c, 0x0f, 0xc4, 0x21, 0x94, 0xdd, 0x9d, 0x6a, 0xd0,
	0x6b, 0x08, 0xf5, 0x22, 0x55, 0x9d, 0xeb, 0x49, 0x46, 0xc1, 0x45, 0xd7, 0x38, 0x4f, 0xfa, 0x4b,
	0x4a, 0x0c, 0x87, 0x23, 0xd5, 0xdc, 0xd7, 0x39, 0x47, 0x57, 0xf0, 0x61, 0x5e, 0x70, 0x69, 0x63,
	0x0d, 0x5d, 0xe0, 0x3d, 0x0a, 0x50, 0xd7, 0x38, 0xa7, 0x5a, 0xb2, 0x03, 0x70, 0xb8, 0xa7, 0xa0,
	0x0c, 0x9e, 0xaa, 0xfc, 0x2a, 0x2a, 0x78, 0x19, 0x89, 0x64, 0x6a, 0x1d, 0xb9, 0xc0, 0x3b, 0x09,
	0x3e, 0xcb, 0x07, 0xfd, 0xd5, 0x38, 0x2f, 0x66, 0x71, 0x15, 0x2d, 0x26, 0x84, 0x89, 0xd4, 0x67,
	0xa2, 0x4c, 0x45, 0xb9, 0xfb, 0xbc, 0x2c, 0xa7, 0x73, 0x5f, 0x86, 0x97, 0xe4, 0x03, 0x67, 0x5d,
	0xe3, 0x5c, 0xf4, 0xb6, 0xf9, 0xeb, 0x86, 0xc3, 0xc7, 0xf2, 0xc7, 0xfd, 0xbe, 0x0f, 0x3e, 0xfd,
	0x6c, 0x6d, 0xb0, 0x6a, 0x6d, 0xb0, 0x6e, 0x6d, 0xf0, 0xbb, 0xb5, 0xc1, 0xf7, 0xad, 0x6d, 0xac,
	0xb7, 0xb6, 0xb1, 0xd9, 0xda, 0xc6, 0x17, 0xaf, 0x97, 0x26, 0x5f, 0x9c, 0x45, 0x34, 0xce, 0x54,
	0xe5, 0x7f, 0xdb, 0x5f, 0x9e, 0xca, 0x9c, 0x1c, 0xab, 0x33, 0x7a, 0xf5, 0x27, 0x00, 0x00, 0xff,
	0xff, 0x52, 0x27, 0x52, 0x75, 0x95, 0x02, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if len(this.ClaimParams) != len(that1.ClaimParams) {
		return false
	}
	for i := range this.ClaimParams {
		a := this.ClaimParams[i]
		b := that1.ClaimParams[i]
		if !(&a).Equal(&b) {
			return false
		}
	}
	return true
}
func (this *ClaimParams) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ClaimParams)
	if !ok {
		that2, ok := that.(ClaimParams)
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
	if this.VotePeriod != that1.VotePeriod {
		return false
	}
	if this.ClaimType != that1.ClaimType {
		return false
	}
	if this.Prevote != that1.Prevote {
		return false
	}
	if !this.VoteThreshold.Equal(that1.VoteThreshold) {
		return false
	}
	return true
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
	if len(m.ClaimParams) > 0 {
		for k := range m.ClaimParams {
			v := m.ClaimParams[k]
			baseI := i
			{
				size, err := (&v).MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintParams(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintParams(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ClaimParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.VoteThreshold.Size()
		i -= size
		if _, err := m.VoteThreshold.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Prevote {
		i--
		if m.Prevote {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.ClaimType) > 0 {
		i -= len(m.ClaimType)
		copy(dAtA[i:], m.ClaimType)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ClaimType)))
		i--
		dAtA[i] = 0x12
	}
	if m.VotePeriod != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.VotePeriod))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
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
	if len(m.ClaimParams) > 0 {
		for k, v := range m.ClaimParams {
			_ = k
			_ = v
			l = v.Size()
			mapEntrySize := 1 + len(k) + sovParams(uint64(len(k))) + 1 + l + sovParams(uint64(l))
			n += mapEntrySize + 1 + sovParams(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *ClaimParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VotePeriod != 0 {
		n += 1 + sovParams(uint64(m.VotePeriod))
	}
	l = len(m.ClaimType)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	if m.Prevote {
		n += 2
	}
	l = m.VoteThreshold.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClaimParams == nil {
				m.ClaimParams = make(map[string]ClaimParams)
			}
			var mapkey string
			mapvalue := &ClaimParams{}
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowParams
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthParams
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthParams
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowParams
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthParams
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthParams
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &ClaimParams{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipParams(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthParams
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.ClaimParams[mapkey] = *mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *ClaimParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: ClaimParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotePeriod", wireType)
			}
			m.VotePeriod = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VotePeriod |= uint64(b&0x7F) << shift
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
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prevote", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
			m.Prevote = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VoteThreshold", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VoteThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)