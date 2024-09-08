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
		os.Getenv("UA_TO_RU_TG_BOT_TOKEN"),
		"Привіт. Я перекладаю всі повідомлення з української на російську мову.",
		"Сталася помилка під час перекладу повідомлення.",
		translator.New(
			"UK",
			"RU",
			&client.DeeplClient{},
		),
	)

	bot.PollTgApiAndRespond()
}
