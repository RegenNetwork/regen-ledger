// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: regen/ecocredit/basket/v1/query.proto

package basketv1

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
	Query_Basket_FullMethodName         = "/regen.ecocredit.basket.v1.Query/Basket"
	Query_Baskets_FullMethodName        = "/regen.ecocredit.basket.v1.Query/Baskets"
	Query_BasketBalances_FullMethodName = "/regen.ecocredit.basket.v1.Query/BasketBalances"
	Query_BasketBalance_FullMethodName  = "/regen.ecocredit.basket.v1.Query/BasketBalance"
	Query_BasketFee_FullMethodName      = "/regen.ecocredit.basket.v1.Query/BasketFee"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Basket queries one basket by denom.
	Basket(ctx context.Context, in *QueryBasketRequest, opts ...grpc.CallOption) (*QueryBasketResponse, error)
	// Baskets lists all baskets in the ecocredit module.
	Baskets(ctx context.Context, in *QueryBasketsRequest, opts ...grpc.CallOption) (*QueryBasketsResponse, error)
	// BasketBalances lists the balance of each credit batch in the basket.
	BasketBalances(ctx context.Context, in *QueryBasketBalancesRequest, opts ...grpc.CallOption) (*QueryBasketBalancesResponse, error)
	// BasketBalance queries the balance of a specific credit batch in the basket.
	BasketBalance(ctx context.Context, in *QueryBasketBalanceRequest, opts ...grpc.CallOption) (*QueryBasketBalanceResponse, error)
	// BasketFee returns the basket creation fee. If not set, a basket creation
	// fee is not required.
	//
	// Since Revision 2
	BasketFee(ctx context.Context, in *QueryBasketFeeRequest, opts ...grpc.CallOption) (*QueryBasketFeeResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Basket(ctx context.Context, in *QueryBasketRequest, opts ...grpc.CallOption) (*QueryBasketResponse, error) {
	out := new(QueryBasketResponse)
	err := c.cc.Invoke(ctx, Query_Basket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Baskets(ctx context.Context, in *QueryBasketsRequest, opts ...grpc.CallOption) (*QueryBasketsResponse, error) {
	out := new(QueryBasketsResponse)
	err := c.cc.Invoke(ctx, Query_Baskets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BasketBalances(ctx context.Context, in *QueryBasketBalancesRequest, opts ...grpc.CallOption) (*QueryBasketBalancesResponse, error) {
	out := new(QueryBasketBalancesResponse)
	err := c.cc.Invoke(ctx, Query_BasketBalances_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BasketBalance(ctx context.Context, in *QueryBasketBalanceRequest, opts ...grpc.CallOption) (*QueryBasketBalanceResponse, error) {
	out := new(QueryBasketBalanceResponse)
	err := c.cc.Invoke(ctx, Query_BasketBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BasketFee(ctx context.Context, in *QueryBasketFeeRequest, opts ...grpc.CallOption) (*QueryBasketFeeResponse, error) {
	out := new(QueryBasketFeeResponse)
	err := c.cc.Invoke(ctx, Query_BasketFee_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Basket queries one basket by denom.
	Basket(context.Context, *QueryBasketRequest) (*QueryBasketResponse, error)
	// Baskets lists all baskets in the ecocredit module.
	Baskets(context.Context, *QueryBasketsRequest) (*QueryBasketsResponse, error)
	// BasketBalances lists the balance of each credit batch in the basket.
	BasketBalances(context.Context, *QueryBasketBalancesRequest) (*QueryBasketBalancesResponse, error)
	// BasketBalance queries the balance of a specific credit batch in the basket.
	BasketBalance(context.Context, *QueryBasketBalanceRequest) (*QueryBasketBalanceResponse, error)
	// BasketFee returns the basket creation fee. If not set, a basket creation
	// fee is not required.
	//
	// Since Revision 2
	BasketFee(context.Context, *QueryBasketFeeRequest) (*QueryBasketFeeResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Basket(context.Context, *QueryBasketRequest) (*QueryBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Basket not implemented")
}
func (UnimplementedQueryServer) Baskets(context.Context, *QueryBasketsRequest) (*QueryBasketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Baskets not implemented")
}
func (UnimplementedQueryServer) BasketBalances(context.Context, *QueryBasketBalancesRequest) (*QueryBasketBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BasketBalances not implemented")
}
func (UnimplementedQueryServer) BasketBalance(context.Context, *QueryBasketBalanceRequest) (*QueryBasketBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BasketBalance not implemented")
}
func (UnimplementedQueryServer) BasketFee(context.Context, *QueryBasketFeeRequest) (*QueryBasketFeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BasketFee not implemented")
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

func _Query_Basket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBasketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Basket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Basket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Basket(ctx, req.(*QueryBasketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Baskets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBasketsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Baskets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Baskets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Baskets(ctx, req.(*QueryBasketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BasketBalances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBasketBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BasketBalances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BasketBalances_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BasketBalances(ctx, req.(*QueryBasketBalancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BasketBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBasketBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BasketBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BasketBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BasketBalance(ctx, req.(*QueryBasketBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BasketFee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBasketFeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BasketFee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_BasketFee_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BasketFee(ctx, req.(*QueryBasketFeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.ecocredit.basket.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Basket",
			Handler:    _Query_Basket_Handler,
		},
		{
			MethodName: "Baskets",
			Handler:    _Query_Baskets_Handler,
		},
		{
			MethodName: "BasketBalances",
			Handler:    _Query_BasketBalances_Handler,
		},
		{
			MethodName: "BasketBalance",
			Handler:    _Query_BasketBalance_Handler,
		},
		{
			MethodName: "BasketFee",
			Handler:    _Query_BasketFee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/ecocredit/basket/v1/query.proto",
}
