package payment

import (
	"errors"

	"github.com/avivbintangaringga/ayobayar/types"
)

type Service struct {
	paymentRepo       types.PaymentRepository
	paymentMethodRepo types.PaymentMethodRepository
	paymentProcessors map[string]types.UpstreamPaymentProcessor
}

func NewService(
	paymentRepo types.PaymentRepository,
	paymentMethodRepo types.PaymentMethodRepository,
	paymentProcessors map[string]types.UpstreamPaymentProcessor,
) *Service {
	return &Service{
		paymentRepo:       paymentRepo,
		paymentMethodRepo: paymentMethodRepo,
		paymentProcessors: paymentProcessors,
	}
}

func (s *Service) GetPaymentList() ([]types.Payment, error) {
	return s.paymentRepo.List()
}

func (s *Service) GetPaymentDetail(id string) (*types.Payment, error) {
	return s.paymentRepo.FindById(id)
}

func (s *Service) CreatePayment(data types.Payment) (*types.Payment, *types.UpstreamPaymentResult, error) {
	// Check if payment processor is available
	pp := s.paymentProcessors[data.PaymentMethodId]
	if pp == nil {
		return nil, nil, errors.New("payment method is not available")
	}

	// Request payment processor to create payment
	res, err := pp.RequestPayment(data)
	if err != nil {
		return nil, nil, err
	}

	payment, err := s.paymentRepo.Create(data)
	if err != nil {
		return nil, nil, err
	}

	return payment, &res, nil
}
