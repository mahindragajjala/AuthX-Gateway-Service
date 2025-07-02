package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"apigateway/apigateway/proto"
	"apigateway/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func StartGRPCClient(req models.SignupRequest) {
	conn, err := grpc.Dial("172.20.78.91:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)

	CallSignupUser(client, req)
}

func CallSendUser(client proto.UserServiceClient) {
	fmt.Println("Calling SendUser...")

	req := &proto.UserRequest{
		Email:    "mahindra@example.com",
		Password: "golang123",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.SendUser(ctx, req)
	if err != nil {
		log.Fatalf("SendUser error: %v", err)
	}

	fmt.Println("Response:", res.GetMessage())
}

func CallSignupUser(client proto.UserServiceClient, userdata models.SignupRequest) {
	fmt.Println("Calling SignupUser...")

	req := &proto.SignupRequest{
		Email:    userdata.Email,
		Password: userdata.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.SignupUser(ctx, req)
	if err != nil {
		log.Fatalf("SignupUser error: %v", err)
	}

	if res.GetError() != "" {
		fmt.Println("❌ Error:", res.GetError())
	} else {
		fmt.Println("✅ Success:", res.GetMessage())
	}
}
