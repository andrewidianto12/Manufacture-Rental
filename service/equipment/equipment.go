package equipment_service

import "time"

type Equipment struct {
	ID           uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryID   uint       `gorm:"column:category_id" json:"category_id"`
	Name         string     `gorm:"column:name;not null" json:"name"`
	Description  string     `gorm:"column:description" json:"description"`
	DailyRate    float64    `gorm:"column:daily_rate;not null" json:"daily_rate"`
	Status       string     `gorm:"column:status;default:available" json:"status"`
	PurchaseDate *time.Time `gorm:"column:purchase_date" json:"purchase_date,omitempty"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Equipment) TableName() string {
	return "equipment"
}
