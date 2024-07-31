package repositories

import (
	"project/application/core/entities"
	"project/application/core/ports"

	"gorm.io/gorm"
)

type ChallengeRepository struct {
	db *gorm.DB
}

func NewChallengeRepository(db *gorm.DB) ports.ChallengeRepository {
	return &ChallengeRepository{db}
}

func (r *ChallengeRepository) CreateChallenge(challenge *entities.Challenge) error {
	return r.db.Create(challenge).Error
}

func (r *ChallengeRepository) GetChallengeByID(id int) (*entities.Challenge, error) {
	var challenge entities.Challenge
	if err := r.db.First(&challenge, id).Error; err != nil {
		return nil, err
	}
	return &challenge, nil
}

func (r *ChallengeRepository) ListChallenges(limit, offset int) ([]entities.Challenge, error) {
	var challenges []entities.Challenge
	if err := r.db.Limit(limit).Offset(offset).Find(&challenges).Error; err != nil {
		return nil, err
	}
	return challenges, nil
}
