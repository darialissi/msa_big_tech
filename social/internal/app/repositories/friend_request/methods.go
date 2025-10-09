package friend_request

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
)


func (r *Repository) Save(ctx context.Context, in *models.FriendRequest) (*models.FriendRequest, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	if v1, v2, v3 := row.FromUserID.String(), row.ToUserID.String(), models.FriendRequestStatus(row.Status).String(); v1 == "" || v2 == "" || v3 == "UNKNOWN" {
		return nil, fmt.Errorf(
			"invalid args: row.FromUserID=%s, row.ToUserID=%s, row.Status=%d %s", 
			v1, 
			v2,
			row.Status,
			v3,
		)
	}

	query := r.sb.
		Insert(friendRequestsTable).
		Columns(
			friendRequestsTableColumnFromUserID, 
			friendRequestsTableColumnToUserID, 
			friendRequestsTableColumnStatus,
			).
		Values(row.FromUserID, row.ToUserID, row.Status).
		Suffix("RETURNING " + strings.Join(friendRequestsTableColumns, ","))

	var outRow FriendRequestRow
	if err := r.pool.Getx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	return ToModel(&outRow), nil
}

func (r *Repository) UpdateStatus(ctx context.Context, in *models.FriendRequest) (*models.FriendRequest, error) {
	row, err := FromModel(in)

	if err != nil {
		return nil, err
	}

	if v1, v2 := row.ID.String(), models.FriendRequestStatus(row.Status).String(); v1 == "" || v2 == "UNKNOWN" {
		return nil, fmt.Errorf(
			"invalid args: row.ID=%s, row.Status=%d %s", 
			v1, 
			row.Status,
			v2,
		)
	}

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

func (r *Repository) FetchById(ctx context.Context, reqId models.FriendRequestID) (*models.FriendRequest, error) {

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

func (r *Repository) FetchManyByUserId(ctx context.Context, userId models.UserID) ([]*models.FriendRequest, error) {

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