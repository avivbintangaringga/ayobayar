package types

type Payment struct {
	Id   string `json:"id"`
	Desc string `json:"desc"`
}

type PaymentService interface {
	GetPaymentList() ([]Payment, error)
	GetPaymentDetail(id string) (*Payment, error)
	CreatePayment(data Payment) (*Payment, error)
}
