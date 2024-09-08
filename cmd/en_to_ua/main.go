package main

import (
	"github.com/joho/godotenv"
	"os"
	"tg_translate_bots/internal/tgbot"
	"tg_translate_bots/internal/translator"
	"tg_translate_bots/internal/translator/client"
)

func main() {
	_ = godotenv.Load(".env")

	bot := tgbot.New(
		os.Getenv("EN_TO_UA_TG_BOT_TOKEN"),
		"Hello. I translate all messages from English to Ukrainian.",
		"An error occurred while translating the message.",
		translator.New(
			"EN",
			"UK",
			&client.DeeplClient{},
		),
	)

	bot.PollTgApiAndRespond()
}
