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

	social_grpc "github.com/darialissi/msa_big_tech/social/internal/app/controllers/grpc"
	friend_repo "github.com/darialissi/msa_big_tech/social/internal/app/repositories/friend"
	friend_req_repo "github.com/darialissi/msa_big_tech/social/internal/app/repositories/friend_request"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases"
	social "github.com/darialissi/msa_big_tech/social/pkg"
)

func main() {

	appEnvs := config.AppConfig()
	dbEnvs := config.DbConfig(appEnvs.GetMode())

	if appErr, dbErr := appEnvs.Validate(), dbEnvs.Validate(); appErr != nil || dbErr != nil {
		log.Fatalf("failed to load env: appErr=%v dbErr=%v", appErr, dbErr)
	}

	ctx := context.Background()
	conn, err := postgres.NewConnectionPool(ctx, dbEnvs.DSN(),
		postgres.WithMaxConnIdleTime(time.Minute),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	txMngr := transaction_manager.New(conn)
	friendRepo := friend_repo.NewRepository(txMngr)
	friendReqRepo := friend_req_repo.NewRepository(txMngr)

	// DI
	deps := usecases.Deps{
		RepoFriend:    friendRepo,
		RepoFriendReq: friendReqRepo,
		TxMan:         txMngr,
	}

	socialUC, err := usecases.NewSocialUsecase(deps)
	if err != nil {
		log.Fatalf("failed to create social usecase: %v", err)
	}

	implementation := social_grpc.NewServer(socialUC) // наша реализация сервера

	lis, err := net.Listen("tcp", ":50055")
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
