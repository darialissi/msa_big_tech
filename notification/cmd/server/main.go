package main

import (
	"context"
	"log"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/darialissi/msa_big_tech/lib/config"
	"github.com/darialissi/msa_big_tech/lib/postgres"
	"github.com/darialissi/msa_big_tech/lib/postgres/transaction_manager"

	"github.com/darialissi/msa_big_tech/notification/internal/app/adapters/consumer"
	inbox "github.com/darialissi/msa_big_tech/notification/internal/app/modules/inbox"
	inbox_repo "github.com/darialissi/msa_big_tech/notification/internal/app/repositories/inbox"
	"github.com/darialissi/msa_big_tech/notification/internal/app/usecases"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

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

	conn, err := postgres.NewConnectionPool(ctx, dbEnvs.DSN(),
		postgres.WithMaxConnIdleTime(time.Minute),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// adapters/repositories
	dedup := consumer.NewInMemoryDeduper(ctx, 24*time.Hour)
	txManager := transaction_manager.New(conn)
	inboxRepo := inbox_repo.NewRepository(txManager)

	// usecases
	handler := usecases.NewUsecase()

	worker := inbox.NewInboxWorker(inboxRepo, txManager, handler,
		inbox.WithBatchSize(10),
		inbox.WithMaxAttempts(10),
		inbox.WithPollInterval(10*time.Second),
	)

	go worker.Run(ctx)

	processor := inbox.NewProcessor(inbox.Deps{
		Repository: inboxRepo,
	})

	kafkaConsumerGroup := kfEnvs.GetConsumerGroup()
	kafkaConsumerName := kfEnvs.GetConsumerName()

	if kafkaConsumerGroup == "" || kafkaConsumerName == "" {
		log.Fatalf("no defined KAFKA_CONSUMER_GROUP or KAFKA_CONSUMER_NAME: kfEnvs=%v", kfEnvs)
	}

	consumer, err := consumer.NewInboxConsumer(
		strings.Split(kfEnvs.GetBrokers(), ","),
		kafkaConsumerGroup,
		kafkaConsumerName,
		dedup,
		processor,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()
	if err := consumer.Run(
		ctx,
		kfEnvs.GetFrReqTopic(),
	); err != nil && ctx.Err() == nil {
		log.Println("consumer stopped:", err)
	}

	log.Println("shutdown")
}
