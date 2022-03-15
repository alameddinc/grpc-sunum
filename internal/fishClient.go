package internal

import (
	"context"
	"crypto/rand"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/big"
	"server/pkg/protoGo"
	"time"
)

type FishClient struct {
	Username string
	Score    uint
	conn     *grpc.ClientConn
	Client   protoGo.FishServiceClient
}

func NewFishClient(username string) FishClient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := protoGo.NewFishServiceClient(conn)
	return FishClient{
		Username: username,
		Score:    0,
		conn:     conn,
		Client:   c,
	}
}

func (c *FishClient) Close() error {
	return c.conn.Close()
}

func (c *FishClient) Register() error {
	registerStream, err := c.Client.Register(context.Background(), &protoGo.RequestRegister{Username: c.Username})
	defer registerStream.CloseSend()
	if err != nil {
		return err
	}
	waitRegister := make(chan error)
	go func() {
		for {
			recv, err := registerStream.Recv()
			if err == io.EOF {
				close(waitRegister)
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Println(recv.Message)
		}
	}()
	registerStream.CloseSend()
	if err := <-waitRegister; err != nil {
		return err
	}
	return nil
}

func (c *FishClient) ListenAndCatch() error {
	ctx := context.Background()
	stream, err := c.Client.TryToCatch(ctx)
	if err != nil {
		return err
	}
	errChan := make(chan error)
	// Read Bidirectional
	go c.listen(stream, errChan)
	c.catch(stream)
	stream.CloseSend()
	return <-errChan
}

func (c *FishClient) listen(stream protoGo.FishService_TryToCatchClient, errChan chan error) {
	for {
		select {
		case <-stream.Context().Done():
			errChan <- nil
			return
		default:
			in, err := stream.Recv()
			if err == io.EOF {
				errChan <- nil
				return
			}
			if err != nil {
				errChan <- err
				return
			}
			if in.Username == c.Username {
				c.Score++
				log.Printf("%s %d adet balık yakaladı!", in.Username, c.Score)
			}
			if in.Status {
				log.Printf("%s Kazandı!", in.Username)
				stream.Context().Done()
				errChan <- err
				return
			}
		}
	}
	errChan <- nil
	return
}

func (c *FishClient) catch(stream protoGo.FishService_TryToCatchClient) {
	for {
		select {
		case <-stream.Context().Done():
			log.Println("Adana")
			return
		default:
			log.Println(1)
			time.Sleep(25 * time.Millisecond)
			nX, err := rand.Int(rand.Reader, big.NewInt(10))
			nY, err := rand.Int(rand.Reader, big.NewInt(10))
			if err != nil {
				panic(err)
			}
			stream.Send(&protoGo.RequestMessage{Username: c.Username, X: nX.Uint64(), Y: nY.Uint64()})
		}
	}
}

func (c *FishClient) HighScore() {
	res, err := c.Client.HighScore(context.Background(), &protoGo.RequestHighScore{Username: c.Username})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Users)
}