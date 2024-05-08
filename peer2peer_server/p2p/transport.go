package p2p

// Transport is anything that handlers communication between the nodes
// in the network. It can be of the form (TCP, UDP, websockets)
type Transport interface {
	ListenAndAccept() error
}

// Peer is an interface that represents remote node.
type Peer interface {
	Connect(addr string) (bool, error)
	Write(data []byte) error
	Receive(data []byte) error
}
