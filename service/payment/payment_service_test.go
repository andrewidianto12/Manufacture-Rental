package payment_service

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type paymentRepoMock struct {
	mock.Mock
}

func (m *paymentRepoMock) CreatePayment(payment *Payment) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *paymentRepoMock) GetAllPayments() ([]Payment, error) {
	args := m.Called()
	payments, _ := args.Get(0).([]Payment)
	return payments, args.Error(1)
}

func (m *paymentRepoMock) GetPaymentByID(id uint) (*Payment, error) {
	args := m.Called(id)
	payment, _ := args.Get(0).(*Payment)
	return payment, args.Error(1)
}

func (m *paymentRepoMock) DeletePayment(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreatePayment_Success(t *testing.T) {
	repoMock := new(paymentRepoMock)
	service := NewPaymentService(repoMock)

	input := CreatePaymentRequest{
		RentalID: 1,
		Amount:   12500000,
		Method:   "bank_transfer",
	}

	repoMock.On("CreatePayment", mock.AnythingOfType("*payment_service.Payment")).Return(nil).Once()

	result, err := service.CreatePayment(input)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, input.RentalID, result.RentalID)
	assert.Equal(t, input.Amount, result.Amount)
	assert.Equal(t, input.Method, result.Method)
	assert.Equal(t, "paid", result.Status)
	repoMock.AssertExpectations(t)
}

func TestCreatePayment_InvalidInput(t *testing.T) {
	repoMock := new(paymentRepoMock)
	service := NewPaymentService(repoMock)

	result, err := service.CreatePayment(CreatePaymentRequest{
		RentalID: 0,
		Amount:   0,
		Method:   "",
	})

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "input payment tidak valid")
	repoMock.AssertNotCalled(t, "CreatePayment", mock.Anything)
}

func TestCreatePayment_RepoError(t *testing.T) {
	repoMock := new(paymentRepoMock)
	service := NewPaymentService(repoMock)

	input := CreatePaymentRequest{
		RentalID: 1,
		Amount:   2000000,
		Method:   "cash",
	}

	repoMock.On("CreatePayment", mock.AnythingOfType("*payment_service.Payment")).Return(errors.New("db error")).Once()

	result, err := service.CreatePayment(input)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "db error")
	repoMock.AssertExpectations(t)
}

func TestGetPaymentByID_InvalidID(t *testing.T) {
	repoMock := new(paymentRepoMock)
	service := NewPaymentService(repoMock)

	result, err := service.GetPaymentByID(0)

	assert.Error(t, err)
	assert.Nil(t, result)
	assert.EqualError(t, err, "id tidak valid")
	repoMock.AssertNotCalled(t, "GetPaymentByID", mock.Anything)
}

func TestDeletePayment_InvalidID(t *testing.T) {
	repoMock := new(paymentRepoMock)
	service := NewPaymentService(repoMock)

	err := service.DeletePayment(0)

	assert.Error(t, err)
	assert.EqualError(t, err, "id tidak valid")
	repoMock.AssertNotCalled(t, "DeletePayment", mock.Anything)
}
