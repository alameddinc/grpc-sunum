package fishServer

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"io"
	"log"
	"server/pkg/protoGo"
	"sync"
	"time"
)

type FishServer struct {
	protoGo.UnimplementedFishServiceServer
	UserScores      map[string]uint
	Users           map[protoGo.FishService_TryToCatchServer]bool
	userLimit       int
	gameScoreConfig uint
	mutex           *sync.Mutex
	completedGame   chan bool
}

func NewFishServer(grpcServer *grpc.Server) FishServer {
	fishServer := FishServer{
		UserScores:      map[string]uint{},
		Users:           map[protoGo.FishService_TryToCatchServer]bool{},
		userLimit:       3,
		gameScoreConfig: 25,
		mutex:           new(sync.Mutex),
		completedGame:   make(chan bool),
	}
	protoGo.RegisterFishServiceServer(grpcServer, &fishServer)
	return fishServer
}
func (s *FishServer) Register(request *protoGo.RequestRegister, registerServer protoGo.FishService_RegisterServer) error {
	if _, ok := s.UserScores[request.Username]; ok {
		return errors.New("User is exist!")
	}
	s.mutex.Lock()
	s.UserScores[request.Username] = 0
	s.mutex.Unlock()
	for len(s.UserScores) < s.userLimit {
		registerServer.Send(&protoGo.ResponseRegister{Message: "Users Waiting...", Status: false})
		time.Sleep(time.Second)
	}
	registerServer.Send(&protoGo.ResponseRegister{Message: "Starting...", Status: true})
	return nil
}

func (s *FishServer) TryToCatch(fStream protoGo.FishService_TryToCatchServer) error {
	for {
		select {
		case <-s.completedGame:
			return nil
		default:
			in, err := fStream.Recv()
			if err == io.EOF {
				s.mutex.Lock()
				delete(s.Users, fStream)
				s.mutex.Unlock()
				log.Print("Disconnection")
				s.completedGame <- true
				break
			}
			if err != nil {
				return err
			}
			if _, ok := s.Users[fStream]; !ok {
				s.Users[fStream] = true
			}
			if in.X == in.Y {
				s.mutex.Lock()
				s.UserScores[in.Username]++
				s.mutex.Unlock()
				fStream.Send(&protoGo.ResponseMessage{Username: in.Username, Status: false})
				if s.UserScores[in.Username] >= s.gameScoreConfig {
					for k, _ := range s.Users {
						k.Send(&protoGo.ResponseMessage{Username: in.Username, Status: true})
					}
					break
				}
			}
		}
	}
	return nil
}

func (s *FishServer) HighScore(ctx context.Context, request *protoGo.RequestHighScore) (*protoGo.ResponseHighScore, error) {
	var highScore int64 = 0
	if v, ok := s.UserScores[request.Username]; ok {
		highScore = int64(v)
	}
	response := protoGo.ResponseHighScore{
		Users: map[string]int64{
			request.Username: highScore,
		},
		YourRank: 1,
		Status:   true,
	}
	s.mutex.Lock()
	delete(s.UserScores, request.Username)
	s.mutex.Unlock()
	return &response, nil
}
