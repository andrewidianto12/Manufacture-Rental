package equipment_service

type EquipmentRepo interface {
	CreateEquipment(equipment *Equipment) error
	GetAllEquipment() ([]Equipment, error)
	GetEquipmentByID(id uint) (*Equipment, error)
	DeleteEquipment(id uint) error
}

type EquipmentService interface {
	CreateEquipment(input CreateEquipmentRequest) (*Equipment, error)
	GetAllEquipment() ([]Equipment, error)
	GetEquipmentByID(id uint) (*Equipment, error)
	DeleteEquipment(id uint) error
}
