package controller

import (
	"net/http"

	"github.com/Manizmn84/GoTravel/internal/application/service"
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
	return c.JSON(http.StatusAccepted, map[string]string{"msg": "ok"})
}
