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
