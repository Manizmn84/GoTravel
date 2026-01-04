package entity

import (
	"github.com/Manizmn84/GoTravel/internal/infrastructure/database"
)

type Route struct {
	database.Model

	Trips []Trip `gorm:"many2many:trip_routes"`

	Distance int64
	BaseTime int

	OriginID      uint
	OriginAirport Airport `gorm:"foreignKey:OriginID"`

	DestinationID      uint
	DestinationAirport Airport `gorm:"foreignKey:DestinationID"`
}
