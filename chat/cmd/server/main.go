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

	errors_mw "github.com/darialissi/msa_big_tech/chat/internal/app/middleware/errors"
	chat_grpc "github.com/darialissi/msa_big_tech/chat/internal/app/controllers/grpc"
	chat_repo "github.com/darialissi/msa_big_tech/chat/internal/app/repositories/chat"
	chat_member_repo "github.com/darialissi/msa_big_tech/chat/internal/app/repositories/chat_member"
	message_repo "github.com/darialissi/msa_big_tech/chat/internal/app/repositories/message"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases"
	chat "github.com/darialissi/msa_big_tech/chat/pkg"
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

	chatRepo := chat_repo.NewRepository(txMngr)
	chatMemberRepo := chat_member_repo.NewRepository(txMngr)
	messageRepo := message_repo.NewRepository(txMngr)

	// DI
	deps := usecases.Deps{
		RepoChat:       chatRepo,
		RepoMessage:    messageRepo,
		RepoChatMember: chatMemberRepo,
		TxMan:          txMngr,
	}

	if err := deps.Valid(); err != nil {
		log.Fatalf("failed to create chat usecase: %v", err)
	}

	chatUC := usecases.NewChatUsecase(deps)

	implementation := chat_grpc.NewServer(chatUC) // наша реализация сервера

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			errors_mw.RecoveryUnaryInterceptor(), // сначала recovery для перехвата паник
			errors_mw.ErrorsUnaryInterceptor(),   // затем errors для конвертации ошибок
		),
	)
	chat.RegisterChatServiceServer(server, implementation) // регистрация обработчиков

	reflection.Register(server) // регистрируем дополнительные обработчики

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
