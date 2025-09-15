package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"  

	"msa_big_tech/auth/pkg/api/proto"
)


func main() {
	conn, err := grpc.NewClient("localhost:8082",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	cli := auth.NewAuthServiceClient(conn)

	// Grpc вызовы в другие сервисы
	{
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

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
}
