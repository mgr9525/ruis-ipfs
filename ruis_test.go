package ruisIpfs

import (
	"os"
	"testing"
)

func Test1(t *testing.T) {
	HttpErr = make(chan error)
	IpfsRepoPath = "D:\\tmp\\ipfs"
	//ruisBitswap.MFilter = mgr.RuisFter
	cmds := []string{os.Args[0], "daemon", "--init" /*, "--enable-gc"*/}
	MainRet(cmds)
}
