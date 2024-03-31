// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: sim.proto

package SimHelper

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

// SimActiveClient is the client API for SimActive service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimActiveClient interface {
	// ///////////////////////////////////
	// //////// METHODS FOR Sim //////////
	// ///////////////////////////////////
	AddSim(ctx context.Context, in *AddSimRequest, opts ...grpc.CallOption) (*AddSimResponse, error)
	DeleteSim(ctx context.Context, in *DeleteSimRequest, opts ...grpc.CallOption) (*DeleteSimResponse, error)
	ActivateSim(ctx context.Context, in *ActivateSimRequest, opts ...grpc.CallOption) (*ActivateSimResponse, error)
	SetSimBlocked(ctx context.Context, in *SSBRequest, opts ...grpc.CallOption) (*SSBResponse, error)
	GetSimList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SimList, error)
	GetFreeServices(ctx context.Context, in *GetFreeServRequest, opts ...grpc.CallOption) (*GetFreeServResponse, error)
	GetUsedServices(ctx context.Context, in *GetUsedServRequest, opts ...grpc.CallOption) (*GetUsedServResponse, error)
	// ///////////////////////////////////
	// //////// METHODS FOR Services /////
	// ///////////////////////////////////
	AddService(ctx context.Context, in *AddServiceRequest, opts ...grpc.CallOption) (*AddServiceResponse, error)
	DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error)
	GetAllServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GASResponse, error)
	// ////////////////////////////////////
	// //////// METHODS FOR Used //////////
	// ////////////////////////////////////
	UseSimForService(ctx context.Context, in *USFSRequest, opts ...grpc.CallOption) (*USFSResponse, error)
	// ////////////////////////////////////
	// //////// METHODS FOR Provider //////
	// ////////////////////////////////////
	GetProviderList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProviderList, error)
}

type simActiveClient struct {
	cc grpc.ClientConnInterface
}

func NewSimActiveClient(cc grpc.ClientConnInterface) SimActiveClient {
	return &simActiveClient{cc}
}

