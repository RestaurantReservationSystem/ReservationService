// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: reservation.proto

package genproto

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

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*Void, error)
	UpdateReservation(ctx context.Context, in *UpdateReservationRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteReservation(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Void, error)
	GetByIdReservation(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*ReservationResponse, error)
	GetAllReservation(ctx context.Context, in *GetAllReservationRequest, opts ...grpc.CallOption) (*ReservationsResponse, error)
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Void, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteOrder(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Void, error)
	GetByIdOrder(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*OrderResponse, error)
	GetAllOrder(ctx context.Context, in *GetAllOrderRequest, opts ...grpc.CallOption) (*OrdersResponse, error)
	CreateRestaurant(ctx context.Context, in *CreateRestaurantRequest, opts ...grpc.CallOption) (*Void, error)
	GetByIdRestaurant(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RestaurantResponse, error)
	GetAllRestaurants(ctx context.Context, in *GetAllRestaurantRequest, opts ...grpc.CallOption) (*RestaurantsResponse, error)
	UpdateRestaurant(ctx context.Context, in *UpdateRestaurantRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteRestaurant(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Void, error)
	CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*Void, error)
	UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*Void, error)
	DeleteMenu(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Void, error)
	GetByIdMenu(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MenuResponse, error)
	GetAllMenu(ctx context.Context, in *GetAllMenuRequest, opts ...grpc.CallOption) (*MenusResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) CreateReservation(ctx context.Context, in *CreateReservationRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/CreateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateReservation(ctx context.Context, in *UpdateReservationRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/UpdateReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteReservation(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/DeleteReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetByIdReservation(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*ReservationResponse, error) {
	out := new(ReservationResponse)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/GetByIdReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllReservation(ctx context.Context, in *GetAllReservationRequest, opts ...grpc.CallOption) (*ReservationsResponse, error) {
	out := new(ReservationsResponse)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/GetAllReservation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/UpdateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteOrder(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/DeleteOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetByIdOrder(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/GetByIdOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllOrder(ctx context.Context, in *GetAllOrderRequest, opts ...grpc.CallOption) (*OrdersResponse, error) {
	out := new(OrdersResponse)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/GetAllOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CreateRestaurant(ctx context.Context, in *CreateRestaurantRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/CreateRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetByIdRestaurant(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*RestaurantResponse, error) {
	out := new(RestaurantResponse)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/GetByIdRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllRestaurants(ctx context.Context, in *GetAllRestaurantRequest, opts ...grpc.CallOption) (*RestaurantsResponse, error) {
	out := new(RestaurantsResponse)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/GetAllRestaurants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateRestaurant(ctx context.Context, in *UpdateRestaurantRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/UpdateRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteRestaurant(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/DeleteRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CreateMenu(ctx context.Context, in *CreateMenuRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/CreateMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) UpdateMenu(ctx context.Context, in *UpdateMenuRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/UpdateMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteMenu(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/DeleteMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetByIdMenu(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MenuResponse, error) {
	out := new(MenuResponse)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/GetByIdMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllMenu(ctx context.Context, in *GetAllMenuRequest, opts ...grpc.CallOption) (*MenusResponse, error) {
	out := new(MenusResponse)
	err := c.cc.Invoke(ctx, "/protos.ReservationService/GetAllMenu", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	CreateReservation(context.Context, *CreateReservationRequest) (*Void, error)
	UpdateReservation(context.Context, *UpdateReservationRequest) (*Void, error)
	DeleteReservation(context.Context, *IdRequest) (*Void, error)
	GetByIdReservation(context.Context, *IdRequest) (*ReservationResponse, error)
	GetAllReservation(context.Context, *GetAllReservationRequest) (*ReservationsResponse, error)
	CreateOrder(context.Context, *CreateOrderRequest) (*Void, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*Void, error)
	DeleteOrder(context.Context, *IdRequest) (*Void, error)
	GetByIdOrder(context.Context, *IdRequest) (*OrderResponse, error)
	GetAllOrder(context.Context, *GetAllOrderRequest) (*OrdersResponse, error)
	CreateRestaurant(context.Context, *CreateRestaurantRequest) (*Void, error)
	GetByIdRestaurant(context.Context, *IdRequest) (*RestaurantResponse, error)
	GetAllRestaurants(context.Context, *GetAllRestaurantRequest) (*RestaurantsResponse, error)
	UpdateRestaurant(context.Context, *UpdateRestaurantRequest) (*Void, error)
	DeleteRestaurant(context.Context, *IdRequest) (*Void, error)
	CreateMenu(context.Context, *CreateMenuRequest) (*Void, error)
	UpdateMenu(context.Context, *UpdateMenuRequest) (*Void, error)
	DeleteMenu(context.Context, *IdRequest) (*Void, error)
	GetByIdMenu(context.Context, *IdRequest) (*MenuResponse, error)
	GetAllMenu(context.Context, *GetAllMenuRequest) (*MenusResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) CreateReservation(context.Context, *CreateReservationRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateReservation not implemented")
}
func (UnimplementedReservationServiceServer) UpdateReservation(context.Context, *UpdateReservationRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReservation not implemented")
}
func (UnimplementedReservationServiceServer) DeleteReservation(context.Context, *IdRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservation not implemented")
}
func (UnimplementedReservationServiceServer) GetByIdReservation(context.Context, *IdRequest) (*ReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdReservation not implemented")
}
func (UnimplementedReservationServiceServer) GetAllReservation(context.Context, *GetAllReservationRequest) (*ReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservation not implemented")
}
func (UnimplementedReservationServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedReservationServiceServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedReservationServiceServer) DeleteOrder(context.Context, *IdRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedReservationServiceServer) GetByIdOrder(context.Context, *IdRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdOrder not implemented")
}
func (UnimplementedReservationServiceServer) GetAllOrder(context.Context, *GetAllOrderRequest) (*OrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOrder not implemented")
}
func (UnimplementedReservationServiceServer) CreateRestaurant(context.Context, *CreateRestaurantRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRestaurant not implemented")
}
func (UnimplementedReservationServiceServer) GetByIdRestaurant(context.Context, *IdRequest) (*RestaurantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdRestaurant not implemented")
}
func (UnimplementedReservationServiceServer) GetAllRestaurants(context.Context, *GetAllRestaurantRequest) (*RestaurantsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRestaurants not implemented")
}
func (UnimplementedReservationServiceServer) UpdateRestaurant(context.Context, *UpdateRestaurantRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestaurant not implemented")
}
func (UnimplementedReservationServiceServer) DeleteRestaurant(context.Context, *IdRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRestaurant not implemented")
}
func (UnimplementedReservationServiceServer) CreateMenu(context.Context, *CreateMenuRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMenu not implemented")
}
func (UnimplementedReservationServiceServer) UpdateMenu(context.Context, *UpdateMenuRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenu not implemented")
}
func (UnimplementedReservationServiceServer) DeleteMenu(context.Context, *IdRequest) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMenu not implemented")
}
func (UnimplementedReservationServiceServer) GetByIdMenu(context.Context, *IdRequest) (*MenuResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByIdMenu not implemented")
}
func (UnimplementedReservationServiceServer) GetAllMenu(context.Context, *GetAllMenuRequest) (*MenusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMenu not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_CreateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CreateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/CreateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CreateReservation(ctx, req.(*CreateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/UpdateReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateReservation(ctx, req.(*UpdateReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/DeleteReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteReservation(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetByIdReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetByIdReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/GetByIdReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetByIdReservation(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/GetAllReservation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllReservation(ctx, req.(*GetAllReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/DeleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteOrder(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetByIdOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetByIdOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/GetByIdOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetByIdOrder(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/GetAllOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllOrder(ctx, req.(*GetAllOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CreateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CreateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/CreateRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CreateRestaurant(ctx, req.(*CreateRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetByIdRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetByIdRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/GetByIdRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetByIdRestaurant(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/GetAllRestaurants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllRestaurants(ctx, req.(*GetAllRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRestaurantRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/UpdateRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateRestaurant(ctx, req.(*UpdateRestaurantRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/DeleteRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteRestaurant(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CreateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CreateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/CreateMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CreateMenu(ctx, req.(*CreateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_UpdateMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).UpdateMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/UpdateMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).UpdateMenu(ctx, req.(*UpdateMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/DeleteMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteMenu(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetByIdMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetByIdMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/GetByIdMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetByIdMenu(ctx, req.(*IdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllMenu_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMenuRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllMenu(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.ReservationService/GetAllMenu",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllMenu(ctx, req.(*GetAllMenuRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateReservation",
			Handler:    _ReservationService_CreateReservation_Handler,
		},
		{
			MethodName: "UpdateReservation",
			Handler:    _ReservationService_UpdateReservation_Handler,
		},
		{
			MethodName: "DeleteReservation",
			Handler:    _ReservationService_DeleteReservation_Handler,
		},
		{
			MethodName: "GetByIdReservation",
			Handler:    _ReservationService_GetByIdReservation_Handler,
		},
		{
			MethodName: "GetAllReservation",
			Handler:    _ReservationService_GetAllReservation_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _ReservationService_CreateOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _ReservationService_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _ReservationService_DeleteOrder_Handler,
		},
		{
			MethodName: "GetByIdOrder",
			Handler:    _ReservationService_GetByIdOrder_Handler,
		},
		{
			MethodName: "GetAllOrder",
			Handler:    _ReservationService_GetAllOrder_Handler,
		},
		{
			MethodName: "CreateRestaurant",
			Handler:    _ReservationService_CreateRestaurant_Handler,
		},
		{
			MethodName: "GetByIdRestaurant",
			Handler:    _ReservationService_GetByIdRestaurant_Handler,
		},
		{
			MethodName: "GetAllRestaurants",
			Handler:    _ReservationService_GetAllRestaurants_Handler,
		},
		{
			MethodName: "UpdateRestaurant",
			Handler:    _ReservationService_UpdateRestaurant_Handler,
		},
		{
			MethodName: "DeleteRestaurant",
			Handler:    _ReservationService_DeleteRestaurant_Handler,
		},
		{
			MethodName: "CreateMenu",
			Handler:    _ReservationService_CreateMenu_Handler,
		},
		{
			MethodName: "UpdateMenu",
			Handler:    _ReservationService_UpdateMenu_Handler,
		},
		{
			MethodName: "DeleteMenu",
			Handler:    _ReservationService_DeleteMenu_Handler,
		},
		{
			MethodName: "GetByIdMenu",
			Handler:    _ReservationService_GetByIdMenu_Handler,
		},
		{
			MethodName: "GetAllMenu",
			Handler:    _ReservationService_GetAllMenu_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation.proto",
}
