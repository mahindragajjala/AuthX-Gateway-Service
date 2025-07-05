package main

import (
	"log"
	"net"

	"authserver/apigateway/proto"
	"authserver/controllers"
	"authserver/db/mongodb"
	"authserver/db/postgres"
	"authserver/middlewares"

	"google.golang.org/grpc"
)

func init() {
	// Connect to DB
	postgres.ConnectPostgres()
	mongodb.InitMongo()
	postgres.Delete_Table_Data()
}

func main() {
	// Start TCP listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server with middleware chain (auth + logger)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			middlewares.ChainUnaryInterceptors(
				middlewares.UnaryLoggerInterceptor(), // Logs every call
				middlewares.AuthInterceptor(),        // Verifies JWT, injects claims
			),
		),
		grpc.StreamInterceptor(
			middlewares.StreamLoggerInterceptor(),
		),
	)

	// Register UserService
	proto.RegisterUserServiceServer(grpcServer, &controllers.UserServiceServer{})

	log.Println("🚀 gRPC server running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("❌ Failed to serve: %v", err)
	}
}
