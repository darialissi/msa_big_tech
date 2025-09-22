package main

import (
	"context"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/darialissi/msa_big_tech/auth/pkg"
	"github.com/darialissi/msa_big_tech/chat/pkg"
	"github.com/darialissi/msa_big_tech/social/pkg"
	"github.com/darialissi/msa_big_tech/users/pkg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// mux для REST
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}))

	// Настройка подключения к gRPC-серверу
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Регистрируем обработчики для gRPC-Gateway
	err := auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "auth_service:8083", opts)
	if err != nil {
		log.Fatalf("failed to connect to auth_service: %v", err)
	}

	err = chat.RegisterChatServiceHandlerFromEndpoint(ctx, mux, "chat_service:8084", opts)
	if err != nil {
		log.Fatalf("failed to connect to chat_service: %v", err)
	}

	err = social.RegisterSocialServiceHandlerFromEndpoint(ctx, mux, "social_service:8085", opts)
	if err != nil {
		log.Fatalf("failed to connect to social_service: %v", err)
	}

	err = users.RegisterUsersServiceHandlerFromEndpoint(ctx, mux, "users_service:8086", opts)
	if err != nil {
		log.Fatalf("failed to connect to users_service: %v", err)
	}

	// Запускаем HTTP-сервер
	log.Println("gRPC-Gateway server listening on :8087")
	if err := http.ListenAndServe(":8087", mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
