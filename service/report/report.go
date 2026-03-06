package report_service

type DashboardReport struct {
	TotalUsers         int64   `json:"total_users"`
	TotalEquipment     int64   `json:"total_equipment"`
	TotalActiveRentals int64   `json:"total_active_rentals"`
	TotalRevenue       float64 `json:"total_revenue"`
}
