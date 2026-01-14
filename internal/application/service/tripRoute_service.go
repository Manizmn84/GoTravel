package service

import (
	"gorm.io/gorm"
)

type TripRouteService struct {
	db *gorm.DB
}

func NewTripRouteService(db *gorm.DB) *TripRouteService {
	return &TripRouteService{db: db}
}

func (tr *TripRouteService) Create() error {
	return nil
}
