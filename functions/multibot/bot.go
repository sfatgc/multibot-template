package multibot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"gopkg.in/telebot.v3"
)

type TgBotInterface interface {
	GetBotName() string
	GetTgBot() *telebot.Bot
	GetTgWebhook() *telebot.Webhook
	Handle(c telebot.Context) error
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	CheckWebhook() bool
}

type TgBot struct {
	BotName   string
	Error     error
	Verbose   bool
	TgWebhook telebot.Webhook
	TgBot     *telebot.Bot
}

func (b *TgBot) GetBotName() string {
	return b.BotName
}

func (b *TgBot) GetTgBot() *telebot.Bot {
	return b.TgBot
}

func (b *TgBot) GetTgWebhook() *telebot.Webhook {
	return &b.TgWebhook
}

func (b *TgBot) Handle(c telebot.Context) error {
	message := c.Text()
	response := fmt.Sprintf("Abstract bot says: I am %s\n You asked for %s", b.BotName, message)
	_, err := c.Bot().Send(c.Sender(), response)
	return err
}

func (b *TgBot) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var update telebot.Update

	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Panicf("Cannot decode request body to telebot.Update struct. Quitting.")
	}

	b.TgBot.ProcessUpdate(update)

}

func (b TgBot) CheckWebhook() bool {

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
			if b.Verbose {
				log.Printf("Webhook check: Bot %s webhook secret (%s) doesn't match desired value (%s)",
					b.BotName,
					wh.SecretToken,
					b.TgWebhook.SecretToken)
			}
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

func NewBot(bot_name string, bot_telegram_token string, bot_webhook_secret string, bot_webhook_url string) (TgBotInterface, error) {

	tbs := telebot.Settings{Token: bot_telegram_token, Verbose: true}

	tele_bot, err := telebot.NewBot(tbs)
	if err != nil {
		return nil, err
	}

	var bot TgBotInterface

	switch bot_name {
	case "BULGAKTEER":
		bot = &TgBotBulgakteer{
			TgBot{bot_name,
				err,
				tbs.Verbose,
				telebot.Webhook{MaxConnections: 5, SecretToken: bot_webhook_secret, Endpoint: &telebot.WebhookEndpoint{PublicURL: bot_webhook_url}},
				tele_bot},
		}
	case "SFATGC":
		bot = &TgBotSFATGC{
			TgBot{bot_name,
				err,
				tbs.Verbose,
				telebot.Webhook{MaxConnections: 5, SecretToken: bot_webhook_secret, Endpoint: &telebot.WebhookEndpoint{PublicURL: bot_webhook_url}},
				tele_bot},
		}
	default:
		bot = &TgBot{
			bot_name,
			err,
			tbs.Verbose,
			telebot.Webhook{MaxConnections: 5, SecretToken: bot_webhook_secret, Endpoint: &telebot.WebhookEndpoint{PublicURL: bot_webhook_url}},
			tele_bot,
		}
	}

	bot.GetTgBot().Handle(telebot.OnText, bot.Handle)

	if !bot.CheckWebhook() {
		err := bot.GetTgBot().SetWebhook(bot.GetTgWebhook())

		if err != nil {
			return bot, err
		}
	}

	return bot, nil
}
