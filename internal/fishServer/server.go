package fishServer

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"server/pkg/protoGo"
)

type FishServer struct {
	protoGo.UnimplementedFishServiceServer
	UserScores map[string]uint
	Users map[protoGo.FishService_TryToCatchServer]bool
}

func NewFishServer(grpcServer *grpc.Server) FishServer {
	fishServer := FishServer{
		UserScores: map[string]uint{},
		Users: map[protoGo.FishService_TryToCatchServer]bool{},
	}
	protoGo.RegisterFishServiceServer(grpcServer, &fishServer)
	return fishServer
}

func (s *FishServer) TryToCatch(fStream protoGo.FishService_TryToCatchServer) error {
	for {
		in, err := fStream.Recv()
		if err == io.EOF{
			s.Users[fStream] = false
			log.Print("Disconnection")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := s.Users[fStream]; !ok{
			s.Users[fStream] = true
		}
		status := false
		if in.X == in.Y{
			if _, ok := s.UserScores[in.Username]; !ok {
				s.UserScores[in.Username] = 0
			}
			s.UserScores[in.Username]++
			status = true
		}
		for k, v:= range s.Users{
			if v {
				k.Send(&protoGo.ResponseMessage{Username: in.Username, Status: status})
			}
		}
	}
	return nil
}

func (s *FishServer) HighScore(ctx context.Context,request *protoGo.RequestHighScore) (*protoGo.ResponseHighScore, error) {
	log.Println("Adana 2")
	var highScore int64 = 0
	if v, ok := s.UserScores[request.Username]; ok{
		highScore = int64(v)
	}
	response := protoGo.ResponseHighScore{
		Users:    map[string]int64{
			request.Username : highScore,
		},
		YourRank: 1,
		Status:   true,
	}
	return &response, nil
}
