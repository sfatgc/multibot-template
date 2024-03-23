package multibot

import (
	"gopkg.in/telebot.v3"
)

type TgBot struct {
	BotName   string
	Error     error
	TgWebhook telebot.Webhook
	TgBot     *telebot.Bot
}

func NewBot(name string, token string, secret string) (*TgBot, error) {

	tbs := telebot.Settings{Token: token}

	tele_bot, err := telebot.NewBot(tbs)
	if err != nil {
		return nil, err
	}

	tele_bot.Handle(telebot.OnText, func(c telebot.Context) error {
		message := c.Text()
		_, err := tele_bot.Send(c.Sender(), message)
		return err
	})

	bot := TgBot{
		name,
		err,
		telebot.Webhook{MaxConnections: 5, SecretToken: token},
		tele_bot,
	}

	return &bot, nil
}
