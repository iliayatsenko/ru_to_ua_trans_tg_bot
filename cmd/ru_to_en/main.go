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
