package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTCPTransport(t *testing.T) {
	addr := ":4000"
	tt := NewTCPTransport(addr)

	assert.Equal(t, tt.listenerAddress, addr)

	assert.Nil(t, tt.ListenAndAccept())

	select {}
}
