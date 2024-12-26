package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	NbaApiKey     string
	ApiNbaUrl     string
	TelegramToken string
	MockDbData    bool
}

func LoadConfig() *Config {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Printf("Cannot read env file: %v", err)
		}
	}

	mockDbDataStr := os.Getenv("Mock_DB_data")
	mockDbDate, _ := strconv.ParseBool(mockDbDataStr)

	return &Config{
		NbaApiKey:     os.Getenv("API_NBA_key"),
		TelegramToken: os.Getenv("telegram_token"),
		ApiNbaUrl:     os.Getenv("API_NBA_URL"),
		MockDbData:    mockDbDate,
	}
}
