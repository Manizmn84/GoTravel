package entity

import (
	"time"

	"gorm.io/gorm"
)

type Passenger struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Password  string
	Username  string
	email     string `gorm:"email"`
}
