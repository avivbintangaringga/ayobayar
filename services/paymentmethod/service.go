package paymentmethod

import (
	"github.com/avivbintangaringga/ayobayar/types"
)

type Service struct {
	repository types.PaymentMethodRepository
}

func NewService(repository types.PaymentMethodRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetPaymentMethods() ([]types.PaymentMethod, error) {
	return s.repository.List()
}

func (s *Service) GetPaymentMethodById(id string) (*types.PaymentMethod, error) {
	return s.repository.FindById(id)
}
