package usecases

import (
	"errors"

	"github.com/darialissi/msa_big_tech/social/internal/app/models"
	"github.com/darialissi/msa_big_tech/social/internal/app/usecases/dto"
)


type SocialUsecases interface {
	// Отправление запроса "В Друзья"
    SendFriendRequest(req *dto.SendFriendRequest) (*models.FriendRequest, error)
	// Получение списка входящих запросов "В Друзья" пользователя
    ListFriendRequests(userId dto.UserID) ([]*models.FriendRequest, error)
	// Принятие заявки "В Друзья"
    AcceptFriendRequest(reqId dto.FriendRequestID) (*models.FriendRequest, error)
	// Отклонение заявки "В Друзья"
    DeclineFriendRequest(reqId dto.FriendRequestID) (*models.FriendRequest, error)
	// Удаление из "Друзей"
    RemoveFriend(fr *dto.RemoveFriend) (*models.UserFriend, error)
	// Получение списка "Друзей" пользователя
    ListFriends(userId dto.UserID) ([]*models.UserFriend, error)
}

type FriendRequestRepository interface {
    Save(req *dto.SaveFriendRequest) (*models.FriendRequest, error)
    UpdateStatus(req *dto.ChangeStatus) (*models.FriendRequest, error)
    FetchById(reqId dto.FriendRequestID) (*models.FriendRequest, error)
    FetchManyByUserId(userId dto.UserID) ([]*models.FriendRequest, error)
}

type FriendRepository interface {
    Save(fr *dto.SaveFriend) (*models.UserFriend, error)
    Delete(fr *dto.RemoveFriend) (*models.UserFriend, error)
    FetchManyByUserId(userId dto.UserID) ([]*models.UserFriend, error)
}

// Проверка реализации всех методов интерфейса при компиляции
var _ SocialUsecases = (*SocialUsecase)(nil)

type Deps struct {
    RepoFriendReq FriendRequestRepository
    RepoFriend FriendRepository
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
    
    return &SocialUsecase{deps}, nil
}