package grpc_hd

import (
	"context"

	"msa_big_tech/chat/pkg"
)

type server struct {
	// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
	chat.UnimplementedChatServiceServer
}

func NewServer() *server {
	return &server{}
}

func (s *server) CreateDirectChat(ctx context.Context, req *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {

	return &chat.CreateDirectChatResponse{ChatId: 0}, nil
}

func (s *server) GetChat(ctx context.Context, req *chat.GetChatRequest) (*chat.GetChatResponse, error) {

	return &chat.GetChatResponse{Chat: &chat.Chat{ChatId: 0, Name: "hello_chat"}}, nil
}

func (s *server) ListUserChats(ctx context.Context, req *chat.ListUserChatsRequest) (*chat.ListUserChatsResponse, error) {

	return &chat.ListUserChatsResponse{Chats: []*chat.Chat{&chat.Chat{ChatId: 0, Name: "zero_chat"}}}, nil
}

func (s *server) ListChatMembers(ctx context.Context, req *chat.ListChatMembersRequest) (*chat.ListChatMembersResponse, error) {

	return &chat.ListChatMembersResponse{UserIds: []uint64{0, 2, 4}}, nil
}

func (s *server) SendMessage(ctx context.Context, req *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {

	return &chat.SendMessageResponse{Message: &chat.Message{MessageId: 11, Text: req.Text}}, nil
}

func (s *server) ListMessages(ctx context.Context, req *chat.ListMessagesRequest) (*chat.ListMessagesResponse, error) {

	return &chat.ListMessagesResponse{Messages: []*chat.Message{&chat.Message{MessageId: 0, Text: "zero"}}}, nil
}

func (s *server) StreamMessages(ctx context.Context, req *chat.StreamMessagesRequest) (*chat.StreamMessagesResponse, error) {

	return &chat.StreamMessagesResponse{Stream: &chat.Message{}}, nil
}