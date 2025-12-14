package inbox

import (
	"github.com/darialissi/msa_big_tech/lib/postgres/transaction_manager"

	"github.com/Masterminds/squirrel"
)

type Repository struct {
	db transaction_manager.TransactionManagerAPI
	qb squirrel.StatementBuilderType
}

func NewRepository(p transaction_manager.TransactionManagerAPI) *Repository {
	return &Repository{
		db: p,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
