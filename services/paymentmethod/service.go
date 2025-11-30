package paymentmethod

import "github.com/avivbintangaringga/dompetkita/types"

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentMethods() ([]types.PaymentMethod, error) {
	paymentMethods := []types.PaymentMethod{
		{
			Id:       "dompetsaya",
			Name:     "Dompet Saya",
			ImageUrl: "https://www.google.com",
			Category: "QRIS",
			TotalFee: 1000,
		},
	}
	return paymentMethods, nil
}
