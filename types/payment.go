package types

import (
	"errors"
	"time"
)

var (
	ErrPaymentNotFound = errors.New("Payment not found")
)

type Payment struct {
	Id              string    `json:"id"`
	PaymentMethodId string    `json:"payment_method_id"`
	Amount          int64     `json:"amount"`
	Status          string    `json:"status"`
	ExpiryTime      time.Time `json:"expiry_time"`
	CallbackUrl     string    `json:"callback_url"`
	RedirectUrl     string    `json:"redirect_url"`
	MerchantId      string    `json:"merchant_id"`
	MerchantOrderId string    `json:"merchant_order_id"`
	CustomerEmail   string    `json:"customer_email"`
	CustomerName    string    `json:"customer_name"`
	CustomerPhone   string    `json:"customer_phone"`
	ProductDetails  string    `json:"product_details"`
	CreatedAt       time.Time `json:"created_at"`
}

type PaymentRequest struct {
	PaymentMethodId string `json:"payment_method_id" validate:"required,alpha,len=2,uppercase"`
	Amount          int64  `json:"amount" validate:"required,min=1000,max=100000000"`
	ExpiryTime      string `json:"expiry_time" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	CallbackUrl     string `json:"callback_url" validate:"required,url,min=1,max=1024"`
	RedirectUrl     string `json:"redirect_url" validate:"required,url,min=1,max=1024"`
	MerchantId      string `json:"merchant_id" validate:"required,min=1,max=255"`
	MerchantOrderId string `json:"merchant_order_id" validate:"required,min=1,max=255"`
	CustomerEmail   string `json:"customer_email" validate:"required,email,min=1,max=255"`
	CustomerName    string `json:"customer_name" validate:"required,min=1,max=255"`
	CustomerPhone   string `json:"customer_phone" validate:"required,min=5,max=20"`
	ProductDetails  string `json:"product_details" validate:"required,min=1,max=255"`
}

type PaymentResponse struct {
	PaymentUrl  string  `json:"payment_url"`
	QrString    string  `json:"qr_string"`
	PaymentData Payment `json:"payment_data"`
}

type PaymentService interface {
	GetPaymentList() ([]Payment, error)
	GetPaymentDetail(id string) (*Payment, error)
	CreatePayment(data Payment) (*Payment, *UpstreamPaymentResult, error)
}

type PaymentRepository interface {
	Create(data Payment) (*Payment, error)
	FindById(id string) (*Payment, error)
	Update(id string, data Payment) (*Payment, error)
	Delete(id string) error
	List() ([]Payment, error)
}
