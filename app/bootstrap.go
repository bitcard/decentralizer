package app

import (
	"context"
	"github.com/iain17/discovery"
	"github.com/iain17/logger"
	"gx/ipfs/QmNUKMfTHQQpEwE8bUdv5qmKC3ymdW7zw82LFS8D6MQXmu/go-ipfs/core"
	pstore "gx/ipfs/QmPgDWmTmuzvP7QE5zwo1TmjbJme9pmZHNujB2453jkCTr/go-libp2p-peerstore"
	ma "gx/ipfs/QmXY77cVe7rVRQXZZQRioukUM7aRW3BTcAgJe12MCtb3Ji/go-multiaddr"
	"gx/ipfs/QmXYjuNuxVzXKJCfWasQk1RqkhVLDM9jtUKhqc2WPQmFSB/go-libp2p-peer"
	"strings"
	"time"
)

func init() {
	core.DefaultBootstrapConfig = core.BootstrapConfig{
		MinPeerThreshold:  4,
		Period:            30 * time.Second,
		ConnectionTimeout: (30 * time.Second) / 3, // Period / 3
		BootstrapPeers: func() []pstore.PeerInfo {
			return nil
		},
	}
}

func (d *Decentralizer) bootstrap() []pstore.PeerInfo {
	logger.Info("Bootstrapping")
	d.setInfo()
	var peers []pstore.PeerInfo
	for _, peer := range d.d.WaitForPeers(MIN_DISCOVERED_PEERS, 5*time.Minute) {
		peerInfo, err := getInfo(peer)
		if err != nil {
			logger.Warning(err)
			continue
		}
		err = d.i.PeerHost.Connect(context.Background(), *peerInfo)
		if err != nil {
			logger.Warning(err)
			continue
		}
		peers = append(peers, *peerInfo)
	}
	logger.Infof("Bootstrapped %d peers", len(peers))
	return peers
}

func (d *Decentralizer) setInfo() {
	ln := d.d.LocalNode
	addrs := ""
	for _, addr := range d.i.PeerHost.Addrs() {
		addrs += addr.String() + DELIMITER_ADDR
	}

	ln.SetInfo("peerId", d.i.Identity.Pretty())
	ln.SetInfo("addr", addrs)
}

func getInfo(remoteNode *discovery.RemoteNode) (*pstore.PeerInfo, error) {
	sPeerId := remoteNode.GetInfo("peerId")
	peerId, err := peer.IDB58Decode(sPeerId)
	if err != nil {
		return nil, err
	}
	var addrs []ma.Multiaddr
	for _, strAddr := range strings.Split(remoteNode.GetInfo("addr"), DELIMITER_ADDR) {
		addr, err := ma.NewMultiaddr(strAddr)
		if err != nil {
			logger.Warning(err)
			continue
		}
		addrs = append(addrs, addr)
	}
	return &pstore.PeerInfo{
		ID:    peerId,
		Addrs: addrs,
	}, nil
}
