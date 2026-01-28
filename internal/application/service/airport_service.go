package service

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/Manizmn84/GoTravel/internal/model"
	"gorm.io/gorm"
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

func (air *AirportService) Update(airport *entity.Airport) error {
	if err := air.airportRepository.GetAirport(airport.ID); err != nil {
		return errors.New("The airport with that Id is Not Exist.")
	}
	err := air.airportRepository.Update(airport)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (air *AirportService) Exist(id uint) error {
	return air.airportRepository.GetAirport(id)
}

func (air *AirportService) ListAllAirport() *gorm.DB {
	ListAirport := air.airportRepository.ListAllAirport()
	return ListAirport
}

func (s *AirportService) AirportRoutesCount(airportID uint) (*postgress.AirportRouteCount, error) {
	return s.airportRepository.AirportRoutesCount(airportID)
}

func (s *AirportService) List() ([]entity.Airport, error) {
	return s.airportRepository.List()
}

func (s *AirportService) AirportRouteList(airportID uint) ([]model.AirportRouteItem, error) {
	return s.airportRepository.AirportRouteList(airportID)
}

func (s *AirportService) GetAirportDependency(airportID uint) ([]model.AirportDependencyReport, error) {
	return s.airportRepository.AirportDependency(airportID)
}
