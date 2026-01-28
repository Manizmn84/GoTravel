package service

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
)

type ReservationService struct {
	reservationRepository *postgress.ReservationRepository
	passengerRepository   *postgress.PassengerRepository
}

func NewReservationService(
	resRepo *postgress.ReservationRepository,
	passRepo *postgress.PassengerRepository,
) *ReservationService {
	return &ReservationService{
		reservationRepository: resRepo,
		passengerRepository:   passRepo,
	}
}

func (re *ReservationService) CreateReservation(reserve *entity.Reserve) error {

	if err := re.passengerRepository.GetPassenger(reserve.PassengerID); err != nil {
		return errors.New("passenger not found")
	}

	if err := re.reservationRepository.CreateReservation(reserve); err != nil {
		return err
	}

	return nil
}
