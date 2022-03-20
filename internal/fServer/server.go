package fServer

import (
	"context"
	"crypto/rand"
	"errors"
	"io"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/alameddinc/grpc-sunum/pkg/protoGo"
	"google.golang.org/grpc"
)

const (
	yLimit = 1000
	xLimit = 1000
)

type FishServer struct {
	protoGo.UnimplementedFishServiceServer
	userScores      map[string]uint
	gameScoreConfig uint
	mutex           *sync.Mutex
	mutexUser       *sync.Mutex
	users           map[protoGo.FishService_TryToCatchServer]bool
	gameMap         [1000][1000]uint
	userLimit       int
	completedGame   chan bool
	x               int
	y               int
}

func NewFishServer(grpcServer *grpc.Server) FishServer {
	fishServer := FishServer{
		userScores:      map[string]uint{},
		mutex:           new(sync.Mutex),
		mutexUser:       new(sync.Mutex),
		users:           map[protoGo.FishService_TryToCatchServer]bool{},
		userLimit:       4,
		gameScoreConfig: 100,
		gameMap:         [1000][1000]uint{},
		x:               1000,
		y:               1000,
		completedGame:   make(chan bool),
	}
	fishServer.mapGenerator()
	protoGo.RegisterFishServiceServer(grpcServer, &fishServer)
	return fishServer
}
func (s *FishServer) Register(request *protoGo.RequestRegister, registerServer protoGo.FishService_RegisterServer) error {
	if _, ok := s.userScores[request.Username]; ok {
		return errors.New("User is exist!")
	}
	s.mutex.Lock()
	s.userScores[request.Username] = 0
	s.mutex.Unlock()
	for len(s.userScores) < s.userLimit {
		registerServer.Send(&protoGo.ResponseRegister{Message: "users Waiting...", Status: false})
		time.Sleep(time.Millisecond * 200)
	}
	registerServer.Send(&protoGo.ResponseRegister{Message: "Starting...", Status: true})
	return nil
}

func (s *FishServer) mapGenerator() {
	// small fish groups
	for i := 0; i < xLimit; i++ {
		for j := 0; j < yLimit; j++ {
			isFishExist, err := rand.Int(rand.Reader, big.NewInt(5))
			if err != nil {
				continue
			}
			if isFishExist.Int64() == 0 {
				s.gameMap[i][j] = 1
			}
		}
	}
	// big fish groups
	count := 200
	for count > 0 {
		randx, err := rand.Int(rand.Reader, big.NewInt(xLimit-6))
		randy, err := rand.Int(rand.Reader, big.NewInt(yLimit-6))
		if err != nil {
			break
		}
		x := randx.Int64() + 3
		y := randy.Int64() + 3
		s.mutex.Lock()
		if s.gameMap[x][y] > 1 {
			continue
		}
		// Big Fish Group Mask
		//  x   x  15  x   x
		//  x  15  30  15  x
		// 15  30  50  30 15
		//  x  15  30  15  x
		//  x   x  15  x   x
		s.gameMap[x][y] = 50

		s.gameMap[x-1][y] = 30
		s.gameMap[x-2][y] = 15
		s.gameMap[x-1][y-1] = 15
		s.gameMap[x-1][y+1] = 15

		s.gameMap[x+1][y] = 30
		s.gameMap[x+2][y] = 15
		s.gameMap[x+1][y-1] = 15
		s.gameMap[x+1][y+1] = 15

		s.gameMap[x][y-1] = 30
		s.gameMap[x][y-2] = 15
		s.gameMap[x+1][y-1] = 15
		s.gameMap[x-1][y-1] = 15

		s.gameMap[x][y+1] = 30
		s.gameMap[x][y+2] = 15
		s.gameMap[x-1][y+1] = 15
		s.gameMap[x+1][y+1] = 15
		count--
		s.mutex.Unlock()
	}
}

func (s *FishServer) TryToCatch(fStream protoGo.FishService_TryToCatchServer) error {
	for {
		in, err := fStream.Recv()
		if err == io.EOF {
			s.mutexUser.Lock()
			delete(s.users, fStream)
			s.mutexUser.Unlock()
			log.Print("Oyuncu Ayrıldı.")
			break
		}
		if err != nil {
			return err
		}
		if _, ok := s.users[fStream]; !ok {
			s.mutexUser.Lock()
			s.users[fStream] = true
			s.mutexUser.Unlock()
		}
		if s.gameMap[in.X][in.Y] != 0 {
			s.mutex.Lock()
			coordinatePoint := s.gameMap[in.X][in.Y]
			s.gameMap[in.X][in.Y] = 0
			s.userScores[in.Username] += coordinatePoint
			s.mutex.Unlock()
			fStream.Send(&protoGo.ResponseMessage{Username: in.Username, FishCount: uint64(coordinatePoint), Status: false})
			if s.userScores[in.Username] >= s.gameScoreConfig {
				s.sendMessageAllUser(in, coordinatePoint)
				break
			}
		}
	}
	return nil
}

func (s *FishServer) HighScore(ctx context.Context, request *protoGo.RequestHighScore) (*protoGo.ResponseHighScore, error) {
	var highScore int64 = 0
	for len(s.completedGame) > 0 {
		<-s.completedGame
	}
	if v, ok := s.userScores[request.Username]; ok {
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
	delete(s.userScores, request.Username)
	s.mutex.Unlock()
	return &response, nil
}

func (s *FishServer) sendMessageAllUser(in *protoGo.RequestMessage, point uint) {
	s.mutexUser.Lock()
	for k, _ := range s.users {
		k.Send(&protoGo.ResponseMessage{Username: in.Username, FishCount: uint64(point), Status: true})
	}
	s.mutexUser.Unlock()
}
