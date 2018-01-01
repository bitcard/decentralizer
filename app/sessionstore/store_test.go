package sessionstore

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
	"github.com/iain17/decentralizer/pb"
	libp2pPeer "gx/ipfs/QmXYjuNuxVzXKJCfWasQk1RqkhVLDM9jtUKhqc2WPQmFSB/go-libp2p-peer"
)

func TestSessionsStore_FindByPeerId(t *testing.T) {
	self, err := libp2pPeer.IDB58Decode("QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ381")
	store, err := New(10, 1 * time.Hour, self)
	assert.NoError(t, err)
	_, err = store.Insert(&pb.Session {
		Address: 1,
		Port: 1,
		Type: 1,
		PId: "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ381",
		Details: map[string]string{
			"hey": "ho",
		},
	})
	assert.NoError(t, err)
	_, err = store.Insert(&pb.Session {
		Address: 1,
		Port: 2,
		Type: 1,
		PId: "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ382",
		Details: map[string]string{
			"hey": "ho",
		},
	})
	assert.NoError(t, err)

	sessions, err := store.FindByPeerId("QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ381")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(sessions))
	assert.Equal(t, sessions[0].PId, "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ381")
}

func TestSessionsStore_Expire(t *testing.T) {
	self, err := libp2pPeer.IDB58Decode("QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ381")
	store, err := New(10, 1 * time.Second, self)
	assert.NoError(t, err)
	//Our own sessions never expire.
	_, err = store.Insert(&pb.Session {
		Address: 1,
		Port: 1,
		DnId: 1,
		PId: "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ381",
		Details: map[string]string{
			"hey": "ho",
		},
	})
	assert.NoError(t, err)
	//others do.
	_, err = store.Insert(&pb.Session {
		Address: 1,
		Port: 2,
		DnId: 2,
		PId: "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ385",
		Details: map[string]string{
			"hey": "ho",
		},
	})
	assert.NoError(t, err)

	sessions, err := store.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, 2, len(sessions))
	time.Sleep(2 * time.Second)
	sessions, err = store.FindAll()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(sessions))
}

func TestSessionsStore_Limit(t *testing.T) {
	self, err := libp2pPeer.IDB58Decode("QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ381")
	store, err := New(1, 2 * time.Hour, self)
	assert.NoError(t, err)
	//Because self has added this. we'll have 2
	_, err = store.Insert(&pb.Session {
		Address: 1,
		Port: 1,
		DnId: 1,
		PId: "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ381",
		Details: map[string]string{
			"hey": "ho",
		},
	})
	_, err = store.Insert(&pb.Session {
		Address: 1,
		Port: 2,
		DnId: 2,
		PId: "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ388",
		Details: map[string]string{
			"hey": "ho",
		},
	})
	assert.NoError(t, err)
	_, err = store.Insert(&pb.Session {
		Address: 1,
		Port: 3,
		DnId: 3,
		PId: "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ385",
		Details: map[string]string{
			"hey": "ho",
		},
	})
	assert.NoError(t, err)
	time.Sleep(1 * time.Second)
	sessions, err := store.FindAll()
	assert.NoError(t, err)
	//Because the size is actually 1. but we can't delete a session we created.
	assert.Equal(t, 2, len(sessions))
}

func TestSessionsStore_FindByDetails(t *testing.T) {
	self, err := libp2pPeer.IDB58Decode("QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ381")
	store, err := New(10, 1 * time.Hour, self)
	assert.NoError(t, err)
	_, err = store.Insert(&pb.Session {
		Address: 1,
		Port: 1,
		DnId: 1,
		PId: "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ381",
		Details: map[string]string{
			"cool": "yes",
		},
	})
	assert.NoError(t, err)
	_, err = store.Insert(&pb.Session {
		Address: 1,
		Port: 2,
		DnId: 2,
		PId: "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ382",
		Details: map[string]string{
			"cool": "no",
		},
	})
	assert.NoError(t, err)
	sessions, err := store.FindByDetails("cool", "no")
	assert.NoError(t, err)
	assert.Equal(t, sessions[0].PId, "QmTq1jNgbHarKgYkZfLJAmtUewyWYniTupQf7ZYsSFQ382")
	assert.Equal(t, len(sessions), 1)
}