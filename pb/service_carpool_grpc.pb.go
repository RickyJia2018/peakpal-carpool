// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: service_carpool.proto

package pb

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

const (
	Carpool_CreateTrip_FullMethodName      = "/pb.Carpool/CreateTrip"
	Carpool_GetTrip_FullMethodName         = "/pb.Carpool/GetTrip"
	Carpool_ListFutureTrips_FullMethodName = "/pb.Carpool/ListFutureTrips"
	Carpool_ListDriverTrips_FullMethodName = "/pb.Carpool/ListDriverTrips"
	Carpool_UpdateTrip_FullMethodName      = "/pb.Carpool/UpdateTrip"
	Carpool_DeleteTrip_FullMethodName      = "/pb.Carpool/DeleteTrip"
)

// CarpoolClient is the client API for Carpool service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarpoolClient interface {
	CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*CreateTripResponse, error)
	GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*GetTripResponse, error)
	ListFutureTrips(ctx context.Context, in *ListFutureTripsRequest, opts ...grpc.CallOption) (*ListTripsResponse, error)
	ListDriverTrips(ctx context.Context, in *ListDriverTripsRequest, opts ...grpc.CallOption) (*ListTripsResponse, error)
	UpdateTrip(ctx context.Context, in *UpdateTripRequest, opts ...grpc.CallOption) (*GetTripResponse, error)
	DeleteTrip(ctx context.Context, in *DeleteTripRequest, opts ...grpc.CallOption) (*DeleteTripResponse, error)
}

type carpoolClient struct {
	cc grpc.ClientConnInterface
}

func NewCarpoolClient(cc grpc.ClientConnInterface) CarpoolClient {
	return &carpoolClient{cc}
}

