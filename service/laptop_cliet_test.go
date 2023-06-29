package service_test

import (
	"context"
	"net"
	"testing"

	"github.com/binsabit/grpc-course/protogen/pb"
	"github.com/binsabit/grpc-course/samples"
	"github.com/binsabit/grpc-course/serializer"
	"github.com/binsabit/grpc-course/service"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

func TestCreateLaptop(t *testing.T) {
	t.Parallel()
	laptopServer, serverAddress := startTestLaptopServer(t)
	laptopClient := startTestLaptopClient(t, serverAddress)

	laptop := samples.NewLaptop()
	expectedID := laptop.Id
	req := &pb.CreateLaptopRequest{Laptop: laptop}

	res, err := laptopClient.CreateLaptop(context.Background(), req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, expectedID, res.Id)

	//chek the memeory
	other, err := laptopServer.Store.Find(res.Id)
	assert.Nil(t, err)
	assert.NotNil(t, other)

	//check if the same
	requireSameLaptop(t, laptop, other)
}

func startTestLaptopServer(t *testing.T) (*service.LaptopServer, string) {
	laptopStore := service.NewInMemoryLaptopStore()

	laptopServer := service.NewLaptopServer(laptopStore)

	grpcServer := grpc.NewServer()

	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	listener, err := net.Listen("tcp", ":0")
	assert.Nil(t, err)
	go grpcServer.Serve(listener)

	return laptopServer, listener.Addr().String()
}

func startTestLaptopClient(t *testing.T, serverAddr string) pb.LaptopServiceClient {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	assert.Nil(t, err)
	return pb.NewLaptopServiceClient(conn)
}

func requireSameLaptop(t *testing.T, laptop1 *pb.Laptop, laptop2 *pb.Laptop) {
	json1, err := serializer.ProtobufToJSON(laptop1)
	assert.NoError(t, err)

	json2, err := serializer.ProtobufToJSON(laptop2)
	assert.NoError(t, err)
	assert.Equal(t, json1, json2)
}
