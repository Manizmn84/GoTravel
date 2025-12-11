package entity

import (
	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/database"
)

type Passenger struct {
	database.Model

	Email        string `gorm:"type:varchar(100);uniqueIndex;not null"`
	FirstName    string `gorm:"type:varchar(20)"`
	LastName     string `gorm:"type:varchar(20)"`
	Gender       enum.Gender
	Password     string `gorm:"type:varchar(255)"`
	NationalCode string `gorm:"type:varchar(20)"`
	Phone        string `gorm:"type:varchar(20)"`

	Reserves []Reserve `gorm:"foreignKey:PassengerID"`
}

func (Passenger) TableName() string {
	return "passengers"
}
