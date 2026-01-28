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

func (repo *PassengerRepository) GetPassenger(id uint) error {
	var passenger entity.Passenger
	return repo.db.First(&passenger, id).Error
}

func (r *PassengerRepository) List() ([]*entity.Passenger, error) {
	var passengers []*entity.Passenger

	err := r.db.Find(&passengers).Error
	if err != nil {
		return nil, err
	}

	return passengers, nil
}
