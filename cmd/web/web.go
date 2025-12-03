package web

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Newhandler() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	return r
}
