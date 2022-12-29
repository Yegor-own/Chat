package controller

import (
	"github.com/Yegor-own/Chat/internal/domain/model"
	"github.com/Yegor-own/Chat/internal/usecase"
)

type UserController struct {
	usecase.UserUsecase
}

func (uc *UserController) NewUser(name, password string) (*model.User, error) {
	return &model.User{
		Name:     name,
		Password: password,
	}, nil
}
