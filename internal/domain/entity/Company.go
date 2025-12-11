package entity

import (
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/database"
)

type Company struct {
	database.Model

	Name         string `gorm:"type:varchar(20)"`
	Country      string `gorm:"type:varchar(20)"`
	ContactEmail string `gorm:"type:varchar(20)"`

	Trips []Trip `gorm:"foreignKey:CompanyID"`
}
