package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type HelloWorldService struct {
	// define the interface and data type
}

func (rpc *HelloWorldService) SayHello(ctx context.Context, a *HelloRequest) (*HelloReply, error) {
	fmt.Println(a.GetName())
	return &HelloReply{Message: "hello~"}, nil
}

func (rpc *HelloWorldService) mustEmbedUnimplementedHelloWorldServiceServer() {

}

func main() {
	srv := grpc.NewServer()
	RegisterHelloWorldServiceServer(srv, &HelloWorldService{})
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
