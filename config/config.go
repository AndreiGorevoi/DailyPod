package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	API_NBA_key string
}

func LoadConfig() *Config {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Printf("Cannot read env file: %v", err)
		}
	}

	return &Config{
		API_NBA_key: os.Getenv("API_NBA_key"),
	}
}
