package enum

type AircraftModel int

const (
	AirbusA320 AircraftModel = iota
	AirbusA321
	Boeing737
	Boeing747
	Boeing777
	EmbraerE190
)

func (am AircraftModel) String() string {
	switch am {
	case AirbusA320:
		return "AribusA320"
	case AirbusA321:
		return "AirbusA321"
	case Boeing737:
		return "Boeing737"
	case Boeing747:
		return "Boeign747"
	case Boeing777:
		return "Boeing777"
	case EmbraerE190:
		return "EmbraerE190"
	}
	return "Unknown"
}
