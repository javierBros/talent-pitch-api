package ports

import (
	"project/application/core/entities"
)

type UserService interface {
	CreateUser(user *entities.User) error
	GetUserByID(id int) (*entities.User, error)
	ListUsers(limit, offset int) ([]entities.User, error)
}

type ChallengeService interface {
	CreateChallenge(challenge *entities.Challenge) error
	GetChallengeByID(id int) (*entities.Challenge, error)
	ListChallenges(limit, offset int) ([]entities.Challenge, error)
}

type VideoService interface {
	CreateVideo(video *entities.Video) error
	GetVideoByID(id int) (*entities.Video, error)
	ListVideos(limit, offset int) ([]entities.Video, error)
}
