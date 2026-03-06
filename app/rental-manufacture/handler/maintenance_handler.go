package handler

import (
	"net/http"

	maintenance_service "github.com/andrewidianto12/Manufacture-Rental/service/maintenance"
	"github.com/labstack/echo/v4"
)

type MaintenanceHandler struct {
	maintenanceService maintenance_service.MaintenanceService
}

func NewMaintenanceHandler(maintenanceService maintenance_service.MaintenanceService) *MaintenanceHandler {
	return &MaintenanceHandler{maintenanceService: maintenanceService}
}

func (h *MaintenanceHandler) CreateMaintenance(c echo.Context) error {
	var input maintenance_service.CreateMaintenanceRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Input tidak valid", "error": err.Error()})
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Validasi gagal", "error": err.Error()})
	}

	result, err := h.maintenanceService.CreateMaintenance(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Gagal membuat maintenance", "error": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]any{"message": "Berhasil membuat maintenance", "data": result})
}

func (h *MaintenanceHandler) GetAllMaintenance(c echo.Context) error {
	result, err := h.maintenanceService.GetAllMaintenance()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal mengambil maintenance", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"data": result})
}
