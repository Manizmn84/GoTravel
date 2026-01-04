package service

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
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
