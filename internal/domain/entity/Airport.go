package entity

import (
	"time"

	"gorm.io/gorm"
)

type Airport struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name      string `gorm:"type:varchar(20)"`
	Country   string `gorm:"type:varchar(20)"`
	City      string `gorm:"type:varchar(20)"`
	LATA_Code string `gorm:"type:varchar(20)"`

	FromRoutes []Route `gorm:"foreignKey:FromAirportID"`
	ToRoutes   []Route `gorm:"foreignKey:ToAirportID"`
}
