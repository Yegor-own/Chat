package repository

import (
	"errors"
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(name, password string) (*entities.User, error)
	ReadUser()
	UpdateUser()
	DeleteUser()
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(name, password string) (*entities.User, error) {
	user := entities.User{
		Name: name, Password: password,
	}
	res := r.db.Create(&user)
	if res.Error != nil || user.ID <= 0 {
		return nil, errors.New("failed to create user")
	}
	return &user, nil
}

func (r *repository) ReadUser() {
}
func (r *repository) UpdateUser() {
}
func (r *repository) DeleteUser() {
}
