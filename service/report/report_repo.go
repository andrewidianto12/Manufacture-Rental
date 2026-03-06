package report_service

type ReportRepo interface {
	GetDashboardReport() (*DashboardReport, error)
}

type ReportService interface {
	GetDashboardReport() (*DashboardReport, error)
}