func (c *carpoolClient) CreateTrip(ctx context.Context, in *CreateTripRequest, opts ...grpc.CallOption) (*CreateTripResponse, error) {
	out := new(CreateTripResponse)
	err := c.cc.Invoke(ctx, Carpool_CreateTrip_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carpoolClient) GetTrip(ctx context.Context, in *GetTripRequest, opts ...grpc.CallOption) (*GetTripResponse, error) {
	out := new(GetTripResponse)
	err := c.cc.Invoke(ctx, Carpool_GetTrip_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carpoolClient) ListFutureTrips(ctx context.Context, in *ListFutureTripsRequest, opts ...grpc.CallOption) (*ListTripsResponse, error) {
	out := new(ListTripsResponse)
	err := c.cc.Invoke(ctx, Carpool_ListFutureTrips_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carpoolClient) ListDriverTrips(ctx context.Context, in *ListDriverTripsRequest, opts ...grpc.CallOption) (*ListTripsResponse, error) {
	out := new(ListTripsResponse)
	err := c.cc.Invoke(ctx, Carpool_ListDriverTrips_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carpoolClient) UpdateTrip(ctx context.Context, in *UpdateTripRequest, opts ...grpc.CallOption) (*GetTripResponse, error) {
	out := new(GetTripResponse)
	err := c.cc.Invoke(ctx, Carpool_UpdateTrip_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carpoolClient) DeleteTrip(ctx context.Context, in *DeleteTripRequest, opts ...grpc.CallOption) (*DeleteTripResponse, error) {
	out := new(DeleteTripResponse)
	err := c.cc.Invoke(ctx, Carpool_DeleteTrip_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CarpoolServer is the server API for Carpool service.
// All implementations must embed UnimplementedCarpoolServer
// for forward compatibility
type CarpoolServer interface {
	CreateTrip(context.Context, *CreateTripRequest) (*CreateTripResponse, error)
	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
	ListFutureTrips(context.Context, *ListFutureTripsRequest) (*ListTripsResponse, error)
	ListDriverTrips(context.Context, *ListDriverTripsRequest) (*ListTripsResponse, error)
	UpdateTrip(context.Context, *UpdateTripRequest) (*GetTripResponse, error)
	DeleteTrip(context.Context, *DeleteTripRequest) (*DeleteTripResponse, error)
	mustEmbedUnimplementedCarpoolServer()
}

// UnimplementedCarpoolServer must be embedded to have forward compatible implementations.
type UnimplementedCarpoolServer struct {
}

func (UnimplementedCarpoolServer) CreateTrip(context.Context, *CreateTripRequest) (*CreateTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTrip not implemented")
}
func (UnimplementedCarpoolServer) GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrip not implemented")
}
func (UnimplementedCarpoolServer) ListFutureTrips(context.Context, *ListFutureTripsRequest) (*ListTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFutureTrips not implemented")
}
func (UnimplementedCarpoolServer) ListDriverTrips(context.Context, *ListDriverTripsRequest) (*ListTripsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDriverTrips not implemented")
}
func (UnimplementedCarpoolServer) UpdateTrip(context.Context, *UpdateTripRequest) (*GetTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTrip not implemented")
}
func (UnimplementedCarpoolServer) DeleteTrip(context.Context, *DeleteTripRequest) (*DeleteTripResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTrip not implemented")
}
func (UnimplementedCarpoolServer) mustEmbedUnimplementedCarpoolServer() {}

// UnsafeCarpoolServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarpoolServer will
// result in compilation errors.
type UnsafeCarpoolServer interface {
	mustEmbedUnimplementedCarpoolServer()
}

func RegisterCarpoolServer(s grpc.ServiceRegistrar, srv CarpoolServer) {
	s.RegisterService(&Carpool_ServiceDesc, srv)
}

func _Carpool_CreateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarpoolServer).CreateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carpool_CreateTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarpoolServer).CreateTrip(ctx, req.(*CreateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carpool_GetTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarpoolServer).GetTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carpool_GetTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarpoolServer).GetTrip(ctx, req.(*GetTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carpool_ListFutureTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFutureTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarpoolServer).ListFutureTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carpool_ListFutureTrips_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarpoolServer).ListFutureTrips(ctx, req.(*ListFutureTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carpool_ListDriverTrips_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDriverTripsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarpoolServer).ListDriverTrips(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carpool_ListDriverTrips_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarpoolServer).ListDriverTrips(ctx, req.(*ListDriverTripsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carpool_UpdateTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarpoolServer).UpdateTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carpool_UpdateTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarpoolServer).UpdateTrip(ctx, req.(*UpdateTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Carpool_DeleteTrip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTripRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarpoolServer).DeleteTrip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Carpool_DeleteTrip_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarpoolServer).DeleteTrip(ctx, req.(*DeleteTripRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Carpool_ServiceDesc is the grpc.ServiceDesc for Carpool service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Carpool_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Carpool",
	HandlerType: (*CarpoolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTrip",
			Handler:    _Carpool_CreateTrip_Handler,
		},
		{
			MethodName: "GetTrip",
			Handler:    _Carpool_GetTrip_Handler,
		},
		{
			MethodName: "ListFutureTrips",
			Handler:    _Carpool_ListFutureTrips_Handler,
		},
		{
			MethodName: "ListDriverTrips",
			Handler:    _Carpool_ListDriverTrips_Handler,
		},
		{
			MethodName: "UpdateTrip",
			Handler:    _Carpool_UpdateTrip_Handler,
		},
		{
			MethodName: "DeleteTrip",
			Handler:    _Carpool_DeleteTrip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_carpool.proto",
}

const (
	PeakPal_AuthorizeUser_FullMethodName = "/pb.PeakPal/AuthorizeUser"
	PeakPal_GetResort_FullMethodName     = "/pb.PeakPal/GetResort"
)

// PeakPalClient is the client API for PeakPal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PeakPalClient interface {
	AuthorizeUser(ctx context.Context, in *AuthorizeUserRequest, opts ...grpc.CallOption) (*AuthorizeUserResponse, error)
	GetResort(ctx context.Context, in *GetResortRequest, opts ...grpc.CallOption) (*GetResortResponse, error)
}

type peakPalClient struct {
	cc grpc.ClientConnInterface
}

func NewPeakPalClient(cc grpc.ClientConnInterface) PeakPalClient {
	return &peakPalClient{cc}
}

func (c *peakPalClient) AuthorizeUser(ctx context.Context, in *AuthorizeUserRequest, opts ...grpc.CallOption) (*AuthorizeUserResponse, error) {
	out := new(AuthorizeUserResponse)
	err := c.cc.Invoke(ctx, PeakPal_AuthorizeUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *peakPalClient) GetResort(ctx context.Context, in *GetResortRequest, opts ...grpc.CallOption) (*GetResortResponse, error) {
	out := new(GetResortResponse)
	err := c.cc.Invoke(ctx, PeakPal_GetResort_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PeakPalServer is the server API for PeakPal service.
// All implementations must embed UnimplementedPeakPalServer
// for forward compatibility
type PeakPalServer interface {
	AuthorizeUser(context.Context, *AuthorizeUserRequest) (*AuthorizeUserResponse, error)
	GetResort(context.Context, *GetResortRequest) (*GetResortResponse, error)
	mustEmbedUnimplementedPeakPalServer()
}

// UnimplementedPeakPalServer must be embedded to have forward compatible implementations.
type UnimplementedPeakPalServer struct {
}

func (UnimplementedPeakPalServer) AuthorizeUser(context.Context, *AuthorizeUserRequest) (*AuthorizeUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeUser not implemented")
}
func (UnimplementedPeakPalServer) GetResort(context.Context, *GetResortRequest) (*GetResortResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetResort not implemented")
}
func (UnimplementedPeakPalServer) mustEmbedUnimplementedPeakPalServer() {}

// UnsafePeakPalServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PeakPalServer will
// result in compilation errors.
type UnsafePeakPalServer interface {
	mustEmbedUnimplementedPeakPalServer()
}

func RegisterPeakPalServer(s grpc.ServiceRegistrar, srv PeakPalServer) {
	s.RegisterService(&PeakPal_ServiceDesc, srv)
}

func _PeakPal_AuthorizeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeakPalServer).AuthorizeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeakPal_AuthorizeUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeakPalServer).AuthorizeUser(ctx, req.(*AuthorizeUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PeakPal_GetResort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResortRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PeakPalServer).GetResort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PeakPal_GetResort_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PeakPalServer).GetResort(ctx, req.(*GetResortRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PeakPal_ServiceDesc is the grpc.ServiceDesc for PeakPal service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PeakPal_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.PeakPal",
	HandlerType: (*PeakPalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizeUser",
			Handler:    _PeakPal_AuthorizeUser_Handler,
		},
		{
			MethodName: "GetResort",
			Handler:    _PeakPal_GetResort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_carpool.proto",
}
