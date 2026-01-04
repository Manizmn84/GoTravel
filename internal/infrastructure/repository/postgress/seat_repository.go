package postgress

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"gorm.io/gorm"
)

type SeatRepository struct {
	db *gorm.DB
}

func NewSeatRepository(db *gorm.DB) *SeatRepository {
	return &SeatRepository{
		db: db,
	}
}

func (seatRepo *SeatRepository) CreateSeat(seat *entity.Seat) error {
	return seatRepo.db.Create(seat).Error
}
