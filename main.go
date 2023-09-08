package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/tapansaikia/slot-book-server/api"
	"github.com/tapansaikia/slot-book-server/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8080", "the server listen address")
	flag.Parse()

	store := storage.NewMongoStorage()

	server := api.NewServer(*listenAddr, store)
	fmt.Println("Server running on port:", *listenAddr)
	log.Fatal(server.Start())
}
