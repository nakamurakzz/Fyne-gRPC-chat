package main

import (
	"context"
	"grpc-chat-client/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	log.Println("start client...")
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure()) // 非TLS通信
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	chatClient := pb.NewChatServiceClient(conn)
	callChat(chatClient)
}

func callChat(client pb.ChatServiceClient) {
	chat := &pb.ChatMessageRequest{
		ChatMessage: "Hello from the client!",
	}
	response, err := client.Chat(context.Background(), chat)
	if err != nil {
		log.Fatalf("Error when calling Chat: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)
}
