package rental_service

type RentalRepo interface {
	CreateRental(rental *Rental) error
	GetAllRentals() ([]Rental, error)
	GetRentalByID(id uint) (*Rental, error)
	DeleteRental(id uint) error
	CountOverlappingRentals(equipmentID uint, rentalDate, returnDate string) (int64, error)
	GetEquipmentDailyRate(equipmentID uint) (float64, error)
}

type RentalService interface {
	CreateRental(input CreateRentalRequest) (*Rental, error)
	GetAllRentals() ([]Rental, error)
	GetRentalByID(id uint) (*Rental, error)
	DeleteRental(id uint) error
}
