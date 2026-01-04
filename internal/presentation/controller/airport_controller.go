package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/labstack/echo"
)

type AirportController struct {
	airPortService *service.AirportService
}

func NewAirportController(air *service.AirportService) *AirportController {
	return &AirportController{
		airPortService: air,
	}
}

func (air *AirportController) CreateAirport(c echo.Context) error {
	// TODO
	airport := new(entity.Airport)

	if err := c.Bind(airport); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	err := air.airPortService.CreateAirport(airport)

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusAccepted, map[string]string{"msg": "ok"})
}
