package repository

import (
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"gorm.io/gorm"
	"time"
)

type MessageRepository interface {
	CreateMessage(text string, cId uint) (*entities.Message, error)
	ReadMessage(id uint) (*entities.Message, error)
	ReadAllMessages(chatId uint) ([]*entities.Message, error)
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &messageRepository{
		db: db,
	}
}

func (r *messageRepository) CreateMessage(text string, cId uint) (*entities.Message, error) {
	msg := entities.Message{
		Text:      text,
		CreatedAt: time.Now(),
		ChatID:    cId,
	}

	res := r.db.Create(&msg)
	if res.Error != nil {
		return nil, res.Error
	}

	return &msg, nil
}

func (r *messageRepository) ReadMessage(id uint) (*entities.Message, error) {
	msg := entities.Message{ID: id}

	res := r.db.First(&msg)
	if res.Error != nil {
		return nil, res.Error
	}

	return &msg, nil
}

func (r *messageRepository) ReadAllMessages(chatId uint) ([]*entities.Message, error) {
	msgs := []*entities.Message{}

	res := r.db.Find(msgs, "chat_id = ?", chatId)
	if res.Error != nil {
		return nil, res.Error
	}

	return msgs, nil
}
