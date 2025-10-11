package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"

	social "github.com/darialissi/msa_big_tech/social/pkg"
)

func main() {
	conn, err := grpc.NewClient("localhost:8085",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := social.NewSocialServiceClient(conn)

	userId := uuid.New().String()
	frReq := ""

	{
		ctx := context.Background()

		resp, err := cli.SendFriendRequest(ctx, &social.SendFriendRequestRequest{
			UserId: userId,
		})
		if err != nil {
			log.Printf("cli.SendFriendRequest: %s\n", status.Code(err).String())
		} else {
			frReq = resp.FriendRequest.RequestId
			log.Printf("cli.SendFriendRequest: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.ListFriendRequests(ctx, &social.ListFriendRequestsRequest{
			UserId: userId,
		})
		if err != nil {
			log.Printf("cli.ListFriendRequests: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.ListFriendRequests: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.AcceptFriendRequest(ctx, &social.AcceptFriendRequestRequest{
			FriendRequestId: frReq,
		})
		if err != nil {
			log.Printf("cli.AcceptFriendRequest: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.AcceptFriendRequest: %s\n", resp.FriendRequest)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.DeclineFriendRequest(ctx, &social.DeclineFriendRequestRequest{
			FriendRequestId: "00000000-0000-0000-0000-0000000000000",
		})
		if err != nil {
			log.Printf("cli.DeclineFriendRequest: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.DeclineFriendRequest: %s\n", resp.FriendRequest)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.ListFriends(ctx, &social.ListFriendsRequest{
			UserId: userId,
			Limit: 5,
		})
		if err != nil {
			log.Printf("cli.ListFriends: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.ListFriends: %s\n", resp)
		}
	}
}