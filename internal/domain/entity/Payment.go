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
	Amount    int64
	Status    enum.PaymentStatus
}
