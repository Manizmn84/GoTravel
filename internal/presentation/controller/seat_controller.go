package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/labstack/echo"
)

type SeatController struct {
	seatService *service.SeatService
}

func NewSeatController(
	seatService *service.SeatService,
) *SeatController {
	return &SeatController{
		seatService: seatService,
	}
}

func (seatcontroller *SeatController) CreateSeat(c echo.Context) error {
	seat := new(entity.Seat)

	if err := c.Bind(seat); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	err := seatcontroller.seatService.CreateSeat(seat)

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"message": err.Error()})
	}

	return c.JSON(http.StatusAccepted, echo.Map{"message": "the seat was create Successfully"})
}
