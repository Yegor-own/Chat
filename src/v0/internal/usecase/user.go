package usecase

import (
	"github.com/Yegor-own/Chat/src/v0/internal/domain/model"
	"github.com/Yegor-own/Chat/src/v0/internal/repository"
)

type UserUsecase interface {
	NewUser(name, password string) (*model.User, error)
	GetUser(name string) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	RemoveUser(user *model.User) error
}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return userUsecase{
		repository: repo,
	}
}

//func NewUser(name, password string) *model.User {
//	return &model.User{
//		Name:     name,
//		Password: password,
//	}
//}
//
//func GetUser(id int) {
//
//}
