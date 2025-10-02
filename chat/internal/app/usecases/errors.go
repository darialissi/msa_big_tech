package usecases

import (
	"errors"
)

// Ошибки уровня бизнес логики
var (
	// У пользователя нет чатов
	ErrNoUserChats = errors.New("User does not have any chats")
	// Пользователь не писал сообщения
	ErrNoUserMessages = errors.New("User does not have any messages")
	// Чат не существует
	ErrNotExistedChat = errors.New("Chat does not exist")
	// Сообщение не существует
	ErrNotExistedMessage = errors.New("Message does not exist")
)