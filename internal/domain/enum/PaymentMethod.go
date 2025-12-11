package enum

type PaymentMethod int

const (
	Cash PaymentMethod = iota
	PaymentMethodOnline
	Card
)

func (pm PaymentMethod) String() string {
	switch pm {
	case Cash:
		return "Cash"
	case PaymentMethodOnline:
		return "Online"
	case Card:
		return "Card"
	}

	return "Unknown"
}
