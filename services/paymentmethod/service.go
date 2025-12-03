package paymentmethod

import (
	"github.com/avivbintangaringga/dompetkita/types"
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
