package p2p

import (
	"encoding/binary"
	"github.com/libp2p/go-libp2p/core/network"
	"time"
)

func writeCounter(s network.Stream) {
	var counter uint64

	for {
		<-time.After(time.Second)
		counter++

		err := binary.Write(s, binary.BigEndian, counter)
		if err != nil {
			panic(err)
		}
	}
}
