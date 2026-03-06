package handler

import (
	"fmt"
	"net/http"

	payment_service "github.com/andrewidianto12/Manufacture-Rental/service/payment"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	paymentService payment_service.PaymentService
}

func NewPaymentHandler(paymentService payment_service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (h *PaymentHandler) CreatePayment(c echo.Context) error {
	var input payment_service.CreatePaymentRequest
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Input tidak valid", "error": err.Error()})
	}
	if err := c.Validate(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Validasi gagal", "error": err.Error()})
	}

	result, err := h.paymentService.CreatePayment(input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "Gagal membuat payment", "error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]any{"message": "Berhasil membuat payment", "data": result})
}

func (h *PaymentHandler) GetAllPayments(c echo.Context) error {
	result, err := h.paymentService.GetAllPayments()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal mengambil data payment", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"data": result})
}

func (h *PaymentHandler) GetPaymentByID(c echo.Context) error {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "ID tidak valid", "error": err.Error()})
	}

	result, err := h.paymentService.GetPaymentByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{"message": "Payment tidak ditemukan", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"data": result})
}

func (h *PaymentHandler) DeletePayment(c echo.Context) error {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{"message": "ID tidak valid", "error": err.Error()})
	}

	if err := h.paymentService.DeletePayment(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal menghapus payment", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"message": "Berhasil menghapus payment"})
}
