package service

import (
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/repository"
)

type ChatService interface {
	InsertChat(title string) (*entities.Chat, error)
	GetChat(id uint) (*entities.Chat, error)
	UpdateChat(id uint, title string) (*entities.Chat, error)
	RemoveChat(id uint) error
}

type chatService struct {
	repository repository.ChatRepository
}

func NewChatService(r repository.ChatRepository) ChatService {
	return &chatService{
		repository: r,
	}
}

func (s *chatService) InsertChat(title string) (*entities.Chat, error) {
	return s.repository.CreateChat(title)
}

func (s *chatService) GetChat(id uint) (*entities.Chat, error) {
	return s.repository.ReadChat(id)
}

func (s *chatService) UpdateChat(id uint, title string) (*entities.Chat, error) {
	return s.repository.UpdateChat(id, title)
}

func (s *chatService) RemoveChat(id uint) error {
	return s.repository.DeleteChat(id)
}
