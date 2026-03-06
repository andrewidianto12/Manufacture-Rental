package maintenance_service

import (
	"errors"
	"time"
)

type CreateMaintenanceRequest struct {
	EquipmentID uint   `json:"equipment_id" validate:"required,gte=1"`
	Description string `json:"description" validate:"required"`
	StartDate   string `json:"start_date" validate:"required"`
	EndDate     string `json:"end_date" validate:"required"`
}

type maintenanceService struct {
	repo MaintenanceRepo
}

func NewMaintenanceService(repo MaintenanceRepo) MaintenanceService {
	return &maintenanceService{repo: repo}
}

func (s *maintenanceService) CreateMaintenance(input CreateMaintenanceRequest) (*Maintenance, error) {
	startDate, err := time.Parse("2006-01-02", input.StartDate)
	if err != nil {
		return nil, errors.New("format start_date harus YYYY-MM-DD")
	}
	endDate, err := time.Parse("2006-01-02", input.EndDate)
	if err != nil {
		return nil, errors.New("format end_date harus YYYY-MM-DD")
	}
	if !endDate.After(startDate) {
		return nil, errors.New("end_date harus setelah start_date")
	}

	data := &Maintenance{
		EquipmentID: input.EquipmentID,
		Description: input.Description,
		Status:      "scheduled",
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := s.repo.CreateMaintenance(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *maintenanceService) GetAllMaintenance() ([]Maintenance, error) {
	return s.repo.GetAllMaintenance()
}
