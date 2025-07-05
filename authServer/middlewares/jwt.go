package middlewares

import (
	"authserver/utils"
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

/*
// Use a secure, secret key in production
var jwtSecret = []byte("your_super_secret_key")

// JWT Claims structure
type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// AuthInterceptor validates JWT token
func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		// Allow unauthenticated access to Signup and Login
		if info.FullMethod == "/proto.UserService/SignupUser" || info.FullMethod == "/proto.UserService/LoginUser" {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "missing metadata")
		}

		authHeaders := md["authorization"]
		if len(authHeaders) == 0 {
			return nil, status.Error(codes.Unauthenticated, "missing auth token")
		}

		token := strings.TrimPrefix(authHeaders[0], "Bearer ")
		claims, err := validateJWTToken(token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		log.Printf("✅ Authenticated user: %s (%s)", claims.UserID, claims.Email)
		return handler(ctx, req)
	}
}

// Validate and parse JWT token
func validateJWTToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, err
	}

	return claims, nil
} */

type ContextKey string

const (
	ContextEmailKey ContextKey = "email"
	ContextRoleKey  ContextKey = "role"
)

func AuthInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Exclude login/signup
		if info.FullMethod == "/proto.UserService/LoginUser" ||
			info.FullMethod == "/proto.UserService/SignupUser" {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unauthenticated, "Missing metadata")
		}

		authHeader := md.Get("authorization")
		if len(authHeader) == 0 {
			return nil, status.Errorf(codes.Unauthenticated, "Missing authorization header")
		}

		tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")
		claims, err := utils.VerifyAccessToken(tokenStr)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "Invalid token: %v", err)
		}

		// Add to context
		ctx = context.WithValue(ctx, ContextEmailKey, claims.Email)
		ctx = context.WithValue(ctx, ContextRoleKey, claims.Role)

		return handler(ctx, req)
	}
}

// Helper for checking role inside controller
func HasRole(ctx context.Context, expected string) bool {
	role, ok := ctx.Value(ContextRoleKey).(string)
	return ok && role == expected
}
