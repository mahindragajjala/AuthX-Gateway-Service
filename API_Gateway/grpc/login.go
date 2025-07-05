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

// StartGRPCClient handles the SignupRequest and calls the SignupUser gRPC service
func StartGRPCClient_Login(req models.LoginRequest) (*proto.LoginResponse, error) {
	conn, err := grpc.Dial("172.20.78.91:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect to gRPC server: %v", err)
		return nil, err
	}
	defer conn.Close()

	client := proto.NewUserServiceClient(conn)
	return CallLoginUser(client, req)
}

// CallSignupUser invokes the SignupUser gRPC method
func CallLoginUser(client proto.UserServiceClient, userdata models.LoginRequest) (*proto.LoginResponse, error) {
	fmt.Println("Calling SignupUser...")

	req := &proto.LoginRequest{
		Email:    userdata.Email,
		Password: userdata.Password,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := client.LoginUser(ctx, req)
	if err != nil {
		log.Printf("SignupUser error: %v", err)
		return nil, err
	}

	fmt.Println("Response from gRPC:", res)
	return res, nil
}
