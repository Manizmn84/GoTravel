package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/labstack/echo"
)

type AirportController struct {
	airPortService *service.AirportService
}

func NewAirportController(air *service.AirportService) *AirportController {
	return &AirportController{
		airPortService: air,
	}
}

func (air *AirportController) CreateAirport(c echo.Context) error {
	// TODO
	airport := new(entity.Airport)

	if err := c.Bind(airport); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	err := air.airPortService.CreateAirport(airport)

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusAccepted, map[string]string{"msg": "ok"})
}

func (air *AirportController) Update(c echo.Context) error {
	fmt.Printf("this is in controller\n")
	airport := new(entity.Airport)
	id, _ := strconv.Atoi(c.Param("id"))

	if err := air.airPortService.Exist(uint(id)); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	if err := c.Bind(airport); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}
	airport.ID = uint(id)
	err := air.airPortService.Update(airport)

	if err != nil {
		return c.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusAccepted, map[string]string{"msg": "Update Is Succesfully"})

}

func (air *AirportController) ListAllAirport(c echo.Context) error {
	ListAirport := air.airPortService.ListAllAirport()
	return c.JSON(http.StatusAccepted, map[string]interface{}{"msg": ListAirport})
}

func (c *AirportController) List(ctx echo.Context) error {
	airports, err := c.airPortService.List()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": airports,
	})
}

func (c *AirportController) AirportRoutes(ctx echo.Context) error {
	airportIDStr := ctx.QueryParam("airport_id")
	if airportIDStr == "" {
		return ctx.JSON(400, echo.Map{
			"error": "airport_id is required",
		})
	}

	airportID, err := strconv.Atoi(airportIDStr)
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "invalid airport_id",
		})
	}

	data, err := c.airPortService.AirportRoutesCount(uint(airportID))
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": data,
	})
}

func (h *AirportController) AirportRoutesList(c echo.Context) error {
	airportIDStr := c.QueryParam("airport_id")
	airportID, err := strconv.ParseUint(airportIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "invalid airport_id",
		})
	}

	data, err := h.airPortService.AirportRouteList(uint(airportID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}

func (c *AirportController) AirportDependency(ctx echo.Context) error {
	airportIDStr := ctx.QueryParam("airport_id")
	if airportIDStr == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "airport_id is required",
		})
	}

	airportID, err := strconv.ParseUint(airportIDStr, 10, 64)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "invalid airport_id",
		})
	}

	result, err := c.airPortService.GetAirportDependency(uint(airportID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, result)
}
