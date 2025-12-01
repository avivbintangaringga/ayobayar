package main

import (
	"log"
	"net/http"

	"github.com/avivbintangaringga/dompetkita/services/common"
	"github.com/avivbintangaringga/dompetkita/services/payment"
	"github.com/avivbintangaringga/dompetkita/services/paymentmethod"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func registerRoutes(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.RealIP)
	r.Route("/api/v1", func(r chi.Router) {
		// Common
		commonHandler := common.NewHandler()
		r.NotFound(commonHandler.HandleNotFound)
		r.MethodNotAllowed(commonHandler.HandleNotAllowed)

		// Payments
		paymentService := payment.NewService()
		paymentHandler := payment.NewHandler(paymentService)
		r.Get("/payments/{id}", paymentHandler.GetPaymentDetail)
		r.Get("/payments", paymentHandler.ListPayments)
		r.Post("/payments", paymentHandler.PostPayment)

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
