package iptbutil

import (
	"gx/ipfs/QmYHpXQEWuhwgRFBnrf4Ua6AZhcqXCYa7Biv65SLGgTgq5/go-ipfs/repo/config"
)

type IpfsNode interface {
	Init() error
	Kill() error
	Start(args []string) error
	APIAddr() (string, error)
	GetPeerID() string
	RunCmd(args ...string) (string, error)
	Shell() error
	String() string

	GetAttr(string) (string, error)
	SetAttr(string, string) error

	GetConfig() (*config.Config, error)
	WriteConfig(*config.Config) error
}
