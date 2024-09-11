package main

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct holds all our configuration
type Config struct {
	TelegramApiKey string `json:"telegram_api_key"`
}

// Load method loads configuration file to Config struct
func (c *Config) load(configFile string) {
	file, err := os.Open(configFile)

	if err != nil {
		log.Printf("Error while opening config file: %v", err)
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(&c)

	if err != nil {
		log.Printf("Error while decoding JSON: %v", err)
	}
}

func initConfig() *Config {
	c := &Config{}
	c.load("config.json")
	return c
}
