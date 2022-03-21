package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/radmirid/grpc/proto/notification"
	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Notify(ctx context.Context, n *notification.NotificationRequest) (*notification.NotificationResponse, error) {
	fmt.Println(n.Message)
	return &notification.NotificationResponse{Status: "OK"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("error: Failed to Listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	notification.RegisterNotificationServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error: Failed to Serve: %s", err)
	}
}
