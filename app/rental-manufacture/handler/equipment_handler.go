package handler

import (
	"fmt"
	"net/http"

	equipment_service "github.com/andrewidianto12/Manufacture-Rental/service/equipment"
	"github.com/labstack/echo/v4"
)

type EquipmentHandler struct {
	equipmentService equipment_service.EquipmentService
}

func NewEquipmentHandler(equipmentService equipment_service.EquipmentService) *EquipmentHandler {
	return &EquipmentHandler{equipmentService: equipmentService}
}

func (h *EquipmentHandler) CreateEquipment(c echo.Context) error {
	var input equipment_service.CreateEquipmentRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Input tidak valid", "error": err.Error()})
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Validasi gagal", "error": err.Error()})
	}

	result, err := h.equipmentService.CreateEquipment(input)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal membuat equipment", "error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{"message": "Berhasil membuat equipment", "data": result})
}

func (h *EquipmentHandler) GetAllEquipment(c echo.Context) error {
	result, err := h.equipmentService.GetAllEquipment()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal mengambil equipment", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"data": result})
}

func (h *EquipmentHandler) GetEquipmentByID(c echo.Context) error {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "ID tidak valid", "error": err.Error()})
	}

	result, err := h.equipmentService.GetEquipmentByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"message": "Equipment tidak ditemukan", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"data": result})
}

func (h *EquipmentHandler) DeleteEquipment(c echo.Context) error {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "ID tidak valid", "error": err.Error()})
	}

	if err := h.equipmentService.DeleteEquipment(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal menghapus equipment", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"message": "Berhasil menghapus equipment"})
}
