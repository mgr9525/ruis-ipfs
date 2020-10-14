package ruisIpfs

import (
	"context"
	"os"
	"testing"
)

func Test1(t *testing.T) {
	RuisErr = make(chan error)
	RuisRepoPath = "D:\\tmp\\ipfs"
	//insert go.mod
	//replace github.com/ipfs/go-bitswap => github.com/mgr9525/go-bitswap v0.2.4-0.20200521125357-4c7a178d5709
	//ruisBitswap.MFilter = mgr.RuisFter
	RuisBoots = []string{"/dns4/hw.1ydt.cn/tcp/4001/ipfs/QmT3yVqgo8kiCkobKui4BFmYoJYqZA816juPbGEGdzYzow"}
	cmds := []string{os.Args[0], "daemon", "--init" /*, "--enable-gc"*/}
	MainRet(context.Background(), cmds)
}
