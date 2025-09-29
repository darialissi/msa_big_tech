package main

import (
	"log"
	"net"
    "path/filepath"
    "runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"

	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases"
	auth_grpc "github.com/darialissi/msa_big_tech/auth/internal/controllers/grpc"
	auth "github.com/darialissi/msa_big_tech/auth/pkg"
	auth_repo "github.com/darialissi/msa_big_tech/auth/internal/app/repositories/auth"
	token_repo "github.com/darialissi/msa_big_tech/auth/internal/app/repositories/token"
)

// init is invoked before main()
func init() {
    // abs path
    _, filename, _, _ := runtime.Caller(0)
    rootDir := filepath.Join(filepath.Dir(filename), "../..")
    envPath := filepath.Join(rootDir, ".env")

    // loads values from .env into the system
    if err := godotenv.Load(envPath); err != nil {
        log.Print(err)
    }
}

func main() {
    // DI
    authRepo := auth_repo.NewRepository()
    tokenRepo := token_repo.NewRepository()
    
    authUC := usecases.NewAuthUsecase(authRepo, tokenRepo)
	
    implementation := auth_grpc.NewServer(authUC)

	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	auth.RegisterAuthServiceServer(server, implementation) // регистрация обработчиков

	reflection.Register(server) // регистрируем дополнительные обработчики

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// Register:
	// grpcurl -plaintext -d '{"Email": "hi", "Password": "bye"}' localhost:8083 AuthService/Register
}