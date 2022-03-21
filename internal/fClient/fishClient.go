package client

import (
	"context"
	"crypto/rand"
	"io"
	"log"
	"math/big"
	"time"

	"github.com/alameddinc/grpc-sunum/pkg/protoGo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Fishclient struct {
	username string
	score    uint
	conn     *grpc.ClientConn
	client   protoGo.FishServiceClient
}

func NewFishClient(username string) Fishclient {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	c := protoGo.NewFishServiceClient(conn)
	return Fishclient{
		username: username,
		score:    0,
		conn:     conn,
		client:   c,
	}
}

func (c *Fishclient) Close() error {
	return c.conn.Close()
}

func (c *Fishclient) Register() error {
	registerStream, err := c.client.Register(context.Background(), &protoGo.RequestRegister{Username: c.username})
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

func (c *Fishclient) ListenAndCatch() error {
	ctx := context.Background()
	stream, err := c.client.TryToCatch(ctx)
	if err != nil {
		return err
	}
	errChan := make(chan error)
	// Read Bidirectional
	go c.listen(stream, errChan)
	c.catch(stream, errChan)
	stream.CloseSend()
	for len(errChan) > 0 {
		<-errChan
	}
	return nil
}

func (c *Fishclient) listen(stream protoGo.FishService_TryToCatchClient, errChan chan error) {
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
			if in.Username == c.username || in.FishCount > 1 {
				log.Printf("%s +%d ", in.Username, in.FishCount)
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

func (c *Fishclient) catch(stream protoGo.FishService_TryToCatchClient, errChan chan error) {
	for {
		select {
		case <-errChan:
			stream.CloseSend()
			return
		default:
			time.Sleep(100 * time.Millisecond)
			nX, err := rand.Int(rand.Reader, big.NewInt(1000))
			nY, err := rand.Int(rand.Reader, big.NewInt(1000))
			if err != nil {
				panic(err)
			}
			stream.Send(&protoGo.RequestMessage{Username: c.username, X: nX.Uint64(), Y: nY.Uint64()})
		}
	}
}

func (c *Fishclient) Highscore() {
	res, err := c.client.HighScore(context.Background(), &protoGo.RequestHighScore{Username: c.username})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Skorunuz: %d\n", res.Users[c.username])
}
