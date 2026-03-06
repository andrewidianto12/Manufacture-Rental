package handler

import (
	"fmt"
	"net/http"

	equipment_category_service "github.com/andrewidianto12/Manufacture-Rental/service/equipment_category"
	"github.com/labstack/echo/v4"
)

type EquipmentCategoryHandler struct {
	categoryService equipment_category_service.EquipmentCategoryService
}

func NewEquipmentCategoryHandler(categoryService equipment_category_service.EquipmentCategoryService) *EquipmentCategoryHandler {
	return &EquipmentCategoryHandler{categoryService: categoryService}
}

func (h *EquipmentCategoryHandler) CreateCategory(c echo.Context) error {
	var input equipment_category_service.CreateCategoryRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Input tidak valid", "error": err.Error()})
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Validasi gagal", "error": err.Error()})
	}

	result, err := h.categoryService.CreateCategory(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Gagal membuat category", "error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{"message": "Berhasil membuat category", "data": result})
}

func (h *EquipmentCategoryHandler) GetAllCategories(c echo.Context) error {
	result, err := h.categoryService.GetAllCategories()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal mengambil category", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"data": result})
}

func (h *EquipmentCategoryHandler) GetCategoryByID(c echo.Context) error {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "ID tidak valid", "error": err.Error()})
	}

	result, err := h.categoryService.GetCategoryByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"message": "Category tidak ditemukan", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]any{"data": result})
}

func (h *EquipmentCategoryHandler) DeleteCategory(c echo.Context) error {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "ID tidak valid", "error": err.Error()})
	}

	if err := h.categoryService.DeleteCategory(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal menghapus category", "error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]any{"message": "Berhasil menghapus category"})
}
