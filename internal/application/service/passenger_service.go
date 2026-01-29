package service

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/domain/enum"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
)

type PassengerService struct {
	passengerRepository *postgress.PassengerRepository
}

func NewPassengerService(repo *postgress.PassengerRepository) *PassengerService {
	return &PassengerService{passengerRepository: repo}
}

func (passengerService *PassengerService) CreatePassenger(passenger *entity.Passenger) error {
	return passengerService.passengerRepository.CreatePassenger(passenger)
}

func (s *PassengerService) List() ([]*entity.Passenger, error) {
	passengers, err := s.passengerRepository.List()
	if err != nil {
		return nil, err
	}

	return passengers, nil
}

func (s *PassengerService) GetPassengersByGender(gender enum.Gender) ([]entity.Passenger, error) {
	return s.passengerRepository.ListByGender(gender)
}
