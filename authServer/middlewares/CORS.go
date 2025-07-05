package middlewares

import (
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(UnaryLoggerInterceptor()),
	)

	// Register your services here...
	// pb.RegisterUserServiceServer(grpcServer, &UserServiceServer{})

	wrapped := grpcweb.WrapServer(grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return true // Allow all origins
		}),
	)

	httpServer := http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if wrapped.IsGrpcWebRequest(r) || wrapped.IsAcceptableGrpcCorsRequest(r) || wrapped.IsGrpcWebSocketRequest(r) {
				wrapped.ServeHTTP(w, r)
				return
			}
			http.NotFound(w, r)
		}),
	}

	log.Println("🌐 gRPC-Web server listening on :8080")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
