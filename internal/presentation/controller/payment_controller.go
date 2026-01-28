package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/labstack/echo"
)

type PaymentController struct {
	paymentService *service.PaymentService
}

func NewPaymentController(pay *service.PaymentService) *PaymentController {
	return &PaymentController{
		paymentService: pay,
	}
}

func (py *PaymentController) CreatePayment(c echo.Context) error {
	// TODO
	payment := new(entity.Payment)

	if err := c.Bind(payment); err != nil {
		return c.JSON(400, echo.Map{"error": err.Error()})
	}

	err := py.paymentService.CreatePayment(payment)

	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"message": "Payment was created successfully",
	})
}
