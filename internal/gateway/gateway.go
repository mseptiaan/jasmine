package gateway

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/joho/godotenv/autoload"
	"github.com/mseptiaan/jasmine/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"os"
)

func Run(dialAddr string) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient(dialAddr, opts...)
	if err != nil {
		return fmt.Errorf("failed to dial server: %w", err)
	}

	gwmux := runtime.NewServeMux()

	if err = pb.RegisterJasmineEndpointHandler(context.Background(), gwmux, conn); err != nil {
		return fmt.Errorf("failed to register gateway: %w", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "11000"
	}
	gatewayAddr := "0.0.0.0:" + port
	gwServer := &http.Server{
		Addr: gatewayAddr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gwmux.ServeHTTP(w, r)
		}),
	}

	fmt.Printf(`

      _____   ______  ________  ______
  __ / / _ | / __/  |/  /  _/ |/ / __/ Jasmine v%s pId %d
 / // / __ |_\ \/ /|_/ // //    / _/   gRpc Address %s
 \___/_/ |_/___/_/  /_/___/_/|_/___/   restAPI Address %s

`, "1.0.0", os.Getpid(), dialAddr, gatewayAddr)

	return fmt.Errorf("serving gRPC-Gateway server: %w", gwServer.ListenAndServe())
}
