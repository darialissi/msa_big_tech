package usecases

import (
	"errors"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


type ChatUsecases interface {
	// Создание личного чата
    CreateDirectChat(chat *dto.CreateDirectChat) (*models.DirectChat, error)
	// Получение чата
    FetchChat(chatId dto.ChatID) (*models.DirectChat, error)
	// Получение списка чатов пользователя
    ListUserChats(userId dto.UserID) ([]*models.DirectChat, error)
	// Получение списка участников чата
    ListChatMembers(chatId dto.ChatID) ([]*models.ChatParticipant, error)
	// Отправление сообщения
    SendMessage(m *dto.SendMessage) (*models.Message, error)
	// Получение истории сообщений чата
    ListMessages(chatId dto.ChatID) ([]*models.Message, error)
	// Серверный стрим новых сообщений
    StreamMessages(chatId dto.ChatID) ([]*models.Message, error)
}

type ChatRepository interface {
    Save(chat *dto.CreateDirectChat) (*models.DirectChat, error)
    FetchById(chatId dto.ChatID) (*models.DirectChat, error)
    FetchChatMembers(chatId dto.ChatID) ([]*models.ChatParticipant, error)
    FetchManyByUserId(userId dto.UserID) ([]*models.DirectChat, error)
}

type MessageRepository interface {
    Save(message *dto.SendMessage) (*models.Message, error)
    FetchById(messageId dto.MessageID) (*models.Message, error)
    FetchManyByChatId(chatId dto.ChatID) ([]*models.Message, error)
}

// Проверка реализации всех методов интерфейса при компиляции
var _ ChatUsecases = (*ChatUsecase)(nil)

type Deps struct {
    RepoChat ChatRepository
    RepoMessage MessageRepository
}

func (d *Deps) Valid() error {
    if d.RepoChat == nil {
        return errors.New("ChatRepository is required")
    }
    if d.RepoMessage == nil {
        return errors.New("MessageRepository is required")
    }
    return nil
}

type ChatUsecase struct {
    repoChat ChatRepository
    repoMessage MessageRepository
}

func NewChatUsecase(deps Deps) *ChatUsecase {
    return &ChatUsecase{
        repoChat: deps.RepoChat, // переупаковка
        repoMessage: deps.RepoMessage,
    }
}