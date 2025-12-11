package entity

import (
	"github.com/Manizmn84/GoTravel/internal/infrastructure/database"
)

type User struct {
	database.Model
	Email    string `gorm:"type:varchar(50);uniqueIndex"`
	Phone    string `gorm:"type:varchar(13)"`
	Password string `gorm:"type:varchar(255)"`
	Username string `gorm:"type:varchar(20)"`
	Name     string `gorm:"type:varchar(20)"`

	Roles []Role `gorm:"many2many:role_user"`
}

func (User) TableName() string {
	return "users"
}
