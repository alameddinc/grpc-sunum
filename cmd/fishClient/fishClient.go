package main

import (
	"log"
	_ "net/http/pprof"
	"os"
	"server/internal"
)

func main() {
	// Connection
	fishClient := internal.NewFishClient(os.Args[1])
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
	fishClient.HighScore()
}
