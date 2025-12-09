package entity

import (
	"time"

	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"gorm.io/gorm"
)

type Payment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Status    enum.PaymentStatus `gorm:"type:int"`
	Method    enum.PaymentMethod `gorm:"type:int"`
	ReserveID uint               `gorm:"type:uniqueIndex"`
	Reserve   Reserve            `gorm:"foreignKey:ReserveID"`
}

func (Payment) TableName() string {
	return "payments"
}
