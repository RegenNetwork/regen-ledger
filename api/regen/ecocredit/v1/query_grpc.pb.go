// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: regen/ecocredit/v1/query.proto

package ecocreditv1

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
	// ClassesByAdmin queries for all credit classes with a specific admin
	// address.
	ClassesByAdmin(ctx context.Context, in *QueryClassesByAdminRequest, opts ...grpc.CallOption) (*QueryClassesByAdminResponse, error)
	// Class queries for information on a credit class.
	Class(ctx context.Context, in *QueryClassRequest, opts ...grpc.CallOption) (*QueryClassResponse, error)
	// ClassIssuers queries for the addresses of the issuers for a credit class.
	ClassIssuers(ctx context.Context, in *QueryClassIssuersRequest, opts ...grpc.CallOption) (*QueryClassIssuersResponse, error)
	// Projects queries for all projects within a class with pagination.
	Projects(ctx context.Context, in *QueryProjectsRequest, opts ...grpc.CallOption) (*QueryProjectsResponse, error)
	// ProjectsByReferenceId queries for all projects by reference-id with
	// pagination.
	ProjectsByReferenceId(ctx context.Context, in *QueryProjectsByReferenceIdRequest, opts ...grpc.CallOption) (*QueryProjectsByReferenceIdResponse, error)
	// Project queries for information on a project.
	Project(ctx context.Context, in *QueryProjectRequest, opts ...grpc.CallOption) (*QueryProjectResponse, error)
	// ProjectsByAdmin queries for all projects by admin with
	// pagination.
	ProjectsByAdmin(ctx context.Context, in *QueryProjectsByAdminRequest, opts ...grpc.CallOption) (*QueryProjectsByAdminResponse, error)
	// Batches queries for all batches with pagination.
	Batches(ctx context.Context, in *QueryBatchesRequest, opts ...grpc.CallOption) (*QueryBatchesResponse, error)
	// BatchesByIssuer queries all batches issued from a given issuer address.
	BatchesByIssuer(ctx context.Context, in *QueryBatchesByIssuerRequest, opts ...grpc.CallOption) (*QueryBatchesByIssuerResponse, error)
	// BatchesByClass queries all batches issued from a given class.
	BatchesByClass(ctx context.Context, in *QueryBatchesByClassRequest, opts ...grpc.CallOption) (*QueryBatchesByClassResponse, error)
	// BatchesByProject queries for all batches from a given project with
	// pagination.
	BatchesByProject(ctx context.Context, in *QueryBatchesByProjectRequest, opts ...grpc.CallOption) (*QueryBatchesByProjectResponse, error)
	// Batch queries for information on a credit batch.
	Batch(ctx context.Context, in *QueryBatchRequest, opts ...grpc.CallOption) (*QueryBatchResponse, error)
	// Balance queries the balance (both tradable and retired) of a given credit
	// batch for a given account address.
	Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error)
	// Balances queries all credit balances the given account holds.
	Balances(ctx context.Context, in *QueryBalancesRequest, opts ...grpc.CallOption) (*QueryBalancesResponse, error)
	// Supply queries the tradable and retired supply of a credit batch.
	Supply(ctx context.Context, in *QuerySupplyRequest, opts ...grpc.CallOption) (*QuerySupplyResponse, error)
	// CreditTypes returns the list of allowed types that credit classes can have.
	// See Types/CreditType for more details.
	CreditTypes(ctx context.Context, in *QueryCreditTypesRequest, opts ...grpc.CallOption) (*QueryCreditTypesResponse, error)
	// Params queries the ecocredit module parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Classes(ctx context.Context, in *QueryClassesRequest, opts ...grpc.CallOption) (*QueryClassesResponse, error) {
	out := new(QueryClassesResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/Classes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ClassesByAdmin(ctx context.Context, in *QueryClassesByAdminRequest, opts ...grpc.CallOption) (*QueryClassesByAdminResponse, error) {
	out := new(QueryClassesByAdminResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/ClassesByAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Class(ctx context.Context, in *QueryClassRequest, opts ...grpc.CallOption) (*QueryClassResponse, error) {
	out := new(QueryClassResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/Class", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ClassIssuers(ctx context.Context, in *QueryClassIssuersRequest, opts ...grpc.CallOption) (*QueryClassIssuersResponse, error) {
	out := new(QueryClassIssuersResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/ClassIssuers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Projects(ctx context.Context, in *QueryProjectsRequest, opts ...grpc.CallOption) (*QueryProjectsResponse, error) {
	out := new(QueryProjectsResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/Projects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProjectsByReferenceId(ctx context.Context, in *QueryProjectsByReferenceIdRequest, opts ...grpc.CallOption) (*QueryProjectsByReferenceIdResponse, error) {
	out := new(QueryProjectsByReferenceIdResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/ProjectsByReferenceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Project(ctx context.Context, in *QueryProjectRequest, opts ...grpc.CallOption) (*QueryProjectResponse, error) {
	out := new(QueryProjectResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/Project", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ProjectsByAdmin(ctx context.Context, in *QueryProjectsByAdminRequest, opts ...grpc.CallOption) (*QueryProjectsByAdminResponse, error) {
	out := new(QueryProjectsByAdminResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/ProjectsByAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Batches(ctx context.Context, in *QueryBatchesRequest, opts ...grpc.CallOption) (*QueryBatchesResponse, error) {
	out := new(QueryBatchesResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/Batches", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BatchesByIssuer(ctx context.Context, in *QueryBatchesByIssuerRequest, opts ...grpc.CallOption) (*QueryBatchesByIssuerResponse, error) {
	out := new(QueryBatchesByIssuerResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/BatchesByIssuer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BatchesByClass(ctx context.Context, in *QueryBatchesByClassRequest, opts ...grpc.CallOption) (*QueryBatchesByClassResponse, error) {
	out := new(QueryBatchesByClassResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/BatchesByClass", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BatchesByProject(ctx context.Context, in *QueryBatchesByProjectRequest, opts ...grpc.CallOption) (*QueryBatchesByProjectResponse, error) {
	out := new(QueryBatchesByProjectResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/BatchesByProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Batch(ctx context.Context, in *QueryBatchRequest, opts ...grpc.CallOption) (*QueryBatchResponse, error) {
	out := new(QueryBatchResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/Batch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Balance(ctx context.Context, in *QueryBalanceRequest, opts ...grpc.CallOption) (*QueryBalanceResponse, error) {
	out := new(QueryBalanceResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/Balance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Balances(ctx context.Context, in *QueryBalancesRequest, opts ...grpc.CallOption) (*QueryBalancesResponse, error) {
	out := new(QueryBalancesResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/Balances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Supply(ctx context.Context, in *QuerySupplyRequest, opts ...grpc.CallOption) (*QuerySupplyResponse, error) {
	out := new(QuerySupplyResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/Supply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CreditTypes(ctx context.Context, in *QueryCreditTypesRequest, opts ...grpc.CallOption) (*QueryCreditTypesResponse, error) {
	out := new(QueryCreditTypesResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/CreditTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/regen.ecocredit.v1.Query/Params", in, out, opts...)
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
	// ClassesByAdmin queries for all credit classes with a specific admin
	// address.
	ClassesByAdmin(context.Context, *QueryClassesByAdminRequest) (*QueryClassesByAdminResponse, error)
	// Class queries for information on a credit class.
	Class(context.Context, *QueryClassRequest) (*QueryClassResponse, error)
	// ClassIssuers queries for the addresses of the issuers for a credit class.
	ClassIssuers(context.Context, *QueryClassIssuersRequest) (*QueryClassIssuersResponse, error)
	// Projects queries for all projects within a class with pagination.
	Projects(context.Context, *QueryProjectsRequest) (*QueryProjectsResponse, error)
	// ProjectsByReferenceId queries for all projects by reference-id with
	// pagination.
	ProjectsByReferenceId(context.Context, *QueryProjectsByReferenceIdRequest) (*QueryProjectsByReferenceIdResponse, error)
	// Project queries for information on a project.
	Project(context.Context, *QueryProjectRequest) (*QueryProjectResponse, error)
	// ProjectsByAdmin queries for all projects by admin with
	// pagination.
	ProjectsByAdmin(context.Context, *QueryProjectsByAdminRequest) (*QueryProjectsByAdminResponse, error)
	// Batches queries for all batches with pagination.
	Batches(context.Context, *QueryBatchesRequest) (*QueryBatchesResponse, error)
	// BatchesByIssuer queries all batches issued from a given issuer address.
	BatchesByIssuer(context.Context, *QueryBatchesByIssuerRequest) (*QueryBatchesByIssuerResponse, error)
	// BatchesByClass queries all batches issued from a given class.
	BatchesByClass(context.Context, *QueryBatchesByClassRequest) (*QueryBatchesByClassResponse, error)
	// BatchesByProject queries for all batches from a given project with
	// pagination.
	BatchesByProject(context.Context, *QueryBatchesByProjectRequest) (*QueryBatchesByProjectResponse, error)
	// Batch queries for information on a credit batch.
	Batch(context.Context, *QueryBatchRequest) (*QueryBatchResponse, error)
	// Balance queries the balance (both tradable and retired) of a given credit
	// batch for a given account address.
	Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error)
	// Balances queries all credit balances the given account holds.
	Balances(context.Context, *QueryBalancesRequest) (*QueryBalancesResponse, error)
	// Supply queries the tradable and retired supply of a credit batch.
	Supply(context.Context, *QuerySupplyRequest) (*QuerySupplyResponse, error)
	// CreditTypes returns the list of allowed types that credit classes can have.
	// See Types/CreditType for more details.
	CreditTypes(context.Context, *QueryCreditTypesRequest) (*QueryCreditTypesResponse, error)
	// Params queries the ecocredit module parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Classes(context.Context, *QueryClassesRequest) (*QueryClassesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Classes not implemented")
}
func (UnimplementedQueryServer) ClassesByAdmin(context.Context, *QueryClassesByAdminRequest) (*QueryClassesByAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClassesByAdmin not implemented")
}
func (UnimplementedQueryServer) Class(context.Context, *QueryClassRequest) (*QueryClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Class not implemented")
}
func (UnimplementedQueryServer) ClassIssuers(context.Context, *QueryClassIssuersRequest) (*QueryClassIssuersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClassIssuers not implemented")
}
func (UnimplementedQueryServer) Projects(context.Context, *QueryProjectsRequest) (*QueryProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Projects not implemented")
}
func (UnimplementedQueryServer) ProjectsByReferenceId(context.Context, *QueryProjectsByReferenceIdRequest) (*QueryProjectsByReferenceIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectsByReferenceId not implemented")
}
func (UnimplementedQueryServer) Project(context.Context, *QueryProjectRequest) (*QueryProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Project not implemented")
}
func (UnimplementedQueryServer) ProjectsByAdmin(context.Context, *QueryProjectsByAdminRequest) (*QueryProjectsByAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProjectsByAdmin not implemented")
}
func (UnimplementedQueryServer) Batches(context.Context, *QueryBatchesRequest) (*QueryBatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Batches not implemented")
}
func (UnimplementedQueryServer) BatchesByIssuer(context.Context, *QueryBatchesByIssuerRequest) (*QueryBatchesByIssuerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchesByIssuer not implemented")
}
func (UnimplementedQueryServer) BatchesByClass(context.Context, *QueryBatchesByClassRequest) (*QueryBatchesByClassResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchesByClass not implemented")
}
func (UnimplementedQueryServer) BatchesByProject(context.Context, *QueryBatchesByProjectRequest) (*QueryBatchesByProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchesByProject not implemented")
}
func (UnimplementedQueryServer) Batch(context.Context, *QueryBatchRequest) (*QueryBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Batch not implemented")
}
func (UnimplementedQueryServer) Balance(context.Context, *QueryBalanceRequest) (*QueryBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balance not implemented")
}
func (UnimplementedQueryServer) Balances(context.Context, *QueryBalancesRequest) (*QueryBalancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Balances not implemented")
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
		FullMethod: "/regen.ecocredit.v1.Query/Classes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Classes(ctx, req.(*QueryClassesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ClassesByAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClassesByAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ClassesByAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/ClassesByAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ClassesByAdmin(ctx, req.(*QueryClassesByAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Class_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Class(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/Class",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Class(ctx, req.(*QueryClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ClassIssuers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClassIssuersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ClassIssuers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/ClassIssuers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ClassIssuers(ctx, req.(*QueryClassIssuersRequest))
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
		FullMethod: "/regen.ecocredit.v1.Query/Projects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Projects(ctx, req.(*QueryProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProjectsByReferenceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProjectsByReferenceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProjectsByReferenceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/ProjectsByReferenceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProjectsByReferenceId(ctx, req.(*QueryProjectsByReferenceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Project_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Project(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/Project",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Project(ctx, req.(*QueryProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ProjectsByAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryProjectsByAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ProjectsByAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/ProjectsByAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ProjectsByAdmin(ctx, req.(*QueryProjectsByAdminRequest))
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
		FullMethod: "/regen.ecocredit.v1.Query/Batches",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Batches(ctx, req.(*QueryBatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BatchesByIssuer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBatchesByIssuerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BatchesByIssuer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/BatchesByIssuer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BatchesByIssuer(ctx, req.(*QueryBatchesByIssuerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BatchesByClass_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBatchesByClassRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BatchesByClass(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/BatchesByClass",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BatchesByClass(ctx, req.(*QueryBatchesByClassRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BatchesByProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBatchesByProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BatchesByProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/BatchesByProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BatchesByProject(ctx, req.(*QueryBatchesByProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Batch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Batch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/Batch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Batch(ctx, req.(*QueryBatchRequest))
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
		FullMethod: "/regen.ecocredit.v1.Query/Balance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balance(ctx, req.(*QueryBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Balances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBalancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Balances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.ecocredit.v1.Query/Balances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Balances(ctx, req.(*QueryBalancesRequest))
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
		FullMethod: "/regen.ecocredit.v1.Query/Supply",
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
		FullMethod: "/regen.ecocredit.v1.Query/CreditTypes",
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
		FullMethod: "/regen.ecocredit.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.ecocredit.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Classes",
			Handler:    _Query_Classes_Handler,
		},
		{
			MethodName: "ClassesByAdmin",
			Handler:    _Query_ClassesByAdmin_Handler,
		},
		{
			MethodName: "Class",
			Handler:    _Query_Class_Handler,
		},
		{
			MethodName: "ClassIssuers",
			Handler:    _Query_ClassIssuers_Handler,
		},
		{
			MethodName: "Projects",
			Handler:    _Query_Projects_Handler,
		},
		{
			MethodName: "ProjectsByReferenceId",
			Handler:    _Query_ProjectsByReferenceId_Handler,
		},
		{
			MethodName: "Project",
			Handler:    _Query_Project_Handler,
		},
		{
			MethodName: "ProjectsByAdmin",
			Handler:    _Query_ProjectsByAdmin_Handler,
		},
		{
			MethodName: "Batches",
			Handler:    _Query_Batches_Handler,
		},
		{
			MethodName: "BatchesByIssuer",
			Handler:    _Query_BatchesByIssuer_Handler,
		},
		{
			MethodName: "BatchesByClass",
			Handler:    _Query_BatchesByClass_Handler,
		},
		{
			MethodName: "BatchesByProject",
			Handler:    _Query_BatchesByProject_Handler,
		},
		{
			MethodName: "Batch",
			Handler:    _Query_Batch_Handler,
		},
		{
			MethodName: "Balance",
			Handler:    _Query_Balance_Handler,
		},
		{
			MethodName: "Balances",
			Handler:    _Query_Balances_Handler,
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/ecocredit/v1/query.proto",
}
