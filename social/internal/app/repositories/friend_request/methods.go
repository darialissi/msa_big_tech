package friend_request

import (
	"context"
	"strings"

	"github.com/Masterminds/squirrel"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (r *Repository) Save(ctx context.Context, in *dto.SaveFriendRequest) (*models.FriendRequest, error) {
	row := FromSaveDTO(in)

	query := r.sb.
		Insert(friendRequestsTable).
		Columns(
			friendRequestsTableColumnFromUserID, 
			friendRequestsTableColumnToUserID, 
			friendRequestsTableColumnStatus,
			).
		Values(row.Values()...).
		Suffix("RETURNING " + strings.Join(friendRequestsTableColumns, ","))

	var outRow FriendRequestRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) UpdateStatus(ctx context.Context, in *dto.UpdateFriendRequest) (*models.FriendRequest, error) {
	row := FromUpdateDTO(in)

	query := r.sb.
		Update(friendRequestsTable).
		Set(friendRequestsTableColumnStatus, row.Status).
		Where(squirrel.Eq{friendRequestsTableColumnID: row.ID}).
		Suffix("RETURNING " + strings.Join(friendRequestsTableColumns, ","))

	var outRow FriendRequestRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchById(ctx context.Context, reqId dto.FriendRequestID) (*models.FriendRequest, error) {

	query := r.sb.
		Select(friendRequestsTableColumns...).
		From(friendRequestsTable).
		Where(squirrel.Eq{friendRequestsTableColumnID: string(reqId)})

	var outRow FriendRequestRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil { // возвращает запись или ошибку при отсутствии
		return nil, err // ошибка, если не найдена запись
	}

	return ToModel(&outRow), nil
}

func (r *Repository) FetchManyByUserId(ctx context.Context, userId dto.UserID) ([]*models.FriendRequest, error) {

	query := r.sb.
		Select(friendRequestsTableColumns...).
		From(friendRequestsTable).
		Where(squirrel.Eq{friendRequestsTableColumnToUserID: string(userId)}) // входящие запросы

	var outRows []FriendRequestRow
	if err := r.pool.Selectx(ctx, &outRows, query); err != nil { // возвращает слайс
		return nil, err // только ошибка БД
	}

	res := make([]*models.FriendRequest, len(outRows))

	for i, outRow := range outRows {
		res[i] = ToModel(&outRow)
	}

	return res, nil
}