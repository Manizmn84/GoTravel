package entity

import (
	"time"

	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"gorm.io/gorm"
)

type Seat struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Type   enum.SeatType
	Number int

	TripID uint `gorm:"unique"`
	Trip   Trip `gorm:"foreignKey:TripID"`

	ClassID   uint      `gorm:"unique"`
	FareClass FareClass `gorm:"foreignKey:ClassID"`
}

func (Seat) TableName() string {
	return "seats"
}
