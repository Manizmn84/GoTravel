package service

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/labstack/echo"
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

func (re *ReservationService) CreateReservation(c echo.Context) error {
	reserve := new(entity.Reserve)

	if err := c.Bind(reserve); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	err := re.passengerRepository.GetPassenger(reserve.PassengerID)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "passenger not found"})
	}

	if err := re.reservationRepository.CreateReservation(reserve); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, reserve)
}
