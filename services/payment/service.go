package payment

import (
	"github.com/avivbintangaringga/dompetkita/types"
)

type Service struct {
	paymentRepo       types.PaymentRepository
	paymentMethodRepo types.PaymentMethodRepository
}

func NewService(paymentRepo types.PaymentRepository, paymentMethodRepo types.PaymentMethodRepository) *Service {
	return &Service{
		paymentRepo:       paymentRepo,
		paymentMethodRepo: paymentMethodRepo,
	}
}

func (s *Service) GetPaymentList() ([]types.Payment, error) {
	return s.paymentRepo.List()
}

func (s *Service) GetPaymentDetail(id string) (*types.Payment, error) {
	return s.paymentRepo.FindById(id)
}

func (s *Service) CreatePayment(data types.PaymentRequest) (*types.Payment, error) {
	payment := types.Payment{
		Desc:   data.Desc,
		Amount: data.Amount,
	}
	return s.paymentRepo.Create(payment)
}
