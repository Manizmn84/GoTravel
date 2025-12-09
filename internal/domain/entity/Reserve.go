package entity

import (
	"time"

	"gorm.io/gorm"
)

type Reserve struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	TotalAmount float64 `gorm:"type:decimal(10,2)"`

	PassengerID uint
	Passenger   Passenger `gorm:"foreignKey:PassengerID"`
}

func (Reserve) TableName() string {
	return "reserves"
}
