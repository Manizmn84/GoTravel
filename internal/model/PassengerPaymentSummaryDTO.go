package model

type PassengerPaymentSummaryDTO struct {
	PassengerID uint    `json:"passenger_id"`
	FirstName   string  `json:"first_name"`
	LastName    string  `json:"last_name"`
	Email       string  `json:"email"`
	Gender      int     `json:"gender"`
	TotalAmount float64 `json:"total_amount"`
}
