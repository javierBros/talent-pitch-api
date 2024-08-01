package services

import (
	"project/application/core/entities"
	"project/application/core/ports"
)

type UserService struct {
	repo ports.UserRepository
}

// NewUserService crea una nueva instancia de UserService.
func NewUserService(repo ports.UserRepository) *UserService {
	return &UserService{repo}
}

// CreateUser crea un nuevo usuario.
func (s *UserService) CreateUser(user *entities.User) error {
	return s.repo.CreateUser(user)
}

// GetUserByID obtiene un usuario por ID.
func (s *UserService) GetUserByID(id int) (*entities.User, error) {
	return s.repo.GetUserByID(id)
}

// ListUsers lista los usuarios con paginación.
func (s *UserService) ListUsers(limit, offset int) ([]entities.User, error) {
	return s.repo.ListUsers(limit, offset)
}

// DeleteUser elimina un usuario por ID.
func (s *UserService) DeleteUser(id int) error {
	_, err := s.repo.GetUserByID(id)
	if err != nil {
		return err
	}
	return s.repo.DeleteUser(id)
}
