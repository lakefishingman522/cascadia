// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cascadia/inflation/v1/inflation_params.proto

package types

import (
	fmt "fmt"
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

type InflationControlParams struct {
	// multiplier value
	Lambda github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=lambda,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"lambda"`
	// w360 defines the weight about avg Token Price of 360 days
	W360 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=w360,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"w360"`
	// w180 defines the weight about avg Token Price of 180 days
	W180 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=w180,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"w180"`
	// w180 defines the weight about avg Token Price of 90 days
	W90 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=w90,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"w90"`
	// w180 defines the weight about avg Token Price of 30 days
	W30 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=w30,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"w30"`
	// w14 defines the weight about avg Token Price of 14 days
	W14 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=w14,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"w14"`
	// w7 defines the weight about avg Token Price of a week
	W7 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=w7,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"w7"`
	// w1 defines the weight about avg Token Price of a day
	W1 github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=w1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"w1"`
}

func (m *InflationControlParams) Reset()         { *m = InflationControlParams{} }
func (m *InflationControlParams) String() string { return proto.CompactTextString(m) }
func (*InflationControlParams) ProtoMessage()    {}
func (*InflationControlParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_69c84a992623bbb9, []int{0}
}
func (m *InflationControlParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationControlParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationControlParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationControlParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationControlParams.Merge(m, src)
}
func (m *InflationControlParams) XXX_Size() int {
	return m.Size()
}
func (m *InflationControlParams) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationControlParams.DiscardUnknown(m)
}

var xxx_messageInfo_InflationControlParams proto.InternalMessageInfo

func init() {
	proto.RegisterType((*InflationControlParams)(nil), "cascadia.inflation.v1.InflationControlParams")
}

func init() {
	proto.RegisterFile("cascadia/inflation/v1/inflation_params.proto", fileDescriptor_69c84a992623bbb9)
}

var fileDescriptor_69c84a992623bbb9 = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x4e, 0x2c, 0x4e,
	0x4e, 0x4c, 0xc9, 0x4c, 0xd4, 0xcf, 0xcc, 0x4b, 0xcb, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f,
	0x33, 0x44, 0x70, 0xe2, 0x0b, 0x12, 0x8b, 0x12, 0x73, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2,
	0x85, 0x44, 0x61, 0xaa, 0xf5, 0xe0, 0x0a, 0xf4, 0xca, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3,
	0xc1, 0x2a, 0xf4, 0x41, 0x2c, 0x88, 0x62, 0xa5, 0x75, 0x2c, 0x5c, 0x62, 0x9e, 0x30, 0x65, 0xce,
	0xf9, 0x79, 0x25, 0x45, 0xf9, 0x39, 0x01, 0x60, 0xd3, 0x84, 0xdc, 0xb8, 0xd8, 0x72, 0x12, 0x73,
	0x93, 0x52, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x9d, 0xf4, 0x4e, 0xdc, 0x93, 0x67, 0xb8,
	0x75, 0x4f, 0x5e, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x39,
	0xbf, 0x38, 0x37, 0xbf, 0x18, 0x4a, 0xe9, 0x16, 0xa7, 0x64, 0xeb, 0x97, 0x54, 0x16, 0xa4, 0x16,
	0xeb, 0xb9, 0xa4, 0x26, 0x07, 0x41, 0x75, 0x0b, 0x39, 0x71, 0xb1, 0x94, 0x1b, 0x9b, 0x19, 0x48,
	0x30, 0x91, 0x65, 0x0a, 0x58, 0x2f, 0xd8, 0x0c, 0x43, 0x0b, 0x03, 0x09, 0x66, 0x32, 0xcd, 0x30,
	0xb4, 0x30, 0x10, 0x72, 0xe0, 0x62, 0x2e, 0xb7, 0x34, 0x90, 0x60, 0x21, 0xcb, 0x08, 0x90, 0x56,
	0xb0, 0x09, 0xc6, 0x06, 0x12, 0xac, 0x64, 0x9a, 0x60, 0x0c, 0x31, 0xc1, 0xd0, 0x44, 0x82, 0x8d,
	0x4c, 0x13, 0x0c, 0x4d, 0x84, 0xec, 0xb8, 0x98, 0xca, 0xcd, 0x25, 0xd8, 0xc9, 0x32, 0x80, 0xa9,
	0xdc, 0x1c, 0xac, 0xdf, 0x50, 0x82, 0x83, 0x4c, 0xfd, 0x86, 0x4e, 0x41, 0x27, 0x1e, 0xc9, 0x31,
	0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb,
	0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x81, 0x6c, 0x0a, 0x34, 0x09, 0xa6, 0xe5, 0x97, 0xe6,
	0xa5, 0x40, 0x12, 0x2c, 0x3c, 0x0d, 0x57, 0x20, 0xa5, 0x62, 0xb0, 0xd9, 0x49, 0x6c, 0xe0, 0xb4,
	0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x83, 0x90, 0xeb, 0x8a, 0xe8, 0x02, 0x00, 0x00,
}

