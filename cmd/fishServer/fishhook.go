package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"server/internal/fishServer"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	fishServer.NewFishServer(grpcServer)
	log.Print("gRPC Fish Server Starting... localhost:9000")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
