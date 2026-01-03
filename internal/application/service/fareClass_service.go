package service

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/labstack/echo"
)

type FareClassService struct {
	fareClassRepository *postgress.FareClassRepository
}

func NewFareClassService(
	fareRepo *postgress.FareClassRepository,
) *FareClassService {
	return &FareClassService{fareClassRepository: fareRepo}
}

func (fareService *FareClassService) CreateFareClass(c echo.Context) error {
	fare := new(entity.FareClass)

	if err := c.Bind(fare); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	err := fareService.fareClassRepository.CreateFareClass(fare)

	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, fare)
}
