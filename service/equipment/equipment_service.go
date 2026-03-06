package equipment_service

import (
	"errors"
	"strings"
)

type CreateEquipmentRequest struct {
	CategoryID  uint    `json:"category_id" validate:"required,gte=1"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	DailyRate   float64 `json:"daily_rate" validate:"required,gt=0"`
	Status      string  `json:"status" validate:"omitempty,oneof=available rented maintenance"`
}

type equipmentService struct {
	repo EquipmentRepo
}

func NewEquipmentService(repo EquipmentRepo) EquipmentService {
	return &equipmentService{repo: repo}
}

func (s *equipmentService) CreateEquipment(input CreateEquipmentRequest) (*Equipment, error) {
	if strings.TrimSpace(input.Name) == "" || input.DailyRate <= 0 || input.CategoryID == 0 {
		return nil, errors.New("input equipment tidak valid")
	}

	status := input.Status
	if status == "" {
		status = "available"
	}

	equipment := &Equipment{
		CategoryID:  input.CategoryID,
		Name:        input.Name,
		Description: input.Description,
		DailyRate:   input.DailyRate,
		Status:      status,
	}

	if err := s.repo.CreateEquipment(equipment); err != nil {
		return nil, err
	}

	return equipment, nil
}

func (s *equipmentService) GetAllEquipment() ([]Equipment, error) {
	return s.repo.GetAllEquipment()
}

func (s *equipmentService) GetEquipmentByID(id uint) (*Equipment, error) {
	return s.repo.GetEquipmentByID(id)
}

func (s *equipmentService) DeleteEquipment(id uint) error {
	if id == 0 {
		return errors.New("id tidak valid")
	}
	return s.repo.DeleteEquipment(id)
}
