package entity

import (
	"time"

	"gorm.io/gorm"
)

type Airline struct {
	ID           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Name         string
	Country      string
	SupportEmail string
}
