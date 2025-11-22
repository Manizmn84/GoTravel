package entity

import (
	"time"

	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"gorm.io/gorm"
)

type Trip struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Date      time.Time      `gorm:"type:date"`
	Type      enum.TripType
}
