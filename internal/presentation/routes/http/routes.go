package http

import (
	"github.com/Manizmn84/GoTravel/internal/presentation/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Router struct {
	Echo *echo.Echo
}

func NewRouter(
	passengerController *controller.PassengerController,
	reservationController *controller.ReservationController,
	paymentController *controller.PaymentController,
	companyController *controller.CompanyController,
	airportController *controller.AirportController,
	fareClassController *controller.FareClassController,
	tripController *controller.TripController,
	seatController *controller.SeatController,
	routeController *controller.RouteController,
) *Router {

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			echo.GET,
			echo.POST,
			echo.PUT,
			echo.DELETE,
			echo.OPTIONS,
		},
	}))

	e.POST("/CreatePassenger", passengerController.CreatePassenger)
	e.POST("/CreateReservation", reservationController.CreateReservation)
	e.POST("/CreatePayment", paymentController.CreatePayment)
	e.POST("/CreateCompany", companyController.CreateCompany)
	e.POST("/CreateAirport", airportController.CreateAirport)
	e.POST("/CreateFareClass", fareClassController.CreateFareClass)
	e.POST("/CreateTrip", tripController.CreateTrip)
	e.POST("/CreateSeat", seatController.CreateSeat)
	e.POST("/CreateRoute", routeController.CreateRoute)
	e.PUT("/UpdateAirport/:id", airportController.Update)
	e.GET("/airports", airportController.List)
	e.GET("/reports/airport-routes", airportController.AirportRoutes)
	e.GET("/reports/airport-routes-list", airportController.AirportRoutesList)
	e.GET("/airlines", companyController.List)
	e.GET("/reports/airline-flight-count", companyController.AirlineFlightCount)
	e.GET("/reports/airport-dependency", airportController.AirportDependency)
	e.GET("/passengers", passengerController.List)
	e.GET("/companies", companyController.List)
	e.GET("/passengersGender", passengerController.ListByGender)
	e.GET("/passengers/by-payment", passengerController.GetByPaymentStatus)

	return &Router{
		Echo: e,
	}
}

func (r *Router) Start(address string) error {
	return r.Echo.Start(address)
}
