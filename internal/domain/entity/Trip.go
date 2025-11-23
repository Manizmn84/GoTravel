package entity

import (
	"time"

	"gorm.io/gorm"
)

type Trip struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Date      time.Time      `gorm:"type:date"`
	Capacity  int
	Filled    int

	CompanyId int
	Company   Company `gorm:"foreignKey:CompanyId"`

	Seats []Seat

	RouteId int
	Route   Route `gorm:"foreignKey:RouteId"`
}
