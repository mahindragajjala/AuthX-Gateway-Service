package main

import (
	"log"
	"net"

	"authserver/apigateway/proto"
	"authserver/controllers"
	"authserver/utils"

	"google.golang.org/grpc"
)

/* func main() {
	// Connect to PostgreSQL
	if err := utils.DBConnection(); err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}

	//Start listening
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	//Create gRPC server with logging middleware
	grpcServer := grpc.NewServer(utils.LoggingMiddleware())

	//Register service
	proto.RegisterUserServiceServer(grpcServer, &controllers.UserServiceServer{})

	fmt.Println("🚀 gRPC Server running on port 50051...")

	//Serve gRPC
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
*/
/*
func main() {
	createdb.CreateDatabase_Manual()
}
*/

func main() {
	// Initialize DB
	if err := utils.DBConnection(); err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	// Listen
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("❌ Failed to listen: %v", err)
	}

	// Create gRPC server with logging middleware
	//grpcServer := grpc.NewServer(utils.LoggingMiddleware())
	grpcServer := grpc.NewServer()

	// Register service
	proto.RegisterUserServiceServer(grpcServer, &controllers.UserServiceServer{})

	log.Println("🚀 gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
