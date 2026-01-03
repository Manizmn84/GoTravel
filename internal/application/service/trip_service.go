package service

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/labstack/echo"
)

type TripService struct {
	tripRepository    *postgress.TripRepository
	companyRepository *postgress.CompanyRepository
}

func NewTripService(
	tripRepo *postgress.TripRepository,
	companyRepo *postgress.CompanyRepository,
) *TripService {
	return &TripService{
		tripRepository:    tripRepo,
		companyRepository: companyRepo,
	}
}

func (tripService *TripService) CreateTrip(c echo.Context) error {
	trip := new(entity.Trip)

	if err := c.Bind(trip); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	err := tripService.companyRepository.GetCompany(trip.CompanyID)

	if err != nil {
		return c.JSON(404, map[string]string{"error": "company not found"})
	}

	if err := tripService.tripRepository.CreateTrip(trip); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, trip)
}
