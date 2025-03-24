package config

import (
	"os"

	"github.com/joho/godotenv"
)

func loadEnvVars() {
	godotenv.Load()
}

func getEnvVar(key string) string {
	return os.Getenv(key)
}
