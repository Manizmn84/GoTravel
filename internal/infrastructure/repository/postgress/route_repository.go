package postgress

import (
	"errors"

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
	if routeRepo.db.Where("(origin_id = ? AND destination_id = ?) OR (origin_id = ? AND destination_id = ?)",
		route.OriginID, route.DestinationID, route.DestinationID, route.OriginID).
		First(&entity.Route{}).Error == nil {
		return errors.New("The route is Exist.")
	}
	return routeRepo.db.Create(route).Error
}
