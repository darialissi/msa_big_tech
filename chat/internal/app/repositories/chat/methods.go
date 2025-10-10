package chat

import (
	"context"
	"fmt"
	"strings"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
)


func (r *Repository) Save(ctx context.Context, in *models.DirectChat) (*models.DirectChat, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	if v1 := row.CreatorID.String(); v1 == "" {
		return nil, fmt.Errorf(
			"invalid args: row.CreatorID=%s", 
			v1,
		)
	}

	query := r.sb.
		Insert(chatsTable).
		Columns(
			chatsTableColumnCreatorID, 
			).
		Values(row.CreatorID).
		Suffix("RETURNING " + strings.Join(chatsTableColumns, ","))

	var outRow ChatRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchById(ctx context.Context, chatId models.ChatID) (*models.DirectChat, error) {

	query := r.sb.
		Select(chatsTableColumns...).
		From(chatsTable).
		Where(squirrel.Eq{chatsTableColumnID: string(chatId)})

	var outRow ChatRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // запись не найдена
			return nil, nil
		}
    	return nil, err
	}

	return ToModel(&outRow), nil
}