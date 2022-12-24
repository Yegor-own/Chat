package presenter

import "github.com/Yegor-own/Chat/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
