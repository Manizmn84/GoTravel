package enum

type PaymentMethod int

const (
	Cash PaymentMethod = iota
	Online
	Card
)
