package entity

import (
	"time"

	"gorm.io/gorm"
)

type Company struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name         string `gorm:"type:varchar(20)"`
	Country      string `gorm:"type:varchar(20)"`
	ContactEmail string `gorm:"type:varchar(20)"`

	Trips []Trip `gorm:"foreignKey:CompanyID"`
}
