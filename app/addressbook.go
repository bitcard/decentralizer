package app

import (
	"github.com/iain17/decentralizer/pb"
	"github.com/iain17/logger"
	"gx/ipfs/QmT6n4mspWYEya864BhCUJEgyxiRfmiSY9ruQwTUNpRKaM/protobuf/proto"
	inet "gx/ipfs/QmNa31VPzC561NWwRsJLE7nGYZYuuD2QfpK2b1q9BK54J1/go-libp2p-net"
	Peer "gx/ipfs/QmXYjuNuxVzXKJCfWasQk1RqkhVLDM9jtUKhqc2WPQmFSB/go-libp2p-peer"
	"github.com/Pallinder/go-randomdata"
	pstore "gx/ipfs/QmPgDWmTmuzvP7QE5zwo1TmjbJme9pmZHNujB2453jkCTr/go-libp2p-peerstore"
	ma "gx/ipfs/QmXY77cVe7rVRQXZZQRioukUM7aRW3BTcAgJe12MCtb3Ji/go-multiaddr"
	"time"
	"github.com/iain17/framed"
	"github.com/iain17/decentralizer/app/peerstore"
)

func (d *Decentralizer) initAddressbook() {
	var err error
	d.peers, err = peerstore.New(MAX_CONTACTS, time.Duration((EXPIRE_TIME_CONTACT*1.5)*time.Second), d.i.Identity)
	if err != nil {
		logger.Fatal(err)
	}
	d.i.PeerHost.SetStreamHandler(GET_PEER_REQ, d.getPeerResponse)
	d.downloadPeers()
	d.saveSelf()
	go func() {
		d.WaitTilEnoughPeers()
		d.connectPreviousPeers()
	}()
	go d.provideSelf()
	d.cron.Every(30).Seconds().Do(d.uploadPeers)
	d.cron.Every(5).Minutes().Do(d.provideSelf)
}

func (d *Decentralizer) downloadPeers() {
	data, err := Base.ReadFile(ADDRESS_BOOK_FILE)
	if err != nil {
		//logger.Warningf("Could not restore address book: %v", err)
		return
	}
	var addressbook pb.DNAddressbook
	err = proto.Unmarshal(data, &addressbook)
	if err != nil {
		logger.Warningf("Could not restore address book: %v", err)
		return
	}
	for _, peer := range addressbook.Peers {
		err = d.peers.Upsert(peer)
		if err != nil {
			logger.Warningf("Error saving peer: %s", peer.PId)
			continue
		}
	}
	logger.Info("Restored address book")
}

func (d *Decentralizer) provideSelf() {
	d.WaitTilEnoughPeers()
	peer, err := d.FindByPeerId("self")
	if err != nil {
		logger.Warningf("Could not provide self: %v", err)
		return
	}
	d.b.Provide(getDecentralizedIdKey(peer.DnId))
	logger.Debug("Provided self")
}

func (d *Decentralizer) uploadPeers() {
	if !d.addressBookChanged {
		return
	}
	peers, err := d.peers.FindAll()
	if err != nil {
		logger.Warningf("Could not save address book: %v", err)
		return
	}
	data, err := proto.Marshal(&pb.DNAddressbook{
		Peers: peers,
	})
	if err != nil {
		logger.Warningf("Could not save address book: %v", err)
		return
	}
	err = Base.WriteFile(ADDRESS_BOOK_FILE, data)
	if err != nil {
		logger.Warningf("Could not save address book: %v", err)
		return
	}
	d.addressBookChanged = false
	logger.Info("Saved address book")
}

//Connect to our previous peers
func (d *Decentralizer) connectPreviousPeers() error {
	logger.Info("Connecting to previous peers...")
	peers, err := d.peers.FindAll()
	if err != nil {
		return err
	}
	for _, peer := range peers {
		pId, err := d.decodePeerId(peer.PId)
		if err != nil {
			continue
		}
		var addrs []ma.Multiaddr
		for _, rawAddr := range peer.Addrs {
			addr, err := ma.NewMultiaddr(rawAddr)
			if err != nil {
				continue
			}
			addrs = append(addrs, addr)
		}
		err = d.i.PeerHost.Connect(d.i.Context(), pstore.PeerInfo{
			ID: pId,
			Addrs: addrs,
		})
		if err != nil {
			logger.Warning(err)
		}
	}
	return nil
}

