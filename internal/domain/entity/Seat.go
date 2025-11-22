package entity

import (
	"time"

	"gorm.io/gorm"
)

type Seat struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Capacity  int64
	Filled    int64
}
