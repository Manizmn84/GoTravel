package service

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/infrastructure/repository/postgress"
)

type RouteService struct {
	routeRepository   *postgress.RouteRepository
	airportRepository *postgress.AirportRepository
}

func NewRouteService(
	routeRepository *postgress.RouteRepository,
	airportRepository *postgress.AirportRepository,
) *RouteService {
	return &RouteService{
		routeRepository:   routeRepository,
		airportRepository: airportRepository,
	}
}

func (routeService *RouteService) Create(route *entity.Route) error {
	if route.OriginID == route.DestinationID {
		return errors.New("the Origin and Destination is same.")
	}

	err := routeService.airportRepository.GetAirport(route.OriginID)
	if err != nil {
		return errors.New("the Origin Airport Not Found.")
	}

	err = routeService.airportRepository.GetAirport(route.DestinationID)

	if err != nil {
		return errors.New("the Destination Airport Not Found.")
	}

	if route.BaseTime <= 0 {
		return errors.New("The Base Time is not Positive.")
	}

	if route.Distance <= 0 {
		return errors.New("The Distance Should be Positive.")
	}

	err = routeService.routeRepository.Create(route)
	if err != nil {
		return err
	}

	return nil
}
