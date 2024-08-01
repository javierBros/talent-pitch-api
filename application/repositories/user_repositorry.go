package repositories

import (
	"project/application/core/entities"
	"project/application/core/ports"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository crea una nueva instancia de UserRepository.
func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &UserRepository{db}
}

// CreateUser crea un nuevo usuario en la base de datos.
func (r *UserRepository) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

// GetUserByID obtiene un usuario por ID de la base de datos.
func (r *UserRepository) GetUserByID(id int) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ListUsers lista los usuarios con paginación.
func (r *UserRepository) ListUsers(limit, offset int) ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// DeleteUser elimina un usuario por ID de la base de datos.
func (r *UserRepository) DeleteUser(id int) error {
	return r.db.Delete(&entities.User{}, id).Error
}
