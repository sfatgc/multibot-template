package multibot

import (
	"reflect"

	"gopkg.in/telebot.v3"
)

type TgBot struct {
	BotName   string
	Error     error
	TgWebhook telebot.Webhook
	TgBot     *telebot.Bot
}

func NewBot(bot_name string, bot_telegram_token string, bot_webhook_secret string, bot_webhook_url string) (*TgBot, error) {

	tbs := telebot.Settings{Token: bot_telegram_token}

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
		bot_name,
		err,
		telebot.Webhook{MaxConnections: 5, SecretToken: bot_telegram_token, Endpoint: &telebot.WebhookEndpoint{PublicURL: bot_webhook_url}},
		tele_bot,
	}

	if !bot.CheckWebhook() {
		err := bot.TgBot.SetWebhook(&bot.TgWebhook)

		if err != nil {
			return &bot, err
		}
	}

	return &bot, nil
}

func (b *TgBot) CheckWebhook() bool {
	tb := b.TgBot
	wh, err := tb.Webhook()

	if err == nil {
		if wh.Endpoint.PublicURL != b.TgWebhook.Endpoint.PublicURL {
			return false
		}
		if wh.SecretToken != b.TgWebhook.SecretToken {
			return false
		}
		if wh.MaxConnections != b.TgWebhook.MaxConnections {
			return false
		}
		if !reflect.DeepEqual(wh.AllowedUpdates, b.TgWebhook.AllowedUpdates) {
			return false
		}

		return true
	}

	return false
}
