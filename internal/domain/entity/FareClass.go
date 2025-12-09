package entity

import (
	"time"

	"gorm.io/gorm"
)

type FareClass struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name             string `gorm:"type:varchar(20)"`
	BaseMultiplier   int
	BaggageAllowance int

	Seats []Seat `gorm:"foreignKey:ClassID"`
}

func (FareClass) TableName() string {
	return "fareclasses"
}
