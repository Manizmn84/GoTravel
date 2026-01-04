package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
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
	far := new(entity.FareClass)

	if err := c.Bind(far); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	err := fare.fareClassService.CreateFareClass(far)

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusAccepted, echo.Map{"message": "FareClass was Create Successfully"})
}
