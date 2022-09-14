package config

import (
	"os"

	"github.com/joho/godotenv"
)

func GetEnvValue(key string) string {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}
