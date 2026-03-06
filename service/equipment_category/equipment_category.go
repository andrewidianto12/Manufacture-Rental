package equipment_category_service

import "time"

type EquipmentCategory struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"column:name;not null" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (EquipmentCategory) TableName() string {
	return "equipment_categories"
}
