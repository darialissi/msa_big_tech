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

	users_grpc "github.com/darialissi/msa_big_tech/users/internal/app/controllers/grpc"
	errors_mw "github.com/darialissi/msa_big_tech/users/internal/app/middleware/errors"
	user_repo "github.com/darialissi/msa_big_tech/users/internal/app/repositories/user"
	"github.com/darialissi/msa_big_tech/users/internal/app/usecases"
	users "github.com/darialissi/msa_big_tech/users/pkg"
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
	usersRepo := user_repo.NewRepository(txMngr)

	usersUC := usecases.NewUsersUsecase(usersRepo, txMngr)

	implementation := users_grpc.NewServer(usersUC) // наша реализация сервера

	lis, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			errors_mw.RecoveryUnaryInterceptor(), // сначала recovery для перехвата паник
			errors_mw.ErrorsUnaryInterceptor(),   // затем errors для конвертации ошибок
		),
	)
	users.RegisterUsersServiceServer(server, implementation) // регистрация обработчиков

	reflection.Register(server) // регистрируем дополнительные обработчики

	log.Printf("server listening at %v", lis.Addr())
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
