package services

import (
	"errors"
	"project/application/core/entities"
	"project/application/core/ports"
)

type VideoService struct {
	videoRepo ports.VideoRepository
	userRepo  ports.UserRepository
}

func NewVideoService(videoRepo ports.VideoRepository, userRepo ports.UserRepository) *VideoService {
	return &VideoService{videoRepo, userRepo}
}

func (s *VideoService) CreateVideo(video *entities.Video) error {
	if _, err := s.userRepo.GetUserByID(video.UserID); err != nil {
		return errors.New("user not found")
	}

	return s.videoRepo.CreateVideo(video)
}

func (s *VideoService) GetVideoByID(id int) (*entities.Video, error) {
	return s.videoRepo.GetVideoByID(id)
}

func (s *VideoService) ListVideos(limit, offset int) ([]entities.Video, error) {
	return s.videoRepo.ListVideos(limit, offset)
}
