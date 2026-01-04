package service

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
)

type SeatService struct {
	seatRepository        *postgress.SeatRepository
	fareRepository        *postgress.FareClassRepository
	reservationRepository *postgress.ReservationRepository
	tripRepository        *postgress.TripRepository
}

func NewSeatService(
	seatRepo *postgress.SeatRepository,
	fareRepo *postgress.FareClassRepository,
	reserveRepo *postgress.ReservationRepository,
	tripRepo *postgress.TripRepository,
) *SeatService {
	return &SeatService{
		seatRepository:        seatRepo,
		fareRepository:        fareRepo,
		reservationRepository: reserveRepo,
		tripRepository:        tripRepo,
	}
}

func (seatService *SeatService) CreateSeat(seat *entity.Seat) error {
	err := seatService.fareRepository.GetFareClass(seat.ClassID)

	if err != nil {
		return errors.New("the fare class is not exist")
	}

	err = seatService.reservationRepository.GetReservation(seat.ReserveID)

	if err != nil {
		return errors.New("the Reserve is not exist")
	}

	err = seatService.tripRepository.GetTrip(seat.TripID)

	if err != nil {
		return errors.New("the trip is not exist")
	}

	err = seatService.seatRepository.CreateSeat(seat)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
