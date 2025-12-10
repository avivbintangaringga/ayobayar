package main

import (
	"net/http"

	"github.com/avivbintangaringga/ayobayar/types"
	"github.com/avivbintangaringga/ayobayar/web/paymentlistpage"
	"github.com/avivbintangaringga/ayobayar/web/static"
	"github.com/go-chi/chi/v5"
)

func NewWebHandler(app *app) http.Handler {
	r := chi.NewRouter()

	staticPrefix := "/static/"
	staticHandler := static.NewHandler(app.staticFiles, staticPrefix)
	r.Get("/static/*", staticHandler.ServeStatic)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		payments := []types.Payment{}
		paymentlistpage.Page(payments).Render(r.Context(), w)
	})

	return r
}
