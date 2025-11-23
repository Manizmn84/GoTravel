package entity

import (
	"time"

	"gorm.io/gorm"
)

type Route struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Trips []Trip

	FromAirportId int
	FromAirport   Airport `gorm:"foreignKey:fromAirportId"`

	ToAirportId int
	ToAirport   Airport `gorm:"foreignKey:toAirport"`
}
