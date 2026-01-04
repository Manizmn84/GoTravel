package service

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
)

type AirportService struct {
	airportRepository *postgress.AirportRepository
}

func NewAirportService(
	airRepo *postgress.AirportRepository,
) *AirportService {
	return &AirportService{airportRepository: airRepo}
}

func (air *AirportService) CreateAirport(airport *entity.Airport) error {

	err := air.airportRepository.CreateAirport(airport)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
