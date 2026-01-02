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

func (s *UserService) GetUserByEmail(email string) (*model.Lib_user, error) {
	if email == "" {
		return nil, errors.New("Email is required")
	}

	data, err := s.UserRepository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("user not found")
	}
	return data, nil
}
