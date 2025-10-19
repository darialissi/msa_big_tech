package chat

import (
	"github.com/Masterminds/squirrel"
	"github.com/darialissi/msa_big_tech/lib/postgres/transaction_manager"
)

type Repository struct {
	db transaction_manager.TransactionManagerAPI
	sb squirrel.StatementBuilderType
}

func NewRepository(txManager transaction_manager.TransactionManagerAPI) *Repository {
	return &Repository{
		db: txManager,
		sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
