package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"  

	"msa_big_tech/auth/pkg"
	"msa_big_tech/chat/pkg"
)


func main() {
	conn, err := grpc.NewClient("localhost:8083",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli_auth := auth.NewAuthServiceClient(conn)

	// Grpc вызовы в auth
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := cli_auth.Register(ctx, &auth.RegisterRequest{
			Email:   "client_email",
			Password: "client_password",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			log.Printf("register user id: %d\n", resp.UserId)
		}
	}

	cli_chat := chat.NewChatServiceClient(conn)

	// Grpc вызовы в chat
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		resp, err := cli_chat.CreateDirectChat(ctx, &chat.CreateDirectChatRequest{
			ParticipantId:   0,
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			log.Printf("chat id: %d\n", resp.ChatId)
		}
	}
}
