package postgress

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/model"
	"gorm.io/gorm"
)

type AirportRepository struct {
	db *gorm.DB
}

func NewAirportRepository(db *gorm.DB) *AirportRepository {
	return &AirportRepository{db: db}
}

func (ap *AirportRepository) CreateAirport(airport *entity.Airport) error {
	if ap.db.Where("Country = ? AND LATA_Code = ?", airport.Country, airport.LATA_Code).First(&entity.Airport{}).Error == nil {
		return errors.New("This Instance is Exist.")
	}
	return ap.db.Create(airport).Error
}

func (ap *AirportRepository) GetAirport(id uint) error {
	return ap.db.First(&entity.Airport{}, id).Error
}

func (ap *AirportRepository) ListAllAirport() *gorm.DB {
	AllAirport := []entity.Airport{}
	return ap.db.Find(&AllAirport)
}

func (ap *AirportRepository) Update(airport *entity.Airport) error {
	var air entity.Airport
	ap.db.Where("ID = ?", airport.ID).First(&air)
	return ap.db.Model(&air).Updates(airport).Error
}

type AirportRouteCount struct {
	AirportID     uint
	AirportName   string
	IncomingCount int64
	OutgoingCount int64
}

func (r *AirportRepository) AirportRoutesCount(airportID uint) (*AirportRouteCount, error) {
	var result AirportRouteCount

	err := r.db.
		Table("airports a").
		Select(`
			a.id as airport_id,
			a.name as airport_name,
			COUNT(DISTINCT rin.id) as incoming_count,
			COUNT(DISTINCT rout.id) as outgoing_count
		`).
		Joins("LEFT JOIN routes rin ON rin.destination_id = a.id").
		Joins("LEFT JOIN routes rout ON rout.origin_id = a.id").
		Where("a.id = ?", airportID).
		Group("a.id, a.name").
		Scan(&result).Error

	return &result, err
}

func (r *AirportRepository) List() ([]entity.Airport, error) {
	var airports []entity.Airport
	err := r.db.Find(&airports).Error
	return airports, err
}

// internal/report/repository/airport_routes.go
func (r *AirportRepository) AirportRouteList(airportID uint) ([]model.AirportRouteItem, error) {
	var routes []model.AirportRouteItem

	query := `
		SELECT
			r.id AS route_id,
			o.name AS origin_airport,
			d.name AS destination_airport,
			CASE
				WHEN r.origin_id = ? THEN 'OUTGOING'
				ELSE 'INCOMING'
			END AS route_type
		FROM routes r
		JOIN airports o ON o.id = r.origin_id
		JOIN airports d ON d.id = r.destination_id
		WHERE r.origin_id = ? OR r.destination_id = ?
		ORDER BY route_type
	`

	err := r.db.Raw(query, airportID, airportID, airportID).Scan(&routes).Error
	return routes, err
}

func (r *AirportRepository) AirportDependency(airportID uint) ([]model.AirportDependencyReport, error) {
	var result []model.AirportDependencyReport

	err := r.db.Table("airports a").
		Select(`
			a.id AS airport_id,
			a.name AS airport_name,
			c.id AS company_id,
			c.name AS company_name,
			COUNT(DISTINCT t.id) AS flight_count,
			COALESCE(
				COUNT(DISTINCT t.id)::float / NULLIF(total.total_flights, 0),
				0
			) AS dependency_ratio
		`).
		Joins(`
			LEFT JOIN routes r 
			  ON r.origin_id = a.id 
			  OR r.destination_id = a.id
		`).
		Joins(`
			LEFT JOIN trip_routes tr 
			  ON tr.route_id = r.id
		`).
		Joins(`
			LEFT JOIN trips t 
			  ON t.id = tr.trip_id
		`).
		Joins(`
			LEFT JOIN companies c 
			  ON c.id = t.company_id
		`).
		Joins(`
			LEFT JOIN (
				SELECT
					a2.id AS airport_id,
					COUNT(DISTINCT t2.id) AS total_flights
				FROM airports a2
				LEFT JOIN routes r2 
				  ON r2.origin_id = a2.id 
				  OR r2.destination_id = a2.id
				LEFT JOIN trip_routes tr2 
				  ON tr2.route_id = r2.id
				LEFT JOIN trips t2 
				  ON t2.id = tr2.trip_id
				GROUP BY a2.id
			) total ON total.airport_id = a.id
		`).
		Where("a.id = ?", airportID).
		Group("a.id, a.name, c.id, c.name, total.total_flights").
		Order("dependency_ratio DESC").
		Scan(&result).Error

	return result, err
}
