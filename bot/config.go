package bot

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Config struct {
	TelegramKey string `yaml:"telegram_key"`
	OwnerId     int64  `yaml:"owner_id"`
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

func initConfig() *Config {
	c := &Config{}
	c.load("data/config.yaml")
	return c
}
