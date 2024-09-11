package main

import (
	"log"

	"aintissuebot/bot"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	bot.Run()
}
