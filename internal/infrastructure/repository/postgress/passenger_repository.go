package postgress

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"gorm.io/gorm"
)

type PassengerRepository struct {
	db *gorm.DB
}

func NewPassengerRepository(db *gorm.DB) *PassengerRepository {
	return &PassengerRepository{db: db}
}

func (repo *PassengerRepository) CreatePassenger(passenger *entity.Passenger) error {
	return repo.db.Create(passenger).Error
}
