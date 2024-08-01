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

// NewVideoService crea una nueva instancia de VideoService.
func NewVideoService(videoRepo ports.VideoRepository, userRepo ports.UserRepository) *VideoService {
	return &VideoService{videoRepo, userRepo}
}

// CreateVideo crea un nuevo video.
func (s *VideoService) CreateVideo(video *entities.Video) error {
	if _, err := s.userRepo.GetUserByID(video.UserID); err != nil {
		return errors.New("user not found")
	}

	return s.videoRepo.CreateVideo(video)
}

// GetVideoByID obtiene un video por ID.
func (s *VideoService) GetVideoByID(id int) (*entities.Video, error) {
	return s.videoRepo.GetVideoByID(id)
}

// ListVideos lista los videos con paginación.
func (s *VideoService) ListVideos(limit, offset int) ([]entities.Video, error) {
	return s.videoRepo.ListVideos(limit, offset)
}

// DeleteVideo elimina un video por ID.
func (s *VideoService) DeleteVideo(id int) error {
	_, err := s.videoRepo.GetVideoByID(id)
	if err != nil {
		return err
	}
	return s.videoRepo.DeleteVideo(id)
}
