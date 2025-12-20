package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/labstack/echo"
)

type PassengerController struct {
	passenger_service *service.PassengerService
}

func NewPassengerController(ps *service.PassengerService) *PassengerController {
	return &PassengerController{
		passenger_service: ps,
	}
}

func (pc *PassengerController) CreatePassenger(c echo.Context) error {
	var passenger entity.Passenger

	if err := c.Bind(&passenger); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	if passenger.Email == "" || passenger.Phone == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "email and phone are required",
		})
	}

	if err := pc.passenger_service.CreatePassenger(&passenger); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, passenger)
}
