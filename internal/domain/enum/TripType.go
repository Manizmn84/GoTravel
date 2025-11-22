package enum

type TripType int

const (
	TripUnknown TripType = iota
	TripAir
	TripRail
)
