package mappers

import (
	"project/application/core/domain"
	"project/application/core/entities"
)

func ToUserResponse(user *entities.User) domain.UserResponse {
	return domain.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		ImagePath: user.ImagePath,
	}
}

func ToUserResponses(users []entities.User) []domain.UserResponse {
	userResponses := make([]domain.UserResponse, len(users))
	for i, user := range users {
		userResponses[i] = ToUserResponse(&user)
	}
	return userResponses
}

func ToUserEntity(req *domain.CreateUserRequest) *entities.User {
	return &entities.User{
		Name:      req.Name,
		Email:     req.Email,
		ImagePath: req.ImagePath,
	}
}
