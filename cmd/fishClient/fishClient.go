package main

import (
	"log"
	_ "net/http/pprof"
	"os"

	client "github.com/alameddinc/grpc-sunum/internal/fClient"
)

func main() {
	// Connection
	fishClient := client.NewFishClient(os.Args[1])
	defer fishClient.Close()
	// Register (Server side stream RPC)
	if err := fishClient.Register(); err != nil {
		log.Fatal(err)
	}
	// Bidirection Connection...
	if err := fishClient.ListenAndCatch(); err != nil {
		log.Fatal(err)
	}
	// Simple RPC
	fishClient.Highscore()
}
