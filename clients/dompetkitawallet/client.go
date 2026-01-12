package dompetkitawallet

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/avivbintangaringga/ayobayar/config"
	"github.com/avivbintangaringga/ayobayar/types"
)

type Client struct {
	client  *http.Client
	baseUrl string
}

type Response struct {
	PaymentId  string `json:"payment_id"`
	PaymentUrl string `json:"payment_url"`
	Status     string `json:"status"`
}

func NewClient() *Client {
	baseUrl := "https://prod.dompetkita.my.id/"
	if config.Env.DevMode {
		baseUrl = "https://dev.dompetkita.my.id/"
	}

	return &Client{
		baseUrl: baseUrl,
		client: &http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) RequestPayment(req types.Payment) (types.UpstreamPaymentResult, error) {
	data := []string{}
	payload := new(bytes.Buffer)

	err := json.NewEncoder(payload).Encode(data)
	if err != nil {
		return types.UpstreamPaymentResult{}, err
	}

	resp, err := c.client.Post(c.baseUrl+"/api/v1/payment", "application/json", payload)
	if err != nil {
		return types.UpstreamPaymentResult{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return types.UpstreamPaymentResult{}, errors.New("failed to send payment request")
	}

	var content Response
	err = json.NewDecoder(resp.Body).Decode(&content)
	if err != nil {
		return types.UpstreamPaymentResult{}, err
	}

	result := types.UpstreamPaymentResult{
		PaymentId:  content.PaymentId,
		PaymentUrl: content.PaymentUrl,
		QrContent:  "",
		Status:     content.Status,
	}

	return result, nil
}

func (c *Client) IsPaymentSuccess(paymentId string) (bool, error) {
	// TODO: Implement
	return true, nil
}

func (c *Client) AcknowledgePayment(paymentId string) error {
	// TODO: Implement
	return nil
}

func (c *Client) GetPaymentResult(paymentId string) (types.UpstreamPaymentResult, error) {
	// TODO: Implement
	return types.UpstreamPaymentResult{
		PaymentId:  "test-id",
		PaymentUrl: "https://dompetkita.my.id/payment/test-id",
		QrContent:  "DUMMY QR CONTENT",
		Status:     "PENDING",
	}, nil
}
