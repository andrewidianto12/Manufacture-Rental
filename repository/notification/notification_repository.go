package notification

import (
	notification_service "github.com/andrewidianto12/Manufacture-Rental/service/notification"
	"gorm.io/gorm"
)

type NotificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) *NotificationRepository {
	return &NotificationRepository{db: db}
}

func (r *NotificationRepository) CreateNotification(data *notification_service.Notification) error {
	return r.db.Create(data).Error
}

func (r *NotificationRepository) GetAllNotifications() ([]notification_service.Notification, error) {
	var result []notification_service.Notification
	err := r.db.Order("id desc").Find(&result).Error
	return result, err
}
