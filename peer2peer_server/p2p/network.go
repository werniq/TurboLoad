package p2p

import (
	"context"
	"fmt"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/werniq/peer2peer_turboload/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Network struct holds all nodes in our network
type Network struct {
	// Nodes are basically our network
	Nodes map[int]host.Host

	ParentNode host.Host

	Count int

	// mux will be used for concurrency-safety for modifications of Nodes
	mux sync.RWMutex
}

// InitNetwork
// Creates Network struct, which will (concurrently-safe) store every node in network
func InitNetwork() *Network {
	n := &Network{
		Nodes: make(map[int]host.Host),
		Count: 1,
		mux:   sync.RWMutex{},
	}

	// initializing main node
	n.Nodes[n.Count] = InitParentNode()
	n.ParentNode = n.Nodes[n.Count]

	return n
}

// InsertNode adds new node into network and updates counter
func (n *Network) InsertNode(h host.Host) {
	n.mux.Lock()
	n.Nodes[n.Count] = h
	n.Count += 1
	n.mux.Unlock()
}

// AddNodeToNetwork creates new node and inserts it to the main network
// based on user ip address
func (n *Network) AddNodeToNetwork(userIp, hostTcpPort string) {
	host := n.InitNode()

	n.InsertNode(host)

	logger.InfoLogger.Println(host.Addrs())
	logger.InfoLogger.Println(host.ID())

	host.SetStreamHandler(ProtocolId, func(s network.Stream) {
		go writeCounter(s)
		go readCounter(s)
	})
	defer host.Close()

	peerAddr := fmt.Sprintf("/ip4/%s/tcp/%s/p2p/%s", userIp, hostTcpPort, host.ID().String())

	if host.Addrs() != nil {
		// Parse the multiaddr string.
		peerMA, err := multiaddr.NewMultiaddr(peerAddr)
		if err != nil {
			panic(err)
		}

		peerAddrInfo, err := peer.AddrInfoFromP2pAddr(peerMA)
		if err != nil {
			panic(err)
		}

		// Connect to the node at the given address.
		if err = host.Connect(context.Background(), *peerAddrInfo); err != nil {
			panic(err)
		}
		fmt.Println("Connected to", peerAddrInfo.String())

		s, err := host.NewStream(context.Background(), peerAddrInfo.ID)
		if err != nil {
			panic(err)
		}

		// Start the write and read threads.
		go writeCounter(s)
		go readCounter(s)

		file, err := os.Open("./files/1GB.bin")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Buffer to read file chunks
		buffer := make([]byte, 1024)
		for {
			bytesRead, err := file.Read(buffer)
			if err != nil {
				break // End of file
			}

			if bytesRead > 0 {
				_, err = s.Write(buffer[:bytesRead])
				if err != nil {
					panic(err)
				}
			}
		}
	}

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGKILL, syscall.SIGINT)
	<-sigCh
}

// Listen
func (n *Network) Listen() {
	// By

}
