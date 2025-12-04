package payment

import (
	"github.com/avivbintangaringga/ayobayar/types"
	"github.com/teris-io/shortid"
)

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

type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) Create(data types.Payment) (*types.Payment, error) {
	payment := data
	payment.Id = shortid.MustGenerate()
	payment.Status = "PENDING"
	payments = append(payments, payment)

	return &payment, nil
}

func (r *Repository) FindById(id string) (*types.Payment, error) {
	for _, p := range payments {
		if p.Id == id {
			return &p, nil
		}
	}

	return nil, types.ErrPaymentNotFound
}

func (r *Repository) Update(id string, data types.Payment) (*types.Payment, error) {
	return nil, nil
}
func (r *Repository) Delete(id string) error {
	return nil
}
func (r *Repository) List() ([]types.Payment, error) {
	return payments, nil
}
