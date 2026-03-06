package notification_service

import "errors"

type CreateNotificationRequest struct {
	UserID    uint   `json:"user_id" validate:"required,gte=1"`
	Channel   string `json:"channel" validate:"required,oneof=email sms whatsapp"`
	Recipient string `json:"recipient" validate:"required"`
	Subject   string `json:"subject"`
	Message   string `json:"message" validate:"required"`
}

type notificationService struct {
	repo NotificationRepo
}

func NewNotificationService(repo NotificationRepo) NotificationService {
	return &notificationService{repo: repo}
}

func (s *notificationService) CreateNotification(input CreateNotificationRequest) (*Notification, error) {
	if input.UserID == 0 || input.Recipient == "" || input.Message == "" {
		return nil, errors.New("input notification tidak valid")
	}

	data := &Notification{
		UserID:    input.UserID,
		Channel:   input.Channel,
		Recipient: input.Recipient,
		Subject:   input.Subject,
		Message:   input.Message,
		Status:    "sent",
	}

	if err := s.repo.CreateNotification(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (s *notificationService) GetAllNotifications() ([]Notification, error) {
	return s.repo.GetAllNotifications()
}
