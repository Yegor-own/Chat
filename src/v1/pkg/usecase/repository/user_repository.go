package repository

import (
	"errors"
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"gorm.io/gorm"
	"reflect"
)

type UserRepository interface {
	CreateUser(name, password string) (*entities.User, error)
	ReadUser(id uint) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(id uint) error
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
	if res.Error != nil {
		return nil, errors.New("failed to create user")
	}
	return &user, nil
}

func (r *repository) ReadUser(id uint) (*entities.User, error) {
	user := entities.User{ID: id}
	r.db.First(&user)
	if reflect.ValueOf(user.Name).IsZero() {
		return nil, errors.New("failed to read user")
	}
	return &user, nil
}
func (r *repository) UpdateUser(user *entities.User) (*entities.User, error) {
	r.db.Save(user)
	return user, nil
}
func (r *repository) DeleteUser(id uint) error {
	user := entities.User{ID: id}
	r.db.Delete(&user)
	return nil
}
