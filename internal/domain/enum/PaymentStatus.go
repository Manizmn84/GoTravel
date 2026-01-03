package enum

type PaymentStatus int

const (
	PayUnknown PaymentStatus = iota
	PayPaid
	PayNotPaid
)

func (ps PaymentStatus) String() string {
	switch ps {
	case PayPaid:
		return "Paid"
	case PayNotPaid:
		return "Not Paid"
	}

	return "Unknown"
}
