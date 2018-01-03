package ipfs

import (
	"gx/ipfs/QmYHpXQEWuhwgRFBnrf4Ua6AZhcqXCYa7Biv65SLGgTgq5/go-ipfs/core/coreapi"
	"bytes"
	"errors"
	"gx/ipfs/QmYHpXQEWuhwgRFBnrf4Ua6AZhcqXCYa7Biv65SLGgTgq5/go-ipfs/core"
	"github.com/iain17/logger"
	"gx/ipfs/QmYHpXQEWuhwgRFBnrf4Ua6AZhcqXCYa7Biv65SLGgTgq5/go-ipfs/core/coreapi/interface"
	"gx/ipfs/QmYHpXQEWuhwgRFBnrf4Ua6AZhcqXCYa7Biv65SLGgTgq5/go-ipfs/core/coreunix"
	Path "gx/ipfs/QmYHpXQEWuhwgRFBnrf4Ua6AZhcqXCYa7Biv65SLGgTgq5/go-ipfs/path"
	"io/ioutil"
	"context"
	libp2pPeer "gx/ipfs/QmXYjuNuxVzXKJCfWasQk1RqkhVLDM9jtUKhqc2WPQmFSB/go-libp2p-peer"
)

//Simplifies all the interactions with IPFS.
type FilesAPI struct {
	ctx					   context.Context
	i 				   	   *core.IpfsNode
	api					   iface.CoreAPI
}

func NewFilesAPI(ctx context.Context, node *core.IpfsNode, api iface.CoreAPI) (*FilesAPI, error) {
	instance := &FilesAPI{
		ctx: ctx,
		i: node,
		api: api,
	}
	return instance, nil
}

func (d *FilesAPI) SavePeerFile(name string, data []byte) (string, error) {
	logger.Infof("Saving peer file %s", name)
	location, path, err := coreunix.AddWrapped(d.i, bytes.NewBuffer(data), name)
	if err != nil {
		return "", err
	}
	ph := Path.FromCid(path.Cid())
	if err != nil {
		return "", err
	}
	err = FilePublish(d.i, ph)
	return "/ipfs/"+location, err
}

func (d *FilesAPI) GetPeerFiles(owner libp2pPeer.ID) ([]*iface.Link, error) {
	logger.Infof("Get peer files of peer id %s", owner.Pretty())
	rawPath := "/ipns/" + owner.Pretty()
	pth := coreapi.ResolvedPath(rawPath, nil, nil)
	//d.api.ResolveNode(d.i.Context(), pth)
	//return coreapi.NewCoreAPI(d.i).Unixfs().Ls(d.i.Context(), pth)
	return d.api.Unixfs().Ls(d.i.Context(), pth)
}

func (d *FilesAPI) GetPeerFile(owner libp2pPeer.ID, name string) ([]byte, error) {
	logger.Infof("Get peer file %s of peer id %s", name, owner.String())
	files, err := d.GetPeerFiles(owner)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		if file.Name == name {
			return d.GetFile(file.Cid.String())
		}
	}
	return nil, errors.New("could not find peer file")
}

//Path could be "/ipfs/QmQy2Dw4Wk7rdJKjThjYXzfFJNaRKRHhHP5gHHXroJMYxk"
func (d *FilesAPI) GetFile(path string) ([]byte, error) {
	logger.Infof("Get file: %s", path)

	pth := coreapi.ResolvedPath(path, nil, nil)
	_, err := d.api.ResolvePath(d.i.Context(), pth)
	if err != nil {
		return nil, err
	}
	r, err := d.api.Unixfs().Cat(d.i.Context(), pth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(r)
}
