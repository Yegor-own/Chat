package usecase

import "github.com/Yegor-own/Chat/internal/domain/model"

type UserUsecase interface {
	NewUser(name, password string) (*model.User, error)
	GetUser(name string) (*model.User, error)
	UpdateUser(user *model.User) (*model.User, error)
	DeleteUser(user *model.User) error
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
