package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	godotenv.Load()

}

func getEnvVar(key string) string {
	return os.Getenv(key)
}
