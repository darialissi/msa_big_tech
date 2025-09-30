package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	grpc_hd "github.com/darialissi/msa_big_tech/chat/internal/app/controllers/grpc"
	chat "github.com/darialissi/msa_big_tech/chat/pkg"
)


func main() {
	implementation := grpc_hd.NewServer() // наша реализация сервера

	lis, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	chat.RegisterChatServiceServer(server, implementation) // регистрация обработчиков

	reflection.Register(server) // регистрируем дополнительные обработчики

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}