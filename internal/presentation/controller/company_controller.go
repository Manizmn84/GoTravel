package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
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
	return c.JSON(http.StatusAccepted, map[string]string{"msg": "ok"})
}
