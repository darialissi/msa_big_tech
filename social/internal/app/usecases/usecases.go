package usecases

import (
	"context"
	"errors"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)

type SocialUsecases interface {
	// Отправление запроса "В Друзья"
	SendFriendRequest(ctx context.Context, req *dto.SendFriendRequest) (*models.FriendRequest, error)
	// Получение списка входящих запросов "В Друзья" пользователя
	ListFriendRequests(ctx context.Context, userId dto.UserID) ([]*models.FriendRequest, error)
	// Принятие заявки "В Друзья"
	AcceptFriendRequest(ctx context.Context, reqId dto.FriendRequestID) (*models.FriendRequest, error)
	// Отклонение заявки "В Друзья"
	DeclineFriendRequest(ctx context.Context, reqId dto.FriendRequestID) (*models.FriendRequest, error)
	// Удаление из "Друзей"
	RemoveFriend(ctx context.Context, fr *dto.RemoveFriend) (*models.UserFriend, error)
	// Получение списка "Друзей" пользователя с пагинацией
	ListFriends(ctx context.Context, lf *dto.ListFriends) ([]*models.UserFriend, *models.Cursor, error)
}

type FriendRequestRepository interface {
	Save(ctx context.Context, in *models.FriendRequest) (*models.FriendRequest, error)
	UpdateStatus(ctx context.Context, in *models.FriendRequest) (*models.FriendRequest, error)
	FetchById(ctx context.Context, reqId models.FriendRequestID) (*models.FriendRequest, error)
	FetchManyByUserId(ctx context.Context, userId models.UserID) ([]*models.FriendRequest, error)
}

type FriendRepository interface {
	Save(ctx context.Context, in *models.UserFriend) (*models.UserFriend, error)
	Delete(ctx context.Context, in *models.UserFriend) (*models.UserFriend, error)
	FetchManyByUserId(ctx context.Context, userId models.UserID) ([]*models.UserFriend, error)
	FetchManyByUserIdCursor(ctx context.Context, userId models.UserID, cursor *models.Cursor) ([]*models.UserFriend, *models.Cursor, error)
}

type OutboxRepository interface {
	SaveFriendRequestCreatedID(ctx context.Context, id models.FriendRequestID) error
	SaveFriendRequestUpdatedID(ctx context.Context, id models.FriendRequestID) error
}

// Проверка реализации всех методов интерфейса при компиляции
var _ SocialUsecases = (*SocialUsecase)(nil)

type TxManager interface {
	RunReadCommitted(ctx context.Context, f func(txCtx context.Context) error) error
}

type Deps struct {
	RepoFriendReq FriendRequestRepository
	RepoFriend    FriendRepository
	RepoOutbox    OutboxRepository
	TxMan         TxManager
}

type SocialUsecase struct {
	Deps // встраивание
}

func NewSocialUsecase(deps Deps) (*SocialUsecase, error) {
	if deps.RepoFriendReq == nil {
		return nil, errors.New("FriendRequestRepository is required")
	}
	if deps.RepoFriend == nil {
		return nil, errors.New("FriendRepository is required")
	}
	if deps.RepoOutbox == nil {
		return nil, errors.New("OutboxRepository is required")
	}
	if deps.TxMan == nil {
		return nil, errors.New("TxMan is required")
	}

	return &SocialUsecase{deps}, nil
}
