package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
	"tg_translate_bots/internal/translator"
	"time"
)

const inlineResponseTimeout = 750 * time.Millisecond

var staticKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("/help"),
		tgbotapi.NewKeyboardButton("/discover"),
	),
)

type TgBot struct {
	Name                  string
	Link                  string
	token                 string
	greeting              string
	errorMsg              string
	translator            *translator.Translator
	otherBotsDiscoverFunc func() map[string]string
	inlineResponsesTimers map[string]*time.Timer
}

func New(
	name,
	link,
	token,
	greeting,
	errorMsg string,
	translator *translator.Translator,
	otherBotsDiscoverFunc func() map[string]string,
) *TgBot {
	return &TgBot{
		Name:                  name,
		Link:                  link,
		token:                 token,
		greeting:              greeting,
		errorMsg:              errorMsg,
		translator:            translator,
		otherBotsDiscoverFunc: otherBotsDiscoverFunc,
		inlineResponsesTimers: make(map[string]*time.Timer),
	}
}

func (tb *TgBot) PollTgApiAndRespond() {
	bot, err := tgbotapi.NewBotAPI(tb.token)
	if err != nil {
		log.Println(err)
		return
	}

	bot.Debug, _ = strconv.ParseBool(os.Getenv("BOT_DEBUG"))

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot is up to.
		if update.InlineQuery != nil {
			tb.handleInlineQuery(update, bot)
		} else if update.Message != nil {
			if update.Message.IsCommand() {
				tb.handleCommand(update, bot)
			} else {
				tb.handleMessage(update, bot)
			}
		}
	}
}

func (tb *TgBot) handleMessage(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	replyMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	replyMsg.ReplyMarkup = staticKeyboard

	translated, err := tb.translator.Translate(update.Message.Text)
	if err != nil {
		log.Println(err)
		replyMsg.Text = tb.errorMsg
	} else {
		replyMsg.Text = translated
	}
	replyMsg.ReplyToMessageID = update.Message.MessageID

	if _, err := bot.Send(replyMsg); err != nil {
		log.Println(err)
	}
}

func (tb *TgBot) handleCommand(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	replyMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
	replyMsg.ReplyMarkup = staticKeyboard

	switch update.Message.Command() {
	case "discover":
		otherBots := tb.otherBotsDiscoverFunc()

		otherBotButtons := [][]tgbotapi.InlineKeyboardButton{}
		for name, link := range otherBots {
			if name == tb.Name {
				continue
			}
			otherBotButtons = append(otherBotButtons, tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL(name, link)))
		}

		var otherBotsKeyboard = tgbotapi.NewInlineKeyboardMarkup(
			otherBotButtons...,
		)
		replyMsg.ReplyMarkup = otherBotsKeyboard
		replyMsg.Text = "\xF0\x9F\x93\x98" // blue book emoji
	case "help":
		fallthrough
	case "start":
		fallthrough
	default:
		replyMsg.Text = tb.greeting
	}

	if _, err := bot.Send(replyMsg); err != nil {
		log.Println(err)
	}
}

func (tb *TgBot) handleInlineQuery(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	inlineResponseTimer := time.AfterFunc(inlineResponseTimeout, func() {
		var replyText string
		translated, err := tb.translator.Translate(update.InlineQuery.Query)
		if err != nil {
			log.Println(err)
			replyText = tb.errorMsg
		} else {
			replyText = translated
		}
		inlineQueryResponse := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, replyText, replyText)
		inline := tgbotapi.InlineConfig{
			InlineQueryID: update.InlineQuery.ID,
			Results:       []interface{}{inlineQueryResponse},
		}

		if _, err := bot.Send(inline); err != nil {
			log.Println(err)
		}
	})

	prevTimer, ok := tb.inlineResponsesTimers[update.InlineQuery.ID]
	if ok {
		prevTimer.Stop()
		delete(tb.inlineResponsesTimers, update.InlineQuery.ID)
	}

	tb.inlineResponsesTimers[update.InlineQuery.ID] = inlineResponseTimer
}
