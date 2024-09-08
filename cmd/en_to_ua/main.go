package main

import (
	"github.com/joho/godotenv"
	"iliayatsenko1708/ru_to_ua_trans_tg_bot/internal/tgbot"
	"iliayatsenko1708/ru_to_ua_trans_tg_bot/internal/translator"
	"iliayatsenko1708/ru_to_ua_trans_tg_bot/internal/translator/client"
	"os"
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
