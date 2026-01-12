package types

import "errors"

var (
	ErrPaymentMethodNotFound = errors.New("Payment method not found")
)

type PaymentMethod struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	SmallImageUrl string `json:"small_image_url"`
	BigImageUrl   string `json:"big_image_url"`
	TotalFee      int    `json:"total_fee"`
	Category      string `json:"category"`
	IsAvailable   bool   `json:"is_available"`
}

type PaymentMethodService interface {
	GetPaymentMethods() ([]PaymentMethod, error)
	GetPaymentMethodById(id string) (*PaymentMethod, error)
}

type PaymentMethodRepository interface {
	// Create(data PaymentMethod) (*PaymentMethod, error)
	FindById(id string) (*PaymentMethod, error)
	// Update(id string, data PaymentMethod) (*PaymentMethod, error)
	// Delete(id string) error
	List() ([]PaymentMethod, error)
}
