package telegram

import (
	"encoding/json"
	"homework/internal/app/models"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "start"
	commandList  = "list"
	// commandListUser = "list/{id:[0-9]+}"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	b.bot.Send(msg)

}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Я не знаю такой команды")
	switch message.Command() {
	case commandStart:
		msg.Text = "Ты ввел комманду СТАРТ"
		_, err := b.bot.Send(msg)
		return err
	case commandList:
		list := b.repositories.ListUser()
		bytes, err := json.Marshal(list)
		if err != nil {
			log.Fatalf("ERROR %s", err)
		}

		msg.Text = string(bytes)
		_, err = b.bot.Send(msg)
		return err
	// case commandListUser:
	// 	msg.Text = fmt.Sprint(b.repositories.FindById())
	// 	_, err := b.bot.Send(msg)
	// 	return err
	default:
		_, err := b.bot.Send(msg)
		return err
	}
}

func (b *Bot) UnmarshalBody(message *tgbotapi.Message) bool {

	type botText models.User

	if err := json.Unmarshal([]byte(message.Text), &botText); err != nil {
		return false
	}
	return true
}