func (c *simActiveClient) AddSim(ctx context.Context, in *AddSimRequest, opts ...grpc.CallOption) (*AddSimResponse, error) {
	out := new(AddSimResponse)
	err := c.cc.Invoke(ctx, "/SimActive/AddSim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) DeleteSim(ctx context.Context, in *DeleteSimRequest, opts ...grpc.CallOption) (*DeleteSimResponse, error) {
	out := new(DeleteSimResponse)
	err := c.cc.Invoke(ctx, "/SimActive/DeleteSim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) ActivateSim(ctx context.Context, in *ActivateSimRequest, opts ...grpc.CallOption) (*ActivateSimResponse, error) {
	out := new(ActivateSimResponse)
	err := c.cc.Invoke(ctx, "/SimActive/ActivateSim", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) SetSimBlocked(ctx context.Context, in *SSBRequest, opts ...grpc.CallOption) (*SSBResponse, error) {
	out := new(SSBResponse)
	err := c.cc.Invoke(ctx, "/SimActive/SetSimBlocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) GetSimList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*SimList, error) {
	out := new(SimList)
	err := c.cc.Invoke(ctx, "/SimActive/GetSimList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) GetFreeServices(ctx context.Context, in *GetFreeServRequest, opts ...grpc.CallOption) (*GetFreeServResponse, error) {
	out := new(GetFreeServResponse)
	err := c.cc.Invoke(ctx, "/SimActive/GetFreeServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) GetUsedServices(ctx context.Context, in *GetUsedServRequest, opts ...grpc.CallOption) (*GetUsedServResponse, error) {
	out := new(GetUsedServResponse)
	err := c.cc.Invoke(ctx, "/SimActive/GetUsedServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) AddService(ctx context.Context, in *AddServiceRequest, opts ...grpc.CallOption) (*AddServiceResponse, error) {
	out := new(AddServiceResponse)
	err := c.cc.Invoke(ctx, "/SimActive/AddService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) DeleteService(ctx context.Context, in *DeleteServiceRequest, opts ...grpc.CallOption) (*DeleteServiceResponse, error) {
	out := new(DeleteServiceResponse)
	err := c.cc.Invoke(ctx, "/SimActive/DeleteService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) GetAllServices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GASResponse, error) {
	out := new(GASResponse)
	err := c.cc.Invoke(ctx, "/SimActive/GetAllServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) UseSimForService(ctx context.Context, in *USFSRequest, opts ...grpc.CallOption) (*USFSResponse, error) {
	out := new(USFSResponse)
	err := c.cc.Invoke(ctx, "/SimActive/UseSimForService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simActiveClient) GetProviderList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ProviderList, error) {
	out := new(ProviderList)
	err := c.cc.Invoke(ctx, "/SimActive/GetProviderList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimActiveServer is the server API for SimActive service.
// All implementations must embed UnimplementedSimActiveServer
// for forward compatibility
type SimActiveServer interface {
	// ///////////////////////////////////
	// //////// METHODS FOR Sim //////////
	// ///////////////////////////////////
	AddSim(context.Context, *AddSimRequest) (*AddSimResponse, error)
	DeleteSim(context.Context, *DeleteSimRequest) (*DeleteSimResponse, error)
	ActivateSim(context.Context, *ActivateSimRequest) (*ActivateSimResponse, error)
	SetSimBlocked(context.Context, *SSBRequest) (*SSBResponse, error)
	GetSimList(context.Context, *Empty) (*SimList, error)
	GetFreeServices(context.Context, *GetFreeServRequest) (*GetFreeServResponse, error)
	GetUsedServices(context.Context, *GetUsedServRequest) (*GetUsedServResponse, error)
	// ///////////////////////////////////
	// //////// METHODS FOR Services /////
	// ///////////////////////////////////
	AddService(context.Context, *AddServiceRequest) (*AddServiceResponse, error)
	DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error)
	GetAllServices(context.Context, *Empty) (*GASResponse, error)
	// ////////////////////////////////////
	// //////// METHODS FOR Used //////////
	// ////////////////////////////////////
	UseSimForService(context.Context, *USFSRequest) (*USFSResponse, error)
	// ////////////////////////////////////
	// //////// METHODS FOR Provider //////
	// ////////////////////////////////////
	GetProviderList(context.Context, *Empty) (*ProviderList, error)
	mustEmbedUnimplementedSimActiveServer()
}

// UnimplementedSimActiveServer must be embedded to have forward compatible implementations.
type UnimplementedSimActiveServer struct {
}

func (UnimplementedSimActiveServer) AddSim(context.Context, *AddSimRequest) (*AddSimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSim not implemented")
}
func (UnimplementedSimActiveServer) DeleteSim(context.Context, *DeleteSimRequest) (*DeleteSimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSim not implemented")
}
func (UnimplementedSimActiveServer) ActivateSim(context.Context, *ActivateSimRequest) (*ActivateSimResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivateSim not implemented")
}
func (UnimplementedSimActiveServer) SetSimBlocked(context.Context, *SSBRequest) (*SSBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSimBlocked not implemented")
}
func (UnimplementedSimActiveServer) GetSimList(context.Context, *Empty) (*SimList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSimList not implemented")
}
func (UnimplementedSimActiveServer) GetFreeServices(context.Context, *GetFreeServRequest) (*GetFreeServResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreeServices not implemented")
}
func (UnimplementedSimActiveServer) GetUsedServices(context.Context, *GetUsedServRequest) (*GetUsedServResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsedServices not implemented")
}
func (UnimplementedSimActiveServer) AddService(context.Context, *AddServiceRequest) (*AddServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddService not implemented")
}
func (UnimplementedSimActiveServer) DeleteService(context.Context, *DeleteServiceRequest) (*DeleteServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteService not implemented")
}
func (UnimplementedSimActiveServer) GetAllServices(context.Context, *Empty) (*GASResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllServices not implemented")
}
func (UnimplementedSimActiveServer) UseSimForService(context.Context, *USFSRequest) (*USFSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UseSimForService not implemented")
}
func (UnimplementedSimActiveServer) GetProviderList(context.Context, *Empty) (*ProviderList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProviderList not implemented")
}
func (UnimplementedSimActiveServer) mustEmbedUnimplementedSimActiveServer() {}

// UnsafeSimActiveServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimActiveServer will
// result in compilation errors.
type UnsafeSimActiveServer interface {
	mustEmbedUnimplementedSimActiveServer()
}

func RegisterSimActiveServer(s grpc.ServiceRegistrar, srv SimActiveServer) {
	s.RegisterService(&SimActive_ServiceDesc, srv)
}

func _SimActive_AddSim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).AddSim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/AddSim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).AddSim(ctx, req.(*AddSimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_DeleteSim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).DeleteSim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/DeleteSim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).DeleteSim(ctx, req.(*DeleteSimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_ActivateSim_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateSimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).ActivateSim(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/ActivateSim",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).ActivateSim(ctx, req.(*ActivateSimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_SetSimBlocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).SetSimBlocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/SetSimBlocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).SetSimBlocked(ctx, req.(*SSBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_GetSimList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).GetSimList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/GetSimList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).GetSimList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_GetFreeServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFreeServRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).GetFreeServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/GetFreeServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).GetFreeServices(ctx, req.(*GetFreeServRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_GetUsedServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsedServRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).GetUsedServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/GetUsedServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).GetUsedServices(ctx, req.(*GetUsedServRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_AddService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).AddService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/AddService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).AddService(ctx, req.(*AddServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_DeleteService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).DeleteService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/DeleteService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).DeleteService(ctx, req.(*DeleteServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_GetAllServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).GetAllServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/GetAllServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).GetAllServices(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_UseSimForService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(USFSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).UseSimForService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/UseSimForService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).UseSimForService(ctx, req.(*USFSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimActive_GetProviderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimActiveServer).GetProviderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/SimActive/GetProviderList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimActiveServer).GetProviderList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// SimActive_ServiceDesc is the grpc.ServiceDesc for SimActive service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimActive_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SimActive",
	HandlerType: (*SimActiveServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSim",
			Handler:    _SimActive_AddSim_Handler,
		},
		{
			MethodName: "DeleteSim",
			Handler:    _SimActive_DeleteSim_Handler,
		},
		{
			MethodName: "ActivateSim",
			Handler:    _SimActive_ActivateSim_Handler,
		},
		{
			MethodName: "SetSimBlocked",
			Handler:    _SimActive_SetSimBlocked_Handler,
		},
		{
			MethodName: "GetSimList",
			Handler:    _SimActive_GetSimList_Handler,
		},
		{
			MethodName: "GetFreeServices",
			Handler:    _SimActive_GetFreeServices_Handler,
		},
		{
			MethodName: "GetUsedServices",
			Handler:    _SimActive_GetUsedServices_Handler,
		},
		{
			MethodName: "AddService",
			Handler:    _SimActive_AddService_Handler,
		},
		{
			MethodName: "DeleteService",
			Handler:    _SimActive_DeleteService_Handler,
		},
		{
			MethodName: "GetAllServices",
			Handler:    _SimActive_GetAllServices_Handler,
		},
		{
			MethodName: "UseSimForService",
			Handler:    _SimActive_UseSimForService_Handler,
		},
		{
			MethodName: "GetProviderList",
			Handler:    _SimActive_GetProviderList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sim.proto",
}
