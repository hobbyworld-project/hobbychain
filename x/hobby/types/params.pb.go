// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: hobbychain/hobby/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	Exchange DenomExchange `protobuf:"bytes,1,opt,name=exchange,proto3" json:"denom_exchange"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_25ffc6096aadb924, []int{0}
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

func (m *Params) GetExchange() DenomExchange {
	if m != nil {
		return m.Exchange
	}
	return DenomExchange{}
}

type DenomExchange struct {
	// type of coin exchange from
	FromDenom string `protobuf:"bytes,1,opt,name=from_denom,json=fromDenom,proto3" json:"from_denom,omitempty"`
	// type of coin exchange to
	ToDenom string `protobuf:"bytes,2,opt,name=to_denom,json=toDenom,proto3" json:"to_denom,omitempty"`
	// exchange rate
	ExchangeRatio github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=exchange_ratio,json=exchangeRatio,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"exchange_ratio"`
	// allow list
	AllowList []string `protobuf:"bytes,4,rep,name=allow_list,json=allowList,proto3" json:"allow_list,omitempty"`
}

func (m *DenomExchange) Reset()         { *m = DenomExchange{} }
func (m *DenomExchange) String() string { return proto.CompactTextString(m) }
func (*DenomExchange) ProtoMessage()    {}
func (*DenomExchange) Descriptor() ([]byte, []int) {
	return fileDescriptor_25ffc6096aadb924, []int{1}
}
func (m *DenomExchange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomExchange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomExchange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomExchange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomExchange.Merge(m, src)
}
func (m *DenomExchange) XXX_Size() int {
	return m.Size()
}
func (m *DenomExchange) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomExchange.DiscardUnknown(m)
}

var xxx_messageInfo_DenomExchange proto.InternalMessageInfo

func (m *DenomExchange) GetFromDenom() string {
	if m != nil {
		return m.FromDenom
	}
	return ""
}

func (m *DenomExchange) GetToDenom() string {
	if m != nil {
		return m.ToDenom
	}
	return ""
}

func (m *DenomExchange) GetAllowList() []string {
	if m != nil {
		return m.AllowList
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "hobbychain.hobby.Params")
	proto.RegisterType((*DenomExchange)(nil), "hobbychain.hobby.DenomExchange")
}

func init() { proto.RegisterFile("hobbychain/hobby/params.proto", fileDescriptor_25ffc6096aadb924) }

var fileDescriptor_25ffc6096aadb924 = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcd, 0xc8, 0x4f, 0x4a,
	0xaa, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x07, 0x33, 0xf5, 0x0b, 0x12, 0x8b, 0x12, 0x73, 0x8b,
	0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x10, 0xd2, 0x7a, 0x60, 0xa6, 0x94, 0x48, 0x7a,
	0x7e, 0x7a, 0x3e, 0x58, 0x52, 0x1f, 0xc4, 0x82, 0xa8, 0x93, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd,
	0x2f, 0x8e, 0x87, 0x48, 0x40, 0x38, 0x10, 0x29, 0xa5, 0x64, 0x2e, 0xb6, 0x00, 0xb0, 0x91, 0x42,
	0xc1, 0x5c, 0x1c, 0xa9, 0x15, 0xc9, 0x19, 0x89, 0x79, 0xe9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xdc, 0x46, 0xf2, 0x7a, 0xe8, 0xe6, 0xeb, 0xb9, 0xa4, 0xe6, 0xe5, 0xe7, 0xba, 0x42, 0x95, 0x39,
	0x89, 0x9d, 0xb8, 0x27, 0xcf, 0xf0, 0xea, 0x9e, 0x3c, 0x5f, 0x0a, 0x48, 0x38, 0x1e, 0xa6, 0x3d,
	0x08, 0x6e, 0x90, 0x15, 0xcb, 0x8c, 0x05, 0xf2, 0x0c, 0x4a, 0x67, 0x19, 0xb9, 0x78, 0x51, 0x74,
	0x0a, 0xc9, 0x72, 0x71, 0xa5, 0x15, 0xe5, 0xe7, 0xc6, 0x83, 0x35, 0x82, 0xad, 0xe3, 0x0c, 0xe2,
	0x04, 0x89, 0x80, 0x95, 0x09, 0x49, 0x72, 0x71, 0x94, 0xe4, 0x43, 0x25, 0x99, 0xc0, 0x92, 0xec,
	0x25, 0xf9, 0x10, 0xa9, 0x64, 0x2e, 0x3e, 0x98, 0xe9, 0xf1, 0x45, 0x89, 0x25, 0x99, 0xf9, 0x12,
	0xcc, 0x20, 0x05, 0x4e, 0x36, 0x20, 0xb7, 0xdc, 0xba, 0x27, 0xaf, 0x96, 0x9e, 0x59, 0x92, 0x51,
	0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0x0b, 0xf5, 0x29, 0x94, 0xd2, 0x2d, 0x4e, 0xc9, 0xd6, 0x2f, 0xa9,
	0x2c, 0x48, 0x2d, 0xd6, 0x73, 0x49, 0x4d, 0xbe, 0xb4, 0x45, 0x97, 0x0b, 0x1a, 0x10, 0x2e, 0xa9,
	0xc9, 0x41, 0xbc, 0x70, 0xb7, 0x83, 0x8c, 0x04, 0x39, 0x2f, 0x31, 0x27, 0x27, 0xbf, 0x3c, 0x3e,
	0x27, 0xb3, 0xb8, 0x44, 0x82, 0x45, 0x81, 0x19, 0xe4, 0x3c, 0xb0, 0x88, 0x4f, 0x66, 0x71, 0x89,
	0x53, 0xc0, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1,
	0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x99, 0x21, 0xd9, 0x0e,
	0x0e, 0xb1, 0xf2, 0xfc, 0xa2, 0x9c, 0x14, 0xdd, 0x82, 0xa2, 0xfc, 0xac, 0xd4, 0xe4, 0x12, 0x7d,
	0xa4, 0xe8, 0xac, 0x80, 0x46, 0x28, 0xd8, 0x45, 0x49, 0x6c, 0xe0, 0xd8, 0x30, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xa5, 0xfe, 0x1f, 0x38, 0xf1, 0x01, 0x00, 0x00,
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
		size, err := m.Exchange.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *DenomExchange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomExchange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomExchange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AllowList) > 0 {
		for iNdEx := len(m.AllowList) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowList[iNdEx])
			copy(dAtA[i:], m.AllowList[iNdEx])
			i = encodeVarintParams(dAtA, i, uint64(len(m.AllowList[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.ExchangeRatio.Size()
		i -= size
		if _, err := m.ExchangeRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ToDenom) > 0 {
		i -= len(m.ToDenom)
		copy(dAtA[i:], m.ToDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.ToDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromDenom) > 0 {
		i -= len(m.FromDenom)
		copy(dAtA[i:], m.FromDenom)
		i = encodeVarintParams(dAtA, i, uint64(len(m.FromDenom)))
		i--
		dAtA[i] = 0xa
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
	l = m.Exchange.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func (m *DenomExchange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = len(m.ToDenom)
	if l > 0 {
		n += 1 + l + sovParams(uint64(l))
	}
	l = m.ExchangeRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	if len(m.AllowList) > 0 {
		for _, s := range m.AllowList {
			l = len(s)
			n += 1 + l + sovParams(uint64(l))
		}
	}
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
				return fmt.Errorf("proto: wrong wireType = %d for field Exchange", wireType)
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
			if err := m.Exchange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *DenomExchange) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: DenomExchange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomExchange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromDenom", wireType)
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
			m.FromDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToDenom", wireType)
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
			m.ToDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExchangeRatio", wireType)
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
			if err := m.ExchangeRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowList", wireType)
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
			m.AllowList = append(m.AllowList, string(dAtA[iNdEx:postIndex]))
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
