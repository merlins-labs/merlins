// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: merlins/downtime-detector/v1beta1/query.proto

package queryproto

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/osmosis-labs/osmosis/v17/x/downtime-detector/types"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Query for has it been at least $RECOVERY_DURATION units of time,
// since the chain has been down for $DOWNTIME_DURATION.
type RecoveredSinceDowntimeOfLengthRequest struct {
	Downtime types.Downtime `protobuf:"varint,1,opt,name=downtime,proto3,enum=merlins.downtimedetector.v1beta1.Downtime" json:"downtime,omitempty" yaml:"downtime"`
	Recovery time.Duration  `protobuf:"bytes,2,opt,name=recovery,proto3,stdduration" json:"recovery" yaml:"recovery_duration"`
}

func (m *RecoveredSinceDowntimeOfLengthRequest) Reset()         { *m = RecoveredSinceDowntimeOfLengthRequest{} }
func (m *RecoveredSinceDowntimeOfLengthRequest) String() string { return proto.CompactTextString(m) }
func (*RecoveredSinceDowntimeOfLengthRequest) ProtoMessage()    {}
func (*RecoveredSinceDowntimeOfLengthRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b748b3d07fa8b8cb, []int{0}
}
func (m *RecoveredSinceDowntimeOfLengthRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecoveredSinceDowntimeOfLengthRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecoveredSinceDowntimeOfLengthRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecoveredSinceDowntimeOfLengthRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveredSinceDowntimeOfLengthRequest.Merge(m, src)
}
func (m *RecoveredSinceDowntimeOfLengthRequest) XXX_Size() int {
	return m.Size()
}
func (m *RecoveredSinceDowntimeOfLengthRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveredSinceDowntimeOfLengthRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveredSinceDowntimeOfLengthRequest proto.InternalMessageInfo

func (m *RecoveredSinceDowntimeOfLengthRequest) GetDowntime() types.Downtime {
	if m != nil {
		return m.Downtime
	}
	return types.Downtime_DURATION_30S
}

func (m *RecoveredSinceDowntimeOfLengthRequest) GetRecovery() time.Duration {
	if m != nil {
		return m.Recovery
	}
	return 0
}

type RecoveredSinceDowntimeOfLengthResponse struct {
	SuccesfullyRecovered bool `protobuf:"varint,1,opt,name=succesfully_recovered,json=succesfullyRecovered,proto3" json:"succesfully_recovered,omitempty"`
}

func (m *RecoveredSinceDowntimeOfLengthResponse) Reset() {
	*m = RecoveredSinceDowntimeOfLengthResponse{}
}
func (m *RecoveredSinceDowntimeOfLengthResponse) String() string { return proto.CompactTextString(m) }
func (*RecoveredSinceDowntimeOfLengthResponse) ProtoMessage()    {}
func (*RecoveredSinceDowntimeOfLengthResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b748b3d07fa8b8cb, []int{1}
}
func (m *RecoveredSinceDowntimeOfLengthResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RecoveredSinceDowntimeOfLengthResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RecoveredSinceDowntimeOfLengthResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RecoveredSinceDowntimeOfLengthResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecoveredSinceDowntimeOfLengthResponse.Merge(m, src)
}
func (m *RecoveredSinceDowntimeOfLengthResponse) XXX_Size() int {
	return m.Size()
}
func (m *RecoveredSinceDowntimeOfLengthResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RecoveredSinceDowntimeOfLengthResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RecoveredSinceDowntimeOfLengthResponse proto.InternalMessageInfo

func (m *RecoveredSinceDowntimeOfLengthResponse) GetSuccesfullyRecovered() bool {
	if m != nil {
		return m.SuccesfullyRecovered
	}
	return false
}

func init() {
	proto.RegisterType((*RecoveredSinceDowntimeOfLengthRequest)(nil), "merlins.downtimedetector.v1beta1.RecoveredSinceDowntimeOfLengthRequest")
	proto.RegisterType((*RecoveredSinceDowntimeOfLengthResponse)(nil), "merlins.downtimedetector.v1beta1.RecoveredSinceDowntimeOfLengthResponse")
}

func init() {
	proto.RegisterFile("merlins/downtime-detector/v1beta1/query.proto", fileDescriptor_b748b3d07fa8b8cb)
}

var fileDescriptor_b748b3d07fa8b8cb = []byte{
	// 483 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4f, 0x8b, 0xd3, 0x40,
	0x14, 0x6f, 0x0a, 0xca, 0x12, 0x41, 0x21, 0xae, 0xd0, 0x2d, 0x92, 0x94, 0xa0, 0xb2, 0x2c, 0x74,
	0x86, 0x76, 0x0f, 0xa2, 0x37, 0xeb, 0x82, 0x2e, 0x08, 0x62, 0xbc, 0x29, 0x52, 0x26, 0xd3, 0xd7,
	0xec, 0x40, 0x32, 0x93, 0xcd, 0x4c, 0xaa, 0xb9, 0x7a, 0xf2, 0x28, 0x78, 0xf1, 0x23, 0xed, 0x71,
	0xc1, 0x8b, 0xa7, 0xaa, 0xad, 0x9f, 0x60, 0x3f, 0x80, 0x48, 0x32, 0x99, 0x50, 0x8a, 0x6c, 0x16,
	0x3c, 0x25, 0xc3, 0xef, 0xcf, 0x7b, 0xbf, 0xf7, 0x66, 0xec, 0xa1, 0x90, 0x89, 0x90, 0x4c, 0xe2,
	0x99, 0x78, 0xcf, 0x15, 0x4b, 0x60, 0x38, 0x03, 0x05, 0x54, 0x89, 0x0c, 0x2f, 0x46, 0x21, 0x28,
	0x32, 0xc2, 0xa7, 0x39, 0x64, 0x05, 0x4a, 0x33, 0xa1, 0x84, 0x33, 0xa8, 0xe9, 0xc8, 0xd0, 0x0d,
	0x1b, 0xd5, 0xec, 0xfe, 0x6e, 0x24, 0x22, 0x51, 0x91, 0x71, 0xf9, 0xa7, 0x75, 0x7d, 0xdc, 0x5e,
	0x26, 0x02, 0x0e, 0xa5, 0xb3, 0x16, 0x3c, 0x6a, 0x17, 0x18, 0x64, 0x3a, 0xcb, 0x33, 0xa2, 0x98,
	0xe0, 0xb5, 0xd4, 0xa5, 0x95, 0x16, 0x87, 0x44, 0x42, 0x43, 0xa6, 0x82, 0x19, 0xfc, 0x60, 0x13,
	0xaf, 0xc2, 0x35, 0xac, 0x94, 0x44, 0x8c, 0x6f, 0x7a, 0xdd, 0x8d, 0x84, 0x88, 0x62, 0xc0, 0x24,
	0x65, 0x98, 0x70, 0x2e, 0x54, 0x05, 0x9a, 0x26, 0xf7, 0x6a, 0xb4, 0x3a, 0x85, 0xf9, 0x1c, 0x13,
	0x5e, 0x18, 0x48, 0x17, 0x99, 0xea, 0x49, 0xe8, 0x83, 0xe9, 0x6f, 0x5b, 0xb5, 0xd5, 0xbf, 0xb7,
	0x8d, 0x97, 0x21, 0xa5, 0x22, 0x49, 0xaa, 0x09, 0xfe, 0x2f, 0xcb, 0xbe, 0x1f, 0x00, 0x15, 0x0b,
	0xc8, 0x60, 0xf6, 0x9a, 0x71, 0x0a, 0x47, 0xf5, 0x28, 0x5e, 0xce, 0x5f, 0x00, 0x8f, 0xd4, 0x49,
	0x00, 0xa7, 0x39, 0x48, 0xe5, 0xbc, 0xb5, 0x77, 0xcc, 0x94, 0x7a, 0xd6, 0xc0, 0xda, 0xbf, 0x39,
	0x3e, 0x40, 0x6d, 0x1b, 0x44, 0xc6, 0x6c, 0x72, 0xfb, 0x62, 0xe9, 0xdd, 0x2a, 0x48, 0x12, 0x3f,
	0xf6, 0x0d, 0xd9, 0x0f, 0x1a, 0xc3, 0xd2, 0x3c, 0xd3, 0x5d, 0x14, 0xbd, 0xee, 0xc0, 0xda, 0xbf,
	0x31, 0xde, 0x43, 0xba, 0x75, 0x64, 0x5a, 0x47, 0x47, 0x75, 0xb4, 0xc9, 0xbd, 0xb3, 0xa5, 0xd7,
	0xb9, 0x58, 0x7a, 0x3d, 0xed, 0x67, 0x84, 0xcd, 0xee, 0xfc, 0xaf, 0x3f, 0x3c, 0x2b, 0x68, 0x0c,
	0xfd, 0x77, 0xf6, 0x83, 0xb6, 0x88, 0x32, 0x15, 0x5c, 0x82, 0x73, 0x68, 0xdf, 0x91, 0x39, 0xa5,
	0x20, 0xe7, 0x79, 0x1c, 0x17, 0xd3, 0xcc, 0xa8, 0xaa, 0xc0, 0x3b, 0xc1, 0xee, 0x06, 0xd8, 0x38,
	0x8e, 0x3f, 0x75, 0xed, 0x6b, 0xaf, 0xca, 0xd5, 0x3b, 0x7f, 0x2c, 0xdb, 0xbd, 0xbc, 0x92, 0xf3,
	0xac, 0x7d, 0x66, 0x57, 0x5a, 0x47, 0xff, 0xf9, 0xff, 0x1b, 0xe9, 0xd0, 0xfe, 0xf1, 0xc7, 0x6f,
	0xbf, 0xbf, 0x74, 0x9f, 0x3a, 0x4f, 0xae, 0xf0, 0xb0, 0x2e, 0xb7, 0x9c, 0xd0, 0xb3, 0x95, 0x6b,
	0x9d, 0xaf, 0x5c, 0xeb, 0xe7, 0xca, 0xb5, 0x3e, 0xaf, 0xdd, 0xce, 0xf9, 0xda, 0xed, 0x7c, 0x5f,
	0xbb, 0x9d, 0x37, 0xc7, 0x11, 0x53, 0x27, 0x79, 0x88, 0xa8, 0x48, 0x4c, 0x99, 0x61, 0x4c, 0x42,
	0xd9, 0xd4, 0x5c, 0x8c, 0x1e, 0xe2, 0x0f, 0xff, 0xa8, 0x4c, 0x63, 0x06, 0x5c, 0xe9, 0xb7, 0xa5,
	0xaf, 0xc2, 0xf5, 0xea, 0x73, 0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x6d, 0x2e, 0x4b, 0x6f,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	RecoveredSinceDowntimeOfLength(ctx context.Context, in *RecoveredSinceDowntimeOfLengthRequest, opts ...grpc.CallOption) (*RecoveredSinceDowntimeOfLengthResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) RecoveredSinceDowntimeOfLength(ctx context.Context, in *RecoveredSinceDowntimeOfLengthRequest, opts ...grpc.CallOption) (*RecoveredSinceDowntimeOfLengthResponse, error) {
	out := new(RecoveredSinceDowntimeOfLengthResponse)
	err := c.cc.Invoke(ctx, "/merlins.downtimedetector.v1beta1.Query/RecoveredSinceDowntimeOfLength", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	RecoveredSinceDowntimeOfLength(context.Context, *RecoveredSinceDowntimeOfLengthRequest) (*RecoveredSinceDowntimeOfLengthResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) RecoveredSinceDowntimeOfLength(ctx context.Context, req *RecoveredSinceDowntimeOfLengthRequest) (*RecoveredSinceDowntimeOfLengthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RecoveredSinceDowntimeOfLength not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_RecoveredSinceDowntimeOfLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RecoveredSinceDowntimeOfLengthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RecoveredSinceDowntimeOfLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/merlins.downtimedetector.v1beta1.Query/RecoveredSinceDowntimeOfLength",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RecoveredSinceDowntimeOfLength(ctx, req.(*RecoveredSinceDowntimeOfLengthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "merlins.downtimedetector.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RecoveredSinceDowntimeOfLength",
			Handler:    _Query_RecoveredSinceDowntimeOfLength_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "merlins/downtime-detector/v1beta1/query.proto",
}

func (m *RecoveredSinceDowntimeOfLengthRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecoveredSinceDowntimeOfLengthRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecoveredSinceDowntimeOfLengthRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.Recovery, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.Recovery):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintQuery(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.Downtime != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Downtime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *RecoveredSinceDowntimeOfLengthResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RecoveredSinceDowntimeOfLengthResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RecoveredSinceDowntimeOfLengthResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SuccesfullyRecovered {
		i--
		if m.SuccesfullyRecovered {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RecoveredSinceDowntimeOfLengthRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Downtime != 0 {
		n += 1 + sovQuery(uint64(m.Downtime))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.Recovery)
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *RecoveredSinceDowntimeOfLengthResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.SuccesfullyRecovered {
		n += 2
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RecoveredSinceDowntimeOfLengthRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: RecoveredSinceDowntimeOfLengthRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecoveredSinceDowntimeOfLengthRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Downtime", wireType)
			}
			m.Downtime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Downtime |= types.Downtime(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recovery", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.Recovery, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *RecoveredSinceDowntimeOfLengthResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: RecoveredSinceDowntimeOfLengthResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RecoveredSinceDowntimeOfLengthResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SuccesfullyRecovered", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
			m.SuccesfullyRecovered = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
