package enum

type PaymentStatus int

const (
	PayUnknown PaymentStatus = iota
	PayPaid
	PayNotPaid
)
