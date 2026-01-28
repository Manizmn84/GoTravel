package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/labstack/echo"
)

type ReservationController struct {
	reservationService *service.ReservationService
}

func NewReservationController(rs *service.ReservationService) *ReservationController {
	return &ReservationController{
		reservationService: rs,
	}
}

func (rs *ReservationController) CreateReservation(c echo.Context) error {
	reserve := new(entity.Reserve)

	if err := c.Bind(reserve); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if err := rs.reservationService.CreateReservation(reserve); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "reservation created successfully",
	})
}
