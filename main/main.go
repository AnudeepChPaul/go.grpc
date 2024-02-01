package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "go.grpc/s_hello"
	hello "go.grpc/s_hello/handlers"
	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "The server port")

func setupServer() *grpc.Server {
	server := grpc.NewServer()

	pb.RegisterGreeterServer(server, &hello.GreeterServerImpl{})
	pb.RegisterAuthenticatorServer(server, &hello.AuthenticatorServerImpl{})

	return server
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := setupServer()

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
