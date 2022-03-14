package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	_ "net/http/pprof"
	"server/pkg/protoGo"
	"time"
)

func main() {
	var conn *grpc.ClientConn
	userLocalScore := 0
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := protoGo.NewFishServiceClient(conn)
	// Bidirection Connection...
	stream, err := c.TryToCatch(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	waitC := make(chan struct{})
	// Read Bidirectional
	go func() {
		// todo select ile dene
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				close(waitC)
				break
			}
			if err != nil {
				log.Fatal(err)
				break
			}
			if in.Status {
				log.Print("Caught!")
				userLocalScore++
			}
		}
		log.Print("Listener Closed!")
		return
	}()
	// write Bidirectional
	for i := 0; i < 10000; i++ {
		stream.Send(&protoGo.RequestMessage{Username: "alameddin", X: uint64(i), Y: uint64(i)})
	}
	// closing...
	time.Sleep(5 * time.Second)
	stream.CloseSend()
	<-waitC
	// Simple RPC
	res, err := c.HighScore(context.Background(), &protoGo.RequestHighScore{Username: "alameddin"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
