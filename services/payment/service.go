package payment

import (
	"errors"

	"github.com/avivbintangaringga/dompetkita/types"
)

type service struct{}

var payments = []types.Payment{
	{
		Id:     "1",
		Desc:   "payment 1",
		Amount: 15000,
		Status: "SUCCESS",
	},
	{
		Id:     "2",
		Desc:   "payment 2",
		Amount: 143000,
		Status: "PENDING",
	},
}

func NewService() *service {
	return &service{}
}

func (s *service) GetPaymentList() ([]types.Payment, error) {
	return payments, nil
}

func (s *service) GetPaymentDetail(id string) (*types.Payment, error) {
	var data *types.Payment

	for _, i := range payments {
		if i.Id == id {
			data = &i
			break
		}
	}

	if data == nil {
		return nil, errors.New("item not found")
	}

	return data, nil
}

func (s *service) CreatePayment(data types.Payment) (*types.Payment, error) {
	return nil, nil
}
