// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: competition_service.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Competition_ListCompetitions_FullMethodName  = "/songcontestrater.Competition/ListCompetitions"
	Competition_GetCompetition_FullMethodName    = "/songcontestrater.Competition/GetCompetition"
	Competition_CreateCompetition_FullMethodName = "/songcontestrater.Competition/CreateCompetition"
	Competition_UpdateCompetition_FullMethodName = "/songcontestrater.Competition/UpdateCompetition"
	Competition_DeleteCompetition_FullMethodName = "/songcontestrater.Competition/DeleteCompetition"
)

// CompetitionClient is the client API for Competition service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompetitionClient interface {
	ListCompetitions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListCompetitionsResponse, error)
	GetCompetition(ctx context.Context, in *GetCompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error)
	CreateCompetition(ctx context.Context, in *CreateCompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error)
	UpdateCompetition(ctx context.Context, in *UpdateCompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error)
	DeleteCompetition(ctx context.Context, in *DeleteCompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error)
}

type competitionClient struct {
	cc grpc.ClientConnInterface
}

func NewCompetitionClient(cc grpc.ClientConnInterface) CompetitionClient {
	return &competitionClient{cc}
}

func (c *competitionClient) ListCompetitions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListCompetitionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCompetitionsResponse)
	err := c.cc.Invoke(ctx, Competition_ListCompetitions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionClient) GetCompetition(ctx context.Context, in *GetCompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompetitionResponse)
	err := c.cc.Invoke(ctx, Competition_GetCompetition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionClient) CreateCompetition(ctx context.Context, in *CreateCompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompetitionResponse)
	err := c.cc.Invoke(ctx, Competition_CreateCompetition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionClient) UpdateCompetition(ctx context.Context, in *UpdateCompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompetitionResponse)
	err := c.cc.Invoke(ctx, Competition_UpdateCompetition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *competitionClient) DeleteCompetition(ctx context.Context, in *DeleteCompetitionRequest, opts ...grpc.CallOption) (*CompetitionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompetitionResponse)
	err := c.cc.Invoke(ctx, Competition_DeleteCompetition_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompetitionServer is the server API for Competition service.
// All implementations must embed UnimplementedCompetitionServer
// for forward compatibility.
type CompetitionServer interface {
	ListCompetitions(context.Context, *emptypb.Empty) (*ListCompetitionsResponse, error)
	GetCompetition(context.Context, *GetCompetitionRequest) (*CompetitionResponse, error)
	CreateCompetition(context.Context, *CreateCompetitionRequest) (*CompetitionResponse, error)
	UpdateCompetition(context.Context, *UpdateCompetitionRequest) (*CompetitionResponse, error)
	DeleteCompetition(context.Context, *DeleteCompetitionRequest) (*CompetitionResponse, error)
	mustEmbedUnimplementedCompetitionServer()
}

// UnimplementedCompetitionServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCompetitionServer struct{}

func (UnimplementedCompetitionServer) ListCompetitions(context.Context, *emptypb.Empty) (*ListCompetitionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCompetitions not implemented")
}
func (UnimplementedCompetitionServer) GetCompetition(context.Context, *GetCompetitionRequest) (*CompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCompetition not implemented")
}
func (UnimplementedCompetitionServer) CreateCompetition(context.Context, *CreateCompetitionRequest) (*CompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCompetition not implemented")
}
func (UnimplementedCompetitionServer) UpdateCompetition(context.Context, *UpdateCompetitionRequest) (*CompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCompetition not implemented")
}
func (UnimplementedCompetitionServer) DeleteCompetition(context.Context, *DeleteCompetitionRequest) (*CompetitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCompetition not implemented")
}
func (UnimplementedCompetitionServer) mustEmbedUnimplementedCompetitionServer() {}
func (UnimplementedCompetitionServer) testEmbeddedByValue()                     {}

// UnsafeCompetitionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompetitionServer will
// result in compilation errors.
type UnsafeCompetitionServer interface {
	mustEmbedUnimplementedCompetitionServer()
}

func RegisterCompetitionServer(s grpc.ServiceRegistrar, srv CompetitionServer) {
	// If the following call pancis, it indicates UnimplementedCompetitionServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Competition_ServiceDesc, srv)
}

func _Competition_ListCompetitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServer).ListCompetitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Competition_ListCompetitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServer).ListCompetitions(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Competition_GetCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCompetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServer).GetCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Competition_GetCompetition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServer).GetCompetition(ctx, req.(*GetCompetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Competition_CreateCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCompetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServer).CreateCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Competition_CreateCompetition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServer).CreateCompetition(ctx, req.(*CreateCompetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Competition_UpdateCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCompetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServer).UpdateCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Competition_UpdateCompetition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServer).UpdateCompetition(ctx, req.(*UpdateCompetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Competition_DeleteCompetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCompetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompetitionServer).DeleteCompetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Competition_DeleteCompetition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompetitionServer).DeleteCompetition(ctx, req.(*DeleteCompetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Competition_ServiceDesc is the grpc.ServiceDesc for Competition service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Competition_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "songcontestrater.Competition",
	HandlerType: (*CompetitionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCompetitions",
			Handler:    _Competition_ListCompetitions_Handler,
		},
		{
			MethodName: "GetCompetition",
			Handler:    _Competition_GetCompetition_Handler,
		},
		{
			MethodName: "CreateCompetition",
			Handler:    _Competition_CreateCompetition_Handler,
		},
		{
			MethodName: "UpdateCompetition",
			Handler:    _Competition_UpdateCompetition_Handler,
		},
		{
			MethodName: "DeleteCompetition",
			Handler:    _Competition_DeleteCompetition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "competition_service.proto",
}
