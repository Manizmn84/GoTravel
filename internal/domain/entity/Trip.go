package entity

import (
	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/database"
)

type Trip struct {
	database.Model

	CompanyID int
	Company   Company `gorm:"foreignKey:CompanyID"`

	Seats []Seat `gorm:"foreignKey:TripID"`

	BaseFare float64 `gorm:"type:decimal(10,2)"`

	AircraftModel enum.AircraftModel

	Routes []Route `gorm:"many2many:trip_routes"`
}
