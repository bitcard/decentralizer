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
	Hash string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
}

func (m *GetServiceRequest) Reset()                    { *m = GetServiceRequest{} }
func (m *GetServiceRequest) String() string            { return proto.CompactTextString(m) }
func (*GetServiceRequest) ProtoMessage()               {}
func (*GetServiceRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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
func (*GetServiceResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

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

type Peer struct {
	Ip      string            `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Port    uint32            `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	RpcPort uint32            `protobuf:"varint,3,opt,name=rpcPort" json:"rpcPort,omitempty"`
	Details map[string]string `protobuf:"bytes,4,rep,name=details" json:"details,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Peer) Reset()                    { *m = Peer{} }
func (m *Peer) String() string            { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()               {}
func (*Peer) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *Peer) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *Peer) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Peer) GetRpcPort() uint32 {
	if m != nil {
		return m.RpcPort
	}
	return 0
}

func (m *Peer) GetDetails() map[string]string {
	if m != nil {
		return m.Details
	}
	return nil
}

func init() {
	proto.RegisterType((*GetServiceRequest)(nil), "pb.getServiceRequest")
	proto.RegisterType((*GetServiceResponse)(nil), "pb.getServiceResponse")
	proto.RegisterType((*Peer)(nil), "pb.peer")
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

func init() { proto.RegisterFile("rpc.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0xdd, 0x4d, 0xda, 0x9a, 0xd1, 0x8a, 0x0e, 0x2a, 0xa1, 0x07, 0x09, 0xb9, 0xd8, 0x53,
	0x84, 0x7a, 0x91, 0x9e, 0xc4, 0x3f, 0xd4, 0x63, 0x59, 0xc1, 0x7b, 0x12, 0x07, 0x1b, 0x0c, 0xcd,
	0xb8, 0xbb, 0x29, 0xf4, 0x63, 0xf9, 0x0d, 0x25, 0x9b, 0x44, 0xaa, 0xbd, 0xbd, 0xf7, 0xdb, 0x07,
	0x6f, 0x78, 0x0b, 0x81, 0xe6, 0x3c, 0x61, 0x5d, 0xd9, 0x0a, 0x25, 0x67, 0xf1, 0x35, 0x9c, 0x7d,
	0x90, 0x7d, 0x25, 0xbd, 0x29, 0x72, 0x52, 0xf4, 0x55, 0x93, 0xb1, 0x88, 0xe0, 0xaf, 0x52, 0xb3,
	0x0a, 0x45, 0x24, 0xa6, 0x81, 0x72, 0x3a, 0x7e, 0x03, 0xdc, 0x0d, 0x1a, 0xae, 0xd6, 0x86, 0x30,
	0x82, 0xa1, 0x26, 0x53, 0x97, 0xd6, 0x65, 0x8f, 0x66, 0x87, 0x09, 0x67, 0x09, 0x13, 0x69, 0xd5,
	0x71, 0xbc, 0x82, 0x41, 0xe3, 0x4d, 0x28, 0x23, 0xef, 0x4f, 0xa0, 0xc5, 0xf1, 0xb7, 0x00, 0xbf,
	0x51, 0x78, 0x02, 0xb2, 0xe0, 0xae, 0x52, 0x16, 0xdc, 0x1c, 0xc1, 0x95, 0xb6, 0xa1, 0x8c, 0xc4,
	0x74, 0xac, 0x9c, 0xc6, 0x10, 0x46, 0x9a, 0xf3, 0x65, 0x83, 0x3d, 0x87, 0x7b, 0x8b, 0x37, 0x30,
	0x7a, 0x27, 0x9b, 0x16, 0xa5, 0x09, 0x7d, 0x57, 0x74, 0xd1, 0x17, 0x25, 0x4f, 0x2d, 0x7f, 0x5e,
	0x5b, 0xbd, 0x55, 0x7d, 0x6a, 0x32, 0x87, 0xe3, 0xdd, 0x07, 0x3c, 0x05, 0xef, 0x93, 0xb6, 0x5d,
	0x7f, 0x23, 0xf1, 0x1c, 0x06, 0x9b, 0xb4, 0xac, 0xc9, 0x5d, 0x10, 0xa8, 0xd6, 0xcc, 0xe5, 0x9d,
	0x98, 0x2d, 0xc0, 0xd3, 0x9c, 0xe3, 0x3d, 0x8c, 0xd5, 0xf2, 0x71, 0xf1, 0xbb, 0x0a, 0xba, 0xce,
	0xbd, 0x39, 0x27, 0x97, 0xff, 0x71, 0x3b, 0x5e, 0x7c, 0xf0, 0x20, 0x5f, 0x44, 0x36, 0x74, 0x9f,
	0x71, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x0d, 0x72, 0x9b, 0x99, 0x01, 0x00, 0x00,
}
