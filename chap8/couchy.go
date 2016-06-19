package main

import (
	"log"

	// This is not the official client so I don't
	// care if this works actually
	"github.com/davecgh/go-spew/spew"
)

func main() {
	conn, err := couchbase.Connect("http://localhost:8091/")
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}

	for _, pn := range conn.Info.Pools {
		spew.Dump(pn.Name)
		spew.Dump(pn.URI)
		// fmt.Printf("Found pool:  %s -> %s\n", pn.Name, pn.URI)
	}
}
