package repositories

import (
	"project/application/core/entities"
	"project/application/core/ports"

	"gorm.io/gorm"
)

type ChallengeRepository struct {
	db *gorm.DB
}

// NewChallengeRepository crea una nueva instancia de ChallengeRepository.
func NewChallengeRepository(db *gorm.DB) ports.ChallengeRepository {
	return &ChallengeRepository{db}
}

// CreateChallenge crea un nuevo desafío en la base de datos.
func (r *ChallengeRepository) CreateChallenge(challenge *entities.Challenge) error {
	return r.db.Create(challenge).Error
}

// GetChallengeByID obtiene un desafío por ID de la base de datos.
func (r *ChallengeRepository) GetChallengeByID(id int) (*entities.Challenge, error) {
	var challenge entities.Challenge
	if err := r.db.First(&challenge, id).Error; err != nil {
		return nil, err
	}
	return &challenge, nil
}

// ListChallenges lista los desafíos con paginación.
func (r *ChallengeRepository) ListChallenges(limit, offset int) ([]entities.Challenge, error) {
	var challenges []entities.Challenge
	if err := r.db.Limit(limit).Offset(offset).Find(&challenges).Error; err != nil {
		return nil, err
	}
	return challenges, nil
}

// DeleteChallenge elimina un desafío por ID de la base de datos.
func (r *ChallengeRepository) DeleteChallenge(id int) error {
	return r.db.Delete(&entities.Challenge{}, id).Error
}
