package controllers

import (
	"authserver/apigateway/proto"
	auth "authserver/authentication"
	"context"
	"fmt"
)

type UserServiceServer struct {
	proto.UnimplementedUserServiceServer
}

func (s *UserServiceServer) SendUser(ctx context.Context, req *proto.UserRequest) (*proto.UserResponse, error) {
	fmt.Println("Received SendUser:", req.Email)
	return &proto.UserResponse{
		Message: "User received: " + req.Email,
	}, nil
}

func (s *UserServiceServer) SignupUser(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {
	fmt.Println("Received SignupUser:", req.Email)
	// Example logic
	if req.Email == "" || req.Password == "" {
		return &proto.SignupResponse{
			Error: "Email or password missing",
		}, nil
	}

	//Check Email
	if !auth.IsValidEmail(req.Email) {
		fmt.Println("error Invalid email format")
		return &proto.SignupResponse{
			Message: "Invalid Email" + req.Email,
		}, nil
	}

	//Check Password
	if len(req.Password) < 6 {
		fmt.Println("Password is to short")
		return &proto.SignupResponse{
			Message: "Password To Short" + req.Email,
		}, nil
	}
	fmt.Println("----------->")
	//Check if email already exist
	exists, err := auth.IsEmailExists(req.Email)
	if err != nil {
		return &proto.SignupResponse{
			Message: "Database error",
		}, nil
	}
	fmt.Println(exists)
	if exists {
		return &proto.SignupResponse{
			Message: "Already exists!" + req.Email,
		}, nil
	}
	fmt.Println("------->")
	return &proto.SignupResponse{
		Message: "User signed up: " + req.Email,
	}, nil
}
