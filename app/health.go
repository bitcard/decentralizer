package app

import (
	"errors"
	"fmt"
	"github.com/iain17/timeout"
	"time"
	"context"
	"github.com/iain17/decentralizer/vars"
)

func (d *Decentralizer) Health(WaitForMinConnections bool) (bool, int, error) {
	//if !d.d.IsReady() {
	//	return false, 0, nil
	//}

	numPeers := len(d.i.PeerHost.Network().Peers())
	if WaitForMinConnections {
		if numPeers < vars.MIN_CONNECTED_PEERS {
			timeout.Do(func(ctx context.Context) {
				for {
					select {
					case <-ctx.Done():
						return
					default:
						numPeers := len(d.i.PeerHost.Network().Peers())
						if numPeers >= vars.MIN_CONNECTED_PEERS {
							return
						}
						time.Sleep(100 * time.Millisecond)
					}
				}
			}, 5*time.Second)
		}
		numPeers = len(d.i.PeerHost.Network().Peers())
		if numPeers < vars.MIN_CONNECTED_PEERS {
			percentage := 0.0
			if numPeers > 0 {
				total := float64(vars.MIN_CONNECTED_PEERS)
				percentage = float64(numPeers) / total * 100
			}
			return false, numPeers, errors.New(fmt.Sprintf("Bootstrapping to ADNA. %.2f %% complete", percentage))
		}
	}

	if d.publisherRecord == nil {
		return false, numPeers, errors.New(fmt.Sprintf("Waiting for publisher file..."))
	}
	if !d.publisherDefinition.Status {
		return false, numPeers, errors.New("closed")
	}
	return true, numPeers, nil
}