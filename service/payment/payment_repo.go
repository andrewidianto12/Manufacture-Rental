package payment_service

type PaymentRepo interface {
	CreatePayment(payment *Payment) error
	GetAllPayments() ([]Payment, error)
	GetPaymentByID(id uint) (*Payment, error)
	DeletePayment(id uint) error
}

type PaymentService interface {
	CreatePayment(input CreatePaymentRequest) (*Payment, error)
	GetAllPayments() ([]Payment, error)
	GetPaymentByID(id uint) (*Payment, error)
	DeletePayment(id uint) error
}
