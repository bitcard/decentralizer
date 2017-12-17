package api

import (
	"context"
	"github.com/iain17/decentralizer/pb"
)

//
// Address book
//
// Create or update a peer. Takes peer info, returns if it was a success.
func (s *Server) UpsertPeer(ctx context.Context, request *pb.RPCUpsertPeerRequest) (*pb.RPCUpsertPeerResponse, error) {
	err := s.app.UpsertPeer(request.Peer.PId, request.Peer.Details)
	return &pb.RPCUpsertPeerResponse{}, err
}

// Get peer ids. takes a key and value to filter the peers by details. If left empty this filter will not apply and all will be fetched.
func (s *Server) GetPeerIds(ctx context.Context, request *pb.RPCGetPeerIdsRequest) (*pb.RPCGetPeerIdsResponse, error) {
	var peers []*pb.Peer
	var err error
	if request.Key == "" && request.Value == "" {
		peers, err = s.app.GetPeers()
	} else {
		peers, err = s.app.GetPeersByDetails(request.Key, request.Value)
	}
	var peerIds []string
	for _, peer := range peers {
		peerIds = append(peerIds, peer.PId)
	}
	return &pb.RPCGetPeerIdsResponse{
		PeerIds: peerIds,
	}, err
}

// Get an individual peer. Takes either a peer id or decentralizer id and returns the peer info.
func (s *Server) GetPeer(ctx context.Context, request *pb.RPCGetPeerRequest) (*pb.RPCGetPeerResponse, error) {
	var peer *pb.Peer
	var err error
	if request.PId != "" {
		peer, err = s.app.FindByPeerId(request.PId)
	}
	if peer == nil && request.DnId != 0 {
		peer, err = s.app.FindByDecentralizedId(request.DnId)
	}
	return &pb.RPCGetPeerResponse{
		Peer: peer,
	}, err
}