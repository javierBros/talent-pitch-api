package mappers

import (
	"github.com/talent-pitch-api/application/core/domain"
	"github.com/talent-pitch-api/application/core/entities"
)

func ToVideoResponse(video *entities.Video) domain.VideoResponse {
	return domain.VideoResponse{
		ID:          video.ID,
		Title:       video.Title,
		Description: video.Description,
		URL:         video.URL,
		UserID:      video.UserID,
	}
}

func ToVideoResponses(videos []entities.Video) []domain.VideoResponse {
	videoResponses := make([]domain.VideoResponse, len(videos))
	for i, video := range videos {
		videoResponses[i] = ToVideoResponse(&video)
	}
	return videoResponses
}

func ToVideoEntity(req *domain.CreateVideoRequest) *entities.Video {
	return &entities.Video{
		Title:       req.Title,
		Description: req.Description,
		URL:         req.URL,
		UserID:      req.UserID,
	}
}
