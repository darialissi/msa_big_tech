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

	query := r.sb.
		Insert(messagesTable).
		Columns(
			messagesTableColumnChatID,
			messagesTableColumnSenderID,
			messagesTableColumnText,
		).
		Values(row.ChatID, row.SenderID, row.Text).
		Suffix("RETURNING " + strings.Join(messagesTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow MessageRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchById(ctx context.Context, messageId models.MessageID) (*models.Message, error) {

	query := r.sb.
		Select(messagesTableColumns...).
		From(messagesTable).
		Where(squirrel.Eq{messagesTableColumnID: string(messageId)})

	pool := r.db.GetQueryEngine(ctx)

	var outRow MessageRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
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

	// Курсорная пагинация по created_at DESC
	if cur := cursor.NextCursor; cur != "" {
		cursorTime, err := time.Parse(time.RFC3339Nano, cur)
		if err != nil {
			return nil, nil, fmt.Errorf("invalid cursor: %w", err)
		}
		query = query.Where(squirrel.Lt{messagesTableColumnCreatedAt: cursorTime})
	}

	query = query.
		OrderBy(messagesTableColumnCreatedAt + " DESC"). // Сортируем по дате DESC
		Limit(cursor.Limit)

	pool := r.db.GetQueryEngine(ctx)

	// Выполняем запрос и формируем результат
	var outRows []MessageRow
	if err := pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
		return nil, nil, err // только ошибка БД
	}

	res := make([]*models.Message, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	// Перезаписываем NextCursor
	if len(res) > 0 {
		cursor.NextCursor = res[len(res)-1].CreatedAt.Format(time.RFC3339Nano)
	}

	return res, cursor, nil
}

func (r *Repository) StreamMany(ctx context.Context, chatId models.ChatID, sinceUx time.Time) ([]*models.Message, error) {

	query := r.sb.
		Select(messagesTableColumns...).
		From(messagesTable).
		Where(squirrel.Eq{messagesTableColumnChatID: string(chatId)}).
		Where(squirrel.Gt{messagesTableColumnCreatedAt: sinceUx})

	pool := r.db.GetQueryEngine(ctx)

	var outRows []MessageRow
	if err := pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
		return nil, err // только ошибка БД
	}

	res := make([]*models.Message, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	return res, nil
}
