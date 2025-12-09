package entity

import (
	"time"

	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"gorm.io/gorm"
)

type Passenger struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Email        string `gorm:"type:varchar(100);uniqueIndex;not null"`
	FirstName    string `gorm:"type:varchar(20)"`
	LastName     string `gorm:"type:varchar(20)"`
	Gender       enum.Gender
	Password     string `gorm:"type:varchar(255)"`
	NationalCode string `gorm:"type:varchar(20)"`
	Phone        string `gorm:"type:varchar(20)"`

	Reserves []Reserve `gorm:"foreignKey:PassengerId"`
}

func (Passenger) TableName() string {
	return "passengers"
}
