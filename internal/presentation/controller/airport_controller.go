package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
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
	return c.JSON(http.StatusAccepted, map[string]string{"msg": "ok"})
}
