package chat_grpc

import (
	"context"

	"github.com/darialissi/msa_big_tech/chat/pkg"
)

func (s *service) CreateDirectChat(ctx context.Context, req *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {

	return &chat.CreateDirectChatResponse{ChatId: 0}, nil
}

func (s *service) GetChat(ctx context.Context, req *chat.GetChatRequest) (*chat.GetChatResponse, error) {

	return &chat.GetChatResponse{Chat: &chat.Chat{}}, nil
}

func (s *service) ListUserChats(ctx context.Context, req *chat.ListUserChatsRequest) (*chat.ListUserChatsResponse, error) {

	return &chat.ListUserChatsResponse{Chats: []*chat.Chat{&chat.Chat{}}}, nil
}

func (s *service) ListChatMembers(ctx context.Context, req *chat.ListChatMembersRequest) (*chat.ListChatMembersResponse, error) {

	return &chat.ListChatMembersResponse{UserIds: []uint64{0, 2, 4}}, nil
}

func (s *service) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {

	return &chat.SendMessageResponse{Message: &chat.Message{}}, nil
}

func (s *service) ListMessages(ctx context.Context, req *chat.ListMessagesRequest) (*chat.ListMessagesResponse, error) {

	return &chat.ListMessagesResponse{Messages: []*chat.Message{&chat.Message{}}}, nil
}

func (s *service) StreamMessages(ctx context.Context, req *chat.StreamMessagesRequest) (*chat.StreamMessagesResponse, error) {

	return &chat.StreamMessagesResponse{Stream: &chat.Message{}}, nil
}