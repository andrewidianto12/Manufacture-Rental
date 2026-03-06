package notification_service

import "time"

type Notification struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"column:user_id;not null" json:"user_id"`
	Channel   string    `gorm:"column:channel;not null" json:"channel"`
	Recipient string    `gorm:"column:recipient;not null" json:"recipient"`
	Subject   string    `gorm:"column:subject" json:"subject"`
	Message   string    `gorm:"column:message" json:"message"`
	Status    string    `gorm:"column:status;default:sent" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}
