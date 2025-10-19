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
	// В чате нет сообщений
	ErrNoMessages = errors.New("Chat has no messages")
	// Сообщение не существует
	ErrNotExistedMessage = errors.New("Message does not exist")
)
