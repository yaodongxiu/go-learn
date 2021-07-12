package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := NewHelloWorldServiceClient(conn)
	resp, err := client.SayHello(context.Background(), &HelloRequest{Name: "hello world!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(resp)
}
