package p2p

import (
	"fmt"
	"github.com/werniq/peer2peer_turboload/logger"
	"net"
	"sync"
)

type UDPTransport struct {
	listenerAddr string
	listener     *net.UDPConn

	my    sync.RWMutex
	peers map[net.Addr]Peer
}

// NewUDPTransport initializes new listener on client's system to listen on listenerAddr with
// connection type = UDP
func NewUDPTransport(listenerAddr string) *UDPTransport {
	return &UDPTransport{
		listenerAddr: listenerAddr,
		listener:     nil,
		my:           sync.RWMutex{},
		peers:        nil,
	}
}

// SendDataToRemoteAddress will be used to send data through udp connection
func (u *UDPTransport) SendDataToRemoteAddress(data chan []byte, addr string) {
	remoteAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		logger.ErrorLogger.Fatalln("Error resolving UDP address", err)
	}

	for d := range data {
		_, err = u.listener.WriteToUDP(d, remoteAddr)
		if err != nil {
			logger.ErrorLogger.Fatalln("Error resolving UDP address", err)
		}
	}
}

func (u *UDPTransport) ListenAndAccept() {
	addr := fmt.Sprintf("127.0.0.1%s", u.listenerAddr)

	n, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		logger.ErrorLogger.Fatalln("Error resolving UDP address", err)
	}

	u.listener, err = net.ListenUDP("udp", n)
	if err != nil {
		logger.ErrorLogger.Fatalln("Error resolving UDP address", err)
	}

}

func (u *UDPTransport) acceptLoop() {
	for {
	}
}
