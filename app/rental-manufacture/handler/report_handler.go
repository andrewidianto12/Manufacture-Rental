package handler

import (
	"net/http"

	report_service "github.com/andrewidianto12/Manufacture-Rental/service/report"
	"github.com/labstack/echo/v4"
)

type ReportHandler struct {
	reportService report_service.ReportService
}

func NewReportHandler(reportService report_service.ReportService) *ReportHandler {
	return &ReportHandler{reportService: reportService}
}

func (h *ReportHandler) GetDashboardReport(c echo.Context) error {
	result, err := h.reportService.GetDashboardReport()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{"message": "Gagal mengambil report", "error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]any{"data": result})
}
