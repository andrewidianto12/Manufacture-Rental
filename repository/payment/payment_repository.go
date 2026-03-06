package payment

import (
	payment_service "github.com/andrewidianto12/Manufacture-Rental/service/payment"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) CreatePayment(payment *payment_service.Payment) error {
	return r.db.Create(payment).Error
}

func (r *PaymentRepository) GetAllPayments() ([]payment_service.Payment, error) {
	var payments []payment_service.Payment
	err := r.db.Order("id desc").Find(&payments).Error
	return payments, err
}

func (r *PaymentRepository) GetPaymentByID(id uint) (*payment_service.Payment, error) {
	var payment payment_service.Payment
	if err := r.db.First(&payment, id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *PaymentRepository) DeletePayment(id uint) error {
	return r.db.Delete(&payment_service.Payment{}, id).Error
}
