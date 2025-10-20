package friend

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/darialissi/msa_big_tech/lib/postgres"
	"github.com/jackc/pgx/v5"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
)

func (r *Repository) Save(ctx context.Context, in *models.UserFriend) (*models.UserFriend, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	userID := row.UserID.String()
	friendID := row.FriendID.String()

	// Сортируем ID
	if userID > friendID {
		row.UserID, row.FriendID = row.FriendID, row.UserID
	}

	query := r.sb.
		Insert(friendsTable).
		Columns(
			friendsTableColumnUserID,
			friendsTableColumnFriendID,
		).
		Values(row.UserID, row.FriendID).
		Suffix("RETURNING " + strings.Join(friendsTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow FriendRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if postgres.IsUniqueViolation(err) {
			return nil, fmt.Errorf("friendship already exist")
		}
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) Delete(ctx context.Context, in *models.UserFriend) (*models.UserFriend, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	userID := row.UserID.String()
	friendID := row.FriendID.String()

	if userID == "" || friendID == "" {
		return nil, fmt.Errorf(
			"invalid args: row.UserID=%s, rrow.FriendID=%s",
			userID,
			friendID,
		)
	}

	// Сортируем ID
	if userID > friendID {
		row.UserID, row.FriendID = row.FriendID, row.UserID
	}

	query := r.sb.
		Delete(friendsTable).
		Where(squirrel.And{
			squirrel.Eq{friendsTableColumnUserID: row.UserID},
			squirrel.Eq{friendsTableColumnFriendID: row.FriendID},
		}).
		Suffix("RETURNING " + strings.Join(friendsTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow FriendRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) { // запись не найдена
			return nil, nil
		}
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchManyByUserId(ctx context.Context, userId models.UserID) ([]*models.UserFriend, error) {

	query := r.sb.
		Select(friendsTableColumns...).
		From(friendsTable).
		Where(squirrel.Or{
			squirrel.Eq{friendsTableColumnUserID: string(userId)},
			squirrel.Eq{friendsTableColumnFriendID: string(userId)},
		})

	pool := r.db.GetQueryEngine(ctx)

	var outRows []FriendRow
	if err := pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
		return nil, err // только ошибка БД
	}

	res := make([]*models.UserFriend, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	return res, nil
}

func (r *Repository) FetchManyByUserIdCursor(ctx context.Context, userId models.UserID, cursor *models.Cursor) ([]*models.UserFriend, *models.Cursor, error) {

	query := r.sb.
		Select(friendsTableColumns...).
		From(friendsTable).
		Where(squirrel.Or{
			squirrel.Eq{friendsTableColumnUserID: string(userId)},
			squirrel.Eq{friendsTableColumnFriendID: string(userId)},
		})

	// Курсорная пагинация по created_at ASC
	if cur := cursor.NextCursor; cur != "" {
        cursorTime, err := time.Parse(time.RFC3339Nano, cur)
        if err != nil {
            return nil, nil, fmt.Errorf("invalid cursor: %w", err)
        }
		query = query.Where(squirrel.Gt{friendsTableColumnCreatedAt: cursorTime})
	}

	query = query.
		OrderBy(friendsTableColumnCreatedAt). // Сортируем по дате создания
		Limit(cursor.Limit)

	pool := r.db.GetQueryEngine(ctx)

	// Выполняем запрос и формируем результат
	var outRows []FriendRow
	if err := pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
		return nil, nil, err // только ошибка БД
	}

	res := make([]*models.UserFriend, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	// Перезаписываем NextCursor
	if len(res) > 0 {
		cursor.NextCursor = res[len(res)-1].CreatedAt.Format(time.RFC3339Nano)
	}

	return res, cursor, nil
}
