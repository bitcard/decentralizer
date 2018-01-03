//Takes care of asserting if services are still using the this background service. If its not being used, close.
package api

import (
	"google.golang.org/grpc"
	"context"
)

func (s *Server) AliveStreamInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		s.Wg.Add(1)
		go func() {
			<- stream.Context().Done()
			s.Wg.Done()
		}()
		return handler(srv, stream)
	}
}

func (s *Server) AliveUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		s.Wg.Add(1)
		defer s.Wg.Done()
		return handler(ctx, req)
	}
}