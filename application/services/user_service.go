package services

import (
	"project/application/core/entities"
	"project/application/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(user *entities.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id int) (*entities.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) ListUsers(limit, offset int) ([]entities.User, error) {
	return s.repo.ListUsers(limit, offset)
}
