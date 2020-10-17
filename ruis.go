package ruisIpfs

import (
	"github.com/ipfs/go-ipfs/core"
	"github.com/ipfs/go-ipfs/core/coreapi"
	coreiface "github.com/ipfs/interface-go-ipfs-core"
)

var (
	nodeInstance *core.IpfsNode

	RuisErr      chan error
	RuisRepoPath = ""
	RuisPeerId   string
	RuisKey      string
	RuisBoots    []string

	RuisStorageMax         string
	RuisStorageGCWatermark int64
)

func GetNode() *core.IpfsNode {
	return nodeInstance
}
func GetApi() (coreiface.CoreAPI, error) {
	return coreapi.NewCoreAPI(nodeInstance)
}
