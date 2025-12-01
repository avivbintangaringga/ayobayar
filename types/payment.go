package types

type Payment struct {
	Id     string `json:"id"`
	Amount int    `json:"amount"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}

type PaymentRequest struct {
	Desc        string `json:"desc"`
	Amount      int    `json:"amount"`
	CallbackUrl string `json:"callback_url"`
}

type PaymentService interface {
	GetPaymentList() ([]Payment, error)
	GetPaymentDetail(id string) (*Payment, error)
	CreatePayment(data Payment) (*Payment, error)
}
