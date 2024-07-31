package repositories

import (
	"project/application/core/entities"
	"project/application/core/ports"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) GetUserByID(id int) (*entities.User, error) {
	var user entities.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) ListUsers(limit, offset int) ([]entities.User, error) {
	var users []entities.User
	if err := r.db.Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
