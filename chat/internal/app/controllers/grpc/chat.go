package chat_grpc

import (
	"context"

	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
	chat "github.com/darialissi/msa_big_tech/chat/pkg"
)

func (s *service) CreateDirectChat(ctx context.Context, req *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {

    // TODO: получить id из jwt MW
	userId := "00000000-0000-0000-0000-0000000000000"

	form := &dto.CreateDirectChat{
		CreatorID: dto.UserID(userId),
		ParticipantID: dto.UserID(req.ParticipantId),
	}

	chatResponse, err := s.ChatUsecase.CreateDirectChat(form)

	if err != nil {
		return nil, err
	}

	return &chat.CreateDirectChatResponse{ChatId: string(chatResponse.ID)}, nil
}

func (s *service) GetChat(ctx context.Context, req *chat.GetChatRequest) (*chat.GetChatResponse, error) {

	res, err := s.ChatUsecase.FetchChat(dto.ChatID(req.ChatId))

	if err != nil {
		return nil, err
	}

	chatResponse := &chat.Chat{
		ChatId: string(res.ID),
		CreatorId: string(res.CreatorID),
		ParticipantId: string(res.ParticipantID),
	}

	return &chat.GetChatResponse{Chat: chatResponse}, nil
}

func (s *service) ListUserChats(ctx context.Context, req *chat.ListUserChatsRequest) (*chat.ListUserChatsResponse, error) {

	res, err := s.ChatUsecase.ListUserChats(dto.UserID(req.UserId))

	if err != nil {
		return nil, err
	}

	chatResponses := make([]*chat.Chat, len(res))

    for i, ch := range res {
        chatResponses[i] = &chat.Chat{
			ChatId: string(ch.ID),
			CreatorId: string(ch.CreatorID),
			ParticipantId: string(ch.ParticipantID),
		}
    }

	return &chat.ListUserChatsResponse{Chats: chatResponses}, nil
}

func (s *service) ListChatMembers(ctx context.Context, req *chat.ListChatMembersRequest) (*chat.ListChatMembersResponse, error) {

	res, err := s.ChatUsecase.ListChatMembers(dto.ChatID(req.ChatId))

	if err != nil {
		return nil, err
	}

	chatMembers := make([]*chat.ChatMember, len(res))

    for i, m := range res {
        chatMembers[i] = &chat.ChatMember{
			UserId: string(m.UserID),
			Role: m.Role,
		}
	}

	return &chat.ListChatMembersResponse{Members: chatMembers}, nil
}

func (s *service) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {

    // TODO: получить id из jwt MW
	userId := "00000000-0000-0000-0000-0000000000000"

	form := &dto.SendMessage{
		Text: req.Text,
		FromUserID: dto.UserID(userId),
		ToChatID: dto.ChatID(req.ChatId),
	}

	res, err := s.ChatUsecase.SendMessage(form)

	if err != nil {
		return nil, err
	}

	chatMessage := &chat.Message{
		MessageId: string(res.ID),
		Text: res.Text,
		FromUserId: string(res.FromUserID),
		ToChatId: string(res.ChatID),
	}

	return &chat.SendMessageResponse{Message: chatMessage}, nil
}

func (s *service) ListMessages(ctx context.Context, req *chat.ListMessagesRequest) (*chat.ListMessagesResponse, error) {

	res, err := s.ChatUsecase.ListMessages(dto.ChatID(req.ChatId))

	if err != nil {
		return nil, err
	}

	chatMessages := make([]*chat.Message, len(res))

    for i, mes := range res {
        chatMessages[i] = &chat.Message{
			MessageId: string(mes.ID),
			Text: mes.Text,
			FromUserId: string(mes.FromUserID),
			ToChatId: string(mes.ChatID),

		}
    }

	return &chat.ListMessagesResponse{Messages: chatMessages}, nil
}

func (s *service) StreamMessages(ctx context.Context, req *chat.StreamMessagesRequest) (*chat.StreamMessagesResponse, error) {

	return &chat.StreamMessagesResponse{Stream: &chat.Message{}}, nil
}