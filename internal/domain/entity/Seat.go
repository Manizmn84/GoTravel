package entity

import (
	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/database"
)

type Seat struct {
	database.Model

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
