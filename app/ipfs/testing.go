package ipfs

import (
	"context"
	"gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/core"
	"gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/core/mock"
	namesys "gx/ipfs/QmTxUjSZnG7WmebrX2U7furEPNSy33pLgA53PtpJYJSZSn/go-ipfs/namesys"
	pstore "gx/ipfs/QmYijbtjCxFEjSXaudaQAUz3LN5VKLssm8WCUsRoqzXmQR/go-libp2p-peerstore"
	"gx/ipfs/Qma23bpHwQrQyvKeBemaeJh7sAoRHggPkgnge1B9489ff5/go-libp2p/p2p/net/mock"
	"time"
)

func FakeNewIPFSNodes(ctx context.Context, numPeers int) []*core.IpfsNode {
	// create network
	mn := mocknet.New(ctx)

	var nodes []*core.IpfsNode
	for i := 0; i < numPeers; i++ {
		n, err := core.NewNode(ctx, &core.BuildCfg{
			Online:    true,
			Permament: true,
			Host:      coremock.MockHostOption(mn),
		})
		if err != nil {
			panic(err)
		}
		n.Namesys = namesys.NewNameSystem(n.Routing, n.Repo.Datastore(), 0)
		nodes = append(nodes, n)
	}

	mn.LinkAll()

	// connect them
	for _, n1 := range nodes {
		for _, n2 := range nodes {
			if n1 == n2 {
				continue
			}
			p2 := n2.PeerHost.Peerstore().PeerInfo(n2.PeerHost.ID())
			if err := n1.PeerHost.Connect(ctx, p2); err != nil {
				panic(err)
			}
		}
	}

	bsinf := core.BootstrapConfigWithPeers(
		[]pstore.PeerInfo{
			nodes[0].Peerstore.PeerInfo(nodes[0].Identity),
		},
	)

	for _, n := range nodes[1:] {
		if err := n.Bootstrap(bsinf); err != nil {
			panic(err)
		}
	}
	time.Sleep(1 * time.Second)

	return nodes
}
