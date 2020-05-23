module ruisIpfs

go 1.14

require (
	github.com/hashicorp/go-multierror v1.1.0
	github.com/ipfs/go-ipfs v0.5.1
	github.com/ipfs/go-ipfs-cmds v0.2.2
	github.com/ipfs/go-ipfs-config v0.5.3
	github.com/ipfs/go-ipfs-files v0.0.8
	github.com/ipfs/go-ipfs-util v0.0.1
	github.com/ipfs/go-log v1.0.4
	github.com/ipfs/go-metrics-prometheus v0.0.2
	github.com/jbenet/goprocess v0.1.4
	github.com/libp2p/go-libp2p-loggables v0.1.0
	github.com/libp2p/go-socket-activation v0.0.2
	github.com/multiformats/go-multiaddr v0.2.1
	github.com/multiformats/go-multiaddr-dns v0.2.0
	github.com/multiformats/go-multiaddr-net v0.1.5
	github.com/prometheus/client_golang v1.5.1
)

replace github.com/ipfs/go-bitswap => github.com/mgr9525/go-bitswap v0.2.4-0.20200521125357-4c7a178d5709
