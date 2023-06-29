package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/binsabit/grpc-course/protogen/pb"
	"github.com/binsabit/grpc-course/samples"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()

	log.Printf("dial server at address %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("could not establish connections:%v", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)

	laptop := samples.NewLaptop()
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := laptopClient.CreateLaptop(ctx, req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Println("laptop Alredy exists")
		} else {
			log.Fatal("could not creat laptop")
		}
		return
	}
	log.Printf("laptop is create with id: %s", res.Id)
}
