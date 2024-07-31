package utils

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

func LoadEnvVars() {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading variables from .env file", "err", err.Error())
		os.Exit(1)
		return
	}
}
