package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/labstack/echo"
)

type CompanyController struct {
	companyService *service.CompanyService
}

func NewCompanyController(com *service.CompanyService) *CompanyController {
	return &CompanyController{
		companyService: com,
	}
}

func (com *CompanyController) CreateCompany(c echo.Context) error {
	company := new(entity.Company)

	if err := c.Bind(company); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	err := com.companyService.CreateCompany(company)

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "Company was created successfully",
	})
}
