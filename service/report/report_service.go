package report_service

type reportService struct {
	repo ReportRepo
}

func NewReportService(repo ReportRepo) ReportService {
	return &reportService{repo: repo}
}

func (s *reportService) GetDashboardReport() (*DashboardReport, error) {
	return s.repo.GetDashboardReport()
}
