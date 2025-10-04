package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"    
	"google.golang.org/protobuf/encoding/protojson"           

	"github.com/darialissi/msa_big_tech/chat/pkg"                                              
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

	{
		ctx := context.Background()

		resp, err := cli.CreateDirectChat(ctx, &chat.CreateDirectChatRequest{
			ParticipantId: "00000000-0000-0000-0000-0000000000000",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			log.Printf("chat id: %d\n", resp.ChatId)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.GetChat(ctx, &chat.GetChatRequest{
			ChatId: "00000000-0000-0000-0000-0000000000000",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			chat, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("chat: %s", string(chat))
			}
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.ListUserChats(ctx, &chat.ListUserChatsRequest{
			UserId: "00000000-0000-0000-0000-0000000000000",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			chats, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("chats: %s", string(chats))
			}
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.ListChatMembers(ctx, &chat.ListChatMembersRequest{
			ChatId: "00000000-0000-0000-0000-0000000000000",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			members, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("members ids: %s", string(members))
			}
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.SendMessage(ctx, &chat.SendMessageRequest{
			ChatId: "00000000-0000-0000-0000-0000000000000",
			Text: "hi from client",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			message, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("message: %s", string(message))
			}
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.ListMessages(ctx, &chat.ListMessagesRequest{
			ChatId: "00000000-0000-0000-0000-0000000000000",
			Limit: 5,
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			messages, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("messages: %s", string(messages))
			}
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.StreamMessages(ctx, &chat.StreamMessagesRequest{
			ChatId: "00000000-0000-0000-0000-0000000000000",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			resp, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("resp stream: %s", string(resp))
			}
		}
	}
}