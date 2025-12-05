package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"net/http"

	"github.com/avivbintangaringga/ayobayar/clients/dompetkitawallet"
	"github.com/avivbintangaringga/ayobayar/config"
	"github.com/avivbintangaringga/ayobayar/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type app struct {
	addr              string
	paymentProcessors map[string]types.UpstreamPaymentProcessor
}

func main() {
	paymentProcessors := make(map[string]types.UpstreamPaymentProcessor)
	paymentProcessors["QD"] = dompetkitawallet.NewClient()

	app := &app{
		addr:              fmt.Sprintf(":%d", config.Env.Port),
		paymentProcessors: paymentProcessors,
	}

	err := startServer(app)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		os.Exit(1)
	}
}

func startServer(app *app) error {
	r := chi.NewRouter()
	r.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.RealIP,
		middleware.Compress(5),
		middleware.StripSlashes,
		middleware.Timeout(60*time.Second),
		middleware.CleanPath,
	)

	apiHandler := NewApiHandler(app)
	r.Mount("/api/v1", apiHandler)

	webHandler := NewWebHandler()
	r.Mount("/", webHandler)

	log.Println("Starting server on address", app.addr)
	return http.ListenAndServe(app.addr, r)
}
