package service

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/Manizmn84/GoTravel/internal/model"
)

type CompanyService struct {
	companyRepository *postgress.CompanyRepository
}

func NewCompanyService(
	companyRepo *postgress.CompanyRepository,
) *CompanyService {
	return &CompanyService{companyRepository: companyRepo}
}

func (com *CompanyService) CreateCompany(company *entity.Company) error {

	err := com.companyRepository.CreateCompany(company)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (s *CompanyService) AirlineFlightCount(airlineID uint) (*model.AirlineFlightCount, error) {
	return s.companyRepository.AirlineFlightCount(airlineID)
}

func (s *CompanyService) List() ([]entity.Company, error) {
	return s.companyRepository.List()
}

func (s *CompanyService) GetCompanyAirports(companyID uint) ([]model.CompanyAirportDTO, error) {
	return s.companyRepository.GetAirportsByCompany(companyID)
}
