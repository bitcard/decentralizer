package app

import (
	"github.com/iain17/discovery"
	"github.com/iain17/logger"
	pstore "gx/ipfs/QmXauCuJzmzapetmC6W4TuDJLL1yFFrVzSHoWv8YdbmnxH/go-libp2p-peerstore"
	"net"
	"errors"
	"gx/ipfs/QmNmJZL7FQySMtE2BQuLMuZg2EB2CLEunJJUSVSc9YnnbV/go-libp2p-host"
	manet "gx/ipfs/QmRK2LxanhK2gZq6k6R7vk5ZoYZk8ULSSTB7FzDsMUX6CB/go-multiaddr-net"
	ma "gx/ipfs/QmWWQ2Txc2c6tqjsBpzg5Ar652cHPGNsQQp2SejkNmkUMb/go-multiaddr"
	"gx/ipfs/QmZoWKhxUmZ2seW4BzX6fJkNR8hh9PsGModr7q171yq2SS/go-libp2p-peer"
	libp2pnet "gx/ipfs/QmXfkENeeBvh3zYA51MaSdGUdBjhQ99cP5WQe8zgr6wchG/go-libp2p-net"
	"strings"
	"time"
)

func (d *Decentralizer) initDiscovery() error {
	addrs, err := getAddrs(d.i.PeerHost)
	if err != nil {
		return err
	}
	d.d, err = discovery.New(d.ctx, d.n, MAX_DISCOVERED_PEERS, d.peerDiscovered, d.limitedConnection, map[string]string{
		"peerId": d.i.Identity.Pretty(),
		"addr": addrs,
	})
	if err != nil {
		logger.Fatal(err)
	}
	d.cron.Every(10).Seconds().Do(func() {
		d.setSelfAddrs()
		d.setReachableAddrs()
	})
	return nil
}

func (d *Decentralizer) setSelfAddrs() {
	addrs, err := getAddrs(d.i.PeerHost)
	if err != nil {
		logger.Warning(err)
		return
	}
	d.d.LocalNode.SetInfo("addr", addrs)
}

func (d *Decentralizer) setReachableAddrs() {
	for _, peer := range d.d.WaitForPeers(MIN_CONNECTED_PEERS, 10*time.Second) {
		peerInfo, err := remoteNodeToPeerInfo(peer)
		if err != nil {
			//logger.Warning(err)
			peer.Close()
			continue
		}

		if d.i.PeerHost.Network().Connectedness(peerInfo.ID) == libp2pnet.Connected {
			peer.SetInfo("addr", serializeAddrs(d.i.PeerHost.Peerstore().Addrs(peerInfo.ID)))
		}
	}
}

func (d *Decentralizer) peerDiscovered(peer *discovery.RemoteNode) {
	info, err := remoteNodeToPeerInfo(peer)
	if err != nil {
		logger.Warning(err)
		return
	}
	d.i.HandlePeerFound(*info)
}

func getDialableListenAddrs(ph host.Host) ([]ma.Multiaddr, error) {
	var out []ma.Multiaddr
	for _, addr := range ph.Addrs() {
		na, err := manet.ToNetAddr(addr)
		if err != nil {
			continue
		}
		if _, ok := na.(*net.TCPAddr); ok {
			out = append(out, addr)
		}
	}
	if len(out) == 0 {
		return nil, errors.New("failed to find good external addr from peerhost")
	}
	return out, nil
}

func getAddrs(ph host.Host) (string, error) {
	maAddrs, err := getDialableListenAddrs(ph)
	if err != nil {
		return "", err
	}
	return serializeAddrs(maAddrs), nil
}

func serializeAddrs(multiAddrs []ma.Multiaddr) string {
	if multiAddrs == nil {
		return ""
	}
	addrs := ""
	for _, addr := range multiAddrs {
		addrs += addr.String() + DELIMITER_ADDR
	}
	return addrs
}

func unSerializeAddrs(addrText string) []ma.Multiaddr {
	var addrs []ma.Multiaddr
	rawAddr := strings.Split(addrText, DELIMITER_ADDR)
	for _, strAddr := range rawAddr {
		addr, err := ma.NewMultiaddr(strAddr)
		if err != nil && addr != nil {
			logger.Warning(err)
			continue
		}
		addrs = append(addrs, addr)
	}
	return addrs
}

func remoteNodeToPeerInfo(remoteNode *discovery.RemoteNode) (*pstore.PeerInfo, error) {
	sPeerId := remoteNode.GetInfo("peerId")
	peerId, err := peer.IDB58Decode(sPeerId)
	if err != nil {
		return nil, err
	}
	addrText := remoteNode.GetInfo("addr")
	addrs := unSerializeAddrs(addrText)
	if len(addrs) == 0 {
		return nil, errors.New("no addr set")
	}
	return &pstore.PeerInfo{
		ID:    peerId,
		Addrs: addrs,
	}, nil
}