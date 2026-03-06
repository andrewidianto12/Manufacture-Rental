package rental_service

import (
	"errors"
	"time"
)

type CreateRentalRequest struct {
	UserID      uint   `json:"user_id" validate:"required,gte=1"`
	EquipmentID uint   `json:"equipment_id" validate:"required,gte=1"`
	RentalDate  string `json:"rental_date" validate:"required"`
	ReturnDate  string `json:"return_date" validate:"required"`
}

type rentalService struct {
	repo RentalRepo
}

func NewRentalService(repo RentalRepo) RentalService {
	return &rentalService{repo: repo}
}

func (s *rentalService) CreateRental(input CreateRentalRequest) (*Rental, error) {
	startDate, err := time.Parse("2006-01-02", input.RentalDate)
	if err != nil {
		return nil, errors.New("format rental_date harus YYYY-MM-DD")
	}

	endDate, err := time.Parse("2006-01-02", input.ReturnDate)
	if err != nil {
		return nil, errors.New("format return_date harus YYYY-MM-DD")
	}

	if !endDate.After(startDate) {
		return nil, errors.New("return_date harus setelah rental_date")
	}

	overlapCount, err := s.repo.CountOverlappingRentals(input.EquipmentID, input.RentalDate, input.ReturnDate)
	if err != nil {
		return nil, err
	}
	if overlapCount > 0 {
		return nil, errors.New("equipment sedang tidak tersedia di rentang tanggal tersebut")
	}

	dailyRate, err := s.repo.GetEquipmentDailyRate(input.EquipmentID)
	if err != nil {
		return nil, err
	}

	durationDays := int(endDate.Sub(startDate).Hours() / 24)
	if durationDays <= 0 {
		durationDays = 1
	}

	rental := &Rental{
		UserID:      input.UserID,
		EquipmentID: input.EquipmentID,
		RentalDate:  startDate,
		ReturnDate:  endDate,
		TotalCost:   float64(durationDays) * dailyRate,
		Status:      "active",
	}

	if err := s.repo.CreateRental(rental); err != nil {
		return nil, err
	}

	return rental, nil
}

func (s *rentalService) GetAllRentals() ([]Rental, error) {
	return s.repo.GetAllRentals()
}

func (s *rentalService) GetRentalByID(id uint) (*Rental, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}
	return s.repo.GetRentalByID(id)
}

func (s *rentalService) DeleteRental(id uint) error {
	if id == 0 {
		return errors.New("id tidak valid")
	}
	return s.repo.DeleteRental(id)
}
