package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *Message) error
}

type GOBDecoder struct{}

func (d GOBDecoder) Decode(r io.Reader, msg *Message) error {
	return gob.NewDecoder(r).Decode(msg)
}

type DefaultDecoder struct {
}

func (n DefaultDecoder) Decode(r io.Reader, msg *Message) error {
	buf := make([]byte, 1024)
	l, err := r.Read(buf)
	if err != nil {
		return err
	}

	msg.Payload = buf[:l]

	return nil
}
