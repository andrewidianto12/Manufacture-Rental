package rental

import (
	rental_service "github.com/andrewidianto12/Manufacture-Rental/service/rental"
	"gorm.io/gorm"
)

type RentalRepository struct {
	db *gorm.DB
}

func NewRentalRepository(db *gorm.DB) *RentalRepository {
	return &RentalRepository{db: db}
}

func (r *RentalRepository) CreateRental(rental *rental_service.Rental) error {
	return r.db.Create(rental).Error
}

func (r *RentalRepository) GetAllRentals() ([]rental_service.Rental, error) {
	var rentals []rental_service.Rental
	err := r.db.Order("id desc").Find(&rentals).Error
	return rentals, err
}

func (r *RentalRepository) GetRentalByID(id uint) (*rental_service.Rental, error) {
	var rental rental_service.Rental
	if err := r.db.First(&rental, id).Error; err != nil {
		return nil, err
	}
	return &rental, nil
}

func (r *RentalRepository) DeleteRental(id uint) error {
	return r.db.Delete(&rental_service.Rental{}, id).Error
}

func (r *RentalRepository) CountOverlappingRentals(equipmentID uint, rentalDate, returnDate string) (int64, error) {
	var count int64
	err := r.db.Model(&rental_service.Rental{}).
		Where("equipment_id = ? AND status = ? AND rental_date < ? AND return_date > ?", equipmentID, "active", returnDate, rentalDate).
		Count(&count).Error
	return count, err
}

func (r *RentalRepository) GetEquipmentDailyRate(equipmentID uint) (float64, error) {
	var dailyRate float64
	err := r.db.Table("equipment").Select("daily_rate").Where("id = ?", equipmentID).Scan(&dailyRate).Error
	return dailyRate, err
}
