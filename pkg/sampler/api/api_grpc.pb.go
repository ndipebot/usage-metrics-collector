// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pkg/sampler/api/api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	grpc_health_v1 "google.golang.org/grpc/health/grpc_health_v1"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsClient interface {
	ListMetrics(ctx context.Context, in *ListMetricsRequest, opts ...grpc.CallOption) (*ListMetricsResponse, error)
}

type metricsClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsClient(cc grpc.ClientConnInterface) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) ListMetrics(ctx context.Context, in *ListMetricsRequest, opts ...grpc.CallOption) (*ListMetricsResponse, error) {
	out := new(ListMetricsResponse)
	err := c.cc.Invoke(ctx, "/containerd.api.Metrics/ListMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServer is the server API for Metrics service.
// All implementations must embed UnimplementedMetricsServer
// for forward compatibility
type MetricsServer interface {
	ListMetrics(context.Context, *ListMetricsRequest) (*ListMetricsResponse, error)
	mustEmbedUnimplementedMetricsServer()
}

// UnimplementedMetricsServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (UnimplementedMetricsServer) ListMetrics(context.Context, *ListMetricsRequest) (*ListMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetrics not implemented")
}
func (UnimplementedMetricsServer) mustEmbedUnimplementedMetricsServer() {}

// UnsafeMetricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServer will
// result in compilation errors.
type UnsafeMetricsServer interface {
	mustEmbedUnimplementedMetricsServer()
}

func RegisterMetricsServer(s grpc.ServiceRegistrar, srv MetricsServer) {
	s.RegisterService(&Metrics_ServiceDesc, srv)
}

func _Metrics_ListMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).ListMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.api.Metrics/ListMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).ListMetrics(ctx, req.(*ListMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Metrics_ServiceDesc is the grpc.ServiceDesc for Metrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.api.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMetrics",
			Handler:    _Metrics_ListMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/sampler/api/api.proto",
}

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest, opts ...grpc.CallOption) (*grpc_health_v1.HealthCheckResponse, error)
	IsLeader(ctx context.Context, in *grpc_health_v1.HealthCheckRequest, opts ...grpc.CallOption) (*grpc_health_v1.HealthCheckResponse, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *grpc_health_v1.HealthCheckRequest, opts ...grpc.CallOption) (*grpc_health_v1.HealthCheckResponse, error) {
	out := new(grpc_health_v1.HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/containerd.api.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) IsLeader(ctx context.Context, in *grpc_health_v1.HealthCheckRequest, opts ...grpc.CallOption) (*grpc_health_v1.HealthCheckResponse, error) {
	out := new(grpc_health_v1.HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/containerd.api.Health/IsLeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
// All implementations must embed UnimplementedHealthServer
// for forward compatibility
type HealthServer interface {
	Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error)
	IsLeader(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error)
	mustEmbedUnimplementedHealthServer()
}

// UnimplementedHealthServer must be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (UnimplementedHealthServer) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedHealthServer) IsLeader(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsLeader not implemented")
}
func (UnimplementedHealthServer) mustEmbedUnimplementedHealthServer() {}

// UnsafeHealthServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthServer will
// result in compilation errors.
type UnsafeHealthServer interface {
	mustEmbedUnimplementedHealthServer()
}

func RegisterHealthServer(s grpc.ServiceRegistrar, srv HealthServer) {
	s.RegisterService(&Health_ServiceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc_health_v1.HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.api.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*grpc_health_v1.HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_IsLeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(grpc_health_v1.HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).IsLeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containerd.api.Health/IsLeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).IsLeader(ctx, req.(*grpc_health_v1.HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Health_ServiceDesc is the grpc.ServiceDesc for Health service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Health_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.api.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
		{
			MethodName: "IsLeader",
			Handler:    _Health_IsLeader_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/sampler/api/api.proto",
}

// MetricsCollectorClient is the client API for MetricsCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsCollectorClient interface {
	PushMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsCollector_PushMetricsClient, error)
}

type metricsCollectorClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsCollectorClient(cc grpc.ClientConnInterface) MetricsCollectorClient {
	return &metricsCollectorClient{cc}
}

func (c *metricsCollectorClient) PushMetrics(ctx context.Context, opts ...grpc.CallOption) (MetricsCollector_PushMetricsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetricsCollector_ServiceDesc.Streams[0], "/containerd.api.MetricsCollector/PushMetrics", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsCollectorPushMetricsClient{stream}
	return x, nil
}

type MetricsCollector_PushMetricsClient interface {
	Send(*ListMetricsResponse) error
	Recv() (*ConfigurePush, error)
	grpc.ClientStream
}

type metricsCollectorPushMetricsClient struct {
	grpc.ClientStream
}

func (x *metricsCollectorPushMetricsClient) Send(m *ListMetricsResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsCollectorPushMetricsClient) Recv() (*ConfigurePush, error) {
	m := new(ConfigurePush)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsCollectorServer is the server API for MetricsCollector service.
// All implementations must embed UnimplementedMetricsCollectorServer
// for forward compatibility
type MetricsCollectorServer interface {
	PushMetrics(MetricsCollector_PushMetricsServer) error
	mustEmbedUnimplementedMetricsCollectorServer()
}

// UnimplementedMetricsCollectorServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsCollectorServer struct {
}

func (UnimplementedMetricsCollectorServer) PushMetrics(MetricsCollector_PushMetricsServer) error {
	return status.Errorf(codes.Unimplemented, "method PushMetrics not implemented")
}
func (UnimplementedMetricsCollectorServer) mustEmbedUnimplementedMetricsCollectorServer() {}

// UnsafeMetricsCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsCollectorServer will
// result in compilation errors.
type UnsafeMetricsCollectorServer interface {
	mustEmbedUnimplementedMetricsCollectorServer()
}

func RegisterMetricsCollectorServer(s grpc.ServiceRegistrar, srv MetricsCollectorServer) {
	s.RegisterService(&MetricsCollector_ServiceDesc, srv)
}

func _MetricsCollector_PushMetrics_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsCollectorServer).PushMetrics(&metricsCollectorPushMetricsServer{stream})
}

type MetricsCollector_PushMetricsServer interface {
	Send(*ConfigurePush) error
	Recv() (*ListMetricsResponse, error)
	grpc.ServerStream
}

type metricsCollectorPushMetricsServer struct {
	grpc.ServerStream
}

func (x *metricsCollectorPushMetricsServer) Send(m *ConfigurePush) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsCollectorPushMetricsServer) Recv() (*ListMetricsResponse, error) {
	m := new(ListMetricsResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsCollector_ServiceDesc is the grpc.ServiceDesc for MetricsCollector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsCollector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "containerd.api.MetricsCollector",
	HandlerType: (*MetricsCollectorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PushMetrics",
			Handler:       _MetricsCollector_PushMetrics_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pkg/sampler/api/api.proto",
}
