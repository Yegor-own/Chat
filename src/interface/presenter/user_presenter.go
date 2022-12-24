package presenter

import "github.com/Yegor-own/Chat/domain/model"

type userPresenter struct{}

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}

func NewUserPresenter() UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(u []*model.User) []*model.User {
	for _, user := range u {
		user.Name = "Mr." + user.Name
	}
	return u
}
