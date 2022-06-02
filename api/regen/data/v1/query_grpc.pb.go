// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: regen/data/v1/query.proto

package datav1

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
	// ByIRI queries data based on IRI.
	ByIRI(ctx context.Context, in *QueryByIRIRequest, opts ...grpc.CallOption) (*QueryByIRIResponse, error)
	// ByHash queries data based on ContentHash.
	ByHash(ctx context.Context, in *QueryByHashRequest, opts ...grpc.CallOption) (*QueryByHashResponse, error)
	// ByAttestor queries data based on attestor.
	ByAttestor(ctx context.Context, in *QueryByAttestorRequest, opts ...grpc.CallOption) (*QueryByAttestorResponse, error)
	// IRIByHash queries IRI based on ContentHash.
	IRIByHash(ctx context.Context, in *QueryIRIByHashRequest, opts ...grpc.CallOption) (*QueryIRIByHashResponse, error)
	// IRIByRawHash queries IRI based on ContentHash_Raw properties.
	IRIByRawHash(ctx context.Context, in *QueryIRIByRawHashRequest, opts ...grpc.CallOption) (*QueryIRIByRawHashResponse, error)
	// IRIByGraphHash queries IRI based on ContentHash_Graph properties.
	IRIByGraphHash(ctx context.Context, in *QueryIRIByGraphHashRequest, opts ...grpc.CallOption) (*QueryIRIByGraphHashResponse, error)
	// HashByIRI queries ContentHash based on IRI.
	HashByIRI(ctx context.Context, in *QueryHashByIRIRequest, opts ...grpc.CallOption) (*QueryHashByIRIResponse, error)
	// AttestorsByIRI queries attestors based on IRI.
	AttestorsByIRI(ctx context.Context, in *QueryAttestorsByIRIRequest, opts ...grpc.CallOption) (*QueryAttestorsByIRIResponse, error)
	// AttestorsByHash queries attestors based on ContentHash.
	AttestorsByHash(ctx context.Context, in *QueryAttestorsByHashRequest, opts ...grpc.CallOption) (*QueryAttestorsByHashResponse, error)
	// Resolver queries information about a resolver based on ID.
	Resolver(ctx context.Context, in *QueryResolverRequest, opts ...grpc.CallOption) (*QueryResolverResponse, error)
	// ResolversByIRI queries resolvers based on IRI.
	ResolversByIRI(ctx context.Context, in *QueryResolversByIRIRequest, opts ...grpc.CallOption) (*QueryResolversByIRIResponse, error)
	// ResolversByHash queries resolvers based on ContentHash.
	ResolversByHash(ctx context.Context, in *QueryResolversByHashRequest, opts ...grpc.CallOption) (*QueryResolversByHashResponse, error)
	// ResolversByUrl queries resolvers based on URL.
	ResolversByUrl(ctx context.Context, in *QueryResolversByUrlRequest, opts ...grpc.CallOption) (*QueryResolversByUrlResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ByIRI(ctx context.Context, in *QueryByIRIRequest, opts ...grpc.CallOption) (*QueryByIRIResponse, error) {
	out := new(QueryByIRIResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/ByIRI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ByHash(ctx context.Context, in *QueryByHashRequest, opts ...grpc.CallOption) (*QueryByHashResponse, error) {
	out := new(QueryByHashResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/ByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ByAttestor(ctx context.Context, in *QueryByAttestorRequest, opts ...grpc.CallOption) (*QueryByAttestorResponse, error) {
	out := new(QueryByAttestorResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/ByAttestor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IRIByHash(ctx context.Context, in *QueryIRIByHashRequest, opts ...grpc.CallOption) (*QueryIRIByHashResponse, error) {
	out := new(QueryIRIByHashResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/IRIByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IRIByRawHash(ctx context.Context, in *QueryIRIByRawHashRequest, opts ...grpc.CallOption) (*QueryIRIByRawHashResponse, error) {
	out := new(QueryIRIByRawHashResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/IRIByRawHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IRIByGraphHash(ctx context.Context, in *QueryIRIByGraphHashRequest, opts ...grpc.CallOption) (*QueryIRIByGraphHashResponse, error) {
	out := new(QueryIRIByGraphHashResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/IRIByGraphHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HashByIRI(ctx context.Context, in *QueryHashByIRIRequest, opts ...grpc.CallOption) (*QueryHashByIRIResponse, error) {
	out := new(QueryHashByIRIResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/HashByIRI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AttestorsByIRI(ctx context.Context, in *QueryAttestorsByIRIRequest, opts ...grpc.CallOption) (*QueryAttestorsByIRIResponse, error) {
	out := new(QueryAttestorsByIRIResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/AttestorsByIRI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) AttestorsByHash(ctx context.Context, in *QueryAttestorsByHashRequest, opts ...grpc.CallOption) (*QueryAttestorsByHashResponse, error) {
	out := new(QueryAttestorsByHashResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/AttestorsByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Resolver(ctx context.Context, in *QueryResolverRequest, opts ...grpc.CallOption) (*QueryResolverResponse, error) {
	out := new(QueryResolverResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/Resolver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ResolversByIRI(ctx context.Context, in *QueryResolversByIRIRequest, opts ...grpc.CallOption) (*QueryResolversByIRIResponse, error) {
	out := new(QueryResolversByIRIResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/ResolversByIRI", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ResolversByHash(ctx context.Context, in *QueryResolversByHashRequest, opts ...grpc.CallOption) (*QueryResolversByHashResponse, error) {
	out := new(QueryResolversByHashResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/ResolversByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ResolversByUrl(ctx context.Context, in *QueryResolversByUrlRequest, opts ...grpc.CallOption) (*QueryResolversByUrlResponse, error) {
	out := new(QueryResolversByUrlResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1.Query/ResolversByUrl", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// ByIRI queries data based on IRI.
	ByIRI(context.Context, *QueryByIRIRequest) (*QueryByIRIResponse, error)
	// ByHash queries data based on ContentHash.
	ByHash(context.Context, *QueryByHashRequest) (*QueryByHashResponse, error)
	// ByAttestor queries data based on attestor.
	ByAttestor(context.Context, *QueryByAttestorRequest) (*QueryByAttestorResponse, error)
	// IRIByHash queries IRI based on ContentHash.
	IRIByHash(context.Context, *QueryIRIByHashRequest) (*QueryIRIByHashResponse, error)
	// IRIByRawHash queries IRI based on ContentHash_Raw properties.
	IRIByRawHash(context.Context, *QueryIRIByRawHashRequest) (*QueryIRIByRawHashResponse, error)
	// IRIByGraphHash queries IRI based on ContentHash_Graph properties.
	IRIByGraphHash(context.Context, *QueryIRIByGraphHashRequest) (*QueryIRIByGraphHashResponse, error)
	// HashByIRI queries ContentHash based on IRI.
	HashByIRI(context.Context, *QueryHashByIRIRequest) (*QueryHashByIRIResponse, error)
	// AttestorsByIRI queries attestors based on IRI.
	AttestorsByIRI(context.Context, *QueryAttestorsByIRIRequest) (*QueryAttestorsByIRIResponse, error)
	// AttestorsByHash queries attestors based on ContentHash.
	AttestorsByHash(context.Context, *QueryAttestorsByHashRequest) (*QueryAttestorsByHashResponse, error)
	// Resolver queries information about a resolver based on ID.
	Resolver(context.Context, *QueryResolverRequest) (*QueryResolverResponse, error)
	// ResolversByIRI queries resolvers based on IRI.
	ResolversByIRI(context.Context, *QueryResolversByIRIRequest) (*QueryResolversByIRIResponse, error)
	// ResolversByHash queries resolvers based on ContentHash.
	ResolversByHash(context.Context, *QueryResolversByHashRequest) (*QueryResolversByHashResponse, error)
	// ResolversByUrl queries resolvers based on URL.
	ResolversByUrl(context.Context, *QueryResolversByUrlRequest) (*QueryResolversByUrlResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) ByIRI(context.Context, *QueryByIRIRequest) (*QueryByIRIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByIRI not implemented")
}
func (UnimplementedQueryServer) ByHash(context.Context, *QueryByHashRequest) (*QueryByHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByHash not implemented")
}
func (UnimplementedQueryServer) ByAttestor(context.Context, *QueryByAttestorRequest) (*QueryByAttestorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ByAttestor not implemented")
}
func (UnimplementedQueryServer) IRIByHash(context.Context, *QueryIRIByHashRequest) (*QueryIRIByHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IRIByHash not implemented")
}
func (UnimplementedQueryServer) IRIByRawHash(context.Context, *QueryIRIByRawHashRequest) (*QueryIRIByRawHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IRIByRawHash not implemented")
}
func (UnimplementedQueryServer) IRIByGraphHash(context.Context, *QueryIRIByGraphHashRequest) (*QueryIRIByGraphHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IRIByGraphHash not implemented")
}
func (UnimplementedQueryServer) HashByIRI(context.Context, *QueryHashByIRIRequest) (*QueryHashByIRIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HashByIRI not implemented")
}
func (UnimplementedQueryServer) AttestorsByIRI(context.Context, *QueryAttestorsByIRIRequest) (*QueryAttestorsByIRIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttestorsByIRI not implemented")
}
func (UnimplementedQueryServer) AttestorsByHash(context.Context, *QueryAttestorsByHashRequest) (*QueryAttestorsByHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttestorsByHash not implemented")
}
func (UnimplementedQueryServer) Resolver(context.Context, *QueryResolverRequest) (*QueryResolverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resolver not implemented")
}
func (UnimplementedQueryServer) ResolversByIRI(context.Context, *QueryResolversByIRIRequest) (*QueryResolversByIRIResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolversByIRI not implemented")
}
func (UnimplementedQueryServer) ResolversByHash(context.Context, *QueryResolversByHashRequest) (*QueryResolversByHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolversByHash not implemented")
}
func (UnimplementedQueryServer) ResolversByUrl(context.Context, *QueryResolversByUrlRequest) (*QueryResolversByUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolversByUrl not implemented")
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

func _Query_ByIRI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryByIRIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ByIRI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/ByIRI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ByIRI(ctx, req.(*QueryByIRIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/ByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ByHash(ctx, req.(*QueryByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ByAttestor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryByAttestorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ByAttestor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/ByAttestor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ByAttestor(ctx, req.(*QueryByAttestorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IRIByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIRIByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IRIByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/IRIByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IRIByHash(ctx, req.(*QueryIRIByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IRIByRawHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIRIByRawHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IRIByRawHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/IRIByRawHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IRIByRawHash(ctx, req.(*QueryIRIByRawHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IRIByGraphHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIRIByGraphHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IRIByGraphHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/IRIByGraphHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IRIByGraphHash(ctx, req.(*QueryIRIByGraphHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HashByIRI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHashByIRIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HashByIRI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/HashByIRI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HashByIRI(ctx, req.(*QueryHashByIRIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AttestorsByIRI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAttestorsByIRIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AttestorsByIRI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/AttestorsByIRI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AttestorsByIRI(ctx, req.(*QueryAttestorsByIRIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_AttestorsByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAttestorsByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).AttestorsByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/AttestorsByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).AttestorsByHash(ctx, req.(*QueryAttestorsByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Resolver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryResolverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Resolver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/Resolver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Resolver(ctx, req.(*QueryResolverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ResolversByIRI_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryResolversByIRIRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ResolversByIRI(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/ResolversByIRI",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ResolversByIRI(ctx, req.(*QueryResolversByIRIRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ResolversByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryResolversByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ResolversByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/ResolversByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ResolversByHash(ctx, req.(*QueryResolversByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ResolversByUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryResolversByUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ResolversByUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1.Query/ResolversByUrl",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ResolversByUrl(ctx, req.(*QueryResolversByUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.data.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByIRI",
			Handler:    _Query_ByIRI_Handler,
		},
		{
			MethodName: "ByHash",
			Handler:    _Query_ByHash_Handler,
		},
		{
			MethodName: "ByAttestor",
			Handler:    _Query_ByAttestor_Handler,
		},
		{
			MethodName: "IRIByHash",
			Handler:    _Query_IRIByHash_Handler,
		},
		{
			MethodName: "IRIByRawHash",
			Handler:    _Query_IRIByRawHash_Handler,
		},
		{
			MethodName: "IRIByGraphHash",
			Handler:    _Query_IRIByGraphHash_Handler,
		},
		{
			MethodName: "HashByIRI",
			Handler:    _Query_HashByIRI_Handler,
		},
		{
			MethodName: "AttestorsByIRI",
			Handler:    _Query_AttestorsByIRI_Handler,
		},
		{
			MethodName: "AttestorsByHash",
			Handler:    _Query_AttestorsByHash_Handler,
		},
		{
			MethodName: "Resolver",
			Handler:    _Query_Resolver_Handler,
		},
		{
			MethodName: "ResolversByIRI",
			Handler:    _Query_ResolversByIRI_Handler,
		},
		{
			MethodName: "ResolversByHash",
			Handler:    _Query_ResolversByHash_Handler,
		},
		{
			MethodName: "ResolversByUrl",
			Handler:    _Query_ResolversByUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "regen/data/v1/query.proto",
}
