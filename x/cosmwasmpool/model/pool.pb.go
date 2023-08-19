// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlins/cosmwasmpool/v1beta1/model/pool.proto

package model

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
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

type CosmWasmPool struct {
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`
	PoolId          uint64 `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	CodeId          uint64 `protobuf:"varint,3,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty"`
	InstantiateMsg  []byte `protobuf:"bytes,4,opt,name=instantiate_msg,json=instantiateMsg,proto3" json:"instantiate_msg,omitempty" yaml:"instantiate_msg"`
}

func (m *CosmWasmPool) Reset()      { *m = CosmWasmPool{} }
func (*CosmWasmPool) ProtoMessage() {}
func (*CosmWasmPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0cb64564a744af1, []int{0}
}
func (m *CosmWasmPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CosmWasmPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CosmWasmPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CosmWasmPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CosmWasmPool.Merge(m, src)
}
func (m *CosmWasmPool) XXX_Size() int {
	return m.Size()
}
func (m *CosmWasmPool) XXX_DiscardUnknown() {
	xxx_messageInfo_CosmWasmPool.DiscardUnknown(m)
}

var xxx_messageInfo_CosmWasmPool proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CosmWasmPool)(nil), "merlins.cosmwasmpool.v1beta1.CosmWasmPool")
}

func init() {
	proto.RegisterFile("merlins/cosmwasmpool/v1beta1/model/pool.proto", fileDescriptor_a0cb64564a744af1)
}

var fileDescriptor_a0cb64564a744af1 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x3f, 0x4f, 0xc2, 0x40,
	0x18, 0xc6, 0x7b, 0x8a, 0x18, 0x1b, 0x22, 0xda, 0x18, 0x41, 0x34, 0x2d, 0xe9, 0xc4, 0x42, 0x2f,
	0xc4, 0x41, 0xc3, 0x26, 0x24, 0x26, 0x0c, 0x26, 0xa6, 0x8b, 0x89, 0x4b, 0x73, 0xfd, 0x63, 0x6d,
	0xd2, 0xe3, 0x25, 0xbc, 0x07, 0xea, 0x37, 0x70, 0x74, 0x74, 0xe4, 0x43, 0xf8, 0x21, 0x8c, 0x13,
	0xa3, 0x13, 0x21, 0xf0, 0x0d, 0xf8, 0x04, 0xe6, 0x7a, 0x25, 0x41, 0xb6, 0x3e, 0xbf, 0xe7, 0xd7,
	0xdc, 0xdd, 0xfb, 0xea, 0x4d, 0x40, 0x0e, 0x98, 0x20, 0x0d, 0x00, 0xf9, 0x0b, 0x43, 0x3e, 0x00,
	0x48, 0xe9, 0xb8, 0xe5, 0x47, 0x82, 0xb5, 0x28, 0x87, 0x30, 0x4a, 0xa9, 0x44, 0xce, 0x60, 0x08,
	0x02, 0x8c, 0x8b, 0x5c, 0x77, 0x36, 0x75, 0x27, 0xd7, 0x6b, 0x67, 0x41, 0x56, 0x7b, 0x99, 0x4b,
	0x55, 0x50, 0x3f, 0xd6, 0x4e, 0x62, 0x88, 0x41, 0x71, 0xf9, 0x95, 0x53, 0x2b, 0x06, 0x88, 0xd3,
	0x88, 0x66, 0xc9, 0x1f, 0x3d, 0x51, 0x91, 0xf0, 0x08, 0x05, 0xe3, 0x03, 0x25, 0xd8, 0x73, 0xa2,
	0x97, 0xba, 0x80, 0xfc, 0x81, 0x21, 0xbf, 0x07, 0x48, 0x8d, 0x5b, 0xfd, 0x28, 0x80, 0xbe, 0x18,
	0xb2, 0x40, 0x78, 0x2c, 0x0c, 0x87, 0x11, 0x62, 0x95, 0xd4, 0x49, 0xe3, 0xa0, 0x73, 0xbe, 0x9a,
	0x59, 0x95, 0x37, 0xc6, 0xd3, 0xb6, 0xbd, 0x6d, 0xd8, 0x6e, 0x79, 0x8d, 0x6e, 0x14, 0x31, 0x2a,
	0xfa, 0xbe, 0xbc, 0xba, 0x97, 0x84, 0xd5, 0x9d, 0x3a, 0x69, 0x14, 0xdc, 0xa2, 0x8c, 0xbd, 0x50,
	0x16, 0x01, 0x84, 0x91, 0x2c, 0x76, 0x55, 0x21, 0x63, 0x2f, 0x34, 0xba, 0x7a, 0x39, 0xe9, 0xa3,
	0x60, 0x7d, 0x91, 0x30, 0x11, 0x79, 0x1c, 0xe3, 0x6a, 0xa1, 0x4e, 0x1a, 0xa5, 0x4e, 0x6d, 0x35,
	0xb3, 0x4e, 0xd5, 0xc1, 0x5b, 0x82, 0xed, 0x1e, 0x6e, 0x90, 0x3b, 0x8c, 0xdb, 0xc7, 0xef, 0x13,
	0x4b, 0xfb, 0x9c, 0x58, 0xda, 0xcf, 0x57, 0x73, 0x4f, 0x3e, 0xa8, 0xd7, 0x71, 0xbf, 0x17, 0x26,
	0x99, 0x2e, 0x4c, 0x32, 0x5f, 0x98, 0xe4, 0x63, 0x69, 0x6a, 0xd3, 0xa5, 0xa9, 0xfd, 0x2e, 0x4d,
	0xed, 0xf1, 0x3a, 0x4e, 0xc4, 0xf3, 0xc8, 0x77, 0x02, 0xe0, 0x34, 0x9f, 0x7b, 0x33, 0x65, 0x3e,
	0xae, 0x03, 0x1d, 0xb7, 0xae, 0xe8, 0xeb, 0xff, 0xcd, 0x65, 0x1b, 0xf3, 0x8b, 0xd9, 0xf4, 0x2e,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x59, 0xc9, 0x13, 0x28, 0xde, 0x01, 0x00, 0x00,
}

