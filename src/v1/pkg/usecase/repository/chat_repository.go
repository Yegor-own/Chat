package repository

import (
	"errors"
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"gorm.io/gorm"
	"reflect"
)

type ChatRepository interface {
	CreateChat(title string) (*entities.Chat, error)
	ReadChat(id uint) (*entities.Chat, error)
	UpdateChat(id uint, title string) (*entities.Chat, error)
	DeleteChat(id uint) error
}

func NewChatRepository(db *gorm.DB) ChatRepository {
	return &chatRepository{
		db: db,
	}
}

func (r *chatRepository) CreateChat(title string) (*entities.Chat, error) {
	chat := entities.Chat{Title: title}
	res := r.db.Create(&chat)

	if res.Error != nil {
		return nil, res.Error
	}

	if reflect.ValueOf(chat.ID).IsZero() {
		return nil, errors.New("failed to create chat")
	}

	return &chat, nil
}

func (r *chatRepository) ReadChat(id uint) (*entities.Chat, error) {
	chat := entities.Chat{ID: id}
	res := r.db.First(&chat)

	if res.Error != nil {
		return nil, res.Error
	}

	if reflect.ValueOf(chat.Title).IsZero() {
		return nil, errors.New("failed to find chat")
	}

	return &chat, nil
}

func (r *chatRepository) UpdateChat(id uint, title string) (*entities.Chat, error) {
	chat := entities.Chat{
		ID:    id,
		Title: title,
	}
	res := r.db.Save(&chat)

	if res.Error != nil {
		return nil, res.Error
	}

	return &chat, nil
}
func (r *chatRepository) DeleteChat(id uint) error {
	chat := entities.Chat{ID: id}
	res := r.db.Delete(&chat)

	return res.Error
}
