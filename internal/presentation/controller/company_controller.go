package controller

import (
	"net/http"
	"strconv"

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

// internal/report/handler/airline_flights.go
func (h *CompanyController) AirlineFlightCount(c echo.Context) error {
	airlineIDStr := c.QueryParam("airline_id")

	airlineID, err := strconv.ParseUint(airlineIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid airline_id",
		})
	}

	data, err := h.companyService.AirlineFlightCount(uint(airlineID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (h *CompanyController) List(c echo.Context) error {
	airlines, err := h.companyService.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": airlines,
	})
}
