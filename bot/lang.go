package bot

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type Language struct {
	Test string `yaml:"test"`
}

// Load method loads configuration file to Config struct
func (l *Language) load(langFile string) {
	file, err := os.Open(langFile)

	if err != nil {
		log.Printf("Error while opening lant file: %v", err)
	}

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&l)

	if err != nil {
		log.Printf("Error while decoding YAML: %v", err)
	}
}

func initLang() *Language {
	l := &Language{}
	l.load("langs/en.yaml")
	return l
}
