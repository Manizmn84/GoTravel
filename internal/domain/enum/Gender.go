package enum

type Gender int

const (
	GenderUnknown Gender = iota
	GenderMale
	GenderFemale
)

func (g Gender) String() string {
	switch g {
	case GenderMale:
		return "Male"
	case GenderFemale:
		return "Female"
	}
	return "Unknown"
}
