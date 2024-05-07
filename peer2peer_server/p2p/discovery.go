package p2p

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p/core/discovery"
	"github.com/libp2p/go-libp2p/core/peer"
)

type DiscoveryNotifee struct{}

// Discovery interface will be used to search for peers
type Discovery interface {
	FindPeers(ctx context.Context, ns string, opts ...discovery.Option) (<-chan peer.AddrInfo, error)
}

type Advertiser interface {
}

type Discoverer interface {
	Discovery
	Advertiser
}

func (n *DiscoveryNotifee) HandlePeerFound(peerInfo peer.AddrInfo) {
	fmt.Println("found peer", peerInfo.String())
}
