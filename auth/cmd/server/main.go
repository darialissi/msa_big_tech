package main

import (
	"context"
	"log"
	"net"
	"path/filepath"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"

	auth_grpc "github.com/darialissi/msa_big_tech/auth/internal/app/controllers/grpc"
	auth_repo "github.com/darialissi/msa_big_tech/auth/internal/app/repositories/auth"
	token_repo "github.com/darialissi/msa_big_tech/auth/internal/app/repositories/token"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases"
	auth "github.com/darialissi/msa_big_tech/auth/pkg"
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
	ctx := context.Background()

	pool, err := NewPostgresConnection(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

    authRepo := auth_repo.NewRepository(pool)
    tokenRepo := token_repo.NewRepository()
    
    deps := usecases.Deps{
        RepoAuth:  authRepo,
        RepoToken: tokenRepo,
    }
    
    authUC, err := usecases.NewAuthUsecase(deps)
    if err != nil {
        log.Fatalf("failed to create auth usecase: %v", err)
    }
	
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