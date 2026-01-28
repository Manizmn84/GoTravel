package model

type AirlineFlightCount struct {
	AirlineID   uint   `json:"airline_id"`
	AirlineName string `json:"airline_name"`
	FlightCount int64  `json:"flight_count"`
}
