package entity

import (
	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/database"
)

type Role struct {
	database.Model
	Name     string `gorm:"type:varchar(50)"`
	UserType enum.UserType

	Users []User `gorm:"many2many:role_user"`
}

func (Role) TableName() string {
	return "roles"
}
