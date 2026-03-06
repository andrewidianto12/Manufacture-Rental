package maintenance_service

import "time"

type Maintenance struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	EquipmentID uint      `gorm:"column:equipment_id;not null" json:"equipment_id"`
	Description string    `gorm:"column:description" json:"description"`
	Status      string    `gorm:"column:status;default:scheduled" json:"status"`
	StartDate   time.Time `gorm:"column:start_date" json:"start_date"`
	EndDate     time.Time `gorm:"column:end_date" json:"end_date"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Maintenance) TableName() string {
	return "maintenance"
}
