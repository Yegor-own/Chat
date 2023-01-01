package service

import "github.com/Yegor-own/Chat/src/v1/pkg/usecase/repository"

type UserService interface {
	InsertUser()
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{
		repository: r,
	}
}

func (s *userService) InsertUser() {

}