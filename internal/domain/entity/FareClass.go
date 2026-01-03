package entity

import (
	"github.com/Manizmn84/GoTravel/internal/infrastructure/database"
)

type FareClass struct {
	database.Model

	Name             string `gorm:"type:varchar(20)"`
	BaseMultiplier   int
	BaggageAllowance int

	Seats []Seat `gorm:"foreignKey:ClassID"`
}

func (FareClass) TableName() string {
	return "fareclasses"
}
