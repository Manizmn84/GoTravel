package service

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/labstack/echo"
)

type CompanyService struct {
	companyRepository *postgress.CompanyRepository
}

func NewCompanyService(
	companyRepo *postgress.CompanyRepository,
) *CompanyService {
	return &CompanyService{companyRepository: companyRepo}
}

func (com *CompanyService) CreateCompany(c echo.Context) error {
	company := new(entity.Company)

	if err := c.Bind(company); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	err := com.companyRepository.CreateCompany(company)

	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, company)
}
