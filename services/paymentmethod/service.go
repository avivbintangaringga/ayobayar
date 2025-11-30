package paymentmethod

import (
	"github.com/avivbintangaringga/dompetkita/types"
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentMethods() ([]types.PaymentMethod, error) {
	paymentMethods := []types.PaymentMethod{
		{
			Id:       "qris-dompetsaya",
			Name:     "Dompet Saya",
			ImageUrl: "https://www.google.com",
			Category: "QRIS",
			TotalFee: 1000,
		},
		{
			Id:       "bt-dompetsaya",
			Name:     "Dompet Saya",
			ImageUrl: "https://www.google.com",
			Category: "BT",
			TotalFee: 3000,
		},
	}

	return paymentMethods, nil
}
