package entity

import (
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/database"
)

type Admin struct {
	database.Model
	Email    string `gorm:"type:varchar(50);uniqueIndex"`
	Phone    string `gorm:"type:varchar(13)"`
	Password string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(20)"`
	Name     string `gorm:"type:varchar(20)"`
}

func (Admin) TableName() string {
	return "admins"
}
