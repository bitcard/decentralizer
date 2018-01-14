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

func TestDecentralizer_getPublisherDefinition(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx,5)
	master := fakeNew(ctx, nodes[0], true)
	assert.NotNil(t, master)
	slave := fakeNew(ctx, nodes[1], false)
	assert.NotNil(t, slave)

	definition := &pb.PublisherDefinition{
		Status: true,
		Files: map[string][]byte{
			"hello.txt": []byte("Hard work, by these words guarded. Please don't steal."),
		},
		Details: map[string]string{
			"data": "wtf",
		},
	}

	err := master.PublishPublisherRecord(definition)
	assert.NoError(t, err)
	assert.NotNil(t, master.publisherRecord)
	time.Sleep(500 * time.Millisecond)

	slave.updatePublisherDefinition()

	assert.NotNil(t, slave.publisherRecord)
	if slave.publisherRecord != nil {
		assert.Equal(t, []byte("Hard work, by these words guarded. Please don't steal."), slave.publisherDefinition.Files["hello.txt"])
	}
}

func TestDecentralizer_publishPublisherUpdate(t *testing.T) {
	const num = 20
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx, num)
	master := fakeNew(ctx, nodes[0], true)
	assert.NotNil(t, master)

	//Start slaves
	var slaves []*Decentralizer
	for i := 1; i < num; i++ {
		slave := fakeNew(ctx, nodes[i], false)
		assert.NotNil(t, slave)
		slaves = append(slaves, slave)
	}

	definition := &pb.PublisherDefinition{
		Status: true,
		Details: map[string]string{
			"cool": "1",
		},
	}
	//start master
	err := master.PublishPublisherRecord(definition)
	assert.NoError(t, err)

	//A slave can't publish.
	err = slaves[0].PublishPublisherRecord(definition)
	assert.Error(t, err)

	//Check the rolling update
	numNodesOnOldUpdate := num
	refreshes := 0
	for numNodesOnOldUpdate > 0 {
		numNodesOnOldUpdate = 0
		for i := 0; i < num - 1; i++ {
			slaves[i].updatePublisherDefinition()
			if slaves[i].publisherDefinition == nil {
				numNodesOnOldUpdate++
			}
		}
		if numNodesOnOldUpdate == 0 {
			break
		}
		logger.Infof("Number of nodes still on old update %d", numNodesOnOldUpdate)
		refreshes++
		time.Sleep(1 * time.Second)
	}
	assert.True(t, refreshes < 10, "It should take less than 10 (actual=%d) refreshes to get all nodes updated", refreshes)
	time.Sleep(1 * time.Second)
	//Do a update
	definition = &pb.PublisherDefinition{
		Status: true,
		Details: map[string]string{
			"cool": "2",
		},
	}
	err = master.PublishPublisherRecord(definition)
	assert.NoError(t, err)

	//Check the rolling update
	numNodesOnOldUpdate = num
	refreshes = 0
	for numNodesOnOldUpdate > 0 {

		numNodesOnOldUpdate = 0
		for i := 0; i < num - 1; i++ {
			slaves[i].updatePublisherDefinition()
			if slaves[i].publisherDefinition.Details["cool"] == "1" {
				numNodesOnOldUpdate++
			} else {
				//People that have updated. Should also be able to set this DHT value.
				err = slaves[i].PushPublisherUpdate()
				assert.NoError(t, err)
			}
		}
		if numNodesOnOldUpdate == 0 {
			break
		}
		logger.Infof("Number of nodes still on old update %d", numNodesOnOldUpdate)
		refreshes++
		time.Sleep(1 * time.Second)
	}
	assert.True(t, refreshes < 10, "It should take less than 10 (actual=%d) refreshes to get all nodes updated", refreshes)
}

//If the publisher has set the network to status false. Stop the process.
func TestDecentralizer_publishStopper(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nodes := ipfs.FakeNewIPFSNodes(ctx,2)
	app1 := fakeNew(ctx, nodes[0], false)
	assert.NotNil(t, app1)
	// Mocked publisher update
	app1.publisherDefinition = &pb.PublisherDefinition{
		Status: false,
	}
	app1.runPublisherInstructions()
}