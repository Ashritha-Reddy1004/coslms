// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/lms/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetStudentsRequest struct {
}

func (m *GetStudentsRequest) Reset()         { *m = GetStudentsRequest{} }
func (m *GetStudentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStudentsRequest) ProtoMessage()    {}
func (*GetStudentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aaac713669af5ad, []int{0}
}
func (m *GetStudentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetStudentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetStudentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetStudentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentsRequest.Merge(m, src)
}
func (m *GetStudentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetStudentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentsRequest proto.InternalMessageInfo

type GetStudentsResponse struct {
	Students []*AddStudentRequest `protobuf:"bytes,1,rep,name=students,proto3" json:"students,omitempty"`
}

func (m *GetStudentsResponse) Reset()         { *m = GetStudentsResponse{} }
func (m *GetStudentsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStudentsResponse) ProtoMessage()    {}
func (*GetStudentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aaac713669af5ad, []int{1}
}
func (m *GetStudentsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetStudentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetStudentsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetStudentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStudentsResponse.Merge(m, src)
}
func (m *GetStudentsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetStudentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStudentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStudentsResponse proto.InternalMessageInfo

func (m *GetStudentsResponse) GetStudents() []*AddStudentRequest {
	if m != nil {
		return m.Students
	}
	return nil
}

type GetLeaveRequestsRequest struct {
}

func (m *GetLeaveRequestsRequest) Reset()         { *m = GetLeaveRequestsRequest{} }
func (m *GetLeaveRequestsRequest) String() string { return proto.CompactTextString(m) }
func (*GetLeaveRequestsRequest) ProtoMessage()    {}
func (*GetLeaveRequestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aaac713669af5ad, []int{2}
}
func (m *GetLeaveRequestsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLeaveRequestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLeaveRequestsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLeaveRequestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLeaveRequestsRequest.Merge(m, src)
}
func (m *GetLeaveRequestsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLeaveRequestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLeaveRequestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLeaveRequestsRequest proto.InternalMessageInfo

type GetLeaveRequestsResponse struct {
	Leaverequests []*ApplyLeaveRequest `protobuf:"bytes,1,rep,name=leaverequests,proto3" json:"leaverequests,omitempty"`
}

func (m *GetLeaveRequestsResponse) Reset()         { *m = GetLeaveRequestsResponse{} }
func (m *GetLeaveRequestsResponse) String() string { return proto.CompactTextString(m) }
func (*GetLeaveRequestsResponse) ProtoMessage()    {}
func (*GetLeaveRequestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aaac713669af5ad, []int{3}
}
func (m *GetLeaveRequestsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLeaveRequestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLeaveRequestsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLeaveRequestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLeaveRequestsResponse.Merge(m, src)
}
func (m *GetLeaveRequestsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLeaveRequestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLeaveRequestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLeaveRequestsResponse proto.InternalMessageInfo

func (m *GetLeaveRequestsResponse) GetLeaverequests() []*ApplyLeaveRequest {
	if m != nil {
		return m.Leaverequests
	}
	return nil
}

type GetLeaveApprovedRequestsRequest struct {
}

func (m *GetLeaveApprovedRequestsRequest) Reset()         { *m = GetLeaveApprovedRequestsRequest{} }
func (m *GetLeaveApprovedRequestsRequest) String() string { return proto.CompactTextString(m) }
func (*GetLeaveApprovedRequestsRequest) ProtoMessage()    {}
func (*GetLeaveApprovedRequestsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aaac713669af5ad, []int{4}
}
func (m *GetLeaveApprovedRequestsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLeaveApprovedRequestsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLeaveApprovedRequestsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLeaveApprovedRequestsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLeaveApprovedRequestsRequest.Merge(m, src)
}
func (m *GetLeaveApprovedRequestsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetLeaveApprovedRequestsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLeaveApprovedRequestsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetLeaveApprovedRequestsRequest proto.InternalMessageInfo

type GetLeaveApprovedRequestsResponse struct {
	Leaverequests []*ApplyLeaveRequest `protobuf:"bytes,1,rep,name=leaverequests,proto3" json:"leaverequests,omitempty"`
}

func (m *GetLeaveApprovedRequestsResponse) Reset()         { *m = GetLeaveApprovedRequestsResponse{} }
func (m *GetLeaveApprovedRequestsResponse) String() string { return proto.CompactTextString(m) }
func (*GetLeaveApprovedRequestsResponse) ProtoMessage()    {}
func (*GetLeaveApprovedRequestsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4aaac713669af5ad, []int{5}
}
func (m *GetLeaveApprovedRequestsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetLeaveApprovedRequestsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetLeaveApprovedRequestsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetLeaveApprovedRequestsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLeaveApprovedRequestsResponse.Merge(m, src)
}
func (m *GetLeaveApprovedRequestsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetLeaveApprovedRequestsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLeaveApprovedRequestsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLeaveApprovedRequestsResponse proto.InternalMessageInfo

func (m *GetLeaveApprovedRequestsResponse) GetLeaverequests() []*ApplyLeaveRequest {
	if m != nil {
		return m.Leaverequests
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStudentsRequest)(nil), "cosmos.lms.v1beta1.GetStudentsRequest")
	proto.RegisterType((*GetStudentsResponse)(nil), "cosmos.lms.v1beta1.GetStudentsResponse")
	proto.RegisterType((*GetLeaveRequestsRequest)(nil), "cosmos.lms.v1beta1.GetLeaveRequestsRequest")
	proto.RegisterType((*GetLeaveRequestsResponse)(nil), "cosmos.lms.v1beta1.GetLeaveRequestsResponse")
	proto.RegisterType((*GetLeaveApprovedRequestsRequest)(nil), "cosmos.lms.v1beta1.GetLeaveApprovedRequestsRequest")
	proto.RegisterType((*GetLeaveApprovedRequestsResponse)(nil), "cosmos.lms.v1beta1.GetLeaveApprovedRequestsResponse")
}

func init() { proto.RegisterFile("cosmos/lms/v1beta1/query.proto", fileDescriptor_4aaac713669af5ad) }

var fileDescriptor_4aaac713669af5ad = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0xcf, 0xc9, 0x2d, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x2f,
	0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x82, 0xc8, 0xeb, 0xe5,
	0xe4, 0x16, 0xeb, 0x41, 0xe5, 0xa5, 0xa4, 0xb1, 0xe8, 0x29, 0xa9, 0x80, 0x68, 0x90, 0x92, 0x49,
	0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49,
	0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0xc8, 0x2a, 0x89, 0x70, 0x09, 0xb9, 0xa7, 0x96, 0x04, 0x97,
	0x94, 0xa6, 0xa4, 0xe6, 0x95, 0x14, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x28, 0x45, 0x70,
	0x09, 0xa3, 0x88, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a, 0x39, 0x72, 0x71, 0x14, 0x43, 0xc5,
	0x24, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x54, 0xf5, 0x30, 0x9d, 0xa3, 0xe7, 0x98, 0x92, 0x02,
	0xd5, 0x0a, 0x35, 0x2f, 0x08, 0xae, 0x4d, 0x49, 0x92, 0x4b, 0xdc, 0x3d, 0xb5, 0xc4, 0x27, 0x35,
	0xb1, 0x2c, 0x15, 0x2a, 0x09, 0xb7, 0x34, 0x9d, 0x4b, 0x02, 0x53, 0x0a, 0x6a, 0xb3, 0x37, 0x17,
	0x6f, 0x0e, 0x48, 0xa2, 0x08, 0x2a, 0x81, 0xd7, 0xfa, 0x82, 0x82, 0x9c, 0x4a, 0x64, 0x63, 0x82,
	0x50, 0xf5, 0x2a, 0x29, 0x72, 0xc9, 0xc3, 0x2c, 0x72, 0x2c, 0x28, 0x28, 0xca, 0x2f, 0x4b, 0x4d,
	0x41, 0x77, 0x4b, 0x3e, 0x97, 0x02, 0x6e, 0x25, 0x34, 0x70, 0x93, 0xd1, 0x0b, 0x66, 0x2e, 0xd6,
	0x40, 0x50, 0x34, 0x0b, 0xb5, 0x33, 0x72, 0x71, 0x23, 0x05, 0xbe, 0x90, 0x1a, 0x36, 0xf3, 0x30,
	0xe3, 0x4c, 0x4a, 0x9d, 0xa0, 0x3a, 0x88, 0xbb, 0x95, 0xd4, 0x9b, 0x2e, 0x3f, 0x99, 0xcc, 0xa4,
	0x28, 0x24, 0xaf, 0x8f, 0x25, 0xd9, 0xa4, 0xa7, 0x96, 0xc0, 0xe2, 0x4a, 0x68, 0x36, 0x23, 0x97,
	0x00, 0x7a, 0x8c, 0x08, 0x69, 0xe3, 0xb0, 0x06, 0x5b, 0x94, 0x4a, 0xe9, 0x10, 0xa7, 0x18, 0xea,
	0x30, 0x1d, 0xb0, 0xc3, 0xd4, 0x84, 0x54, 0x70, 0x38, 0x0c, 0x25, 0xc4, 0x84, 0xb6, 0x31, 0x22,
	0xd2, 0x0b, 0x7a, 0x1c, 0x09, 0x19, 0xe3, 0xb3, 0x18, 0x47, 0xa4, 0x4b, 0x99, 0x90, 0xa6, 0x09,
	0xea, 0x6a, 0x7d, 0xb0, 0xab, 0x35, 0x85, 0xd4, 0x71, 0xb8, 0x3a, 0x11, 0xaa, 0x11, 0xe6, 0x70,
	0x27, 0xd5, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2,
	0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xe2, 0xae, 0x00, 0x6b,
	0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x67, 0x50, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4c, 0x61, 0x43, 0xb3, 0x11, 0x04, 0x00, 0x00,
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
	GetStudents(ctx context.Context, in *GetStudentsRequest, opts ...grpc.CallOption) (*GetStudentsResponse, error)
	GetLeaveRequests(ctx context.Context, in *GetLeaveRequestsRequest, opts ...grpc.CallOption) (*GetLeaveRequestsResponse, error)
	GetLeaveApprovedRequests(ctx context.Context, in *GetLeaveApprovedRequestsRequest, opts ...grpc.CallOption) (*GetLeaveApprovedRequestsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetStudents(ctx context.Context, in *GetStudentsRequest, opts ...grpc.CallOption) (*GetStudentsResponse, error) {
	out := new(GetStudentsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.lms.v1beta1.Query/GetStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetLeaveRequests(ctx context.Context, in *GetLeaveRequestsRequest, opts ...grpc.CallOption) (*GetLeaveRequestsResponse, error) {
	out := new(GetLeaveRequestsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.lms.v1beta1.Query/GetLeaveRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetLeaveApprovedRequests(ctx context.Context, in *GetLeaveApprovedRequestsRequest, opts ...grpc.CallOption) (*GetLeaveApprovedRequestsResponse, error) {
	out := new(GetLeaveApprovedRequestsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.lms.v1beta1.Query/GetLeaveApprovedRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	GetStudents(context.Context, *GetStudentsRequest) (*GetStudentsResponse, error)
	GetLeaveRequests(context.Context, *GetLeaveRequestsRequest) (*GetLeaveRequestsResponse, error)
	GetLeaveApprovedRequests(context.Context, *GetLeaveApprovedRequestsRequest) (*GetLeaveApprovedRequestsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetStudents(ctx context.Context, req *GetStudentsRequest) (*GetStudentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudents not implemented")
}
func (*UnimplementedQueryServer) GetLeaveRequests(ctx context.Context, req *GetLeaveRequestsRequest) (*GetLeaveRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaveRequests not implemented")
}
func (*UnimplementedQueryServer) GetLeaveApprovedRequests(ctx context.Context, req *GetLeaveApprovedRequestsRequest) (*GetLeaveApprovedRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaveApprovedRequests not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetStudents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStudentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetStudents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.lms.v1beta1.Query/GetStudents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetStudents(ctx, req.(*GetStudentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetLeaveRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaveRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetLeaveRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.lms.v1beta1.Query/GetLeaveRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetLeaveRequests(ctx, req.(*GetLeaveRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetLeaveApprovedRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLeaveApprovedRequestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetLeaveApprovedRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.lms.v1beta1.Query/GetLeaveApprovedRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetLeaveApprovedRequests(ctx, req.(*GetLeaveApprovedRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.lms.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStudents",
			Handler:    _Query_GetStudents_Handler,
		},
		{
			MethodName: "GetLeaveRequests",
			Handler:    _Query_GetLeaveRequests_Handler,
		},
		{
			MethodName: "GetLeaveApprovedRequests",
			Handler:    _Query_GetLeaveApprovedRequests_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/lms/v1beta1/query.proto",
}

func (m *GetStudentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStudentsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetStudentsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetStudentsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStudentsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetStudentsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Students) > 0 {
		for iNdEx := len(m.Students) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Students[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetLeaveRequestsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLeaveRequestsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLeaveRequestsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetLeaveRequestsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLeaveRequestsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLeaveRequestsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Leaverequests) > 0 {
		for iNdEx := len(m.Leaverequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Leaverequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetLeaveApprovedRequestsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLeaveApprovedRequestsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLeaveApprovedRequestsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *GetLeaveApprovedRequestsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetLeaveApprovedRequestsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetLeaveApprovedRequestsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Leaverequests) > 0 {
		for iNdEx := len(m.Leaverequests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Leaverequests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
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
func (m *GetStudentsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetStudentsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Students) > 0 {
		for _, e := range m.Students {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *GetLeaveRequestsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetLeaveRequestsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Leaverequests) > 0 {
		for _, e := range m.Leaverequests {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *GetLeaveApprovedRequestsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetLeaveApprovedRequestsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Leaverequests) > 0 {
		for _, e := range m.Leaverequests {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetStudentsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetStudentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStudentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *GetStudentsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetStudentsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStudentsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Students", wireType)
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
			m.Students = append(m.Students, &AddStudentRequest{})
			if err := m.Students[len(m.Students)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GetLeaveRequestsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetLeaveRequestsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLeaveRequestsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *GetLeaveRequestsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetLeaveRequestsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLeaveRequestsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leaverequests", wireType)
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
			m.Leaverequests = append(m.Leaverequests, &ApplyLeaveRequest{})
			if err := m.Leaverequests[len(m.Leaverequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GetLeaveApprovedRequestsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetLeaveApprovedRequestsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLeaveApprovedRequestsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *GetLeaveApprovedRequestsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetLeaveApprovedRequestsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetLeaveApprovedRequestsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Leaverequests", wireType)
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
			m.Leaverequests = append(m.Leaverequests, &ApplyLeaveRequest{})
			if err := m.Leaverequests[len(m.Leaverequests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
