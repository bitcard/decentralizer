package app

import (
	"gx/ipfs/QmNUKMfTHQQpEwE8bUdv5qmKC3ymdW7zw82LFS8D6MQXmu/go-ipfs/core/coreapi"
	//"gx/ipfs/QmNp85zy9RLrQ5oQD4hPyS39ezrrXpcaa7R4Y9kxdWQLLQ/go-cid"
	"bytes"
	"github.com/iain17/decentralizer/app/ipfs"
	"gx/ipfs/QmNUKMfTHQQpEwE8bUdv5qmKC3ymdW7zw82LFS8D6MQXmu/go-ipfs/core/coreapi/interface"
	"gx/ipfs/QmNUKMfTHQQpEwE8bUdv5qmKC3ymdW7zw82LFS8D6MQXmu/go-ipfs/core/coreunix"
	Path "gx/ipfs/QmNUKMfTHQQpEwE8bUdv5qmKC3ymdW7zw82LFS8D6MQXmu/go-ipfs/path"
	Peer "gx/ipfs/QmXYjuNuxVzXKJCfWasQk1RqkhVLDM9jtUKhqc2WPQmFSB/go-libp2p-peer"
	"io/ioutil"
)

func (d *Decentralizer) SavePeerFile(name string, data []byte) (string, error) {
	location, path, err := coreunix.AddWrapped(d.i, bytes.NewBuffer(data), name)
	if err != nil {
		return "", err
	}
	err = ipfs.FilePublish(d.i, Path.FromCid(path.Cid()))
	if err != nil {
		return "", err
	}
	return location, nil
}

//id, err := Peer.IDFromString(peer)
func (d *Decentralizer) GetPeerFiles(id Peer.ID) ([]*iface.Link, error) {
	api := coreapi.NewCoreAPI(d.i)
	rawPath := "/ipns/" + id.Pretty()
	pth := coreapi.ResolvedPath(rawPath, nil, nil)
	return api.Unixfs().Ls(d.i.Context(), pth)
}

//Path could be "/ipfs/QmQy2Dw4Wk7rdJKjThjYXzfFJNaRKRHhHP5gHHXroJMYxk"
func (d *Decentralizer) GetFile(path string) ([]byte, error) {
	api := coreapi.NewCoreAPI(d.i)

	pth := coreapi.ResolvedPath(path, nil, nil)
	_, err := api.ResolvePath(d.i.Context(), pth)
	if err != nil {
		return nil, err
	}
	r, err := api.Unixfs().Cat(d.i.Context(), pth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}
