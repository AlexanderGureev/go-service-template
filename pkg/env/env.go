package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Load() {
	if err := godotenv.Load(); err != nil {
		panic(".env file is not found")
	}
}

func Get(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		return defaultValue
	}

	return val
}
