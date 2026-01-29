package postgress

import (
	"errors"

	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"github.com/Manizmn84/GoTravel/internal/model"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (co *CompanyRepository) CreateCompany(company *entity.Company) error {

	if co.db.Where("Name = ? AND Country = ?", company.Name, company.Country).First(&entity.Company{}).Error == nil {
		return errors.New("The Instance is Exist.")
	}

	return co.db.Create(company).Error
}

func (co *CompanyRepository) GetCompany(id uint) error {
	var company entity.Company
	return co.db.First(&company, id).Error
}

// internal/report/repository/airline_flights.go
func (r *CompanyRepository) AirlineFlightCount(companyID uint) (*model.AirlineFlightCount, error) {
	var result model.AirlineFlightCount

	err := r.db.Table("companies c").
		Select(`
			c.id   AS airline_id,
			c.name AS airline_name,
			COUNT(t.id) AS flight_count
		`).
		Joins("LEFT JOIN trips t ON t.company_id = c.id").
		Where("c.id = ?", companyID).
		Group("c.id, c.name").
		Scan(&result).Error

	return &result, err
}

func (r *CompanyRepository) List() ([]entity.Company, error) {
	var airlines []entity.Company

	err := r.db.Find(&airlines).Error
	if err != nil {
		return nil, err
	}

	return airlines, nil
}

func (r *CompanyRepository) GetAirportsByCompany(companyID uint) ([]model.CompanyAirportDTO, error) {
	var airports []model.CompanyAirportDTO

	err := r.db.
		Table("airports a").
		Select(`
			DISTINCT a.id,
			a.name,
			a.city,
			a.country,
			a.lata_code as iata_code
		`).
		Joins(`
			JOIN routes r 
			  ON r.origin_id = a.id OR r.destination_id = a.id
		`).
		Joins(`
			JOIN trip_routes tr ON tr.route_id = r.id
		`).
		Joins(`
			JOIN trips t ON t.id = tr.trip_id
		`).
		Where("t.company_id = ?", companyID).
		Scan(&airports).Error

	return airports, err
}
