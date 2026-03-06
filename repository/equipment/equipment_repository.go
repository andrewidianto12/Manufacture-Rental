package equipment

import (
	equipment_service "github.com/andrewidianto12/Manufacture-Rental/service/equipment"
	"gorm.io/gorm"
)

type EquipmentRepository struct {
	db *gorm.DB
}

func NewEquipmentRepository(db *gorm.DB) *EquipmentRepository {
	return &EquipmentRepository{db: db}
}

func (r *EquipmentRepository) CreateEquipment(equipment *equipment_service.Equipment) error {
	return r.db.Create(equipment).Error
}

func (r *EquipmentRepository) GetAllEquipment() ([]equipment_service.Equipment, error) {
	var equipments []equipment_service.Equipment
	err := r.db.Order("id desc").Find(&equipments).Error
	return equipments, err
}

func (r *EquipmentRepository) GetEquipmentByID(id uint) (*equipment_service.Equipment, error) {
	var equipment equipment_service.Equipment
	if err := r.db.First(&equipment, id).Error; err != nil {
		return nil, err
	}
	return &equipment, nil
}

func (r *EquipmentRepository) DeleteEquipment(id uint) error {
	return r.db.Delete(&equipment_service.Equipment{}, id).Error
}
