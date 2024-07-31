package entities

import (
	"time"
)

type Challenge struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"size:255"`
	Description string `gorm:"type:text"`
	Difficulty  int
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
