package user_service

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"column:username;unique;not null" json:"username"`
	Email     string    `gorm:"column:email;unique;not null" json:"email"`
	Password  string    `gorm:"column:password;not null" json:"-"`
	Fullname  string    `gorm:"column:full_name" json:"full_name"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Role      string    `gorm:"column:role;default:user" json:"role"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}
