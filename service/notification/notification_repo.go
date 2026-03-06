package notification_service

type NotificationRepo interface {
	CreateNotification(data *Notification) error
	GetAllNotifications() ([]Notification, error)
}

type NotificationService interface {
	CreateNotification(input CreateNotificationRequest) (*Notification, error)
	GetAllNotifications() ([]Notification, error)
}
