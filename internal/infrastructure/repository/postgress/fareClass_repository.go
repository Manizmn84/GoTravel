package postgress

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"gorm.io/gorm"
)

type FareClassRepository struct {
	db *gorm.DB
}

func NewFareClassRepository(db *gorm.DB) *FareClassRepository {
	return &FareClassRepository{db: db}
}

func (fare *FareClassRepository) CreateFareClass(fareClass *entity.FareClass) error {
	return fare.db.Create(fareClass).Error
}

func (fare *FareClassRepository) GetFareClass(id uint) error {
	var fareClass entity.FareClass
	return fare.db.First(&fareClass, id).Error
}
