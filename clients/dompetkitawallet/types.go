package dompetkitawallet

import "net/http"

type Client struct {
	client  *http.Client
	baseUrl string
}

type DompetKitaWalletResponse struct {
	PaymentId  string `json:"payment_id"`
	PaymentUrl string `json:"payment_url"`
	Status     string `json:"status"`
}
