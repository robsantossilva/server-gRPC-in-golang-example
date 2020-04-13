package main

import (
	"context"
	"log"

	"github.com/robsantossilva/grpc-hello-go/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	request := &pb.HelloRequest{
		Name: "Robson",
	}

	res, err := client.Hello(context.Background(), request)

	if err != nil {
		log.Fatalf("Error during the execution: %v", err)
	}

	log.Println(res.Msg)
}