func (m *InflationControlParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationControlParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationControlParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.W1.Size()
		i -= size
		if _, err := m.W1.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflationParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.W7.Size()
		i -= size
		if _, err := m.W7.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflationParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.W14.Size()
		i -= size
		if _, err := m.W14.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflationParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.W30.Size()
		i -= size
		if _, err := m.W30.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflationParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.W90.Size()
		i -= size
		if _, err := m.W90.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflationParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.W180.Size()
		i -= size
		if _, err := m.W180.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflationParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.W360.Size()
		i -= size
		if _, err := m.W360.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflationParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Lambda.Size()
		i -= size
		if _, err := m.Lambda.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflationParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintInflationParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovInflationParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InflationControlParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Lambda.Size()
	n += 1 + l + sovInflationParams(uint64(l))
	l = m.W360.Size()
	n += 1 + l + sovInflationParams(uint64(l))
	l = m.W180.Size()
	n += 1 + l + sovInflationParams(uint64(l))
	l = m.W90.Size()
	n += 1 + l + sovInflationParams(uint64(l))
	l = m.W30.Size()
	n += 1 + l + sovInflationParams(uint64(l))
	l = m.W14.Size()
	n += 1 + l + sovInflationParams(uint64(l))
	l = m.W7.Size()
	n += 1 + l + sovInflationParams(uint64(l))
	l = m.W1.Size()
	n += 1 + l + sovInflationParams(uint64(l))
	return n
}

func sovInflationParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInflationParams(x uint64) (n int) {
	return sovInflationParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InflationControlParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInflationParams
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
			return fmt.Errorf("proto: InflationControlParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationControlParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lambda", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflationParams
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
				return ErrInvalidLengthInflationParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflationParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Lambda.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field W360", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflationParams
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
				return ErrInvalidLengthInflationParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflationParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.W360.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field W180", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflationParams
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
				return ErrInvalidLengthInflationParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflationParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.W180.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field W90", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflationParams
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
				return ErrInvalidLengthInflationParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflationParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.W90.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field W30", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflationParams
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
				return ErrInvalidLengthInflationParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflationParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.W30.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field W14", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflationParams
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
				return ErrInvalidLengthInflationParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflationParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.W14.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field W7", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflationParams
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
				return ErrInvalidLengthInflationParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflationParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.W7.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field W1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflationParams
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
				return ErrInvalidLengthInflationParams
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflationParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.W1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInflationParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInflationParams
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
func skipInflationParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInflationParams
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
					return 0, ErrIntOverflowInflationParams
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
					return 0, ErrIntOverflowInflationParams
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
				return 0, ErrInvalidLengthInflationParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInflationParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInflationParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInflationParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInflationParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInflationParams = fmt.Errorf("proto: unexpected end of group")
)
