package rental_service

import "time"

type Rental struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"column:user_id;not null" json:"user_id"`
	EquipmentID uint      `gorm:"column:equipment_id;not null" json:"equipment_id"`
	RentalDate  time.Time `gorm:"column:rental_date;not null" json:"rental_date"`
	ReturnDate  time.Time `gorm:"column:return_date;not null" json:"return_date"`
	TotalCost   float64   `gorm:"column:total_cost" json:"total_cost"`
	Status      string    `gorm:"column:status;default:active" json:"status"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
