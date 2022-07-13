package main

import (
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
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
