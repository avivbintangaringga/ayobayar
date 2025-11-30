package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Addr string
}

var Env = loadEnv()

func loadEnv() *config {
	godotenv.Load()

	return &config{
		Addr: getEnvString("ADDR", ":8080"),
	}
}

func getEnvString(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
