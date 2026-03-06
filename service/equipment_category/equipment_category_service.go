package equipment_category_service

import (
	"errors"
	"strings"
)

type CreateCategoryRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type equipmentCategoryService struct {
	repo EquipmentCategoryRepo
}

func NewEquipmentCategoryService(repo EquipmentCategoryRepo) EquipmentCategoryService {
	return &equipmentCategoryService{repo: repo}
}

func (s *equipmentCategoryService) CreateCategory(input CreateCategoryRequest) (*EquipmentCategory, error) {
	if strings.TrimSpace(input.Name) == "" {
		return nil, errors.New("name wajib diisi")
	}

	category := &EquipmentCategory{
		Name:        input.Name,
		Description: input.Description,
	}

	if err := s.repo.CreateCategory(category); err != nil {
		return nil, err
	}

	return category, nil
}

func (s *equipmentCategoryService) GetAllCategories() ([]EquipmentCategory, error) {
	return s.repo.GetAllCategories()
}

func (s *equipmentCategoryService) GetCategoryByID(id uint) (*EquipmentCategory, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}
	return s.repo.GetCategoryByID(id)
}

func (s *equipmentCategoryService) DeleteCategory(id uint) error {
	if id == 0 {
		return errors.New("id tidak valid")
	}
	return s.repo.DeleteCategory(id)
}
