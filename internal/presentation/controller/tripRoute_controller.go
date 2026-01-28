package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/labstack/echo"
)

type TripRouteController struct {
	tripRouteServcie *service.TripRouteService
}

func NewTripRouteController(tripRouteService *service.TripRouteService) *TripRouteController {
	return &TripRouteController{tripRouteServcie: tripRouteService}
}

func (tripRoute *TripRouteController) Create(c echo.Context) error {
	trip_Route := new(struct {
		trip_id  int `gorm:"not null"`
		route_id int `gorm:"not null"`
	})

	if err := c.Bind(&trip_Route); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid request body",
		})
	}

	return nil
}
