package entity

import (
	"github.com/Manizmn84/GoTravel/internal/infrastructure/database"
)

type Reserve struct {
	database.Model

	TotalAmount float64 `gorm:"type:decimal(10,2)"`

	PassengerID uint
	Passenger   Passenger `gorm:"foreignKey:PassengerID"`
}

func (Reserve) TableName() string {
	return "reserves"
}
