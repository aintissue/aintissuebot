package main

import (
	"log"

	"github.com/aintissue/aintissuebot/bot"
)

// var conf *Config

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// conf = initConfig()

	log.Println("Test.")

	bot.Run()
}
