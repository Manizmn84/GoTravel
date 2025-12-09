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

	Trips []Trip `gorm:"many2many:trip_routes"`

	FromAirportID uint
	FromAirport   Airport `gorm:"foreignKey:FromAirportID"`

	ToAirportID uint
	ToAirport   Airport `gorm:"foreignKey:ToAirportID"`
}
