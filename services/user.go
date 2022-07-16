package services

import (
	"context"
	"github.com/danilobandeira29/grpc/pb"
	"io"
	"log"
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
	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)
	stream.Send(&pb.UserResultStream{
		Status: "User inserted",
		User:   &pb.User{},
	})
	time.Sleep(time.Second * 3)
	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id:    "1",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})
	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	var users []*pb.User
	for {
		requestStream, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				Users: users,
			})
		}
		if err != nil {
			log.Fatalf("Could not recive the stream AddUsers: %v", err)
		}
		users = append(users, &pb.User{
			Id:    requestStream.GetId(),
			Name:  requestStream.GetName(),
			Email: requestStream.GetEmail(),
		})
	}
}
