package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/binsabit/grpc-course/protogen/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LaptopServer struct {
	Store LaptopStore
	pb.LaptopServiceServer
}
type LaptopService struct {
}

func NewLaptopServer(store LaptopStore) *LaptopServer {
	return &LaptopServer{Store: store}
}

func (s *LaptopServer) CreateLaptop(ctx context.Context, req *pb.CreateLaptopRequest) (*pb.CreateLaptopResponse, error) {
	laptop := req.GetLaptop()
	time.Sleep(6 * time.Second)
	log.Printf("recieved a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not valid UUID: %v", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}
	///save laptom in data base -- int real project
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("deadline is exceeding")
		return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeding")
	}
	err := s.Store.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrorAlreadyExist) {
			code = codes.AlreadyExists
		}
		return nil, status.Errorf(code, "cannot cannot store laptop with ID: %s, %v", laptop.Id, err)

	}

	res := &pb.CreateLaptopResponse{Id: laptop.Id}
	log.Printf("saved laptop with id:%s", laptop.Id)
	return res, nil
}
