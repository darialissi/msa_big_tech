package chat

import (
	"context"
	"errors"
	"strings"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
)

func (r *Repository) Save(ctx context.Context, in *models.DirectChat) (*models.DirectChat, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	query := r.sb.
		Insert(chatsTable).
		Columns(
			chatsTableColumnCreatorID,
		).
		Values(row.CreatorID).
		Suffix("RETURNING " + strings.Join(chatsTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow ChatRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchById(ctx context.Context, chatId models.ChatID) (*models.DirectChat, error) {

	query := r.sb.
		Select(chatsTableColumns...).
		From(chatsTable).
		Where(squirrel.Eq{chatsTableColumnID: string(chatId)})

	pool := r.db.GetQueryEngine(ctx)

	var outRow ChatRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // запись не найдена
			return nil, nil
		}
		return nil, err
	}

	return ToModel(&outRow), nil
}
