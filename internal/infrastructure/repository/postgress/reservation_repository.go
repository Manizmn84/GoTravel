package postgress

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db *gorm.DB) *ReservationRepository {
	return &ReservationRepository{db: db}
}

func (re *ReservationRepository) CreateReservation(reservation *entity.Reserve) error {
	return re.db.Create(reservation).Error
}

func (re *ReservationRepository) GetReservation(id uint) error {
	var reservation entity.Reserve
	return re.db.First(&reservation, id).Error
}
