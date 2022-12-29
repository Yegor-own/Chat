package usecase

import "github.com/Yegor-own/Chat/internal/domain/model"

type MessageUsecase interface {
	NewMessage(text string, authorId uint, chatId uint) (*model.Message, error)
}
