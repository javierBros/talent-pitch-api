package services

import (
	"errors"
	"project/application/core/entities"
	"project/application/core/ports"
)

type ChallengeService struct {
	challengeRepo ports.ChallengeRepository
	userRepo      ports.UserRepository
}

func NewChallengeService(challengeRepo ports.ChallengeRepository, userRepo ports.UserRepository) *ChallengeService {
	return &ChallengeService{challengeRepo, userRepo}
}

func (s *ChallengeService) CreateChallenge(challenge *entities.Challenge) error {
	if _, err := s.userRepo.GetUserByID(challenge.UserID); err != nil {
		return errors.New("user not found")
	}

	return s.challengeRepo.CreateChallenge(challenge)
}

func (s *ChallengeService) GetChallengeByID(id int) (*entities.Challenge, error) {
	return s.challengeRepo.GetChallengeByID(id)
}

func (s *ChallengeService) ListChallenges(limit, offset int) ([]entities.Challenge, error) {
	return s.challengeRepo.ListChallenges(limit, offset)
}
