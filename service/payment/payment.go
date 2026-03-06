package payment_service

import "time"

type Payment struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	RentalID  uint       `gorm:"column:rental_id;not null" json:"rental_id"`
	Amount    float64    `gorm:"column:amount;not null" json:"amount"`
	Method    string     `gorm:"column:method;not null" json:"method"`
	Status    string     `gorm:"column:status;default:pending" json:"status"`
	PaidAt    *time.Time `gorm:"column:paid_at" json:"paid_at,omitempty"`
	CreatedAt time.Time  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
