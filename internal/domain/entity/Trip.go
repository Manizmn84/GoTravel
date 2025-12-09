package entity

import (
	"time"

	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"gorm.io/gorm"
)

type Trip struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Date      time.Time      `gorm:"type:date"`

	CompanyID int
	Company   Company `gorm:"foreignKey:CompanyID"`

	Seats []Seat `gorm:"foreignKey:TripID"`

	BaseFare float64 `gorm:"type:decimal(10,2)"`

	AircraftModel enum.AircraftModel

	Routes []Route `gorm:"many2many:trip_routes"`
}
