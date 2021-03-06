package peerstore

import (
	libp2pPeer "gx/ipfs/QmdVrMn1LhB4ybb8hMVaMLXnA8XRSewMnK6YqXKXoTcRvN/go-libp2p-peer"
	"hash/fnv"
)

//libp2p peer id to uint64. Some apps expect some identification in the form of an integer. This will make it so we are compatible.
func PeerToDnId(id libp2pPeer.ID) (pId string, dID uint64) {
	h := fnv.New32a()
	pId = id.Pretty()
	h.Write([]byte(pId))
	dID = uint64(h.Sum32())
	return
}
