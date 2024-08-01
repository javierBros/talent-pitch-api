package repositories

import (
	"github.com/talent-pitch-api/application/core/entities"
	"github.com/talent-pitch-api/application/core/ports"

	"gorm.io/gorm"
)

type VideoRepository struct {
	db *gorm.DB
}

// NewVideoRepository crea una nueva instancia de VideoRepository.
func NewVideoRepository(db *gorm.DB) ports.VideoRepository {
	return &VideoRepository{db}
}

// CreateVideo crea un nuevo video en la base de datos.
func (r *VideoRepository) CreateVideo(video *entities.Video) error {
	return r.db.Create(video).Error
}

// GetVideoByID obtiene un video por ID de la base de datos.
func (r *VideoRepository) GetVideoByID(id int) (*entities.Video, error) {
	var video entities.Video
	if err := r.db.First(&video, id).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

// ListVideos lista los videos con paginación.
func (r *VideoRepository) ListVideos(limit, offset int) ([]entities.Video, error) {
	var videos []entities.Video
	if err := r.db.Limit(limit).Offset(offset).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

// DeleteVideo elimina un video por ID de la base de datos.
func (r *VideoRepository) DeleteVideo(id int) error {
	return r.db.Delete(&entities.Video{}, id).Error
}
