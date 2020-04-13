package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/robsantossilva/grpc-hello-go/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (*server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	result := "Hello " + request.GetName()

	res := &pb.HelloResponse{
		Msg: result,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8080")

	if err != nil {
		log.Fatalf("Failed to list %v", err)
	}

	fmt.Printf("Server is running in localhost:8080")

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &server{})
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
