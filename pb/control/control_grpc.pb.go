// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: control.proto

package control_pb

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

const (
	Control_ClientCmdSrv_FullMethodName         = "/control_pb.Control/ClientCmdSrv"
	Control_CreateTableResultSrv_FullMethodName = "/control_pb.Control/CreateTableResultSrv"
	Control_QuitTableSrv_FullMethodName         = "/control_pb.Control/QuitTableSrv"
	Control_TableEndSrv_FullMethodName          = "/control_pb.Control/TableEndSrv"
)

// ControlClient is the client API for Control service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlClient interface {
	ClientCmdSrv(ctx context.Context, in *ControlMessage, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateTableResultSrv(ctx context.Context, in *CreateTableResult, opts ...grpc.CallOption) (*emptypb.Empty, error)
	QuitTableSrv(ctx context.Context, in *QuitTable, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TableEndSrv(ctx context.Context, in *TableEnd, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type controlClient struct {
	cc grpc.ClientConnInterface
}

func NewControlClient(cc grpc.ClientConnInterface) ControlClient {
	return &controlClient{cc}
}

func (c *controlClient) ClientCmdSrv(ctx context.Context, in *ControlMessage, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Control_ClientCmdSrv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) CreateTableResultSrv(ctx context.Context, in *CreateTableResult, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Control_CreateTableResultSrv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) QuitTableSrv(ctx context.Context, in *QuitTable, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Control_QuitTableSrv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlClient) TableEndSrv(ctx context.Context, in *TableEnd, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Control_TableEndSrv_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServer is the server API for Control service.
// All implementations must embed UnimplementedControlServer
// for forward compatibility
type ControlServer interface {
	ClientCmdSrv(context.Context, *ControlMessage) (*emptypb.Empty, error)
	CreateTableResultSrv(context.Context, *CreateTableResult) (*emptypb.Empty, error)
	QuitTableSrv(context.Context, *QuitTable) (*emptypb.Empty, error)
	TableEndSrv(context.Context, *TableEnd) (*emptypb.Empty, error)
	mustEmbedUnimplementedControlServer()
}

// UnimplementedControlServer must be embedded to have forward compatible implementations.
type UnimplementedControlServer struct {
}

func (UnimplementedControlServer) ClientCmdSrv(context.Context, *ControlMessage) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientCmdSrv not implemented")
}
func (UnimplementedControlServer) CreateTableResultSrv(context.Context, *CreateTableResult) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTableResultSrv not implemented")
}
func (UnimplementedControlServer) QuitTableSrv(context.Context, *QuitTable) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuitTableSrv not implemented")
}
func (UnimplementedControlServer) TableEndSrv(context.Context, *TableEnd) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TableEndSrv not implemented")
}
func (UnimplementedControlServer) mustEmbedUnimplementedControlServer() {}

// UnsafeControlServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlServer will
// result in compilation errors.
type UnsafeControlServer interface {
	mustEmbedUnimplementedControlServer()
}

func RegisterControlServer(s grpc.ServiceRegistrar, srv ControlServer) {
	s.RegisterService(&Control_ServiceDesc, srv)
}

func _Control_ClientCmdSrv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ControlMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).ClientCmdSrv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_ClientCmdSrv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).ClientCmdSrv(ctx, req.(*ControlMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_CreateTableResultSrv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableResult)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).CreateTableResultSrv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_CreateTableResultSrv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).CreateTableResultSrv(ctx, req.(*CreateTableResult))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_QuitTableSrv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuitTable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).QuitTableSrv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_QuitTableSrv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).QuitTableSrv(ctx, req.(*QuitTable))
	}
	return interceptor(ctx, in, info, handler)
}

func _Control_TableEndSrv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TableEnd)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServer).TableEndSrv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Control_TableEndSrv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServer).TableEndSrv(ctx, req.(*TableEnd))
	}
	return interceptor(ctx, in, info, handler)
}

// Control_ServiceDesc is the grpc.ServiceDesc for Control service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Control_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "control_pb.Control",
	HandlerType: (*ControlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientCmdSrv",
			Handler:    _Control_ClientCmdSrv_Handler,
		},
		{
			MethodName: "CreateTableResultSrv",
			Handler:    _Control_CreateTableResultSrv_Handler,
		},
		{
			MethodName: "QuitTableSrv",
			Handler:    _Control_QuitTableSrv_Handler,
		},
		{
			MethodName: "TableEndSrv",
			Handler:    _Control_TableEndSrv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "control.proto",
}
