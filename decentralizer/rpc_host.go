package decentralizer

import (
	"golang.org/x/net/context"
	"github.com/iain17/dht-hello/decentralizer/pb"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"fmt"
	logger "github.com/Sirupsen/logrus"
	"net"
	"github.com/pkg/errors"
)

/*
- The rpc server is a TCP GRPC server that is used to exchange messages between nodes.
 */
func (d *decentralizer) listenRpcServer() error {
	lis, err := getTcpConn()
	if err != nil {
		return err
	}
	port := lis.Addr().(*net.TCPAddr).Port
	s := grpc.NewServer()
	pb.RegisterDecentralizerServer(s, d)
	reflection.Register(s)
	d.rpcPort = uint16(port)

	go func() {
		if err := s.Serve(lis); err != nil {
			panic(fmt.Sprintf("failed to serve: %v", err))
		}
	}()
	logger.Infof("RPC server listening at %d", port)
	return nil
}

func (d *decentralizer) RPCGetService(ctx context.Context, req *pb.GetServiceRequest) (*pb.GetServiceResponse, error) {
	service := d.services[req.Hash]
	if service != nil {
		return nil, errors.New("No such service registered under that hash")
	}
	return &pb.GetServiceResponse{
		Result: &pb.Peer{
			Ip: service.self.IP,
			Port: uint32(service.self.Port),
			RpcPort: uint32(service.self.RPCPort),
			Details: service.self.Details,
		},
		//TODO others.
	}, nil
}