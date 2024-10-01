package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var CFG *Configuration

func Init() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	CFG = &Configuration{
		os.Getenv("DB_URL"),
	}

	return nil
}
