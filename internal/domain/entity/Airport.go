package entity

import "github.com/Manizmn84/GoTravel/internal/infrastructure/database"

type Airport struct {
	database.Model

	Name      string `gorm:"type:varchar(20)"`
	Country   string `gorm:"type:varchar(20)"`
	City      string `gorm:"type:varchar(20)"`
	LATA_Code string `gorm:"type:varchar(20)"`

	FromRoutes []Route `gorm:"foreignKey:OriginID"`
	ToRoutes   []Route `gorm:"foreignKey:DestinationID"`
}
