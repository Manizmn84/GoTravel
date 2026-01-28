package postgress

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
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
