package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"

	chat "github.com/darialissi/msa_big_tech/chat/pkg"
)

func main() {
	conn, err := grpc.NewClient("localhost:8084",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := chat.NewChatServiceClient(conn)

	chatId := ""
	userId := uuid.New().String()

	{
		ctx := context.Background()

		resp, err := cli.CreateDirectChat(ctx, &chat.CreateDirectChatRequest{
			ParticipantId: userId,
		})
		if err != nil {
			log.Printf("cli.CreateDirectChat: %s\n", status.Code(err).String())
		} else {
			chatId = resp.ChatId
			log.Printf("cli.CreateDirectChat: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.GetChat(ctx, &chat.GetChatRequest{
			ChatId: chatId,
		})
		if err != nil {
			log.Printf("cli.GetChat: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.GetChat: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.ListUserChats(ctx, &chat.ListUserChatsRequest{
			UserId: userId,
		})
		if err != nil {
			log.Printf("cli.ListUserChats: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.ListUserChats: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.ListChatMembers(ctx, &chat.ListChatMembersRequest{
			ChatId: chatId,
		})
		if err != nil {
			log.Printf("cli.ListChatMembers: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.ListChatMembers: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.SendMessage(ctx, &chat.SendMessageRequest{
			ChatId: chatId,
			Text:   "hi from script",
		})
		if err != nil {
			log.Printf("cli.SendMessage: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.SendMessage: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.ListMessages(ctx, &chat.ListMessagesRequest{
			ChatId: chatId,
			Limit:  5,
		})
		if err != nil {
			log.Printf("cli.ListMessages: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.ListMessages: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.StreamMessages(ctx, &chat.StreamMessagesRequest{
			ChatId:      chatId,
			SinceUnixMs: 1760210300,
		})
		if err != nil {
			log.Printf("cli.StreamMessages: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.StreamMessages: %s\n", resp)
		}
	}
}
