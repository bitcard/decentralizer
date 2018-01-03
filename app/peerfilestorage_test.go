package app

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"context"
	"github.com/iain17/decentralizer/app/ipfs"
	"time"
	"github.com/iain17/logger"
)

//One user saves a file. The other gets it by its hash.
func TestDecentralizer_SaveGetFile(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx,2)
	app1 := fakeNew(ctx, nodes[0], false)
	assert.NotNil(t, app1)
	app2 := fakeNew(ctx, nodes[1], false)
	assert.NotNil(t, app2)

	message := []byte("Low in coupling and high in cohesion.")

	cid, err := app2.SavePeerFile("test.txt", message)
	assert.NoError(t, err)
	assert.NotNil(t, cid)

	data, err := app1.getIPFSFile(cid)
	assert.NoError(t, err)
	assert.Equal(t, string(message), string(data))
}

//One user saves a file. The other gets it by its name and the peer id that saved it.
func TestDecentralizer_SaveGetUserFile(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx,2)
	app1 := fakeNew(ctx, nodes[0], false)
	assert.NotNil(t, app1)
	app2 := fakeNew(ctx, nodes[1], false)
	assert.NotNil(t, app2)

	message := []byte("Hey ho this is cool.")
	filename := "test.txt"

	_, err := app1.SavePeerFile(filename, message)
	assert.NoError(t, err)

	file, err := app2.filesApi.GetPeerFile(app1.i.Identity, filename)
	assert.NoError(t, err)
	assert.Equal(t, string(message), string(file), "Should work fine calling directly IPFS")

	file, err = app2.GetPeerFile(app1.i.Identity.Pretty(), filename)
	assert.NoError(t, err)
	assert.Equal(t, string(message), string(file), "Should work fine calling IPFS through cache layer")

	_, err = app2.GetPeerFile(app1.i.Identity.Pretty(), "random shit")
	assert.Error(t, err, "404. doesn't exist")
}

func TestDecentralizer_GetPeerFileUpdated(t *testing.T) {
	FILE_EXPIRE = 0
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx,2)
	app1 := fakeNew(ctx, nodes[0], false)
	assert.NotNil(t, app1)
	app2 := fakeNew(ctx, nodes[1], false)
	assert.NotNil(t, app2)

	message := []byte("Simplicity is the ultimate sophistication ~ Leonardo Da Vinci")
	updatedMessage := []byte("The mass of men lead lives of quiet desperation. What is called resignation is confirmed desperation. ~Henry David Thoreau")
	filename := "test.txt"

	_, err := app1.SavePeerFile(filename, message)
	assert.NoError(t, err)

	_, err = app1.SavePeerFile(filename, updatedMessage)
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	var result []byte
	for i:= 0; i < 10; i++ {
		result, err = app2.GetPeerFile(app1.i.Identity.Pretty(), filename)
		assert.NoError(t, err)
		if string(updatedMessage) == string(result) {
			break
		}
		time.Sleep(1 * time.Second)
	}
	assert.Equal(t, string(updatedMessage), string(result))
}

func TestDecentralizer_GetPeerFileCache(t *testing.T) {
	FILE_EXPIRE = 0
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx,2)
	app1 := fakeNew(ctx, nodes[0], false)
	assert.NotNil(t, app1)
	app2 := fakeNew(ctx, nodes[1], false)
	assert.NotNil(t, app2)

	message := []byte("Simplicity is the ultimate sophistication ~ Leonardo Da Vinci")
	filename := "test.txt"
	_, err := app1.SavePeerFile(filename, message)
	assert.NoError(t, err)

	var result []byte
	for i:= 0; i < 10; i++ {
		result, err = app2.GetPeerFile(app1.i.Identity.Pretty(), filename)
		assert.NoError(t, err)
		if string(message) == string(result) {
			break
		}
		logger.Infof("%s != %s", message, result)
		time.Sleep(1 * time.Second)
	}
	assert.Equal(t, string(message), string(result))

	//Now app1 goes offline. Can app2 still get his data from cache?
	app1.Stop()
	result, err = app2.GetPeerFile(app1.i.Identity.Pretty(), filename)
	assert.NoError(t, err)
	assert.Equal(t, string(message), string(result), "app1 is offline. But I can still fetch his data.")
}
