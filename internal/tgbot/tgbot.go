package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"iliayatsenko1708/ru_to_ua_trans_tg_bot/internal/translator"
	"log"
)

var staticKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/help"),
		tgbotapi.NewKeyboardButton("/otherBots"),
	),
)

type TgBot struct {
	token      string
	greeting   string
	errorMsg   string
	translator *translator.Translator
}

func New(token, greeting, errorMsg string, translator *translator.Translator) *TgBot {
	return &TgBot{
		token:      token,
		greeting:   greeting,
		errorMsg:   errorMsg,
		translator: translator,
	}
}

func (t *TgBot) PollTgApiAndRespond() {
	bot, err := tgbotapi.NewBotAPI(t.token)
	if err != nil {
		log.Println(err)
		return
	}

	bot.Debug = true

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		replyMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		replyMsg.ReplyMarkup = staticKeyboard

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "otherBots":
				var otherBotsKeyboard = tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonURL("Ru-Ua", "tg://resolve?domain=@RuToUaTranslatorBot"),
						tgbotapi.NewInlineKeyboardButtonURL("Ua-Ru", "tg://resolve?domain=@UaToRuTranslatorBot"),
						tgbotapi.NewInlineKeyboardButtonURL("En-Ru", "tg://resolve?domain=@EnToRuTranslatorBot"),
						tgbotapi.NewInlineKeyboardButtonURL("Ru-En", "tg://resolve?domain=@RuToEnTranslatorBot"),
						tgbotapi.NewInlineKeyboardButtonURL("En-Ua", "tg://resolve?domain=@EnToUaTranslatorBot"),
						tgbotapi.NewInlineKeyboardButtonURL("Ua-En", "tg://resolve?domain=@UaToEnTranslatorBot"),
					),
				)
				replyMsg.ReplyMarkup = otherBotsKeyboard
				replyMsg.Text = "\xF0\x9F\x93\x98" // blue book emoji
			case "help":
				fallthrough
			case "start":
				fallthrough
			default:
				replyMsg.Text = t.greeting
			}
		} else {
			translated, err := t.translator.Translate(update.Message.Text)
			if err != nil {
				log.Println(err)
				replyMsg.Text = "Произошла ошибка при переводе сообщения."
			} else {
				replyMsg.Text = translated
			}
			replyMsg.ReplyToMessageID = update.Message.MessageID
		}

		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		if _, err := bot.Send(replyMsg); err != nil {
			log.Println(err)
		}
	}
}
