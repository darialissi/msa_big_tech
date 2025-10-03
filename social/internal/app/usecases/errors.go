package usecases

import (
	"errors"
)

// Ошибки уровня бизнес логики
var (
	// У пользователя нет входящих запросов "В Друзья"
	ErrUserNoFriendRequestsIn = errors.New("User has no incoming friend requests")
	// У пользователя нет исходящих запросов "В Друзья"
	ErrUserNoFriendRequestsOut = errors.New("User has no outcoming friend requests")
	// У пользователя нет "Друзей"
	ErrUserNoFriends = errors.New("User has no friends")
	// Пользователь отсутствует в "Друзьях"
	ErrNoFriend = errors.New("No friend found")
	// Запроса "В Друзья" не существует
	ErrNoFriendRequest = errors.New("Friend request does not exist")
	// Недоступный статус перехода
	ErrBadStatus = errors.New("Transition is not available")
)