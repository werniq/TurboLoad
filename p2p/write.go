package p2p

import (
	"encoding/binary"
	"fmt"
	"github.com/libp2p/go-libp2p/core/network"
)

func readCounter(s network.Stream) {
	for {
		var counter uint64

		err := binary.Read(s, binary.BigEndian, &counter)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Received %d from %s\n", counter, s.ID())
	}
}
