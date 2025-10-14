package usecases

import (
	"errors"
    "context"
    "time"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
)


type ChatUsecases interface {
	// Создание личного чата
    CreateDirectChat(ctx context.Context, chat *dto.CreateDirectChat) (*models.DirectChat, error)
	// Получение чата
    FetchChat(ctx context.Context, chatId dto.ChatID) (*models.DirectChat, error)
	// Получение списка чатов пользователя
    ListUserChats(ctx context.Context, userId dto.UserID) ([]*models.ChatMember, error)
	// Получение списка участников чата
    ListChatMembers(ctx context.Context, chatId dto.ChatID) ([]*models.ChatMember, error)
	// Отправление сообщения
    SendMessage(ctx context.Context, m *dto.SendMessage) (*models.Message, error)
	// Получение истории сообщений чата
    ListMessages(ctx context.Context, lm *dto.ListMessages) ([]*models.Message, *models.Cursor, error)
	// Серверный стрим новых сообщений
    StreamMessages(ctx context.Context, st *dto.StreamMessages) ([]*models.Message, error)
}

type ChatRepository interface {                                
    Save(ctx context.Context, in *models.DirectChat) (*models.DirectChat, error)
    FetchById(ctx context.Context, in models.ChatID) (*models.DirectChat, error)
}

type ChatMemberRepository interface {       
    Save(ctx context.Context, in *models.ChatMember) (*models.ChatMember, error)    
    SaveMultiple(ctx context.Context, members []*models.ChatMember) ([]*models.ChatMember, error)                        
    FetchManyByUserId(ctx context.Context, chatId models.UserID) ([]*models.ChatMember, error)                 
    FetchManyByChatId(ctx context.Context, chatId models.ChatID) ([]*models.ChatMember, error)
}

type MessageRepository interface {
    Save(ctx context.Context, in *models.Message) (*models.Message, error)
    FetchById(ctx context.Context, messageId models.MessageID) (*models.Message, error)
    FetchManyByChatIdCursor(ctx context.Context, chatId models.ChatID, cursor *models.Cursor) ([]*models.Message, *models.Cursor, error)
    StreamMany(ctx context.Context, chatId models.ChatID, sinceUx time.Time) ([]*models.Message, error)
}

// Проверка реализации всех методов интерфейса при компиляции
var _ ChatUsecases = (*ChatUsecase)(nil)

type Deps struct {
    RepoChat ChatRepository
    RepoMessage MessageRepository
	RepoChatMember ChatMemberRepository
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
	repoChatMember ChatMemberRepository
}

func NewChatUsecase(deps Deps) *ChatUsecase {
    return &ChatUsecase{
        repoChat: deps.RepoChat, // переупаковка
        repoMessage: deps.RepoMessage,
        repoChatMember: deps.RepoChatMember,
    }
}