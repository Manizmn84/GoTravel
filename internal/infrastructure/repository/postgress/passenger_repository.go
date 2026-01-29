package postgress

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"github.com/Manizmn84/GoTravel/internal/model"
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

func (r *PassengerRepository) ListByGender(gender enum.Gender) ([]entity.Passenger, error) {
	var passengers []entity.Passenger

	err := r.db.
		Preload("Reserves").
		Where("gender = ?", gender).
		Find(&passengers).Error

	return passengers, err
}

func (r *PassengerRepository) FindByPaymentStatus(
	status enum.PaymentStatus,
) ([]model.PassengerPaymentSummaryDTO, error) {

	var result []model.PassengerPaymentSummaryDTO

	err := r.db.
		Table("passengers p").
		Select(`
			p.id as passenger_id,
			p.first_name,
			p.last_name,
			p.email,
			p.gender,
			SUM(r.total_amount) as total_amount
		`).
		Joins("JOIN reserves r ON r.passenger_id = p.id").
		Joins("JOIN payments pay ON pay.reserve_id = r.id").
		Where("pay.status = ?", status).
		Group("p.id, p.first_name, p.last_name, p.email, p.gender").
		Having("SUM(r.total_amount) > 0").
		Scan(&result).Error

	return result, err
}
