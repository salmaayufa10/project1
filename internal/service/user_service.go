package service

import (
	"errors"
	"library/internal/model"
	"library/internal/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(UserRepository *repository.UserRepository) *UserService {
	return &UserService{
		UserRepository: UserRepository,
	}
}

func (s *UserService) CreateUser(user *model.Lib_user) error {
	if user.Email == "" || user.Name == "" || user.Password == "" {
		return errors.New("Email, Name, Password are required")
	}

	err := s.UserRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
