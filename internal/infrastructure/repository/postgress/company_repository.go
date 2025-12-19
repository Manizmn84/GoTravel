package postgress

import (
	"github.com/Manizmn84/GoTravel/internal/domain/entity"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{db: db}
}

func (co *CompanyRepository) CreateCompany(company *entity.Company) error {
	return co.db.Create(company).Error
}
