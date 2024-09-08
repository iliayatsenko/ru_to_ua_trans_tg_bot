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
		os.Getenv("RU_TO_EN_TG_BOT_TOKEN"),
		"Привет. Я перевожу все сообщения с русского на английский язык.",
		"Произошла ошибка при переводе сообщения.",
		translator.New(
			"RU",
			"EN",
			&client.DeeplClient{},
		),
	)

	bot.PollTgApiAndRespond()
}
