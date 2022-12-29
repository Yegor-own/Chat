package usecase

import "github.com/Yegor-own/Chat/internal/domain/model"

type ChatUsecase interface {
	NewChat(title string, users [2]*model.User) (*model.Chat, error)
	GetChat(id uint) (*model.Chat, error)
	UpdateChat(chat *model.Chat) (*model.Chat, error)
	DeleteChat(chat *model.Chat) error
}
