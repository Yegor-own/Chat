package service

import (
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/repository"
)

type MessageService interface {
	InsertMessage(text string, cId uint) (*entities.Message, error)
	GetMessage(id uint) (*entities.Message, error)
	GetAllMessages(chatId uint) ([]*entities.Message, error)
}

type messageService struct {
	repository repository.MessageRepository
}

func NewMessageService(r repository.MessageRepository) MessageService {
	return &messageService{
		repository: r,
	}
}

func (s *messageService) InsertMessage(text string, cId uint) (*entities.Message, error) {
	return s.repository.CreateMessage(text, cId)
}

func (s *messageService) GetMessage(id uint) (*entities.Message, error) {
	return s.repository.ReadMessage(id)
}

func (s *messageService) GetAllMessages(chatId uint) ([]*entities.Message, error) {
	return s.repository.ReadAllMessages(chatId)
}
