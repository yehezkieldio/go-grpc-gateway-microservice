package main

import (
	"context"
	"log/slog"
	"net"

	proto "github.com/yehezkieldio/go-grpc-gateway-microservice/proto/health/v1"
	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedHealthServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) Check(context.Context, *proto.CheckRequest) (*proto.CheckResponse, error) {
	return &proto.CheckResponse{
		Status: proto.CheckResponse_SERVING_STATUS_SERVING,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		slog.Error("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterHealthServiceServer(s, NewServer())
	slog.Info("server listening at 8080")
	go func() {
		if err := s.Serve(lis); err != nil {
			slog.Error("failed to serve: %v", err)
		}
	}()

	slog.Info("server started")

	<-make(chan struct{})

}