package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
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
	// TODO
	return c.JSON(http.StatusAccepted, map[string]string{
		"msg": "accept",
	})
}
