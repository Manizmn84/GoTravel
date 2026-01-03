package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/labstack/echo"
)

type FareClassController struct {
	fareClassService *service.FareClassService
}

func NewFareClassController(fare *service.FareClassService) *FareClassController {
	return &FareClassController{
		fareClassService: fare,
	}
}

func (fare *FareClassController) CreateFareClass(c echo.Context) error {
	// TODO
	return c.JSON(http.StatusAccepted, map[string]string{"msg": "ok"})
}
