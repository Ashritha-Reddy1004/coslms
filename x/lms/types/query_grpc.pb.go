// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/cosmos/lms/v1beta1/query.proto

package types

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	GetStudents(ctx context.Context, in *GetStudentsRequest, opts ...grpc.CallOption) (*GetStudentsResponse, error)
	GetLeaveRequests(ctx context.Context, in *GetLeaveRequestsRequest, opts ...grpc.CallOption) (*GetLeaveRequestsResponse, error)
	GetLeaveApprovedRequests(ctx context.Context, in *GetLeaveApprovedRequestsRequest, opts ...grpc.CallOption) (*GetLeaveApprovedRequestsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetStudents(ctx context.Context, in *GetStudentsRequest, opts ...grpc.CallOption) (*GetStudentsResponse, error) {
	out := new(GetStudentsResponse)
	err := c.cc.Invoke(ctx, "/lms.v1beta1.Query/GetStudents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetLeaveRequests(ctx context.Context, in *GetLeaveRequestsRequest, opts ...grpc.CallOption) (*GetLeaveRequestsResponse, error) {
	out := new(GetLeaveRequestsResponse)
	err := c.cc.Invoke(ctx, "/lms.v1beta1.Query/GetLeaveRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetLeaveApprovedRequests(ctx context.Context, in *GetLeaveApprovedRequestsRequest, opts ...grpc.CallOption) (*GetLeaveApprovedRequestsResponse, error) {
	out := new(GetLeaveApprovedRequestsResponse)
	err := c.cc.Invoke(ctx, "/lms.v1beta1.Query/GetLeaveApprovedRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	GetStudents(context.Context, *GetStudentsRequest) (*GetStudentsResponse, error)
	GetLeaveRequests(context.Context, *GetLeaveRequestsRequest) (*GetLeaveRequestsResponse, error)
	GetLeaveApprovedRequests(context.Context, *GetLeaveApprovedRequestsRequest) (*GetLeaveApprovedRequestsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) GetStudents(context.Context, *GetStudentsRequest) (*GetStudentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStudents not implemented")
}
func (UnimplementedQueryServer) GetLeaveRequests(context.Context, *GetLeaveRequestsRequest) (*GetLeaveRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaveRequests not implemented")
}
func (UnimplementedQueryServer) GetLeaveApprovedRequests(context.Context, *GetLeaveApprovedRequestsRequest) (*GetLeaveApprovedRequestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaveApprovedRequests not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
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
		FullMethod: "/lms.v1beta1.Query/GetStudents",
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
		FullMethod: "/lms.v1beta1.Query/GetLeaveRequests",
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
		FullMethod: "/lms.v1beta1.Query/GetLeaveApprovedRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetLeaveApprovedRequests(ctx, req.(*GetLeaveApprovedRequestsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lms.v1beta1.Query",
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
	Metadata: "proto/cosmos/lms/v1beta1/query.proto",
}
