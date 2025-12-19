package service

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/labstack/echo"
)

type AirportService struct {
	airportRepository *postgress.AirportRepository
}

func NewAirportService(
	airRepo *postgress.AirportRepository,
) *AirportService {
	return &AirportService{airportRepository: airRepo}
}

func (air *AirportService) CreateAirport(c echo.Context) error {
	airport := new(entity.Airport)

	if err := c.Bind(airport); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	err := air.airportRepository.CreateAirport(airport)

	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, airport)
}
