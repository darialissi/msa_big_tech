package main

import (
	"context"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/darialissi/msa_big_tech/lib/config"
	"github.com/darialissi/msa_big_tech/lib/postgres"
	"github.com/darialissi/msa_big_tech/lib/postgres/transaction_manager"

	auth_grpc "github.com/darialissi/msa_big_tech/auth/internal/app/controllers/grpc"
	auth_repo "github.com/darialissi/msa_big_tech/auth/internal/app/repositories/auth"
	token_repo "github.com/darialissi/msa_big_tech/auth/internal/app/repositories/token"
	"github.com/darialissi/msa_big_tech/auth/internal/app/usecases"
	auth "github.com/darialissi/msa_big_tech/auth/pkg"
)

func main() {
	appEnvs := config.AppConfig()
	dbEnvs := config.DbConfig(appEnvs.GetMode())
	jwtEnvs := config.JWTConfig()

	if appErr, dbErr, jwtErr := appEnvs.Validate(), dbEnvs.Validate(), jwtEnvs.Validate(); appErr != nil || dbErr != nil || jwtErr != nil {
		log.Fatalf("failed to load env: appErr=%v dbErr=%v jwtErr=%v", appErr, dbErr, jwtErr)
	}

	// TODO: вынести в middleware
	jwtSecret := jwtEnvs.GetSecret()

	// прокидываем jwtSecret в контекст
	ctx := context.WithValue(context.Background(), "jwtSecret", jwtSecret)

	conn, err := postgres.NewConnectionPool(ctx, dbEnvs.DSN(),
		postgres.WithMaxConnIdleTime(time.Minute),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	txMngr := transaction_manager.New(conn)

	// DI
	authRepo := auth_repo.NewRepository(txMngr)
	tokenRepo := token_repo.NewRepository()

	deps := usecases.Deps{
		RepoAuth:  authRepo,
		RepoToken: tokenRepo,
		TxMan: txMngr,
	}

	authUC, err := usecases.NewAuthUsecase(deps)
	if err != nil {
		log.Fatalf("failed to create auth usecase: %v", err)
	}

	implementation := auth_grpc.NewServer(authUC)

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TODO: добавить JWTInterceptor
	server := grpc.NewServer()
	// server := grpc.NewServer(
	//     grpc.UnaryInterceptor(interceptors.JWTInterceptor(jwtSecret)),
	// )

	auth.RegisterAuthServiceServer(server, implementation) // регистрация обработчиков

	reflection.Register(server) // регистрируем дополнительные обработчики

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
