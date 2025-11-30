package payment

import "github.com/avivbintangaringga/dompetkita/types"

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentList() ([]types.Payment, error) {
	payments := []types.Payment{
		{
			Id:   "1",
			Desc: "payment 1",
		},
		{
			Id:   "2",
			Desc: "payment 2",
		},
	}

	return payments, nil
}
