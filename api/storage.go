package api

import (
	"context"
	"github.com/iain17/decentralizer/pb"
	"github.com/iain17/logger"
	"time"
	//"github.com/ipfs/go-ipfs/namesys"
	"github.com/giantswarm/retry-go"
)

//
// Storage
//
// Write a user file. Takes a file name and the data it should save.
func (s *Server) WritePeerFile(ctx context.Context, request *pb.RPCWritePeerFileRequest) (*pb.RPCWritePeerFileResponse, error) {
	_, err := s.App.SavePeerFile(request.Name, request.File)
	if err != nil {
		logger.Warning(err)
	}
	return &pb.RPCWritePeerFileResponse{
		Success: err != nil,
	}, err
}

// Get a user file. Takes a file name, returns the file.
func (s *Server) GetPeerFile(ctx context.Context, request *pb.RPCGetPeerFileRequest) (*pb.RPCGetPeerFileResponse, error) {
	time_start := time.Now()
	var file []byte
	err := retry.Do(func() error {
		var err error
		file, err = s.App.GetPeerFile(request.PId, request.Name)
		return err
	},
	retry.RetryChecker(func(err error) bool {
		return true
	}),
	retry.Timeout(10 * time.Second),
	retry.Sleep(3 * time.Second))
	if err != nil {
		logger.Warning(err)
	}
	logger.Infof("Responded get peer file request in: %s", time.Since(time_start).String())
	return &pb.RPCGetPeerFileResponse{
		File: file,
	}, err
}