package p2p

import "net"

// Message is data that being send over each transport between nodes
type Message struct {
	From    net.Addr
	Payload []byte
}
