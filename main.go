package main

import (
	"github.com/joho/godotenv"
	"os"
	"sync"
	"tg_translate_bots/internal/tgbot"
	"tg_translate_bots/internal/translator"
	"tg_translate_bots/internal/translator/client"
)

func main() {
	_ = godotenv.Load(".env")

	var botsPtr *[]*tgbot.TgBot
	botsDiscoverFunc := func() map[string]string {
		botsMap := make(map[string]string)
		for _, bot := range *botsPtr {
			botsMap[bot.Name] = bot.Link
		}
		return botsMap
	}

	var bots = []*tgbot.TgBot{
		tgbot.New(
			"En-Ru",
			"https://t.me/EnToRuTranslatorBot",
			os.Getenv("EN_TO_RU_TG_BOT_TOKEN"),
			"Hello. I translate all messages from English to Russian.",
			"An error occurred while translating the message.",
			translator.New(
				"EN",
				"RU",
				&client.DeeplClient{},
			),
			botsDiscoverFunc,
		),
		tgbot.New(
			"En-Ua",
			"https://t.me/EnToUaTranslatorBot",
			os.Getenv("EN_TO_UA_TG_BOT_TOKEN"),
			"Hello. I translate all messages from English to Ukrainian.",
			"An error occurred while translating the message.",
			translator.New(
				"EN",
				"UK",
				&client.DeeplClient{},
			),
			botsDiscoverFunc,
		),
		tgbot.New(
			"Ru-En",
			"https://t.me/RuToEnTranslatorBot",
			os.Getenv("RU_TO_EN_TG_BOT_TOKEN"),
			"Привет. Я перевожу все сообщения с русского на английский язык.",
			"Произошла ошибка при переводе сообщения.",
			translator.New(
				"RU",
				"EN",
				&client.DeeplClient{},
			),
			botsDiscoverFunc,
		),
		tgbot.New(
			"Ru-Ua",
			"https://t.me/RuToUaTranslatorBot",
			os.Getenv("RU_TO_UA_TG_BOT_TOKEN"),
			"Привет. Я перевожу все сообщения с русского на украинский язык.",
			"Произошла ошибка при переводе сообщения.",
			translator.New(
				"RU",
				"UK",
				&client.DeeplClient{},
			),
			botsDiscoverFunc,
		),
		tgbot.New(
			"Ua-En",
			"https://t.me/UaToEnTranslatorBot",
			os.Getenv("UA_TO_EN_TG_BOT_TOKEN"),
			"Привіт. Я перекладаю всі повідомлення з української на англійську мову.",
			"Сталася помилка під час перекладу повідомлення.",
			translator.New(
				"UK",
				"EN",
				&client.DeeplClient{},
			),
			botsDiscoverFunc,
		),
		tgbot.New(
			"Ua-Ru",
			"https://t.me/UaToRuTranslatorBot",
			os.Getenv("UA_TO_RU_TG_BOT_TOKEN"),
			"Привіт. Я перекладаю всі повідомлення з української на російську мову.",
			"Сталася помилка під час перекладу повідомлення.",
			translator.New(
				"UK",
				"RU",
				&client.DeeplClient{},
			),
			botsDiscoverFunc,
		),
	}

	botsPtr = &bots

	var wg sync.WaitGroup
	wg.Add(len(bots))

	for _, bot := range bots {
		go func(bot *tgbot.TgBot) {
			defer wg.Done()
			bot.PollTgApiAndRespond()
		}(bot)
	}

	wg.Wait()
}
