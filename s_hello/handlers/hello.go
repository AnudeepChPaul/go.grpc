package s_hello

import (
	"context"
	pb "go.grpc/s_hello"
)

type GreeterServerImpl struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServerImpl) SayHelloAgain(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello again " + in.GetName()}, nil
}
