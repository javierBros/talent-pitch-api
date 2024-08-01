package mappers

import (
	"github.com/talent-pitch-api/application/core/domain"
	"github.com/talent-pitch-api/application/core/entities"
)

func ToChallengeResponse(challenge *entities.Challenge) domain.ChallengeResponse {
	return domain.ChallengeResponse{
		ID:          challenge.ID,
		Title:       challenge.Title,
		Description: challenge.Description,
		Difficulty:  challenge.Difficulty,
		UserID:      challenge.UserID,
	}
}

func ToChallengeResponses(challenges []entities.Challenge) []domain.ChallengeResponse {
	challengeResponses := make([]domain.ChallengeResponse, len(challenges))
	for i, challenge := range challenges {
		challengeResponses[i] = ToChallengeResponse(&challenge)
	}
	return challengeResponses
}

func ToChallengeEntity(req *domain.CreateChallengeRequest) *entities.Challenge {
	return &entities.Challenge{
		Title:       req.Title,
		Description: req.Description,
		Difficulty:  req.Difficulty,
		UserID:      req.UserID,
	}
}
