package main

import (
	"log"
	"os"

	"github.com/avivbintangaringga/dompetkita/config"
)

type app struct {
	addr string
}

func main() {
	app := &app{
		addr: config.Env.Addr,
	}

	err := startServer(app)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		os.Exit(1)
	}
}
