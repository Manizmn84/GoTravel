package service

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
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

func (tripService *TripService) CreateTrip(trip *entity.Trip) error {

	err := tripService.companyRepository.GetCompany(trip.CompanyID)

	if err != nil {
		return errors.New("company not found")
	}

	if err := tripService.tripRepository.CreateTrip(trip); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
