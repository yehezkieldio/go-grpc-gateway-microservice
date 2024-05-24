package main

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	proto "github.com/yehezkieldio/go-grpc-gateway-microservice/proto/health/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func newHealthGateway(ctx context.Context) (http.Handler, error)  {
	health := "0.0.0.0:8080"

	mux := runtime.NewServeMux()
	dialOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := proto.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, health, dialOpts)
	if err != nil {
		return nil, err
	}

	return mux, nil

}

func main() {
	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := http.NewServeMux()

	gw, err := newHealthGateway(ctx)
	if err != nil {
		slog.Error("failed to create gateway: %v", err)
	}

	mux.Handle("/", gw)

	slog.Info("gateway listening at 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		slog.Error("failed to serve: %v", err)
	}

	slog.Info("gateway started")


}