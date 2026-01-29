package controller

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/domain/enum"
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

func (c *PassengerController) List(ctx echo.Context) error {
	passengers, err := c.passengerService.List()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": passengers,
	})
}

func (c *PassengerController) ListByGender(cxt echo.Context) error {
	genderStr := cxt.QueryParam("gender")

	var gender enum.Gender
	switch genderStr {
	case "Male":
		gender = enum.GenderMale
	case "Female":
		gender = enum.GenderFemale
	default:
		gender = enum.GenderUnknown
	}

	passengers, err := c.passengerService.GetPassengersByGender(gender)
	if err != nil {
		return cxt.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return cxt.JSON(http.StatusOK, echo.Map{
		"data": passengers,
	})
}
