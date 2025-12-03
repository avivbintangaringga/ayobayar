package types

import "errors"

var (
	ErrPaymentNotFound = errors.New("Payment not found")
)

type Payment struct {
	Id              string `json:"id"`
	Amount          int    `json:"amount"`
	Desc            string `json:"desc"`
	Status          string `json:"status"`
	CallbackUrl     string `json:"callback_url"`
	MerchantId      string `json:"merchant_id"`
	MerchantOrderId string `json:"merchant_order_id"`
	UserEmail       string `json:"email"`
	UserName        string `json:"name"`
	PaymentMethodId string `json:"payment_method_id"`
}

type PaymentRequest struct {
	Desc            string `json:"desc" validate:"required,min=1,max=255"`
	Amount          int    `json:"amount" validate:"required,min=1000"`
	CallbackUrl     string `json:"callback_url" validate:"required,url"`
	MerchantId      string `json:"merchant_id" validate:"required,min=1,max=255"`
	MerchantOrderId string `json:"merchant_order_id" validate:"required,min=1,max=255"`
	UserEmail       string `json:"email" validate:"required,email"`
	UserName        string `json:"name" validate:"required,min=1,max=255"`
	PaymentMethodId string `json:"payment_method_id" validate:"required,min=1,max=255"`
}

type PaymentService interface {
	GetPaymentList() ([]Payment, error)
	GetPaymentDetail(id string) (*Payment, error)
	CreatePayment(data Payment) (*Payment, error)
}

type PaymentRepository interface {
	Create(data Payment) (*Payment, error)
	FindById(id string) (*Payment, error)
	Update(id string, data Payment) (*Payment, error)
	Delete(id string) error
	List() ([]Payment, error)
}
