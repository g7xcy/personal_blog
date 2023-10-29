package util

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var Env map[string]string

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}
	Env = map[string]string{
		"DB_USER": "",
		"DB_PWD":  "",
		"DB_URL":  "",
		"DB_PORT": "",
		"DB":      "",
		"PORT":    "",
		"URL":     "",
	}

	for k := range Env {
		Env[k] = os.Getenv(k)
		if Env[k] == "" {
			return fmt.Errorf("environment variable %q is missing", k)
		}
	}

	return nil
}
