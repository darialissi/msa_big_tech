package chat_grpc

import (
	"context"

	"buf.build/go/protovalidate"

	"github.com/google/uuid"

	"github.com/darialissi/msa_big_tech/chat/internal/app/models"
	"github.com/darialissi/msa_big_tech/chat/internal/app/usecases/dto"
	chat "github.com/darialissi/msa_big_tech/chat/pkg"
)

func (s *service) CreateDirectChat(ctx context.Context, req *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {

	v, err := protovalidate.New()
	if err != nil {
		return nil, models.ErrValidationFailed
	}

	if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	// TODO: получить id из jwt MW
	userId := uuid.New().String()

	form := &dto.CreateDirectChat{
		CreatorID: dto.UserID(userId),
		MemberID:  dto.UserID(req.ParticipantId),
	}

	chatResponse, err := s.ChatUsecase.CreateDirectChat(ctx, form)

	if err != nil {
		return nil, err
	}

	return &chat.CreateDirectChatResponse{ChatId: string(chatResponse.ID)}, nil
}

func (s *service) GetChat(ctx context.Context, req *chat.GetChatRequest) (*chat.GetChatResponse, error) {

	v, err := protovalidate.New()
	if err != nil {
		return nil, models.ErrValidationFailed
	}

	if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	res, err := s.ChatUsecase.FetchChat(ctx, dto.ChatID(req.ChatId))

	if err != nil {
		return nil, err
	}

	chatResponse := &chat.Chat{
		ChatId:    string(res.ID),
		CreatorId: string(res.CreatorID),
	}

	return &chat.GetChatResponse{Chat: chatResponse}, nil
}

func (s *service) ListUserChats(ctx context.Context, req *chat.ListUserChatsRequest) (*chat.ListUserChatsResponse, error) {

	v, err := protovalidate.New()
	if err != nil {
		return nil, models.ErrValidationFailed
	}

	if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	res, err := s.ChatUsecase.ListUserChats(ctx, dto.UserID(req.UserId))

	if err != nil {
		return nil, err
	}

	chatResponses := make([]*chat.ChatMember, len(res))

	for i, ch := range res {
		chatResponses[i] = &chat.ChatMember{
			UserId: string(ch.UserID),
			ChatId: string(ch.ChatID),
		}
	}

	return &chat.ListUserChatsResponse{Chats: chatResponses}, nil
}

func (s *service) ListChatMembers(ctx context.Context, req *chat.ListChatMembersRequest) (*chat.ListChatMembersResponse, error) {

	v, err := protovalidate.New()
	if err != nil {
		return nil, models.ErrValidationFailed
	}

	if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	res, err := s.ChatUsecase.ListChatMembers(ctx, dto.ChatID(req.ChatId))

	if err != nil {
		return nil, err
	}

	chatMembers := make([]*chat.ChatMember, len(res))

	for i, m := range res {
		chatMembers[i] = &chat.ChatMember{
			UserId: string(m.UserID),
			ChatId: string(m.ChatID),
		}
	}

	return &chat.ListChatMembersResponse{Members: chatMembers}, nil
}

func (s *service) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {

	v, err := protovalidate.New()
	if err != nil {
		return nil, models.ErrValidationFailed
	}

	if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	// TODO: получить id из jwt MW
	userId := uuid.New().String()

	form := &dto.SendMessage{
		Text:     req.Text,
		SenderID: dto.UserID(userId),
		ChatID:   dto.ChatID(req.ChatId),
	}

	res, err := s.ChatUsecase.SendMessage(ctx, form)

	if err != nil {
		return nil, err
	}

	chatMessage := &chat.Message{
		MessageId: string(res.ID),
		Text:      res.Text,
		SenderId:  string(res.SenderID),
		ChatId:    string(res.ChatID),
	}

	return &chat.SendMessageResponse{Message: chatMessage}, nil
}

func (s *service) ListMessages(ctx context.Context, req *chat.ListMessagesRequest) (*chat.ListMessagesResponse, error) {

	v, err := protovalidate.New()
	if err != nil {
		return nil, models.ErrValidationFailed
	}

	if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	form := &dto.ListMessages{
		ChatID: dto.ChatID(req.ChatId),
		Limit:  req.Limit,
		Cursor: req.Cursor,
	}

	messages, nextCur, err := s.ChatUsecase.ListMessages(ctx, form)

	if err != nil {
		return nil, err
	}

	chatMessages := make([]*chat.Message, len(messages))

	for i, mes := range messages {
		chatMessages[i] = &chat.Message{
			MessageId: string(mes.ID),
			Text:      mes.Text,
			SenderId:  string(mes.SenderID),
			ChatId:    string(mes.ChatID),
		}
	}

	return &chat.ListMessagesResponse{Messages: chatMessages, NextCursor: string(nextCur.NextCursor)}, nil
}

func (s *service) StreamMessages(ctx context.Context, req *chat.StreamMessagesRequest) (*chat.StreamMessagesResponse, error) {

	v, err := protovalidate.New()
	if err != nil {
		return nil, models.ErrValidationFailed
	}

	if err = v.Validate(req); err != nil {
		return nil, models.ErrValidationFailed
	}

	form := &dto.StreamMessages{
		ChatID:    dto.ChatID(req.ChatId),
		SinceUnix: req.SinceUnixMs,
	}

	res, err := s.ChatUsecase.StreamMessages(ctx, form)

	if err != nil {
		return nil, err
	}

	chatMessages := make([]*chat.Message, len(res))

	for i, mes := range res {
		chatMessages[i] = &chat.Message{
			MessageId: string(mes.ID),
			Text:      mes.Text,
			SenderId:  string(mes.SenderID),
			ChatId:    string(mes.ChatID),
		}
	}

	return &chat.StreamMessagesResponse{Stream: chatMessages}, nil
}
