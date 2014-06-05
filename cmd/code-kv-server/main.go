package main

import (
	"flag"
	"log"

	"code/store"
	"code/entries"
	"code/randomaddr"
)

var (
	addr = flag.String("addr", "localhost:rand", "server listen address")
)

func main() {
	flag.Parse()

	*addr = randomaddr.Resolve(*addr)

	s := store.NewStorage()

	log.Printf("key-value store serving on %s", *addr)

	e := entries.ServeBackSingle(*addr, s, nil)
	if e != nil {
		log.Fatal(e)
	}
}
