package main

import (
	"github.com/werniq/peer2peer_turboload/logger"
	"github.com/werniq/peer2peer_turboload/p2p"
	"log"
	"net"
	"time"
)

type Node struct {
	ID   string
	Addr string
}

var mainNode *Node
var nodes chan *Node

func main() {
	// Steps to implement:
	// 1. Create peer-to-peer network
	//network := createNetwork()
	//
	//// 2. Launch main node
	//mainNode = &Node{
	//	ID:   "main",
	//	Addr: "localhost:8081",
	//}
	//go launchMainNode(network)
	//
	//// 4. Listen for new nodes
	//go listenForNewNodes(network)
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}

	tt := p2p.NewTCPTransport(tcpOpts)

	if err := tt.ListenAndAccept(); err != nil {
		log.Fatalln(err)
	}

	// Keep the main goroutine alive
	select {}
}

// createNetwork creates main 'server' that will send files to all newly joined
// participants in the network
func createNetwork() *net.UDPConn {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		logger.ErrorLogger.Fatalln("Error resolving UDP address:", err)
		return nil
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		logger.ErrorLogger.Fatalln("Error listening on udp connection:", err)
		return nil
	}

	logger.InfoLogger.Println("Successfully initialized network.")
	return conn
}

// launchMainNode will send files to new participants in the network
func launchMainNode(conn *net.UDPConn) {
	// Send files to joined participants
	for {
		for node := range nodes {
			logger.ErrorLogger.Fatalln("Sending file to node:", node.ID)
			// Simulate sending file
			// TODO: implement sending file to new node
			time.Sleep(time.Second)
			// after file is sent, remove node from network
		}
		// Wait before sending files again
		time.Sleep(time.Second * 5)
	}
}

// listenForNewNodes will handle adding nodes to the pool and sending files to them
func listenForNewNodes(conn *net.UDPConn) {
	for {
		buffer := make([]byte, 1024)
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			logger.ErrorLogger.Fatalln("Error reading from UDP:", err)
			continue
		}

		nodeID := string(buffer[:n])
		newNode := &Node{
			ID:   nodeID,
			Addr: addr.String(),
		}

		nodes <- newNode
		logger.InfoLogger.Println("New node joined:", newNode.ID)
	}
}
