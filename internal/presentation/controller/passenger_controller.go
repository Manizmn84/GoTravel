package controller

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
)

type PassengerController struct {
	passengerService *service.PassengerService
}

func NewPassengerController(ps *service.PassengerService) *PassengerController {
	return &PassengerController{
		passengerService: ps,
	}
}

func (pc *PassengerController) CreatePassenger(c echo.Context) error {
	var passenger entity.Passenger

	if err := c.Bind(&passenger); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid request body",
		})
	}

	if passenger.Email == "" || passenger.Phone == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "email and phone are required",
		})
	}

	if err := pc.passengerService.CreatePassenger(&passenger); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, passenger)
}
