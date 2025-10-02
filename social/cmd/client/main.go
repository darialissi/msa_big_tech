package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"    
	"google.golang.org/protobuf/encoding/protojson"           

	"github.com/darialissi/msa_big_tech/social/pkg"                                              
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

	{
		ctx := context.Background()

		resp, err := cli.SendFriendRequest(ctx, &social.SendFriendRequestRequest{
			UserId:   0,
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			log.Printf("request id: %d\n", resp.FriendRequest.RequestId)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.ListFriendRequests(ctx, &social.ListFriendRequestsRequest{
			UserId:   0,
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			requests, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("requests: %s", string(requests))
			}
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.AcceptFriendRequest(ctx, &social.AcceptFriendRequestRequest{
			FriendRequestId: 10,
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			log.Printf("request id: %d; request status: %s\n", resp.FriendRequest.RequestId, resp.FriendRequest.Status)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.DeclineFriendRequest(ctx, &social.DeclineFriendRequestRequest{
			FriendRequestId: 10,
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			log.Printf("request id: %d; request status: %s\n", resp.FriendRequest.RequestId, resp.FriendRequest.Status)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.RemoveFriend(ctx, &social.RemoveFriendRequest{
			UserId: 0,
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			log.Printf("response: %s\n", resp)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.ListFriends(ctx, &social.ListFriendsRequest{
			UserId:   0,
			Limit: 5,
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			// для Marshal proto сообщений в JSON необходимо использовать пакет protojson
			ids, err := protojson.Marshal(resp)
			if err != nil {
				log.Fatalf("protojson.Marshal error: %v", err)
			} else {
				log.Printf("friends user ids: %s", string(ids))
			}
		}
	}
}