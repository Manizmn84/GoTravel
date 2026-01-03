package service

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
	"github.com/labstack/echo"
)

type PaymentService struct {
	paymentRepository     *postgress.PaymentRepository
	reservationRepository *postgress.ReservationRepository
}

func NewPaymentService(
	payRepo *postgress.PaymentRepository,
	reservRepo *postgress.ReservationRepository,
) *PaymentService {
	return &PaymentService{
		paymentRepository:     payRepo,
		reservationRepository: reservRepo,
	}
}

func (pay *PaymentService) CreatePayment(c echo.Context) error {
	payment := new(entity.Payment)

	if err := c.Bind(payment); err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	err := pay.reservationRepository.GetReservation(payment.ReserveID)

	if err != nil {
		return c.JSON(404, map[string]string{"error": "reservation not found"})
	}

	if err := pay.paymentRepository.CreatePayment(payment); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(201, payment)
}
