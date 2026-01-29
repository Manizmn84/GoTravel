package model

type CompanyAirportDTO struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	City     string `json:"city"`
	Country  string `json:"country"`
	IATACode string `json:"iata_code"`
}
