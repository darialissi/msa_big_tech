package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	users "github.com/darialissi/msa_big_tech/users/pkg"
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

	userId := ""

	{
		ctx := context.Background()

		resp, err := cli.CreateProfile(ctx, &users.CreateProfileRequest{
			Nickname:  "test_me",
			Bio:       "info",
			AvatarUrl: "http://avatar",
		})
		if err != nil {
			log.Printf("cli.CreateProfile: %s\n", status.Code(err).String())
		} else {
			userId = resp.UserProfile.Id
			log.Printf("cli.CreateProfile: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.UpdateProfile(ctx, &users.UpdateProfileRequest{
			UserId:    userId,
			Nickname:  "test_me_0",
			Bio:       "",
			AvatarUrl: "http://avatar.me",
		})
		if err != nil {
			log.Printf("cli.UpdateProfile: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.UpdateProfile: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.GetProfileByID(ctx, &users.GetProfileByIDRequest{
			UserId: userId,
		})
		if err != nil {
			log.Printf("cli.GetProfileByID: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.GetProfileByID: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.GetProfileByNickname(ctx, &users.GetProfileByNicknameRequest{
			Nickname: "test_me_0",
		})
		if err != nil {
			log.Printf("cli.GetProfileByNickname: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.GetProfileByNickname: %s\n", resp)
		}
	}
}
