package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/labstack/echo"
)

type RouteController struct {
	routeService *service.RouteService
}

func NewRouteController(
	routeService *service.RouteService,
) *RouteController {
	return &RouteController{
		routeService: routeService,
	}
}

func (routeController *RouteController) CreateRoute(c echo.Context) error {
	route := new(entity.Route)

	if err := c.Bind(route); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	err := routeController.routeService.Create(route)

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "The Route Was Create Successfully",
	})
}
