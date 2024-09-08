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
