// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: todos.proto

package models

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TodosClient is the client API for Todos service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodosClient interface {
	Add(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*MutationResponse, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllResponse, error)
	Delete(ctx context.Context, in *TodoId, opts ...grpc.CallOption) (*MutationResponse, error)
	Update(ctx context.Context, in *TodoUpdate, opts ...grpc.CallOption) (*MutationResponse, error)
}

type todosClient struct {
	cc grpc.ClientConnInterface
}

func NewTodosClient(cc grpc.ClientConnInterface) TodosClient {
	return &todosClient{cc}
}

func (c *todosClient) Add(ctx context.Context, in *Todo, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/model.Todos/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/model.Todos/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Delete(ctx context.Context, in *TodoId, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/model.Todos/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todosClient) Update(ctx context.Context, in *TodoUpdate, opts ...grpc.CallOption) (*MutationResponse, error) {
	out := new(MutationResponse)
	err := c.cc.Invoke(ctx, "/model.Todos/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodosServer is the server API for Todos service.
// All implementations must embed UnimplementedTodosServer
// for forward compatibility
type TodosServer interface {
	Add(context.Context, *Todo) (*MutationResponse, error)
	List(context.Context, *emptypb.Empty) (*GetAllResponse, error)
	Delete(context.Context, *TodoId) (*MutationResponse, error)
	Update(context.Context, *TodoUpdate) (*MutationResponse, error)
	mustEmbedUnimplementedTodosServer()
}

// UnimplementedTodosServer must be embedded to have forward compatible implementations.
type UnimplementedTodosServer struct {
}

func (UnimplementedTodosServer) Add(context.Context, *Todo) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedTodosServer) List(context.Context, *emptypb.Empty) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTodosServer) Delete(context.Context, *TodoId) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTodosServer) Update(context.Context, *TodoUpdate) (*MutationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTodosServer) mustEmbedUnimplementedTodosServer() {}

// UnsafeTodosServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodosServer will
// result in compilation errors.
type UnsafeTodosServer interface {
	mustEmbedUnimplementedTodosServer()
}

func RegisterTodosServer(s grpc.ServiceRegistrar, srv TodosServer) {
	s.RegisterService(&Todos_ServiceDesc, srv)
}

func _Todos_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Todo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Add(ctx, req.(*Todo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Delete(ctx, req.(*TodoId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Todos_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodoUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodosServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/model.Todos/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodosServer).Update(ctx, req.(*TodoUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

// Todos_ServiceDesc is the grpc.ServiceDesc for Todos service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Todos_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "model.Todos",
	HandlerType: (*TodosServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Todos_Add_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Todos_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Todos_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Todos_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todos.proto",
}
