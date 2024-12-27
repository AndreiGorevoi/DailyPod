package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	NbaApiKey     string
	ApiNbaUrl     string `json:"api_nba_url"`
	TelegramToken string
	MockDbData    bool `json:"mock_db_data"`
}

func LoadConfig() *Config {
	env := determineEnv()
	cfg := readConfigFile(env)

	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			log.Printf("Cannot read env file: %v", err)
		}
	}
	cfg.NbaApiKey = os.Getenv("API_NBA_key")
	cfg.TelegramToken = os.Getenv("telegram_token")

	return cfg
}

func determineEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return "dev" // Default to dev environment
	}
	return env
}

func readConfigFile(env string) *Config {
	config := &Config{}
	name := fmt.Sprintf("./config/config_%s.json", env)
	f, err := os.Open(name)
	if err != nil {
		log.Printf("Cannot read %s config. Default values for config will be used.", name)
		return config
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		log.Printf("Cannot read data from %s config name. Default values for config will be used.", name)
		return config
	}

	err = json.Unmarshal(data, config)
	if err != nil {
		log.Printf("Cannot unmarshal data from %s config name. Default values for config will be used.", name)
	}

	return config
}
