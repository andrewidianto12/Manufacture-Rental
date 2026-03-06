package maintenance

import (
	maintenance_service "github.com/andrewidianto12/Manufacture-Rental/service/maintenance"
	"gorm.io/gorm"
)

type MaintenanceRepository struct {
	db *gorm.DB
}

func NewMaintenanceRepository(db *gorm.DB) *MaintenanceRepository {
	return &MaintenanceRepository{db: db}
}

func (r *MaintenanceRepository) CreateMaintenance(data *maintenance_service.Maintenance) error {
	return r.db.Create(data).Error
}

func (r *MaintenanceRepository) GetAllMaintenance() ([]maintenance_service.Maintenance, error) {
	var result []maintenance_service.Maintenance
	err := r.db.Order("id desc").Find(&result).Error
	return result, err
}
