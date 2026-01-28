package model

type AirportDependencyReport struct {
	AirportID      uint    `json:"airport_id"`
	AirportName    string  `json:"airport_name"`
	CompanyID      uint    `json:"company_id"`
	CompanyName    string  `json:"company_name"`
	FlightCount    int     `json:"flight_count"`
	DependencyRate float64 `json:"dependency_ratio"`
}
