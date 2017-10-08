package discovery

import (
	"github.com/iain17/dht"
	"sync"
	"github.com/op/go-logging"
	"context"
	"time"
	"net"
	"encoding/hex"
)

type DiscoveryDHT struct {
	node      *dht.Server
	announce *dht.Announce
	localNode *LocalNode
	context context.Context
	ih [20]byte
	lastPeers map[string]bool
	mutex     sync.Mutex

	logger *logging.Logger
}

func (d *DiscoveryDHT) Init(ctx context.Context, ln *LocalNode) (err error) {
	d.logger = logging.MustGetLogger("DiscoveryDHT")
	d.localNode = ln
	d.context = ctx
	d.lastPeers = map[string]bool{}

	conn, err := net.ListenPacket("udp", ":0")
	if err != nil {
		return err
	}
	d.node, err = dht.NewServer(&dht.ServerConfig{
		Conn: conn,
		StartingNodes: dht.GlobalBootstrapAddrs,
	})
	d.ih = d.localNode.discovery.network.InfoHash()
	if err != nil {
		return
	}
	go d.Run()
	return
}

func (d *DiscoveryDHT) Stop() {
	if d.announce != nil {
		d.announce.Close()
	}
	d.node.Close()
}

func (d *DiscoveryDHT) Run() {
	defer d.Stop()
	d.request()
	if d.announce == nil {
		d.logger.Error("Can't initiate DHT.")
		return
	}

	for {
		select {
		case <-d.context.Done():
			return
		case peers, ok := <-d.announce.Peers:
			if !ok {
				d.announce.Close()
				time.Sleep(30 * time.Second)
				d.request()
				continue
			}
			if !d.localNode.netTableService.isEnoughPeers() {
				for _, peer := range peers.Peers {
					go d.addPeer(&peer)
				}
			}
		}
	}
}

func (d *DiscoveryDHT) request() {
	d.logger.Debugf("sending request '%s'", hex.EncodeToString(d.ih[:]))
	var err error
	d.announce, err = d.node.Announce(d.ih, d.localNode.port, false)
	if err != nil {
		d.logger.Warning(err)
	}
}

func (d *DiscoveryDHT) addPeer(peer *dht.Peer) {
	d.mutex.Lock()
	key := peer.String()
	if d.lastPeers[key] {
		d.mutex.Unlock()
		return
	}
	d.lastPeers[key] = true
	d.mutex.Unlock()

	d.logger.Debugf("new potential DHT peer %q discovered", peer)
	d.localNode.netTableService.GetNewConnChan() <- &net.UDPAddr{
		IP:   peer.IP[:],
		Port: int(peer.Port),
	}
}