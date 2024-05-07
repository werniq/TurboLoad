package p2p

import (
	"100gombs/logger"
	"context"
	"flag"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"os"
	"os/signal"
	"syscall"
)

const (
	ProtocolId = "example/"
)

func CreateHost() {
	peerAddr := flag.String("peer-address", "", "peer address")
	flag.Parse()

	host, err := libp2p.New()
	if err != nil {
		logger.ErrorLogger.Fatalln(err)
	}

	host.SetStreamHandler(ProtocolId, func(s network.Stream) {
		go writeCounter(s)
		go readCounter(s)
	})

	fmt.Println(host.Addrs())

	if *peerAddr != "" {
		// Parse the multiaddr string.
		peerMA, err := multiaddr.NewMultiaddr(*peerAddr)
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
	}

	host.SetStreamHandler(ProtocolId, func(s network.Stream) {
		go writeCounter(s)
		go readCounter(s)
	})

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGKILL, syscall.SIGINT)
	<-sigCh
}
