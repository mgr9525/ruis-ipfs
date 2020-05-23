// +build linux

package ruisIpfs

import (
	daemon "github.com/coreos/go-systemd/v22/daemon"
)

func notifyReady() {
	_, _ = daemon.SdNotify(false, daemon.SdNotifyReady)
}

func notifyStopping() {
	_, _ = daemon.SdNotify(false, daemon.SdNotifyStopping)
}
