package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/darialissi/msa_big_tech/auth/pkg"
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
			Email:    "me@example.com",
			Password: "paSS123",
		})
		if err != nil {
			log.Printf("cli.Register: %s\n", status.Code(err).String())
		} else {
			log.Printf("cli.Register: %s\n", resp)
		}
	}
}
