package services

import (
	"context"
	"github.com/danilobandeira29/grpc/pb"
	"time"
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

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserStream{
		Status: "Init",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)
	stream.Send(&pb.UserStream{
		Status: "User inserted",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)
	stream.Send(&pb.UserStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "1",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})
	return nil
}
