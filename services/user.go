package services

import (
	"context"
	"github.com/danilobandeira29/grpc/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer // evitar que o server der erro por implementar um service que não existe no protofile
}

// NewUserService uma especie de construtor. assim posso chamar por services.NewUserService() no server
func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	return &pb.User{
		Id:    req.GetId(),
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}