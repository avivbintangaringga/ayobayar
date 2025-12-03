package paymentmethod

import "github.com/avivbintangaringga/dompetkita/types"

type Repository struct {
	paymentMethods []types.PaymentMethod
}

func NewRepository() *Repository {
	return &Repository{
		paymentMethods: []types.PaymentMethod{
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
		},
	}
}

func (r *Repository) List() ([]types.PaymentMethod, error) {
	return r.paymentMethods, nil
}

func (r *Repository) FindById(id string) (*types.PaymentMethod, error) {
	for _, pm := range r.paymentMethods {
		if pm.Id == id {
			return &pm, nil
		}
	}
	return nil, types.ErrPaymentMethodNotFound
}
