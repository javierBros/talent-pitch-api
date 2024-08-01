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

// NewChallengeService crea una nueva instancia de ChallengeService.
func NewChallengeService(challengeRepo ports.ChallengeRepository, userRepo ports.UserRepository) *ChallengeService {
	return &ChallengeService{challengeRepo, userRepo}
}

// CreateChallenge crea un nuevo desafío.
func (s *ChallengeService) CreateChallenge(challenge *entities.Challenge) error {
	if _, err := s.userRepo.GetUserByID(challenge.UserID); err != nil {
		return errors.New("user not found")
	}

	return s.challengeRepo.CreateChallenge(challenge)
}

// GetChallengeByID obtiene un desafío por ID.
func (s *ChallengeService) GetChallengeByID(id int) (*entities.Challenge, error) {
	return s.challengeRepo.GetChallengeByID(id)
}

// ListChallenges lista los desafíos con paginación.
func (s *ChallengeService) ListChallenges(limit, offset int) ([]entities.Challenge, error) {
	return s.challengeRepo.ListChallenges(limit, offset)
}

// DeleteChallenge elimina un desafío por ID.
func (s *ChallengeService) DeleteChallenge(id int) error {
	_, err := s.challengeRepo.GetChallengeByID(id)
	if err != nil {
		return err
	}
	return s.challengeRepo.DeleteChallenge(id)
}
