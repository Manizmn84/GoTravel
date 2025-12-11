package enum

type UserType int

const (
	Admin UserType = iota
	Customer
)

func (ut UserType) String() string {
	switch ut {
	case Admin:
		return "Admin"
	case Customer:
		return "Customer"
	}

	return "Unknown"
}
