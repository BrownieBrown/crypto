package config

import (
	"crypto/pkg/logger"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		logger.Log.Error("Error loading .env file")
	}
}

func GetEnvVar(key string) string {
	return os.Getenv(key)
}
