package middlewares

import (
	"context"

	"google.golang.org/grpc"
)

// Chains multiple Unary interceptors
func ChainUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Wrap the final handler with all interceptors
		chain := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			curr := interceptors[i]
			chain = wrapUnaryHandler(curr, info, chain)
		}
		return chain(ctx, req)
	}
}

func wrapUnaryHandler(
	interceptor grpc.UnaryServerInterceptor,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) grpc.UnaryHandler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return interceptor(ctx, req, info, handler)
	}
}
