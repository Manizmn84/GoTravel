package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/labstack/echo"
)

type TripController struct {
	tripService *service.TripService
}

func NewTripController(ts *service.TripService) *TripController {
	return &TripController{
		tripService: ts,
	}
}
func (tripcCon *TripController) CreateTrip(c echo.Context) error {
	trip := new(entity.Trip)

	if err := c.Bind(trip); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	err := tripcCon.tripService.CreateTrip(trip)

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusAccepted, echo.Map{"message": "Trip was Create Successfully"})

}
