package main

import (
	"log"
	"stj-ecommerce/api"
)

func main() {
	server, err := api.NewServer(".")
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}
	server.Start()
}
