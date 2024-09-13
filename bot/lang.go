package bot

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Language struct {
	StartMsg string `yaml:"startMsg"`
}

// Load method loads configuration file to Config struct
func (l *Language) load(langFile string) {
	file, err := os.Open(langFile)

	if err != nil {
		loge(err)
	}

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&l)

	if err != nil {
		loge(err)
	}
}

func initLang() *Language {
	l := &Language{}
	l.load("langs/en.yaml")
	return l
}
