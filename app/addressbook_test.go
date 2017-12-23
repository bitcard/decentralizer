package app

import (
	"testing"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/iain17/decentralizer/app/ipfs"
)

func TestDecentralizer_FindSelf(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx,2)
	app1 := fakeNew(nodes[0], false)
	assert.NotNil(t, app1)

	err := app1.UpsertPeer("self", map[string]string{
		"quote": "these violent delights have violent ends",
	})
	assert.NoError(t, err)

	peer, err := app1.FindByPeerId("self")
	assert.NoError(t, err)
	assert.NotNil(t, peer)
	assert.Equal(t, peer.Details["quote"], "these violent delights have violent ends")
}

func TestDecentralizer_FindByPeerIdAndUpdate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx,2)
	app1 := fakeNew(nodes[0], false)
	assert.NotNil(t, app1)
	app2 := fakeNew(nodes[1], false)
	assert.NotNil(t, app2)

	//Simple set
	err := app1.UpsertPeer("self", map[string]string{
		"quote": "these violent delights have cool beginnings",
	})
	assert.NoError(t, err)

	//Update
	err = app2.UpsertPeer("self", map[string]string{
		"quote": "these violent delights have violent ends",
	})
	assert.NoError(t, err)
	err = app2.UpsertPeer("self", map[string]string{
		"quote": "these violent delights have violent beginnings",
	})

	//Find myself app1 find app1.
	peer, err := app1.FindByPeerId(app1.i.Identity.Pretty())
	assert.NoError(t, err)
	assert.NotNil(t, peer)
	assert.Equal(t, "these violent delights have cool beginnings", peer.Details["quote"])

	//App 1 find app 2.
	peer, err = app1.FindByPeerId(app2.i.Identity.Pretty())
	assert.NoError(t, err)
	assert.NotNil(t, peer)
	assert.Equal(t, "these violent delights have violent beginnings", peer.Details["quote"])
}

func TestDecentralizer_FindUnknownId(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx,2)
	app1 := fakeNew(nodes[0], false)
	assert.NotNil(t, app1)
	app2 := fakeNew(nodes[1], false)
	assert.NotNil(t, app2)

	err := app2.UpsertPeer("self", map[string]string{
		"quote": "these violent delights have violent ends",
	})
	assert.NoError(t, err)

	peer, err := app2.FindByPeerId("self")
	assert.NoError(t, err)
	assert.NotNil(t, peer)

	DnId := peer.DnId
	//Now have app2 find us with just DnId.
	peer, err = app1.FindByDecentralizedId(DnId)
	assert.NoError(t, err)
	assert.NotNil(t, peer)
	assert.Equal(t, peer.Details["quote"], "these violent delights have violent ends")
}