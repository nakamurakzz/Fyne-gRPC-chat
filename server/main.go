package server

import (
	"fyne-grpc-chat/proto/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ChatServer struct {
	pb.UnimplementedChatServiceServer
}

func main() {
	log.Println("start server...")

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	server.Serve(lis)
}
