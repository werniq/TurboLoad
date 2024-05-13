package p2p

import (
	"github.com/werniq/peer2peer_turboload/logger"
	"net"
	"sync"
)

type TCPTransport struct {
	listenerAddress string
	listener        net.Listener
	//handshakeFunc   HandshakeFunc
	TCPTransportOpts

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

type TCPTransportOpts struct {
	Decoder       Decoder
	ListenAddr    string
	HandshakeFunc HandshakeFunc
}

// TCPPeer represents the node (remote) which listen and accepts data
// over tcp connection
type TCPPeer struct {
	conn     net.Conn
	outbound bool
}

// NewTCPPeer initializes new peer which listens on tcp connection
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

// NewTCPTransport
func NewTCPTransport(opts TCPTransportOpts) *TCPTransport {
	return &TCPTransport{
		//handshakeFunc:   NOPHandshakeFunc,
		//listenerAddress: addr,
		TCPTransportOpts: opts,
	}
}

// ListenAndAccept basically listens on tcp network and accepts incoming connection if it matches
// the rules
func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddr)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

// startAcceptLoop listens for all incoming to listener connections and
// verifies them
func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			logger.InfoLogger.Println("Error while accepting TCP:  ", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error
	peer := NewTCPPeer(conn, true)

	if err = t.HandshakeFunc(peer); err != nil {
		conn.Close()
		logger.ErrorLogger.Println("TCP handshake error: ", err.Error())
		return
	}

	msg := &Message{}
	for {
		if err = t.Decoder.Decode(conn, msg); err != nil {
			logger.ErrorLogger.Printf("TCP decoding error: %v\n", err)
			continue
		}

		msg.From = conn.RemoteAddr()

		logger.InfoLogger.Printf("New message: %+v\n", msg)
	}
}
