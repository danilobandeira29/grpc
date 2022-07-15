package main

import (
	"context"
	"fmt"
	"github.com/danilobandeira29/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()
	client := pb.NewUserServiceClient(connection)
	AddUserVerbose(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "1234",
		Name:  "Danilo Bandeira",
		Email: "danilobandeira29@gmail.com",
	}
	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC call AddUser: %v", err)
	}
	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	resStream, err := client.AddUserVerbose(context.Background(), &pb.User{
		Id:    "1234",
		Name:  "Danilo Bandeira",
		Email: "danilobandeira29@gmail.com",
	})
	if err != nil {
		log.Fatalf("Could not make gRPC call AddUserStream: %v", err)
	}
	for {
		message, err := resStream.Recv()
		if err == io.EOF {
			fmt.Println("End!")
			break
		}
		if err != nil {
			log.Fatalf("Could not recive the stream: %v", err)
		}
		fmt.Printf("Status: %v\nUser: %v\n\n", message.Status, message.User)
	}
}
