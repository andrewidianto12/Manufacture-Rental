package maintenance_service

type MaintenanceRepo interface {
	CreateMaintenance(data *Maintenance) error
	GetAllMaintenance() ([]Maintenance, error)
}

type MaintenanceService interface {
	CreateMaintenance(input CreateMaintenanceRequest) (*Maintenance, error)
	GetAllMaintenance() ([]Maintenance, error)
}
