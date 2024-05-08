package p2p

type HandshakeFunc func(any) error

func NOPHandshakeFunc(a any) error {
	return nil
}
