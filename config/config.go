package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	Port        int
	DatabaseUrl string
	DevMode     bool
}

var Env = loadEnv()

func loadEnv() *config {
	godotenv.Load()

	return &config{
		Port:        getEnvInt("PORT", 8080),
		DatabaseUrl: getEnvString("DATABASE_URL", ""),
		DevMode:     getEnvString("DEV_MODE", "true") == "true",
	}
}

func getEnvString(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return intValue
}
