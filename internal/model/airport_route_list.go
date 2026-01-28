package model

type AirportRouteItem struct {
	RouteID            uint   `json:"route_id"`
	OriginAirport      string `json:"origin_airport"`
	DestinationAirport string `json:"destination_airport"`
	RouteType          string `json:"route_type"` // INCOMING | OUTGOING
}
