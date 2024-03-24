package multibot

import (
	"log"
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

	tbs := telebot.Settings{Token: bot_telegram_token, Verbose: true}

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
		telebot.Webhook{MaxConnections: 5, SecretToken: bot_webhook_secret, Endpoint: &telebot.WebhookEndpoint{PublicURL: bot_webhook_url}},
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

	result := true

	tb := b.TgBot
	wh, err := tb.Webhook()

	if err == nil {

		log.Printf("Webhook check: Bot %s webhook exists", b.BotName)

		if wh.Endpoint == nil || b.TgWebhook.Endpoint == nil {
			result = false
		} else if wh.Endpoint.PublicURL != b.TgWebhook.Endpoint.PublicURL {
			log.Printf("Webhook check: Bot %s webhook URL (%s) doesn't match desired value (%s)",
				b.BotName,
				wh.Endpoint.PublicURL,
				b.TgWebhook.Endpoint.PublicURL)
			result = false
		}

		// TODO: Remove secret value logging after debug
		if wh.SecretToken != b.TgWebhook.SecretToken {
			log.Printf("Webhook check: Bot %s webhook secret (%s) doesn't match desired value (%s)",
				b.BotName,
				wh.SecretToken,
				b.TgWebhook.SecretToken)
			result = false
		}

		if wh.MaxConnections != b.TgWebhook.MaxConnections {
			log.Printf("Webhook check: Bot %s webhook MaxConnections (%d) doesn't match desired value (%d)",
				b.BotName,
				wh.MaxConnections,
				b.TgWebhook.MaxConnections)
			result = false
		}

		if !reflect.DeepEqual(wh.AllowedUpdates, b.TgWebhook.AllowedUpdates) {
			log.Printf("Webhook check: Bot %s webhook AllowedUpdates (%v) doesn't match desired value (%v)",
				b.BotName,
				wh.AllowedUpdates,
				b.TgWebhook.AllowedUpdates)
			result = false
		}

	} else {
		log.Printf("Webhook check: Bot %s webhook is not set", b.BotName)
		result = false
	}

	return result
}
