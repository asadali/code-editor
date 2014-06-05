package code 

import (
	"net"
	"net/rpc"
	"net/http"

	//"fmt"
)
func NewBinClient(backs []string) BinStorage {
	return &BinSClient{backs : backs}
}

func NewFront(s BinStorage) Server {
	return &AnybaseService{ bin : s }
}

func (self *BinSClient) Bin(binName string) Storage {
	return NewDSClient(self.backs, binName)
}

// Serve as a backend based on the given configuration
func ServeBack(b *BackConfig) error {

	server := rpc.NewServer()
	err := server.Register(b.Store)
	if err != nil {
		if b.Ready != nil {
			b.Ready <- false
		}
		return err
	}

	listener, err := net.Listen("tcp", b.Addr)
	if err != nil {
		if b.Ready != nil {
			b.Ready <-false
		}
		return err
	}
	if b.Ready != nil {
		b.Ready <- true
	}
	return http.Serve(listener, server)
}