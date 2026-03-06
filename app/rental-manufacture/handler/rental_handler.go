package handler

import (
	"fmt"
	"net/http"

	rental_service "github.com/andrewidianto12/Manufacture-Rental/service/rental"
	"github.com/labstack/echo/v4"
)

type RentalHandler struct {
	rentalService rental_service.RentalService
}

func NewRentalHandler(rentalService rental_service.RentalService) *RentalHandler {
	return &RentalHandler{rentalService: rentalService}
}

func (h *RentalHandler) CreateRental(c echo.Context) error {
	var input rental_service.CreateRentalRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Input tidak valid", "error": err.Error()})
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Validasi gagal", "error": err.Error()})
	}

	result, err := h.rentalService.CreateRental(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Gagal membuat rental", "error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{"message": "Berhasil membuat rental", "data": result})
}

func (h *RentalHandler) GetAllRentals(c echo.Context) error {
	result, err := h.rentalService.GetAllRentals()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal mengambil data rental", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"data": result})
}

func (h *RentalHandler) GetRentalByID(c echo.Context) error {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "ID tidak valid", "error": err.Error()})
	}

	result, err := h.rentalService.GetRentalByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"message": "Rental tidak ditemukan", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"data": result})
}

func (h *RentalHandler) DeleteRental(c echo.Context) error {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "ID tidak valid", "error": err.Error()})
	}

	if err := h.rentalService.DeleteRental(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal menghapus rental", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"message": "Berhasil menghapus rental"})
}
