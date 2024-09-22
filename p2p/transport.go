package p2p

// Peer is an interface that represents a remote node on the network
type Peer interface{}

// Transport is anything that handles communication between nodes on the network
// can be any form like a TCP , UDP, websocket
type Transport interface {
	ListenAndAccept() error
}
