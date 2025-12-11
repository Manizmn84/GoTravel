package entity

import (
	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/database"
)

type Payment struct {
	database.Model

	Status    enum.PaymentStatus `gorm:"type:int"`
	Method    enum.PaymentMethod `gorm:"type:int"`
	ReserveID uint               `gorm:"type:uniqueIndex"`
	Reserve   Reserve            `gorm:"foreignKey:ReserveID"`
}

func (Payment) TableName() string {
	return "payments"
}
