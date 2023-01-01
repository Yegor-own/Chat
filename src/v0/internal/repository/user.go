package repository

import (
	"github.com/Yegor-own/Chat/src/v0/internal/domain/model"
)

type UserRepository struct{}

func (uc UserRepository) NewUser(name, password string) (*model.User, error) {
	var user = &model.User{
		Name:     name,
		Password: password,
	}

	res := DB.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (uc UserRepository) GetUser(name string) (*model.User, error) {
	var user *model.User
	user.Name = name

	res := DB.First(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (uc UserRepository) UpdateUser(user *model.User) (*model.User, error) {
	res := DB.Save(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

func (uc UserRepository) RemoveUser(user *model.User) error {
	return nil
}
