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
		os.Getenv("UA_TO_EN_TG_BOT_TOKEN"),
		"Привіт. Я перекладаю всі повідомлення з української на англійську мову.",
		"Сталася помилка під час перекладу повідомлення.",
		translator.New(
			"UK",
			"EN",
			&client.DeeplClient{},
		),
	)

	bot.PollTgApiAndRespond()
}
