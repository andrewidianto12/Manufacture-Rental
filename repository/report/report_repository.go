package report

import (
	report_service "github.com/andrewidianto12/Manufacture-Rental/service/report"
	"gorm.io/gorm"
)

type ReportRepository struct {
	db *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (r *ReportRepository) GetDashboardReport() (*report_service.DashboardReport, error) {
	result := &report_service.DashboardReport{}

	if err := r.db.Table("users").Count(&result.TotalUsers).Error; err != nil {
		return nil, err
	}
	if err := r.db.Table("equipment").Count(&result.TotalEquipment).Error; err != nil {
		return nil, err
	}
	if err := r.db.Table("rentals").Where("status = ?", "active").Count(&result.TotalActiveRentals).Error; err != nil {
		return nil, err
	}
	if err := r.db.Table("payments").Select("COALESCE(SUM(amount), 0)").Scan(&result.TotalRevenue).Error; err != nil {
		return nil, err
	}

	return result, nil
}
