package util

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type envError struct {
	key   string
	error string
}

func (e envError) Error() string {
	return fmt.Sprintf("%v: %v", e.key, e.error)
}

func LoadFromEnv(key string) (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	value := os.Getenv(key)
	if value == "" {
		return value, envError{key, "No value found in env"}
	}

	return value, nil
}
