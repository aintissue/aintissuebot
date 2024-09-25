package bot

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

// Language stores language variables
type Language struct {
	StartMsg      string `yaml:"startMsg"`
	RegisterDone  string `yaml:"registerDone"`
	RegisterError string `yaml:"registerError"`
	LinkMsg       string `yaml:"linkMsg"`
}

// Load method loads configuration file to Language struct
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

// Initializes language
func initLang() *Language {
	l := &Language{}
	l.load(LangPath)
	return l
}
