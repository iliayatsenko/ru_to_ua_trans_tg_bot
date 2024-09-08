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
		os.Getenv("RU_TO_UA_TG_BOT_TOKEN"),
		"Привет. Я перевожу все сообщения с русского на украинский язык.",
		"Произошла ошибка при переводе сообщения.",
		translator.New(
			"RU",
			"UK",
			&client.DeeplClient{},
		),
	)

	bot.PollTgApiAndRespond()
}
