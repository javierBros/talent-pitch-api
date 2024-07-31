package entities

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255;unique"`
	ImagePath string `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
