package ruisIpfs

import (
	cmds "github.com/ipfs/go-ipfs-cmds"
	"github.com/ipfs/go-ipfs/commands"
)

var (
	IpfsRepoPath = `D:\tmp\ipfs`
	IpfsId       string
	IpfsKey      string
	IpfsBoots    []string

	IpfsStorageMax         string
	IpfsStorageGCWatermark int64

	daemonReq  *cmds.Request
	daemonCmds *commands.Context

	HttpErr chan error
)

func GetDaemonReq() *cmds.Request {
	return daemonReq
}
func GetDaemonCmd() *commands.Context {
	return daemonCmds
}
