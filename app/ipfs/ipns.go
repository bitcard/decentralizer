package ipfs

import (
	"gx/ipfs/QmYHpXQEWuhwgRFBnrf4Ua6AZhcqXCYa7Biv65SLGgTgq5/go-ipfs/core"
	"gx/ipfs/QmYHpXQEWuhwgRFBnrf4Ua6AZhcqXCYa7Biv65SLGgTgq5/go-ipfs/path"
	"github.com/iain17/timeout"
	"time"
	"context"
)

func FilePublish(n *core.IpfsNode, pth path.Path) error {
	// verify the path exists
	_, err := core.Resolve(n.Context(), n.Namesys, n.Resolver, pth)
	if err != nil {
		return err
	}
	k, err := n.GetKey("self")
	if err != nil {
		return err
	}

	completed := false
	timeout.Do(func(ctx context.Context) {
		err = n.Namesys.Publish(n.Context(), k, pth)
		completed = true
	}, 15*time.Second)
	//if !completed {
	//	err = errors.New("could not publish file in under 15 seconds. Check if you are connected to enough peers")
	//}
	return err
}
