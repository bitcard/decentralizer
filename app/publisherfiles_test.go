package app

import (
	"testing"
	"context"
	"github.com/iain17/decentralizer/app/ipfs"
	"github.com/stretchr/testify/assert"
	"github.com/iain17/decentralizer/pb"
	"time"
	"github.com/iain17/logger"
)

func TestDecentralizer_updatePublisherDefinition(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx,6)
	master := fakeNew(nodes[0], true)
	assert.NotNil(t, master)

	definition := &pb.PublisherDefinition{
		Files: map[string][]byte{
			"hello.txt": []byte("Hard work, by these words guarded. Please don't steal."),
		},
	}
	err := master.publishPublisherUpdate(definition)
	assert.NoError(t, err)
	assert.NotNil(t, master.publisherUpdate)

	//Now start the slave
	slave := fakeNew(nodes[1], false)
	assert.NotNil(t, slave)
	time.Sleep(3 * time.Second)
	assert.NotNil(t, slave.publisherUpdate)
	assert.Equal(t, []byte("Hard work, by these words guarded. Please don't steal."), slave.publisherUpdate.Definition.Files["hello.txt"])
}

func TestDecentralizer_publishPublisherUpdate(t *testing.T) {
	const num = 5
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx, num)
	master := fakeNew(nodes[0], true)
	assert.NotNil(t, master)

	definition := &pb.PublisherDefinition{
		Details: map[string]string{
			"cool": "1",
		},
	}
	err := master.publishPublisherUpdate(definition)
	assert.NoError(t, err)

	//Now start the slaves. Master got an update.
	var slaves []*Decentralizer
	for i := 1; i < num; i++ {
		slave := fakeNew(nodes[i], false)
		assert.NotNil(t, slave)
		slaves = append(slaves, slave)
	}
	//A slave can't publish.
	err = slaves[0].publishPublisherUpdate(definition)
	assert.Error(t, err)

	for i := 0; i < num - 1; i++ {
		slaves[i].updatePublisherDefinition()
		assert.NotNil(t, slaves[i].publisherUpdate)
	}

	time.Sleep(1 * time.Second)
	//Do a update
	definition = &pb.PublisherDefinition{
		Details: map[string]string{
			"cool": "2",
		},
	}
	err = master.publishPublisherUpdate(definition)
	assert.NoError(t, err)

	//Check the rolling update
	numNodesOnOldUpdate := num
	refreshes := 0
	for numNodesOnOldUpdate > 0 {
		numNodesOnOldUpdate = 0
		for i := 0; i < num - 1; i++ {
			slaves[i].updatePublisherDefinition()
			if slaves[i].publisherUpdate.Definition.Details["cool"] == "1" {
				numNodesOnOldUpdate++
			}
		}
		logger.Infof("Number of nodes still on old update %d", numNodesOnOldUpdate)
		refreshes++
	}
	assert.True(t, refreshes < 4, "It should take less than 4 refreshes to get all nodes updated")
}