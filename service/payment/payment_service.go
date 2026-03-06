package payment_service

import (
	"errors"
	"strings"
	"time"
)

type CreatePaymentRequest struct {
	RentalID uint    `json:"rental_id" validate:"required,gte=1"`
	Amount   float64 `json:"amount" validate:"required,gt=0"`
	Method   string  `json:"method" validate:"required"`
}

type paymentService struct {
	repo PaymentRepo
}

func NewPaymentService(repo PaymentRepo) PaymentService {
	return &paymentService{repo: repo}
}

func (s *paymentService) CreatePayment(input CreatePaymentRequest) (*Payment, error) {
	if input.RentalID == 0 || input.Amount <= 0 || strings.TrimSpace(input.Method) == "" {
		return nil, errors.New("input payment tidak valid")
	}

	paidAt := time.Now()
	payment := &Payment{
		RentalID: input.RentalID,
		Amount:   input.Amount,
		Method:   input.Method,
		Status:   "paid",
		PaidAt:   &paidAt,
	}

	if err := s.repo.CreatePayment(payment); err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *paymentService) GetAllPayments() ([]Payment, error) {
	return s.repo.GetAllPayments()
}

func (s *paymentService) GetPaymentByID(id uint) (*Payment, error) {
	if id == 0 {
		return nil, errors.New("id tidak valid")
	}
	return s.repo.GetPaymentByID(id)
}

func (s *paymentService) DeletePayment(id uint) error {
	if id == 0 {
		return errors.New("id tidak valid")
	}
	return s.repo.DeletePayment(id)
}
