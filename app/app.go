package app

import (
	"context"
	"errors"
	"github.com/hashicorp/golang-lru"
	"github.com/iain17/decentralizer/app/ipfs"
	"github.com/iain17/decentralizer/app/peerstore"
	"github.com/iain17/decentralizer/app/sessionstore"
	"github.com/iain17/decentralizer/pb"
	"github.com/iain17/discovery"
	"github.com/iain17/discovery/network"
	"github.com/iain17/kvcache/lttlru"
	"github.com/iain17/logger"
	"github.com/iain17/stime"
	"github.com/jasonlvhit/gocron"
	"github.com/shibukawa/configdir"
	"github.com/spf13/afero"
	//"gx/ipfs/QmemVjhp1UuWPQqrWSvPcaqH3QJRMjMqNm4T2RULMkDDQe/go-libp2p-swarm"
	libp2pPeer "gx/ipfs/QmdVrMn1LhB4ybb8hMVaMLXnA8XRSewMnK6YqXKXoTcRvN/go-libp2p-peer"
	"gx/ipfs/QmebqVUQQqQFhg74FtQFszUJo22Vpr3e8qBAkvvV4ho9HH/go-ipfs/core"
	"gx/ipfs/QmebqVUQQqQFhg74FtQFszUJo22Vpr3e8qBAkvvV4ho9HH/go-ipfs/core/coreapi"
	coreiface "gx/ipfs/QmebqVUQQqQFhg74FtQFszUJo22Vpr3e8qBAkvvV4ho9HH/go-ipfs/core/coreapi/interface"
	"hash"
	"hash/crc32"
	"net"
	"os"
	"sync"
	"time"
	"github.com/iain17/decentralizer/vars"
)

type Decentralizer struct {
	mutex             sync.Mutex
	ctx               context.Context
	n                 *network.Network
	cron              *gocron.Scheduler
	cronChan          chan bool
	d                 *discovery.Discovery
	i                 *core.IpfsNode
	b                 *ipfs.BitswapService
	ip                *net.IP
	api               coreiface.CoreAPI
	limitedConnection bool
	fs                afero.Fs //general file system

	//Peer ids that did not respond to our queries.
	ignore         *lttlru.LruWithTTL
	crcTable       hash.Hash32
	unmarshalCache *lru.Cache //We unmarshal the same data over and over. Let us cache this.
	connected	   bool//We have connected

	//Storage
	filesApi       *ipfs.FilesAPI
	anythingToPublish bool
	peerFileSystem afero.Fs

	//Matchmaking
	matchmakingMutex            sync.Mutex
	searchMutex                 sync.Mutex
	sessionQueries              chan sessionRequest
	sessions                    map[uint64]*sessionstore.Store
	sessionIdToSessionType      map[uint64]uint64
	sessionIdToSessionTypeMutex sync.RWMutex
	searches                    *lttlru.LruWithTTL

	//addressbook
	peers *peerstore.Store

	//messaging
	directMessageChannels map[uint32]chan *pb.RPCDirectMessage

	//Publisher files
	publisherRecord     *pb.DNPublisherRecord
	publisherDefinition *pb.PublisherDefinition
}

var configPath configdir.ConfigDir
var Base = getBasePath()

func getBasePath() *configdir.Config {
	paths := configPath.QueryFolders(configdir.Global)
	if len(paths) == 0 {
		panic(errors.New("queryFolder request failed"))
	}
	return paths[0]
}

func Reset() {
	os.RemoveAll(configPath.QueryCacheFolder().Path)
	os.RemoveAll(getBasePath().Path + "/ipfs")
}

func New(ctx context.Context, networkStr string, privateKey bool, limitedConnection bool, profile string) (*Decentralizer, error) {
	configPath = configdir.New("ECorp", profile)
	var n *network.Network
	var err error
	if privateKey {
		n, err = network.UnmarshalFromPrivateKey(networkStr)
	} else {
		n, err = network.Unmarshal(networkStr)
	}
	if err != nil {
		return nil, err
	}
	swarmKey, err := Asset("static/swarm.key")
	if err != nil {
		return nil, err
	}

	if stime.IsBadNetwork() {
		logger.Info("Detected a bad network. Enabling limited connection mode")
		limitedConnection = true
	}

	ipfsPath := configPath.QueryCacheFolder().Path + "/ipfs"
	logger.Infof("IPFS path: %s", ipfsPath)
	logger.Infof("Cache path: %s", configPath.QueryCacheFolder().Path)
	i, err := ipfs.OpenIPFSRepo(ctx, ipfsPath, limitedConnection, swarmKey)
	if err != nil {
		return nil, err
	}
	b, err := ipfs.NewBitSwap(i)
	if err != nil {
		return nil, err
	}
	ignore, err := lttlru.NewTTL(vars.MAX_IGNORE)
	if err != nil {
		return nil, err
	}
	unmarshalCache, err := lru.New(vars.MAX_UNMARSHAL_CACHE)
	if err != nil {
		return nil, err
	}
	paths := configPath.QueryFolders(configdir.Global)
	if len(paths) == 0 {
		return nil, errors.New("could not resolve config path")
	}
	base := afero.NewBasePathFs(afero.NewOsFs(), paths[0].Path)
	layer := afero.NewMemMapFs()
	instance := &Decentralizer{
		limitedConnection: limitedConnection,
		ctx:               ctx,
		cron:              gocron.NewScheduler(),
		n:                 n,
		i:                 i,
		b:                 b,
		api:               coreapi.NewCoreAPI(i),
		directMessageChannels: make(map[uint32]chan *pb.RPCDirectMessage),
		ignore:                ignore,
		unmarshalCache:        unmarshalCache,
		crcTable:              crc32.NewIEEE(),
		fs:                    afero.NewCacheOnReadFs(base, layer, 30*time.Minute),
	}
	instance.initializeComponents(false)
	return instance, err
}

func (s *Decentralizer) initializeComponents(testing bool) {
	if !testing {
		s.initDiscovery()
		s.initBootstrap()
	}
	s.initStorage()
	s.initMatchmaking()
	s.initMessaging()
	s.initAddressbook()
	s.initPublisherFiles()
	s.cronChan = s.cron.Start()
}

func (s *Decentralizer) decodePeerId(id string) (libp2pPeer.ID, error) {
	if id == "self" {
		return s.i.Identity, nil
	}
	return libp2pPeer.IDB58Decode(id)
}

func (d *Decentralizer) GetIP() net.IP {
	if d.ip != nil {
		return *d.ip
	}
	self, err := d.peers.FindByPeerId("self")
	if err != nil || self.Details["ip"] == "" {
		return net.ParseIP("127.0.0.1")
	}
	ip := net.ParseIP(self.Details["ip"])
	d.ip = &ip
	return ip
}

func (d *Decentralizer) IsEnoughPeers() bool {
	if d.i == nil {
		return false
	}
	lenPeers := len(d.i.PeerHost.Network().Peers())
	return lenPeers >= vars.MIN_CONNECTED_PEERS
}

func (d *Decentralizer) WaitTilEnoughPeers() {
	for {
		if d.IsEnoughPeers() {
			break
		}
		time.Sleep(300 * time.Millisecond)
	}
}

func (s *Decentralizer) Stop() {
	s.cronChan <- false
	if s.i != nil {
		s.i.Close()
	}
}

func (d *Decentralizer) clearBackOff(id libp2pPeer.ID) {
	//TODO: Fix this again
	//snet, ok := d.i.PeerHost.Network().(*swarm.Network)
	//if ok {
	//	snet.Swarm().Backoff().Clear(id)
	//}
}
