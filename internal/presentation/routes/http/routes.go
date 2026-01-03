package http

import (
	"github.com/labstack/echo"

	"github.com/Manizmn84/GoTravel/internal/presentation/controller"
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
) *Router {

	e := echo.New()

	e.POST("/CreatePassenger", passengerController.CreatePassenger)
	e.POST("/CreateReservation", reservationController.CreateReservation)
	e.POST("/CreatePayment", paymentController.CreatePayment)
	e.POST("/CreateCompany", companyController.CreateCompany)
	e.POST("/CreateAirport", airportController.CreateAirport)
	e.POST("/CreateFareClass", fareClassController.CreateFareClass)

	return &Router{
		Echo: e,
	}
}

func (r *Router) Start(address string) error {
	return r.Echo.Start(address)
}
