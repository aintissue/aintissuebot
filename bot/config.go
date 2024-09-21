package bot

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

// Config stores configuration vars
type Config struct {
	TelegramKey string `yaml:"telegram_key"`
	OwnerId     int64  `yaml:"owner_id"`
	Dev         bool   `yaml:"dev"`
}

// Load method loads configuration file to Config struct
func (c *Config) load(configFile string) {
	file, err := os.Open(configFile)

	if err != nil {
		loge(err)
	}

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&c)

	if err != nil {
		loge(err)
	}
}

// Initializes configuration
func initConfig() *Config {
	c := &Config{}
	c.load("data/config.yaml")
	return c
}
