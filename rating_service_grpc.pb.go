// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: rating_service.proto

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
	Rating_ListRatings_FullMethodName     = "/songcontestrater.Rating/ListRatings"
	Rating_ListUserRatings_FullMethodName = "/songcontestrater.Rating/ListUserRatings"
	Rating_GetRating_FullMethodName       = "/songcontestrater.Rating/GetRating"
	Rating_CreateRating_FullMethodName    = "/songcontestrater.Rating/CreateRating"
	Rating_UpdateRating_FullMethodName    = "/songcontestrater.Rating/UpdateRating"
	Rating_DeleteRating_FullMethodName    = "/songcontestrater.Rating/DeleteRating"
)

// RatingClient is the client API for Rating service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingClient interface {
	ListRatings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRatingsResponse, error)
	ListUserRatings(ctx context.Context, in *ListUserRatingsRequest, opts ...grpc.CallOption) (*ListRatingsResponse, error)
	GetRating(ctx context.Context, in *GetRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	CreateRating(ctx context.Context, in *CreateRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	UpdateRating(ctx context.Context, in *UpdateRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error)
	DeleteRating(ctx context.Context, in *DeleteRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error)
}

type ratingClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingClient(cc grpc.ClientConnInterface) RatingClient {
	return &ratingClient{cc}
}

func (c *ratingClient) ListRatings(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListRatingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRatingsResponse)
	err := c.cc.Invoke(ctx, Rating_ListRatings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingClient) ListUserRatings(ctx context.Context, in *ListUserRatingsRequest, opts ...grpc.CallOption) (*ListRatingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRatingsResponse)
	err := c.cc.Invoke(ctx, Rating_ListUserRatings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingClient) GetRating(ctx context.Context, in *GetRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, Rating_GetRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingClient) CreateRating(ctx context.Context, in *CreateRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, Rating_CreateRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingClient) UpdateRating(ctx context.Context, in *UpdateRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, Rating_UpdateRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingClient) DeleteRating(ctx context.Context, in *DeleteRatingRequest, opts ...grpc.CallOption) (*RatingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RatingResponse)
	err := c.cc.Invoke(ctx, Rating_DeleteRating_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServer is the server API for Rating service.
// All implementations must embed UnimplementedRatingServer
// for forward compatibility.
type RatingServer interface {
	ListRatings(context.Context, *emptypb.Empty) (*ListRatingsResponse, error)
	ListUserRatings(context.Context, *ListUserRatingsRequest) (*ListRatingsResponse, error)
	GetRating(context.Context, *GetRatingRequest) (*RatingResponse, error)
	CreateRating(context.Context, *CreateRatingRequest) (*RatingResponse, error)
	UpdateRating(context.Context, *UpdateRatingRequest) (*RatingResponse, error)
	DeleteRating(context.Context, *DeleteRatingRequest) (*RatingResponse, error)
	mustEmbedUnimplementedRatingServer()
}

// UnimplementedRatingServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRatingServer struct{}

func (UnimplementedRatingServer) ListRatings(context.Context, *emptypb.Empty) (*ListRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRatings not implemented")
}
func (UnimplementedRatingServer) ListUserRatings(context.Context, *ListUserRatingsRequest) (*ListRatingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserRatings not implemented")
}
func (UnimplementedRatingServer) GetRating(context.Context, *GetRatingRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRating not implemented")
}
func (UnimplementedRatingServer) CreateRating(context.Context, *CreateRatingRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRating not implemented")
}
func (UnimplementedRatingServer) UpdateRating(context.Context, *UpdateRatingRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRating not implemented")
}
func (UnimplementedRatingServer) DeleteRating(context.Context, *DeleteRatingRequest) (*RatingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRating not implemented")
}
func (UnimplementedRatingServer) mustEmbedUnimplementedRatingServer() {}
func (UnimplementedRatingServer) testEmbeddedByValue()                {}

// UnsafeRatingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServer will
// result in compilation errors.
type UnsafeRatingServer interface {
	mustEmbedUnimplementedRatingServer()
}

func RegisterRatingServer(s grpc.ServiceRegistrar, srv RatingServer) {
	// If the following call pancis, it indicates UnimplementedRatingServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Rating_ServiceDesc, srv)
}

func _Rating_ListRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServer).ListRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rating_ListRatings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServer).ListRatings(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rating_ListUserRatings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRatingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServer).ListUserRatings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rating_ListUserRatings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServer).ListUserRatings(ctx, req.(*ListUserRatingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rating_GetRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServer).GetRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rating_GetRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServer).GetRating(ctx, req.(*GetRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rating_CreateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServer).CreateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rating_CreateRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServer).CreateRating(ctx, req.(*CreateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rating_UpdateRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServer).UpdateRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rating_UpdateRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServer).UpdateRating(ctx, req.(*UpdateRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rating_DeleteRating_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRatingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServer).DeleteRating(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Rating_DeleteRating_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServer).DeleteRating(ctx, req.(*DeleteRatingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Rating_ServiceDesc is the grpc.ServiceDesc for Rating service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rating_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "songcontestrater.Rating",
	HandlerType: (*RatingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRatings",
			Handler:    _Rating_ListRatings_Handler,
		},
		{
			MethodName: "ListUserRatings",
			Handler:    _Rating_ListUserRatings_Handler,
		},
		{
			MethodName: "GetRating",
			Handler:    _Rating_GetRating_Handler,
		},
		{
			MethodName: "CreateRating",
			Handler:    _Rating_CreateRating_Handler,
		},
		{
			MethodName: "UpdateRating",
			Handler:    _Rating_UpdateRating_Handler,
		},
		{
			MethodName: "DeleteRating",
			Handler:    _Rating_DeleteRating_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rating_service.proto",
}
