// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

package pb

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

type GetServiceRequest struct {
	Me   *Peer  `protobuf:"bytes,1,opt,name=me" json:"me,omitempty"`
	Hash string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
}

func (m *GetServiceRequest) Reset()                    { *m = GetServiceRequest{} }
func (m *GetServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetServiceRequest) ProtoMessage()               {}
func (*GetServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *GetServiceRequest) GetMe() *Peer {
	if m != nil {
		return m.Me
	}
	return nil
}

func (m *GetServiceRequest) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

type GetServiceResponse struct {
	Result *Peer   `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
	Peers  []*Peer `protobuf:"bytes,2,rep,name=peers" json:"peers,omitempty"`
}

func (m *GetServiceResponse) Reset()                    { *m = GetServiceResponse{} }
func (m *GetServiceResponse) String() string            { return proto.CompactTextString(m) }
func (*GetServiceResponse) ProtoMessage()               {}
func (*GetServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *GetServiceResponse) GetResult() *Peer {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *GetServiceResponse) GetPeers() []*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*GetServiceRequest)(nil), "pb.getServiceRequest")
	proto.RegisterType((*GetServiceResponse)(nil), "pb.getServiceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Rpc service

type RpcClient interface {
	RPCGetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
}

type rpcClient struct {
	cc *grpc.ClientConn
}

func NewRpcClient(cc *grpc.ClientConn) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) RPCGetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	out := new(GetServiceResponse)
	err := grpc.Invoke(ctx, "/pb.rpc/RPCGetService", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Rpc service

type RpcServer interface {
	RPCGetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
}

func RegisterRpcServer(s *grpc.Server, srv RpcServer) {
	s.RegisterService(&_Rpc_serviceDesc, srv)
}

func _Rpc_RPCGetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).RPCGetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.rpc/RPCGetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).RPCGetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RPCGetService",
			Handler:    _Rpc_RPCGetService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x92, 0xe2, 0xc9, 0xcd, 0x4f, 0x49, 0xcd,
	0x29, 0x86, 0x88, 0x28, 0x39, 0x72, 0x09, 0xa6, 0xa7, 0x96, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26,
	0xa7, 0x06, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x49, 0x70, 0x31, 0xe5, 0xa6, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x71, 0xe8, 0x15, 0x24, 0xe9, 0x15, 0xa4, 0xa6, 0x16, 0x05, 0x31,
	0xe5, 0xa6, 0x0a, 0x09, 0x71, 0xb1, 0x64, 0x24, 0x16, 0x67, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70,
	0x06, 0x81, 0xd9, 0x4a, 0x61, 0x5c, 0x42, 0xc8, 0x46, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x0a,
	0x29, 0x70, 0xb1, 0x15, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x60, 0x98, 0x03, 0x15, 0x17, 0x92, 0xe3,
	0x62, 0x05, 0xf1, 0x8b, 0x25, 0x98, 0x14, 0x98, 0x51, 0x14, 0x40, 0x84, 0x8d, 0xdc, 0xb9, 0x98,
	0x8b, 0x0a, 0x92, 0x85, 0x1c, 0xb8, 0x78, 0x83, 0x02, 0x9c, 0xdd, 0xe1, 0x36, 0x08, 0x89, 0x82,
	0x14, 0x62, 0x38, 0x5a, 0x4a, 0x0c, 0x5d, 0x18, 0xe2, 0x10, 0x25, 0x06, 0x27, 0x26, 0x0f, 0xc6,
	0x24, 0x36, 0xb0, 0x77, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72, 0xec, 0xdd, 0x82, 0x0d,
	0x01, 0x00, 0x00,
}