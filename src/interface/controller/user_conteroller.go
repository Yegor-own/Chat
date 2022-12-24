package controller

import (
	"github.com/Yegor-own/Chat/domain/model"
	"github.com/Yegor-own/Chat/usecase/interactor"
	"net/http"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	GetUsers(c Context) error
}

func NewUserController(us interactor.UserInteractor) UserController {
	return &userController{us}
}

func (us *userController) GetUsers(c Context) error {
	var u []*model.User
	u, err := us.userInteractor.Get(u)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, u)
}
