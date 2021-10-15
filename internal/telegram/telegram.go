package telegram

import (
	"fmt"
	"homework/internal/app/repositories"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	bot          *tgbotapi.BotAPI
	repositories *repositories.UserRepository
}

func NewBot(bot *tgbotapi.BotAPI, repositories *repositories.UserRepository) *Bot {
	return &Bot{bot: bot, repositories: repositories}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates, err := b.initUpdatesChannel()
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}
	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
			continue
		}
		if i, err := strconv.Atoi(fmt.Sprint(update.Message)); err == nil {
			b.repositories.FindById(i)
			continue
		}
		b.handleMessage(update.Message)
	}
}

func (b *Bot) initUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}
