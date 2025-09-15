package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"msa_big_tech/auth/pkg/api/proto"
)

func main() {
	ctx := context.Background()

	// mux для REST
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}))

	// Настройка подключения к gRPC-серверу
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Регистрируем обработчики для gRPC-Gateway
	err := auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "localhost:8082", opts)
	if err != nil {
		log.Fatalf("failed to start gateway: %v", err)
	}

	// Запускаем HTTP-сервер
	log.Println("gRPC-Gateway server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
