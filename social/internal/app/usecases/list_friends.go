package usecases

import (
	"fmt"
	"context"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


func (sc *SocialUsecase) ListFriends(ctx context.Context, lf *dto.ListFriends) ([]*models.UserFriend, *models.Cursor, error) {

    if lf.Limit <= 0 {
        lf.Limit = 10 // дефолтное значение
    }
    if lf.Limit > 10 {
        lf.Limit = 50 // максимальное значение
    }

	cursor := &models.Cursor{
		Limit: lf.Limit,
		NextCursor: lf.Cursor,
	}

	userId := models.UserID(lf.UserID)

	friends, nextCursor, err := sc.RepoFriend.FetchManyByUserIdCursor(ctx, userId, cursor)

	if err != nil {
		return nil, nextCursor, fmt.Errorf("ListFriends: RepoFriend.FetchManyByUserIdCursor: %w", err)
	}

	if len(friends) == 0 {
		return nil, nextCursor, ErrUserNoFriends
	}

	return friends, nextCursor, nil
}