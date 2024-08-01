package ports

import (
	"project/application/core/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetUserByID(id int) (*entities.User, error)
	ListUsers(limit, offset int) ([]entities.User, error)
	DeleteUser(id int) error
}

type ChallengeRepository interface {
	CreateChallenge(challenge *entities.Challenge) error
	GetChallengeByID(id int) (*entities.Challenge, error)
	ListChallenges(limit, offset int) ([]entities.Challenge, error)
	DeleteChallenge(id int) error
}

type VideoRepository interface {
	CreateVideo(video *entities.Video) error
	GetVideoByID(id int) (*entities.Video, error)
	ListVideos(limit, offset int) ([]entities.Video, error)
	DeleteVideo(id int) error
}