//Save ourself at least in the address book.
func (d *Decentralizer) saveSelf() error {
	self, err := d.peers.FindByPeerId(d.i.Identity.Pretty())
	if err != nil || self == nil {
		//Add self
		err = d.UpsertPeer(d.i.Identity.Pretty(), map[string]string{
			"name": randomdata.SillyName(),
		})
		if err != nil {
			return err
		}
		d.uploadPeers()
	}
	return nil
}

func (d *Decentralizer) UpsertPeer(pId string, details map[string]string) error {
	err := d.peers.Upsert(&pb.Peer{
		PId:     pId,
		Details: details,
	})
	d.addressBookChanged = true
	return err
}

func (d *Decentralizer) GetPeersByDetails(key, value string) ([]*pb.Peer, error) {
	return d.peers.FindByDetails(key, value)
}

func (d *Decentralizer) GetPeers() ([]*pb.Peer, error) {
	return d.peers.FindAll()
}

func (d *Decentralizer) FindByPeerId(peerId string) (p *pb.Peer, err error) {
	p, err = d.peers.FindByPeerId(peerId)
	if err != nil {
		var id Peer.ID
		id, err = d.decodePeerId(peerId)
		if err != nil {
			return nil, err
		}
		p, err = d.getPeerRequest(id)
		if err != nil {
			return nil, err
		}
		d.peers.Upsert(p)
	}
	return p, err
}

//Request peer info from an external peer.
func (d *Decentralizer) getPeerRequest(peer Peer.ID) (*pb.Peer, error) {
	stream, err := d.i.PeerHost.NewStream(d.i.Context(), peer, GET_PEER_REQ)
	if err != nil {
		return nil, err
	}
	stream.SetDeadline(time.Now().Add(MESSAGE_DEADLINE))
	defer stream.Close()

	//Request
	reqData, err := proto.Marshal(&pb.DNPeerRequest{})
	if err != nil {
		return nil, err
	}
	err = framed.Write(stream, reqData)
	if err != nil {
		return nil, err
	}

	//Response
	resData, err := framed.Read(stream)
	if err != nil {
		return nil, err
	}
	var response pb.DNPeerResponse
	err = proto.Unmarshal(resData, &response)
	if err != nil {
		return nil, err
	}

	//Save addr so we can quickly connect to our contacts
	if response.Peer != nil {
		info := d.i.Peerstore.PeerInfo(peer)
		response.Peer.Addrs = []string{}
		for _, addr := range info.Addrs {
			response.Peer.Addrs = append(response.Peer.Addrs, addr.String())
		}
		d.i.Peerstore.AddAddrs(peer, ma.Split(stream.Conn().RemoteMultiaddr()), 3 * 24 * time.Hour)//Save it for 3 days.
	}
	return response.Peer, nil
}

func (d *Decentralizer) FindByDecentralizedId(decentralizedId uint64) (*pb.Peer, error) {
	peer, err := d.peers.FindByDecentralizedId(decentralizedId)
	if err != nil || peer == nil {
		peerId, err := d.resolveDecentralizedId(decentralizedId)
		if err != nil {
			return nil, err
		}
		return d.FindByPeerId(peerId.Pretty())
	}
	return peer, err
}

//Called when an external peer asks for our peer info.
func (d *Decentralizer) getPeerResponse(stream inet.Stream) {
	reqData, err := framed.Read(stream)
	if err != nil {
		logger.Error(err)
		return
	}
	var request pb.DNPeerRequest
	err = proto.Unmarshal(reqData, &request)
	if err != nil {
		logger.Error(err)
		return
	}
	peer, err := d.peers.FindByPeerId(d.i.Identity.Pretty())
	if err != nil {
		logger.Error(err)
		return
	}

	//Response
	response, err := proto.Marshal(&pb.DNPeerResponse{
		Peer: peer,
	})
	if err != nil {
		logger.Error(err)
		return
	}
	err = framed.Write(stream, response)
	if err != nil {
		logger.Error(err)
		return
	}
}
