package usecase

import (
	model2 "github.com/Yegor-own/Chat/src/v0/internal/domain/model"
)

type ChatUsecase interface {
	NewChat(title string, users [2]*model2.User) (*model2.Chat, error)
	GetChat(id uint) (*model2.Chat, error)
	UpdateChat(chat *model2.Chat) (*model2.Chat, error)
	DeleteChat(chat *model2.Chat) error
}
