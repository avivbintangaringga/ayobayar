package main

import (
	"log"
	"net/http"

	"github.com/avivbintangaringga/dompetkita/services/payment"
	"github.com/avivbintangaringga/dompetkita/services/paymentmethod"
	"github.com/go-chi/chi/v5"
)

func registerRoutes(r *chi.Mux) {
	r.Route("/api/v1", func(r chi.Router) {
		// Payments
		paymentService := payment.NewService()
		paymentHandler := payment.NewHandler(paymentService)
		r.Get("/payments", paymentHandler.ListPayments)

		// Payment Methods
		paymentMethodService := paymentmethod.NewService()
		paymentMethodHandler := paymentmethod.NewHandler(paymentMethodService)
		r.Get("/paymentmethods", paymentMethodHandler.GetPaymentMethods)
	})
}

func startServer(app *app) error {
	r := chi.NewRouter()
	registerRoutes(r)
	log.Println("Starting server on address", app.addr)
	return http.ListenAndServe(app.addr, r)
}
