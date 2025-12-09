package main

import (
	"net/http"

	"github.com/avivbintangaringga/ayobayar/web/basepage"
	"github.com/go-chi/chi/v5"
)

func NewWebHandler() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		basepage.BasePage("HELLO WORLD").Render(r.Context(), w)
	})

	return r
}
