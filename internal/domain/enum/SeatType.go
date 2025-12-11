package enum

type SeatType int

const (
	SeatUnknown SeatType = iota
	SeatEconomy
	SeatBusiness
)

func (st SeatType) String() string {
	switch st {
	case SeatEconomy:
		return "Economy"
	case SeatBusiness:
		return "Business"
	}

	return "Unknown"
}
