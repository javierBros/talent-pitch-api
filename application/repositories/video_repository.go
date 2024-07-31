package repositories

import (
	"project/application/core/entities"
	"project/application/core/ports"

	"gorm.io/gorm"
)

type VideoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) ports.VideoRepository {
	return &VideoRepository{db}
}

func (r *VideoRepository) CreateVideo(video *entities.Video) error {
	return r.db.Create(video).Error
}

func (r *VideoRepository) GetVideoByID(id int) (*entities.Video, error) {
	var video entities.Video
	if err := r.db.First(&video, id).Error; err != nil {
		return nil, err
	}
	return &video, nil
}

func (r *VideoRepository) ListVideos(limit, offset int) ([]entities.Video, error) {
	var videos []entities.Video
	if err := r.db.Limit(limit).Offset(offset).Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
