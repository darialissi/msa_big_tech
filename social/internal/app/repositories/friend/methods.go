package friend

import (
	"context"
	"strings"

	"github.com/Masterminds/squirrel"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (r *Repository) Save(ctx context.Context, in *dto.SaveFriend) (*models.UserFriend, error) {
	row := FromSaveDTO(in)

	query := r.sb.
		Insert(friendsTable).
		Columns(
			friendsTableColumnUserID, 
			friendsTableColumnFriendID,
			).
		Values(row.Values()...).
		Suffix("RETURNING " + strings.Join(friendsTableColumns, ","))

	var outRow FriendRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) Delete(ctx context.Context, in *dto.RemoveFriend) (*models.UserFriend, error) {
	row := FromDeleteDTO(in)

	query := r.sb.
		Delete(friendsTable).
		Where(squirrel.Eq{friendsTableColumnUserID: row.UserID}).
		Where(squirrel.Eq{friendsTableColumnFriendID: row.FriendID}).
		Suffix("RETURNING " + strings.Join(friendsTableColumns, ","))

	var outRow FriendRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchManyByUserId(ctx context.Context, userId dto.UserID) ([]*models.UserFriend, error) {

	query := r.sb.
		Select(friendsTableColumns...).
		From(friendsTable).
		Where(squirrel.Eq{friendsTableColumnUserID: string(userId)})

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