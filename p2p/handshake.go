package p2p

type HandshakeFunc func(any) error

func NOPHandshakeFunc() HandshakeFunc { return nil }
