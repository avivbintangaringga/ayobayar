package types

type Payment struct {
	Id   string `json:"id"`
	Desc string `json:"desc"`
}

type PaymentService interface {
	GetPaymentList() ([]Payment, error)
}
