package equipment_category_service

type EquipmentCategoryRepo interface {
	CreateCategory(category *EquipmentCategory) error
	GetAllCategories() ([]EquipmentCategory, error)
	GetCategoryByID(id uint) (*EquipmentCategory, error)
	DeleteCategory(id uint) error
}

type EquipmentCategoryService interface {
	CreateCategory(input CreateCategoryRequest) (*EquipmentCategory, error)
	GetAllCategories() ([]EquipmentCategory, error)
	GetCategoryByID(id uint) (*EquipmentCategory, error)
	DeleteCategory(id uint) error
}
