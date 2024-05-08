package p2p

import (
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/multiformats/go-multiaddr"
	"github.com/werniq/peer2peer_turboload/logger"
	"net"
)

type Listener struct {
	listenerAddress string
	listener        net.Listener
	//     peers map[string]Peer
	peers map[net.Addr]Peer
}

var (
	MainNodeAddr = []multiaddr.Multiaddr{}
)

const (
	ProtocolId = "example/"
)

func InitParentNode() host.Host {
	//peerAddr := flag.String("peer-address", "", "peer address")
	//flag.Parse()

	node, err := libp2p.New()
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
	}

	MainNodeAddr = node.Addrs()

	return node
}

// InitNode
// Initializes new node in our P2P network
//
// After initialization we know that we need to send a file to new node, depending on
func (n *Network) InitNode() host.Host {
	node, err := libp2p.New()
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
	}

	return node
}
