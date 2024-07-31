package entities

import (
	"time"
)

type Video struct {
	ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"size:255"`
	Description string `gorm:"type:text"`
	URL         string `gorm:"size:255"`
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
