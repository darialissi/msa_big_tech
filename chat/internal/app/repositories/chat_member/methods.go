package chat_member

import (
	"context"
	"strings"

	"github.com/Masterminds/squirrel"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
)

func (r *Repository) Save(ctx context.Context, in *models.ChatMember) (*models.ChatMember, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	query := r.sb.
		Insert(chatMembersTable).
		Columns(
			chatMembersTableColumnChatID,
			chatMembersTableColumnUserID,
		).
		Values(row.ChatID, row.UserID).
		Suffix("RETURNING " + strings.Join(chatMembersTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow ChatMemberRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) SaveMultiple(ctx context.Context, members []*models.ChatMember) ([]*models.ChatMember, error) {
	if len(members) == 0 {
		return nil, nil
	}

	query := r.sb.
		Insert(chatMembersTable).
		Columns(
			chatMembersTableColumnChatID,
			chatMembersTableColumnUserID,
		)

	// Добавляем значения для каждой записи
	for _, member := range members {
		row, err := FromModel(member)

		if err != nil {
			return nil, err
		}

		query = query.Values(row.ChatID, row.UserID)
	}

	query = query.Suffix("RETURNING " + strings.Join(chatMembersTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRows []ChatMemberRow
	if err := pool.Selectx(ctx, &outRows, query); err != nil {
		return nil, err
	}

	res := make([]*models.ChatMember, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	return res, nil
}

func (r *Repository) FetchManyByChatId(ctx context.Context, chatId models.ChatID) ([]*models.ChatMember, error) {

	query := r.sb.
		Select(chatMembersTableColumns...).
		From(chatMembersTable).
		Where(squirrel.Eq{chatMembersTableColumnChatID: string(chatId)})

	pool := r.db.GetQueryEngine(ctx)

	var outRows []ChatMemberRow
	if err := pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
		return nil, err // только ошибка БД
	}

	res := make([]*models.ChatMember, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	return res, nil
}

func (r *Repository) FetchManyByUserId(ctx context.Context, userId models.UserID) ([]*models.ChatMember, error) {

	query := r.sb.
		Select(chatMembersTableColumns...).
		From(chatMembersTable).
		Where(squirrel.Eq{chatMembersTableColumnUserID: string(userId)})

	pool := r.db.GetQueryEngine(ctx)

	var outRows []ChatMemberRow
	if err := pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
		return nil, err // только ошибка БД
	}

	res := make([]*models.ChatMember, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	return res, nil
}
