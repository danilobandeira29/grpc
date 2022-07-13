package main

import (
	"github.com/danilobandeira29/grpc/pb"
	"github.com/danilobandeira29/grpc/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051") //ficará escutando a essa porta, golang já tem imbutido no pacote net
	if err != nil {
		log.Fatalf("Could not listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService()) // registrando o service AddUser
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
