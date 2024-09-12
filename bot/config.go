package bot

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	TelegramKey string `yaml:"telegram_key"`
}

// Load method loads configuration file to Config struct
func (c *Config) load(configFile string) {
	file, err := os.Open(configFile)

	if err != nil {
		log.Printf("Error while opening config file: %v", err)
	}

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&c)

	if err != nil {
		log.Printf("Error while decoding JSON: %v", err)
	}
}

func initConfig() *Config {
	c := &Config{}
	c.load("config.yaml")
	return c
}
