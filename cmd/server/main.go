package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/binsabit/grpc-course/protogen/pb"
	"github.com/binsabit/grpc-course/service"
	"google.golang.org/grpc"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()

	log.Println(*port)
	laptopStore := service.NewInMemoryLaptopStore()
	laptopServer := service.NewLaptopServer(laptopStore)

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	address := fmt.Sprintf("0.0.0.0:%d", *port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("could not listent to address: %s, \n error:%v", address, err)
	}
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("could not serve at address:%v", err)

	}

}
