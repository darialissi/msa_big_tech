package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/darialissi/msa_big_tech/users/pkg"
)

func main() {
	conn, err := grpc.NewClient("localhost:8086",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := users.NewUsersServiceClient(conn)

	{
		ctx := context.Background()

		resp, err := cli.CreateProfile(ctx, &users.CreateProfileRequest{
			Nickname: "???",
			Bio: "I'm client",
			AvatarUrl: "http://avatar",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			profile, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("user profile: %s", string(profile))
			}
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.UpdateProfile(ctx, &users.UpdateProfileRequest{
			Nickname: "???",
			Bio: "I'm client",
			AvatarUrl: "http://avatar",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			profile, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("user profile: %s", string(profile))
			}
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.GetProfileByID(ctx, &users.GetProfileByIDRequest{
			UserId: "00000000-0000-0000-0000-0000000000000",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			profile, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("user profile: %s", string(profile))
			}
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.GetProfileByNickname(ctx, &users.GetProfileByNicknameRequest{
			Nickname: "new_client_nickname",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			profile, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("user profile: %s", string(profile))
			}
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.SearchByQuery(ctx, &users.SearchByQueryRequest{
			Query: "nickname=='client_nickname'",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			profiles, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("user profile: %s", string(profiles))
			}
		}
	}
}