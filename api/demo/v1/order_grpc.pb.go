// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: demo/v1/order.proto

package v1

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
	OrderService_ConfirmOrder_FullMethodName = "/v1.OrderService/ConfirmOrder"
	OrderService_CreateOrder_FullMethodName  = "/v1.OrderService/CreateOrder"
	OrderService_PayOrder_FullMethodName     = "/v1.OrderService/PayOrder"
	OrderService_OrderInfo_FullMethodName    = "/v1.OrderService/OrderInfo"
	OrderService_OrderList_FullMethodName    = "/v1.OrderService/OrderList"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	//确认订单
	ConfirmOrder(ctx context.Context, in *ConfirmOrderRequest, opts ...grpc.CallOption) (*ConfirmOrderReply, error)
	//提交订单
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderReply, error)
	//支付订单
	PayOrder(ctx context.Context, in *PayOrderRequest, opts ...grpc.CallOption) (*PayOrderReply, error)
	//订单详情
	OrderInfo(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*OrderInfoReply, error)
	//订单列表
	OrderList(ctx context.Context, in *OrderListRequest, opts ...grpc.CallOption) (*OrderListReply, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) ConfirmOrder(ctx context.Context, in *ConfirmOrderRequest, opts ...grpc.CallOption) (*ConfirmOrderReply, error) {
	out := new(ConfirmOrderReply)
	err := c.cc.Invoke(ctx, OrderService_ConfirmOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderReply, error) {
	out := new(CreateOrderReply)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) PayOrder(ctx context.Context, in *PayOrderRequest, opts ...grpc.CallOption) (*PayOrderReply, error) {
	out := new(PayOrderReply)
	err := c.cc.Invoke(ctx, OrderService_PayOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OrderInfo(ctx context.Context, in *OrderInfoRequest, opts ...grpc.CallOption) (*OrderInfoReply, error) {
	out := new(OrderInfoReply)
	err := c.cc.Invoke(ctx, OrderService_OrderInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) OrderList(ctx context.Context, in *OrderListRequest, opts ...grpc.CallOption) (*OrderListReply, error) {
	out := new(OrderListReply)
	err := c.cc.Invoke(ctx, OrderService_OrderList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	//确认订单
	ConfirmOrder(context.Context, *ConfirmOrderRequest) (*ConfirmOrderReply, error)
	//提交订单
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderReply, error)
	//支付订单
	PayOrder(context.Context, *PayOrderRequest) (*PayOrderReply, error)
	//订单详情
	OrderInfo(context.Context, *OrderInfoRequest) (*OrderInfoReply, error)
	//订单列表
	OrderList(context.Context, *OrderListRequest) (*OrderListReply, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) ConfirmOrder(context.Context, *ConfirmOrderRequest) (*ConfirmOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmOrder not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) PayOrder(context.Context, *PayOrderRequest) (*PayOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PayOrder not implemented")
}
func (UnimplementedOrderServiceServer) OrderInfo(context.Context, *OrderInfoRequest) (*OrderInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderInfo not implemented")
}
func (UnimplementedOrderServiceServer) OrderList(context.Context, *OrderListRequest) (*OrderListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderList not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_ConfirmOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ConfirmOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ConfirmOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ConfirmOrder(ctx, req.(*ConfirmOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_PayOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PayOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).PayOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_PayOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).PayOrder(ctx, req.(*PayOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OrderInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_OrderInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderInfo(ctx, req.(*OrderInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_OrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_OrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderList(ctx, req.(*OrderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConfirmOrder",
			Handler:    _OrderService_ConfirmOrder_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "PayOrder",
			Handler:    _OrderService_PayOrder_Handler,
		},
		{
			MethodName: "OrderInfo",
			Handler:    _OrderService_OrderInfo_Handler,
		},
		{
			MethodName: "OrderList",
			Handler:    _OrderService_OrderList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "demo/v1/order.proto",
}