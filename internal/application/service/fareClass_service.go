package service

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
)

type FareClassService struct {
	fareClassRepository *postgress.FareClassRepository
}

func NewFareClassService(
	fareRepo *postgress.FareClassRepository,
) *FareClassService {
	return &FareClassService{fareClassRepository: fareRepo}
}

func (fareService *FareClassService) CreateFareClass(fare *entity.FareClass) error {

	err := fareService.fareClassRepository.CreateFareClass(fare)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}
