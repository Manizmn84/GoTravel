package service

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/labstack/echo"
)

type PassengerService struct {
	passengerRepository *postgress.PassengerRepository
}

func NewPassengerService(repo *postgress.PassengerRepository) *PassengerService {
	return &PassengerService{passengerRepository: repo}
}

func (passengerService *PassengerService) CreatePassenger(c echo.Context) error {
	passenger := new(entity.Passenger)

	if err := c.Bind(passenger); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	err := passengerService.passengerRepository.CreatePassenger(passenger)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, passenger)
}
