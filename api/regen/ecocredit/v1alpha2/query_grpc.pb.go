// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: regen/ecocredit/v1alpha2/query.proto

package ecocreditv1alpha2

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
	// Classes queries for all credit classes with pagination.
	Classes(ctx context.Context, in *QueryClassesRequest, opts ...grpc.CallOption) (*QueryClassesResponse, error)
	// ClassInfo queries for information on a credit class.
	ClassInfo(ctx context.Context, in *QueryClassInfoRequest, opts ...grpc.CallOption) (*QueryClassInfoResponse, error)
	// Projects queries for all projects within a class with pagination.
	Projects(ctx context.Context, in *QueryProjectsRequest, opts ...grpc.CallOption) (*QueryProjectsResponse, error)
	// ClassInfo queries for information on a project.
	ProjectInfo(ctx context.Context, in *QueryProjectInfoRequest, opts ...grpc.CallOption) (*QueryProjectInfoResponse, error)
	// Batches queries for all batches in the given project with pagination.
	Batches(ctx context.Context, in *QueryBatchesRequest, opts ...grpc.CallOption) (*QueryBatchesResponse, error)
	// BatchInfo queries for information on a credit batch.
	BatchInfo(ctx context.Context, in *QueryBatchInfoRequest, opts ...grpc.CallOption) (*QueryBatchInfoResponse, error)
	// Balance queries the balance (both tradable and retired) of a given credit
	// batch for a given account.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// Supply queries the tradable and retired supply of a credit batch.
	Supply(ctx context.Context, in *QuerySupplyRequest, opts ...grpc.CallOption) (*QuerySupplyResponse, error)
	// CreditTypes returns the list of allowed types that credit classes can have.
	// See Types/CreditType for more details.
	CreditTypes(ctx context.Context, in *QueryCreditTypesRequest, opts ...grpc.CallOption) (*QueryCreditTypesResponse, error)
	// Params queries the ecocredit module parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// SellOrder queries a sell order by its ID
	SellOrder(ctx context.Context, in *QuerySellOrderRequest, opts ...grpc.CallOption) (*QuerySellOrderResponse, error)
	// SellOrders queries a paginated list of all sell orders
	SellOrders(ctx context.Context, in *QuerySellOrdersRequest, opts ...grpc.CallOption) (*QuerySellOrdersResponse, error)
	// SellOrdersByDenom queries a paginated list of all sell orders of a specific ecocredit denom
	SellOrdersByBatchDenom(ctx context.Context, in *QuerySellOrdersByBatchDenomRequest, opts ...grpc.CallOption) (*QuerySellOrdersByBatchDenomResponse, error)
	// SellOrdersByAddress queries a paginated list of all sell orders from a specific address
	SellOrdersByAddress(ctx context.Context, in *QuerySellOrdersByAddressRequest, opts ...grpc.CallOption) (*QuerySellOrdersByAddressResponse, error)
	// BuyOrder queries a buy order by its id
	BuyOrder(ctx context.Context, in *QueryBuyOrderRequest, opts ...grpc.CallOption) (*QueryBuyOrderResponse, error)
	// BuyOrders queries a paginated list of all buy orders
	BuyOrders(ctx context.Context, in *QueryBuyOrdersRequest, opts ...grpc.CallOption) (*QueryBuyOrdersResponse, error)
	// BuyOrdersByAddress queries a paginated list of buy orders by creator address
	BuyOrdersByAddress(ctx context.Context, in *QueryBuyOrdersByAddressRequest, opts ...grpc.CallOption) (*QueryBuyOrdersByAddressResponse, error)
	// AllowedAskDenoms queries all denoms allowed to be set in the AskPrice of a sell order
	AllowedAskDenoms(ctx context.Context, in *QueryAllowedAskDenomsRequest, opts ...grpc.CallOption) (*QueryAllowedAskDenomsResponse, error)
	// Basket queries one basket by denom.
	Basket(ctx context.Context, in *QueryBasketRequest, opts ...grpc.CallOption) (*QueryBasketResponse, error)
	// Baskets lists all baskets in the ecocredit module.
	Baskets(ctx context.Context, in *QueryBasketsRequest, opts ...grpc.CallOption) (*QueryBasketsResponse, error)
	// BasketCredits lists all ecocredits inside a given basket.
	BasketCredits(ctx context.Context, in *QueryBasketCreditsRequest, opts ...grpc.CallOption) (*QueryBasketCreditsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Classes(ctx context.Context, in *QueryClassesRequest, opts ...grpc.CallOption) (*QueryClassesResponse, error) {
	out := new(QueryClassesResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/Classes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ClassInfo(ctx context.Context, in *QueryClassInfoRequest, opts ...grpc.CallOption) (*QueryClassInfoResponse, error) {
	out := new(QueryClassInfoResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/ClassInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Projects(ctx context.Context, in *QueryProjectsRequest, opts ...grpc.CallOption) (*QueryProjectsResponse, error) {
	out := new(QueryProjectsResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/Projects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProjectInfo(ctx context.Context, in *QueryProjectInfoRequest, opts ...grpc.CallOption) (*QueryProjectInfoResponse, error) {
	out := new(QueryProjectInfoResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/ProjectInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Batches(ctx context.Context, in *QueryBatchesRequest, opts ...grpc.CallOption) (*QueryBatchesResponse, error) {
	out := new(QueryBatchesResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/Batches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BatchInfo(ctx context.Context, in *QueryBatchInfoRequest, opts ...grpc.CallOption) (*QueryBatchInfoResponse, error) {
	out := new(QueryBatchInfoResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/BatchInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Supply(ctx context.Context, in *QuerySupplyRequest, opts ...grpc.CallOption) (*QuerySupplyResponse, error) {
	out := new(QuerySupplyResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/Supply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CreditTypes(ctx context.Context, in *QueryCreditTypesRequest, opts ...grpc.CallOption) (*QueryCreditTypesResponse, error) {
	out := new(QueryCreditTypesResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/CreditTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SellOrder(ctx context.Context, in *QuerySellOrderRequest, opts ...grpc.CallOption) (*QuerySellOrderResponse, error) {
	out := new(QuerySellOrderResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/SellOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SellOrders(ctx context.Context, in *QuerySellOrdersRequest, opts ...grpc.CallOption) (*QuerySellOrdersResponse, error) {
	out := new(QuerySellOrdersResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/SellOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SellOrdersByBatchDenom(ctx context.Context, in *QuerySellOrdersByBatchDenomRequest, opts ...grpc.CallOption) (*QuerySellOrdersByBatchDenomResponse, error) {
	out := new(QuerySellOrdersByBatchDenomResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/SellOrdersByBatchDenom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SellOrdersByAddress(ctx context.Context, in *QuerySellOrdersByAddressRequest, opts ...grpc.CallOption) (*QuerySellOrdersByAddressResponse, error) {
	out := new(QuerySellOrdersByAddressResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/SellOrdersByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BuyOrder(ctx context.Context, in *QueryBuyOrderRequest, opts ...grpc.CallOption) (*QueryBuyOrderResponse, error) {
	out := new(QueryBuyOrderResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/BuyOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BuyOrders(ctx context.Context, in *QueryBuyOrdersRequest, opts ...grpc.CallOption) (*QueryBuyOrdersResponse, error) {
	out := new(QueryBuyOrdersResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/BuyOrders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BuyOrdersByAddress(ctx context.Context, in *QueryBuyOrdersByAddressRequest, opts ...grpc.CallOption) (*QueryBuyOrdersByAddressResponse, error) {
	out := new(QueryBuyOrdersByAddressResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/BuyOrdersByAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AllowedAskDenoms(ctx context.Context, in *QueryAllowedAskDenomsRequest, opts ...grpc.CallOption) (*QueryAllowedAskDenomsResponse, error) {
	out := new(QueryAllowedAskDenomsResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/AllowedAskDenoms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Basket(ctx context.Context, in *QueryBasketRequest, opts ...grpc.CallOption) (*QueryBasketResponse, error) {
	out := new(QueryBasketResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/Basket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Baskets(ctx context.Context, in *QueryBasketsRequest, opts ...grpc.CallOption) (*QueryBasketsResponse, error) {
	out := new(QueryBasketsResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/Baskets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BasketCredits(ctx context.Context, in *QueryBasketCreditsRequest, opts ...grpc.CallOption) (*QueryBasketCreditsResponse, error) {
	out := new(QueryBasketCreditsResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1alpha2.Query/BasketCredits", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Classes queries for all credit classes with pagination.
	Classes(context.Context, *QueryClassesRequest) (*QueryClassesResponse, error)
	// ClassInfo queries for information on a credit class.
	ClassInfo(context.Context, *QueryClassInfoRequest) (*QueryClassInfoResponse, error)
	// Projects queries for all projects within a class with pagination.
	Projects(context.Context, *QueryProjectsRequest) (*QueryProjectsResponse, error)
	// ClassInfo queries for information on a project.
	ProjectInfo(context.Context, *QueryProjectInfoRequest) (*QueryProjectInfoResponse, error)
	// Batches queries for all batches in the given project with pagination.
	Batches(context.Context, *QueryBatchesRequest) (*QueryBatchesResponse, error)
	// BatchInfo queries for information on a credit batch.
	BatchInfo(context.Context, *QueryBatchInfoRequest) (*QueryBatchInfoResponse, error)
	// Balance queries the balance (both tradable and retired) of a given credit
	// batch for a given account.
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// Supply queries the tradable and retired supply of a credit batch.
	Supply(context.Context, *QuerySupplyRequest) (*QuerySupplyResponse, error)
	// CreditTypes returns the list of allowed types that credit classes can have.
	// See Types/CreditType for more details.
	CreditTypes(context.Context, *QueryCreditTypesRequest) (*QueryCreditTypesResponse, error)
	// Params queries the ecocredit module parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// SellOrder queries a sell order by its ID
	SellOrder(context.Context, *QuerySellOrderRequest) (*QuerySellOrderResponse, error)
	// SellOrders queries a paginated list of all sell orders
	SellOrders(context.Context, *QuerySellOrdersRequest) (*QuerySellOrdersResponse, error)
	// SellOrdersByDenom queries a paginated list of all sell orders of a specific ecocredit denom
	SellOrdersByBatchDenom(context.Context, *QuerySellOrdersByBatchDenomRequest) (*QuerySellOrdersByBatchDenomResponse, error)
	// SellOrdersByAddress queries a paginated list of all sell orders from a specific address
	SellOrdersByAddress(context.Context, *QuerySellOrdersByAddressRequest) (*QuerySellOrdersByAddressResponse, error)
	// BuyOrder queries a buy order by its id
	BuyOrder(context.Context, *QueryBuyOrderRequest) (*QueryBuyOrderResponse, error)
	// BuyOrders queries a paginated list of all buy orders
	BuyOrders(context.Context, *QueryBuyOrdersRequest) (*QueryBuyOrdersResponse, error)
	// BuyOrdersByAddress queries a paginated list of buy orders by creator address
	BuyOrdersByAddress(context.Context, *QueryBuyOrdersByAddressRequest) (*QueryBuyOrdersByAddressResponse, error)
	// AllowedAskDenoms queries all denoms allowed to be set in the AskPrice of a sell order
	AllowedAskDenoms(context.Context, *QueryAllowedAskDenomsRequest) (*QueryAllowedAskDenomsResponse, error)
	// Basket queries one basket by denom.
	Basket(context.Context, *QueryBasketRequest) (*QueryBasketResponse, error)
	// Baskets lists all baskets in the ecocredit module.
	Baskets(context.Context, *QueryBasketsRequest) (*QueryBasketsResponse, error)
	// BasketCredits lists all ecocredits inside a given basket.
	BasketCredits(context.Context, *QueryBasketCreditsRequest) (*QueryBasketCreditsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Classes(context.Context, *QueryClassesRequest) (*QueryClassesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Classes not implemented")
}
func (UnimplementedQueryServer) ClassInfo(context.Context, *QueryClassInfoRequest) (*QueryClassInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClassInfo not implemented")
}
func (UnimplementedQueryServer) Projects(context.Context, *QueryProjectsRequest) (*QueryProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Projects not implemented")
}
func (UnimplementedQueryServer) ProjectInfo(context.Context, *QueryProjectInfoRequest) (*QueryProjectInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectInfo not implemented")
}
func (UnimplementedQueryServer) Batches(context.Context, *QueryBatchesRequest) (*QueryBatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Batches not implemented")
}
func (UnimplementedQueryServer) BatchInfo(context.Context, *QueryBatchInfoRequest) (*QueryBatchInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchInfo not implemented")
}
func (UnimplementedQueryServer) Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedQueryServer) Supply(context.Context, *QuerySupplyRequest) (*QuerySupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Supply not implemented")
}
func (UnimplementedQueryServer) CreditTypes(context.Context, *QueryCreditTypesRequest) (*QueryCreditTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreditTypes not implemented")
}
func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) SellOrder(context.Context, *QuerySellOrderRequest) (*QuerySellOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOrder not implemented")
}
func (UnimplementedQueryServer) SellOrders(context.Context, *QuerySellOrdersRequest) (*QuerySellOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOrders not implemented")
}
func (UnimplementedQueryServer) SellOrdersByBatchDenom(context.Context, *QuerySellOrdersByBatchDenomRequest) (*QuerySellOrdersByBatchDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOrdersByBatchDenom not implemented")
}
func (UnimplementedQueryServer) SellOrdersByAddress(context.Context, *QuerySellOrdersByAddressRequest) (*QuerySellOrdersByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SellOrdersByAddress not implemented")
}
func (UnimplementedQueryServer) BuyOrder(context.Context, *QueryBuyOrderRequest) (*QueryBuyOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyOrder not implemented")
}
func (UnimplementedQueryServer) BuyOrders(context.Context, *QueryBuyOrdersRequest) (*QueryBuyOrdersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyOrders not implemented")
}
func (UnimplementedQueryServer) BuyOrdersByAddress(context.Context, *QueryBuyOrdersByAddressRequest) (*QueryBuyOrdersByAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyOrdersByAddress not implemented")
}
func (UnimplementedQueryServer) AllowedAskDenoms(context.Context, *QueryAllowedAskDenomsRequest) (*QueryAllowedAskDenomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllowedAskDenoms not implemented")
}
func (UnimplementedQueryServer) Basket(context.Context, *QueryBasketRequest) (*QueryBasketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Basket not implemented")
}
func (UnimplementedQueryServer) Baskets(context.Context, *QueryBasketsRequest) (*QueryBasketsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Baskets not implemented")
}
func (UnimplementedQueryServer) BasketCredits(context.Context, *QueryBasketCreditsRequest) (*QueryBasketCreditsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BasketCredits not implemented")
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

func _Query_Classes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClassesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Classes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/Classes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Classes(ctx, req.(*QueryClassesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ClassInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClassInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ClassInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/ClassInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ClassInfo(ctx, req.(*QueryClassInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Projects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Projects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/Projects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Projects(ctx, req.(*QueryProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProjectInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProjectInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProjectInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/ProjectInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProjectInfo(ctx, req.(*QueryProjectInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Batches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Batches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/Batches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Batches(ctx, req.(*QueryBatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BatchInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBatchInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BatchInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/BatchInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BatchInfo(ctx, req.(*QueryBatchInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Balance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Supply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Supply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/Supply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Supply(ctx, req.(*QuerySupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CreditTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryCreditTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CreditTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/CreditTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CreditTypes(ctx, req.(*QueryCreditTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SellOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySellOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SellOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/SellOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SellOrder(ctx, req.(*QuerySellOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SellOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySellOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SellOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/SellOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SellOrders(ctx, req.(*QuerySellOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SellOrdersByBatchDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySellOrdersByBatchDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SellOrdersByBatchDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/SellOrdersByBatchDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SellOrdersByBatchDenom(ctx, req.(*QuerySellOrdersByBatchDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SellOrdersByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySellOrdersByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SellOrdersByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/SellOrdersByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SellOrdersByAddress(ctx, req.(*QuerySellOrdersByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BuyOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBuyOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BuyOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/BuyOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BuyOrder(ctx, req.(*QueryBuyOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BuyOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBuyOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BuyOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/BuyOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BuyOrders(ctx, req.(*QueryBuyOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BuyOrdersByAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBuyOrdersByAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BuyOrdersByAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/BuyOrdersByAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BuyOrdersByAddress(ctx, req.(*QueryBuyOrdersByAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AllowedAskDenoms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllowedAskDenomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AllowedAskDenoms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/AllowedAskDenoms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AllowedAskDenoms(ctx, req.(*QueryAllowedAskDenomsRequest))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/regen.ecocredit.v1alpha2.Query/Basket",
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
		FullMethod: "/regen.ecocredit.v1alpha2.Query/Baskets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Baskets(ctx, req.(*QueryBasketsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BasketCredits_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBasketCreditsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BasketCredits(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1alpha2.Query/BasketCredits",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BasketCredits(ctx, req.(*QueryBasketCreditsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.ecocredit.v1alpha2.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Classes",
			Handler:    _Query_Classes_Handler,
		},
		{
			MethodName: "ClassInfo",
			Handler:    _Query_ClassInfo_Handler,
		},
		{
			MethodName: "Projects",
			Handler:    _Query_Projects_Handler,
		},
		{
			MethodName: "ProjectInfo",
			Handler:    _Query_ProjectInfo_Handler,
		},
		{
			MethodName: "Batches",
			Handler:    _Query_Batches_Handler,
		},
		{
			MethodName: "BatchInfo",
			Handler:    _Query_BatchInfo_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "Supply",
			Handler:    _Query_Supply_Handler,
		},
		{
			MethodName: "CreditTypes",
			Handler:    _Query_CreditTypes_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SellOrder",
			Handler:    _Query_SellOrder_Handler,
		},
		{
			MethodName: "SellOrders",
			Handler:    _Query_SellOrders_Handler,
		},
		{
			MethodName: "SellOrdersByBatchDenom",
			Handler:    _Query_SellOrdersByBatchDenom_Handler,
		},
		{
			MethodName: "SellOrdersByAddress",
			Handler:    _Query_SellOrdersByAddress_Handler,
		},
		{
			MethodName: "BuyOrder",
			Handler:    _Query_BuyOrder_Handler,
		},
		{
			MethodName: "BuyOrders",
			Handler:    _Query_BuyOrders_Handler,
		},
		{
			MethodName: "BuyOrdersByAddress",
			Handler:    _Query_BuyOrdersByAddress_Handler,
		},
		{
			MethodName: "AllowedAskDenoms",
			Handler:    _Query_AllowedAskDenoms_Handler,
		},
		{
			MethodName: "Basket",
			Handler:    _Query_Basket_Handler,
		},
		{
			MethodName: "Baskets",
			Handler:    _Query_Baskets_Handler,
		},
		{
			MethodName: "BasketCredits",
			Handler:    _Query_BasketCredits_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/ecocredit/v1alpha2/query.proto",
}