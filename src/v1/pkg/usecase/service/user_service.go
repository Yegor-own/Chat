package service

import (
	"github.com/Yegor-own/Chat/src/v1/pkg/entities"
	"github.com/Yegor-own/Chat/src/v1/pkg/usecase/repository"
)

type UserService interface {
	InsertUser(name, password string) (*entities.User, error)
	GetUser(id uint) (*entities.User, error)
	GetUserByName(name string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	RemoveUser(id uint) error
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repository: r,
	}
}

func (s *userService) InsertUser(name, password string) (*entities.User, error) {
	return s.repository.CreateUser(name, password)
}

func (s *userService) GetUser(id uint) (*entities.User, error) {
	return s.repository.ReadUser(id)
}

func (s *userService) GetUserByName(name string) (*entities.User, error) {
	return s.repository.ReadUserByName(name)
}

func (s *userService) GetUserByEmail(email string) (*entities.User, error) {
	return s.repository.ReadUserByEmail(email)
}

func (s *userService) UpdateUser(user *entities.User) (*entities.User, error) {
	return s.repository.UpdateUser(user)
}

func (s *userService) RemoveUser(id uint) error {
	return s.repository.DeleteUser(id)
}
