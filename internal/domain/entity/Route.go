package entity

import (
	"github.com/Manizmn84/GoTravel/internal/infrastructure/database"
)

type Route struct {
	database.Model

	Trips []Trip `gorm:"many2many:trip_routes"`

	FromAirportID uint
	FromAirport   Airport `gorm:"foreignKey:FromAirportID"`

	ToAirportID uint
	ToAirport   Airport `gorm:"foreignKey:ToAirportID"`
}
