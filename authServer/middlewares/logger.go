package middlewares

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func UnaryLoggerInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		start := time.Now()
		resp, err = handler(ctx, req)
		duration := time.Since(start)

		log.Printf("📍 gRPC Method: %s | Duration: %v | Error: %v", info.FullMethod, duration, err)
		return resp, err
	}
}

func StreamLoggerInterceptor() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		start := time.Now()
		err := handler(srv, ss)
		duration := time.Since(start)

		log.Printf("📡 gRPC Stream Method: %s | Duration: %v | Error: %v", info.FullMethod, duration, err)
		return err
	}
}
