package main

import (
	"context"
	"grpc-chat-server/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
}

func (c *ChatServer) Chat(context.Context, *pb.ChatMessageRequest) (*pb.ChatMessageResponse, error) {
	log.Println("Chat called")
	chat := &pb.ChatMessageResponse{
		Message: "Hello from the server!",
	}
	return chat, nil
}

func main() {
	log.Println("start server...")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()

	pb.RegisterChatServiceServer(server, &ChatServer{})

	log.Println("Server is running at port 8080")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
