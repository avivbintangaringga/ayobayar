package types

type PaymentMethod struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
	TotalFee int    `json:"total_fee"`
	Category string `json:"category"`
}

type PaymentMethodService interface {
	GetPaymentMethods() ([]PaymentMethod, error)
}
