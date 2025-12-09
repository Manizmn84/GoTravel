package entity

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Email     string         `gorm:"type:varchar(50);uniqueIndex"`
	Phone     string         `gorm:"type:varchar(13)"`
	Password  string         `gorm:"type:varchar(20)"`
	Username  string         `gorm:"type:varchar(20)"`
	Name      string         `gorm:"type:varchar(20)"`
}
