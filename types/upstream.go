package types

type UpstreamPaymentResult struct {
	PaymentId  string `json:"payment_id"`
	PaymentUrl string `json:"payment_url"`
	QrContent  string `json:"qr_content"`
	Status     string `json:"status"`
}

type UpstreamPaymentProcessor interface {
	RequestPayment(req Payment) (UpstreamPaymentResult, error)
	IsPaymentSuccess(paymentId string) (bool, error)
	AcknowledgePayment(paymentId string) error
}
