package service

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
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
