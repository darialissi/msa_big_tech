package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"               

	"msa_big_tech/auth/pkg/api/proto"                                              
)

func main() {
	conn, err := grpc.NewClient("localhost:8083",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := auth.NewAuthServiceClient(conn)

	{
		ctx := context.Background()

		resp, err := cli.Register(ctx, &auth.RegisterRequest{
			Email:   "client_email",
			Password: "client_password",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			log.Printf("register user id: %d\n", resp.UserId)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.Login(ctx, &auth.LoginRequest{
			Email:   "client_email",
			Password: "client_password",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			log.Printf("login user id: %d\n", resp.UserId)
		}
	}

	{
		ctx := context.Background()

		resp, err := cli.Refresh(ctx, &auth.RefreshRequest{
			RefreshToken: "client_refresh_token",
		})
		if err != nil {
			log.Fatalln(status.Code(err).String())
		} else {
			log.Printf("refresh token for user id: %d\n", resp.UserId)
		}
	}
}