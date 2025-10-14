package friend

import (
	"context"
	"strings"
	"fmt"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
)


func (r *Repository) Save(ctx context.Context, in *models.UserFriend) (*models.UserFriend, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	if v1, v2 := row.UserID.String(), row.FriendID.String(); v1 == "" || v2 == "" {
		return nil, fmt.Errorf(
			"invalid args: row.UserID=%s, rrow.FriendID=%s", 
			v1,
			v2,
		)
	}

	query := r.sb.
		Insert(friendsTable).
		Columns(
			friendsTableColumnUserID, 
			friendsTableColumnFriendID,
			).
		Values(row.UserID, row.FriendID).
		Suffix("RETURNING " + strings.Join(friendsTableColumns, ","))

	var outRow FriendRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) Delete(ctx context.Context, in *models.UserFriend) (*models.UserFriend, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	if v1, v2 := row.UserID.String(), row.FriendID.String(); v1 == "" || v2 == "" {
		return nil, fmt.Errorf(
			"invalid args: row.UserID=%s, row.FriendID=%s", 
			v1,
			v2,
		)
	}

	query := r.sb.
		Delete(friendsTable).
		Where(squirrel.Or{
			squirrel.Eq{friendsTableColumnUserID: row.UserID},
			squirrel.Eq{friendsTableColumnFriendID: row.FriendID},
		}).
		Where(squirrel.Or{
			squirrel.Eq{friendsTableColumnUserID: row.FriendID},
			squirrel.Eq{friendsTableColumnFriendID: row.UserID},
		}).
		Suffix("RETURNING " + strings.Join(friendsTableColumns, ","))

	var outRow FriendRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
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

	var outRows []FriendRow
	if err := r.pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
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

	// добавляем курсорную пагинацию
	if cur := string(cursor.NextCursor); cur != "" {
		query = query.Where(squirrel.Gt{friendsTableColumnUserID: cur})
	}

	query = query.
        OrderBy(friendsTableColumnUserID).
        Limit(cursor.Limit)

	// выполняем запрос и формируем результат
	var outRows []FriendRow
	if err := r.pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
		return nil, nil, err // только ошибка БД
	}

	res := make([]*models.UserFriend, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	// перезаписываем NextCursor
	if len(res) > 0 {
		cursor.NextCursor = res[len(res)-1].UserID
	}

	return res, cursor, nil
}