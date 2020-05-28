package ruisIpfs

import (
	"os"
	"testing"
)

func Test1(t *testing.T) {
	HttpErr = make(chan error)
	IpfsRepoPath = "D:\\tmp\\ipfs"
	//ruisBitswap.MFilter = mgr.RuisFter
	IpfsBoots = []string{"/dns4/hw.1ydt.cn/tcp/4001/ipfs/QmT3yVqgo8kiCkobKui4BFmYoJYqZA816juPbGEGdzYzow"}
	cmds := []string{os.Args[0], "daemon", "--init" /*, "--enable-gc"*/}
	MainRet(cmds)
}
