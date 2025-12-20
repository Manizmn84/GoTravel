package postgress

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"gorm.io/gorm"
)

type TripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) *TripRepository {
	return &TripRepository{db: db}
}

func (tripRepo *TripRepository) CreateTrip(trip *entity.Trip) error {
	return tripRepo.db.Create(trip).Error
}
