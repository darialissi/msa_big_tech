package main

import (
	"context"
	"log"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/darialissi/msa_big_tech/lib/config"
	"github.com/darialissi/msa_big_tech/lib/kafka"
	"github.com/darialissi/msa_big_tech/lib/postgres"
	"github.com/darialissi/msa_big_tech/lib/postgres/transaction_manager"

	friend_req_event "github.com/darialissi/msa_big_tech/social/internal/app/adapters/friend_request_event_handler"
	social_grpc "github.com/darialissi/msa_big_tech/social/internal/app/controllers/grpc"
	errors_mw "github.com/darialissi/msa_big_tech/social/internal/app/middleware/errors"
	outbox "github.com/darialissi/msa_big_tech/social/internal/app/modules/outbox"
	friend_repo "github.com/darialissi/msa_big_tech/social/internal/app/repositories/friend"
	friend_req_repo "github.com/darialissi/msa_big_tech/social/internal/app/repositories/friend_request"
	outbox_repo "github.com/darialissi/msa_big_tech/social/internal/app/repositories/outbox"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases"
	social "github.com/darialissi/msa_big_tech/social/pkg"
)

func main() {

	appEnvs := config.AppConfig()
	dbEnvs := config.DbConfig(appEnvs.GetMode())
	kfEnvs := config.KfConfig(appEnvs.GetMode())

	if appErr := appEnvs.Validate(); appErr != nil {
		log.Fatalf("failed to load env: appErr=%v", appErr)
	}

	if dbErr := dbEnvs.Validate(); dbErr != nil {
		log.Fatalf("failed to load env: dbErr=%v", dbErr)
	}

	if kfErr := kfEnvs.Validate(); kfErr != nil {
		log.Fatalf("failed to load env: kfErr=%v", kfErr)
	}

	ctx := context.Background()
	conn, err := postgres.NewConnectionPool(ctx, dbEnvs.DSN(),
		postgres.WithMaxConnIdleTime(time.Minute),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	producer, err := kafka.NewNewSyncProducer(strings.Split(kfEnvs.GetBrokers(), ","), nil)
	if err != nil {
		log.Fatal(err)
	}

	friendRequestEventsHandler := friend_req_event.NewKafkaFriendRequestBatchHandler(producer,
		friend_req_event.WithMaxBatchSize(100),
		friend_req_event.WithTopic(kfEnvs.GetFrReqTopic()),
	)

	txMngr := transaction_manager.New(conn)
	friendRepo := friend_repo.NewRepository(txMngr)
	friendReqRepo := friend_req_repo.NewRepository(txMngr)
	outboxRepo := outbox_repo.NewRepository(txMngr)

	worker := outbox.NewOutboxFriendRequestWorker(outboxRepo, txMngr, friendRequestEventsHandler,
		outbox.WithBatchSize(10),
		outbox.WithMaxRetry(10),
		outbox.WithRetryInterval(30*time.Second),
		outbox.WithWindow(time.Hour),
		outbox.WithPollInterval(10*time.Second), // явно задаем интервал опроса
	)

	go worker.Run(ctx)

	outboxProcessor := outbox.NewProcessor(outbox.Deps{
		Repository: outboxRepo,
	})

	// DI
	deps := usecases.Deps{
		RepoFriend:    friendRepo,
		RepoFriendReq: friendReqRepo,
		RepoOutbox:    outboxProcessor,
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

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			errors_mw.RecoveryUnaryInterceptor(), // сначала recovery для перехвата паник
			errors_mw.ErrorsUnaryInterceptor(),   // затем errors для конвертации ошибок
		),
	)
	social.RegisterSocialServiceServer(server, implementation) // регистрация обработчиков

	reflection.Register(server) // регистрируем дополнительные обработчики

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
