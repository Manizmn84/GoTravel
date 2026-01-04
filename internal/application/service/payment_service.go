package service

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
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

func (pay *PaymentService) CreatePayment(payment *entity.Payment) error {

	err := pay.reservationRepository.GetReservation(payment.ReserveID)

	if err != nil {
		return errors.New("reservation not found")
	}

	if err := pay.paymentRepository.CreatePayment(payment); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
