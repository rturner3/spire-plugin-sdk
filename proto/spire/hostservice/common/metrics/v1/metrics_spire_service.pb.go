// Code generated by protoc-gen-go-spire. DO NOT EDIT.

package metricsv1

import (
	pluginsdk "github.com/spiffe/spire-plugin-sdk/pluginsdk"
	grpc "google.golang.org/grpc"
)

func MetricsServiceServer(server MetricsServer) pluginsdk.ServiceServer {
	return metricsServiceServer{MetricsServer: server}
}

type metricsServiceServer struct {
	MetricsServer
}

func (s metricsServiceServer) GRPCServiceName() string {
	return "spire.hostservice.common.metrics.v1.Metrics"
}

func (s metricsServiceServer) RegisterServer(server *grpc.Server) interface{} {
	RegisterMetricsServer(server, s.MetricsServer)
	return s.MetricsServer
}

type MetricsServiceClient struct {
	MetricsClient
}

func (c *MetricsServiceClient) IsInitialized() bool {
	return c.MetricsClient != nil
}

func (c *MetricsServiceClient) GRPCServiceName() string {
	return "spire.hostservice.common.metrics.v1.Metrics"
}

func (c *MetricsServiceClient) InitClient(conn grpc.ClientConnInterface) interface{} {
	c.MetricsClient = NewMetricsClient(conn)
	return c.MetricsClient
}