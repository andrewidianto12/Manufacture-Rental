package equipment_category

import (
	equipment_category_service "github.com/andrewidianto12/Manufacture-Rental/service/equipment_category"
	"gorm.io/gorm"
)

type EquipmentCategoryRepository struct {
	db *gorm.DB
}

func NewEquipmentCategoryRepository(db *gorm.DB) *EquipmentCategoryRepository {
	return &EquipmentCategoryRepository{db: db}
}

func (r *EquipmentCategoryRepository) CreateCategory(category *equipment_category_service.EquipmentCategory) error {
	return r.db.Create(category).Error
}

func (r *EquipmentCategoryRepository) GetAllCategories() ([]equipment_category_service.EquipmentCategory, error) {
	var categories []equipment_category_service.EquipmentCategory
	err := r.db.Order("id desc").Find(&categories).Error
	return categories, err
}

func (r *EquipmentCategoryRepository) GetCategoryByID(id uint) (*equipment_category_service.EquipmentCategory, error) {
	var category equipment_category_service.EquipmentCategory
	if err := r.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *EquipmentCategoryRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&equipment_category_service.EquipmentCategory{}, id).Error
}
