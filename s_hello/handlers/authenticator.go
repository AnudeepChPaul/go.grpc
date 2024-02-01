package s_hello

import (
	"context"
	"github.com/google/uuid"
	pb "go.grpc/s_hello"
)

type AuthenticatorServerImpl struct {
	pb.UnimplementedAuthenticatorServer
}

func (s *AuthenticatorServerImpl) Auth(ctx context.Context, in *pb.TryAuthRequest) (*pb.Authenticated, error) {
	return &pb.Authenticated{Token: uuid.NewString(), Name: in.GetName()}, nil
}
