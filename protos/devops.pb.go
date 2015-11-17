// Code generated by protoc-gen-go.
// source: devops.proto
// DO NOT EDIT!

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DevopsResponse_StatusCode int32

const (
	DevopsResponse_UNDEFINED DevopsResponse_StatusCode = 0
	DevopsResponse_SUCCESS   DevopsResponse_StatusCode = 1
	DevopsResponse_FAILURE   DevopsResponse_StatusCode = 2
)

var DevopsResponse_StatusCode_name = map[int32]string{
	0: "UNDEFINED",
	1: "SUCCESS",
	2: "FAILURE",
}
var DevopsResponse_StatusCode_value = map[string]int32{
	"UNDEFINED": 0,
	"SUCCESS":   1,
	"FAILURE":   2,
}

func (x DevopsResponse_StatusCode) String() string {
	return proto.EnumName(DevopsResponse_StatusCode_name, int32(x))
}

type BuildResult_StatusCode int32

const (
	BuildResult_UNDEFINED BuildResult_StatusCode = 0
	BuildResult_SUCCESS   BuildResult_StatusCode = 1
	BuildResult_FAILURE   BuildResult_StatusCode = 2
)

var BuildResult_StatusCode_name = map[int32]string{
	0: "UNDEFINED",
	1: "SUCCESS",
	2: "FAILURE",
}
var BuildResult_StatusCode_value = map[string]int32{
	"UNDEFINED": 0,
	"SUCCESS":   1,
	"FAILURE":   2,
}

func (x BuildResult_StatusCode) String() string {
	return proto.EnumName(BuildResult_StatusCode_name, int32(x))
}

type DevopsResponse struct {
	Status DevopsResponse_StatusCode `protobuf:"varint,1,opt,name=status,enum=protos.DevopsResponse_StatusCode" json:"status,omitempty"`
	Msg    []byte                    `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *DevopsResponse) Reset()         { *m = DevopsResponse{} }
func (m *DevopsResponse) String() string { return proto.CompactTextString(m) }
func (*DevopsResponse) ProtoMessage()    {}

type BuildResult struct {
	Status         BuildResult_StatusCode  `protobuf:"varint,1,opt,name=status,enum=protos.BuildResult_StatusCode" json:"status,omitempty"`
	Msg            string                  `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
	DeploymentSpec *ChainletDeploymentSpec `protobuf:"bytes,3,opt,name=deploymentSpec" json:"deploymentSpec,omitempty"`
}

func (m *BuildResult) Reset()         { *m = BuildResult{} }
func (m *BuildResult) String() string { return proto.CompactTextString(m) }
func (*BuildResult) ProtoMessage()    {}

func (m *BuildResult) GetDeploymentSpec() *ChainletDeploymentSpec {
	if m != nil {
		return m.DeploymentSpec
	}
	return nil
}

func init() {
	proto.RegisterEnum("protos.DevopsResponse_StatusCode", DevopsResponse_StatusCode_name, DevopsResponse_StatusCode_value)
	proto.RegisterEnum("protos.BuildResult_StatusCode", BuildResult_StatusCode_name, BuildResult_StatusCode_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Devops service

type DevopsClient interface {
	// Build the chaincode package.
	Build(ctx context.Context, in *ChainletSpec, opts ...grpc.CallOption) (*ChainletDeploymentSpec, error)
	// Deploy the chaincode package to the chain.
	Deploy(ctx context.Context, in *ChainletSpec, opts ...grpc.CallOption) (*ChainletDeploymentSpec, error)
	// Invoke chaincode.
	Invoke(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*DevopsResponse, error)
	// Invoke chaincode.
	Query(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*DevopsResponse, error)
}

type devopsClient struct {
	cc *grpc.ClientConn
}

func NewDevopsClient(cc *grpc.ClientConn) DevopsClient {
	return &devopsClient{cc}
}

func (c *devopsClient) Build(ctx context.Context, in *ChainletSpec, opts ...grpc.CallOption) (*ChainletDeploymentSpec, error) {
	out := new(ChainletDeploymentSpec)
	err := grpc.Invoke(ctx, "/protos.Devops/Build", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Deploy(ctx context.Context, in *ChainletSpec, opts ...grpc.CallOption) (*ChainletDeploymentSpec, error) {
	out := new(ChainletDeploymentSpec)
	err := grpc.Invoke(ctx, "/protos.Devops/Deploy", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Invoke(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*DevopsResponse, error) {
	out := new(DevopsResponse)
	err := grpc.Invoke(ctx, "/protos.Devops/Invoke", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devopsClient) Query(ctx context.Context, in *ChaincodeInvocationSpec, opts ...grpc.CallOption) (*DevopsResponse, error) {
	out := new(DevopsResponse)
	err := grpc.Invoke(ctx, "/protos.Devops/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Devops service

type DevopsServer interface {
	// Build the chaincode package.
	Build(context.Context, *ChainletSpec) (*ChainletDeploymentSpec, error)
	// Deploy the chaincode package to the chain.
	Deploy(context.Context, *ChainletSpec) (*ChainletDeploymentSpec, error)
	// Invoke chaincode.
	Invoke(context.Context, *ChaincodeInvocationSpec) (*DevopsResponse, error)
	// Invoke chaincode.
	Query(context.Context, *ChaincodeInvocationSpec) (*DevopsResponse, error)
}

func RegisterDevopsServer(s *grpc.Server, srv DevopsServer) {
	s.RegisterService(&_Devops_serviceDesc, srv)
}

func _Devops_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChainletSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).Build(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChainletSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).Deploy(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChaincodeInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).Invoke(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Devops_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	in := new(ChaincodeInvocationSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	out, err := srv.(DevopsServer).Query(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Devops_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.Devops",
	HandlerType: (*DevopsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Build",
			Handler:    _Devops_Build_Handler,
		},
		{
			MethodName: "Deploy",
			Handler:    _Devops_Deploy_Handler,
		},
		{
			MethodName: "Invoke",
			Handler:    _Devops_Invoke_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _Devops_Query_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
