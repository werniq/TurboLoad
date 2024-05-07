package p2p

import (
	"github.com/libp2p/go-libp2p/core/network"
	"io"
	"os"
)

func writeCounter(s network.Stream) {
	// Open the file
	var n int

	out, err := os.Open("../files/1GB.bin")
	if err != nil {
		panic(err)
	}

	defer out.Close()

	buffer := make([]byte, 1024)

	for {
		// Read from the file into the buffer
		n, err = out.Read(buffer)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}

		// Write the buffer to the stream
		_, err = s.Write(buffer[:n])
		if err != nil {
			panic(err)
		}
	}
}
