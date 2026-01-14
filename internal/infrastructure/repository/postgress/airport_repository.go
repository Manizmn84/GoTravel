package postgress

import (
	"errors"

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
	if ap.db.Where("Country = ? AND LATA_Code = ?", airport.Country, airport.LATA_Code).First(&entity.Airport{}).Error == nil {
		return errors.New("This Instance is Exist.")
	}
	return ap.db.Create(airport).Error
}

func (ap *AirportRepository) GetAirport(id uint) error {
	return ap.db.First(&entity.Airport{}, id).Error
}
