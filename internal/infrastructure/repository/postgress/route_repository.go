package postgress

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"gorm.io/gorm"
)

type RouteRepository struct {
	db *gorm.DB
}

func NewRouteRepository(db *gorm.DB) *RouteRepository {
	return &RouteRepository{
		db: db,
	}
}

func (routeRepo *RouteRepository) Create(route *entity.Route) error {
	return routeRepo.db.Create(route).Error
}
