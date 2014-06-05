package entries

import (
	"code"
)

// Makes a front end that talks to one single backend
// Used in Lab1
func MakeFrontSingle(back string) code.Server {
	return code.NewFront(code.NewBinClient([]string{back}))
}

// Serve as a single backend.
// Listen on addr, using s as underlying storage.
func ServeBackSingle(addr string, s code.Storage, ready chan<- bool) error {
	back := &code.BackConfig{
		Addr:  addr,
		Store: s,
		Ready: ready,
	}

	return code.ServeBack(back)
}
