package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/iain17/decentralizer/app/peerstore"
	"github.com/iain17/decentralizer/app/sessionstore"
	"github.com/iain17/decentralizer/pb"
	"github.com/iain17/decentralizer/utils"
	"github.com/iain17/logger"
	"github.com/iain17/timeout"
	"gx/ipfs/QmT6n4mspWYEya864BhCUJEgyxiRfmiSY9ruQwTUNpRKaM/protobuf/proto"
	inet "gx/ipfs/QmU4vCDZTPLDqSDKguWbHCiUe46mZUtmM2g2suBZ9NE8ko/go-libp2p-net"
	peer "gx/ipfs/QmWNY7dV54ZDYmTA1ykVdwNCqC11mpU4zSUp6XDpLTH9eG/go-libp2p-peer"
	"sync"
	"time"
	"encoding/hex"
)

func (d *Decentralizer) getKey(sessionType uint64) string {
	var ih [20]byte
	if d.n != nil {
		ih = d.n.InfoHash()
	}
	return fmt.Sprintf("%s_MATCHMAKING_%d", hex.EncodeToString(ih[:]), sessionType)
}

func (d *Decentralizer) initMatchmaking() {
	d.i.PeerHost.SetStreamHandler(GET_SESSION_REQ, d.getSessionResponse)
}

func (d *Decentralizer) getSessionStorage(sessionType uint64) *sessionstore.Store {
	if d.sessions[sessionType] == nil {
		var err error
		d.sessions[sessionType], err = sessionstore.New(MAX_SESSIONS, time.Duration((EXPIRE_TIME_SESSION*1.5)*time.Second), d.i.Identity)
		if err != nil {
			return nil
		}
	}
	return d.sessions[sessionType]
}

func (d *Decentralizer) UpsertSession(sessionType uint64, name string, port uint32, details map[string]string) (uint64, error) {
	sessions := d.getSessionStorage(sessionType)
	pId, dId := peerstore.PeerToDnId(d.i.Identity)
	info := &pb.Session{
		DnId:    dId,
		PId:     pId,
		Type:    sessionType,
		Name:    name,
		Address: uint32(utils.Inet_aton(d.GetIP())),
		Port:    port,
		Details: details,
	}
	sessionId, err := sessions.Insert(info)
	if err != nil {
		return 0, err
	}
	d.sessionIdToSessionType[sessionId] = sessionType
	err = d.b.Provide(d.getKey(sessionType))
	return sessionId, err
}

func (d *Decentralizer) DeleteSession(sessionId uint64) error {
	if d.sessionIdToSessionType[sessionId] == 0 {
		return errors.New("no such session exists")
	}
	sessionType := d.sessionIdToSessionType[sessionId]
	sessions := d.getSessionStorage(sessionType)
	return sessions.Remove(sessionId)
}

func (d *Decentralizer) GetSession(sessionId uint64) (*pb.Session, error) {
	if d.sessionIdToSessionType[sessionId] == 0 {
		return nil, errors.New("no such session exists")
	}
	sessionType := d.sessionIdToSessionType[sessionId]
	sessions := d.getSessionStorage(sessionType)
	return sessions.FindSessionId(sessionId)
}

func (d *Decentralizer) GetSessions(sessionType uint64) ([]*pb.Session, error) {
	sessions := d.getSessionStorage(sessionType)
	timeout.Do(func(ctx context.Context) {
		d.refreshSessions(sessionType)
	}, 10*time.Second)
	return sessions.FindAll()
}

func (d *Decentralizer) GetSessionsByDetails(sessionType uint64, key, value string) ([]*pb.Session, error) {
	sessions := d.getSessionStorage(sessionType)
	timeout.Do(func(ctx context.Context) {
		d.refreshSessions(sessionType)
	}, 10*time.Second)
	return sessions.FindByDetails(key, value)
}

func (d *Decentralizer) refreshSessions(sessionType uint64) {
	var wg sync.WaitGroup
	sessionsStorage := d.getSessionStorage(sessionType)
	seen := map[string]bool{}
	providers := d.b.Find(d.getKey(sessionType), MAX_SESSIONS)
	logger.Infof("Found %d providers for %d", len(providers), sessionType)
	for provider := range providers {
		//Stop any duplicates
		id := provider.String()
		if seen[id] {
			continue
		}
		seen[id] = true

		wg.Add(1)
		go func(id peer.ID) {
			defer wg.Done()
			sessions, err := d.getSessionsRequest(id, sessionType)
			if err != nil {
				return
			}
			logger.Infof("Received sessions %d from %s", len(sessions), id.Pretty())
			for _, session := range sessions {
				sessionId, err := sessionsStorage.Insert(session)
				if err != nil {
					return
				}
				d.sessionIdToSessionType[sessionId] = sessionType
			}
		}(provider)
	}
	wg.Wait()
}

func (d *Decentralizer) getSessionResponse(stream inet.Stream) {
	reqData, err := Read(stream)
	if err != nil {
		logger.Error(err)
		return
	}
	var request pb.DNSessionRequest
	err = proto.Unmarshal(reqData, &request)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("Someone requested our sessions of type %d...", request.Type)
	sessionsStorage := d.getSessionStorage(request.Type)
	sessions, err := sessionsStorage.FindByPeerId(d.i.Identity.Pretty())
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("Sending %d sessions back of type %d", len(sessions), request.Type)
	//Response
	response, err := proto.Marshal(&pb.DNSessionResponse{
		Results: sessions,
	})
	if err != nil {
		logger.Error(err)
		return
	}
	err = Write(stream, response)
	if err != nil {
		logger.Error(err)
		return
	}
}

func (d *Decentralizer) getSessionsRequest(peer peer.ID, sessionType uint64) ([]*pb.Session, error) {
	stream, err := d.i.PeerHost.NewStream(d.i.Context(), peer, GET_SESSION_REQ)
	if err != nil {
		return nil, err
	}
	defer stream.Close()
	//Request
	reqData, err := proto.Marshal(&pb.DNSessionRequest{
		Type: sessionType,
	})
	if err != nil {
		return nil, err
	}
	err = Write(stream, reqData)
	if err != nil {
		return nil, err
	}

	//Response
	resData, err := Read(stream)
	if err != nil {
		return nil, err
	}
	var response pb.DNSessionResponse
	err = proto.Unmarshal(resData, &response)
	if err != nil {
		return nil, err
	}
	return response.Results, nil
}
