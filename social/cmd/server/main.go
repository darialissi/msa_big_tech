package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"msa_big_tech/social/pkg/api/proto"
	grpc_hd "msa_big_tech/social/internal/handlers/grpc"
)


func main() {
	implementation := grpc_hd.NewServer() // наша реализация сервера

	lis, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	social.RegisterSocialServiceServer(server, implementation) // регистрация обработчиков

	reflection.Register(server) // регистрируем дополнительные обработчики

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}