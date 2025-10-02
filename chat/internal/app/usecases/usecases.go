package usecases

import (
	"errors"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


type ChatUsecases interface {
	// Создание личного чата
    CreateDirectChat(userId dto.UserID) (models.ChatID, error)
	// Получение чата
    FetchChat(chatId dto.ChatID) (*models.Chat, error)
	// Получение списка чатов пользователя
    ListUserChats(userId dto.UserID) ([]*models.Chat, error)
	// Получение списка участников чата
    ListChatMembers(chatId dto.ChatID) ([]models.UserID, error)
	// Отправление сообщения
    SendMessage(m *dto.SendMessage) (*models.Message, error)
	// Получение истории сообщений чата
    ListMessages(chatId dto.ChatID) ([]*models.Message, error)
	// Серверный стрим новых сообщений
    StreamMessages(chatId dto.ChatID) ([]*models.Message, error)
}

type ChatRepository interface {
    Save(chat *dto.CreateChat) (models.ChatID, error)
    FetchById(chatId dto.ChatID) (*models.Chat, error)
    FetchManyByUserId(userId dto.UserID) ([]*models.Chat, error)
}

type MessageRepository interface {
    Save(message *dto.SendMessage) (models.MessageID, error)
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