package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

// DB connection
func DBConnection() error {
	var err error
	dsn := "host=localhost port=5432 user=mahindra password=1234 dbname=auth_db sslmode=disable"
	Db, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open DB: %v", err)
	}
	if err = Db.Ping(); err != nil {
		return fmt.Errorf("failed to ping DB: %v", err)
	}
	log.Println("✅ PostgreSQL connected successfully!")
	return nil
}

/*
// gRPC Unary Logging Interceptor
func GRPCLoggerInterceptor(
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	log.Printf("🔔 Incoming gRPC call: %s", info.FullMethod)
	resp, err := handler(req)
	if err != nil {
		log.Printf("❌ Error handling gRPC call %s: %v", info.FullMethod, err)
	}
	return resp, err
}

// Return gRPC middleware
func LoggingMiddleware() grpc.ServerOption {
	return grpc.UnaryInterceptor(GRPCLoggerInterceptor)
}
*/
