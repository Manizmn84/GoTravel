package postgress

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"gorm.io/gorm"
)

type AirportRepository struct {
	db *gorm.DB
}

func NewAirportRepository(db *gorm.DB) *AirportRepository {
	return &AirportRepository{db: db}
}

func (ap *AirportRepository) CreateAirport(airport *entity.Airport) error {
	return ap.db.Create(airport).Error
}