func (m *CosmWasmPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CosmWasmPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CosmWasmPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InstantiateMsg) > 0 {
		i -= len(m.InstantiateMsg)
		copy(dAtA[i:], m.InstantiateMsg)
		i = encodeVarintPool(dAtA, i, uint64(len(m.InstantiateMsg)))
		i--
		dAtA[i] = 0x22
	}
	if m.CodeId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.CodeId))
		i--
		dAtA[i] = 0x18
	}
	if m.PoolId != 0 {
		i = encodeVarintPool(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintPool(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPool(dAtA []byte, offset int, v uint64) int {
	offset -= sovPool(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CosmWasmPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	if m.PoolId != 0 {
		n += 1 + sovPool(uint64(m.PoolId))
	}
	if m.CodeId != 0 {
		n += 1 + sovPool(uint64(m.CodeId))
	}
	l = len(m.InstantiateMsg)
	if l > 0 {
		n += 1 + l + sovPool(uint64(l))
	}
	return n
}

func sovPool(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPool(x uint64) (n int) {
	return sovPool(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CosmWasmPool) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPool
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
			return fmt.Errorf("proto: CosmWasmPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CosmWasmPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeId", wireType)
			}
			m.CodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstantiateMsg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPool
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
				return ErrInvalidLengthPool
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPool
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InstantiateMsg = append(m.InstantiateMsg[:0], dAtA[iNdEx:postIndex]...)
			if m.InstantiateMsg == nil {
				m.InstantiateMsg = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPool(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPool
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
func skipPool(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
					return 0, ErrIntOverflowPool
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
				return 0, ErrInvalidLengthPool
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPool
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPool
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPool        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPool          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPool = fmt.Errorf("proto: unexpected end of group")
)
