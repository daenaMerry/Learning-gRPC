package main

import (
	"fmt"
	"log"
	"net"

	"github.com/daenaMerry/Learning-gRPC/chat"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("gRPC Server started...")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service := &chat.Server{}
	chat.RegisterChatServiceServer(grpcServer, service)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
