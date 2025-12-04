package api

import (
	"net/http"

	"github.com/avivbintangaringga/ayobayar/services/common"
	"github.com/avivbintangaringga/ayobayar/services/payment"
	"github.com/avivbintangaringga/ayobayar/services/paymentmethod"
	"github.com/go-chi/chi/v5"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()

	// Common
	commonHandler := common.NewHandler()
	r.NotFound(commonHandler.HandleNotFound)
	r.MethodNotAllowed(commonHandler.HandleNotAllowed)

	// Payment Methods
	paymentMethodRepo := paymentmethod.NewRepository()
	paymentMethodService := paymentmethod.NewService(paymentMethodRepo)
	paymentMethodHandler := paymentmethod.NewHandler(paymentMethodService)
	r.Get("/paymentmethods", paymentMethodHandler.GetPaymentMethods)

	// Payments
	paymentRepo := payment.NewRepository()
	paymentService := payment.NewService(paymentRepo, paymentMethodRepo)
	paymentHandler := payment.NewHandler(paymentService)
	r.Get("/payments/{id}", paymentHandler.GetPaymentDetail)
	r.Get("/payments", paymentHandler.ListPayments)
	r.Post("/payments", paymentHandler.PostPayment)

	return r
}
