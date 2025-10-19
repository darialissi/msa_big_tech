package message

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
)

func (r *Repository) Save(ctx context.Context, in *models.Message) (*models.Message, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	if v1, v2 := row.ChatID.String(), row.SenderID.String(); v1 == "" || v2 == "" {
		return nil, fmt.Errorf(
			"invalid args: row.ChatID=%s, row.SenderID=%s",
			v1,
			v2,
		)
	}

	query := r.sb.
		Insert(messagesTable).
		Columns(
			messagesTableColumnChatID,
			messagesTableColumnSenderID,
			messagesTableColumnText,
		).
		Values(row.ChatID, row.SenderID, row.Text).
		Suffix("RETURNING " + strings.Join(messagesTableColumns, ","))

	var outRow MessageRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchById(ctx context.Context, messageId models.MessageID) (*models.Message, error) {

	query := r.sb.
		Select(messagesTableColumns...).
		From(messagesTable).
		Where(squirrel.Eq{messagesTableColumnID: string(messageId)})

	var outRow MessageRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // запись не найдена
			return nil, nil
		}
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchManyByChatIdCursor(ctx context.Context, chatId models.ChatID, cursor *models.Cursor) ([]*models.Message, *models.Cursor, error) {

	query := r.sb.
		Select(messagesTableColumns...).
		From(messagesTable).
		Where(squirrel.Eq{messagesTableColumnChatID: string(chatId)})

	// добавляем курсорную пагинацию
	if cur := string(cursor.NextCursor); cur != "" {
		query = query.Where(squirrel.Lt{messagesTableColumnChatID: cur})
	}

	query = query.
		OrderBy(messagesTableColumnChatID + " DESC").
		Limit(cursor.Limit)

	// выполняем запрос и формируем результат
	var outRows []MessageRow
	if err := r.pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
		return nil, nil, err // только ошибка БД
	}

	res := make([]*models.Message, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	// перезаписываем NextCursor
	if len(res) > 0 {
		cursor.NextCursor = res[len(res)-1].ID
	}

	return res, cursor, nil
}

func (r *Repository) StreamMany(ctx context.Context, chatId models.ChatID, sinceUx time.Time) ([]*models.Message, error) {

	query := r.sb.
		Select(messagesTableColumns...).
		From(messagesTable).
		Where(squirrel.Eq{messagesTableColumnChatID: string(chatId)}).
		Where(squirrel.Gt{messagesTableColumnCreatedAt: sinceUx})

	var outRows []MessageRow
	if err := r.pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
		return nil, err // только ошибка БД
	}

	res := make([]*models.Message, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	return res, nil
}
