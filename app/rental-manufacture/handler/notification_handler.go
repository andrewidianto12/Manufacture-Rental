package handler

import (
	"net/http"

	notification_service "github.com/andrewidianto12/Manufacture-Rental/service/notification"
	"github.com/labstack/echo/v4"
)

type NotificationHandler struct {
	notificationService notification_service.NotificationService
}

func NewNotificationHandler(notificationService notification_service.NotificationService) *NotificationHandler {
	return &NotificationHandler{notificationService: notificationService}
}

func (h *NotificationHandler) CreateNotification(c echo.Context) error {
	var input notification_service.CreateNotificationRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Input tidak valid", "error": err.Error()})
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Validasi gagal", "error": err.Error()})
	}

	result, err := h.notificationService.CreateNotification(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Gagal membuat notification", "error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]any{"message": "Berhasil membuat notification", "data": result})
}

func (h *NotificationHandler) GetAllNotifications(c echo.Context) error {
	result, err := h.notificationService.GetAllNotifications()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal mengambil notification", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"data": result})
}
