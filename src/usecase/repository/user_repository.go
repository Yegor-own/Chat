package repository

import "github.com/Yegor-own/Chat/domain/model"

type UserRepository interface {
	FinAll(u []*model.User) ([]*model.User, error)
}
