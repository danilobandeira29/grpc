package main

import (
	"context"
	"fmt"
	"github.com/danilobandeira29/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()
	client := pb.NewUserServiceClient(connection)
	AddUserStreamBoth(client)
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

func AddUsers(client pb.UserServiceClient) {
	requests := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Danilo Bandeira",
			Email: "danilobandeira29@gmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Ana Banana",
			Email: "ana@banana.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Maria de Fatima",
			Email: "maria@fatima.com",
		},
	}
	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Could not send stream to AddUsers: %v", err)
	}
	for i, user := range requests {
		stream.Send(user)
		fmt.Println("Sending User", user.GetName())
		if i == len(requests)-1 {
			fmt.Println("End of Users")
		}
		time.Sleep(time.Second * 3)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Could not Close and Recv the response for AddUsers: %v", err)
	}
	log.Println("Result", res)
}

func AddUserStreamBoth(client pb.UserServiceClient) {
	requestStream, err := client.AddUserStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Could not make gRPC call AddUserStreamBoth: %v", err)
	}
	requests := []*pb.User{
		&pb.User{
			Id:    "1",
			Name:  "Danilo Bandeira",
			Email: "danilobandeira29@gmail.com",
		},
		&pb.User{
			Id:    "2",
			Name:  "Ana Banana",
			Email: "ana@banana.com",
		},
		&pb.User{
			Id:    "3",
			Name:  "Maria de Fatima",
			Email: "maria@fatima.com",
		},
	}
	wait := make(chan int)
	go func() {
		for _, user := range requests {
			err := requestStream.Send(user)
			if err != nil {
				log.Fatalf("Could not Send gRPC call AddUserStreamBoth: %v", err)
			}
			fmt.Println("Sending user", user.GetName())
			time.Sleep(time.Second * 3)
		}
		requestStream.CloseSend()
	}()
	go func() {
		for {
			response, err := requestStream.Recv()
			if err == io.EOF {
				fmt.Println("EOF")
				break
			}
			if err != nil {
				log.Fatalf("Could not receive stream AddUserStreamBoth: %v", err)
			}
			fmt.Printf("Receiving User: %v with Status: %v\n", response.GetUser().GetName(), response.GetStatus())
		}
		close(wait)
	}()
	<-wait
}
