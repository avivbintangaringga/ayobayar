package main

import (
	"net/http"

	"github.com/avivbintangaringga/ayobayar/services/payment"
	"github.com/avivbintangaringga/ayobayar/services/paymentmethod"
	"github.com/avivbintangaringga/ayobayar/web/paymentlistpage"
	"github.com/avivbintangaringga/ayobayar/web/paymentpage"
	"github.com/avivbintangaringga/ayobayar/web/static"
	"github.com/go-chi/chi/v5"
)

func NewWebHandler(app *app) http.Handler {
	r := chi.NewRouter()

	staticPrefix := "/static/"
	staticHandler := static.NewHandler(app.staticFiles, staticPrefix)
	r.Get("/static/*", staticHandler.ServeStatic)

	paymentMethodRepo := paymentmethod.NewRepository(app.db)
	paymentRepo := payment.NewRepository(app.db)
	paymentService := payment.NewService(paymentRepo, paymentMethodRepo, app.paymentProcessors)

	paymentListPage := paymentlistpage.NewHandler(paymentService)
	r.Get("/", paymentListPage.Handle)

	paymentMethodService := paymentmethod.NewService(paymentMethodRepo)
	paymentPage := paymentpage.NewHandler(paymentService, paymentMethodService, app.paymentProcessors)
	r.Get("/payment/{id}", paymentPage.Handle)

	return r
}
